-- name: CreateUserMarketProvider :exec
INSERT INTO user_market_providers (
  user_uuid, provider
) VALUES ($1, $2);

-- name: GetUserMarketProvider :one
SELECT * FROM user_market_providers WHERE user_uuid = $1;
