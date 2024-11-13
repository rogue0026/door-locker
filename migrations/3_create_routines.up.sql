BEGIN;

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
            --door_locks.image,
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

COMMIT;