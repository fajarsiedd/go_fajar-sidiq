-- table users
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255)
);

-- table notes
CREATE TABLE notes (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    title VARCHAR(255),
    content TEXT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- table categories
CREATE TABLE categories (
    id INT PRIMARY KEY AUTO_INCREMENT,
    category_name VARCHAR(255)
);

-- table category_notes
CREATE TABLE category_notes (
    id INT PRIMARY KEY AUTO_INCREMENT,
    note_id INT,
    category_id INT,
    FOREIGN KEY (note_id) REFERENCES notes(id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- table pages
CREATE TABLE pages (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    title VARCHAR(255),
    content TEXT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- table media
CREATE TABLE media (
    id INT PRIMARY KEY AUTO_INCREMENT,
    note_id INT,
    page_id INT,
    media_type VARCHAR(255),
    media_url TEXT,
    FOREIGN KEY (note_id) REFERENCES notes(id),
    FOREIGN KEY (page_id) REFERENCES pages(id)
);

-- table tasks
CREATE TABLE tasks (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    title VARCHAR(255),
    description TEXT,
    status VARCHAR(255),
    due_date DATE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);