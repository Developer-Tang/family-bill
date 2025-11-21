# 用户与权限管理模块 API 文档

##  模块概述

用户与权限管理模块是系统的基础，负责用户身份认证、权限控制和安全审计，确保系统数据安全和访问控制。本模块提供了完整的用户生命周期管理、家庭组管理、账本权限控制等功能的API接口。

##  接口清单

| 功能模块 | 接口路径 | 方法 | 功能描述 |
|---------|---------|------|--------|
| **用户管理** | `/api/v1/users/register` | `POST` | 用户注册 |
| | `/api/v1/users/login` | `POST` | 用户登录 |
| | `/api/v1/users/reset-password` | `POST` | 密码找回 |
| | `/api/v1/users/profile` | `GET` | 获取个人信息 |
| | `/api/v1/users/profile` | `PUT` | 更新个人信息 |
| | `/api/v1/users/avatar` | `POST` | 更新头像 |
| | `/api/v1/users/logout` | `POST` | 用户登出 |
| **家庭组管理** | `/api/v1/families` | `GET` | 获取家庭组列表 |
| | `/api/v1/families` | `POST` | 创建家庭组 |
| | `/api/v1/families/:id` | `PUT` | 更新家庭组信息 |
| | `/api/v1/families/:id` | `DELETE` | 删除家庭组 |
| | `/api/v1/families/:id/invite` | `POST` | 邀请成员加入 |
| | `/api/v1/families/:id/leave` | `POST` | 退出家庭组 |
| | `/api/v1/families/:id/members` | `GET` | 获取成员列表 |
| | `/api/v1/families/:id/members/:userId` | `DELETE` | 移除成员 |
| **账本管理** | `/api/v1/books` | `GET` | 获取账本列表 |
| | `/api/v1/books` | `POST` | 创建账本 |
| | `/api/v1/books/:id` | `PUT` | 更新账本信息 |
| | `/api/v1/books/:id` | `DELETE` | 删除账本 |
| | `/api/v1/books/:id/access` | `GET` | 获取账本访问权限 |
| **权限控制** | `/api/v1/books/:id/permissions` | `GET` | 获取成员权限 |
| | `/api/v1/books/:id/permissions/:userId` | `PUT` | 设置成员权限 |
| **角色管理** | `/api/v1/roles` | `GET` | 获取角色列表 |
| | `/api/v1/roles` | `POST` | 创建角色 |
| | `/api/v1/roles/:id` | `PUT` | 更新角色 |
| | `/api/v1/roles/:id` | `DELETE` | 删除角色 |
| **操作日志** | `/api/v1/logs/operations` | `GET` | 获取操作日志 |
| **安全设置** | `/api/v1/security/2fa` | `POST` | 设置两步验证 |
| | `/api/v1/security/devices` | `GET` | 获取登录设备列表 |
| | `/api/v1/security/devices/:id` | `DELETE` | 下线设备 |

##  详细接口说明

###  用户注册

#### 请求

```http
POST /api/v1/users/register
Content-Type: application/json

{
  "username": "example_user",
  "email": "user@example.com",
  "phone": "13800138000",
  "password": "secure_password",
  "verification_code": "123456",
  "source": "mobile/web"
}
```

#### 响应

```
# 成功
{
  "code": 200,
  "message": "注册成功",
  "data": {
    "user_id": 1,
    "username": "example_user",
    "email": "user@example.com",
    "phone": "138****8000",
    "token": "jwt_token_string"
  }
}

# 失败
{
  "code": 400,
  "message": "注册失败：用户名已存在",
  "data": null
}
```

###  用户登录

#### 请求

```http
POST /api/v1/users/login
Content-Type: application/json

{
  "account": "user@example.com",  // 邮箱或手机号
  "password": "secure_password",
  "login_type": "password",  // password, verification_code, third_party
  "third_party_info": null    // 第三方登录信息，可选
}
```

#### 响应

