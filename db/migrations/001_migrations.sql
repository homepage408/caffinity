CREATE TABLE
    cafes (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        tagline TEXT,
        address TEXT,
        city VARCHAR(100),
        latitude DECIMAL(10, 6),
        longitude DECIMAL(10, 6),
        open_hours VARCHAR(100),
        phone VARCHAR(20),
        instagram VARCHAR(100),
        rating DECIMAL(2, 1) DEFAULT 0,
        reviews INT DEFAULT 0,
        price_level INT CHECK (price_level BETWEEN 1 AND 5),
        hero_img VARCHAR(10),
        vibe_emoji VARCHAR(10),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    cafe_approvals (
        id SERIAL PRIMARY KEY,
        cafe_id INT REFERENCES cafes (id) ON DELETE CASCADE,
        approved BOOLEAN DEFAULT FALSE,
        approved_reason TEXT,
        approved_at TIMESTAMP,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    menus (
        id SERIAL PRIMARY KEY,
        cafe_id INT REFERENCES cafes (id) ON DELETE CASCADE NOT NULL,
        name VARCHAR(255) NOT NULL,
        price INT NOT NULL,
        strength INT CHECK (strength BETWEEN 1 AND 5),
        is_safe BOOLEAN DEFAULT TRUE,
        description TEXT,
        image VARCHAR(10),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    photos (
        id SERIAL PRIMARY KEY,
        is_primary BOOLEAN DEFAULT FALSE,
        position INT DEFAULT 0,
        cafe_id INT NOT NULL REFERENCES cafes (id) ON DELETE CASCADE,
        url TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    tags (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) UNIQUE NOT NULL
    );

CREATE TABLE
    cafe_tags (
        cafe_id INT REFERENCES cafes (id) ON DELETE CASCADE,
        tag_id INT REFERENCES tags (id) ON DELETE CASCADE,
        PRIMARY KEY (cafe_id, tag_id)
    );

CREATE TABLE
    facilities (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) UNIQUE NOT NULL
    );

CREATE TABLE
    cafe_facilities (
        cafe_id INT REFERENCES cafes (id) ON DELETE CASCADE,
        facility_id INT REFERENCES facilities (id) ON DELETE CASCADE,
        PRIMARY KEY (cafe_id, facility_id)
    );

CREATE INDEX idx_cafes_city ON cafes (city);

CREATE INDEX idx_menus_cafe_id ON menus (cafe_id);

CREATE INDEX idx_photos_cafe_id ON photos (cafe_id);

CREATE INDEX idx_cafe_facilities_cafe_id ON cafe_facilities (cafe_id);