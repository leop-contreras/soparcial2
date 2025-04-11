CREATE TABLE users (
    id INT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL
);

INSERT INTO users (id, name, email) VALUES
(1234, 'Leo', 'cont@example.com'),
(5678, 'Raul', 'bec@example.com');