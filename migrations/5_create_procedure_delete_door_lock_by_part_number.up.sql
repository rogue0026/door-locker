CREATE PROCEDURE delete_door_lock_by_part_number(_part_number VARCHAR(30))
    LANGUAGE plpgsql
AS
$$
BEGIN
    DELETE
    FROM locks
    WHERE part_number = _part_number;
END;
$$;