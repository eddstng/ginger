CREATE TABLE orders_items (
    id SERIAL NOT NULL PRIMARY KEY,
    item_id INT NOT NULL REFERENCES items(id), -- FK reference to items table
    order_id INT NOT NULL REFERENCES orders(id), -- FK reference to orders table
    quantity INT NOT NULL DEFAULT 1,
    customizations jsonb DEFAULT NULL, -- Store customizations here in JSON format
    price DECIMAL(10, 2) NOT NULL -- Store the price of the item including customizations at the time of the order
);

INSERT INTO orders_items 
(item_id, order_id, quantity, customizations, price) VALUES
(1,       1,        1,        '[{"name_eng": "add bb sauce", "name_oth": "gaseejup", "price": 2.00}, {"name_eng": "no msg", "name_oth": "jaojing", "price": 0.00}]', 4.5),
(2,       1,        1,        '[{"name_eng": "less oil", "name_oth": "sewyao", "price": 0.00}]', 3                                                                    ),
(3,       2,        1,        NULL, 5                                                                                                                                 );
