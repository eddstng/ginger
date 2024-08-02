CREATE TABLE items (
    id SERIAL PRIMARY KEY, 
    menu_id INTEGER DEFAULT 0, -- Multiple items can share the same menu_id. This may happen for items with variations.
    category_id INTEGER DEFAULT 1 REFERENCES categories(id),
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
(menu_id, category_id, price, name_eng,                     name_oth,     special,  alcohol,   custom,      variant,     variant_default, variant_price_charge) VALUES 
(1,       2,           5.99,  'Spring Rolls',               '春卷',       false,     false,     false,      '',          false,            0),
(2,       3,           4.99,  'Hot and Sour Soup',          '酸辣汤',      false,     false,     false,      'Small',     true,            0),
(3,       3,           4.99,  'Hot and Sour Soup',          '酸辣汤',      false,     false,     false,      'Large',     false,            4),
(4,       4,           6.99,  'Chicken Egg Foo Yung',       '雞芙蓉蛋',    false,     false,     false,      '',          false,            0),
(5,       5,           7.99,  'Stir-fried Bok Choy',        '炒青菜',      false,     false,     false,      '',          false,            0),
(6,       6,           12.99,  'Salt and Pepper Shrimp',    '椒鹽蝦',      false,     false,     false,      '',          false,            0),
(7,       7,           13.99,  'Stir-fried Scallops',       '炒帶子',      false,     false,     false,      '',          false,            0),
(8,       8,           14.99,  'Beef Hot Pot',              '牛肉煲',      false,     false,     false,      '',          false,            0),
(9,       9,           9.99,  'Sweet and Sour Pork',        '糖醋排骨',    false,     false,     false,      '',          false,            0),
(10,       10,          10.99,  'Beef with Broccoli',        '西蘭花牛肉',  false,     false,     false,      '',           false,            0),
(11,      11,          8.99,  'Kung Pao Chicken',           '宫保鸡丁',    false,     false,     false,      '',          false,            0),
(12,      12,          7.99,  'BBQ Pork Over Rice',         '叉燒飯',      false,     false,     false,      '',          false,            0),
(13,      13,          8.99,  'Yangzhou Fried Rice',        '扬州炒饭',    false,     false,     false,      '',           false,           0),
(14,      14,          9.99,  'Chicken Chow Mein',          '雞肉炒麵',    false,     false,     false,      '',          false,            0),
(15,      15,          6.99,  'Beef Noodle Soup',           '牛肉湯麵',    false,     false,     false,      '',           false,            0),
(16,      16,          5.99,  'Pork Congee',                '豬肉粥',      false,     false,     false,      '',          false,            0),
(17,      17,          15.99,  'General Tso''s Chicken',    '左宗棠雞',    false,     false,     false,      '',          false,            0),
(18,      18,          2.99,  'Local Beer',                 '本地啤酒',    false,     true,      false,      '',            false,           0);
