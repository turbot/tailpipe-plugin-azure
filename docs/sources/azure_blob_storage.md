---
title: "Source: azure_blob_storage - Collect logs from Azure Blob Storage"
description: "Allows users to collect logs from Azure Blob Storage."
---

# Source: azure_blob_storage - Collect logs from Azure Blob Storage

Azure Blob Storage is a cloud storage solution available in Microsoft Azure. It is designed to store large amounts of unstructured data, such as logs, backups, and media files. Blob Storage is optimized for large-scale workloads and provides cost-effective, secure, and highly available storage.

Using this source, you can collect, filter, and analyze logs stored in Azure Blob Storage, enabling system monitoring, security investigations, and compliance reporting.

## Example Configurations

### Collect Logs from an Azure Blob Storage Container

```hcl
connection "azure" "logging_account" {
  account_name = "my-azure-account"
}

partition "azure_activity_log" "my_logs" {
  source "azure_blob_storage" {
    connection   = connection.azure.logging_account
    account_name = "my-azure-account"
    container    = "logs-container"
  }
}
```

### Collect Logs with a Prefix Filter

```hcl
partition "azure_activity_log" "my_logs_prefix" {
  source "azure_blob_storage" {
    connection   = connection.azure.logging_account
    account_name = "my-azure-account"
    container    = "logs-container"
    prefix       = "logs/2024/"
  }
}
```

### Collect Logs with Specific File Extensions

```hcl
partition "azure_activity_log" "my_logs_extensions" {
  source "azure_blob_storage" {
    connection   = connection.azure.logging_account
    account_name = "my-azure-account"
    container    = "logs-container"
    extensions   = [".json", ".log"]
  }
}
```

## Arguments

| Argument      | Required | Default                  | Description                                                                                                                |
|--------------|----------|--------------------------|----------------------------------------------------------------------------------------------------------------------------|
| connection   | Yes      |                          | The connection to use for accessing the Azure account.                                                                     |
| account_name | Yes      |                          | The name of the Azure Blob Storage account to collect logs from.                                                           |
| container    | Yes      |                          | The name of the container where logs are stored.                                                                           |
| prefix       | No       | Root of the container   | The prefix to filter objects in the container.                                                                             |
| extensions   | No       | All files                | The file extensions to collect (e.g., `.json`, `.log`).                                                                    |

### Table Defaults

The following tables define their own default values for certain source arguments:

- **[azure_activity_log](https://tailpipe.io/plugins/turbot/azure/tables/azure_activity_log#azure_blob_storage)**

