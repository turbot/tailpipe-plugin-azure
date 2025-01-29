## Activity Examples

### Daily activity trends

Count events per day to identify activity trends over time.

```sql
select
  strftime(event_timestamp, '%Y-%m-%d') AS event_date,
  count(*) AS event_count
from
  azure_activity_log
group by
  event_date
order by
  event_date asc;
```

### Top 10 events

List the 10 most frequently called events.

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

### Top 10 events (exclude read-only)

List the top 10 most frequently called events, excluding read-only events.

```sql
select
  resource_provider_name as event_source,
  operation_name as event_name,
  count(*) as event_count
from
  azure_activity_log
where
  status != 'Succeeded'
group by
  resource_provider_name,
  operation_name
order by
  event_count desc
limit 10;
```

### Top events by subscription

Count and group events by subscription ID, event source, and event name to analyze activity across subscriptions.

```sql
select
  resource_provider_name as event_source,
  operation_name as event_name,
  subscription_id,
  count(*) as event_count
from
  azure_activity_log
group by
  resource_provider_name,
  operation_name,
  subscription_id
order by
  event_count desc;
```

### Top error codes

Identify the most frequent error codes.

```sql
select
  sub_status as error_code,
  count(*) as event_count
from
  azure_activity_log
where
  sub_status is not null
group by
  sub_status
order by
  event_count desc;
```

## Detection Examples

### Azure Key Vault secret access

Detect when secrets in Azure Key Vault are accessed.

```sql
select
  event_timestamp,
  resource_provider_name,
  operation_name,
  caller,
  resource_id,
  status
from
  azure_activity_log
where
  resource_provider_name = 'Microsoft.KeyVault'
  and operation_name like '%Secret%'
order by
  event_timestamp desc;
```

### Azure Activity Log retention policy changes

Detect when Activity Log retention policies are modified.

```sql
select
  event_timestamp,
  resource_provider_name,
  operation_name,
  caller,
  subscription_id,
  status
from
  azure_activity_log
where
  resource_provider_name = 'Microsoft.Insights'
  and operation_name like '%RetentionPolicy%'
order by
  event_timestamp desc;
```

### Unauthorized login attempts

Find failed login attempts that may indicate unauthorized access attempts.

```sql
select
  event_timestamp,
  operation_name,
  caller,
  resource_id,
  status
from
  azure_activity_log
where
  operation_name = 'SignInLogs'
  and status != 'Succeeded'
order by
  event_timestamp desc;
```

### Root activity

Track any actions performed by privileged root accounts.

```sql
select
  event_timestamp,
  operation_name,
  caller,
  resource_provider_name,
  subscription_id
from
  azure_activity_log
where
  caller = 'Root'
order by
  event_timestamp desc;
```

### Activity in unapproved regions

Identify actions occurring in Azure regions outside an approved list.

```sql
select
  event_timestamp,
  resource_provider_name,
  operation_name,
  caller,
  subscription_id
from
  azure_activity_log
where
  resource_id not like '%/locations/eastus%' -- Example filtering for specific regions
order by
  event_timestamp desc;
```

### Activity from unapproved IP addresses

Flag activity originating from IP addresses outside an approved list.

```sql
select
  event_timestamp,
  resource_provider_name,
  operation_name,
  caller,
  tp_source_ip
from
  azure_activity_log
where
  tp_source_ip not in ('192.168.1.1', '10.0.0.2')
order by
  event_timestamp desc;
```

## Operational Examples

### Network security group rule updates

Track changes to network security group rules.

```sql
select
  event_timestamp,
  resource_provider_name,
  operation_name,
  caller,
  resource_id,
  status
from
  azure_activity_log
where
  resource_provider_name = 'Microsoft.Network'
  and operation_name like '%SecurityRule%'
order by
  event_timestamp desc;
```

### Azure role assignments

List events where a user has added or removed role assignments.

```sql
select
  event_timestamp,
  resource_provider_name,
  operation_name,
  caller,
  resource_id,
  status
from
  azure_activity_log
where
  resource_provider_name = 'Microsoft.Authorization'
  and operation_name like '%roleAssignment%'
order by
  event_timestamp desc;
```

## Volume Examples

### High volume of storage account access requests

Detect unusually high access activity to Azure Storage accounts.

```sql
select
  caller,
  count(*) as event_count,
  date_trunc('minute', event_timestamp) as event_minute
from
  azure_activity_log
where
  resource_provider_name = 'Microsoft.Storage'
  and operation_name like '%StorageAccount%'
group by
  caller,
  event_minute
having
  count(*) > 100
order by
  event_count desc;
```

### Excessive Azure role assumptions

Identify role assignments occurring at an unusually high frequency.

```sql
select
  caller,
  count(*) as event_count,
  date_trunc('hour', event_timestamp) as event_hour
from
  azure_activity_log
where
  resource_provider_name = 'Microsoft.Authorization'
  and operation_name like '%roleAssignment%'
group by
  caller,
  event_hour
having
  count(*) > 10
order by
  event_hour desc,
  event_count desc;
```

