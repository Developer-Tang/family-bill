# 系统管理模块 API 文档

##  模块概述

系统管理模块提供家庭记账系统的核心管理功能，包括用户设置、分类管理、账本管理、系统日志、系统配置和通知管理等。本模块为用户提供个性化设置、分类管理和系统状态监控等基础功能的API接口。

##  接口清单

| 功能模块 | 接口路径 | 方法 | 功能描述 |
|---------|---------|------|--------|
| **用户设置** | `/api/v1/user/settings` | `GET` | 获取用户设置 |
| | `/api/v1/user/settings` | `PUT` | 更新用户设置 |
| | `/api/v1/user/profile` | `GET` | 获取用户资料 |
| | `/api/v1/user/profile` | `PUT` | 更新用户资料 |
| | `/api/v1/user/avatar` | `POST` | 上传用户头像 |
| | `/api/v1/user/password` | `PUT` | 修改密码 |
| **分类管理** | `/api/v1/categories` | `GET` | 获取分类列表 |
| | `/api/v1/categories` | `POST` | 创建分类 |
| | `/api/v1/categories/:id` | `PUT` | 更新分类 |
| | `/api/v1/categories/:id` | `DELETE` | 删除分类 |
| | `/api/v1/categories/import` | `POST` | 导入分类 |
| | `/api/v1/categories/export` | `GET` | 导出分类 |
| **账本管理** | `/api/v1/books` | `GET` | 获取账本列表 |
| | `/api/v1/books` | `POST` | 创建账本 |
| | `/api/v1/books/:id` | `GET` | 获取账本详情 |
| | `/api/v1/books/:id` | `PUT` | 更新账本 |
| | `/api/v1/books/:id` | `DELETE` | 删除账本 |
| | `/api/v1/books/:id/switch` | `POST` | 切换当前账本 |
| **系统日志** | `/api/v1/logs/operation` | `GET` | 获取操作日志 |
| | `/api/v1/logs/error` | `GET` | 获取错误日志 |
| | `/api/v1/logs/system` | `GET` | 获取系统日志 |
| | `/api/v1/logs/user-activity` | `GET` | 获取用户活动日志 |
| **系统配置** | `/api/v1/system/config` | `GET` | 获取系统配置 |
| | `/api/v1/system/config` | `PUT` | 更新系统配置 |
| | `/api/v1/system/info` | `GET` | 获取系统信息 |
| | `/api/v1/system/check-update` | `GET` | 检查更新 |
| | `/api/v1/system/version` | `GET` | 获取版本信息 |
| **通知管理** | `/api/v1/notifications` | `GET` | 获取通知列表 |
| | `/api/v1/notifications/:id` | `PUT` | 标记通知已读 |
| | `/api/v1/notifications/read-all` | `PUT` | 标记所有通知已读 |
| | `/api/v1/notifications/:id` | `DELETE` | 删除通知 |
| | `/api/v1/notifications/settings` | `GET` | 获取通知设置 |
| | `/api/v1/notifications/settings` | `PUT` | 更新通知设置 |

##  详细接口说明

###  获取用户设置

#### 请求

```http
GET /api/v1/user/settings
Authorization: Bearer jwt_token_string
```

#### 响应

```
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "settings_id": 1,
    "user_id": 1,
    "theme": "light",  // light, dark, system
    "language": "zh-CN",
    "date_format": "YYYY-MM-DD",
    "time_format": "24h",  // 12h, 24h
    "currency_format": "symbol_before",  // symbol_before, symbol_after
    "decimal_places": 2,
    "auto_sync": true,
    "notifications": {
      "email": true,
      "push": true,
      "budget_alerts": true,
      "reminders": true,
      "updates": true
    },
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-05T12:30:00Z"
  }
}
```

###  更新用户设置

#### 请求

```http
PUT /api/v1/user/settings
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "theme": "dark",
  "language": "zh-CN",
  "date_format": "YYYY/MM/DD",
  "time_format": "24h",
  "currency_format": "symbol_before",
  "decimal_places": 2,
  "auto_sync": true,
  "notifications": {
    "email": true,
    "push": false,
    "budget_alerts": true,
    "reminders": true,
    "updates": false
  }
}
```

#### 响应

```
{
  "code": 200,
  "message": "更新成功",
  "data": {
    "settings_id": 1,
    "user_id": 1,
    "theme": "dark",
    "language": "zh-CN",
    "date_format": "YYYY/MM/DD",
    "time_format": "24h",
    "currency_format": "symbol_before",
    "decimal_places": 2,
    "auto_sync": true,
    "notifications": {
      "email": true,
      "push": false,
      "budget_alerts": true,
      "reminders": true,
      "updates": false
    },
    "updated_at": "2023-01-05T12:35:00Z"
  }
}
```

