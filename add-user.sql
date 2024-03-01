DROP USER if EXISTS 'admin'@'localhost';
CREATE USER 'admin'@'localhost' IDENTIFIED BY 'morphemetime';
GRANT ALL on recordings.* TO 'admin'@'localhost' WITH GRANT OPTION;
show create user 'admin'@'localhost';