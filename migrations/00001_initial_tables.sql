-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    nickname VARCHAR(30)
);

CREATE TABLE tournaments
(
    id SERIAL PRIMARY KEY, 
    name TEXT UNIQUE
);

CREATE TABLE events
(
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE,
    tournament_id INTEGER REFERENCES tournaments (id)
);

CREATE TABLE brackets
(
    id SERIAL PRIMARY KEY,
    event_id INTEGER REFERENCES events (id)
);

CREATE TABLE pools 
(
    id SERIAL PRIMARY KEY,
    bracket_id INTEGER REFERENCES brackets (id)
);

CREATE TABLE sets
(
    id SERIAL PRIMARY KEY,
    player1_id INTEGER REFERENCES users (id),
    player2_id INTEGER REFERENCES users (id),
    pool_id INTEGER
    -- pool_id INTEGER REFERENCES pools (id)
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE sets;
DROP TABLE pools;
DROP TABLE brackets;
DROP TABLE events;
DROP TABLE tournaments;
DROP TABLE users;
-- +goose StatementEnd
