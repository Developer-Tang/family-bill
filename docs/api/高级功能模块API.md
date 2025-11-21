# 高级功能模块 API 文档

##  模块概述

高级功能模块提供家庭记账系统的扩展功能，增强用户体验和系统智能化水平。本模块包括预算管理、标签管理、共享账本、账单提醒、智能分类和汇率换算等功能的API接口。

##  接口清单

| 功能模块 | 接口路径 | 方法 | 功能描述 |
|---------|---------|------|--------|
| **预算管理** | `/api/v1/budgets` | `GET` | 获取预算列表 |
| | `/api/v1/budgets` | `POST` | 创建预算 |
| | `/api/v1/budgets/:id` | `PUT` | 更新预算 |
| | `/api/v1/budgets/:id` | `DELETE` | 删除预算 |
| | `/api/v1/budgets/:id/status` | `GET` | 获取预算执行状态 |
| | `/api/v1/budgets/summary` | `GET` | 获取预算汇总信息 |
| **标签管理** | `/api/v1/tags` | `GET` | 获取标签列表 |
| | `/api/v1/tags` | `POST` | 创建标签 |
| | `/api/v1/tags/:id` | `PUT` | 更新标签 |
| | `/api/v1/tags/:id` | `DELETE` | 删除标签 |
| | `/api/v1/tags/batch` | `POST` | 批量创建标签 |
| | `/api/v1/tags/recommend` | `GET` | 获取推荐标签 |
| **共享账本** | `/api/v1/shared/books` | `GET` | 获取共享账本列表 |
| | `/api/v1/shared/books/:bookId/members` | `GET` | 获取账本成员列表 |
| | `/api/v1/shared/books/:bookId/members` | `POST` | 添加共享成员 |
| | `/api/v1/shared/books/:bookId/members/:userId` | `PUT` | 修改成员权限 |
| | `/api/v1/shared/books/:bookId/members/:userId` | `DELETE` | 移除共享成员 |
| | `/api/v1/shared/invitations` | `GET` | 获取邀请列表 |
| | `/api/v1/shared/invitations` | `POST` | 创建共享邀请 |
| | `/api/v1/shared/invitations/:id` | `PUT` | 处理邀请 |
| **账单提醒** | `/api/v1/reminders` | `GET` | 获取提醒列表 |
| | `/api/v1/reminders` | `POST` | 创建提醒 |
| | `/api/v1/reminders/:id` | `PUT` | 更新提醒 |
| | `/api/v1/reminders/:id` | `DELETE` | 删除提醒 |
| | `/api/v1/reminders/:id/activate` | `PUT` | 激活/停用提醒 |
| | `/api/v1/reminders/upcoming` | `GET` | 获取即将到来的提醒 |
| **智能分类** | `/api/v1/smart/category-predict` | `POST` | 预测交易分类 |
| | `/api/v1/smart/recurring-detection` | `GET` | 检测周期性交易 |
| | `/api/v1/smart/spending-pattern` | `GET` | 获取消费模式分析 |
| | `/api/v1/smart/anomaly-detection` | `GET` | 异常消费检测 |
| **汇率换算** | `/api/v1/currency/rates` | `GET` | 获取汇率列表 |
| | `/api/v1/currency/convert` | `POST` | 货币转换 |
| | `/api/v1/currency/supported` | `GET` | 获取支持的货币列表 |
| | `/api/v1/currency/refresh` | `POST` | 刷新汇率 |

##  详细接口说明

###  创建预算

#### 请求

```http
POST /api/v1/budgets
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "name": "月度餐饮预算",
  "amount": 5000,
  "category_ids": [5, 6],  // 关联的分类IDs
  "period": "monthly",  // monthly, yearly, custom
  "start_date": "2023-01-01",
  "end_date": "2023-01-31",
  "notify_threshold": 80,  // 百分比，达到此值时通知
  "is_recurring": true
}
```

#### 响应

```
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "budget_id": 1,
    "name": "月度餐饮预算",
    "amount": 5000,
    "category_ids": [5, 6],
    "period": "monthly",
    "start_date": "2023-01-01",
    "end_date": "2023-01-31",
    "notify_threshold": 80,
    "is_recurring": true,
    "status": "active",
    "created_at": "2023-01-01T00:00:00Z"
  }
}
```

###  获取预算执行状态

#### 请求

```http
GET /api/v1/budgets/:id/status?date=2023-01-15
Authorization: Bearer jwt_token_string
```

#### 响应

