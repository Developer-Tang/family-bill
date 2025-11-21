# 收支记录模块 API 文档

##  模块概述

收支记录模块是系统的核心功能，负责记录用户的日常收支活动，支持多种记账方式和记录类型。本模块提供了完整的收支记录管理、快速记账、周期记账、批量记账等功能的API接口。

##  接口清单

| 功能模块 | 接口路径 | 方法 | 功能描述 |
|---------|---------|------|--------|
| **基础记账** | `/api/v1/transactions` | `GET` | 获取收支记录列表 |
| | `/api/v1/transactions` | `POST` | 创建收支记录 |
| | `/api/v1/transactions/:id` | `GET` | 获取收支记录详情 |
| | `/api/v1/transactions/:id` | `PUT` | 更新收支记录 |
| | `/api/v1/transactions/:id` | `DELETE` | 删除收支记录 |
| | `/api/v1/transactions/:id/lock` | `PUT` | 锁定/解锁收支记录 |
| **快速记账** | `/api/v1/quick-transactions` | `POST` | 创建快速记账记录 |
| | `/api/v1/transaction-templates` | `GET` | 获取记账模板 |
| | `/api/v1/transaction-templates` | `POST` | 创建记账模板 |
| | `/api/v1/transaction-templates/:id` | `PUT` | 更新记账模板 |
| | `/api/v1/transaction-templates/:id` | `DELETE` | 删除记账模板 |
| **周期记账** | `/api/v1/recurring-transactions` | `GET` | 获取周期记账列表 |
| | `/api/v1/recurring-transactions` | `POST` | 创建周期记账 |
| | `/api/v1/recurring-transactions/:id` | `PUT` | 更新周期记账 |
| | `/api/v1/recurring-transactions/:id` | `DELETE` | 删除周期记账 |
| | `/api/v1/recurring-transactions/:id/trigger` | `POST` | 手动触发周期记账 |
| **日历记账** | `/api/v1/calendar/transactions` | `GET` | 获取日历视图收支记录 |
| **批量记账** | `/api/v1/transactions/batch` | `POST` | 批量创建收支记录 |
| | `/api/v1/transactions/import` | `POST` | 导入外部账单 |
| **标签管理** | `/api/v1/tags` | `GET` | 获取标签列表 |
| | `/api/v1/tags` | `POST` | 创建标签 |
| | `/api/v1/tags/:id` | `PUT` | 更新标签 |
| | `/api/v1/tags/:id` | `DELETE` | 删除标签 |
| **离线记账** | `/api/v1/offline/sync` | `POST` | 同步离线记账数据 |
| **图片识别** | `/api/v1/ocr/receipt` | `POST` | 账单图片识别 |

##  详细接口说明

###  创建收支记录

#### 请求

```http
POST /api/v1/transactions
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "type": "expense",  // expense 支出, income 收入
  "amount": 100.50,
  "currency": "CNY",
  "account_id": 1,
  "category_id": 5,
  "subcategory_id": null,
  "date": "2023-01-05",
  "time": "12:30",
  "memo": "午餐",
  "tags": [1, 2],
  "member_id": 1,
  "attachments": ["base64_encoded_image_1", "base64_encoded_image_2"],
  "location": {
    "latitude": 39.9042,
    "longitude": 116.4074,
    "address": "北京市朝阳区"
  }
}
```

#### 响应

```
# 成功
{
  "code": 200,
  "message": "记账成功",
  "data": {
    "transaction_id": 1,
    "type": "expense",
    "amount": 100.50,
    "currency": "CNY",
    "account_name": "我的工资卡",
    "category_name": "餐饮",
    "date": "2023-01-05",
    "time": "12:30",
    "memo": "午餐",
    "created_at": "2023-01-05T12:30:00Z"
  }
}
```

###  获取收支记录列表

#### 请求

```http
GET /api/v1/transactions?book_id=1&type=expense&start_date=2023-01-01&end_date=2023-01-31&category_id=5&page=1&limit=20
Authorization: Bearer jwt_token_string
```

#### 响应

```
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total": 150,
    "items": [
      {
        "transaction_id": 1,
        "type": "expense",
        "amount": 100.50,
        "currency": "CNY",
        "account_id": 1,
        "account_name": "我的工资卡",
        "category_id": 5,
        "category_name": "餐饮",
        "date": "2023-01-05",
        "memo": "午餐",
        "tags": ["日常", "午餐"],
        "is_locked": false,
        "created_at": "2023-01-05T12:30:00Z"
      },
      {
        "transaction_id": 2,
        "type": "expense",
        "amount": 50.00,
        "currency": "CNY",
        "account_id": 2,
        "account_name": "支付宝",
        "category_id": 6,
        "category_name": "交通",
        "date": "2023-01-05",
        "memo": "打车",
        "tags": ["交通"],
        "is_locked": false,
        "created_at": "2023-01-05T15:45:00Z"
      }
    ]
  }
}
```

###  创建周期记账

#### 请求

```http
POST /api/v1/recurring-transactions
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "type": "expense",
  "amount": 3000.00,
  "currency": "CNY",
  "account_id": 1,
  "category_id": 10,
  "name": "房租",
  "frequency": "monthly",  // daily, weekly, monthly, yearly
  "interval": 1,
  "start_date": "2023-01-01",
  "end_date": "2023-12-31",
  "next_date": "2023-02-01",
  "memo": "每月房租",
  "tags": [3],
  "auto_create": true,
  "notify_before_days": 3
}
```

#### 响应

