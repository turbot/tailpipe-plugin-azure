# Azure Plugin for Tailpipe

[Tailpipe](https://tailpipe.io) is an open-source CLI tool that allows you to collect logs and query them with SQL.

[Azure](https://azure.microsoft.com) provides on-demand cloud computing platforms and APIs to authenticated customers on a metered pay-as-you-go basis.

The [Azure Plugin for Tailpipe](https://hub.tailpipe.io/plugins/turbot/azure) allows you to collect and query Azure logs using SQL to track activity, monitor trends, detect anomalies, and more!

- **[Get started →](https://hub.tailpipe.io/plugins/turbot/azure)**
- Documentation: [Table definitions & examples](https://hub.tailpipe.io/plugins/turbot/azure/tables)
- Community: [Join #tailpipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/tailpipe-plugin-azure/issues)

Collect and query logs:
![image](docs/images/azure_activity_log_terminal.png)

## Getting Started

Install Tailpipe from the [downloads](https://tailpipe.io/downloads) page:

```sh
# MacOS
brew install turbot/tap/tailpipe
```

```sh
# Linux or Windows (WSL)
sudo /bin/sh -c "$(curl -fsSL https://tailpipe.io/install/tailpipe.sh)"
```

Install the plugin:

```sh
tailpipe plugin install azure
```

Configure your [connection credentials](https://hub.tailpipe.io/plugins/turbot/azure#connection-credentials), table partition, and data source ([examples](https://hub.tailpipe.io/plugins/turbot/azure/tables/azure_activity_log#example-configurations)):

```sh
vi ~/.tailpipe/config/azure.tpc
```

```hcl
connection "azure" "my_subscription" {
  tenant_id       = "00000000-0000-0000-0000-000000000000"
  subscription_id = "00000000-0000-0000-0000-000000000000"
  client_id       = "00000000-0000-0000-0000-000000000000"
  client_secret   = "my plaintext secret"
}

partition "azure_activity_log" "my_logs" {
  source "azure_blob_storage" {
    connection   = connection.azure.my_subscription
    account_name = "storage_account_name"
    container    = "container_name"
  }
}
```

Download, enrich, and save logs from your source ([examples](https://tailpipe.io/docs/reference/cli/collect)):

```sh
tailpipe collect azure_activity_log
```

Enter interactive query mode:

```sh
tailpipe query
```

Run a query:

```sql
select
  resource_type,
  operation_name,
  count(*) as operation_count
from
  azure_activity_log
group by
  resource_type,
  operation_name
order by
  operation_count desc;
```

```sh
+-----------------------------------------------------------+------------------------------------------------------------------+-----------------+
| resource_type                                             | operation_name                                                   | operation_count |
+-----------------------------------------------------------+------------------------------------------------------------------+-----------------+
| Microsoft.Resources/deployments                           | Microsoft.Resources/deployments/write                            | 86              |
| Microsoft.Resources/deployments                           | Microsoft.Resources/deployments/validate/action                  | 58              |
| Microsoft.Compute/virtualMachines                         | Microsoft.Authorization/policies/auditIfNotExists/action         | 54              |
| Microsoft.Compute/virtualMachines                         | Microsoft.Authorization/policies/audit/action                    | 36              |
| Microsoft.Sql/servers                                     | Microsoft.Authorization/policies/auditIfNotExists/action         | 25              |
| Microsoft.Sql/servers/databases                           | Microsoft.Sql/servers/databases/read                             | 20              |
| MICROSOFT.CDN/profiles                                    | Microsoft.Resourcehealth/healthevent/Activated/action            | 18              |
+-----------------------------------------------------------+------------------------------------------------------------------+-----------------+
```

## Detections as Code with Powerpipe

Pre-built dashboards and detections for the Azure plugin are available in [Powerpipe](https://powerpipe.io) mods, helping you monitor and analyze activity across your Azure subscriptions.

For example, the [Azure Activity Log Detections mod](https://hub.powerpipe.io/mods/turbot/tailpipe-mod-azure-activity-log-detections) scans your activity logs for anomalies, such as a SQL server firewall rule getting updated or a change in your virtual networks.

Dashboards and detections are [open source](https://github.com/topics/tailpipe-mod), allowing easy customization and collaboration.

To get started, choose a mod from the [Powerpipe Hub](https://hub.powerpipe.io/?engines=tailpipe&q=azure).

![image](docs/images/azure_activity_log_mitre_dashboard.png)

## Developing

Prerequisites:

- [Tailpipe](https://tailpipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/tailpipe-plugin-azure.git
cd tailpipe-plugin-azure
```

After making your local changes, build the plugin, which automatically installs the new version to your `~/.tailpipe/plugins` directory:

```sh
make
```

Re-collect your data:

```sh
tailpipe collect azure_activity_log
```

Try it!

```sh
tailpipe query
> .inspect azure_activity_log
```

## Open Source & Contributing

This repository is published under the [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0) (source code) and [CC BY-NC-ND](https://creativecommons.org/licenses/by-nc-nd/2.0/) (docs) licenses. Please see our [code of conduct](https://github.com/turbot/.github/blob/main/CODE_OF_CONDUCT.md). We look forward to collaborating with you!

[Tailpipe](https://tailpipe.io) is a product produced from this open source software, exclusively by [Turbot HQ, Inc](https://turbot.com). It is distributed under our commercial terms. Others are allowed to make their own distribution of the software, but cannot use any of the Turbot trademarks, cloud services, etc. You can learn more in our [Open Source FAQ](https://turbot.com/open-source).

## Get Involved

**[Join #tailpipe on Slack →](https://turbot.com/community/join)**

Want to help but don't know where to start? Pick up one of the `help wanted` issues:

- [Tailpipe](https://github.com/turbot/tailpipe/labels/help%20wanted)
- [Azure Plugin](https://github.com/turbot/tailpipe-plugin-azure/labels/help%20wanted)
