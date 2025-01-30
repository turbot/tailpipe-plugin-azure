package blob_storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"github.com/elastic/go-grok"

	typehelpers "github.com/turbot/go-kit/types"
	"github.com/turbot/pipe-fittings/v2/filter"
	"github.com/turbot/tailpipe-plugin-azure/config"
	"github.com/turbot/tailpipe-plugin-sdk/artifact_source"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
	"github.com/turbot/tailpipe-plugin-sdk/types"
)

const AzureBlobStorageSourceIdentifier = "azure_blob_storage"

// AzureBlobStorageSource is a [ArtifactSource] implementation that reads artifacts from an Azure Blob Storage container
type AzureBlobStorageSource struct {
	artifact_source.ArtifactSourceImpl[*AzureBlobStorageSourceConfig, *config.AzureConnection]

	client    *container.Client
	errorList []error
}

func (s *AzureBlobStorageSource) Init(ctx context.Context, params *row_source.RowSourceParams, opts ...row_source.RowSourceOption) error {
	// call base to parse config and apply options
	if err := s.ArtifactSourceImpl.Init(ctx, params, opts...); err != nil {
		return err
	}

	client, err := s.getClient(ctx, s.Config.Container)
	if err != nil {
		return fmt.Errorf("failed to get Azure Blob Storage client: %w", err)
	}
	s.client = client

	slog.Info("Initialized AzureBlobStorageSource", "account_name", s.Config.AccountName, "container", s.Config.Container)
	return nil
}

func (s *AzureBlobStorageSource) Identifier() string {
	return AzureBlobStorageSourceIdentifier
}

func (s *AzureBlobStorageSource) Close() error {
	// delete any temp files
	_ = os.RemoveAll(s.TempDir)
	return nil
}

func (s *AzureBlobStorageSource) DiscoverArtifacts(ctx context.Context) error {
	var prefix string
	layout := typehelpers.SafeString(s.Config.GetFileLayout())
	// if there are any optional segments, we expand them into all possible combinations
	optionalLayouts := artifact_source.ExpandPatternIntoOptionalAlternatives(layout)

	filterMap := make(map[string]*filter.SqlFilter)

	g := grok.New()
	// add any patterns defined in config
	err := g.AddPatterns(s.Config.GetPatterns())
	if err != nil {
		return fmt.Errorf("failed to add grok patterns: %w", err)
	}

	if s.Config.Prefix != nil {
		prefix = *s.Config.Prefix
		if !strings.HasSuffix(prefix, "/") {
			prefix = prefix + "/"
		}
		var newOptionalLayouts []string
		for _, l := range optionalLayouts {
			newOptionalLayouts = append(newOptionalLayouts, fmt.Sprintf("%s%s", prefix, l))
		}
		optionalLayouts = newOptionalLayouts
	}

	err = s.walk(ctx, prefix, optionalLayouts, filterMap, g)
	if err != nil {
		s.errorList = append(s.errorList, fmt.Errorf("error walking Azure Blob Storage: %w", err))
	}

	if len(s.errorList) > 0 {
		return errors.Join(s.errorList...)
	}

	return nil
}

func (s *AzureBlobStorageSource) DownloadArtifact(ctx context.Context, info *types.ArtifactInfo) error {
	blobClient := s.client.NewBlobClient(info.Name)

	localFilePath := path.Join(s.TempDir, info.Name)
	if err := os.MkdirAll(path.Dir(localFilePath), 0755); err != nil {
		return fmt.Errorf("failed to create directory for file, %w", err)
	}

	resp, err := blobClient.DownloadStream(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to download blob: %w", err)
	}
	defer resp.Body.Close()

	// Get the size of the object
	size := typehelpers.Int64Value(resp.ContentLength)

	outFile, err := os.Create(localFilePath)
	if err != nil {
		return fmt.Errorf("failed to create file, %w", err)
	}
	defer outFile.Close()

	if _, err = io.Copy(outFile, resp.Body); err != nil {
		return fmt.Errorf("failed to write file, %w", err)

	}

	downloadInfo := types.NewDownloadedArtifactInfo(info, localFilePath, size)

	return s.OnArtifactDownloaded(ctx, downloadInfo)
}

func (s *AzureBlobStorageSource) getClient(ctx context.Context, containerName string) (*container.Client, error) {
	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", s.Config.AccountName)

	sess, err := s.Connection.GetSession()
	if err != nil {
		return nil, fmt.Errorf("failed getting Azure Connection session: %w", err)
	}

	client, err := azblob.NewClient(serviceURL, sess.Credential, &azblob.ClientOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed creating Azure Blob Storage client: %w", err)
	}

	c := client.ServiceClient().NewContainerClient(containerName)

	return c, nil
}

func (s *AzureBlobStorageSource) walk(ctx context.Context, prefix string, layouts []string, filterMap map[string]*filter.SqlFilter, g *grok.Grok) error {
	opts := container.ListBlobsHierarchyOptions{
		Prefix: &prefix,
	}
	pager := s.client.NewListBlobsHierarchyPager("/", &opts)

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("error getting next page, %w", err)
		}

		// Directories
		for _, dir := range page.Segment.BlobPrefixes {
			err = s.WalkNode(ctx, *dir.Name, "", layouts, true, g, filterMap)
			if err != nil {
				if errors.Is(err, fs.SkipDir) {
					continue
				}
				return fmt.Errorf("error walking node: %w", err)
			}
			err = s.walk(ctx, *dir.Name, layouts, filterMap, g)
			if err != nil {
				s.errorList = append(s.errorList, err)
			}
		}

		// Files
		for _, obj := range page.Segment.BlobItems {
			err = s.WalkNode(ctx, *obj.Name, "", layouts, false, g, filterMap)
			if err != nil {
				s.errorList = append(s.errorList, fmt.Errorf("error parsing object %s, %w", *obj.Name, err))
			}
		}
	}

	return nil
}
