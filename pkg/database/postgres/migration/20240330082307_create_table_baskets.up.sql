CREATE TABLE baskets
(
    basket_id SERIAL PRIMARY KEY,
    quantity INTEGER,
    session_key VARCHAR(100),
    product_id BIGINT,
    CONSTRAINT fk_product_id
        FOREIGN KEY(product_id)
        REFERENCES products(product_id)
)
