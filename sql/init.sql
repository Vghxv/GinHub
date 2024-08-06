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

CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);

INSERT INTO customers (first_name, last_name, email) VALUES 
('John', 'Doe', 'john212@gmail.com'),
('Jane', 'taylor', 'f23@gmail.com'),
('Mary', 'Smith', 'msmi44@gmail.com');

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    album_id INT NOT NULL,
    quantity INT NOT NULL,
    order_date DATE NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers(id),
    FOREIGN KEY (album_id) REFERENCES albums(id)
);

INSERT INTO orders (customer_id, album_id, quantity, order_date) VALUES 
(1, 1, 1, '2021-01-01'),
(1, 2, 2, '2021-01-02'),
(2, 3, 1, '2021-01-03'),
(3, 1, 1, '2021-01-04'),
(3, 2, 1, '2021-01-05'),
(3, 3, 2, '2021-01-06');
