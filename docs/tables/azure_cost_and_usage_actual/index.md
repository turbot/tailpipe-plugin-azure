---
title: "Tailpipe Table: azure_cost_and_usage_actual - Query Azure Cost and Usage Actual data"
description: "Azure Cost and Usage Actual data provides detailed information about Azure resource usage and costs, including subscription charges, resource consumption, pricing details, and billing information."
---

# Table: azure_cost_and_usage_actual - Query Azure Cost and Usage Actual data

The `azure_cost_and_usage_actual` table allows you to query data from Azure Cost and Usage Actual exports. This table provides detailed information about Azure resource usage and costs, including subscription charges, resource consumption, pricing details, and billing information, enabling cost analysis, budget tracking, and optimization insights across Azure subscriptions.

Limitations and notes:

- The table currently supports exported logs in CSV format (compressed or uncompressed).
- Files must follow the Azure Cost Management export naming pattern: `part_X_YYYY.csv.gz` or `part_X_YYYY.csv.zip`.
- Cost data must be exported from Azure Cost Management and downloaded locally before processing.

## Configure

Create a [partition](https://tailpipe.io/docs/manage/partition) for `azure_cost_and_usage_actual` ([examples](https://hub.tailpipe.io/plugins/turbot/azure/tables/azure_cost_and_usage_actual#example-configurations)):

```sh
vi ~/.tailpipe/config/azure.tpc
```

```hcl
partition "azure_cost_and_usage_actual" "my_costs" {
  source "file" {
    paths = ["/path/to/cost/data/"]
    file_layout = "part_%{INT:part_number}_%{INT:file_number}.csv.(?:gz|zip)"
  }
}
```

## Collect

[Collect](https://tailpipe.io/docs/manage/collection) data for all `azure_cost_and_usage_actual` partitions:

```sh
tailpipe collect azure_cost_and_usage_actual
```

Or for a single partition:

```sh
tailpipe collect azure_cost_and_usage_actual.my_costs
```

## Query

**[Explore example queries for this table →](https://hub.tailpipe.io/plugins/turbot/azure/queries/azure_cost_and_usage_actual)**

### Monthly cost by service

Analyze monthly costs by service to identify spending trends and high-cost services.

```sql
select
  date_trunc('month', date) as month,
  service_family,
  sum(cost_in_usd) as total_cost_usd
from
  azure_cost_and_usage_actual
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
  azure_cost_and_usage_actual
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
  azure_cost_and_usage_actual
group by
  usage_date
order by
  usage_date desc;
```

## Example Configurations

### Collect cost data from local files

Collect Azure Cost Management export files stored locally.

```hcl
partition "azure_cost_and_usage_actual" "my_costs" {
  source "file" {
    paths = ["/path/to/cost/data/"]
    file_layout = "part_%{INT:part_number}_%{INT:file_number}.csv.(?:gz|zip)"
  }
}
```