```
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "budget_id": 1,
    "total_budget": 5000,
    "spent_amount": 1500,
    "remaining_amount": 3500,
    "percentage": 30,
    "is_over_budget": false,
    "daily_average": 100,
    "trend": "normal",  // normal, warning, critical
    "transactions_count": 15,
    "forecast": {
      "estimated_spend": 3000,
      "estimated_remaining": 2000,
      "will_exceed": false
    }
  }
}
```

###  创建标签

#### 请求

```http
POST /api/v1/tags
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "name": "生日礼物",
  "color": "#FF5733",
  "icon": "gift",
  "description": "生日礼物相关消费"
}
```

#### 响应

```
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "tag_id": 1,
    "book_id": 1,
    "name": "生日礼物",
    "color": "#FF5733",
    "icon": "gift",
    "description": "生日礼物相关消费",
    "usage_count": 0,
    "created_at": "2023-01-05T12:30:00Z"
  }
}
```

###  添加共享成员

#### 请求

```http
POST /api/v1/shared/books/:bookId/members
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "email": "family_member@example.com",
  "permission": "editor",  // viewer, editor, manager, owner
  "message": "邀请您加入我们的家庭账本"
}
```

#### 响应

```
{
  "code": 200,
  "message": "邀请已发送",
  "data": {
    "invitation_id": "inv_123",
    "book_id": 1,
    "recipient_email": "family_member@example.com",
    "permission": "editor",
    "status": "pending",  // pending, accepted, declined, expired
    "expires_at": "2023-01-12T12:30:00Z",
    "created_at": "2023-01-05T12:30:00Z"
  }
}
```

###  创建提醒

#### 请求

```http
POST /api/v1/reminders
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "title": "房贷还款提醒",
  "amount": 8000,
  "currency": "CNY",
  "type": "fixed_expense",  // income, expense, fixed_income, fixed_expense
  "recurrence": "monthly",  // one_time, weekly, biweekly, monthly, yearly
  "recurrence_day": 15,  // 每月15日
  "start_date": "2023-01-15",
  "end_date": "2030-12-15",
  "reminder_days": [3, 1, 0],  // 提前3天、1天、当天提醒
  "account_id": 1,
  "category_id": 2,
  "auto_create": true,  // 是否自动创建交易记录
  "is_active": true
}
```

#### 响应

```
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "reminder_id": 1,
    "title": "房贷还款提醒",
    "amount": 8000,
    "currency": "CNY",
    "type": "fixed_expense",
    "recurrence": "monthly",
    "recurrence_day": 15,
    "start_date": "2023-01-15",
    "end_date": "2030-12-15",
    "reminder_days": [3, 1, 0],
    "account_id": 1,
    "category_id": 2,
    "auto_create": true,
    "is_active": true,
    "next_occurrence": "2023-02-15",
    "created_at": "2023-01-05T12:30:00Z"
  }
}
```

###  预测交易分类

#### 请求

```http
POST /api/v1/smart/category-predict
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "transaction": {
    "amount": 128.50,
    "description": "超市购物",
    "account_id": 1,
    "date": "2023-01-05",
    "currency": "CNY"
  },
  "include_probability": true
}
```

#### 响应

```
{
  "code": 200,
  "message": "预测成功",
  "data": {
    "predicted_categories": [
      {
        "category_id": 5,
        "name": "日常购物",
        "probability": 0.85
      },
      {
        "category_id": 6,
        "name": "杂货",
        "probability": 0.10
      },
      {
        "category_id": 7,
        "name": "餐饮",
        "probability": 0.05
      }
    ],
    "confidence_level": "high",  // high, medium, low
    "model_version": "1.0.0"
  }
}
```

###  货币转换

#### 请求

```http
POST /api/v1/currency/convert
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "amount": 100,
  "from_currency": "USD",
  "to_currency": "CNY",
  "date": "2023-01-05"  // 可选，不提供则使用最新汇率
}
```

#### 响应

```
{
  "code": 200,
  "message": "转换成功",
  "data": {
    "amount": 100,
    "from_currency": "USD",
    "to_currency": "CNY",
    "converted_amount": 680.50,
    "exchange_rate": 6.805,
    "rate_date": "2023-01-05",
    "rate_source": "央行汇率",
    "last_updated": "2023-01-05T08:00:00Z"
  }
}
```

##  数据模型

###  预算模型

