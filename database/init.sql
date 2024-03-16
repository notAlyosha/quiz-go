CREATE DATABASE IF NOT EXISTS 'quiz_app_data' DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;
CREATE USER IF NOT EXISTS 'web_quiz'@'%' IDENTIFIED BY 'password_DEV_testing_0';
GRANT SELECT, INSERT, UPDATE ON 'quiz_app_data'.* TO 'web_quiz'@'%';

-- TODO: understatd, how to split users and their privileges between different data bases
-- CREATE DATABASE IF NOT EXISTS 'quiz_app_data_sequrity_data' DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;
-- CREATE USER IF NOT EXISTS 'web_quiz_sequrity_600' IDENTIFIED BY 'enought_strong_password_here';
-- GRANT SELECT, INSERT, UPDATE ON 'web_quiz_sequrity_personal_data'.* TO 'web_quiz_sequrity_600'*

FLUSH PRIVILEGES;