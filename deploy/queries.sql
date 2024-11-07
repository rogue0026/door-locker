CREATE TABLE lock_colors(
                            id SERIAL PRIMARY KEY,
                            name VARCHAR(30) NOT NULL
);

CREATE TABLE lock_categories(
                                id SERIAL PRIMARY KEY,
                                name VARCHAR(30) NOT NULL
);

CREATE TABLE lock_materials(
                               id SERIAL PRIMARY KEY,
                               name VARCHAR(30) NOT NULL
)

CREATE TABLE door_locks (
                            part_number VARCHAR(30) PRIMARY KEY,
                            title VARCHAR(100) NOT NULL,
                            price REAL NOT NULL,
                            sale_price REAL NOT NULL,
                            equipment VARCHAR(256) NOT NULL,
--     color_id INT NOT NULL REFERENCES
                            description VARCHAR(4096) NOT NULL,
--     category_id
                            card_memory INTEGER NOT NULL,
--     material_id
                            has_mobile_application BOOLEAN NOT NULL,
                            power_supply VARCHAR(50) NOT NULL,
                            size VARCHAR(50) NOT NULL,
                            weight INTEGER NOT NULL,
                            door_types_id INTEGER[] NOT NULL,
                            door_thickness_min INTEGER NOT NULL,
                            door_thickness_max INTEGER NOT NULL,
                            rating REAL NOT NULL,
                            quantity INTEGER NOT NULL
);