# 滴滴打人 - 紧急事件响应平台

## 项目简介
滴滴打人是一个基于地理位置的紧急事件响应平台，致力于打造"15分钟安全守护圈"。平台连接专业安保机构与普通用户，提供全方位的安全保护服务。

## 核心功能

### 1. 智能预警系统
- 实时位置追踪
- 异常行为识别
- 自动环境声光分析
- 危险区域热力图展示

### 2. 多级响应机制
1. 自动定位报警
2. 紧急联系人通知
3. 附近志愿者响应
4. 专业安保人员出动
5. 公安系统联动

### 3. 特色服务
- 夜间护送服务
- 家庭暴力应急响应
- 医疗紧急联动
- 防诈骗干预支持

## 技术架构

### 后端技术栈
- 开发语言：Golang
- 数据库：MySQL
- 缓存：Redis
- 消息队列：RabbitMQ
- 搜索引擎：Elasticsearch
- 对象存储：阿里云OSS
- 地图服务：高德地图API

### 系统模块
1. 用户认证模块
2. 位置服务模块
3. 事件处理模块
4. 支付模块
5. 消息推送模块
6. AI分析模块
7. 评价系统模块

## 数据库设计
详细的数据库设计文档请参考 `docs/database.sql`

## 项目结构
```
dididaren/
├── cmd/                    # 主程序入口
├── config/                 # 配置文件
├── internal/               # 内部包
│   ├── handler/           # HTTP处理器
│   ├── middleware/        # 中间件
│   ├── model/            # 数据模型
│   ├── repository/       # 数据访问层
│   └── service/          # 业务逻辑层
├── pkg/                   # 公共包
├── docs/                  # 文档
├── scripts/              # 脚本文件
└── test/                 # 测试文件
```

## 快速开始

### 环境要求
- Go 1.16+
- MySQL 5.7+
- Redis 6.0+
- RabbitMQ 3.8+

### 安装步骤
1. 克隆项目
```bash
git clone https://github.com/yourusername/dididaren.git
```

2. 安装依赖
```bash
go mod download
```

3. 配置环境变量
```bash
cp .env.example .env
# 编辑 .env 文件，填入必要的配置信息
```

4. 初始化数据库
```bash
mysql -u root -p < docs/database.sql
```

5. 运行项目
```bash
go run cmd/main.go
```

## API文档
API文档请参考 `docs/api.md`

## 贡献指南
1. Fork 项目
2. 创建特性分支
3. 提交变更
4. 推送到分支
5. 创建 Pull Request
