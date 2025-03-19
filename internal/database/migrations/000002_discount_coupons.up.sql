CREATE TABLE discount_coupons (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    code VARCHAR(6) NOT NULL,
    discount_percent FLOAT NOT NULL NULL,
    expiration_date TIMESTAMP NOT NULL,
    usage_limit SMALLINT NOT NULL,
    used_count SMALLINT NOT NULL ,
    active BOOL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
)