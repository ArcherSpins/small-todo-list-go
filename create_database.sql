CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    resolved BOOLEAN NOT NULL DEFAULT FALSE
);