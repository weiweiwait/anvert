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
CREATE TABLE gallery (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    username VARCHAR(255),
    title VARCHAR(255),          -- 画作标题
    description TEXT,            -- 画作介绍/描述
    image_url VARCHAR(500),      -- 七牛云图片URL
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES userss(id)
);
CREATE TABLE anthology (
    id          INT PRIMARY KEY AUTO_INCREMENT,
    user_id     INT NOT NULL,                          -- 关联 userss 表
    username    VARCHAR(255) NOT NULL,                  -- 冗余存储用户名（方便诗友圈展示）
    title       VARCHAR(255) DEFAULT '',                -- 文集标题/出处（如"《使至塞上》"）
    content     TEXT NOT NULL,                          -- 诗句正文
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    -- 创建时间
    updated_at  DATETIME DEFAULT NULL,                  -- 更新时间
    FOREIGN KEY (user_id) REFERENCES userss(id)
);
ALTER TABLE gallery ADD COLUMN creator VARCHAR(255) DEFAULT '' COMMENT '画作作者/创作者';
ALTER TABLE gallery ADD COLUMN year VARCHAR(50) DEFAULT '' COMMENT '创作年份';
ALTER TABLE gallery ADD COLUMN material VARCHAR(255) DEFAULT '' COMMENT '材质/媒介';
ALTER TABLE gallery ADD COLUMN size VARCHAR(100) DEFAULT '' COMMENT '尺寸';
