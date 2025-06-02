## Cost Analysis Examples

### Monthly Cost Trends

Track monthly cost trends to identify spending patterns over time.

```sql
select
  date_trunc('month', date) as month,
  sum(cost_in_usd) as monthly_cost_usd
from
  azure_cost_and_usage_details
group by
  month
order by
  month asc;
```

```yaml
folder: Cost and Usage Details
```
### Daily Cost Breakdown
Analyze daily costs to identify unusual spending patterns.
```sql
select
  tp_date as usage_date,
  sum(cost_in_usd) as daily_cost_usd
from
  azure_cost_and_usage_details
group by
  usage_date
order by
  usage_date desc
limit 30;
```

```yaml
folder: Cost and Usage Details
```
### Top 10 Services by Cost
Identify the most expensive services to focus optimization efforts.
```sql
select
  service_family,
  sum(cost_in_usd) as total_cost_usd,
  count(*) as usage_count
from
  azure_cost_and_usage_details
group by
  service_family
order by
  total_cost_usd desc
limit 10;
```

```yaml
folder: Cost and Usage Details
```
### Cost by Meter Category
Break down costs by meter category to understand spending across different types of resources.
```sql
select
  meter_category,
  sum(cost_in_usd) as total_cost_usd,
  count(*) as usage_count
from
  azure_cost_and_usage_details
where
  meter_category is not null
group by
  meter_category
order by
  total_cost_usd desc;
```

```yaml
folder: Cost and Usage Details
```
## Resource Examples
### Top 10 Most Expensive Resources
Identify the most expensive individual resources.
```sql
select
  resource_id,
  resource_group_name,
  product_name,
  sum(cost_in_usd) as total_cost_usd
from
  azure_cost_and_usage_details
where
  resource_id is not null
group by
  resource_id,
  resource_group_name,
  product_name
order by
  total_cost_usd desc
limit 10;
```

```yaml
folder: Cost and Usage Details
```
### Cost by Resource Group
Analyze costs by resource group to identify high-spending areas.
```sql
select
  resource_group_name,
  sum(cost_in_usd) as total_cost_usd,
  count(*) as usage_count
from
  azure_cost_and_usage_details
where
  resource_group_name is not null
group by
  resource_group_name
order by
  total_cost_usd desc;
```

```yaml
folder: Cost and Usage Details
```
### Cost by Resource Location
Analyze costs by location to understand geographical distribution of spending.
```sql
select
  resource_location,
  sum(cost_in_usd) as total_cost_usd,
  count(*) as usage_count
from
  azure_cost_and_usage_details
where
  resource_location is not null
group by
  resource_location
order by
  total_cost_usd desc;
```

```yaml
folder: Cost and Usage Details
```
## Subscription Examples
### Cost by Subscription
Compare costs across different subscriptions.
```sql
select
  subscription_id,
  subscription_name,
  sum(cost_in_usd) as total_cost_usd
from
  azure_cost_and_usage_details
group by
  subscription_id,
  subscription_name
order by
  total_cost_usd desc;
```

```yaml
folder: Cost and Usage Details
```
### Monthly Costs by Subscription
Track monthly costs for each subscription to identify trends.
```sql
select
  subscription_name,
  date_trunc('month', date) as month,
  sum(cost_in_usd) as monthly_cost_usd
from
  azure_cost_and_usage_details
group by
  subscription_name,
  month
order by
  subscription_name,
  month;
```

```yaml
folder: Cost and Usage Details
```
## Optimization Examples
### Resources with Azure Credit Eligibility
Identify resources eligible for Azure credits to optimize cost allocation.
```sql
select
  resource_id,
  resource_group_name,
  product_name,
  sum(cost_in_usd) as total_cost_usd,
  is_azure_credit_eligible
from
  azure_cost_and_usage_details
where
  resource_id is not null
group by
  resource_id,
  resource_group_name,
  product_name,
  is_azure_credit_eligible
order by
  total_cost_usd desc;
```

```yaml
folder: Cost and Usage Details
```
### Cost by Pricing Model
Compare costs across different pricing models to identify optimization opportunities.
```sql
select
  pricing_model,
  sum(cost_in_usd) as total_cost_usd,
  count(*) as usage_count
from
  azure_cost_and_usage_details
where
  pricing_model is not null
group by
  pricing_model
order by
  total_cost_usd desc;
```

```yaml
folder: Cost and Usage Details
```
### Reservation Usage Analysis
Analyze costs related to reservations to evaluate their effectiveness.
```sql
select
  reservation_id,
  reservation_name,
  sum(cost_in_usd) as total_cost_usd,
  count(*) as usage_count
from
  azure_cost_and_usage_details
where
  reservation_id is not null
group by
  reservation_id,
  reservation_name
order by
  total_cost_usd desc;
```

```yaml
folder: Cost and Usage Details
```
### Pay-As-You-Go vs. Effective Price Comparison
Compare pay-as-you-go prices with effective prices to quantify savings from discounts.
```sql
select
  product_name,
  sum(cost_in_usd) as actual_cost,
  sum(payg_cost_in_usd) as payg_cost,
  sum(payg_cost_in_usd - cost_in_usd) as savings
from
  azure_cost_and_usage_details
where
  payg_cost_in_usd is not null
  and cost_in_usd is not null
group by
  product_name
order by
  savings desc;
```

```yaml
folder: Cost and Usage Details
```