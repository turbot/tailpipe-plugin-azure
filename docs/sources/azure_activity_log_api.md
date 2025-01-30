---
title: "Source: azure_activity_log_api - Collect logs from Azure Activity Log API"
description: "Allows users to collect logs from Azure Activity Log API."
---

# Source: azure_activity_log_api - Collect logs from Azure Activity Log API

Azure Activity Log API provides access to activity logs for Azure resources. These logs help track administrative actions, security events, and operational changes within Azure. The API enables users to query, monitor, and analyze activity logs for auditing, compliance, and security investigations.

Using this source, you can collect, filter, and analyze logs retrieved from the Azure Activity Log API to enhance visibility into Azure operations and security monitoring.

## Example Configurations

### Collect Activity Logs for All Subscriptions

```hcl
connection "azure" "logging_account" {
  tenant_id = "my-tenant-id"
}

partition "azure_activity_log" "my_logs" {
  source "azure_activity_log_api" {
    connection = connection.azure.logging_account
  }
}
```

### Collect Activity Logs for a Specific Subscription

```hcl
partition "azure_activity_log" "my_subscription_logs" {
  source "azure_activity_log_api" {
    connection     = connection.azure.logging_account
    subscription_id = "my-subscription-id"
  }
}
```

## Arguments

| Argument         | Required | Default                  | Description                                                                                                                 |
|-----------------|----------|--------------------------|-----------------------------------------------------------------------------------------------------------------------------|
| connection      | Yes      |                          | The connection to use for accessing the Azure account.                                                                     |
| subscription_id | No       | All subscriptions        | The Azure subscription ID from which logs should be retrieved.                                                             |

### Table Defaults

The following tables define their own default values for certain source arguments:

- **[azure_activity_log](https://tailpipe.io/plugins/turbot/azure/tables/azure_activity_log#azure_activity_log_api)**