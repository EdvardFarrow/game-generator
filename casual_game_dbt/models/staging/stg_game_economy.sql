{{ config(
    materialized='table',
    partition_by={
        "field": "event_time",
        "data_type": "timestamp",
        "granularity": "day"
    },
    cluster_by=["currency", "item"]
) }}

with raw_data as (
    select * from {{ source('raw_events', 'events_raw_ext') }}
)

select
    cast(event_id as string) as event_id,
    cast(user_id as string) as user_id,
    cast(platform as string) as platform,
    cast(timestamp as timestamp) as event_time,
    cast(payload.currency as string) as currency,
    cast(payload.amount as int64) as amount,
    cast(payload.item as string) as item
from raw_data
where type = 'economy_spend'
