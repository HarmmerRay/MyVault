# MyVault - 个人博客网站

一个基于React + Go + MySQL + Redis的个人博客网站，主要功能是展示每日活动时间轴。

## 功能特性

1. **时间轴展示** - 展示每日GitHub提交记录等活动
2. **AI总结** - 使用AI自动总结每日活动
3. **详情页面** - 点击时间轴项目查看详情
4. **每日提醒** - 12点前未提交代码时显示提醒
5. **用户系统** - 登录注册、GitHub授权
6. **响应式设计** - 明亮护眼的渐变动效设计

## 技术栈

- **前端**: React + TypeScript + Tailwind CSS + Framer Motion
- **后端**: Go + Gin + GORM
- **数据库**: MySQL 8.0
- **缓存**: Redis 7.0
- **AI服务**: OpenAI API (可扩展本地模型)

## 项目结构

```
MyVault/
├── frontend/          # React前端
│   ├── src/
│   │   ├── components/    # 公共组件
│   │   ├── pages/         # 页面组件
│   │   ├── contexts/      # React Context
│   │   ├── services/      # API服务
│   │   └── types/         # TypeScript类型定义
│   ├── public/
│   └── package.json
├── backend/           # Go后端
│   ├── cmd/              # 主程序入口
│   ├── internal/
│   │   ├── handlers/     # HTTP处理器
│   │   ├── models/       # 数据模型
│   │   ├── services/     # 业务逻辑
│   │   └── middleware/   # 中间件
│   ├── pkg/              # 外部包
│   │   ├── auth/         # 认证相关
│   │   ├── github/       # GitHub集成
│   │   └── ai/           # AI服务
│   ├── configs/          # 配置文件
│   └── go.mod
├── database/          # 数据库迁移文件
├── docker-compose.yml # Docker配置
└── README.md
```

## 快速开始

### 方式一：使用Docker Compose (推荐)

1. **克隆项目**
   ```bash
   git clone <repository-url>
   cd MyVault
   ```

2. **配置环境变量**
   ```bash
   cp backend/.env.example backend/.env
   # 编辑 backend/.env 文件，配置你的GitHub OAuth和OpenAI API Key
   ```

3. **启动所有服务**
   ```bash
   docker-compose up -d
   ```

4. **访问应用**
   - 前端：http://localhost:3000
   - 后端API：http://localhost:8080

### 方式二：本地开发

#### 环境要求

- Node.js 18+
- Go 1.21+
- MySQL 8.0+
- Redis 7.0+

#### 安装步骤

1. **启动数据库服务**
   ```bash
   # 启动MySQL
   mysql -u root -p < database/init.sql
   
   # 启动Redis
   redis-server
   ```

2. **配置环境变量**
   ```bash
   cp backend/.env.example backend/.env
   # 编辑配置文件
   ```

3. **启动后端服务**
   ```bash
   cd backend
   go mod download
   go run cmd/main.go
   ```

4. **启动前端服务**
   ```bash
   cd frontend
   npm install
   npm run dev
   ```

## 配置说明

### 环境变量

在 `backend/.env` 文件中配置以下变量：

```bash
# 数据库配置
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=myvault

# Redis配置
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# JWT配置
JWT_SECRET=your-secret-key

# 服务端口
PORT=8080

# GitHub OAuth配置
GITHUB_CLIENT_ID=your_github_client_id
GITHUB_CLIENT_SECRET=your_github_client_secret

# OpenAI API配置
OPENAI_API_KEY=your_openai_api_key

# 环境
ENVIRONMENT=development
```

### GitHub OAuth 设置

1. 前往 [GitHub Developer Settings](https://github.com/settings/developers)
2. 创建新的 OAuth App
3. 设置回调URL：`http://localhost:3000/auth/github/callback`
4. 将Client ID和Client Secret配置到环境变量

### OpenAI API 设置

1. 前往 [OpenAI Platform](https://platform.openai.com/)
2. 创建API Key
3. 配置到环境变量中

## API 接口文档

### 认证相关

- `POST /api/auth/register` - 用户注册
- `POST /api/auth/login` - 用户登录
- `GET /api/auth/github` - GitHub OAuth登录
- `GET /api/auth/github/callback` - GitHub OAuth回调

### 用户相关

- `GET /api/user` - 获取用户信息
- `PUT /api/user` - 更新用户信息

### 活动相关

- `GET /api/activities` - 获取活动列表
- `GET /api/activities/:id` - 获取活动详情
- `POST /api/activities/sync` - 同步活动数据

## 开发指南

### 添加新的数据源

1. 在 `backend/pkg/` 中创建新的数据源客户端
2. 在 `backend/internal/services/` 中添加相应的服务
3. 更新 `ActivityService` 以支持新的数据源

### 自定义AI模型

1. 实现 `backend/pkg/ai/` 中的接口
2. 在 `AIService` 中添加新的模型支持
3. 更新配置文件以支持新的AI服务

### 添加新的前端页面

1. 在 `frontend/src/pages/` 中创建新页面
2. 在 `App.tsx` 中添加路由
3. 更新导航组件

## 部署

### 生产环境部署

1. **构建前端**
   ```bash
   cd frontend
   npm run build
   ```

2. **构建后端**
   ```bash
   cd backend
   go build -o main cmd/main.go
   ```

3. **使用Docker部署**
   ```bash
   docker-compose -f docker-compose.prod.yml up -d
   ```

### 注意事项

- 确保数据库和Redis服务正常运行
- 检查防火墙设置，确保端口可访问
- 定期备份数据库数据
- 监控系统资源使用情况

## 故障排除

### 常见问题

1. **数据库连接失败**
   - 检查数据库服务是否运行
   - 验证数据库配置信息
   - 确认网络连接

2. **GitHub OAuth失败**
   - 检查Client ID和Secret是否正确
   - 验证回调URL配置
   - 确认GitHub App权限设置

3. **AI服务调用失败**
   - 检查OpenAI API Key是否有效
   - 验证网络连接
   - 检查API使用配额

## 贡献指南

1. Fork 项目
2. 创建功能分支
3. 提交更改
4. 发起Pull Request

## 许可证

MIT License

## 联系方式

如有问题，请提交Issue或联系开发者。