```
# 成功
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "recurring_id": 1,
    "name": "房租",
    "type": "expense",
    "amount": 3000.00,
    "frequency": "monthly",
    "interval": 1,
    "next_date": "2023-02-01",
    "status": "active"
  }
}
```

###  批量创建收支记录

#### 请求

```http
POST /api/v1/transactions/batch
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "transactions": [
    {
      "type": "expense",
      "amount": 15.00,
      "account_id": 2,
      "category_id": 5,
      "date": "2023-01-05",
      "memo": "早餐"
    },
    {
      "type": "expense",
      "amount": 25.00,
      "account_id": 2,
      "category_id": 5,
      "date": "2023-01-05",
      "memo": "晚餐"
    },
    {
      "type": "income",
      "amount": 5000.00,
      "account_id": 1,
      "category_id": 20,
      "date": "2023-01-10",
      "memo": "工资"
    }
  ]
}
```

#### 响应

```
{
  "code": 200,
  "message": "批量记账成功",
  "data": {
    "success_count": 3,
    "failed_count": 0,
    "transaction_ids": [1, 2, 3]
  }
}
```

###  导入外部账单

#### 请求

```http
POST /api/v1/transactions/import
Content-Type: multipart/form-data
Authorization: Bearer jwt_token_string

book_id=1
provider=alipay  // alipay, wechat, bank
file=...  // 上传的账单文件
```

#### 响应

```
{
  "code": 200,
  "message": "导入成功",
  "data": {
    "import_id": "uuid-string",
    "total_count": 50,
    "success_count": 48,
    "failed_count": 2,
    "failed_records": [
      {
        "index": 10,
        "error": "无效的金额格式"
      },
      {
        "index": 45,
        "error": "无法识别的账户"
      }
    ],
    "transaction_ids": [4, 5, ..., 51]
  }
}
```

##  数据模型

###  收支记录模型

```javascript
{
  "transaction_id": 1,    // 记录ID
  "book_id": 1,          // 所属账本ID
  "type": "expense",    // 类型：expense 支出, income 收入
  "amount": 100.50,      // 金额
  "currency": "CNY",    // 货币类型
  "account_id": 1,       // 账户ID
  "account_name": "我的工资卡", // 账户名称
  "category_id": 5,      // 分类ID
  "category_name": "餐饮", // 分类名称
  "subcategory_id": null, // 子分类ID
  "date": "2023-01-05", // 日期
  "time": "12:30",      // 时间
  "memo": "午餐",        // 备注
  "tags": [1, 2],        // 标签ID列表
  "tag_names": ["日常", "午餐"], // 标签名称列表
  "member_id": 1,        // 记账成员ID
  "member_name": "张三", // 记账成员名称
  "attachments": ["url1", "url2"], // 附件URL列表
  "location": {
    "latitude": 39.9042,
    "longitude": 116.4074,
    "address": "北京市朝阳区"
  },
  "is_locked": false,    // 是否锁定
  "created_at": "2023-01-05T12:30:00Z", // 创建时间
  "updated_at": "2023-01-05T12:30:00Z"  // 更新时间
}
```

###  周期记账模型

```javascript
{
  "recurring_id": 1,     // 周期记账ID
  "book_id": 1,          // 所属账本ID
  "name": "房租",        // 名称
  "type": "expense",    // 类型
  "amount": 3000.00,     // 金额
  "currency": "CNY",    // 货币类型
  "account_id": 1,       // 账户ID
  "category_id": 10,     // 分类ID
  "frequency": "monthly", // 频率：daily, weekly, monthly, yearly
  "interval": 1,         // 间隔
  "start_date": "2023-01-01", // 开始日期
  "end_date": "2023-12-31", // 结束日期
  "next_date": "2023-02-01", // 下次执行日期
  "memo": "每月房租",    // 备注
  "tags": [3],           // 标签ID列表
  "auto_create": true,   // 是否自动创建
  "notify_before_days": 3, // 提前通知天数
  "status": "active",   // 状态：active, paused, completed
  "created_at": "2023-01-01T12:00:00Z" // 创建时间
}
```

###  记账模板模型

```javascript
{
  "template_id": 1,      // 模板ID
  "book_id": 1,          // 所属账本ID
  "name": "午餐模板",     // 模板名称
  "type": "expense",    // 类型
  "amount": 30.00,       // 金额
  "account_id": 2,       // 账户ID
  "category_id": 5,      // 分类ID
  "tags": [1, 2],        // 标签ID列表
  "memo": "午餐",        // 备注
  "created_at": "2023-01-01T12:00:00Z" // 创建时间
}
```

##  标签体系

###  标签模型

```javascript
{
  "tag_id": 1,           // 标签ID
  "book_id": 1,          // 所属账本ID
  "name": "日常",        // 标签名称
  "color": "#FF6B6B",    // 标签颜色
  "count": 50,           // 使用次数
  "created_at": "2023-01-01T12:00:00Z" // 创建时间
}
```

###  标签使用说明

- 每个收支记录可以关联多个标签
- 标签按使用频率排序
- 支持标签的批量管理

##  错误码说明

| 错误码 | 描述 |
|-------|------|
| 400 | 请求参数错误 |
| 401 | 未授权，需要登录 |
| 403 | 权限不足 |
| 404 | 收支记录不存在 |
| 409 | 无法修改已锁定的记录 |
| 500 | 服务器内部错误 |
| 800 | 账户余额不足 |
| 801 | 无效的交易类型 |
| 802 | 无效的日期格式 |
| 803 | 文件格式不支持 |
| 804 | 导入文件解析失败 |
| 805 | 周期记账频率无效 |