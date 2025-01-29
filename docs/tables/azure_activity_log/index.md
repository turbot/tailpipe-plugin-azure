---
title: "Tailpipe Table: azure_activity_log - Query Azure Activity Logs"
description: "Azure Activity Logs capture administrative actions, policy changes, and security events within your Azure environment."
---

# Table: azure_activity_log - Query Azure Activity Logs

The `azure_activity_log` table allows you to query data from Azure Activity Logs. This table provides detailed information about API calls, resource modifications, security events, and administrative actions within your Azure environment.

## Configure

Create a [partition](https://tailpipe.io/docs/manage/partition) for `azure_activity_log`:

```sh
vi ~/.tailpipe/config/azure.tpc
```

```hcl
connection "azure" "logging_account" {
  subscription_id = "my-subscription-id"
}

partition "azure_activity_log" "my_logs" {
  source "azure_monitor" {
    connection = connection.azure.logging_account
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

**[Explore 100+ example queries for this table →](https://hub.tailpipe.io/plugins/turbot/azure/queries/azure_activity_log)**

### Root Activity

Find any actions taken by the root user.

```sql
select
  event_timestamp,
  event_name,
  caller,
  resource_group_name,
  resource_provider_name,
  subscription_id
from
  azure_activity_log
where
  caller = 'Root'
order by
  event_timestamp desc;
```

### Top 10 Events

List the top 10 events and how many times they were called.

```sql
select
  resource_provider_name as event_source,
  operation_name as event_name,
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

### High Volume Resource Modifications

Find users generating a high volume of resource modifications to identify potential anomalous activity.

```sql
select
  caller,
  count(*) as event_count,
  date_trunc('minute', event_timestamp) as event_minute
from
  azure_activity_log
where
  category = 'Administrative'
group by
  caller,
  event_minute
having
  count(*) > 100
order by
  event_count desc;
```

## Example Configurations

### Collect logs from Azure Monitor

Collect Azure Activity Logs using Azure Monitor:

```hcl
partition "azure_activity_log" "my_logs" {
  source "azure_monitor" {
    connection = connection.azure.logging_account
  }
}
```

### Exclude Read-Only Events

Use the filter argument in your partition to exclude read-only events and reduce log storage size.

```hcl
partition "azure_activity_log" "my_logs_write" {
  # Avoid saving read-only events
  filter = "status != 'Succeeded'"

  source "azure_monitor" {
    connection = connection.azure.logging_account
  }
}
```

### Collect logs for all subscriptions in a tenant

For a specific tenant, collect logs for all subscriptions.

```hcl
partition "azure_activity_log" "my_logs_tenant" {
  source "azure_monitor"  {
    connection  = connection.azure.logging_account
  }
}
```

### Collect logs for a single subscription

For a specific subscription, collect logs for all resource groups.

```hcl
partition "azure_activity_log" "my_logs_subscription" {
  source "azure_monitor"  {
    connection  = connection.azure.logging_account
    subscription_id = "my-subscription-id"
  }
}
```

## Source Defaults

### azure_monitor

This table sets the following defaults for the [azure_monitor source](https://tailpipe.io/plugins/turbot/azure/sources/azure_monitor#arguments):

| Argument      | Default |
|--------------|---------|
| log_type     | `AzureActivityLog` |

