-- 创建数据库
CREATE DATABASE IF NOT EXISTS myvault CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 允许 root 用户从任何地址连接
ALTER USER 'root'@'%' IDENTIFIED BY '111111';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;

USE myvault;

-- 用户表
CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    avatar VARCHAR(255),
    github_username VARCHAR(100),
    github_id VARCHAR(50),
    access_token TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- 活动表
CREATE TABLE IF NOT EXISTS activities (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    date DATE NOT NULL,
    summary TEXT,
    ai_generated BOOLEAN DEFAULT FALSE,
    has_activity BOOLEAN DEFAULT FALSE,
    commit_count INT DEFAULT 0,
    total_time INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    INDEX idx_user_date (user_id, date)
);

-- 数据源表
CREATE TABLE IF NOT EXISTS data_sources (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    activity_id BIGINT UNSIGNED NOT NULL,
    type VARCHAR(50) NOT NULL,
    data TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (activity_id) REFERENCES activities(id) ON DELETE CASCADE
);

-- 提交记录表
CREATE TABLE IF NOT EXISTS commits (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    activity_id BIGINT UNSIGNED NOT NULL,
    hash VARCHAR(40) NOT NULL,
    message TEXT NOT NULL,
    repository VARCHAR(255),
    author VARCHAR(100),
    time TIMESTAMP NOT NULL,
    files INT DEFAULT 0,
    additions INT DEFAULT 0,
    deletions INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (activity_id) REFERENCES activities(id) ON DELETE CASCADE,
    INDEX idx_activity_time (activity_id, time)
);

-- 仓库表
CREATE TABLE IF NOT EXISTS repositories (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    description TEXT,
    language VARCHAR(50),
    private BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 插入示例数据
INSERT INTO users (username, email, password, avatar) VALUES 
('demo-user', 'demo@example.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'https://via.placeholder.com/150');

-- 获取刚插入的用户ID
SET @user_id = LAST_INSERT_ID();

-- 插入示例活动
INSERT INTO activities (user_id, date, summary, ai_generated, has_activity, commit_count) VALUES
(@user_id, CURDATE(), '今天完成了React组件优化和API接口开发', TRUE, TRUE, 5),
(@user_id, DATE_SUB(CURDATE(), INTERVAL 1 DAY), '实现了用户认证系统和GitHub集成', TRUE, TRUE, 8),
(@user_id, DATE_SUB(CURDATE(), INTERVAL 2 DAY), '今日无编程活动', FALSE, FALSE, 0);

-- 获取活动ID
SET @activity_id = (SELECT id FROM activities WHERE user_id = @user_id AND date = CURDATE());

-- 插入示例提交记录
INSERT INTO commits (activity_id, hash, message, repository, author, time, files, additions, deletions) VALUES
(@activity_id, 'abc123def456', 'feat: 优化Timeline组件性能', 'MyVault', 'demo-user', DATE_SUB(NOW(), INTERVAL 8 HOUR), 3, 45, 12),
(@activity_id, 'def456ghi789', 'fix: 修复用户认证状态问题', 'MyVault', 'demo-user', DATE_SUB(NOW(), INTERVAL 6 HOUR), 2, 23, 8),
(@activity_id, 'ghi789jkl012', 'docs: 更新README文档', 'MyVault', 'demo-user', DATE_SUB(NOW(), INTERVAL 4 HOUR), 1, 15, 3),
(@activity_id, 'jkl012mno345', 'style: 调整CSS样式', 'MyVault', 'demo-user', DATE_SUB(NOW(), INTERVAL 2 HOUR), 4, 32, 18),
(@activity_id, 'mno345pqr678', 'test: 添加单元测试', 'MyVault', 'demo-user', DATE_SUB(NOW(), INTERVAL 1 HOUR), 2, 67, 5);