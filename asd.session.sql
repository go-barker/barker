SELECT users.* FROM `users`
    LEFT OUTER JOIN deliveries
        ON deliveries.telegram_id = users.telegram_id
        AND deliveries.bot_id = users.bot_id
        AND deliveries.campaign_id = 4
    WHERE deliveries.telegram_id IS NULL AND users.deleted_at IS NULL AND users.bot_id = 3 LIMIT 1
