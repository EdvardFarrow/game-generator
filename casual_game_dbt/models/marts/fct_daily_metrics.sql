{{ config(
    materialized='table',
    partition_by={
        "field": "event_date",
        "data_type": "date"
    }
) }}

WITH daily_sessions AS (
    SELECT
        DATE(event_time) AS event_date,
        COUNT(DISTINCT user_id) AS dau,
        COUNT(event_id) AS total_sessions
    FROM {{ ref('stg_game_sessions') }}
    GROUP BY 1
),

daily_economy AS (
    SELECT
        DATE(event_time) AS event_date,
        SUM(
            CASE WHEN currency = 'hard_gem' THEN amount ELSE 0 END
        ) AS hard_currency_spent,
        SUM(
            CASE WHEN currency = 'soft_coin' THEN amount ELSE 0 END
        ) AS soft_currency_spent,
        COUNT(event_id) AS total_transactions
    FROM {{ ref('stg_game_economy') }}
    GROUP BY 1
)

SELECT
    COALESCE(s.event_date, e.event_date) AS event_date,
    COALESCE(s.dau, 0) AS dau,
    COALESCE(s.total_sessions, 0) AS total_sessions,
    COALESCE(e.hard_currency_spent, 0) AS hard_currency_spent,
    COALESCE(e.soft_currency_spent, 0) AS soft_currency_spent,
    COALESCE(e.total_transactions, 0) AS total_transactions
FROM daily_sessions AS s
FULL OUTER JOIN daily_economy AS e
    ON s.event_date = e.event_date