###  创建分类

#### 请求

```http
POST /api/v1/categories
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "name": "日常购物",
  "type": "expense",  // income, expense, transfer
  "parent_id": null,  // 父分类ID，顶级分类为null
  "icon": "shopping-bag",
  "color": "#4CAF50",
  "sort_order": 1,
  "is_system": false,
  "description": "日常购物消费"
}
```

#### 响应

```
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "category_id": 1,
    "book_id": 1,
    "name": "日常购物",
    "type": "expense",
    "parent_id": null,
    "icon": "shopping-bag",
    "color": "#4CAF50",
    "sort_order": 1,
    "is_system": false,
    "description": "日常购物消费",
    "created_at": "2023-01-05T12:30:00Z",
    "updated_at": "2023-01-05T12:30:00Z"
  }
}
```

###  创建账本

#### 请求

```http
POST /api/v1/books
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "name": "我的家庭账本",
  "description": "记录家庭日常收支",
  "currency": "CNY",
  "start_date": "2023-01-01",
  "is_default": true,
  "auto_backup": true,
  "template_id": null  // 使用模板ID，无模板为null
}
```

#### 响应

```
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "book_id": 1,
    "name": "我的家庭账本",
    "description": "记录家庭日常收支",
    "currency": "CNY",
    "start_date": "2023-01-01",
    "is_default": true,
    "auto_backup": true,
    "owner_id": 1,
    "created_at": "2023-01-05T12:30:00Z",
    "updated_at": "2023-01-05T12:30:00Z"
  }
}
```

###  获取操作日志

#### 请求

```http
GET /api/v1/logs/operation?book_id=1&start_date=2023-01-01&end_date=2023-01-31&page=1&page_size=20&action_type=create
Authorization: Bearer jwt_token_string
```

#### 响应

```
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total": 150,
    "page": 1,
    "page_size": 20,
    "logs": [
      {
        "log_id": 1,
        "book_id": 1,
        "user_id": 1,
        "username": "张三",
        "action_type": "create",  // create, update, delete, login, logout
        "entity_type": "transaction",  // transaction, account, category, etc.
        "entity_id": 101,
        "details": "创建交易记录：日常购物 100元",
        "ip_address": "192.168.1.1",
        "user_agent": "Mozilla/5.0...",
        "created_at": "2023-01-05T12:30:00Z"
      },
      // 更多日志...
    ]
  }
}
```

###  获取系统信息

#### 请求

```http
GET /api/v1/system/info
Authorization: Bearer jwt_token_string
```

#### 响应

```
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "system": {
      "version": "1.0.0",
      "build_date": "2023-01-01T00:00:00Z",
      "environment": "production",
      "uptime": "30d 5h 12m",
      "database": {
        "type": "SQLite",
        "version": "3.36.0",
        "size": "25.6MB",
        "status": "healthy"
      },
      "storage": {
        "total": "10GB",
        "used": "2.5GB",
        "free": "7.5GB"
      },
      "server": {
        "name": "family-bill-server",
        "os": "Linux Ubuntu 20.04",
        "memory": {
          "total": "8GB",
          "used": "2.4GB",
          "free": "5.6GB"
        }
      }
    }
  }
}
```

###  获取通知列表

#### 请求

```http
GET /api/v1/notifications?page=1&page_size=20&status=unread
Authorization: Bearer jwt_token_string
```

#### 响应

```
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total": 10,
    "page": 1,
    "page_size": 20,
    "unread_count": 3,
    "notifications": [
      {
        "notification_id": 1,
        "user_id": 1,
        "title": "预算提醒",
        "content": "您的月度餐饮预算已使用80%",
        "type": "budget_alert",  // budget_alert, reminder, system, shared
        "is_read": false,
        "is_pinned": false,
        "related_id": 1,  // 相关实体ID
        "created_at": "2023-01-05T12:30:00Z"
      },
      // 更多通知...
    ]
  }
}
```

##  数据模型

###  用户设置模型

