CREATE TYPE user_role AS ENUM ('user', 'seller', 'admin');

CREATE TABLE
    users (
        id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        public_id UUID NOT NULL DEFAULT gen_random_uuid (),
        email VARCHAR(255) NOT NULL UNIQUE,
        password_hash TEXT NOT NULL,
        full_name VARCHAR(100),
        role user_role NOT NULL DEFAULT 'user',
        is_active BOOLEAN NOT NULL DEFAULT TRUE,
        created_at TIMESTAMP NOT NULL DEFAULT NOW (),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW (),
        UNIQUE (public_id)
    );

CREATE TABLE
    refresh_tokens (
        id UUID PRIMARY KEY,
        user_id BIGINT NOT NULL,
        token_hash TEXT NOT NULL,
        expires_at TIMESTAMP NOT NULL,
        revoked BOOLEAN NOT NULL DEFAULT FALSE,
        created_at TIMESTAMP NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users (id)
    );