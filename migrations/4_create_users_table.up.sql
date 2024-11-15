CREATE TABLE IF NOT EXISTS accounts(
    user_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    login VARCHAR(30) NOT NULL,
    password_hash VARCHAR(100) NOT NULL,
    status VARCHAR(5) NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    birth_date DATE NOT NULL,
    phone_mobile VARCHAR(20) NOT NULL,
    email VARCHAR(50) NOT NULL
);