CREATE TABLE signature (
                       ID INT PRIMARY KEY AUTO_INCREMENT,
                       nickname VARCHAR(255),
                       poem TEXT
);
CREATE TABLE comment(
                        ID INT PRIMARY KEY AUTO_INCREMENT,
                        nickname VARCHAR(255),
                        poem TEXT
);
CREATE TABLE poetical(
                        ID INT PRIMARY KEY AUTO_INCREMENT,
                        nickname VARCHAR(255),
                        poem TEXT
);
CREATE TABLE poetry(
                         ID INT PRIMARY KEY AUTO_INCREMENT,
                         nickname VARCHAR(255),
                         poem TEXT
);
CREATE TABLE userss(
                       id INT PRIMARY KEY AUTO_INCREMENT,
                       username VARCHAR(255),
                       password VARCHAR(255),
                       email VARCHAR(255),
                       avatar_url VARCHAR(500) DEFAULT NULL
);
INSERT INTO userss (username, password, email) VALUES ('李宇翔', '12345678', '12345678@qq.com');
INSERT INTO userss (username, password, email) VALUES ('use2', 'password2', 'user2@example.com');

CREATE TABLE note(
                     id INT PRIMARY KEY AUTO_INCREMENT,
                     user_id INT NOT NULL,
                     user_email VARCHAR(255) NOT NULL,
                     title VARCHAR(255),
                     content TEXT,
                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                     updated_at DATETIME DEFAULT NULL,
                     FOREIGN KEY (user_id) REFERENCES userss(id)
);

-- 如果 userss 表已存在，执行以下语句添加头像字段
ALTER TABLE userss ADD COLUMN avatar_url VARCHAR(500) DEFAULT NULL;