```javascript
{
  "settings_id": 1,          // 设置ID
  "user_id": 1,              // 用户ID
  "theme": "light",         // 主题：light, dark, system
  "language": "zh-CN",      // 语言
  "date_format": "YYYY-MM-DD", // 日期格式
  "time_format": "24h",     // 时间格式：12h, 24h
  "currency_format": "symbol_before", // 货币格式
  "decimal_places": 2,       // 小数位数
  "auto_sync": true,         // 是否自动同步
  "notifications": {         // 通知设置
    "email": true,           // 邮件通知
    "push": true,            // 推送通知
    "budget_alerts": true,   // 预算提醒
    "reminders": true,       // 账单提醒
    "updates": true          // 更新通知
  },
  "created_at": "2023-01-01T00:00:00Z", // 创建时间
  "updated_at": "2023-01-05T12:30:00Z"  // 更新时间
}
```

###  分类模型

```javascript
{
  "category_id": 1,          // 分类ID
  "book_id": 1,              // 账本ID
  "name": "日常购物",         // 分类名称
  "type": "expense",        // 类型：income, expense, transfer
  "parent_id": null,         // 父分类ID
  "icon": "shopping-bag",   // 图标
  "color": "#4CAF50",       // 颜色
  "sort_order": 1,           // 排序顺序
  "is_system": false,        // 是否系统分类
  "description": "日常购物消费", // 描述
  "created_at": "2023-01-05T12:30:00Z", // 创建时间
  "updated_at": "2023-01-05T12:30:00Z"  // 更新时间
}
```

###  账本模型

```javascript
{
  "book_id": 1,              // 账本ID
  "name": "我的家庭账本",     // 账本名称
  "description": "记录家庭日常收支", // 描述
  "currency": "CNY",        // 默认货币
  "start_date": "2023-01-01", // 开始日期
  "is_default": true,        // 是否默认账本
  "auto_backup": true,       // 是否自动备份
  "owner_id": 1,             // 所有者ID
  "created_at": "2023-01-05T12:30:00Z", // 创建时间
  "updated_at": "2023-01-05T12:30:00Z"  // 更新时间
}
```

###  系统日志模型

```javascript
{
  "log_id": 1,               // 日志ID
  "book_id": 1,              // 账本ID（可选）
  "user_id": 1,              // 用户ID
  "username": "张三",        // 用户名
  "action_type": "create",  // 操作类型
  "entity_type": "transaction", // 实体类型
  "entity_id": 101,          // 实体ID
  "details": "创建交易记录：日常购物 100元", // 详细信息
  "ip_address": "192.168.1.1", // IP地址
  "user_agent": "Mozilla/5.0...", // 用户代理
  "created_at": "2023-01-05T12:30:00Z" // 创建时间
}
```

###  通知模型

```javascript
{
  "notification_id": 1,      // 通知ID
  "user_id": 1,              // 用户ID
  "title": "预算提醒",       // 标题
  "content": "您的月度餐饮预算已使用80%", // 内容
  "type": "budget_alert",   // 类型
  "is_read": false,          // 是否已读
  "is_pinned": false,        // 是否置顶
  "related_id": 1,           // 相关实体ID
  "created_at": "2023-01-05T12:30:00Z" // 创建时间
}
```

##  系统管理功能说明

###  用户设置管理

- 支持个性化主题设置（亮色、暗色、跟随系统）
- 多语言支持和本地化设置
- 自定义日期时间和货币显示格式
- 通知偏好设置
- 同步和数据安全设置

###  分类管理系统

- 支持多层级分类管理
- 自定义图标和颜色
- 分类排序和状态管理
- 导入导出功能
- 预设系统分类模板

###  账本管理功能

- 多账本支持
- 账本基本信息管理
- 默认账本设置
- 账本切换功能
- 账本数据统计概览

###  日志和审计功能

- 操作日志记录和查询
- 用户活动追踪
- 系统事件监控
- 错误日志记录
- 日志搜索和过滤

###  系统配置和监控

- 系统版本和更新管理
- 数据库状态监控
- 存储使用情况
- 服务器资源监控
- 系统健康检查

##  错误码说明

| 错误码 | 描述 |
|-------|------|
| 400 | 请求参数错误 |
| 401 | 未授权，需要登录 |
| 403 | 权限不足 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |
| 3000 | 设置更新失败 |
| 3001 | 分类创建失败 |
| 3002 | 账本创建失败 |
| 3003 | 头像上传失败 |
| 3004 | 密码修改失败 |
| 3005 | 日志查询失败 |
| 3006 | 分类删除失败（已使用） |
| 3007 | 账本删除失败（有成员） |
| 3008 | 通知设置更新失败 |
| 3009 | 系统配置更新失败 |
| 3010 | 版本检查失败 |