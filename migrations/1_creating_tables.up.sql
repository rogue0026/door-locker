BEGIN;

CREATE TABLE IF NOT EXISTS lock_colors(
                            id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                            name VARCHAR(30) NOT NULL
);

CREATE TABLE IF NOT EXISTS lock_categories(
                                id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                name VARCHAR(30) NOT NULL,
                                image SMALLINT[] NOT NULL DEFAULT '{1,0,1,0,1,0,1,0,1}'
);

CREATE TABLE IF NOT EXISTS lock_materials(
                               id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                               name VARCHAR(30) NOT NULL
);

CREATE TABLE IF NOT EXISTS door_locks (
                            part_number VARCHAR(30) PRIMARY KEY,
                            title VARCHAR(100) NOT NULL,
                            image SMALLINT[] NOT NULL DEFAULT '{1,0,1,0,1,0,1,0,1}',
                            price REAL NOT NULL,
                            sale_price REAL NOT NULL,
                            equipment VARCHAR(256) NOT NULL,
                            color_id INT NOT NULL REFERENCES lock_colors(id) ON UPDATE CASCADE ON DELETE CASCADE,
                            description VARCHAR(4096) NOT NULL,
                            category_id INT REFERENCES lock_categories(id) ON UPDATE CASCADE ON DELETE CASCADE,
                            card_memory INTEGER NOT NULL,
                            material_id INT NOT NULL REFERENCES lock_materials(id) ON UPDATE CASCADE ON DELETE CASCADE,
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

COMMIT;