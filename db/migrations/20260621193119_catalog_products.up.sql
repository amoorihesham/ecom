CREATE TABLE categories (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        TEXT NOT NULL
);

CREATE TABLE products (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            TEXT NOT NULL,
    category_id     UUID NULL,
    status          TEXT NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_products_category
        FOREIGN KEY (category_id)
        REFERENCES categories(id)
        ON DELETE CASCADE,

    CONSTRAINT chk_products_status
        CHECK (
            status IN (
                'draft',
                'published',
                'unpublished',
                'discontinued'
            )
        )
);

CREATE TABLE variants (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    product_id          UUID NOT NULL,

    sku                 TEXT NOT NULL,

    price_amount        NUMERIC(18,4) NOT NULL,
    price_currency      CHAR(3) NOT NULL,

    CONSTRAINT fk_variants_product
        FOREIGN KEY (product_id)
        REFERENCES products(id)
        ON DELETE CASCADE,

    CONSTRAINT uq_variants_product_sku
        UNIQUE (product_id, sku)
);

CREATE TABLE product_images (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    product_id      UUID NOT NULL,

    file_id         UUID NOT NULL,

    position        INT NOT NULL DEFAULT 0,

    CONSTRAINT fk_product_images_product
        FOREIGN KEY (product_id)
        REFERENCES products(id)
        ON DELETE CASCADE
);

CREATE TABLE product_attributes (
    product_id      UUID NOT NULL DEFAULT gen_random_uuid(),

    attribute_key   TEXT NOT NULL,
    value           TEXT NOT NULL,

    PRIMARY KEY (product_id, attribute_key),

    CONSTRAINT fk_product_attributes_product
        FOREIGN KEY (product_id)
        REFERENCES products(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_products_category_id
    ON products(category_id);

CREATE INDEX idx_variants_product_id
    ON variants(product_id);

CREATE INDEX idx_product_images_product_id
    ON product_images(product_id);

CREATE INDEX idx_product_attributes_product_id
    ON product_attributes(product_id);