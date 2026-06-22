-- USERS
CREATE UNIQUE INDEX idx_users_email ON users (email);

-- REFRESH TOKENS
CREATE INDEX idx_refresh_tokens_user_id ON refresh_tokens (user_id);

CREATE UNIQUE INDEX idx_refresh_tokens_token_hash ON refresh_tokens (token_hash);