```
# 成功
{
  "code": 200,
  "message": "登录成功",
  "data": {
    "user_id": 1,
    "username": "example_user",
    "token": "jwt_token_string",
    "refresh_token": "refresh_token_string",
    "expires_in": 7200
  }
}

# 失败
{
  "code": 401,
  "message": "登录失败：账号或密码错误",
  "data": null
}
```

###  家庭组创建

#### 请求

```http
POST /api/v1/families
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "name": "我的家庭",
  "description": "家庭记账群组",
  "avatar": "base64_encoded_image"  // 可选
}
```

#### 响应

```
# 成功
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "family_id": 1,
    "name": "我的家庭",
    "description": "家庭记账群组",
    "created_at": "2023-01-01T12:00:00Z",
    "members_count": 1
  }
}
```

###  账本创建

#### 请求

```http
POST /api/v1/books
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "family_id": 1,
  "name": "家庭日常账本",
  "description": "记录家庭日常收支",
  "currency": "CNY",
  "start_date": "2023-01-01",
  "end_date": null  // 永久账本为null
}
```

#### 响应

```
# 成功
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "book_id": 1,
    "name": "家庭日常账本",
    "currency": "CNY",
    "created_at": "2023-01-01T12:00:00Z"
  }
}
```

###  设置成员权限

#### 请求

```http
PUT /api/v1/books/:id/permissions/:userId
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "role_id": 2,  // 1:管理员, 2:记账人, 3:查账人
  "permissions": [
    "read_transactions",
    "create_transactions",
    "update_transactions",
    "delete_transactions"
  ]
}
```

#### 响应

```
# 成功
{
  "code": 200,
  "message": "权限设置成功",
  "data": {
    "book_id": 1,
    "user_id": 2,
    "role_id": 2,
    "role_name": "记账人",
    "permissions": ["read_transactions", "create_transactions", "update_transactions"]
  }
}
```

##  数据模型

###  用户模型

```javascript
{
  "user_id": 1,           // 用户ID
  "username": "张三",    // 用户名
  "email": "user@example.com", // 邮箱
  "phone": "138****8000", // 手机号（部分隐藏）
  "avatar": "url_to_avatar", // 头像URL
  "created_at": "2023-01-01T12:00:00Z", // 创建时间
  "last_login_at": "2023-01-10T15:30:00Z" // 最后登录时间
}
```

###  家庭组模型

```javascript
{
  "family_id": 1,         // 家庭组ID
  "name": "我的家庭",    // 家庭组名称
  "description": "家庭记账群组", // 描述
  "avatar": "url_to_avatar", // 头像URL
  "creator_id": 1,        // 创建者ID
  "created_at": "2023-01-01T12:00:00Z", // 创建时间
  "members_count": 4      // 成员数量
}
```

###  账本模型

```javascript
{
  "book_id": 1,           // 账本ID
  "family_id": 1,         // 所属家庭组ID
  "name": "家庭日常账本", // 账本名称
  "description": "记录家庭日常收支", // 描述
  "currency": "CNY",     // 货币类型
  "start_date": "2023-01-01", // 开始日期
  "end_date": null,       // 结束日期（永久账本为null）
  "created_at": "2023-01-01T12:00:00Z" // 创建时间
}
```

###  权限模型

```javascript
{
  "permission_id": 1,     // 权限ID
  "name": "read_transactions", // 权限名称
  "display_name": "查看收支", // 显示名称
  "description": "允许查看所有收支记录", // 描述
  "module": "transaction" // 所属模块
}
```

##  权限说明

- **管理员**：完全控制权限，可管理家庭组成员、创建/删除账本、设置所有成员权限
- **记账人**：可创建、编辑、删除收支记录，但不能修改账本设置或其他成员权限
- **查账人**：仅可查看收支记录和统计报表，不能进行任何修改操作

##  错误码说明

| 错误码 | 描述 |
|-------|------|
| 400 | 请求参数错误 |
| 401 | 未授权，需要登录 |
| 403 | 权限不足 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |
| 600 | 用户已存在 |
| 601 | 用户不存在 |
| 602 | 密码错误 |
| 603 | 验证码错误或已过期 |
| 604 | 家庭组已存在 |
| 605 | 账本已存在 |