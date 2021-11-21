CREATE TABLE user_market_providers (
    user_uuid UUID PRIMARY KEY,
    provider VARCHAR(32) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
