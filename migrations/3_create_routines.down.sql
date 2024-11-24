BEGIN;

DROP FUNCTION IF EXISTS fn_locks_limit_offset(page_number INT, records_per_page INT);

DROP FUNCTION IF EXISTS fn_locks_ordered_by_rating(num_of_records INT);

DROP PROCEDURE IF EXISTS save_door_lock(IN part_number character varying, IN title character varying, IN price real, IN sale_price real, IN equipment character varying, IN color_id integer, IN description character varying, IN category_id integer, IN card_memory integer, IN material_id integer, IN has_mobile_application boolean, IN power_supply character varying, IN size character varying, IN weight integer, IN door_types_id integer[], IN door_thickness_min integer, IN door_thickness_max integer, IN rating real, IN quantity integer);

DROP PROCEDURE IF EXISTS delete_door_lock_by_partnumber(partnumber varchar(30));

COMMIT;