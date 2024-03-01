DROP USER if EXISTS 'apiservice'@'localhost';
CREATE USER 'apiservice'@'localhost' IDENTIFIED BY 'badpassword';
GRANT ALL on recordings.* TO 'apiservice'@'localhost' WITH GRANT OPTION;
show create user 'apiservice'@'localhost';