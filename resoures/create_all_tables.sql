-- =============================================
-- ArtVerse 项目完整建表 SQL
-- 生成时间: 2026-03-31
-- 数据库: MySQL
-- =============================================

-- 1. 用户表
CREATE TABLE userss (
    id         INT PRIMARY KEY AUTO_INCREMENT,
    username   VARCHAR(255),
    password   VARCHAR(255),
    email      VARCHAR(255),
    avatar_url VARCHAR(500) DEFAULT NULL
);

-- 2. 签名表
CREATE TABLE signature (
    ID       INT PRIMARY KEY AUTO_INCREMENT,
    nickname VARCHAR(255),
    poem     TEXT
);

-- 3. 评论表
CREATE TABLE comment (
    ID       INT PRIMARY KEY AUTO_INCREMENT,
    nickname VARCHAR(255),
    poem     TEXT
);

-- 4. 诗作表
CREATE TABLE poetical (
    ID       INT PRIMARY KEY AUTO_INCREMENT,
    nickname VARCHAR(255),
    poem     TEXT
);

-- 5. 诗句表
CREATE TABLE poetry (
    ID       INT PRIMARY KEY AUTO_INCREMENT,
    nickname VARCHAR(255),
    poem     TEXT
);

-- 6. 笔记表
CREATE TABLE note (
    id         INT PRIMARY KEY AUTO_INCREMENT,
    user_id    INT NOT NULL,
    user_email VARCHAR(255) NOT NULL,
    title      VARCHAR(255),
    content    TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES userss(id)
);

-- 7. 画廊表
CREATE TABLE gallery (
    id          INT PRIMARY KEY AUTO_INCREMENT,
    user_id     INT NOT NULL,
    username    VARCHAR(255),
    title       VARCHAR(255)            COMMENT '画作标题',
    description TEXT                    COMMENT '画作介绍/描述',
    image_url   VARCHAR(500)            COMMENT '七牛云图片URL',
    creator     VARCHAR(255) DEFAULT '' COMMENT '画作作者/创作者',
    year        VARCHAR(50)  DEFAULT '' COMMENT '创作年份',
    material    VARCHAR(255) DEFAULT '' COMMENT '材质/媒介',
    size        VARCHAR(100) DEFAULT '' COMMENT '尺寸',
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES userss(id)
);

-- 8. 文集表（诗友圈）
CREATE TABLE anthology (
    id         INT PRIMARY KEY AUTO_INCREMENT,
    user_id    INT NOT NULL                       COMMENT '关联 userss 表',
    username   VARCHAR(255) NOT NULL              COMMENT '冗余存储用户名（方便诗友圈展示）',
    title      VARCHAR(255) DEFAULT ''            COMMENT '文集标题/出处',
    content    TEXT NOT NULL                       COMMENT '诗句正文',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES userss(id)
);

-- 9. 关注表（粉丝关系）
CREATE TABLE user_follow (
    id          INT PRIMARY KEY AUTO_INCREMENT,
    follower_id INT NOT NULL                       COMMENT '关注者（粉丝）的用户ID',
    followed_id INT NOT NULL                       COMMENT '被关注者的用户ID',
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_follow (follower_id, followed_id),
    FOREIGN KEY (follower_id) REFERENCES userss(id),
    FOREIGN KEY (followed_id) REFERENCES userss(id)
);
