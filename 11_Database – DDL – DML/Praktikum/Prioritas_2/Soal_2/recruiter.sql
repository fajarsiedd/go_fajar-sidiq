-- table users
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255),
    description TEXT,
    address TEXT
);

-- table recruiters
CREATE TABLE recruiters (
    id INT PRIMARY KEY AUTO_INCREMENT,
    company_name VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255)
);

-- table jobs
CREATE TABLE jobs (
    id INT PRIMARY KEY AUTO_INCREMENT,
    recruiter_id INT,
    title VARCHAR(255),
    description TEXT,
    status ENUM('open', 'closed'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (recruiter_id) REFERENCES recruiters(id)
);

-- table job_categories
CREATE TABLE job_categories (
    id INT PRIMARY KEY AUTO_INCREMENT,
    job_id INT,
    category_id INT,
    FOREIGN KEY (job_id) REFERENCES jobs(id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- table categories
CREATE TABLE categories (
    id INT PRIMARY KEY AUTO_INCREMENT,
    category_name VARCHAR(255)
);

-- table experiences
CREATE TABLE experiences (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    company_name VARCHAR(255),
    job_title VARCHAR(255),
    start_date DATE,
    end_date DATE,
    description TEXT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- table job_applications
CREATE TABLE job_applications (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    job_id INT,
    cv VARCHAR(255),
    status ENUM('applied', 'interviewed', 'accepted', 'declined'),
    cover_letter TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (job_id) REFERENCES jobs(id)
);