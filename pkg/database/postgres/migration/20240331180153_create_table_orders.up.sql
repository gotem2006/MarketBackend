CREATE TABLE orders
(
    order_id bigserial PRIMARY KEY NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    delivary_address text NOT NULL,
    username text NOT NULL,
    phone text NOT NULL,
    email text NOT NULL,
    products jsonb NOT NULL
)