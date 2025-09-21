CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    nickname VARCHAR(30)
);

CREATE TABLE IF NOT EXISTS sets
(
    id SERIAL PRIMARY KEY,
    player1 INTEGER REFERENCES users (id),
    player2 INTEGER REFERENCES users (id),
    pool_id INTEGER REFERENCES pools (id)
);

CREATE TABLE IF NOT EXISTS pools 
(
    id SERIAL PRIMARY KEY,
    bracket_id INTEGER REFERENCES brackets (id)
);

CREATE TABLE IF NOT EXISTS brackets
(
    id SERIAL PRIMARY KEY,
    event_id INTEGER REFERENCES events (id)
);

CREATE TABLE IF NOT EXISTS events
(
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE,
    tournament_id INTEGER REFERENCES tournament (id)
);

CREATE TABLE IF NOT EXISTS tournaments
(
    id SERIAL PRIMARY KEY, 
    name UNIQUE TEXT
);