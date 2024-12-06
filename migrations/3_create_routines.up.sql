CREATE FUNCTION fn_locks_ordered_by_rating(num_of_records INT)
    RETURNS TABLE (
                      part_number VARCHAR(30),
                      title VARCHAR(100),
                      image TEXT[],
                      price INT,
                      sale_price INT,
                      equipment VARCHAR(256),
                      colors VARCHAR(50)[],
                      description VARCHAR(4096),
                      category VARCHAR(50),
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
                      quantity INT)
LANGUAGE plpgsql
AS
$$
    DECLARE

    BEGIN
        RETURN QUERY
            SELECT
                locks.part_number,
                locks.title,
                locks.image,
                locks.price,
                locks.sale_price,
                locks.equipment,
                locks.colors,
                locks.description,
                c.name as category_name,
                locks.card_memory,
                locks.material,
                locks.has_mobile_application,
                locks.power_supply,
                locks.size,
                locks.weight,
                locks.door_type,
                locks.door_thickness_min,
                locks.door_thickness_max,
                locks.rating,
                locks.quantity
            FROM locks
            JOIN categories c on locks.category_id = c.id
            ORDER BY rating
            LIMIT num_of_records;
    END;
$$;

CREATE PROCEDURE save_lock(
    part_number VARCHAR(30),
    title VARCHAR(100),
    image TEXT[],
    price INT,
    sale_price INT,
    equipment VARCHAR(256),
    colors VARCHAR(50)[],
    description VARCHAR(4096),
    category VARCHAR(50),
    card_memory INT,
    material VARCHAR(40)[],
    has_mobile_application BOOLEAN,
    power_supply INT,
    size VARCHAR(50),
    weight INT,
    door_type VARCHAR(50)[],
    door_thickness_min INT,
    door_thickness_max INT,
    rating REAL,
    quantity INT)
LANGUAGE plpgsql
AS
    $$
        DECLARE
            __category_id INT = 0;
        BEGIN

            SELECT
                id
            INTO __category_id
            FROM categories
            WHERE name = category;

            IF __category_id = 0 THEN
                INSERT INTO categories(name, image) VALUES (category, '{"not defined"}');
            END IF;

            INSERT INTO locks (
                part_number,
                title,
                image,
                price,
                sale_price,
                equipment,
                colors,
                description,
                category_id,
                card_memory,
                material,
                has_mobile_application,
                power_supply,
                size,
                weight,
                door_type,
                door_thickness_min,
                door_thickness_max,
                rating,
                quantity)
            VALUES (
                save_lock.part_number,
                save_lock.title,
                save_lock.image,
                save_lock.price,
                save_lock.sale_price,
                save_lock.equipment,
                save_lock.colors,
                save_lock.description,
                __category_id,
                save_lock.card_memory,
                save_lock.material,
                save_lock.has_mobile_application,
                save_lock.power_supply,
                save_lock.size,
                save_lock.weight,
                save_lock.door_type,
                save_lock.door_thickness_min,
                save_lock.door_thickness_max,
                save_lock.rating,
                save_lock.quantity
            );
        END;
$$;

create procedure delete_door_lock_by_partnumber(partnumber varchar(30))
    language plpgsql
as
$$
begin
    delete
    from locks
    where part_number = partnumber;
end;
$$;