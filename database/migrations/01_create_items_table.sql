CREATE TABLE items (
    id SERIAL PRIMARY KEY, 
    menu_id INTEGER DEFAULT 0, -- Multiple items can share the same menu_id. This may happen for items with variations.
    category_id INTEGER NOT NULL REFERENCES categories(id),
    price DECIMAL(10, 2) DEFAULT 0,
    name_eng VARCHAR(100) NOT NULL,
    name_oth VARCHAR(100) DEFAULT NULL, -- Item name in other language.
    special BOOLEAN DEFAULT FALSE, -- Automatic discounts will not be applied to items that are special.
    alcohol BOOLEAN DEFAULT FALSE, -- An alcohol tax (PST) will be applied to items that are alcoholic.
    custom BOOLEAN DEFAULT FALSE, -- Custom items may be created.
    variant VARCHAR(100) DEFAULT NULL, -- The variation string if applicable.
    variant_default BOOLEAN DEFAULT FALSE, -- Indicates whether this variation is the default variation if applicable.
    variant_price_charge DECIMAL(10,2) DEFAULT 0
);

INSERT INTO items 
(menu_id, category_id, price,     name_eng,           name_oth,    special,   alcohol,   custom,    variant,     variant_default) VALUES 
(1,       1,           2.50,      'Appetizer 1',      'App 1',     FALSE,     FALSE,     FALSE,      NULL,        FALSE         ),
(2,       2,           3.00,      'Soup 1',           'Soup 1',    FALSE,     FALSE,     FALSE,     'Small',      TRUE          ),
(2,       2,           5.00,      'Soup 1',           'Soup 1',    FALSE,     FALSE,     FALSE,     'Large',      FALSE          ),
(3,       3,           12.50,     'Vegetable 1',      'Veg 1',     FALSE,     FALSE,     FALSE,      NULL,        FALSE         ),
(5,       4,           7.50,      'Alcohol 1',        'Alc 1',     FALSE,     TRUE,      FALSE,      NULL,        FALSE         ),
(4,       5,           15.25,     'Combo 1',          'Comb 1',    FALSE,     FALSE,     FALSE,      NULL,        FALSE         );
