-- 创建数据库
CREATE DATABASE IF NOT EXISTS dididaren DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE dididaren;

-- 用户表
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    phone VARCHAR(20) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    real_name VARCHAR(50),
    id_card VARCHAR(18),
    avatar_url VARCHAR(255),
    gender TINYINT COMMENT '0:未知 1:男 2:女',
    status TINYINT DEFAULT 1 COMMENT '1:正常 0:禁用',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 紧急联系人表
CREATE TABLE emergency_contacts (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    name VARCHAR(50) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    relationship VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- 安保人员表
CREATE TABLE security_staff (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    company_name VARCHAR(100),
    license_number VARCHAR(50),
    certification_status TINYINT DEFAULT 0 COMMENT '0:待审核 1:已认证 2:已拒绝',
    rating DECIMAL(2,1) DEFAULT 5.0,
    total_orders INT DEFAULT 0,
    online_status TINYINT DEFAULT 0 COMMENT '0:离线 1:在线',
    current_location_lat DECIMAL(10,6),
    current_location_lng DECIMAL(10,6),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- 报警事件表
CREATE TABLE emergency_events (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    event_type TINYINT COMMENT '1:普通求助 2:家庭暴力 3:医疗急救 4:诈骗干预',
    status TINYINT DEFAULT 0 COMMENT '0:待处理 1:处理中 2:已完成 3:已取消',
    location_lat DECIMAL(10,6) NOT NULL,
    location_lng DECIMAL(10,6) NOT NULL,
    address VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- 事件处理记录表
CREATE TABLE event_handling_records (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    event_id BIGINT NOT NULL,
    handler_id BIGINT NOT NULL,
    handler_type TINYINT COMMENT '1:志愿者 2:安保人员 3:警察',
    action_type TINYINT COMMENT '1:接单 2:到达 3:处理中 4:完成 5:取消',
    remark TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (event_id) REFERENCES emergency_events(id),
    FOREIGN KEY (handler_id) REFERENCES users(id)
);

-- 危险区域表
CREATE TABLE danger_zones (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    location_lat DECIMAL(10,6) NOT NULL,
    location_lng DECIMAL(10,6) NOT NULL,
    radius INT COMMENT '危险区域半径(米)',
    danger_level TINYINT COMMENT '1:低 2:中 3:高',
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 用户设备信息表
CREATE TABLE user_devices (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    device_id VARCHAR(100),
    device_type VARCHAR(50),
    os_version VARCHAR(50),
    app_version VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- 传感器数据表
CREATE TABLE sensor_data (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    device_id VARCHAR(100),
    data_type TINYINT COMMENT '1:加速度 2:声音 3:光线',
    data_value TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- 评价表
CREATE TABLE ratings (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    event_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    staff_id BIGINT NOT NULL,
    rating TINYINT NOT NULL COMMENT '1-5星',
    comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (event_id) REFERENCES emergency_events(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (staff_id) REFERENCES security_staff(id)
);

-- 系统配置表
CREATE TABLE system_configs (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    config_key VARCHAR(50) NOT NULL UNIQUE,
    config_value TEXT NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
); 