## 模块概述

公共服务模块是家庭记账系统的基础支撑组件，为其他业务模块提供共享的底层功能支持。本模块包括认证授权、验证码服务、文件管理、地理位置服务、天气服务和帮助文档等标准化API接口，确保系统功能的统一性和可复用性。

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

**请求**

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

**响应**

```javascript
// 成功
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

### 用户注册

**请求**

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

**响应**

```javascript
// 成功
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

### 获取图形验证码

**请求**

```http
GET /api/v1/captcha/image?width=120&height=40&length=4
```

**响应**

```javascript
// 成功
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

### 上传文件

**请求**

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

**响应**

```javascript
// 成功
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
    "path": "/uploads/files/2023/01/05/file_123.jpg",
    "url": "https://example.com/uploads/files/2023/01/05/file_123.jpg",
    "thumbnail_url": "https://example.com/uploads/files/2023/01/05/thumb_file_123.jpg",
    "entity_type": "transaction",
    "entity_id": 101,
    "description": "发票照片",
    "created_at": "2023-01-05T12:30:00Z"
  }
}
```

### 获取帮助文章列表

**请求**

```http
GET /api/v1/help/articles?category_id=1&page=1&page_size=20&language=zh-CN
Authorization: Bearer jwt_token_string
```

**响应**

```javascript
// 成功
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
      },
      // 更多文章...
    ]
  }
}
```

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