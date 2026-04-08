CREATE TABLE cafes (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    city TEXT NOT NULL,
    address TEXT,
    description TEXT,
    rating FLOAT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE menus (
    id SERIAL PRIMARY KEY,
    cafe_id INT NOT NULL REFERENCES cafes(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    price INT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE photos (
    id SERIAL PRIMARY KEY,
    cafe_id INT NOT NULL REFERENCES cafes(id) ON DELETE CASCADE,
    url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    cafe_id INT NOT NULL REFERENCES cafes(id) ON DELETE CASCADE,
    name TEXT,
    -- nama user (simple dulu, belum auth)
    rating INT CHECK (
        rating >= 1
        AND rating <= 5
    ),
    comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_cafes_city ON cafes(city);
CREATE INDEX idx_menus_cafe_id ON menus(cafe_id);
CREATE INDEX idx_photos_cafe_id ON photos(cafe_id);
CREATE INDEX idx_reviews_cafe_id ON reviews(cafe_id);