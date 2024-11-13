CREATE TABLE lock_colors(
                            id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                            name VARCHAR(30) NOT NULL
);

INSERT INTO lock_colors (name) VALUES ('Черный'), ('Золотой'),('Серебристный'),('Серый'), ('Коричневый');

CREATE TABLE lock_categories(
                                id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                name VARCHAR(30) NOT NULL
);

INSERT INTO lock_categories(name) VALUES ('Для дома'), ('Для квартиры'), ('Для гаража'), ('Для сарая'), ('Для машины'), ('Для Маги');
CREATE TABLE lock_materials(
                               id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                               name VARCHAR(30) NOT NULL
);

INSERT INTO lock_materials(name) VALUES ('Металл'), ('Пластик'), ('Металл/пластик'), ('Дерево'), ('Жопа дракона'), ('Карбон');

CREATE TABLE door_locks (
                            part_number VARCHAR(30) PRIMARY KEY,
                            title VARCHAR(100) NOT NULL,
                            image SMALLINT[],
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

INSERT INTO door_locks
    (part_number,
     title,
     price,
     sale_price,
     equipment,
     color_id,
     description,
     category_id,
     card_memory,
     material_id,
     has_mobile_application,
     power_supply,
     size,
     weight,
     door_types_id,
     door_thickness_min,
     door_thickness_max,
     rating,
     quantity)
VALUES
    ('DL-0001', 'Smart Lock A1', 200.99, 180.99, 'Fingerprint, Keypad', 1, 'A modern smart lock with mobile app access.', 1, 200, 2, TRUE, '4 AA Batteries', '80x60x40 mm', 500, ARRAY[1,2], 35, 50, 4.7, 100),
    ('DL-0002', 'Biometric Lock B2', 250.50, 230.00, 'Card, Fingerprint', 2, 'High-security biometric lock.', 2, 300, 3, FALSE, 'Rechargeable Battery', '85x65x45 mm', 550, ARRAY[1,3], 30, 45, 4.8, 75),
    ('DL-0003', 'Classic Lock C1', 80.00, 75.00, 'Key, Card', 3, 'Classic lock with additional card functionality.', 3, 100, 1, FALSE, 'CR123A Battery', '70x55x35 mm', 450, ARRAY[2,4], 28, 40, 4.2, 120),
    ('DL-0004', 'Electronic Lock D3', 150.99, 140.99, 'Remote, Card', 4, 'Electronic lock with remote access.', 1, 150, 2, TRUE, '4 AA Batteries', '78x58x38 mm', 520, ARRAY[1,2,3], 33, 50, 4.5, 90),
    ('DL-0005', 'Fingerprint Lock E5', 210.00, 195.00, 'Fingerprint, Bluetooth', 2, 'Lock with fingerprint and Bluetooth control.', 2, 250, 4, TRUE, 'USB-C Rechargeable', '82x62x42 mm', 540, ARRAY[1,3,4], 32, 48, 4.6, 60),
    ('DL-0006', 'Secure Lock F4', 190.75, 180.75, 'Remote, Key', 5, 'Secure lock with dual access options.', 3, 0, 3, FALSE, 'Lithium Battery', '76x56x36 mm', 470, ARRAY[1,2], 30, 45, 4.3, 85),
    ('DL-0007', 'Advanced Lock G2', 280.30, 260.20, 'Fingerprint, Card', 4, 'High-end lock for premium security.', 2, 400, 2, TRUE, '4 AA Batteries', '90x70x50 mm', 600, ARRAY[1,4], 34, 52, 4.9, 50),
    ('DL-0008', 'Bluetooth Lock H3', 180.50, 165.50, 'Bluetooth, Key', 1, 'Smart lock with Bluetooth integration.', 3, 150, 1, TRUE, 'Rechargeable Battery', '75x55x35 mm', 490, ARRAY[2,3,4], 28, 42, 4.4, 70),
    ('DL-0009', 'Digital Lock I2', 220.00, 205.00, 'Fingerprint, Key', 2, 'Lock with advanced digital features.', 1, 300, 2, TRUE, 'Lithium Battery', '88x68x48 mm', 530, ARRAY[1,2], 31, 48, 4.6, 65),
    ('DL-0010', 'Traditional Lock J1', 75.50, 70.00, 'Key only', 3, 'Reliable traditional lock.', 4, 0, 1, FALSE, 'No Power Required', '60x50x30 mm', 420, ARRAY[3,4], 25, 35, 4.1, 150),
    ('DL-0011', 'Biometric Lock K5', 290.40, 275.30, 'Fingerprint, Card, Remote', 5, 'High-security biometric lock with remote.', 1, 350, 3, TRUE, 'USB-C Rechargeable', '95x75x55 mm', 580, ARRAY[1,2,4], 35, 55, 4.8, 45),
    ('DL-0012', 'Smart Lock L6', 195.00, 180.00, 'Bluetooth, Fingerprint', 2, 'Smart lock with multiple access points.', 2, 200, 4, TRUE, 'Lithium Battery', '85x65x45 mm', 510, ARRAY[2,3], 30, 45, 4.5, 80),
    ('DL-0013', 'Premium Lock M3', 270.75, 250.75, 'Card, Fingerprint, Keypad', 3, 'Premium lock with enhanced features.', 2, 250, 3, TRUE, '4 AA Batteries', '82x62x42 mm', 540, ARRAY[1,3,4], 33, 50, 4.9, 55),
    ('DL-0014', 'Eco Lock N7', 95.50, 85.00, 'Key only', 4, 'Environment-friendly eco-lock.', 4, 0, 1, FALSE, 'No Power Required', '65x55x35 mm', 460, ARRAY[3,4], 25, 35, 4.2, 130),
    ('DL-0015', 'Smart Entry P2', 160.00, 145.00, 'Remote, Card', 1, 'Smart entry lock with remote access.', 1, 100, 2, TRUE, 'Rechargeable Battery', '77x57x37 mm', 500, ARRAY[1,2,3], 30, 47, 4.4, 95),
    ('DL-0016', 'Biometric Plus Q4', 275.99, 260.99, 'Fingerprint, Bluetooth', 2, 'High-end biometric lock.', 3, 350, 3, TRUE, 'USB-C Rechargeable', '88x68x48 mm', 550, ARRAY[1,3,4], 32, 50, 4.7, 50),
    ('DL-0017', 'Standard Lock R5', 105.00, 95.00, 'Key only', 5, 'Standard lock for all types of doors.', 4, 0, 1, FALSE, 'No Power Required', '68x58x38 mm', 430, ARRAY[3,4], 26, 36, 4.3, 140),
    ('DL-0018', 'Electronic Lock S1', 230.00, 210.00, 'Fingerprint, Remote', 3, 'Electronic lock with fingerprint option.', 2, 300, 2, TRUE, '4 AA Batteries', '84x64x44 mm', 510, ARRAY[1,2], 30, 48, 4.6, 65),
    ('DL-0019', 'Advanced Biometric T6', 300.99, 280.99, 'Fingerprint, Keypad', 4, 'Advanced biometric lock with keypad.', 1, 400, 4, TRUE, 'Rechargeable Battery', '92x72x52 mm', 600, ARRAY[1,3], 36, 54, 4.9, 40),
    ('DL-0020', 'Compact Lock U2', 125.00, 115.00, 'Key only', 3, 'Compact lock for residential doors.', 4, 0, 1, FALSE, 'No Power Required', '65x55x35 mm', 420, ARRAY[2,4], 28, 38, 4.2, 110);

CREATE FUNCTION fn_locks_limit_offset(page_number INT, records_per_page INT)
    RETURNS TABLE(
                     part_number VARCHAR(30),
                     title VARCHAR(100),
                     image SMALLINT[],
                     price REAL,
                     sale_price REAL,
                     equipment VARCHAR(256),
                     color_id INTEGER,
                     description VARCHAR(4096),
                     category_id INTEGER,
                     card_memory INTEGER,
                     material_id INTEGER,
                     has_mobile_application BOOLEAN,
                     power_supply VARCHAR(50),
                     size VARCHAR(50),
                     weight INTEGER,
                     door_types_id INTEGER[],
                     door_thickness_min INTEGER,
                     door_thickness_max INTEGER,
                     rating REAL,
                     quantity INTEGER)
    LANGUAGE plpgsql
AS
$$
DECLARE
    rows_to_skip INT;
BEGIN
    rows_to_skip = (page_number - 1) * records_per_page;
    RETURN QUERY
        SELECT
            door_locks.part_number,
            door_locks.title,
            door_locks.image,
            door_locks.price,
            door_locks.sale_price,
            door_locks.equipment,
            door_locks.color_id,
            door_locks.description,
            door_locks.category_id,
            door_locks.card_memory,
            door_locks.material_id,
            door_locks.has_mobile_application,
            door_locks.power_supply,
            door_locks.size,
            door_locks.weight,
            door_locks.door_types_id,
            door_locks.door_thickness_min,
            door_locks.door_thickness_max,
            door_locks.rating,
            door_locks.quantity
        FROM door_locks
        LIMIT records_per_page
            OFFSET rows_to_skip;
END;
$$;

create procedure save_door_lock(IN part_number character varying, IN title character varying, IN price real, IN sale_price real, IN equipment character varying, IN color_id integer, IN description character varying, IN category_id integer, IN card_memory integer, IN material_id integer, IN has_mobile_application boolean, IN power_supply character varying, IN size character varying, IN weight integer, IN door_types_id integer[], IN door_thickness_min integer, IN door_thickness_max integer, IN rating real, IN quantity integer)
    language plpgsql
as
$$
begin
    insert into door_locks(
        part_number,
        title,
        --image,
        price,
        sale_price,
        equipment,
        color_id,
        description,
        category_id,
        card_memory,
        material_id,
        has_mobile_application,
        power_supply,
        size,
        weight,
        door_types_id,
        door_thickness_min,
        door_thickness_max,
        rating,
        quantity)
    values (
               part_number,
               title,
               --image,
               price,
               sale_price,
               equipment,
               color_id,
               description,
               category_id,
               card_memory,
               material_id,
               has_mobile_application,
               power_supply,
               size,
               weight,
               door_types_id,
               door_thickness_min,
               door_thickness_max,
               rating,
               quantity);
end;
$$;


create procedure delete_door_lock_by_partnumber(partnumber varchar(30))
    language plpgsql
as
$$
begin
    delete
    from door_locks
    where part_number = partnumber;
end;
$$;