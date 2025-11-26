## 模块概述

用户与权限管理模块是系统的基础安全组件，负责用户身份认证、权限控制和安全审计，确保系统数据安全和精细化访问管理。本模块提供完整的用户生命周期管理、家庭组协作、账本权限控制等功能的标准化API接口。

## 接口清单

<!-- tabs:start -->
<!-- tab:用户管理 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/users/register`](#用户注册) | `POST` | 用户注册 |
| [`/api/v1/users/login`](#用户登录) | `POST` | 用户登录 |
| [`/api/v1/users/reset-password`](#密码找回与重置) | `POST` | 密码找回与重置 |
| [`/api/v1/users/profile`](#获取个人信息) | `GET` | 获取个人信息 |
| [`/api/v1/users/profile`](#更新个人信息) | `PUT` | 更新个人信息 |
| [`/api/v1/users/avatar`](#更新头像) | `POST` | 更新头像 |
| [`/api/v1/users/logout`](#用户登出) | `POST` | 用户登出 |

<!-- tab:家庭组管理 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/families`](#获取家庭组列表) | `GET` | 获取家庭组列表 |
| [`/api/v1/families`](#创建家庭组) | `POST` | 创建家庭组 |
| [`/api/v1/families/:id`](#更新家庭组信息) | `PUT` | 更新家庭组信息 |
| [`/api/v1/families/:id`](#删除家庭组) | `DELETE` | 删除家庭组 |
| [`/api/v1/families/:id/invite`](#邀请成员加入) | `POST` | 邀请成员加入 |
| [`/api/v1/families/:id/leave`](#退出家庭组) | `POST` | 退出家庭组 |
| [`/api/v1/families/:id/members`](#获取成员列表) | `GET` | 获取成员列表 |
| [`/api/v1/families/:id/members/:userId`](#移除成员) | `DELETE` | 移除成员 |

<!-- tab:账本管理 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/books`](#获取账本列表) | `GET` | 获取账本列表 |
| [`/api/v1/books`](#创建账本) | `POST` | 创建账本 |
| [`/api/v1/books/:id`](#更新账本信息) | `PUT` | 更新账本信息 |
| [`/api/v1/books/:id`](#删除账本) | `DELETE` | 删除账本 |
| [`/api/v1/books/:id/access`](#获取账本访问权限) | `GET` | 获取账本访问权限 |

<!-- tab:权限控制 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/books/:id/permissions`](#获取成员权限列表) | `GET` | 获取成员权限列表 |
| [`/api/v1/books/:id/permissions/:userId`](#设置成员权限) | `PUT` | 设置成员权限 |
| [`/api/v1/roles`](#获取角色列表) | `GET` | 获取角色列表 |
| [`/api/v1/roles`](#创建自定义角色) | `POST` | 创建自定义角色 |
| [`/api/v1/roles/:id`](#更新角色) | `PUT` | 更新角色 |
| [`/api/v1/roles/:id`](#删除角色) | `DELETE` | 删除角色 |

<!-- tab:系统安全 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/logs/operations`](#获取操作审计日志) | `GET` | 获取操作审计日志 |
| [`/api/v1/security/2fa`](#设置两步验证) | `POST` | 设置两步验证 |
| [`/api/v1/security/devices`](#获取登录设备列表) | `GET` | 获取登录设备列表 |
| [`/api/v1/security/devices/:id`](#下线异常设备) | `DELETE` | 下线异常设备 |

<!-- tabs:end -->

## 详细接口说明

### 用户注册

**功能描述**：用户注册新账号

**请求参数**

| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| username | string | 是 | - | 用户名 | 3-20个字符，支持字母、数字、下划线 |
| email | string | 否 | - | 邮箱 | 有效的邮箱格式，与phone二选一 |
| phone | string | 否 | - | 手机号 | 11位数字，与email二选一 |
| password | string | 是 | - | 密码 | 8-20个字符，包含字母、数字和特殊字符 |
| confirm_password | string | 是 | - | 确认密码 | 与password一致 |
| verification_code | string | 否 | - | 验证码 | 6位数字，当使用手机号注册时必填 |

**请求**

```http
POST /api/v1/users/register
Content-Type: application/json

{
  "username": "example_user",
  "email": "user@example.com",
  "phone": "13800138000",
  "password": "strong_password",
  "confirm_password": "strong_password",
  "verification_code": "123456"
}
```

**成功响应**

```json
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
```

**错误响应**

```json
{
  "code": 600,
  "message": "注册失败：用户名已存在",
  "data": null
}
```

### 用户登录

**功能描述**：用户登录系统

**请求参数**

| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| account | string | 是 | - | 账号 | 邮箱或手机号 |
| password | string | 否 | - | 密码 | 8-20个字符，当login_type为password时必填 |
| login_type | string | 是 | "password" | 登录类型 | 枚举值：password, verification_code, third_party |
| third_party_info | object | 否 | null | 第三方登录信息 | 当login_type为third_party时必填 |

**请求**

```http
POST /api/v1/users/login
Content-Type: application/json

{
  "account": "user@example.com",
  "password": "secure_password",
  "login_type": "password",
  "third_party_info": null
}
```

**成功响应**

```json
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
```

**错误响应**

```json
{
  "code": 401,
  "message": "登录失败：账号或密码错误",
  "data": null
}
```

### 家庭组创建

**功能描述**：创建新的家庭组

**请求参数**

| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| name | string | 是 | - | 家庭组名称 | 1-20个字符 |
| description | string | 否 | "" | 家庭组描述 | 0-100个字符 |
| avatar | string | 否 | "" | 家庭组头像 | base64编码的图片 |

**请求**

```http
POST /api/v1/families
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "name": "我的家庭",
  "description": "家庭记账群组",
  "avatar": "base64_encoded_image"
}
```

**成功响应**

```json
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

**错误响应**

```json
{
  "code": 400,
  "message": "创建失败：家庭组名称已存在",
  "data": null
}
```

### 账本创建

**功能描述**：创建新的账本

**请求参数**

| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| family_id | integer | 是 | - | 家庭组ID | 正整数 |
| name | string | 是 | - | 账本名称 | 1-20个字符 |
| description | string | 否 | "" | 账本描述 | 0-100个字符 |
| currency | string | 否 | "CNY" | 货币类型 | 有效的货币代码 |
| start_date | string | 是 | - | 开始日期 | YYYY-MM-DD格式 |
| end_date | string | 否 | null | 结束日期 | YYYY-MM-DD格式，永久账本为null |

**请求**

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
  "end_date": null
}
```

**成功响应**

```json
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

**错误响应**

```json
{
  "code": 403,
  "message": "创建失败：权限不足",
  "data": null
}
```

### 设置成员权限

**功能描述**：设置账本成员的权限

**请求参数**

| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| id | integer | 是 | - | 账本ID | 正整数 |
| userId | integer | 是 | - | 用户ID | 正整数 |
| role_id | integer | 是 | - | 角色ID | 1:管理员, 2:记账人, 3:查账人 |
| permissions | array | 否 | - | 权限列表 | 权限字符串数组，当role_id为自定义角色时必填 |

**请求**

```http
PUT /api/v1/books/:id/permissions/:userId
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "role_id": 2,
  "permissions": [
    "read_transactions",
    "create_transactions",
    "update_transactions",
    "delete_transactions"
  ]
}
```

**成功响应**

```json
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

**错误响应**

```json
{
  "code": 403,
  "message": "权限设置失败：您没有权限修改其他成员的权限",
  "data": null
}
```

## 数据模型

### 用户模型

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

### 家庭组模型

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

### 账本模型

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

### 权限模型

```javascript
{
  "permission_id": 1,     // 权限ID
  "name": "read_transactions", // 权限名称
  "display_name": "查看收支", // 显示名称
  "description": "允许查看所有收支记录", // 描述
  "module": "transaction" // 所属模块
}
```

## 权限说明

- **管理员**：完全控制权限，可管理家庭组成员、创建/删除账本、设置所有成员权限
- **记账人**：可创建、编辑、删除收支记录，但不能修改账本设置或其他成员权限
- **查账人**：仅可查看收支记录和统计报表，不能进行任何修改操作

## 错误码说明

| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误 | 检查请求参数是否符合要求，包括类型、格式、必填项等 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有执行该操作的权限 |
| 404 | 资源不存在 | 检查请求的资源ID是否正确 |
| 500 | 服务器内部错误 | 请稍后重试，或联系系统管理员 |
| 600 | 用户已存在 | 用户名、邮箱或手机号已被注册，请更换其他账号 |
| 601 | 用户不存在 | 检查账号是否正确，或先注册账号 |
| 602 | 密码错误 | 检查密码是否正确，或使用忘记密码功能重置 |
| 603 | 验证码错误或已过期 | 请获取新的验证码并重新尝试 |
| 604 | 家庭组已存在 | 家庭组名称已存在，请更换其他名称 |
| 605 | 账本已存在 | 账本名称已存在，请更换其他名称 |

## API版本控制策略

1. **版本号规则**：采用语义化版本号（Major.Minor.Patch）
2. **版本升级策略**：
   - 兼容性功能增加或bug修复，升级Patch版本
   - 新增非破坏性功能，升级Minor版本
   - 破坏性变更，升级Major版本
3. **API废弃规则**：
   - 废弃的API将在文档中明确标记
   - 废弃的API将继续支持至少6个月
   - 废弃的API在响应头中添加`X-API-Deprecated`字段
4. **迁移指南**：
   - 对于破坏性变更，将提供详细的迁移指南
   - 迁移指南包含旧API与新API的映射关系
   - 迁移指南包含代码示例

## API使用规范

1. **请求格式**：所有请求必须使用JSON格式
2. **认证方式**：使用JWT令牌进行认证，令牌放在Authorization头中
3. **请求频率限制**：每个API有请求频率限制，默认每分钟60次
4. **错误处理**：客户端应根据错误码进行相应处理
5. **分页规则**：列表接口支持分页，使用`page`和`page_size`参数
6. **排序规则**：列表接口支持排序，使用`sort_by`和`sort_order`参数