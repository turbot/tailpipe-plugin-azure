---
title: "Tailpipe Table: azure_cost_management - Query Azure Cost Management data"
description: "Azure Cost Management data provides detailed information about Azure resource usage and costs, including subscription charges, resource consumption, pricing details, and billing information."
---

# Table: azure_cost_management - Query Azure Cost Management data

The `azure_cost_management` table allows you to query data from Azure Cost Management exports. This table provides detailed information about Azure resource usage and costs, including subscription charges, resource consumption, pricing details, and billing information, enabling cost analysis, budget tracking, and optimization insights across Azure subscriptions.

## Configure

Create a [partition](https://tailpipe.io/docs/manage/partition) for `azure_cost_management` ([examples](#example-configurations)):

```sh
vi ~/.tailpipe/config/azure.tpc
```

```hcl
connection "azure" "cost_account" {
  tenant_id       = "00000000-0000-0000-0000-000000000000"
  subscription_id = "00000000-0000-0000-0000-000000000000"
  client_id       = "00000000-0000-0000-0000-000000000000"
  client_secret   = "my plaintext secret"
}

partition "azure_cost_management" "my_costs" {
  source "azure_blob_storage" {
    connection   = connection.azure.cost_account
    account_name = "storage_account_name"
    container    = "container_name"
  }
}
```

## Collect

[Collect](https://tailpipe.io/docs/manage/collection) data for all `azure_cost_management` partitions:

```sh
tailpipe collect azure_cost_management
```

Or for a single partition:

```sh
tailpipe collect azure_cost_management.my_costs
```

## Query

**[Explore example queries for this table →](https://hub.tailpipe.io/plugins/turbot/azure/queries/azure_cost_management)**

### Monthly cost by service

Analyze monthly costs by service to identify spending trends and high-cost services.

```sql
select
  date_trunc('month', date) as month,
  service_family,
  sum(cost_in_usd) as total_cost_usd
from
  azure_cost_management
group by
  month,
  service_family
order by
  month,
  total_cost_usd desc;
```

### Resource groups with highest costs

Identify resource groups with the highest costs to focus optimization efforts.

```sql
select
  resource_group_name,
  sum(cost_in_usd) as total_cost_usd,
  count(*) as usage_count
from
  azure_cost_management
where
  resource_group_name is not null
group by
  resource_group_name
order by
  total_cost_usd desc
limit 10;
```

### Daily cost trends

Track daily cost trends to identify unusual spending patterns.

```sql
select
  tp_date as usage_date,
  sum(cost_in_usd) as daily_cost_usd
from
  azure_cost_management
group by
  usage_date
order by
  usage_date desc;
```

## Example Configurations

### Collect cost data from a storage account

Collect Azure Cost Management data stored in a storage account.

```hcl
connection "azure" "cost_account" {
  tenant_id       = "00000000-0000-0000-0000-000000000000"
  subscription_id = "00000000-0000-0000-0000-000000000000"
  client_id       = "00000000-0000-0000-0000-000000000000"
  client_secret   = "my plaintext secret"
}

partition "azure_cost_management" "my_costs" {
  source "azure_blob_storage" {
    connection   = connection.azure.cost_account
    account_name = "storage_account_name"
    container    = "container_name"
  }
}
```

### Filter costs by subscription

Use the filter argument to focus on costs from a specific subscription.

```hcl
partition "azure_cost_management" "subscription_costs" {
  filter = "subscription_id = '00000000-0000-0000-0000-000000000000'"

  source "azure_blob_storage" {
    connection   = connection.azure.cost_account
    account_name = "storage_account_name"
    container    = "container_name"
  }
}
```

### Filter costs by date range

Filter costs to a specific date range to analyze spending during a particular period.

```hcl
partition "azure_cost_management" "recent_costs" {
  filter = "date >= '2023-01-01' and date <= '2023-12-31'"

  source "azure_blob_storage" {
    connection   = connection.azure.cost_account
    account_name = "storage_account_name"
    container    = "container_name"
  }
}
```

## Source Defaults

### azure_blob_storage

This table sets the following defaults for the [azure_blob_storage source](https://hub.tailpipe.io/plugins/turbot/azure/sources/azure_blob_storage#arguments):

| Argument    | Default |
|-------------|---------|
| file_layout | `part_%{INT:part_number}_%{INT:file_number}.csv.(?:gz|zip)` |
