CREATE PROCEDURE add_account(
    login VARCHAR(30),
    password_hash VARCHAR(100),
    status VARCHAR(5),
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    birth_date DATE,
    phone_mobile VARCHAR(20),
    email VARCHAR(50))
LANGUAGE plpgsql
AS
$$
    BEGIN
        INSERT INTO accounts(
                             login,
                             password_hash,
                             status,
                             first_name,
                             last_name,
                             birth_date,
                             phone_mobile,
                             email)
        VALUES(login,
               password_hash,
               status,
               first_name,
               last_name,
               birth_date,
               phone_mobile,
               email);
    END;
$$;