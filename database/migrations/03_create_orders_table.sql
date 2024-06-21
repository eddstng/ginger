CREATE TYPE order_type AS ENUM ('IN', 'OUT', 'DELIVERY');

CREATE TABLE orders (
    id SERIAL NOT NULL PRIMARY KEY,
    subtotal DECIMAL(10, 2) NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    gst DECIMAL(10, 2) NOT NULL DEFAULT 0,
    pst DECIMAL(10, 2) NOT NULL DEFAULT 0,
    discount DECIMAL(10, 2) NOT NULL DEFAULT 0,
    timestamp TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    type order_type NOT NULL,
    void BOOLEAN NOT NULL DEFAULT FALSE,
    paid BOOLEAN NOT NULL DEFAULT FALSE,
    customizations jsonb DEFAULT NULL,
    -- customizations_price DECIMAL(10, 2) NOT NULL DEFAULT 0, 
        -- No longer need to use this field - will be handled in the query. 
        -- SELECT 
        --     id,
        --     COALESCE(
        --         (SELECT string_agg(value->>'name_eng', ', ')
        --          FROM jsonb_array_elements(customizations) AS value), '') AS customization_names,
        --     COALESCE(
        --         (SELECT SUM((value->>'price')::numeric)
        --          FROM jsonb_array_elements(customizations) AS value), 0) AS total_customization_price
        -- FROM 
        --     orders
    customer_id INT REFERENCES customers(id)
);

INSERT INTO orders
(subtotal,    total,   gst,   pst,    discount,   type,       customizations,                                                           customer_id ) VALUES
(7.50,        7.87,    0.37,  0.00,   0.00,       'IN',       NULL,                                                                     1           ),
(6.00,        6.30,    0.30,  0.00,   0.00,       'OUT',      '[{"name_eng": "add bb sauce", "name_oth": "gaseejup", "price": 1.00}]',  2           );