CREATE TABLE books (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    title VARCHAR(255) NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    author VARCHAR(100) NOT NULL,
    publisher VARCHAR(100) NOT NULL,
    isbn VARCHAR(100) NOT NULL,
    price FLOAT NOT NULL,
    stock SMALLINT NOT NULL,
    category VARCHAR(36) NOT NULL,
    published_year SMALLINT,
    description VARCHAR(150) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
)