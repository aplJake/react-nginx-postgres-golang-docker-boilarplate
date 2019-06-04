CREATE TABLE messages
(
    id         SERIAL PRIMARY KEY,
    message    VARCHAR(32)              NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);

INSERT INTO messages(id, message, created_at) VALUES(DEFAULT, 'Hello $1 000 000', now());