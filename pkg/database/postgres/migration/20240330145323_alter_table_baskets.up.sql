ALTER TABLE baskets 
ADD CONSTRAINT session_product_unique 
UNIQUE (session_key, product_id);
