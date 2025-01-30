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

### Top 10 failed events

List the top 10 most frequently called events that failed.

```sql
select
  resource_provider_name,
  operation_name,
  count(*) as event_count
from
  azure_activity_log
where
  status = 'Failed'
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
  resource_provider_name,
  operation_name,
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
  sub_status,
  count(*) as event_count
from
  azure_activity_log
where
  sub_status not in ('', 'OK', 'Created', 'Accepted', 'NoContent')
group by
  sub_status
order by
  event_count desc;
```

## Detection Examples

### Activity from unapproved IP addresses

Flag activity originating from IP addresses outside an approved list.

```sql
select
  event_timestamp,
  operation_name,
  resource_id,
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

### Azure role assignments

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

## Volume Examples

### High volume of storage account access requests

Detect unusually high access activity to storage accounts.

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
