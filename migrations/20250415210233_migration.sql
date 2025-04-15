-- +goose Up
CREATE TYPE city_enum AS ENUM ('Москва', 'Санкт-Петербург', 'Казань');
CREATE TYPE reception_status_enum AS ENUM ('in_progress', 'close');
CREATE TYPE product_type_enum AS ENUM ('электроника', 'одежда', 'обувь');
CREATE TYPE user_role_enum AS ENUM ('employee', 'moderator');

CREATE TABLE users(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role user_role_enum NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE pvz(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    city city_enum NOT NULL,
    registration_date TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE receptions(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    pvz_id UUID NOT NULL REFERENCES pvz(id) ON DELETE CASCADE,
    date_time TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    status reception_status_enum NOT NULL DEFAULT 'in_progress'
);

CREATE TABLE products(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    reception_id UUID NOT NULL REFERENCES receptions(id) ON DELETE CASCADE,
    date_time TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    type product_type_enum NOT NULL
);

-- +goose Down
DROP TYPE city_enum;
DROP TYPE reception_status_enum;
DROP TYPE product_type_enum;
DROP TYPE user_role_enum;
DROP TABLE users;
DROP TABLE pvz;
DROP TABLE receptions;
DROP TABLE products;
