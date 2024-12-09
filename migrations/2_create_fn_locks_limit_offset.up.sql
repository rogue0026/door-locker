CREATE FUNCTION fn_locks_limit_offset(page_number INT, records_per_page INT)
    RETURNS TABLE(
                     part_number BIGINT,
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
    rows_to_skip INT;
BEGIN
    rows_to_skip = (page_number - 1) * records_per_page;

    RETURN QUERY
        SELECT
            locks.part_number,
            locks.title,
            locks.image_links,
            locks.price,
            locks.sale_price,
            locks.equipment,
            locks.colors,
            locks.description,
            categories.name as category_name,
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
                 JOIN categories on locks.category_id = categories.id
        LIMIT records_per_page
            OFFSET rows_to_skip;
END;
$$;