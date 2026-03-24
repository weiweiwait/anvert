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
                       email VARCHAR(255)
);
INSERT INTO userss (username, password, email) VALUES ('李宇翔', '12345678', '12345678@qq.com');
INSERT INTO userss (username, password, email) VALUES ('use2', 'password2', 'user2@example.com');