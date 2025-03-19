CREATE TABLE reviews (
    id VARCHAR(36) PRIMARY KEY NOT NULL ,
    sale_id VARCHAR(36) NOT NULL ,
    rating FLOAT NOT NULL ,
    comment VARCHAR(255) NOT NULL ,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (sale_id) REFERENCES sales(id)
)