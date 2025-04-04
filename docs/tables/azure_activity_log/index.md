---
title: "Tailpipe Table: azure_activity_log - Query Azure activity logs"
description: "Azure activity logs capture administrative actions, policy changes, and security events within your Azure environment."
---

# Table: azure_activity_log - Query Azure activity logs

The `azure_activity_log` table allows you to query data from Azure activity logs. This table provides detailed information about API calls, resource modifications, security events, and administrative actions within your Azure environment.

## Configure

Create a [partition](https://tailpipe.io/docs/manage/partition) for `azure_activity_log` ([examples](https://hub.tailpipe.io/plugins/turbot/azure/tables/azure_activity_log#example-configurations)):

```sh
vi ~/.tailpipe/config/azure.tpc
```

```hcl
connection "azure" "logging_account" {
  tenant_id       = "00000000-0000-0000-0000-000000000000"
  subscription_id = "00000000-0000-0000-0000-000000000000"
  client_id       = "00000000-0000-0000-0000-000000000000"
  client_secret   = "my plaintext secret"
}

partition "azure_activity_log" "my_logs" {
  source "azure_blob_storage" {
    connection   = connection.azure.logging_account
    account_name = "storage_account_name"
    container    = "container_name"
  }
}
```

## Collect

[Collect](https://tailpipe.io/docs/manage/collection) logs for all `azure_activity_log` partitions:

```sh
tailpipe collect azure_activity_log
```

Or for a single partition:

```sh
tailpipe collect azure_activity_log.my_logs
```

## Query

**[Explore 40+ example queries for this table →](https://hub.tailpipe.io/plugins/turbot/azure/queries/azure_activity_log)**

### Role assigments

List role assignments to check for unexpected or suspicious role changes.

```sql
select
  event_timestamp,
  resource_id,
  caller,
  resource_group_name,
  subscription_id
from
  azure_activity_log
where
  operation_name = 'Microsoft.Authorization/roleAssignments/write'
order by
  event_timestamp desc;
```

### Top 10 events

List the top 10 events and how many times they were called.

```sql
select
  resource_provider_name,
  operation_name,
  count(*) as event_count
from
  azure_activity_log
group by
  resource_provider_name,
  operation_name
order by
  event_count desc
limit 10;
```

### High volume Storage access requests

Find users generating a high volume of Storage access requests to identify potential anomalous activity.

```sql
select
  caller,
  count(*) as event_count,
  date_trunc('minute', event_timestamp) as event_minute
from
  azure_activity_log
where
  operation_name = 'Microsoft.Storage/storageAccounts/listKeys/action'
group by
  caller,
  event_minute
having
  count(*) > 100
order by
  event_count desc;
```

## Example Configurations

### Collect logs from a storage account

Collect activity logs stored in a storage account that use the [default blob naming convention](https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/activity-log?tabs=powershell#send-to-azure-storage).

```hcl
connection "azure" "my_logging_account" {
  tenant_id       = "00000000-0000-0000-0000-000000000000"
  subscription_id = "00000000-0000-0000-0000-000000000000"
  client_id       = "00000000-0000-0000-0000-000000000000"
  client_secret   = "my plaintext secret"
}

partition "azure_activity_log" "my_logs" {
  source "azure_blob_storage" {
    connection   = connection.azure.my_logging_account
    account_name = "storage_account_name"
    container    = "container_name"
  }
}
```

### Collect logs from Monitor activity logs API

Collect activity logs from the Monitor activity logs API.

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

### Exclude read-only events

Use the filter argument in your partition to exclude specific events and and reduce log storage size.

```hcl
partition "azure_activity_log" "my_logs_filtered" {
  # Avoid saving unnecessary events, which can drastically reduce local log size
  filter = "operation_name != 'Microsoft.Storage/storageAccounts/listKeys/action'"

  source "azure_activity_log_api" {
    connection = connection.azure.my_subscription
  }
}
```

### Collect logs for a single subscription

Collect logs for a specific subscription.

```hcl
partition "azure_activity_log" "my_logs_subscription" {
  source "azure_blob_storage" {
    connection   = connection.azure.my_logging_account
    account_name = "storage_account_name"
    container    = "container_name"
    file_layout  = `/SUBSCRIPTIONS/12345678-1234-1234-1234-123456789012/y=%{YEAR:year}/m=%{MONTHNUM:month}/d=%{MONTHDAY:day}/h=%{HOUR:hour}/m=%{MINUTE:minute}/%{DATA:filename}.json`
  }
}
```

## Source Defaults

### azure_blob_storage

This table sets the following defaults for the [azure_blob_storage source](https://hub.tailpipe.io/plugins/turbot/azure/sources/azure_blob_storage#arguments):

| Argument    | Default |
|-------------|---------|
| file_layout | `/SUBSCRIPTIONS/%{DATA:subscription_id}/y=%{YEAR:year}/m=%{MONTHNUM:month}/d=%{MONTHDAY:day}/h=%{HOUR:hour}/m=%{MINUTE:minute}/%{DATA:filename}.json` |
