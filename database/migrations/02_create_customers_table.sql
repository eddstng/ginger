CREATE TABLE customers (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(30),
    phone VARCHAR(15),
    unit_number VARCHAR(15), -- expecting values like bsm or upstairs
    street_number VARCHAR(15),
    street_name VARCHAR(30),
    -- city VARCHAR(15), -- removing due to rare use
    buzzer_number VARCHAR(15), -- expecting values like bsm or upstairs or A12
    -- address TEXT, -- the server/sql query will build this for the response
      -- CONCAT_WS(' ', COALESCE(unit_number, ''), street_number, street_name, COALESCE(buzzer_number, '')) AS full_address
    note TEXT
);

INSERT INTO customers 
(name,                 phone,           unit_number,  street_number, street_name,   buzzer_number,  note       ) VALUES 
('John Doe',           '604-123-1234',  '',         '5555',        'Powel St',    '',           ''       ),
('Christine StClaire', '123-456-7890',  'A12',        '123',         'Maple St',    'A12',          'good tips'),
('David Hogan',        '778-123-1234',  'BSM',         '5555',        'Powel St',    '',           ''       );