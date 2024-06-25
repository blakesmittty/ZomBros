-- Reset database by dropping, recreating tables, and truncating data
DROP TABLE IF EXISTS players, game_progress, items, player_items;

CREATE TABLE players (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE game_progress (
    id SERIAL PRIMARY KEY,
    player_id INTEGER REFERENCES players(id),
    player_level INTEGER NOT NULL,
    score INTEGER NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    item_name VARCHAR(100) NOT NULL
);

CREATE TABLE player_items (
    id SERIAL PRIMARY KEY,
    player_id INTEGER REFERENCES players(id),
    item_id INTEGER REFERENCES items(id),
    unlocked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
