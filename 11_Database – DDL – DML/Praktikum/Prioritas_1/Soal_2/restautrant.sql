-- Tabel restaurant_types
CREATE TABLE restaurant_types (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255)
);

-- Tabel restaurants
CREATE TABLE restaurants (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    address TEXT,
    restaurant_type_id INT,
    FOREIGN KEY (restaurant_type_id) REFERENCES restaurant_types(id)
);

-- Tabel menus
CREATE TABLE menus (
    id INT PRIMARY KEY AUTO_INCREMENT,
    menu_type_id INT,
    restaurant_id INT,
    name VARCHAR(255),
    description TEXT,
    price INT,
    FOREIGN KEY (menu_type_id) REFERENCES menu_types(id),
    FOREIGN KEY (restaurant_id) REFERENCES restaurants(id)
);

-- Tabel menu_types
CREATE TABLE menu_types (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255)
);

-- Tabel users
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255)
);

-- Tabel user_reviews
CREATE TABLE user_reviews (
    id INT PRIMARY KEY AUTO_INCREMENT,
    restaurant_id INT,
    rating FLOAT,
    description TEXT,
    user_id INT,
    FOREIGN KEY (restaurant_id) REFERENCES restaurants(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);