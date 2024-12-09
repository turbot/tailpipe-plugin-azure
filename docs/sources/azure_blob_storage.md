---
title: "Source: azure_blob_storage - Obtain logs from Azure Blob Storage"
description: "Allows users to collect logs from Azure Blob Storage."
---

# Source: azure_blob_storage - Obtain logs from Azure Blob Storage

Azure Blob Storage is a public cloud storage resource available in Microsoft Azure. It is used to store objects, which consist of data and its descriptive metadata. Blob Storage makes it possible to store and retrieve varying amounts of data, at any time, from anywhere on the web.

## Configuration

| Property | Description | Default |
| - |----------------------------------------------------------------------------------------------|---------------------------|
| `connection` | The connection to use to connect to the Azure account. | - |
| `account_name` | The name of the Azure Blob Storage account to collect logs from. | - |
| `container` | The name of the container to collect logs from. | - |
| `prefix` | The prefix to filter objects in the container. | Defaults to container root. |
| `extensions` | The file extensions to collect. | Defaults to all files. |

