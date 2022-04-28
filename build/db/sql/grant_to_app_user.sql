CREATE USER 'winelist-app'@'%' IDENTIFIED BY 'winelist-password';
GRANT SELECT,INSERT,UPDATE,DELETE ON winelist.* TO 'winelist-app'@'%';