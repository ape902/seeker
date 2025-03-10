# Seeker - 现代化自动化运维平台

<div align="center">

![Seeker Logo](https://example.com/seeker-logo.png)

**强大、高效、安全的企业级自动化运维解决方案**

[![Go Version](https://img.shields.io/badge/Go-1.22.4-blue.svg)](https://golang.org/)
[![React Version](https://img.shields.io/badge/React-19-blue.svg)](https://reactjs.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

</div>

## 📖 项目简介

Seeker是一个功能全面的现代化自动化运维平台，专为中大型企业IT基础设施管理设计。平台整合了主机管理、文件传输、命令执行、服务发现、对象存储管理等核心运维功能，通过直观的界面和强大的自动化能力，帮助运维团队显著提升工作效率，降低运维成本，保障系统稳定性。

![Seeker平台截图](https://example.com/seeker-screenshot.png)

## ✨ 核心功能

### 🖥️ 主机管理
- **多维度信息管理**：全方位记录和展示主机硬件、系统、网络等信息
- **标签化分类管理**：灵活的标签系统，支持自定义分组和筛选
- **实时状态监控**：CPU、内存、磁盘、网络等关键指标实时监控
- **批量操作支持**：基于标签或自定义规则进行批量主机操作

### 📁 文件管理
- **SFTP文件传输**：安全高效的文件传输协议支持
- **文件上传/下载**：支持断点续传和大文件分片传输
- **文件权限管理**：精细化的文件权限控制
- **批量文件操作**：支持批量上传、下载和权限修改

### 🔄 命令执行
- **远程命令实时执行**：支持WebSocket实时展示命令执行过程
- **批量命令下发**：一键向多台主机分发执行命令
- **命令执行结果记录**：详细记录所有命令执行历史和结果
- **定时任务支持**：支持Cron表达式设置周期性任务

### 🔍 服务发现
- **自动服务发现**：基于规则自动发现网络中的服务
- **进程监控**：关键进程状态和资源占用监控
- **端口监控**：服务端口可用性和连接状态监控
- **服务健康检查**：自定义健康检查规则和告警策略

### 🗄️ 对象存储
- **MinIO对象存储管理**：兼容S3协议的对象存储服务
- **文件存储与共享**：支持文件版本控制和访问链接生成
- **存储桶权限管理**：细粒度的存储桶访问控制
- **文件版本控制**：自动保留文件历史版本

### 👥 用户管理
- **多级权限控制**：基于RBAC的权限管理系统
- **用户行为审计**：详细记录用户操作日志
- **安全认证**：支持多因素认证和SSO集成

## 🛠️ 技术栈

### 后端技术
- **Go 1.22.4**：高性能、并发友好的编程语言
- **Gin Web框架**：轻量级高性能HTTP Web框架
- **gRPC微服务通信**：高效的跨服务通信协议
- **GORM数据库ORM**：功能强大的Go语言ORM库
- **MinIO对象存储**：高性能分布式对象存储
- **JWT身份验证**：安全的用户身份验证机制
- **Prometheus监控**：系统指标收集和监控
- **WebSocket实时通信**：支持实时数据推送和交互

### 前端技术
- **React 19**：高效的用户界面构建库
- **TypeScript**：类型安全的JavaScript超集
- **Vite构建工具**：现代化的前端构建工具
- **Ant Design UI组件库**：企业级UI设计语言和React组件库
- **Redux Toolkit状态管理**：简化的Redux状态管理方案
- **React Router路由管理**：声明式路由管理
- **Axios HTTP客户端**：基于Promise的HTTP客户端

## 🏗️ 系统架构

Seeker采用前后端分离的微服务架构，确保系统的高可用性、可扩展性和可维护性。

- **前端架构**
  - 基于React的单页应用(SPA)
  - 多标签页管理，支持同时操作多个功能模块
  - 响应式设计，适配不同设备和屏幕尺寸
  - 模块化组件设计，提高代码复用性

- **后端架构**
  - 基于Go的微服务架构
  - gRPC通信，提供高性能服务间调用
  - RESTful API接口，便于前端和第三方集成
  - 中间件层处理认证、日志、限流等横切关注点

- **存储层**
  - MySQL数据库：存储结构化业务数据
  - MinIO对象存储：处理非结构化数据如日志、文件等
  - Redis缓存：提高频繁访问数据的响应速度

- **监控系统**
  - Prometheus收集系统和应用指标
  - 自定义Agent采集主机和服务数据
  - 可视化仪表盘展示系统状态

## 🚀 快速开始

### 环境要求
- Go 1.22+
- Node.js 18+
- MySQL 8.0+
- MinIO
- Redis (可选，用于缓存)

### 后端部署

1. **克隆代码库**
   ```bash
   git clone https://github.com/yourusername/seeker.git
   cd seeker
   ```

2. **配置环境**
   ```bash
   # 复制配置文件模板
   cp config/seeker.yaml.example config/seeker.yaml
   
   # 编辑配置文件，设置数据库连接信息、MinIO配置等
   vim config/seeker.yaml
   ```

3. **安装依赖**
   ```bash
   go mod tidy
   ```

4. **初始化数据库**
   ```bash
   # 自动创建数据库表结构和初始数据
   go run cmd/seeker/seeker.go migrate
   ```

5. **启动服务**
   ```bash
   # 开发模式启动
   go run cmd/seeker/seeker.go
   
   # 或使用Makefile
   make run
   
   # 生产环境部署
   make build
   ./bin/seeker
   ```

### 前端部署

1. **进入前端目录**
   ```bash
   cd web
   ```

2. **安装依赖**
   ```bash
   # 使用npm
   npm install
   
   # 或使用yarn
   yarn
   ```

3. **配置环境变量**
   ```bash
   # 开发环境配置
   cp .env.example .env.development.local
   
   # 生产环境配置
   cp .env.example .env.production.local
   
   # 编辑配置文件，设置API地址等
   vim .env.development.local
   ```

4. **启动开发服务器**
   ```bash
   # 使用npm
   npm run dev
   
   # 或使用yarn
   yarn dev
   ```

5. **构建生产版本**
   ```bash
   # 使用npm
   npm run build
   
   # 或使用yarn
   yarn build
   ```

## 📚 使用指南

### 登录系统

使用默认管理员账号登录系统：
- **用户名**：admin
- **密码**：admin123

> **安全提示**：首次登录后请立即修改默认密码，并启用多因素认证以提高账户安全性。

### 主要功能模块

1. **用户中心**
   - 用户管理：创建、编辑、禁用用户账号
   - 角色管理：定义角色和权限
   - 个人设置：修改密码、设置多因素认证
   - 操作日志：查看用户操作历史

2. **CMDB管理**
   - **主机管理**：添加、编辑、删除主机，分组管理
   - **文件管理**：SFTP文件传输，文件浏览和操作
   - **命令执行**：远程执行命令，查看执行结果
   - **服务发现**：自动发现和监控服务

3. **存储管理**
   - 存储桶管理：创建、配置存储桶
   - 文件管理：上传、下载、共享文件
   - 权限控制：设置访问策略和权限

4. **监控告警**
   - 系统监控：查看系统资源使用情况
   - 服务监控：监控关键服务状态
   - 告警配置：设置告警规则和通知方式
   - 告警历史：查看历史告警记录

## 💻 开发指南

### 目录结构

```
├── cmd/                # 主要命令入口
│   ├── agent/          # Agent服务入口
│   ├── engine/         # 任务引擎入口
│   ├── seeker/         # 主服务入口
│   └── skctl/          # 命令行工具
├── config/             # 配置文件
├── pkg/                # 核心包
│   ├── api/            # API定义和接口
│   ├── config/         # 配置加载和管理
│   ├── controller/     # 控制器
│   ├── global/         # 全局变量和常量
│   ├── handler/        # 请求处理器
│   ├── initialize/     # 初始化逻辑
│   ├── models/         # 数据模型
│   ├── router/         # 路由定义
│   └── tools/          # 工具函数和辅助库
└── web/                # 前端代码
    ├── public/         # 静态资源
    └── src/            # 源代码
        ├── components/ # 组件
        ├── config/     # 配置
        ├── hooks/      # 自定义Hooks
        ├── layouts/    # 布局组件
        ├── pages/      # 页面
        ├── router/     # 路由
        ├── styles/     # 样式
        └── utils/      # 工具函数
```

### 开发规范

1. **代码风格**
   - 遵循Go和TypeScript/React的最佳实践
   - 使用ESLint和Prettier保持代码风格一致
   - 遵循项目的编码规范和命名约定

2. **提交规范**
   - 使用语义化的提交消息
   - 每个提交专注于单一功能或修复
   - 提交前运行测试确保代码质量

### 贡献指南

1. Fork 项目仓库
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'feat: add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建Pull Request

## 📄 许可证

本项目采用 [MIT 许可证](LICENSE) 进行许可。

## 📞 联系我们

- **项目维护者**：Your Name
- **邮箱**：your.email@example.com
- **项目主页**：https://github.com/yourusername/seeker

---

<div align="center">

**Seeker** - 让运维工作更简单、更高效

</div>
