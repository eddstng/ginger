CREATE TABLE order_item_customizations (
    order_item_id INT NOT NULL REFERENCES order_items(id) ON DELETE CASCADE,
    customization_id INT NOT NULL REFERENCES customizations(id),
    PRIMARY KEY (order_item_id, customization_id)
);

INSERT INTO order_item_customizations (order_item_id, customization_id) VALUES
(1, 1),
(1, 2),
(2, 1),
(3, 1);