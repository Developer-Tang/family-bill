##  模块概述

高级功能模块提供家庭记账系统的扩展功能，增强用户体验和系统智能化水平。本模块包括预算管理、标签管理、共享账本、账单提醒、智能分类和汇率换算等功能的API接口。

##  接口清单

<!-- tabs:start -->
<!-- tab:预算管理 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/budgets`](#获取预算列表) | `GET` | 获取预算列表 |
| [`/api/v1/budgets`](#创建预算) | `POST` | 创建预算 |
| [`/api/v1/budgets/:id`](#更新预算) | `PUT` | 更新预算 |
| [`/api/v1/budgets/:id`](#删除预算) | `DELETE` | 删除预算 |
| [`/api/v1/budgets/:id/status`](#获取预算执行状态) | `GET` | 获取预算执行状态 |
| [`/api/v1/budgets/summary`](#获取预算汇总信息) | `GET` | 获取预算汇总信息 |

<!-- tab:标签管理 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/tags`](#获取标签列表) | `GET` | 获取标签列表 |
| [`/api/v1/tags`](#创建标签) | `POST` | 创建标签 |
| [`/api/v1/tags/:id`](#更新标签) | `PUT` | 更新标签 |
| [`/api/v1/tags/:id`](#删除标签) | `DELETE` | 删除标签 |
| [`/api/v1/tags/batch`](#批量创建标签) | `POST` | 批量创建标签 |
| [`/api/v1/tags/recommend`](#获取推荐标签) | `GET` | 获取推荐标签 |

<!-- tab:共享账本 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/shared/books`](#获取共享账本列表) | `GET` | 获取共享账本列表 |
| [`/api/v1/shared/books/:bookId/members`](#获取账本成员列表) | `GET` | 获取账本成员列表 |
| [`/api/v1/shared/books/:bookId/members`](#添加共享成员) | `POST` | 添加共享成员 |
| [`/api/v1/shared/books/:bookId/members/:userId`](#修改成员权限) | `PUT` | 修改成员权限 |
| [`/api/v1/shared/books/:bookId/members/:userId`](#移除共享成员) | `DELETE` | 移除共享成员 |
| [`/api/v1/shared/invitations`](#获取邀请列表) | `GET` | 获取邀请列表 |
| [`/api/v1/shared/invitations`](#创建共享邀请) | `POST` | 创建共享邀请 |
| [`/api/v1/shared/invitations/:id`](#处理邀请) | `PUT` | 处理邀请 |

<!-- tab:账单提醒 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/reminders`](#获取提醒列表) | `GET` | 获取提醒列表 |
| [`/api/v1/reminders`](#创建提醒) | `POST` | 创建提醒 |
| [`/api/v1/reminders/:id`](#更新提醒) | `PUT` | 更新提醒 |
| [`/api/v1/reminders/:id`](#删除提醒) | `DELETE` | 删除提醒 |
| [`/api/v1/reminders/:id/activate`](#激活停用提醒) | `PUT` | 激活/停用提醒 |
| [`/api/v1/reminders/upcoming`](#获取即将到来的提醒) | `GET` | 获取即将到来的提醒 |

<!-- tab:智能分类 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/smart/category-predict`](#预测交易分类) | `POST` | 预测交易分类 |
| [`/api/v1/smart/recurring-detection`](#检测周期性交易) | `GET` | 检测周期性交易 |
| [`/api/v1/smart/spending-pattern`](#获取消费模式分析) | `GET` | 获取消费模式分析 |
| [`/api/v1/smart/anomaly-detection`](#异常消费检测) | `GET` | 异常消费检测 |

<!-- tab:汇率换算 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/currency/rates`](#获取汇率列表) | `GET` | 获取汇率列表 |
| [`/api/v1/currency/convert`](#货币转换) | `POST` | 货币转换 |
| [`/api/v1/currency/supported`](#获取支持的货币列表) | `GET` | 获取支持的货币列表 |
| [`/api/v1/currency/refresh`](#刷新汇率) | `POST` | 刷新汇率 |

<!-- tabs:end -->

##  详细接口说明

###  创建预算

**功能描述**：创建新的预算记录

**请求参数**

| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| book_id | integer | 是 | - | 账本ID | 正整数 |
| name | string | 是 | - | 预算名称 | 1-50个字符 |
| amount | number | 是 | - | 预算金额 | 大于0 |
| category_ids | array | 否 | [] | 关联的分类IDs | 数组元素为正整数 |
| period | string | 是 | - | 预算周期 | 枚举值：monthly, yearly, custom |
| start_date | string | 是 | - | 开始日期 | YYYY-MM-DD格式 |
| end_date | string | 否 | - | 结束日期 | YYYY-MM-DD格式，仅当period为custom时必填 |
| notify_threshold | integer | 否 | 80 | 通知阈值（百分比） | 1-100之间的整数 |
| is_recurring | boolean | 否 | false | 是否循环 | 布尔值 |

**请求**

```http
POST /api/v1/budgets
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "name": "月度餐饮预算",
  "amount": 5000,
  "category_ids": [5, 6],
  "period": "monthly",
  "start_date": "2023-01-01",
  "end_date": "2023-01-31",
  "notify_threshold": 80,
  "is_recurring": true
}
```

**成功响应**

```json
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

**错误响应**

```json
{
  "code": 400,
  "message": "请求参数错误：预算金额必须大于0",
  "data": null
}
```

###  获取预算执行状态

**功能描述**：获取指定预算的执行状态

**请求参数**

| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| id | integer | 是 | - | 预算ID | 正整数 |
| date | string | 否 | 当前日期 | 查询日期 | YYYY-MM-DD格式 |

**请求**

```http
GET /api/v1/budgets/:id/status?date=2023-01-15
Authorization: Bearer jwt_token_string
```

**成功响应**

```json
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
    "trend": "normal",
    "transactions_count": 15,
    "forecast": {
      "estimated_spend": 3000,
      "estimated_remaining": 2000,
      "will_exceed": false
    }
  }
}
```

**错误响应**

```json
{
  "code": 404,
  "message": "预算不存在",
  "data": null
}
```

###  创建标签

**功能描述**：创建新的标签

**请求参数**

| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| book_id | integer | 是 | - | 账本ID | 正整数 |
| name | string | 是 | - | 标签名称 | 1-30个字符 |
| color | string | 否 | "#3366FF" | 标签颜色 | 十六进制颜色代码 |
| icon | string | 否 | "tag" | 标签图标 | 1-20个字符 |
| description | string | 否 | "" | 标签描述 | 0-100个字符 |

**请求**

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

**成功响应**

```json
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

**错误响应**

```json
{
  "code": 400,
  "message": "请求参数错误：标签名称已存在",
  "data": null
}
```

###  添加共享成员

**功能描述**：邀请成员加入共享账本

**请求参数**

| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| bookId | integer | 是 | - | 账本ID | 正整数 |
| email | string | 是 | - | 被邀请人邮箱 | 有效的邮箱格式 |
| permission | string | 是 | - | 权限级别 | 枚举值：viewer, editor, manager, owner |
| message | string | 否 | "" | 邀请消息 | 0-200个字符 |

**请求**

```http
POST /api/v1/shared/books/:bookId/members
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "email": "family_member@example.com",
  "permission": "editor",
  "message": "邀请您加入我们的家庭账本"
}
```

**成功响应**

```json
{
  "code": 200,
  "message": "邀请已发送",
  "data": {
    "invitation_id": "inv_123",
    "book_id": 1,
    "recipient_email": "family_member@example.com",
    "permission": "editor",
    "status": "pending",
    "expires_at": "2023-01-12T12:30:00Z",
    "created_at": "2023-01-05T12:30:00Z"
  }
}
```

**错误响应**

```json
{
  "code": 403,
  "message": "权限不足：您没有邀请成员的权限",
  "data": null
}
```

###  创建提醒

**功能描述**：创建新的账单提醒

**请求参数**

| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| book_id | integer | 是 | - | 账本ID | 正整数 |
| title | string | 是 | - | 提醒标题 | 1-50个字符 |
| amount | number | 是 | - | 提醒金额 | 大于0 |
| currency | string | 否 | "CNY" | 货币类型 | 有效的货币代码 |
| type | string | 是 | - | 提醒类型 | 枚举值：income, expense, fixed_income, fixed_expense |
| recurrence | string | 是 | - | 重复频率 | 枚举值：one_time, weekly, biweekly, monthly, yearly |
| recurrence_day | integer | 否 | - | 重复日 | 1-31之间的整数，仅当recurrence为monthly时必填 |
| start_date | string | 是 | - | 开始日期 | YYYY-MM-DD格式 |
| end_date | string | 否 | - | 结束日期 | YYYY-MM-DD格式，必须大于等于start_date |
| reminder_days | array | 否 | [1] | 提前提醒天数 | 数组元素为0-30之间的整数 |
| account_id | integer | 是 | - | 账户ID | 正整数 |
| category_id | integer | 是 | - | 分类ID | 正整数 |
| auto_create | boolean | 否 | false | 是否自动创建交易记录 | 布尔值 |
| is_active | boolean | 否 | true | 是否激活 | 布尔值 |

**请求**

```http
POST /api/v1/reminders
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
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
  "is_active": true
}
```

**成功响应**

```json
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

**错误响应**

```json
{
  "code": 400,
  "message": "请求参数错误：结束日期必须大于等于开始日期",
  "data": null
}
```

###  预测交易分类

**功能描述**：根据交易信息预测分类

**请求参数**

| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| book_id | integer | 是 | - | 账本ID | 正整数 |
| transaction | object | 是 | - | 交易信息 | 包含amount, description, account_id, date, currency字段 |
| include_probability | boolean | 否 | false | 是否包含概率 | 布尔值 |

**请求**

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

**成功响应**

```json
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
    "confidence_level": "high",
    "model_version": "1.0.0"
  }
}
```

**错误响应**

```json
{
  "code": 2005,
  "message": "智能分类模型错误：模型加载失败",
  "data": null
}
```

###  货币转换

**功能描述**：进行货币转换

**请求参数**

| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| amount | number | 是 | - | 转换金额 | 大于0 |
| from_currency | string | 是 | - | 源货币 | 有效的货币代码 |
| to_currency | string | 是 | - | 目标货币 | 有效的货币代码 |
| date | string | 否 | 当前日期 | 汇率日期 | YYYY-MM-DD格式 |

**请求**

```http
POST /api/v1/currency/convert
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "amount": 100,
  "from_currency": "USD",
  "to_currency": "CNY",
  "date": "2023-01-05"
}
```

**成功响应**

```json
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

**错误响应**

```json
{
  "code": 2010,
  "message": "汇率数据过期：请刷新汇率数据",
  "data": null
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

##  错误码说明

| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误 | 检查请求参数是否符合要求，包括类型、格式、必填项等 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有执行该操作的权限 |
| 404 | 资源不存在 | 检查请求的资源ID是否正确 |
| 500 | 服务器内部错误 | 请稍后重试，或联系系统管理员 |
| 2000 | 预算创建失败 | 检查预算参数是否正确，特别是金额、日期范围等 |
| 2001 | 标签创建失败 | 检查标签名称是否已存在，或参数格式是否正确 |
| 2002 | 共享权限不足 | 只有管理员或账本所有者可以邀请成员 |
| 2003 | 提醒创建失败 | 检查提醒参数是否正确，特别是日期范围、重复规则等 |
| 2004 | 汇率获取失败 | 检查网络连接，或稍后重试 |
| 2005 | 智能分类模型错误 | 模型加载或预测失败，请稍后重试 |
| 2006 | 预算已过期 | 预算周期已结束，无法进行操作 |
| 2007 | 超出预算限额 | 已超出预算金额，请调整预算或消费 |
| 2008 | 邀请已过期 | 邀请链接已过期，请重新发送邀请 |
| 2009 | 提醒频率过高 | 提醒频率设置过高，请调整提醒规则 |
| 2010 | 汇率数据过期 | 汇率数据已过期，请调用刷新汇率接口更新数据 |

##  API版本控制策略

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

##  API使用规范

1. **请求格式**：所有请求必须使用JSON格式
2. **认证方式**：使用JWT令牌进行认证，令牌放在Authorization头中
3. **请求频率限制**：每个API有请求频率限制，默认每分钟60次
4. **错误处理**：客户端应根据错误码进行相应处理
5. **分页规则**：列表接口支持分页，使用`page`和`page_size`参数
6. **排序规则**：列表接口支持排序，使用`sort_by`和`sort_order`参数