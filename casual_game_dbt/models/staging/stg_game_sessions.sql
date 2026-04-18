{{ config(
    materialized='table',
    partition_by={
        "field": "event_time",
        "data_type": "timestamp",
        "granularity": "day"
    },
    cluster_by=["platform", "country"]
) }}

with raw_data as (
    select * from {{ source('raw_events', 'events_raw_ext') }}
)

select
    cast(event_id as string) as event_id,
    cast(user_id as string) as user_id,
    cast(platform as string) as platform,
    cast(timestamp as timestamp) as event_time,
    cast(payload.app_version as string) as app_version,
    cast(payload.country as string) as country
from raw_data
where type = 'session_start'