---
organization: Turbot
category: ["public cloud"]
icon_url: "/images/plugins/turbot/azure.svg"
brand_color: "#0089D6"
display_name: "Azure"
description: "Tailpipe plugin for collecting and querying various logs from Azure."
og_description: "Collect Azure logs and query them instantly with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/azure-social-graphic.png"
---

# Azure + Tailpipe

[Tailpipe](https://tailpipe.io) is an open-source CLI tool that allows you to collect logs and query them with SQL.

[Azure](https://azure.microsoft.com) provides on-demand cloud computing platforms and APIs to authenticated customers on a metered pay-as-you-go basis.

The [Azure Plugin for Tailpipe](https://hub.tailpipe.io/plugins/turbot/azure) allows you to collect and query Azure logs using SQL to track activity, monitor trends, detect anomalies, and more!

- Documentation: [Table definitions & examples](https://hub.tailpipe.io/plugins/turbot/azure/tables)
- Community: [Join #tailpipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/tailpipe-plugin-azure/issues)

<img src="https://raw.githubusercontent.com/turbot/tailpipe-plugin-azure/main/docs/images/azure_activity_log_terminal.png" width="50%" type="thumbnail"/>
<img src="https://raw.githubusercontent.com/turbot/tailpipe-plugin-azure/main/docs/images/azure_activity_log_mitre_dashboard.png" width="50%" type="thumbnail"/>

## Getting Started

Install Tailpipe from the [downloads](https://tailpipe.io/downloads) page:

```sh
# MacOS
brew install turbot/tap/tailpipe
```

```sh
# Linux or Windows (WSL)
sudo /bin/sh -c "$(curl -fsSL https://tailpipe.io/install/tailpipe.sh)"
```

Install the plugin:

```sh
tailpipe plugin install azure
```

Configure your [connection credentials](https://hub.tailpipe.io/plugins/turbot/azure#connection-credentials), table partition, and data source ([examples](https://hub.tailpipe.io/plugins/turbot/azure/tables/azure_activity_log#example-configurations)):

```sh
vi ~/.tailpipe/config/azure.tpc
```

```hcl
connection "azure" "my_subscription" {
  tenant_id       = "00000000-0000-0000-0000-000000000000"
  subscription_id = "00000000-0000-0000-0000-000000000000"
  client_id       = "00000000-0000-0000-0000-000000000000"
  client_secret   = "my plaintext secret"
}

partition "azure_activity_log" "my_logs" {
  source "azure_blob_storage" {
    connection   = connection.azure.my_subscription
    account_name = "storage_account_name"
    container    = "container_name"
  }
}
```

Download, enrich, and save logs from your source ([examples](https://tailpipe.io/docs/reference/cli/collect)):

```sh
tailpipe collect azure_activity_log
```

Enter interactive query mode:

```sh
tailpipe query
```

Run a query:

```sql
select
  resource_type,
  operation_name,
  count(*) as operation_count
from
  azure_activity_log
group by
  resource_type,
  operation_name
order by
  operation_count desc;
```

```sh
+-----------------------------------------------------------+------------------------------------------------------------------+-----------------+
| resource_type                                             | operation_name                                                   | operation_count |
+-----------------------------------------------------------+------------------------------------------------------------------+-----------------+
| Microsoft.Resources/deployments                           | Microsoft.Resources/deployments/write                            | 86              |
| Microsoft.Resources/deployments                           | Microsoft.Resources/deployments/validate/action                  | 58              |
| Microsoft.Compute/virtualMachines                         | Microsoft.Authorization/policies/auditIfNotExists/action         | 54              |
| Microsoft.Compute/virtualMachines                         | Microsoft.Authorization/policies/audit/action                    | 36              |
| Microsoft.Sql/servers                                     | Microsoft.Authorization/policies/auditIfNotExists/action         | 25              |
| Microsoft.Sql/servers/databases                           | Microsoft.Sql/servers/databases/read                             | 20              |
| MICROSOFT.CDN/profiles                                    | Microsoft.Resourcehealth/healthevent/Activated/action            | 18              |
+-----------------------------------------------------------+------------------------------------------------------------------+-----------------+
```

## Detections as Code with Powerpipe

Pre-built dashboards and detections for the Azure plugin are available in [Powerpipe](https://powerpipe.io) mods, helping you monitor and analyze activity across your Azure subscriptions.

For example, the [Azure Activity Log Detections mod](https://hub.powerpipe.io/mods/turbot/tailpipe-mod-azure-activity-log-detections) scans your activity logs for anomalies, such as a SQL server firewall rule getting updated or a change in your virtual networks.

Dashboards and detections are [open source](https://github.com/topics/tailpipe-mod), allowing easy customization and collaboration.

To get started, choose a mod from the [Powerpipe Hub](https://hub.powerpipe.io/?engines=tailpipe&q=azure).

<img src="https://raw.githubusercontent.com/turbot/tailpipe-plugin-azure/main/docs/images/azure_activity_log_mitre_dashboard.png"/>

## Connection Credentials

### Arguments

| Name                   | Type   | Required | Description                                                                      |
|------------------------|--------|----------|----------------------------------------------------------------------------------|
| `certificate_password` | String | No       | The password for the certificate file, if required.                              |
| `certificate_path`     | String | No       | Path to the certificate file used for authentication instead of a client secret. |
| `client_id`            | String | Yes      | The client ID (Application ID) for authentication using a service principal.     |
| `client_secret`        | String | No       | The client secret for authentication using a service principal.                  |
| `environment`          | String | No       | The Azure environment (e.g., `AzurePublicCloud`, `AzureChinaCloud`).             |
| `password`             | String | No       | The password for authentication when using user-based authentication.            |
| `subscription_id`      | String | Yes      | The Azure subscription ID for resource access.                                   |
| `tenant_id`            | String | Yes      | The Entra ID tenant ID.                                                          |
| `username`             | String | No       | The username for authentication when using user-based authentication.            |

### Client Secret Credentials

You may specify the tenant ID, subscription ID, client ID, and client secret to authenticate:

- `tenant_id`: Specify the tenant to authenticate with.
- `subscription_id`: Specify the subscription to connect to.
- `client_id`: Specify the app client ID to use.
- `client_secret`: Specify the app secret to use.

#### azure.tpc:

```hcl
connection "azure" "azure_via_sp_secret" {
  plugin            = "azure"
  tenant_id         = "00000000-0000-0000-0000-000000000000"
  subscription_id   = "00000000-0000-0000-0000-000000000000"
  client_id         = "00000000-0000-0000-0000-000000000000"
  client_secret     = "my plaintext password"
}
```

### Client Certificate Credentials

You may specify the tenant ID, subscription ID, client ID, certificate path, and certificate password to authenticate:

- `tenant_id`: Specify the tenant to authenticate with.
- `subscription_id`: Specify the subscription to connect to.
- `client_id`: Specify the app client ID to use.
- `certificate_path`: Specify the certificate path to use.
- `certificate_password`: Specify the certificate password to use.

#### azure.tpc:

```hcl
connection "azure" "azure_via_sp_cert" {
  plugin               = "azure"
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
- `subscription_id`: Specify the subscription to connect to.
- `client_id`: Specify the app client ID to use.
- `username`: Specify the username to use.
- `password`: Specify the password to use.

#### azure.tpc:

```hcl
connection "azure" "password_not_recommended" {
  plugin          = "azure"
  tenant_id       = "00000000-0000-0000-0000-000000000000"
  subscription_id = "00000000-0000-0000-0000-000000000000"
  client_id       = "00000000-0000-0000-0000-000000000000"
  username        = "my-username"
  password        = "plaintext password"
}
```

### Azure Managed Identity

Tailpipe works with managed identities (formerly known as Managed Service Identity), provided it is running in Azure, e.g., on a VM. All configuration is handled by Azure. See [Azure Managed Identities](https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/overview) for more details.

- `tenant_id`: Specify the tenant to authenticate with.
- `subscription_id`: Specify the subscription to connect to.
- `client_id`: Specify the app client ID of managed identity to use.

#### azure.tpc:

```hcl
connection "azure" "azure_msi" {
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
