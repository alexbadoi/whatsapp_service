-- Use the newly created database
create database if not exists data_feeds;

-- Use the newly created database
USE data_feeds;

-- Drop the agoda_data table if it exists
DROP TABLE IF EXISTS agoda_data;

-- Create the agoda_data table
CREATE TABLE IF NOT EXISTS agoda_data (
    proposal_detail_id INT PRIMARY KEY,
    proposal_id INT NOT NULL,
    day_count INT NOT NULL,
    agoda_hotel_location_id varchar(50),
    agoda_location_name VARCHAR(255),
    start_dt DATE NULL,
    end_dt DATE NULL,
    number_of_rooms INT,
    number_of_adults INT,
    number_of_children INT,
    children_ages_pipe_delimited VARCHAR(50),
    agoda_hotel_name VARCHAR(255),
    agoda_hodel_id INT,
    agoda_review_score FLOAT,
    agoda_review_count INT,
    best_price_per_night DECIMAL(10, 2),
    agoda_room_pic_pipe_delimited TEXT,
    agoda_hotel_pic_pipe_delimited TEXT,
    sleeps_count INT,
    room_price_per_night DECIMAL(10, 2),
    risk_free_booking_binary TINYINT,
    cancelation_deadline_dt DATE,
    includes_breakfast_bainary TINYINT,
    includeas_lunch_binary TINYINT,
    includes_dinner_binary TINYINT,
    bed_type VARCHAR(255),
    size_sqm FLOAT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    active_flg INT DEFAULT 1
);

-- Drop the user table if it exists
DROP TABLE IF EXISTS user_ref;

-- Create the user table
CREATE TABLE IF NOT EXISTS user_ref (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    second_name VARCHAR(100) NOT NULL,
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    active_flg INT DEFAULT 1
);


-- Insert users Alex, Gopal, and Engadev
INSERT INTO user_ref (first_name, second_name, username, password, active_flg)
VALUES 
('Alex', 'Smith', 'alex', 'useralex', 1),
('Gopal', 'Patel', 'gopal', 'usergopal', 1),
('Engadev', 'Jones', 'engadev', 'userengadev', 1);


-- Drop table if it exists
DROP TABLE IF EXISTS contact_whatsapp;

-- Create table if not exists
CREATE TABLE IF NOT EXISTS contact_whatsapp (
    contact_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NULL,
    origin_country_id INT DEFAULT NULL,
    address_1 VARCHAR(255) DEFAULT NULL,
    address_2 VARCHAR(255) DEFAULT NULL,
    co_phone_cd VARCHAR(10),
    mobile VARCHAR(20),
    email VARCHAR(255) NULL,
    tag_id INT NULL,
    agree_contact INT DEFAULT NULL,
    agree_promo INT DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,
    active_flg INT DEFAULT 1
);

-- Load data from CSV file
LOAD DATA INFILE 'C:\\ProgramData\\MySQL\\MySQL Server 9.0\\Uploads\\contact_whatsapp.csv'
INTO TABLE contact_whatsapp
CHARACTER SET utf8mb4
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 1 ROWS
(contact_id, name, surname, origin_country_id, address_1, address_2, co_phone_cd, mobile, email, tag_id, agree_contact, agree_promo, @created_at, @updated_at, @deleted_at, active_flg)
SET
    created_at = IFNULL(NULLIF(@created_at, ''), CURRENT_TIMESTAMP),
    updated_at = IFNULL(NULLIF(@updated_at, ''), CURRENT_TIMESTAMP),
    deleted_at = NULLIF(@deleted_at, '');



-- Drop table if it exists
DROP TABLE IF EXISTS tag_ref;

-- Create table if not exists
CREATE TABLE IF NOT EXISTS tag_ref (
    tag_id INT AUTO_INCREMENT PRIMARY KEY,
    tag_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,
    active_flg INT DEFAULT 1
);

-- Load data from CSV file
LOAD DATA INFILE 'C:\\ProgramData\\MySQL\\MySQL Server 9.0\\Uploads\\tag_ref.csv'
INTO TABLE tag_ref
CHARACTER SET utf8mb4
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 1 ROWS
(tag_id, tag_name, @created_at, @updated_at, @deleted_at, active_flg)
SET
    created_at = IFNULL(NULLIF(@created_at, ''), CURRENT_TIMESTAMP),
    updated_at = IFNULL(NULLIF(@updated_at, ''), CURRENT_TIMESTAMP),
    deleted_at = NULLIF(@deleted_at, '');
    
    
-- Grant all privileges on the data_feeds database to the users

GRANT ALL PRIVILEGES ON data_feeds.* TO 'gopal'@'%';
GRANT ALL PRIVILEGES ON data_feeds.* TO 'engadev'@'%';

-- Flush privileges to ensure that the changes take effect
FLUSH PRIVILEGES;

select * from user_ref;


