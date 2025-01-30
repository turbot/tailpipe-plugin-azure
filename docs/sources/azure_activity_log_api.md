---
title: "Source: azure_activity_log_api - Collect logs from Azure activity log API"
description: "Allows users to collect logs from Azure activity log API."
---

# Source: azure_activity_log_api - Collect logs from Azure activity log API

Azure activity log API provides access to activity logs for Azure resources. These logs help track administrative actions, security events, and operational changes within Azure. The API enables users to query, monitor, and analyze activity logs for auditing, compliance, and security investigations.

Using this source, you can collect, filter, and analyze logs retrieved from the Azure activity log API to enhance visibility into Azure operations and security monitoring.

## Example Configurations

### Collect activity logs

```hcl
connection "azure" "my_subscription" {
  tenant_id       = "00000000-0000-0000-0000-000000000000"
  subscription_id = "00000000-0000-0000-0000-000000000000"
  client_id       = "00000000-0000-0000-0000-000000000000"
  client_secret   = "my plaintext secret"
}

partition "azure_activity_log" "my_logs" {
  source "azure_activity_log_api" {
    connection = connection.azure.my_subscription
  }
}
```

## Arguments

| Argument        | Required | Default                  | Description                                                                                                                 |
|-----------------|----------|--------------------------|-----------------------------------------------------------------------------------------------------------------------------|
| connection      | No       |                          | The connection to use for accessing the Azure account.                                                                     |
