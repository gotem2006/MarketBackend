CREATE TABLE categories
(
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(50) NOT NULL
);

CREATE TABLE products
(
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(50) NOT NULL,
    product_price  DECIMAL NOT NULL,
    category_id BIGINT,
    CONSTRAINT fk_category_id
        FOREIGN KEY(category_id)
        REFERENCES categories(category_id)
);


CREATE TABLE attributes
(
    attribute_name VARCHAR(50) NOT NULL,
    attribute_value TEXT NOT NULL,
    product_id BIGINT,
    CONSTRAINT fk_product_id
        FOREIGN KEY(product_id)
        REFERENCES products(product_id)
)