```javascript
{
  "budget_id": 1,            // 预算ID
  "book_id": 1,              // 账本ID
  "name": "月度餐饮预算",      // 预算名称
  "amount": 5000,           // 预算金额
  "category_ids": [5, 6],   // 关联分类ID列表
  "period": "monthly",      // 周期：monthly, yearly, custom
  "start_date": "2023-01-01", // 开始日期
  "end_date": "2023-01-31",   // 结束日期
  "notify_threshold": 80,    // 提醒阈值(百分比)
  "is_recurring": true,      // 是否循环
  "status": "active",       // 状态：active, paused, expired
  "created_at": "2023-01-01T00:00:00Z", // 创建时间
  "updated_at": "2023-01-01T00:00:00Z"  // 更新时间
}
```

###  标签模型

```javascript
{
  "tag_id": 1,                // 标签ID
  "book_id": 1,              // 账本ID
  "name": "生日礼物",         // 标签名称
  "color": "#FF5733",       // 标签颜色
  "icon": "gift",           // 标签图标
  "description": "生日礼物相关消费", // 描述
  "usage_count": 10,         // 使用次数
  "created_at": "2023-01-05T12:30:00Z", // 创建时间
  "updated_at": "2023-01-05T12:30:00Z"  // 更新时间
}
```

###  共享成员模型

```javascript
{
  "member_id": 1,            // 成员ID
  "book_id": 1,              // 账本ID
  "user_id": 2,              // 用户ID
  "email": "family_member@example.com", // 邮箱
  "username": "家人",        // 用户名
  "permission": "editor",   // 权限：viewer, editor, manager, owner
  "joined_at": "2023-01-05T12:30:00Z", // 加入时间
  "last_accessed": "2023-01-05T14:30:00Z", // 最后访问时间
  "status": "active"        // 状态：active, suspended
}
```

###  提醒模型

```javascript
{
  "reminder_id": 1,          // 提醒ID
  "book_id": 1,              // 账本ID
  "title": "房贷还款提醒",     // 提醒标题
  "amount": 8000,           // 金额
  "currency": "CNY",        // 货币
  "type": "fixed_expense",  // 类型：income, expense, fixed_income, fixed_expense
  "recurrence": "monthly",  // 重复：one_time, weekly, biweekly, monthly, yearly
  "recurrence_day": 15,      // 重复日（每月）
  "start_date": "2023-01-15", // 开始日期
  "end_date": "2030-12-15",   // 结束日期
  "reminder_days": [3, 1, 0], // 提前提醒天数
  "account_id": 1,           // 账户ID
  "category_id": 2,          // 分类ID
  "auto_create": true,       // 是否自动创建交易
  "is_active": true,         // 是否激活
  "next_occurrence": "2023-02-15", // 下次发生时间
  "created_at": "2023-01-05T12:30:00Z", // 创建时间
  "updated_at": "2023-01-05T12:30:00Z"  // 更新时间
}
```

###  汇率模型

```javascript
{
  "rate_id": 1,              // 汇率ID
  "from_currency": "USD",   // 源货币
  "to_currency": "CNY",     // 目标货币
  "rate": 6.805,             // 汇率
  "date": "2023-01-05",     // 汇率日期
  "source": "央行汇率",       // 汇率来源
  "last_updated": "2023-01-05T08:00:00Z" // 最后更新时间
}
```

##  高级功能说明

###  预算管理

- 支持按周期设置预算（月度、年度、自定义）
- 可关联多个分类
- 支持自动提醒功能
- 提供预算执行状态跟踪和预警
- 支持循环预算设置

###  标签系统

- 支持自定义颜色和图标
- 每个交易可添加多个标签
- 支持标签批量管理
- 提供标签使用统计
- 智能推荐常用标签

###  共享账本机制

- 支持多用户协作记账
- 细粒度权限控制
- 邀请机制和权限管理
- 操作日志追踪
- 支持家庭和小团队使用

###  智能分析能力

- 基于历史数据的交易分类预测
- 周期性交易自动检测
- 个人消费模式分析
- 异常消费识别和提醒
- 消费趋势预测

##  错误码说明

| 错误码 | 描述 |
|-------|------|
| 400 | 请求参数错误 |
| 401 | 未授权，需要登录 |
| 403 | 权限不足 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |
| 2000 | 预算创建失败 |
| 2001 | 标签创建失败 |
| 2002 | 共享权限不足 |
| 2003 | 提醒创建失败 |
| 2004 | 汇率获取失败 |
| 2005 | 智能分类模型错误 |
| 2006 | 预算已过期 |
| 2007 | 超出预算限额 |
| 2008 | 邀请已过期 |
| 2009 | 提醒频率过高 |
| 2010 | 汇率数据过期 |