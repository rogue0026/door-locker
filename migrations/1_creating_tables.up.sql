CREATE TABLE IF NOT EXISTS categories (
    id int generated always as identity primary key,
    name varchar(100),
    image text
);

CREATE TABLE IF NOT EXISTS locks (
    part_number VARCHAR(30) PRIMARY KEY,
    title VARCHAR(100),
    image TEXT[],
    price INT,
    sale_price INT,
    equipment VARCHAR(256),
    colors VARCHAR(50)[],
    description VARCHAR(4096),
    category_id INT DEFAULT 0 REFERENCES categories(id) ON DELETE SET DEFAULT ON UPDATE CASCADE,
    card_memory INTEGER,
    material VARCHAR(40)[],
    has_mobile_application BOOLEAN,
    power_supply INT,
    size VARCHAR(50),
    weight INT,
    door_type VARCHAR(50)[],
    door_thickness_min INT,
    door_thickness_max INT,
    rating REAL,
    quantity INT
);

