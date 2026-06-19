CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE destinations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    country VARCHAR(100) NOT NULL,
    description TEXT,
    estimated_daily_cost INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE trips (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    destination_id INT NOT NULL,
    title VARCHAR(150) NOT NULL,
    duration_days INT NOT NULL,
    budget INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_destination
        FOREIGN KEY(destination_id)
        REFERENCES destinations(id)
);

CREATE TABLE itineraries (
    id SERIAL PRIMARY KEY,
    trip_id INT NOT NULL,
    day_number INT NOT NULL,
    activity VARCHAR(255) NOT NULL,
    location VARCHAR(150),
    estimated_cost INT DEFAULT 0,

    CONSTRAINT fk_trip
        FOREIGN KEY(trip_id)
        REFERENCES trips(id)
        ON DELETE CASCADE
);

-- table ini guna nya liat history promt sama debug response dair ai
CREATE TABLE ai_generations (
    id SERIAL PRIMARY KEY,
    trip_id INT NOT NULL,
    prompt TEXT NOT NULL,
    generated_result TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (trip_id) REFERENCES trips(id)
    ON DELETE CASCADE
);