CREATE DATABASE IF NOT EXISTS db_routes;
USE db_routes;

CREATE TABLE IF NOT EXISTS routes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    source JSON NOT NULL,
    destination JSON NOT NULL
    );