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
    agoda_room_name VARCHAR(25),
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
DROP TABLE IF EXISTS user;

-- Create the user table
CREATE TABLE IF NOT EXISTS user (
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

-- Drop the user table if it exists
DROP TABLE IF EXISTS user;

-- Create the user table
CREATE TABLE IF NOT EXISTS user (
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
INSERT INTO user (first_name, second_name, username, password, active_flg)
VALUES 
('Alex', 'Smith', 'alex', 'useralex', 1),
('Gopal', 'Patel', 'gopal', 'usergopal', 1),
('Engadev', 'Jones', 'engadev', 'userengadev', 1);

