---
organization: Turbot
category: ["public cloud"]
icon_url: "/images/plugins/turbot/azure.svg"
brand_color: "#0089D6"
display_name: "Azure"
name: "azure"
description: "Tailpipe plugin for obtaining and querying logs from Azure."
og_description: "Query Azure logs with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/azure-social-graphic.png"
engines: ["tailpipe"]
---

# Azure + Tailpipe

[Tailpipe](https://tailpipe.io) is an open-source CLI tool that allows you to obtain logs and query then with SQL.

[Azure](https://azure.microsoft.com) provides on-demand cloud computing platforms and APIs to authenticated customers on a metered pay-as-you-go basis.

<!-- TODO: Insert Example -->

## Documentation

- **[Table definitions & examples →](/plugins/turbot/azure/tables)**

## Get started

### Install

Download and install the latest Azure plugin:

```bash
tailpipe plugin install azure
```

### Credentials

| Item        | Description                                                                                                                                                                                                                                                         |
| ----------- |---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Credentials | Use the `az login` command to setup your [Azure Default Connection](https://docs.microsoft.com/en-us/cli/azure/authenticate-azure-cli).                                                                                                                             |
| Permissions | Assign the `Reader and Data Access` (if listing storage account keys) roles to your user or service principal in the subscription.                                                                                                                                  |
| Radius      | Each connection represents a single Azure subscription.                                                                                                                                                                                                             |
| Resolution  | 1. Credentials explicitly set in a tailpipe config file (`~/.tailpipe/config/azure.tpc`).<br />2. Credentials specified in [environment variables](#credentials-from-environment-variables), e.g., `AZURE_SUBSCRIPTION_ID`.<br />3. Credentials from the Azure CLI. |

### Configuration

TODO: Elaborate on the configuration requirements around `partition`, `source` and `connection`.

## Configuring Azure Credentials

The Azure plugin support multiple formats/authentication mechanisms and they are tried in the below order:

1. [Client Secret Credentials](https://docs.microsoft.com/en-us/azure/active-directory/develop/v2-saml-bearer-assertion#prerequisites) if set; otherwise
2. [Client Certificate Credentials](https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-certificate-credentials#register-your-certificate-with-microsoft-identity-platform) if set; otherwise
3. [Resource Owner Password](https://docs.microsoft.com/en-us/azure/active-directory/develop/v2-oauth-ropc) if set; otherwise
4. If no credentials are supplied, then the [az cli](https://docs.microsoft.com/en-us/cli/azure/#:~:text=The%20Azure%20command%2Dline%20interface,with%20an%20emphasis%20on%20automation.) credentials are used

If connection arguments are provided, they will always take precedence over [Azure SDK environment variables](https://github.com/Azure/azure-sdk-for-go/blob/main/documentation/new-version-quickstart.md#setting-environment-variables), and they are tried in the below order:

### Client Secret Credentials

You may specify the tenant ID, subscription ID, client ID, and client secret to authenticate:

- `tenant_id`: Specify the tenant to authenticate with.
- `subscription_id`: Specify the subscription to query.
- `client_id`: Specify the app client ID to use.
- `client_secret`: Specify the app secret to use.

```hcl
connection "azure" "via_sp_secret" {
  tenant_id         = "00000000-0000-0000-0000-000000000000"
  subscription_id   = "00000000-0000-0000-0000-000000000000"
  client_id         = "00000000-0000-0000-0000-000000000000"
  client_secret     = "my plaintext password"
}
```

### Client Certificate Credentials

You may specify the tenant ID, subscription ID, client ID, certificate path, and certificate password to authenticate:

- `tenant_id`: Specify the tenant to authenticate with.
- `subscription_id`: Specify the subscription to query.
- `client_id`: Specify the app client ID to use.
- `certificate_path`: Specify the certificate path to use.
- `certificate_password`: Specify the certificate password to use.

```hcl
connection "azure" "via_sp_cert" {
  tenant_id            = "00000000-0000-0000-0000-000000000000"
  subscription_id      = "00000000-0000-0000-0000-000000000000"
  client_id            = "00000000-0000-0000-0000-000000000000"
  certificate_path     = "path/to/file.pem"
  certificate_password = "my plaintext password"
}
```

### Resource Owner Password

**Note:** This grant type is _not recommended_, use device login instead if you need interactive login.

You may specify the tenant ID, subscription ID, client ID, username, and password to authenticate:

- `tenant_id`: Specify the tenant to authenticate with.
- `subscription_id`: Specify the subscription to query.
- `client_id`: Specify the app client ID to use.
- `username`: Specify the username to use.
- `password`: Specify the password to use.

```hcl
connection "azure" "password_not_recommended" {
  tenant_id       = "00000000-0000-0000-0000-000000000000"
  subscription_id = "00000000-0000-0000-0000-000000000000"
  client_id       = "00000000-0000-0000-0000-000000000000"
  username        = "my-username"
  password        = "plaintext password"
}
```

### Azure Managed Identity

Steampipe works with managed identities (formerly known as Managed Service Identity), provided it is running in Azure, e.g., on a VM. All configuration is handled by Azure. See [Azure Managed Identities](https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/overview) for more details.

- `tenant_id`: Specify the tenant to authenticate with.
- `subscription_id`: Specify the subscription to query.
- `client_id`: Specify the app client ID of managed identity to use.

```hcl
connection "azure" "msi" {
  plugin          = "azure"
  tenant_id       = "00000000-0000-0000-0000-000000000000"
  client_id       = "00000000-0000-0000-0000-000000000000"
  subscription_id = "00000000-0000-0000-0000-000000000000"
}
```

### Azure CLI

If no credentials are specified and the SDK environment variables are not set, the plugin will use the active credentials from the Azure CLI. You can run `az login` to set up these credentials.


### Credentials from Environment Variables

The Azure AD plugin will use the standard Azure environment variables to obtain credentials **only if other arguments (`tenant_id`, `client_id`, `client_secret`, `certificate_path`, etc..) are not specified** in the connection:

```sh
export AZURE_ENVIRONMENT="AZUREPUBLICCLOUD" # Defaults to "AZUREPUBLICCLOUD". Valid environments are "AZUREPUBLICCLOUD", "AZURECHINACLOUD" and "AZUREUSGOVERNMENTCLOUD"
export AZURE_TENANT_ID="00000000-0000-0000-0000-000000000000"
export AZURE_SUBSCRIPTION_ID="00000000-0000-0000-0000-000000000000"
export AZURE_CLIENT_ID="00000000-0000-0000-0000-000000000000"
export AZURE_CLIENT_SECRET="my plaintext secret"
export AZURE_CERTIFICATE_PATH="path/to/file.pem"
export AZURE_CERTIFICATE_PASSWORD="my plaintext password"
```
