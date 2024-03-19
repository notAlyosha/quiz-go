CREATE DATABASE IF NOT EXISTS quiz_app_data DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;
CREATE USER IF NOT EXISTS 'web_quiz'@'%' IDENTIFIED BY 'password_DEV_testing_0';
GRANT SELECT, INSERT, UPDATE ON quiz_app_data.* TO 'web_quiz'@'%';
REVOKE ALL ON quiz_app_data.history FROM 'web_quiz'@'%';
GRANT SELECT ON quiz_app_data.history TO 'web_quiz'@'%';

CREATE DATABASE IF NOT EXISTS quiz_app_data_sequrity_data DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;
CREATE USER IF NOT EXISTS 'web_quiz_sequrity_640'@'127.0.0.1' IDENTIFIED BY 'enought_strong_password_here';
GRANT SELECT, INSERT, UPDATE ON web_quiz_sequrity_personal_data.* TO 'web_quiz_sequrity_640'@'127.0.0.1';

FLUSH PRIVILEGES;

DROP DATABASE IF EXISTS quiz_app_data;
DROP DATABASE IF EXISTS quiz_app_data_sequrity_data;

CREATE TABLE IF NOT EXISTS quiz_app_data.application_settings
(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR NOT NULL UNIQUE,
    value VARCHAR NOT NULL,
    is_deleted BOOLEAN DEFAULT FALSE,
    
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS quiz_app_data.users
(
    id INT NOT NULL AUTO_INCREMENT,
    front_id VARCHAR NOT NULL UNIQUE,
    role VARCHAR NOT NULL,
    logo_url VARCHAR NULL,
    is_deleted BOOLEAN DEFAULT FALSE,
    
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS quiz_app_data.subjects
(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR NOT NULL UNIQUE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS quiz_app_data.groups
(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR,
    is_deleted BOOLEAN DEFAULT FALSE,

    CHECK(name IS NOT NULL),
    CHECK(is_deleted IS NOT NULL),
    
    UNIQUE(name),
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS quiz_app_data.sessions_statuses
(
    id INT NOT NULL AUTO_INCREMENT,
    status VARCHAR NOT NULL UNIQUE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS quiz_app_data.questions_types
( 
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR,
    durability INT DEFAULT 0,
    points INT,
    is_deleted BOOLEAN DEFAULT FALSE,

    CHECK (name IS NOT NULL),
    CHECK (durability IS NOT NULL),
    CHECK (points IS NOT NULL),
    CHECK (is_deleted IS NOT NULL),

    UNIQUE(name),

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS quiz_app_data.groups_subjects
(
    id INT NOT NULL AUTO_INCREMENT,
    subject_id INT NOT NULL,
    group_id INT NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY(id),
    FOREIGN KEY(subject_id) REFERENCES quiz_app_data.subjects(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY(group_id) REFERENCES quiz_app_data.groups(id)
        ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS quiz_app_data.groups_students
(
    id INT NOT NULL AUTO_INCREMENT,
    student_id INT NOT NULL,
    group_id INT NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY(id),
    FOREIGN KEY(student_id) REFERENCES quiz_app_data.users(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY(group_id) REFERENCES quiz_app_data.groups(id)
        ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS quiz_app_data.groups_teachers
(
    id INT NOT NULL AUTO_INCREMENT,
    teacher_id INT NOT NULL,
    group_id INT NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY(id),
    FOREIGN KEY(teacher_id) REFERENCES quiz_app_data.users(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY(group_id) REFERENCES quiz_app_data.groups(id)
        ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS quiz_app_data.quizes
(
    id INT NOT NULL AUTO_INCREMENT,
    subject_id INT,
    name VARCHAR,
    description VARCHAR,
    summary_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    max_points VARCHAR,
    max_add_points VARCHAR,
    is_deleted BOOLEAN DEFAULT FALSE,


    CHECK (subject_id IS NOT NULL),
    CHECK (name IS NOT NULL),
    CHECK (summary_time IS NOT NULL),
    CHECK (max_points IS NOT NULL),
    CHECK (max_points > 0),
    CHECK (max_add_points IS NOT NULL),
    CHECK (max_add_points > 0),
    CHECK (is_deleted IS NOT NULL),
    
    PRIMARY KEY (id),

    FOREIGN KEY (subject_id)
    REFERENCES quiz_app_data.subjects(id)
);

CREATE TABLE IF NOT EXISTS quiz_app_data.quizes_teachers
( 
    id INT NOT NULL AUTO_INCREMENT,
    teacher_id INT,
    quiz_id INT,
    is_deleted BOOLEAN DEFAULT FALSE,

    CHECK (teacher_id IS NOT NULL),
    CHECK (quiz_id IS NOT NULL),
    CHECK (is_deleted IS NOT NULL),

    PRIMARY KEY (id),
    FOREIGN KEY (teacher_id)
    REFERENCES quiz_app_data.users(id),
    FOREIGN KEY (quiz_id)
    REFERENCES quiz_app_data.quizes(id)
);

CREATE TABLE IF NOT EXISTS quiz_app_data.quizes_questions_containers
(
    id INT NOT NULL AUTO_INCREMENT,
    quiz_id INT,
    question_type_id INT,
    question_text VARCHAR,
    picture_URI VARCHAR,
    is_additional BOOLEAN DEFAULT FALSE,
    is_deleted BOOLEAN DEFAULT FALSE,


    CHECK (quiz_id IS NOT NULL),
    CHECK (question_type_id IS NOT NULL),
    CHECK (question_text IS NOT NULL),
    CHECK (is_additional IS NOT NULL),
    CHECK (is_deleted IS NOT NULL),
    
    PRIMARY KEY (id),

    FOREIGN KEY (quiz_id)
    REFERENCES quiz_app_data.quizes(id),

    FOREIGN KEY (question_type_id)
    REFERENCES quiz_app_data.questions_types(id)

);

CREATE TABLE IF NOT EXISTS quiz_app_data.paired_answers_questions_options
(
    id INT NOT NULL AUTO_INCREMENT,
    question_id INT,
    option_text VARCHAR,
    is_deleted BOOLEAN DEFAULT FALSE,

    CHECK (question_id IS NOT NULL),
    CHECK (option_text IS NOT NULL),
    CHECK (is_deleted IS NOT NULL),

    PRIMARY KEY (id),

    FOREIGN KEY(question_id)
    REFERENCES quiz_app_data.quizes_questions_containers(id)

);

CREATE TABLE IF NOT EXISTS quiz_app_data.paired_answers_questions_options_links
(
    id INT NOT NULL AUTO_INCREMENT,
    left_side_id INT,
    right_side_id INT,
    is_deleted BOOLEAN DEFAULT FALSE,

    CHECK (left_side_id IS NOT NULL),
    CHECK (right_side_id IS NOT NULL),
    CHECK (is_deleted IS NOT NULL),

    PRIMARY KEY (id),

    FOREIGN KEY (left_side_id)
    REFERENCES quiz_app_data.paired_answers_questions_options(id),

    FOREIGN KEY (right_side_id)
    REFERENCES quiz_app_data.paired_answers_questions_options(id)
);

CREATE TABLE IF NOT EXISTS quiz_app_data.closed_questions_options
(
    id INT NOT NULL AUTO_INCREMENT,
    question_id INT,
    option_text VARCHAR,
    is_correct_option  BOOLEAN DEFAULT FALSE, 
    is_deleted BOOLEAN DEFAULT FALSE,

    CHECK (question_id IS NOT NULL),
    CHECK (option_text IS NOT NULL),
    CHECK (is_correct_option IS NOT NULL),
    CHECK (is_deleted IS NOT NULL),

    PRIMARY KEY (id),

    FOREIGN KEY(question_id) 
    REFERENCES quiz_app_data.quizes_questions_containers(id)
);

CREATE TABLE IF NOT EXISTS quiz_app_data.sequenced_question_options
(
    id INT NOT NULL AUTO_INCREMENT,
    question_id INT,
    option_text VARCHAR,
    is_deleted BOOLEAN DEFAULT FALSE,

    CHECK (question_id IS NOT NULL),
    CHECK (option_text IS NOT NULL),
    CHECK (is_deleted IS NOT NULL),

    PRIMARY KEY (id),

    FOREIGN KEY(question_id)
    REFERENCES quiz_app_data.quizes_questions_containers(id)
);

CREATE TABLE IF NOT EXISTS quiz_app_data.parts_insert_questions_options
(
    id INT NOT NULL AUTO_INCREMENT,
    question_id INT,
    part_text VARCHAR,
    is_hidden BOOLEAN DEFAULT FALSE,
    is_deleted BOOLEAN DEFAULT FALSE,
    
    CHECK (question_id IS NOT NULL),
    CHECK (part_text IS NOT NULL),
    CHECK (is_deleted IS NOT NULL),

    PRIMARY KEY (id),

    FOREIGN KEY(question_id)
    REFERENCES quiz_app_data.quizes_questions_containers(id)

);

CREATE TABLE IF NOT EXISTS quiz_app_data.geo_with_checking_questions
(
    id INT NOT NULL AUTO_INCREMENT,
    question_id INT,
    longitude INT,
    lantitude INT,
    is_place_visited BOOLEAN DEFAULT FALSE,
    check_q_text VARCHAR,
    check_q_answer VARCHAR,
    is_deleted BOOLEAN DEFAULT FALSE,
    
    CHECK (question_id IS NOT NULL),
    CHECK (longitude IS NOT NULL),
    CHECK (lantitude IS NOT NULL),
    CHECK (is_place_visited IS NOT NULL),
    CHECK (check_q_text IS NOT NULL),
    CHECK (check_q_answer IS NOT NULL),
    CHECK (is_deleted IS NOT NULL),

    PRIMARY KEY (id),

    FOREIGN KEY(question_id)
    REFERENCES quiz_app_data.quizes_questions_containers(id)

);


CREATE TABLE IF NOT EXISTS quiz_app_data.opened_question
(
    id INT NOT NULL AUTO_INCREMENT,
    question_id INT,
    answer_text VARCHAR,
    is_deleted BOOLEAN DEFAULT FALSE,

    CHECK (question_id IS NOT NULL),
    CHECK (answer_text IS NOT NULL),
    CHECK (is_deleted IS NOT NULL),

    PRIMARY KEY (id),

    FOREIGN KEY (question_id)
    REFERENCES quiz_app_data.quizes_questions_containers(id)

);

CREATE TABLE IF NOT EXISTS quiz_app_data.sessions
(
    id INT NOT NULL AUTO_INCREMENT,
    front_id INT NOT NULL UNIQUE,
    teacher_id INT NOT NULL,
    quiz_id INT NOT NULL,
    status_id INT NOT NULL,
    date_time_start TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    date_time_end TIMESTAMP NULL,
    is_summary_exam BOOLEAN NOT NULL DEFAULT FALSE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY(id),
    FOREIGN KEY(teacher_id) REFERENCES quiz_app_data.users(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY(quiz_id) REFERENCES quiz_app_data.quizes(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY(status_id) REFERENCES quiz_app_data.sessions_statuses(id)
        ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS quiz_app_data.sessions_reexam_orders
(
    id INT NOT NULL AUTO_INCREMENT,
    teacher_id INT NOT NULL,
    student_id INT NOT NULL,
    main_session_id INT NOT NULL,
    order_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY(id),
    FOREIGN KEY(teacher_id) REFERENCES quiz_app_data.users(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY(student_id) REFERENCES quiz_app_data.users(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY(main_session_id) REFERENCES quiz_app_data.sessions(id)
        ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS quiz_app_data.students_sessions_results
(
    id INT NOT NULL AUTO_INCREMENT,
    session_id INT NOT NULL,
    student_id INT NOT NULL,
    summaty_mark INT NULL,
    additional_points INT NOT NULL,
    date_time_end TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_bought BOOLEAN NOT NULL DEFAULT FALSE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY(id),
    FOREIGN KEY(session_id) REFERENCES quiz_app_data.sessions(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY(student_id) REFERENCES quiz_app_data.users(id)
        ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS quiz_app_data.sessions_orders
(
    id INT NOT NULL AUTO_INCREMENT,
    reexam_session_id INT NOT NULL,
    session_id INT NOT NULL,
    order_id INT NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY(id),
    FOREIGN KEY(reexam_session_id) REFERENCES quiz_app_data.sessions(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY(session_id) REFERENCES quiz_app_data.sessions(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY(order_id) REFERENCES quiz_app_data.sessions_reexam_orders(id)
        ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS quiz_app_data.chats
(
    id INT NOT NULL AUTO_INCREMENT,
    chat_administrator_id INT NULL,
    description VARCHAR NULL,
    invite_link VARCHAR NULL,
    theme_name VARCHAR NULL,
    is_group_chat BOOLEAN NOT NULL DEFAULT FALSE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY(id),
    FOREIGN KEY(chat_administrator_id)
    REFERENCES quiz_app_data.users(id)
    ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS quiz_app_data.chats_users
(
    id INT NOT NULL AUTO_INCREMENT,
    chat_id INT NOT NULL,
    user_id INT NOT NULL,
    is_deleted_by_chat_administrator BOOLEAN NOT NULL DEFAULT FALSE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY(id),
    FOREIGN KEY(chat_id)
    REFERENCES quiz_app_data.chats(id)
    ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY(user_id)
    REFERENCES quiz_app_data.users(id)
    ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS quiz_app_data.chats_messages
(
    id INT NOT NULL AUTO_INCREMENT,
    chat_id INT NOT NULL,
    user_id INT NOT NULL,
    message_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    message_text VARCHAR NULL,
    file_url VARCHAR NULL,
    is_deleted_by_chat_user BOOLEAN NOT NULL DEFAULT FALSE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY(id),
    FOREIGN KEY(chat_id)
    REFERENCES quiz_app_data.chats(id)
    ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY(user_id)
    REFERENCES quiz_app_data.users(id)
    ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS quiz_app_data.messages_replies
(
    id INT NOT NULL AUTO_INCREMENT,
    message_main_id INT NOT NULL,
    message_reply_id INT NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY(id),
    FOREIGN KEY(message_main_id)
    REFERENCES quiz_app_data.chats_messages(id)
    ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY(message_reply_id)
    REFERENCES quiz_app_data.chats_messages(id)
    ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS quiz_app_data.history
(
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    table_name VARCHAR NOT NULL,
    column_name VARCHAR NOT NULL,
    row_id INT NOT NULL,
    old_value VARCHAR NOT NULL,
    new_value VARCHAR NOT NULL,
    edit_timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN DEFAULT FALSE,
    
    PRIMARY KEY(id),
    FOREIGN KEY(user_id)
    REFERENCES quiz_app_data.users(id)
    ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS quiz_app_data_sequrity_data.entering_attemts
(
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    entering_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_correct BOOLEAN NOT NULL DEFAULT FALSE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY(id),
    FOREIGN KEY(user_id)
    REFERENCES quiz_app_data.users(id)
    ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS quiz_app_data_sequrity_data.users_passwords_history
(
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    value BINARY(255) NOT NULL,
    entering_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY(id),
    FOREIGN KEY(user_id)
    REFERENCES quiz_app_data.users(id)
    ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS quiz_app_data_sequrity_data.profiles
(
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL UNIQUE,
    login VARCHAR NOT NULL UNIQUE,
    salt BINARY(255) NOT NULL,
    name VARCHAR NULL,
    email VARCHAR NULL UNIQUE,
    phone VARCHAR NULL UNIQUE,
    adding_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY(id),
    FOREIGN KEY profiles(user_id)
    REFERENCES quiz_app_data.users(id)
    ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS quiz_app_data_sequrity_data.students_wallets
(
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    points INT NOT NULL DEFAULT 0,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY(id),
    FOREIGN KEY(user_id)
    REFERENCES quiz_app_data.users(id)
    ON UPDATE CASCADE ON DELETE CASCADE
);