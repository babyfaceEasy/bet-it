-- Add UUID extension so that postgres can create it for you
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- set timezone
SET TIMEZONE="Africa/Lagos";

-- create admins table
CREATE TABLE admins (
    id uuid DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE,
    full_name VARCHAR (200) NOT NULL,
    email VARCHAR (200) NOT NULL,
    password VARCHAR (200) NOT NULL,
    super_admin BOOLEAN DEFAULT false,

    UNIQUE (email),
    PRIMARY KEY(id)
);