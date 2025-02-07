---
title: "Source: azure_blob_storage - Collect logs from Azure Blob Storage"
description: "Allows users to collect logs from Azure Blob Storage."
---

# Source: azure_blob_storage - Collect logs from Azure Blob Storage

Azure Blob Storage is a cloud storage solution available in Microsoft Azure. It is designed to store large amounts of unstructured data, such as logs, backups, and media files. Blob Storage is optimized for large-scale workloads and provides cost-effective, secure, and highly available storage.

Using this source, you can collect, filter, and analyze logs stored in Azure Blob Storage, enabling system monitoring, security investigations, and compliance reporting.

## Example Configurations

### Collect activity logs

Collect activity logs for all subscriptions.

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

### Collect logs for a single subscription

Collect logs for a specific subscription.

```hcl
partition "azure_activity_log" "my_logs_subscription" {
  source "azure_blob_storage" {
    connection   = connection.azure.my_subscription
    account_name = "storage_account_name"
    container    = "container_name"
    file_layout  = "/SUBSCRIPTIONS/12345678-1234-1234-1234-123456789012/y=%{YEAR:year}/m=%{MONTHNUM:month}/d=%{MONTHDAY:day}/h=%{HOUR:hour}/m=%{MINUTE:minute}/%{DATA:filename}.json"
  }
}
```

## Arguments

| Argument     | Type               | Required | Default                    | Description                                                                                                              |
|--------------|--------------------|----------|----------------------------|--------------------------------------------------------------------------------------------------------------------------|
| account_name | String             | Yes      |                            | The name of the Storage account to collect logs from.                                                                   |
| connection   | `connection.azure` | No       | `connection.azure.default` | The [Azure connection](https://hub.tailpipe.io/plugins/turbot/azure#connection-credentials) to use to connect to the Azure subscription. |
| container    | String             | Yes      |                            | The name of the Storage container where logs are stored.                                                                |
| file_layout  | String             | No       |                            | The Grok pattern that defines the log file structure.                                                                   |

### Table Defaults

The following tables define their own default values for certain source arguments:

- **[azure_activity_log](https://hub.tailpipe.io/plugins/turbot/azure/tables/azure_activity_log#azure_blob_storage)**
