package config

import (
	"bytes"
	"crypto"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"slices"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	cloudPolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

const PluginName = "azure"

type AzureConnection struct {
	Environment         *string `hcl:"environment"`
	TenantId            *string `hcl:"tenant_id"`
	SubscriptionId      *string `hcl:"subscription_id"`
	ClientId            *string `hcl:"client_id"`
	ClientSecret        *string `hcl:"client_secret"`
	CertificatePath     *string `hcl:"certificate_path"`
	CertificatePassword *string `hcl:"certificate_password"`
	UserName            *string `hcl:"username"`
	Password            *string `hcl:"password"`
}

type AzureConnectionSession struct {
	Credential     azcore.TokenCredential
	SubscriptionID string
	TenantID       string
	ClientOptions  *policy.ClientOptions
}

func (c *AzureConnection) Validate() error {
	// if environment is set, ensure it's a valid value
	if c.Environment != nil {
		validEnvironments := []string{"AZUREPUBLICCLOUD", "AZUREUSGOVERNMENTCLOUD", "AZURECHINACLOUD", "AZUREGERMANCLOUD"}
		if !slices.Contains(validEnvironments, *c.Environment) {
			return fmt.Errorf("invalid environment: %s", *c.Environment)
		}
	}

	// if certificate path is set, ensure certificate password is also set
	if c.CertificatePath != nil && c.CertificatePassword == nil {
		return fmt.Errorf("certificate password is required when certificate path is set")
	}

	// if username is set, ensure password is also set
	if c.UserName != nil && c.Password == nil {
		return fmt.Errorf("password is required when username is set")
	}

	return nil
}

func (c *AzureConnection) Identifier() string {
	return PluginName
}

func (c *AzureConnection) GetSession() (*AzureConnectionSession, error) {
	var cred azcore.TokenCredential
	var cloudConfiguration cloud.Configuration
	var err error

	// environment should be defaulted if not set
	environment := c.getConfigOrEnv(c.Environment, "AZURE_ENVIRONMENT")
	tenantId := c.getConfigOrEnv(c.TenantId, auth.TenantID)
	subscriptionId := c.getConfigOrEnv(c.SubscriptionId, auth.SubscriptionID)
	clientId := c.getConfigOrEnv(c.ClientId, auth.ClientID)
	clientSecret := c.getConfigOrEnv(c.ClientSecret, auth.ClientSecret)
	certificatePath := c.getConfigOrEnv(c.CertificatePath, auth.CertificatePath)
	certificatePassword := c.getConfigOrEnv(c.CertificatePassword, auth.CertificatePassword)
	username := c.getConfigOrEnv(c.UserName, auth.Username)
	password := c.getConfigOrEnv(c.Password, auth.Password)

	switch environment {
	case "AZURECHINACLOUD":
		cloudConfiguration = cloud.AzureChina
	case "AZUREUSGOVERNMENTCLOUD":
		cloudConfiguration = cloud.AzureGovernment
	default:
		cloudConfiguration = cloud.AzurePublic
	}

	clientOptions := policy.ClientOptions{ClientOptions: cloudPolicy.ClientOptions{Cloud: cloudConfiguration}}

	// determine credential type
	switch {
	// client-secret auth
	case tenantId != "" && subscriptionId != "" && clientId != "" && clientSecret != "":
		cred, err = azidentity.NewClientSecretCredential(tenantId, clientId, clientSecret, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create client secret credential: %w", err)
		}
	// certificate auth
	case tenantId != "" && subscriptionId != "" && clientId != "" && certificatePath != "":
		var certs []*x509.Certificate
		var key crypto.PrivateKey
		var passBytes []byte = nil
		var certBytes []byte

		certBytes, err = os.ReadFile(certificatePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read certificate file: %w", err)
		}

		if certificatePassword != "" {
			passBytes = []byte(certificatePassword)
		}

		certs, key, err = azidentity.ParseCertificates(certBytes, passBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse certificate: %w", err)
		}

		cred, err = azidentity.NewClientCertificateCredential(tenantId, clientId, certs, key, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create client certificate credential: %w", err)
		}
	// username/password auth
	case tenantId != "" && subscriptionId != "" && clientId != "" && username != "" && password != "":
		cred, err = azidentity.NewUsernamePasswordCredential(tenantId, clientId, username, password, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create username/password credential: %w", err)
		}
	// managed identity auth
	case tenantId != "" && subscriptionId != "" && clientId != "":
		cred, err = azidentity.NewManagedIdentityCredential(&azidentity.ManagedIdentityCredentialOptions{
			ID: azidentity.ClientID(clientId),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create managed identity credential: %w", err)
		}
	// fallback to CLI auth
	default:
		cred, err = azidentity.NewAzureCLICredential(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create Azure CLI credential: %w", err)
		}
		subscriptionId, err = getSubscriptionIDFromCLINew()
		if err != nil {
			return nil, err
		}
	}

	return &AzureConnectionSession{
		Credential:     cred,
		SubscriptionID: subscriptionId,
		TenantID:       tenantId,
		ClientOptions:  &clientOptions,
	}, nil
}

func (c *AzureConnection) getConfigOrEnv(configValue *string, env string) string {
	if configValue != nil {
		return *configValue
	}

	return os.Getenv(env)
}

// getSubscriptionIDFromCLINew executes Azure CLI to get the subscription ID.
func getSubscriptionIDFromCLINew() (string, error) {
	const azureCLIPath = "AzureCLIPath"

	azureCLIDefaultPathWindows := fmt.Sprintf("%s\\Microsoft SDKs\\Azure\\CLI2\\wbin; %s\\Microsoft SDKs\\Azure\\CLI2\\wbin", os.Getenv("ProgramFiles(x86)"), os.Getenv("ProgramFiles"))
	const azureCLIDefaultPath = "/bin:/sbin:/usr/bin:/usr/local/bin"

	var cliCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cliCmd = exec.Command(fmt.Sprintf("%s\\system32\\cmd.exe", os.Getenv("windir")))
		cliCmd.Env = os.Environ()
		cliCmd.Env = append(cliCmd.Env, fmt.Sprintf("PATH=%s;%s", os.Getenv(azureCLIPath), azureCLIDefaultPathWindows))
		cliCmd.Args = append(cliCmd.Args, "/c", "az")
	} else {
		cliCmd = exec.Command("az")
		cliCmd.Env = os.Environ()
		cliCmd.Env = append(cliCmd.Env, fmt.Sprintf("PATH=%s:%s", os.Getenv(azureCLIPath), azureCLIDefaultPath))
	}
	cliCmd.Args = append(cliCmd.Args, "account", "show", "-o", "json")

	var stderr bytes.Buffer
	cliCmd.Stderr = &stderr

	output, err := cliCmd.Output()
	if err != nil {
		return "", fmt.Errorf("invoking Azure CLI failed with the following error: %v", err)
	}

	var accountResponse struct {
		SubscriptionID string `json:"id"`
	}
	err = json.Unmarshal(output, &accountResponse)
	if err != nil {
		return "", fmt.Errorf("error parsing JSON output: %v", err)
	}

	return accountResponse.SubscriptionID, nil
}
