CREATE TABLE order_items (
    id SERIAL NOT NULL PRIMARY KEY,
    item_id INT NOT NULL REFERENCES items(id), -- FK reference to items table
    order_id INT NOT NULL REFERENCES orders(id), -- FK reference to orders table
    quantity INT NOT NULL DEFAULT 1,
    price DECIMAL(10, 2) NOT NULL -- Store the price of the item including customizations at the time of the order
);

INSERT INTO order_items 
(item_id, order_id, quantity, price) VALUES
(1,       1,        1,         4.5),
(2,       1,        1,         3),
(3,       2,        1,         5);
