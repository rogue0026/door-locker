CREATE TABLE IF NOT EXISTS categories (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100),
    image_link TEXT
);

CREATE TABLE IF NOT EXISTS locks (
    part_number BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    image_links TEXT[] NOT NULL,
    price INT NOT NULL CHECK (price > 0),
    sale_price INT NOT NULL CHECK (sale_price >= 0),
    equipment VARCHAR(512) NOT NULL,
    colors VARCHAR(100)[] NOT NULL,
    description VARCHAR(8192) NOT NULL,
    category_id INT DEFAULT 0 REFERENCES categories(id) ON DELETE SET DEFAULT ON UPDATE CASCADE,
    card_memory INT NOT NULL CHECK (card_memory >= 0),
    material VARCHAR(40)[] NOT NULL,
    has_mobile_application BOOLEAN NOT NULL,
    power_supply INT NOT NULL CHECK (power_supply >= 0),
    size VARCHAR(50) NOT NULL,
    weight INT NOT NULL,
    door_type VARCHAR(50)[] NOT NULL,
    door_thickness_min INT NOT NULL CHECK(door_thickness_min > 0),
    door_thickness_max INT NOT NULL CHECK(door_thickness_max > 0),
    rating REAL NOT NULL,
    quantity INT NOT NULL CHECK (quantity >= 0)
);
