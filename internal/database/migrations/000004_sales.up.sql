CREATE TABLE sales (
    id VARCHAR(36) PRIMARY KEY NOT NULL ,
    code VARCHAR(6) NOT NULL,
    seller_code  VARCHAR(36) NOT NULL ,
    client_phone VARCHAR(11) NOT NULL ,
    quantity SMALLINT NOT NULL,
    discount_code VARCHAR(6) NOT NULL ,
    is_reviewed BOOL DEFAULT false,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    FOREIGN KEY (client_phone) REFERENCES clients(phone),
    FOREIGN KEY (seller_code) REFERENCES sellers(code)
);

CREATE TABLE sale_books (
    sale_id VARCHAR(36) NOT NULL,
    book_id VARCHAR(36) NOT NULL ,
    PRIMARY KEY (sale_id, book_id),
    FOREIGN KEY (sale_id) REFERENCES sales(id) ON DELETE CASCADE,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
)