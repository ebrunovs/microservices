CREATE DATABASE IF NOT EXISTS `order`;
USE `order`;

INSERT INTO stock (product_code, quantity, created_at, updated_at) VALUES
('PROD001', 100, NOW(), NOW()),
('PROD002', 50, NOW(), NOW()),
('PROD003', 200, NOW(), NOW()),
('PROD004', 75, NOW(), NOW()),
('PROD005', 30, NOW(), NOW()),
('LAPTOP01', 25, NOW(), NOW()),
('MOUSE01', 150, NOW(), NOW()),
('KEYBOARD01', 80, NOW(), NOW()),
('MONITOR01', 40, NOW(), NOW()),
('HEADSET01', 60, NOW(), NOW())
ON DUPLICATE KEY UPDATE quantity = VALUES(quantity);