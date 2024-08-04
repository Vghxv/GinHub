CREATE DATABASE music;

\c music

CREATE TABLE albums (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    artist VARCHAR(255) NOT NULL,
    price NUMERIC(10, 2) NOT NULL
);

INSERT INTO albums (title, artist, price) VALUES 
('Blue Train', 'John Coltrane', 56.99),
('Jeru', 'Gerry Mulligan', 17.99),
('Sarah Vaughan and Clifford Brown', 'Sarah Vaughan', 39.99);
