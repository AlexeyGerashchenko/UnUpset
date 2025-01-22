CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       user_name VARCHAR(20),
                       date_of_creation TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE todos (
                       task_id SERIAL PRIMARY KEY,
                       user_id INT NOT NULL,
                       text VARCHAR(255) NOT NULL,
                       status BOOLEAN NOT NULL ,
                       FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE statistics (
                            user_id INT NOT NULL,
                            completed_tasks_count INT DEFAULT 0,
                            timer_start_count INT DEFAULT 0,
                            record_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE timers (
                        user_id INT NOT NULL,
                        timer_length INT NOT NULL,
                        start_timer TIMESTAMP NOT NULL,
                        end_timer TIMESTAMP NOT NULL,
                        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);