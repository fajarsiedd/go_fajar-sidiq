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
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- table categories
CREATE TABLE categories (
    id INT PRIMARY KEY AUTO_INCREMENT,
    category_name VARCHAR(255)
);

-- table note_categories
CREATE TABLE note_categories (
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
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- table media
CREATE TABLE media (
    id INT PRIMARY KEY AUTO_INCREMENT,
    note_id INT,
    page_id INT,
    media_type ENUM('image', 'video', 'audio', 'gif'),
    media_url VARCHAR(255),
    FOREIGN KEY (note_id) REFERENCES notes(id),
    FOREIGN KEY (page_id) REFERENCES pages(id)
);

-- table tasks
CREATE TABLE tasks (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    title VARCHAR(255),
    description TEXT,
    status ENUM('draft', 'ongoing', 'done', 'canceled', 'pending'),
    due_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);