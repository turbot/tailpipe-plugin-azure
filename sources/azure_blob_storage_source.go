package sources

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"

	"github.com/turbot/tailpipe-plugin-azure/config"
	"github.com/turbot/tailpipe-plugin-sdk/artifact_source"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
	"github.com/turbot/tailpipe-plugin-sdk/schema"
	"github.com/turbot/tailpipe-plugin-sdk/types"
)

const AzureBlobStorageSourceIdentifier = "azure_blob_storage"

// register the source from the package init function
func init() {
	row_source.RegisterRowSource[*AzureBlobStorageSource]()
}

// AzureBlobStorageSource is a [ArtifactSource] implementation that reads artifacts from an Azure Blob Storage container
type AzureBlobStorageSource struct {
	artifact_source.ArtifactSourceImpl[*AzureBlobStorageSourceConfig, *config.AzureConnection]

	Extensions types.ExtensionLookup
	client     *azblob.Client
}

func (s *AzureBlobStorageSource) Init(ctx context.Context, configData, connectionData types.ConfigData, opts ...row_source.RowSourceOption) error {
	// call base to parse config and apply options
	if err := s.ArtifactSourceImpl.Init(ctx, configData, connectionData, opts...); err != nil {
		return err
	}

	s.TmpDir = path.Join(artifact_source.BaseTmpDir, fmt.Sprintf("azure-blob-%s-%s", s.Config.AccountName, s.Config.Container))
	s.Extensions = types.NewExtensionLookup(s.Config.Extensions)

	client, err := s.getClient(ctx)
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
	return nil // s.client has no Close / Finalizer method
}

func (s *AzureBlobStorageSource) DiscoverArtifacts(ctx context.Context) error {
	containerClient := s.client.ServiceClient().NewContainerClient(s.Config.Container)
	opts := azblob.ListBlobsFlatOptions{
		Prefix: &s.Config.Prefix,
	}

	pager := containerClient.NewListBlobsFlatPager(&opts)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("failed to list blobs in container: %w", err)
		}

		for _, blob := range page.Segment.BlobItems {
			objPath := *blob.Name
			if s.Extensions.IsValid(objPath) {
				sourceEnrichmentFields := &schema.SourceEnrichment{
					CommonFields: schema.CommonFields{
						TpSourceLocation: &objPath,
						TpSourceName:     &s.Config.Container,
						TpSourceType:     AzureBlobStorageSourceIdentifier,
					},
				}

				info := &types.ArtifactInfo{Name: objPath, OriginalName: objPath, SourceEnrichment: sourceEnrichmentFields}

				if err = s.OnArtifactDiscovered(ctx, info); err != nil {
					// TODO: #error should we continue or fail?
					return fmt.Errorf("failed to notify observers of discovered artifact, %w", err)
				}
			}
		}

	}

	return nil
}

func (s *AzureBlobStorageSource) DownloadArtifact(ctx context.Context, info *types.ArtifactInfo) error {
	blobClient := s.client.ServiceClient().NewContainerClient(s.Config.Container).NewBlobClient(info.Name)

	slog.Error("Container name:", info.Name, "done")
	localFilePath := path.Join(s.TmpDir, info.Name)
	if err := os.MkdirAll(path.Dir(localFilePath), 0755); err != nil {
		return fmt.Errorf("failed to create directory for file, %w", err)
	}

	resp, err := blobClient.DownloadStream(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to download blob: %w", err)
	}
	defer resp.Body.Close()

	outFile, err := os.Create(localFilePath)
	if err != nil {
		return fmt.Errorf("failed to create file, %w", err)
	}
	defer outFile.Close()

	if _, err = io.Copy(outFile, resp.Body); err != nil {
		return fmt.Errorf("failed to write file, %w", err)

	}

	downloadInfo := &types.ArtifactInfo{Name: localFilePath, OriginalName: info.Name, SourceEnrichment: info.SourceEnrichment}

	return s.OnArtifactDownloaded(ctx, downloadInfo)
}

func (s *AzureBlobStorageSource) getClient(ctx context.Context) (*azblob.Client, error) {
	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", s.Config.AccountName)

	sess, err := s.Connection.GetSession()
	if err != nil {
		return nil, fmt.Errorf("failed getting Azure Connection session: %w", err)
	}

	client, err := azblob.NewClient(serviceURL, sess.Credential, &azblob.ClientOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed creating Azure Blob Storage client: %w", err)
	}

	return client, nil
}
