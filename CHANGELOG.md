## v0.6.1 [2025-07-28]

_Dependencies_

- Recompiled plugin with [tailpipe-plugin-sdk v0.9.2](https://github.com/turbot/tailpipe-plugin-sdk/blob/develop/CHANGELOG.md#v092-2025-07-24) that fixes incorrect data ranges for zero‑granularity collections and prevents crashes in certain collection states. ([#88](https://github.com/turbot/tailpipe-plugin-azure/pull/88))

## v0.6.0 [2025-07-02]

_Dependencies_

- Recompiled plugin with [tailpipe-plugin-sdk v0.9.1](https://github.com/turbot/tailpipe-plugin-sdk/blob/develop/CHANGELOG.md#v091-2025-07-02) to support the `--to` flag, directional time-based collection, improved tracking of collected data and fixed collection state issues. ([#81](https://github.com/turbot/tailpipe-plugin-azure/pull/81))

## v0.5.0 [2025-06-20]

_What's new?_

- New tables added
  - [azure_cost_and_usage_actual](https://hub.tailpipe.io/plugins/turbot/azure/tables/azure_cost_and_usage_actual) ([#68](https://github.com/turbot/tailpipe-plugin-azure/pull/68))

## v0.4.2 [2025-06-05]

- Recompiled plugin with [tailpipe-plugin-sdk v0.7.2](https://github.com/turbot/tailpipe-plugin-sdk/blob/develop/CHANGELOG.md#v072-2025-06-04) that fixes an issue where the end time was not correctly recorded for collections using artifact sources. ([#73](https://github.com/turbot/tailpipe-plugin-azure/pull/73))

## v0.4.1 [2025-06-04]

- Recompiled plugin with [tailpipe-plugin-sdk v0.7.1](https://github.com/turbot/tailpipe-plugin-sdk/blob/develop/CHANGELOG.md#v071-2025-06-04) that fixes an issue affecting collections using a file source. ([#71](https://github.com/turbot/tailpipe-plugin-azure/pull/71))

## v0.4.0 [2025-06-03]

_Dependencies_

- Recompiled plugin with [tailpipe-plugin-sdk v0.7.0](https://github.com/turbot/tailpipe-plugin-sdk/blob/develop/CHANGELOG.md#v070-2025-06-03) that improves how collection end times are tracked, helping make future collections more accurate and reliable. ([#70](https://github.com/turbot/tailpipe-plugin-azure/pull/70))

## v0.3.0 [2025-03-03]

_Enhancements_

- Standardized all example query titles to use `Title Case` for consistency. ([#43](https://github.com/turbot/tailpipe-plugin-azure/pull/43))
- Added `folder` front matter to all queries for improved organization and discoverability in the Hub. ([#43](https://github.com/turbot/tailpipe-plugin-azure/pull/43))

## v0.2.0 [2025-02-06]

_Enhancements_

- Updated documentation formatting and enhanced argument descriptions for `azure_activity_log_api` and `azure_blob_storage` sources. ([#34](https://github.com/turbot/tailpipe-plugin-azure/pull/34))

## v0.1.0 [2025-01-30]

_What's new?_

- New tables added
  - [azure_activity_log](https://hub.tailpipe.io/plugins/turbot/azure/tables/azure_activity_log)
- New sources added
  - [azure_activity_log_api](https://hub.tailpipe.io/plugins/turbot/azure/sources/azure_activity_log_api)
  - [azure_blob_storage](https://hub.tailpipe.io/plugins/turbot/azure/sources/azure_blob_storage)
