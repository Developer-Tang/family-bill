## 模块概述

公共服务模块是家庭记账系统的基础支撑组件，为其他业务模块提供共享的底层功能支持。本模块包括认证授权、验证码服务、文件管理、地理位置服务、天气服务和帮助文档等标准化API接口，确保系统功能的统一性和可复用性。

### 设计原则

1. **RESTful规范**：所有API接口严格遵循RESTful设计规范，使用标准HTTP方法和资源导向的URL路径
2. **本地缓存优先**：暂不集成远程缓存机制，如需使用缓存功能，严格限定为本地缓存方案
3. **完整的请求验证**：所有API接口包含严格的请求参数验证，确保数据完整性和安全性
4. **统一的错误处理**：采用统一的错误响应格式和错误码体系，便于客户端处理
5. **清晰的响应格式**：所有API响应采用标准JSON格式，包含明确的数据结构和状态信息
6. **可扩展性**：API设计考虑未来功能扩展，预留合理的扩展点

### 适用场景

- 系统各模块需要统一的认证授权机制
- 用户需要验证码服务进行安全验证
- 系统需要文件上传下载功能
- 用户需要地理位置和天气信息服务
- 用户需要查看帮助文档

## 接口清单

<!-- tabs:start -->
<!-- tab:认证授权 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/auth/login`](#用户登录) | `POST` | 用户登录 |
| [`/api/v1/auth/logout`](#用户登出) | `POST` | 用户登出 |
| [`/api/v1/auth/register`](#用户注册) | `POST` | 用户注册 |
| [`/api/v1/auth/refresh`](#刷新令牌) | `POST` | 刷新令牌 |
| [`/api/v1/auth/verify-email`](#邮箱验证) | `POST` | 邮箱验证 |
| [`/api/v1/auth/reset-password`](#重置密码) | `POST` | 重置密码 |
| [`/api/v1/auth/send-code`](#发送验证码) | `POST` | 发送验证码 |

<!-- tab:验证码服务 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/captcha/image`](#获取图形验证码) | `GET` | 获取图形验证码 |
| [`/api/v1/captcha/verify`](#验证图形验证码) | `POST` | 验证图形验证码 |
| [`/api/v1/captcha/sms`](#发送短信验证码) | `POST` | 发送短信验证码 |
| [`/api/v1/captcha/email`](#发送邮箱验证码) | `POST` | 发送邮箱验证码 |
| [`/api/v1/captcha/voice`](#发送语音验证码) | `POST` | 发送语音验证码 |

<!-- tab:文件服务 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/files/upload`](#上传文件) | `POST` | 上传文件 |
| [`/api/v1/files/:id`](#下载文件) | `GET` | 下载文件 |
| [`/api/v1/files/:id`](#删除文件) | `DELETE` | 删除文件 |
| [`/api/v1/files/preview/:id`](#预览文件) | `GET` | 预览文件 |
| [`/api/v1/files/list`](#获取文件列表) | `GET` | 获取文件列表 |
| [`/api/v1/files/attach/:transactionId`](#关联交易附件) | `POST` | 关联交易附件 |

<!-- tab:帮助文档 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/help/articles`](#获取帮助文章列表) | `GET` | 获取帮助文章列表 |
| [`/api/v1/help/articles/:id`](#获取帮助文章详情) | `GET` | 获取帮助文章详情 |
| [`/api/v1/help/categories`](#获取帮助分类) | `GET` | 获取帮助分类 |
| [`/api/v1/help/search`](#搜索帮助文档) | `GET` | 搜索帮助文档 |
| [`/api/v1/help/feedback`](#提交帮助反馈) | `POST` | 提交帮助反馈 |
<!-- tabs:end -->

## 详细接口说明

### 用户登录

**功能描述**：用户通过邮箱和密码登录系统，获取访问令牌

**请求方法**：POST
**URL路径**：/api/v1/auth/login
**权限要求**：无
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |

**请求体**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| email | string | 是 | - | 用户邮箱 | 有效的邮箱格式 |
| password | string | 是 | - | 用户密码 | 6-20个字符，包含字母和数字 |
| captcha_token | string | 是 | - | 验证码令牌 | 有效的验证码令牌 |
| captcha_code | string | 是 | - | 验证码 | 4位数字 |
| remember_me | boolean | 否 | false | 是否记住登录状态 | 布尔值 |

**请求示例**
```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "captcha_token": "captcha_123",
  "captcha_code": "1234",
  "remember_me": true
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.user | object | 用户信息 |
| data.user.user_id | integer | 用户ID |
| data.user.username | string | 用户名 |
| data.user.email | string | 用户邮箱 |
| data.user.avatar | string | 用户头像URL |
| data.user.created_at | string | 创建时间 |
| data.tokens | object | 认证令牌 |
| data.tokens.access_token | string | 访问令牌 |
| data.tokens.refresh_token | string | 刷新令牌 |
| data.tokens.access_token_expires_in | integer | 访问令牌过期时间（秒） |
| data.tokens.refresh_token_expires_in | integer | 刷新令牌过期时间（秒） |
| data.default_book_id | integer | 默认账本ID |

**响应示例**
```json
{
  "code": 200,
  "message": "登录成功",
  "data": {
    "user": {
      "user_id": 1,
      "username": "张三",
      "email": "user@example.com",
      "avatar": "https://example.com/avatar.jpg",
      "created_at": "2023-01-01T00:00:00Z"
    },
    "tokens": {
      "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "access_token_expires_in": 3600,
      "refresh_token_expires_in": 86400
    },
    "default_book_id": 1
  }
}
```

**错误响应**
```json
{
  "code": 4000,
  "message": "登录失败",
  "data": {
    "error_details": "邮箱或密码错误"
  }
}
```

**本地缓存策略**：登录成功后，本地缓存用户信息和令牌，缓存时间根据remember_me参数决定（true为7天，false为1天）

### 用户登出

**功能描述**：用户退出登录，使当前令牌失效

**请求方法**：POST
**URL路径**：/api/v1/auth/logout
**权限要求**：已登录用户
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |
| Content-Type | string | 是 | application/json |

**请求体**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| refresh_token | string | 是 | - | 刷新令牌 | 有效的刷新令牌 |

**请求示例**
```http
POST /api/v1/auth/logout
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.logout_time | string | 登出时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "登出成功",
  "data": {
    "logout_time": "2023-01-05T12:30:00Z"
  }
}
```

**错误响应**
```json
{
  "code": 401,
  "message": "未授权，需要登录",
  "data": {
    "error_details": "令牌无效或已过期"
  }
}
```

**本地缓存策略**：登出成功后，清除本地缓存的用户信息和令牌

### 刷新令牌

**功能描述**：使用刷新令牌获取新的访问令牌

**请求方法**：POST
**URL路径**：/api/v1/auth/refresh
**权限要求**：无
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | application/json |

**请求体**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| refresh_token | string | 是 | - | 刷新令牌 | 有效的刷新令牌 |

**请求示例**
```http
POST /api/v1/auth/refresh
Content-Type: application/json

{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.access_token | string | 新的访问令牌 |
| data.access_token_expires_in | integer | 访问令牌过期时间（秒） |
| data.refresh_token | string | 新的刷新令牌（可选） |
| data.refresh_token_expires_in | integer | 刷新令牌过期时间（秒）（可选） |

**响应示例**
```json
{
  "code": 200,
  "message": "令牌刷新成功",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "access_token_expires_in": 3600,
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token_expires_in": 86400
  }
}
```

**错误响应**
```json
{
  "code": 401,
  "message": "刷新令牌无效或已过期",
  "data": {
    "error_details": "请重新登录"
  }
}
```

**本地缓存策略**：刷新成功后，更新本地缓存的令牌信息

### 邮箱验证

**功能描述**：验证用户邮箱，激活账号

**请求方法**：POST
**URL路径**：/api/v1/auth/verify-email
**权限要求**：无
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | application/json |

**请求体**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| email | string | 是 | - | 用户邮箱 | 有效的邮箱格式 |
| verification_code | string | 是 | - | 邮箱验证码 | 6位数字 |

**请求示例**
```http
POST /api/v1/auth/verify-email
Content-Type: application/json

{
  "email": "user@example.com",
  "verification_code": "123456"
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.email | string | 验证成功的邮箱 |
| data.verified_at | string | 验证时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "邮箱验证成功",
  "data": {
    "email": "user@example.com",
    "verified_at": "2023-01-05T12:30:00Z"
  }
}
```

**错误响应**
```json
{
  "code": 4002,
  "message": "验证码错误或已过期",
  "data": {
    "error_details": "请重新获取验证码并输入正确的验证码"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 重置密码

**功能描述**：通过邮箱验证码重置用户密码

**请求方法**：POST
**URL路径**：/api/v1/auth/reset-password
**权限要求**：无
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | application/json |

**请求体**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| email | string | 是 | - | 用户邮箱 | 有效的邮箱格式 |
| verification_code | string | 是 | - | 邮箱验证码 | 6位数字 |
| new_password | string | 是 | - | 新密码 | 6-20个字符，包含字母和数字 |
| confirm_password | string | 是 | - | 确认密码 | 必须与new_password相同 |

**请求示例**
```http
POST /api/v1/auth/reset-password
Content-Type: application/json

{
  "email": "user@example.com",
  "verification_code": "123456",
  "new_password": "newpassword123",
  "confirm_password": "newpassword123"
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.email | string | 密码重置成功的邮箱 |
| data.reset_at | string | 密码重置时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "密码重置成功",
  "data": {
    "email": "user@example.com",
    "reset_at": "2023-01-05T12:30:00Z"
  }
}
```

**错误响应**
```json
{
  "code": 4010,
  "message": "密码格式不符合要求",
  "data": {
    "error_details": "密码必须包含字母和数字，长度6-20个字符"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 发送验证码

**功能描述**：发送验证码到用户邮箱或手机

**请求方法**：POST
**URL路径**：/api/v1/auth/send-code
**权限要求**：无
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | application/json |

**请求体**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| contact | string | 是 | - | 联系方式（邮箱或手机号） | 有效的邮箱格式或手机号格式 |
| type | string | 是 | - | 验证码类型 | 枚举值：email, sms |
| purpose | string | 是 | - | 验证码用途 | 枚举值：register, login, reset_password, verify_email |

**请求示例**
```http
POST /api/v1/auth/send-code
Content-Type: application/json

{
  "contact": "user@example.com",
  "type": "email",
  "purpose": "register"
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.contact | string | 发送验证码的联系方式 |
| data.type | string | 验证码类型 |
| data.expires_in | integer | 验证码过期时间（秒） |

**响应示例**
```json
{
  "code": 200,
  "message": "验证码发送成功",
  "data": {
    "contact": "user@example.com",
    "type": "email",
    "expires_in": 300
  }
}
```

**错误响应**
```json
{
  "code": 4008,
  "message": "邮箱发送失败",
  "data": {
    "error_details": "请检查邮箱地址是否正确，稍后重试"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 用户注册

**功能描述**：用户通过邮箱和密码注册新账号

**请求方法**：POST
**URL路径**：/api/v1/auth/register
**权限要求**：无
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |

**请求体**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| username | string | 是 | - | 用户名 | 2-20个字符，支持中文、字母、数字 |
| email | string | 是 | - | 用户邮箱 | 有效的邮箱格式，未被注册 |
| password | string | 是 | - | 用户密码 | 6-20个字符，包含字母和数字 |
| confirm_password | string | 是 | - | 确认密码 | 必须与password相同 |
| phone | string | 否 | - | 手机号码 | 有效的手机号码格式 |
| verification_code | string | 是 | - | 验证码 | 6位数字 |
| captcha_token | string | 是 | - | 验证码令牌 | 有效的验证码令牌 |
| captcha_code | string | 是 | - | 验证码 | 4位数字 |

**请求示例**
```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "username": "张三",
  "email": "user@example.com",
  "password": "password123",
  "confirm_password": "password123",
  "phone": "13800138000",
  "verification_code": "123456",
  "captcha_token": "captcha_123",
  "captcha_code": "5678"
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.user_id | integer | 用户ID |
| data.username | string | 用户名 |
| data.email | string | 用户邮箱 |
| data.created_at | string | 创建时间 |
| data.default_book_created | boolean | 是否创建默认账本 |
| data.verification_status | string | 验证状态：pending, verified |

**响应示例**
```json
{
  "code": 200,
  "message": "注册成功",
  "data": {
    "user_id": 1,
    "username": "张三",
    "email": "user@example.com",
    "created_at": "2023-01-05T12:30:00Z",
    "default_book_created": true,
    "verification_status": "pending"
  }
}
```

**错误响应**
```json
{
  "code": 4001,
  "message": "注册失败",
  "data": {
    "error_details": "邮箱已被注册"
  }
}
```

**本地缓存策略**：注册成功后，本地缓存用户基本信息，缓存时间1小时

### 获取图形验证码

**功能描述**：获取图形验证码，用于登录、注册等场景的安全验证

**请求方法**：GET
**URL路径**：/api/v1/captcha/image
**权限要求**：无
**限流策略**：每分钟60次

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| width | integer | 否 | 120 | 验证码宽度 | 60-200之间的整数 |
| height | integer | 否 | 40 | 验证码高度 | 30-100之间的整数 |
| length | integer | 否 | 4 | 验证码长度 | 4-6之间的整数 |

**请求示例**
```http
GET /api/v1/captcha/image?width=120&height=40&length=4
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.captcha_token | string | 验证码令牌，用于验证时使用 |
| data.image_data | string | 验证码图片的base64编码 |
| data.expires_in | integer | 验证码过期时间（秒） |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "captcha_token": "captcha_123",
    "image_data": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAA...",
    "expires_in": 180
  }
}
```

**错误响应**
```json
{
  "code": 400,
  "message": "请求参数错误",
  "data": {
    "error_details": "验证码宽度必须在60-200之间"
  }
}
```

**本地缓存策略**：图形验证码不缓存，每次请求重新生成

### 验证图形验证码

**功能描述**：验证用户输入的图形验证码是否正确

**请求方法**：POST
**URL路径**：/api/v1/captcha/verify
**权限要求**：无
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | application/json |

**请求体**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| captcha_token | string | 是 | - | 验证码令牌 | 有效的验证码令牌 |
| captcha_code | string | 是 | - | 验证码 | 4-6位字符 |

**请求示例**
```http
POST /api/v1/captcha/verify
Content-Type: application/json

{
  "captcha_token": "captcha_123",
  "captcha_code": "1234"
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.is_valid | boolean | 验证码是否有效 |

**响应示例**
```json
{
  "code": 200,
  "message": "验证码验证成功",
  "data": {
    "is_valid": true
  }
}
```

**错误响应**
```json
{
  "code": 4002,
  "message": "验证码错误或已过期",
  "data": {
    "error_details": "请重新获取验证码并输入正确的验证码"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 发送短信验证码

**功能描述**：发送短信验证码到用户手机

**请求方法**：POST
**URL路径**：/api/v1/captcha/sms
**权限要求**：无
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | application/json |

**请求体**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| phone | string | 是 | - | 手机号码 | 有效的手机号码格式 |
| purpose | string | 是 | - | 验证码用途 | 枚举值：register, login, reset_password, verify_phone |

**请求示例**
```http
POST /api/v1/captcha/sms
Content-Type: application/json

{
  "phone": "13800138000",
  "purpose": "login"
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.phone | string | 发送验证码的手机号码 |
| data.expires_in | integer | 验证码过期时间（秒） |

**响应示例**
```json
{
  "code": 200,
  "message": "短信验证码发送成功",
  "data": {
    "phone": "13800138000",
    "expires_in": 300
  }
}
```

**错误响应**
```json
{
  "code": 4009,
  "message": "手机号格式错误",
  "data": {
    "error_details": "请输入有效的手机号码"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 发送邮箱验证码

**功能描述**：发送邮箱验证码到用户邮箱

**请求方法**：POST
**URL路径**：/api/v1/captcha/email
**权限要求**：无
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | application/json |

**请求体**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| email | string | 是 | - | 用户邮箱 | 有效的邮箱格式 |
| purpose | string | 是 | - | 验证码用途 | 枚举值：register, login, reset_password, verify_email |

**请求示例**
```http
POST /api/v1/captcha/email
Content-Type: application/json

{
  "email": "user@example.com",
  "purpose": "register"
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.email | string | 发送验证码的邮箱 |
| data.expires_in | integer | 验证码过期时间（秒） |

**响应示例**
```json
{
  "code": 200,
  "message": "邮箱验证码发送成功",
  "data": {
    "email": "user@example.com",
    "expires_in": 300
  }
}
```

**错误响应**
```json
{
  "code": 4008,
  "message": "邮箱发送失败",
  "data": {
    "error_details": "请检查邮箱地址是否正确，稍后重试"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 发送语音验证码

**功能描述**：发送语音验证码到用户手机

**请求方法**：POST
**URL路径**：/api/v1/captcha/voice
**权限要求**：无
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | application/json |

**请求体**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| phone | string | 是 | - | 手机号码 | 有效的手机号码格式 |
| purpose | string | 是 | - | 验证码用途 | 枚举值：register, login, reset_password, verify_phone |

**请求示例**
```http
POST /api/v1/captcha/voice
Content-Type: application/json

{
  "phone": "13800138000",
  "purpose": "login"
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.phone | string | 发送验证码的手机号码 |
| data.expires_in | integer | 验证码过期时间（秒） |

**响应示例**
```json
{
  "code": 200,
  "message": "语音验证码发送成功",
  "data": {
    "phone": "13800138000",
    "expires_in": 300
  }
}
```

**错误响应**
```json
{
  "code": 4009,
  "message": "手机号格式错误",
  "data": {
    "error_details": "请输入有效的手机号码"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 上传文件

**功能描述**：上传文件，用于交易记录附件、用户头像等场景

**请求方法**：POST
**URL路径**：/api/v1/files/upload
**权限要求**：已登录用户
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为multipart/form-data |
| Authorization | string | 是 | Bearer JWT令牌 |

**表单数据**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| file | file | 是 | - | 上传的文件 | 支持的文件类型：jpg, jpeg, png, gif, pdf, doc, docx, xls, xlsx |
| book_id | integer | 是 | - | 账本ID | 正整数，账本必须存在且用户有访问权限 |
| entity_type | string | 是 | - | 关联实体类型 | 枚举值：transaction, user, book, category |
| entity_id | integer | 是 | - | 关联实体ID | 正整数，实体必须存在 |
| description | string | 否 | "" | 文件描述 | 0-100个字符 |

**请求示例**
```http
POST /api/v1/files/upload
Content-Type: multipart/form-data
Authorization: Bearer jwt_token_string

file=@/path/to/file.jpg
book_id=1
entity_type=transaction
entity_id=101
description=发票照片
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.file_id | string | 文件ID |
| data.book_id | integer | 账本ID |
| data.user_id | integer | 用户ID |
| data.filename | string | 文件名称 |
| data.original_name | string | 原始文件名 |
| data.mime_type | string | 文件MIME类型 |
| data.size | string | 文件大小（格式化） |
| data.size_bytes | integer | 文件大小（字节） |
| data.path | string | 文件存储路径 |
| data.url | string | 文件访问URL |
| data.thumbnail_url | string | 缩略图URL（仅图片类型） |
| data.entity_type | string | 关联实体类型 |
| data.entity_id | integer | 关联实体ID |
| data.description | string | 文件描述 |
| data.created_at | string | 创建时间 |
| data.updated_at | string | 更新时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "上传成功",
  "data": {
    "file_id": "file_123",
    "book_id": 1,
    "user_id": 1,
    "filename": "发票照片.jpg",
    "original_name": "IMG_1234.jpg",
    "mime_type": "image/jpeg",
    "size": "1.5MB",
    "size_bytes": 1572864,
    "path": "/uploads/files/2023/01/05/file_123.jpg",
    "url": "https://example.com/uploads/files/2023/01/05/file_123.jpg",
    "thumbnail_url": "https://example.com/uploads/files/2023/01/05/thumb_file_123.jpg",
    "entity_type": "transaction",
    "entity_id": 101,
    "description": "发票照片",
    "created_at": "2023-01-05T12:30:00Z",
    "updated_at": "2023-01-05T12:30:00Z"
  }
}
```

**错误响应**
```json
{
  "code": 4005,
  "message": "文件大小超限",
  "data": {
    "error_details": "文件大小不能超过10MB"
  }
}
```

**本地缓存策略**：上传成功后，本地缓存文件基本信息，缓存时间24小时

### 下载文件

**功能描述**：下载指定ID的文件

**请求方法**：GET
**URL路径**：/api/v1/files/:id
**权限要求**：已登录用户
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| id | string | 是 | - | 文件ID | 有效的文件ID |

**请求示例**
```http
GET /api/v1/files/file_123
Authorization: Bearer jwt_token_string
```

**响应**：文件内容，Content-Type根据文件类型自动设置

**错误响应**
```json
{
  "code": 404,
  "message": "文件不存在",
  "data": {
    "error_details": "未找到指定的文件"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接从服务器获取

### 删除文件

**功能描述**：删除指定ID的文件

**请求方法**：DELETE
**URL路径**：/api/v1/files/:id
**权限要求**：已登录用户
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |
| Content-Type | string | 是 | application/json |

**路径参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| id | string | 是 | - | 文件ID | 有效的文件ID |

**请求示例**
```http
DELETE /api/v1/files/file_123
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.file_id | string | 文件ID |
| data.deleted_at | string | 删除时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "文件删除成功",
  "data": {
    "file_id": "file_123",
    "deleted_at": "2023-01-05T12:35:00Z"
  }
}
```

**错误响应**
```json
{
  "code": 403,
  "message": "权限不足",
  "data": {
    "error_details": "您没有权限删除该文件"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 预览文件

**功能描述**：预览指定ID的文件

**请求方法**：GET
**URL路径**：/api/v1/files/preview/:id
**权限要求**：已登录用户
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| id | string | 是 | - | 文件ID | 有效的文件ID |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| width | integer | 否 | - | 预览宽度（仅图片类型） | 正整数 |
| height | integer | 否 | - | 预览高度（仅图片类型） | 正整数 |

**请求示例**
```http
GET /api/v1/files/preview/file_123?width=800&height=600
Authorization: Bearer jwt_token_string
```

**响应**：预览文件内容，Content-Type根据文件类型自动设置

**错误响应**
```json
{
  "code": 400,
  "message": "不支持的文件预览类型",
  "data": {
    "error_details": "该文件类型不支持预览"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接从服务器获取

### 获取文件列表

**功能描述**：获取文件列表，支持按条件筛选

**请求方法**：GET
**URL路径**：/api/v1/files/list
**权限要求**：已登录用户
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| book_id | integer | 是 | - | 账本ID | 正整数，账本必须存在且用户有访问权限 |
| entity_type | string | 否 | - | 关联实体类型 | 枚举值：transaction, user, book, category |
| entity_id | integer | 否 | - | 关联实体ID | 正整数 |
| page | integer | 否 | 1 | 页码 | 正整数 |
| page_size | integer | 否 | 20 | 每页条数 | 1-100之间的整数 |
| type | string | 否 | - | 文件类型 | 枚举值：image, document, other |

**请求示例**
```http
GET /api/v1/files/list?book_id=1&entity_type=transaction&entity_id=101&page=1&page_size=20
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.total | integer | 总文件数 |
| data.page | integer | 当前页码 |
| data.page_size | integer | 每页条数 |
| data.list | array | 文件列表 |
| data.list[].file_id | string | 文件ID |
| data.list[].filename | string | 文件名称 |
| data.list[].original_name | string | 原始文件名 |
| data.list[].mime_type | string | 文件MIME类型 |
| data.list[].size | string | 文件大小（格式化） |
| data.list[].url | string | 文件访问URL |
| data.list[].thumbnail_url | string | 缩略图URL（仅图片类型） |
| data.list[].entity_type | string | 关联实体类型 |
| data.list[].entity_id | integer | 关联实体ID |
| data.list[].description | string | 文件描述 |
| data.list[].created_at | string | 创建时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total": 5,
    "page": 1,
    "page_size": 20,
    "list": [
      {
        "file_id": "file_123",
        "filename": "发票照片.jpg",
        "original_name": "IMG_1234.jpg",
        "mime_type": "image/jpeg",
        "size": "1.5MB",
        "url": "https://example.com/uploads/files/2023/01/05/file_123.jpg",
        "thumbnail_url": "https://example.com/uploads/files/2023/01/05/thumb_file_123.jpg",
        "entity_type": "transaction",
        "entity_id": 101,
        "description": "发票照片",
        "created_at": "2023-01-05T12:30:00Z"
      }
    ]
  }
}
```

**错误响应**
```json
{
  "code": 403,
  "message": "权限不足",
  "data": {
    "error_details": "您没有权限访问该账本的文件列表"
  }
}
```

**本地缓存策略**：缓存1小时

### 关联交易附件

**功能描述**：关联文件到交易记录

**请求方法**：POST
**URL路径**：/api/v1/files/attach/:transactionId
**权限要求**：已登录用户
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |
| Content-Type | string | 是 | application/json |

**路径参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| transactionId | integer | 是 | - | 交易记录ID | 有效的交易记录ID |

**请求体**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| file_ids | array | 是 | - | 文件ID列表 | 有效的文件ID数组 |

**请求示例**
```http
POST /api/v1/files/attach/101
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "file_ids": ["file_123", "file_456"]
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.transaction_id | integer | 交易记录ID |
| data.attached_count | integer | 成功关联的文件数量 |
| data.file_ids | array | 关联的文件ID列表 |

**响应示例**
```json
{
  "code": 200,
  "message": "附件关联成功",
  "data": {
    "transaction_id": 101,
    "attached_count": 2,
    "file_ids": ["file_123", "file_456"]
  }
}
```

**错误响应**
```json
{
  "code": 400,
  "message": "关联失败",
  "data": {
    "error_details": "部分文件关联失败"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 获取帮助文章列表

**功能描述**：获取帮助文章列表，支持按分类、语言等条件筛选

**请求方法**：GET
**URL路径**：/api/v1/help/articles
**权限要求**：无
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 否 | Bearer JWT令牌（可选） |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| category_id | integer | 否 | - | 分类ID | 正整数 |
| page | integer | 否 | 1 | 页码 | 正整数 |
| page_size | integer | 否 | 20 | 每页条数 | 1-100之间的整数 |
| language | string | 否 | "zh-CN" | 语言代码 | 有效的语言代码 |
| keyword | string | 否 | - | 搜索关键词 | 1-50个字符 |
| sort_by | string | 否 | "created_at" | 排序字段 | 枚举值：created_at, updated_at, view_count |
| sort_order | string | 否 | "desc" | 排序顺序 | 枚举值：asc, desc |

**请求示例**
```http
GET /api/v1/help/articles?category_id=1&page=1&page_size=20&language=zh-CN
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.total | integer | 总记录数 |
| data.page | integer | 当前页码 |
| data.page_size | integer | 每页条数 |
| data.articles | array | 文章列表 |
| data.articles[].article_id | integer | 文章ID |
| data.articles[].category_id | integer | 分类ID |
| data.articles[].category_name | string | 分类名称 |
| data.articles[].title | string | 文章标题 |
| data.articles[].subtitle | string | 文章副标题 |
| data.articles[].summary | string | 文章摘要 |
| data.articles[].cover_image | string | 封面图片URL |
| data.articles[].view_count | integer | 浏览次数 |
| data.articles[].created_at | string | 创建时间 |
| data.articles[].updated_at | string | 更新时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total": 50,
    "page": 1,
    "page_size": 20,
    "articles": [
      {
        "article_id": 1,
        "category_id": 1,
        "category_name": "入门指南",
        "title": "如何开始使用家庭记账系统",
        "subtitle": "快速上手，开始记录您的收支",
        "summary": "本文将指导您如何注册账号、创建账本并开始记录第一笔交易...",
        "cover_image": "https://example.com/help/cover1.jpg",
        "view_count": 1500,
        "created_at": "2023-01-01T00:00:00Z",
        "updated_at": "2023-01-02T00:00:00Z"
      }
    ]
  }
}
```

**错误响应**
```json
{
  "code": 400,
  "message": "请求参数错误",
  "data": {
    "error_details": "页码必须为正整数"
  }
}
```

**本地缓存策略**：帮助文章列表本地缓存，缓存时间1小时

### 获取帮助文章详情

**功能描述**：获取指定ID的帮助文章详情

**请求方法**：GET
**URL路径**：/api/v1/help/articles/:id
**权限要求**：无
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 否 | Bearer JWT令牌（可选） |

**路径参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| id | integer | 是 | - | 文章ID | 正整数 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| language | string | 否 | "zh-CN" | 语言代码 | 有效的语言代码 |

**请求示例**
```http
GET /api/v1/help/articles/1?language=zh-CN
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.article_id | integer | 文章ID |
| data.category_id | integer | 分类ID |
| data.category_name | string | 分类名称 |
| data.title | string | 文章标题 |
| data.subtitle | string | 文章副标题 |
| data.summary | string | 文章摘要 |
| data.content | string | 文章内容（HTML格式） |
| data.cover_image | string | 封面图片URL |
| data.view_count | integer | 浏览次数 |
| data.like_count | integer | 点赞次数 |
| data.comment_count | integer | 评论次数 |
| data.created_at | string | 创建时间 |
| data.updated_at | string | 更新时间 |
| data.language | string | 语言代码 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "article_id": 1,
    "category_id": 1,
    "category_name": "入门指南",
    "title": "如何开始使用家庭记账系统",
    "subtitle": "快速上手，开始记录您的收支",
    "summary": "本文将指导您如何注册账号、创建账本并开始记录第一笔交易...",
    "content": "<h1>欢迎使用家庭记账系统</h1><p>...详细内容...</p>",
    "cover_image": "https://example.com/help/cover1.jpg",
    "view_count": 1500,
    "like_count": 150,
    "comment_count": 20,
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-02T00:00:00Z",
    "language": "zh-CN"
  }
}
```

**错误响应**
```json
{
  "code": 404,
  "message": "文章不存在",
  "data": {
    "error_details": "未找到指定的帮助文章"
  }
}
```

**本地缓存策略**：帮助文章详情本地缓存，缓存时间1小时

### 获取帮助分类

**功能描述**：获取帮助文章分类列表

**请求方法**：GET
**URL路径**：/api/v1/help/categories
**权限要求**：无
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 否 | Bearer JWT令牌（可选） |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| language | string | 否 | "zh-CN" | 语言代码 | 有效的语言代码 |

**请求示例**
```http
GET /api/v1/help/categories?language=zh-CN
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | array | 分类列表 |
| data[].category_id | integer | 分类ID |
| data[].name | string | 分类名称 |
| data[].description | string | 分类描述 |
| data[].article_count | integer | 文章数量 |
| data[].created_at | string | 创建时间 |
| data[].updated_at | string | 更新时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": [
    {
      "category_id": 1,
      "name": "入门指南",
      "description": "系统使用入门教程",
      "article_count": 10,
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-02T00:00:00Z"
    }
  ]
}
```

**错误响应**
```json
{
  "code": 500,
  "message": "服务器内部错误",
  "data": {
    "error_details": "获取分类列表失败，请稍后重试"
  }
}
```

**本地缓存策略**：帮助分类列表本地缓存，缓存时间1小时

### 搜索帮助文档

**功能描述**：搜索帮助文章，支持关键词搜索

**请求方法**：GET
**URL路径**：/api/v1/help/search
**权限要求**：无
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 否 | Bearer JWT令牌（可选） |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| keyword | string | 是 | - | 搜索关键词 | 1-50个字符 |
| category_id | integer | 否 | - | 分类ID | 正整数 |
| page | integer | 否 | 1 | 页码 | 正整数 |
| page_size | integer | 否 | 20 | 每页条数 | 1-100之间的整数 |
| language | string | 否 | "zh-CN" | 语言代码 | 有效的语言代码 |

**请求示例**
```http
GET /api/v1/help/search?keyword=记账&category_id=1&page=1&page_size=20&language=zh-CN
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.total | integer | 总搜索结果数 |
| data.page | integer | 当前页码 |
| data.page_size | integer | 每页条数 |
| data.articles | array | 搜索结果列表 |
| data.articles[].article_id | integer | 文章ID |
| data.articles[].category_id | integer | 分类ID |
| data.articles[].category_name | string | 分类名称 |
| data.articles[].title | string | 文章标题 |
| data.articles[].summary | string | 文章摘要 |
| data.articles[].cover_image | string | 封面图片URL |
| data.articles[].view_count | integer | 浏览次数 |
| data.articles[].created_at | string | 创建时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "搜索成功",
  "data": {
    "total": 5,
    "page": 1,
    "page_size": 20,
    "articles": [
      {
        "article_id": 1,
        "category_id": 1,
        "category_name": "入门指南",
        "title": "如何开始使用家庭记账系统",
        "summary": "本文将指导您如何注册账号、创建账本并开始记录第一笔交易...",
        "cover_image": "https://example.com/help/cover1.jpg",
        "view_count": 1500,
        "created_at": "2023-01-01T00:00:00Z"
      }
    ]
  }
}
```

**错误响应**
```json
{
  "code": 400,
  "message": "搜索参数错误",
  "data": {
    "error_details": "搜索关键词不能为空"
  }
}
```

**本地缓存策略**：搜索结果不缓存，每次请求重新搜索

### 提交帮助反馈

**功能描述**：提交帮助反馈或建议

**请求方法**：POST
**URL路径**：/api/v1/help/feedback
**权限要求**：无
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | application/json |
| Authorization | string | 否 | Bearer JWT令牌（可选） |

**请求体**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| type | string | 是 | - | 反馈类型 | 枚举值：question, suggestion, bug, other |
| content | string | 是 | - | 反馈内容 | 10-500个字符 |
| contact | string | 否 | - | 联系方式（邮箱或手机号） | 有效的邮箱格式或手机号格式 |
| user_id | integer | 否 | - | 用户ID（已登录用户自动填充） | 正整数 |

**请求示例**
```http
POST /api/v1/help/feedback
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "type": "suggestion",
  "content": "希望能增加批量导入交易记录的功能",
  "contact": "user@example.com"
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 响应数据 |
| data.feedback_id | string | 反馈ID |
| data.type | string | 反馈类型 |
| data.created_at | string | 提交时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "反馈提交成功",
  "data": {
    "feedback_id": "feedback_123",
    "type": "suggestion",
    "created_at": "2023-01-05T12:30:00Z"
  }
}
```

**错误响应**
```json
{
  "code": 400,
  "message": "反馈内容不符合要求",
  "data": {
    "error_details": "反馈内容长度必须在10-500个字符之间"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

## 数据模型

### 认证令牌模型

```javascript
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "access_token_expires_in": 3600,  // 秒
  "refresh_token_expires_in": 86400, // 秒
  "token_type": "Bearer"
}
```

### 用户信息模型

```javascript
{
  "user_id": 1,
  "username": "张三",
  "email": "user@example.com",
  "phone": "138****8000", // 部分隐藏
  "avatar": "https://example.com/avatar.jpg",
  "created_at": "2023-01-01T00:00:00Z",
  "last_login": "2023-01-05T12:30:00Z",
  "status": "active",  // active, suspended, locked
  "email_verified": true,
  "phone_verified": false,
  "default_book_id": 1
}
```

### 文件信息模型

```javascript
{
  "file_id": "file_123",
  "book_id": 1,
  "user_id": 1,
  "filename": "发票照片.jpg",
  "original_name": "IMG_1234.jpg",
  "mime_type": "image/jpeg",
  "size": "1.5MB",
  "size_bytes": 1572864,
  "path": "/uploads/files/2023/01/05/file_123.jpg",
  "url": "https://example.com/uploads/files/2023/01/05/file_123.jpg",
  "thumbnail_url": "https://example.com/uploads/files/2023/01/05/thumb_file_123.jpg",
  "entity_type": "transaction",
  "entity_id": 101,
  "description": "发票照片",
  "created_at": "2023-01-05T12:30:00Z",
  "updated_at": "2023-01-05T12:30:00Z"
}
```

### 地理位置模型

```javascript
{
  "latitude": 39.9042,
  "longitude": 116.4074,
  "country": "中国",
  "province": "北京市",
  "city": "北京市",
  "district": "东城区",
  "street": "景山前街",
  "address": "北京市东城区景山前街4号",
  "name": "故宫博物院",
  "formatted_address": "中国北京市东城区景山前街4号",
  "postal_code": "100009",
  "location_type": "attraction"
}
```

### 天气信息模型

```javascript
{
  "location": {
    "latitude": 39.9042,
    "longitude": 116.4074,
    "city": "北京市",
    "district": "东城区"
  },
  "current": {
    "temperature": 15,
    "feels_like": 14,
    "humidity": 45,
    "pressure": 1013,
    "wind_speed": 3.5,
    "wind_direction": "东北",
    "condition": "晴",
    "condition_code": "sunny",
    "visibility": 10000,
    "uv_index": 5,
    "air_quality": {
      "aqi": 75,
      "level": "良",
      "pm25": 48,
      "pm10": 65
    },
    "timestamp": "2023-01-05T12:30:00Z"
  }
}
```

### 帮助文章模型

```javascript
{
  "article_id": 1,
  "category_id": 1,
  "category_name": "入门指南",
  "title": "如何开始使用家庭记账系统",
  "subtitle": "快速上手，开始记录您的收支",
  "summary": "本文将指导您如何注册账号、创建账本并开始记录第一笔交易...",
  "content": "<h1>欢迎使用家庭记账系统</h1><p>...详细内容...</p>",
  "cover_image": "https://example.com/help/cover1.jpg",
  "view_count": 1500,
  "like_count": 150,
  "comment_count": 20,
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-02T00:00:00Z",
  "language": "zh-CN",
  "status": "published"
}
```

## 错误码说明

| 错误码 | 描述 |
|-------|------|
| 400 | 请求参数错误 |
| 401 | 未授权，需要登录 |
| 403 | 权限不足 |
| 404 | 资源不存在 |
| 429 | 请求过于频繁 |
| 500 | 服务器内部错误 |
| 4000 | 登录失败 |
| 4001 | 注册失败 |
| 4002 | 验证码错误或已过期 |
| 4003 | 文件上传失败 |
| 4004 | 文件类型不支持 |
| 4005 | 文件大小超限 |
| 4006 | 地理位置解析失败 |
| 4007 | 天气服务调用失败 |
| 4008 | 邮箱发送失败 |
| 4009 | 手机号格式错误 |
| 4010 | 密码格式不符合要求 |