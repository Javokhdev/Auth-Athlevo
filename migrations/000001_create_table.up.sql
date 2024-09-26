CREATE TYPE role_type AS ENUM ('admin', 'user', 'owner','coach');
CREATE TYPE gender_type AS ENUM ('male', 'female');

-- USER TABLE
CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    gym_id UUID NOT NULL,
    username VARCHAR(50) UNIQUE,
    email VARCHAR(100) UNIQUE,
    phone_number VARCHAR(20),
    password VARCHAR(255),
    full_name VARCHAR(100),
    date_of_birth DATE,
    role role_type NOT NULL DEFAULT 'user',
    gender gender_type NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

-- SETTING TABLE
CREATE TABLE IF NOT EXISTS settings (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    privacy_level VARCHAR(50) NOT NULL DEFAULT 'private',
    notification VARCHAR(30) NOT NULL DEFAULT 'on',
    language VARCHAR(255) NOT NULL DEFAULT 'en',
    theme VARCHAR(255) NOT NULL DEFAULT 'light',
    updated_at TIMESTAMP DEFAULT NOW()
);

-- TOKEN
CREATE TABLE IF NOT EXISTS tokens (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    token VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);





