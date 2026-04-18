-- The test will fail if it finds even one row with a negative sum.
SELECT
    event_id,
    user_id,
    amount
FROM {{ ref('stg_game_economy') }}
WHERE amount < 0