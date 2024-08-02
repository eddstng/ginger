CREATE TABLE categories (
    id SERIAL NOT NULL PRIMARY KEY, 
    name_eng VARCHAR(100) NOT NULL UNIQUE, 
    name_oth VARCHAR(100) NOT NULL UNIQUE, 
    description TEXT DEFAULT NULL
);

INSERT INTO categories 
(name_eng,             name_oth) VALUES 
('Custom',            'Custom'),
('Appetizers',         '頭 盤'),
('Soups',              '湯 羹 類'),
('Egg Foo Yung',       '芙 蓉 蛋 類'),
('Vegetables',         '蔬 菜 豆 腐 類'),
('Seafood',            '海 鮮 類'),
('Oysters/Scallops',   '生 蠔 帶 子 類'),
('Hot Pot',            '煲 仔 類'),
('Pork',               '豬 肉 類'),
('Beef',               '牛 肉 類'),
('Chicken',            '雞 肉 類'),
('Over Rice',          '碟 飯 類'),
('Fried Rice',         '炒 飯 類'),
('Chow Mein',          '炒 粉 麵 類'),
('Noodle Soup',        '湯 粉 麵 類'),
('Congee',             '粥 類'),
('Specials',           'Specials'),
('Drinks',             'Drinks');