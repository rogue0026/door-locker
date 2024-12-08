CREATE PROCEDURE save_lock(
    title VARCHAR(150),
    image_links TEXT[],
    price INT,
    sale_price INT,
    equipment VARCHAR(512),
    colors VARCHAR(100)[],
    description VARCHAR(8192),
    category VARCHAR(100),
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
    __category_id INT;
BEGIN

    SELECT
        id
    INTO __category_id
    FROM categories
    WHERE name = category;
    IF NOT FOUND THEN
        INSERT INTO categories(name, image_link) VALUES (category, '{"not defined"}');
        SELECT
            id
        INTO __category_id
        FROM categories
        WHERE name = category;
    END IF;

    INSERT INTO locks (
        title,
        image_links,
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
               save_lock.title,
               save_lock.image_links,
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