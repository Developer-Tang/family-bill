## 模块概述

收支记录模块是系统的核心功能，负责记录用户的日常收支活动，支持多种记账方式和记录类型。本模块提供了完整的收支记录管理、快速记账、周期记账、批量记账等功能的API接口。

### 设计原则

1. **RESTful规范**：所有API接口严格遵循RESTful设计规范，使用标准HTTP方法和资源导向的URL路径
2. **本地缓存优先**：暂不集成远程缓存机制，如需使用缓存功能，严格限定为本地缓存方案
3. **完整的请求验证**：所有API接口包含严格的请求参数验证，确保数据完整性和安全性
4. **统一的错误处理**：采用统一的错误响应格式和错误码体系，便于客户端处理
5. **清晰的响应格式**：所有API响应采用标准JSON格式，包含明确的数据结构和状态信息
6. **可扩展性**：API设计考虑未来功能扩展，预留合理的扩展点

### 适用场景

- 用户需要记录日常收支活动
- 用户需要快速记账功能，提高记账效率
- 用户需要设置周期记账，自动记录固定支出
- 用户需要批量导入外部账单
- 用户需要管理收支记录的标签
- 用户需要在离线状态下记账，后续同步

## 接口清单

<!-- tabs:start -->
<!-- tab:基础记账 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/transactions`](#获取收支记录列表) | `GET` | 获取收支记录列表 |
| [`/api/v1/transactions`](#创建收支记录) | `POST` | 创建收支记录 |
| [`/api/v1/transactions/:id`](#获取收支记录详情) | `GET` | 获取收支记录详情 |
| [`/api/v1/transactions/:id`](#更新收支记录) | `PUT` | 更新收支记录 |
| [`/api/v1/transactions/:id`](#删除收支记录) | `DELETE` | 删除收支记录 |
| [`/api/v1/transactions/:id/lock`](#锁定解锁收支记录) | `PUT` | 锁定/解锁收支记录 |

<!-- tab:快速记账 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/quick-transactions`](#创建快速记账记录) | `POST` | 创建快速记账记录 |
| [`/api/v1/transaction-templates`](#获取记账模板) | `GET` | 获取记账模板 |
| [`/api/v1/transaction-templates`](#创建记账模板) | `POST` | 创建记账模板 |
| [`/api/v1/transaction-templates/:id`](#更新记账模板) | `PUT` | 更新记账模板 |
| [`/api/v1/transaction-templates/:id`](#删除记账模板) | `DELETE` | 删除记账模板 |

<!-- tab:周期记账 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/recurring-transactions`](#获取周期记账列表) | `GET` | 获取周期记账列表 |
| [`/api/v1/recurring-transactions`](#创建周期记账) | `POST` | 创建周期记账 |
| [`/api/v1/recurring-transactions/:id`](#更新周期记账) | `PUT` | 更新周期记账 |
| [`/api/v1/recurring-transactions/:id`](#删除周期记账) | `DELETE` | 删除周期记账 |
| [`/api/v1/recurring-transactions/:id/trigger`](#手动触发周期记账) | `POST` | 手动触发周期记账 |

<!-- tab:其他记账 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/calendar/transactions`](#获取日历视图收支记录) | `GET` | 获取日历视图收支记录 |
| [`/api/v1/transactions/batch`](#批量创建收支记录) | `POST` | 批量创建收支记录 |

<!-- tab:标签管理 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/tags`](#获取标签列表) | `GET` | 获取标签列表 |
| [`/api/v1/tags`](#创建标签) | `POST` | 创建标签 |
| [`/api/v1/tags/:id`](#更新标签) | `PUT` | 更新标签 |
| [`/api/v1/tags/:id`](#删除标签) | `DELETE` | 删除标签 |

<!-- tab:其他 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/transactions/import`](#导入外部账单) | `POST` | 导入外部账单 |
| [`/api/v1/transactions/export`](#导出收支记录) | `GET` | 导出收支记录 |
| [`/api/v1/offline/sync`](#同步离线记账数据) | `POST` | 同步离线记账数据 |
| [`/api/v1/ocr/receipt`](#账单图片识别) | `POST` | 账单图片识别 |

<!-- tabs:end -->

## 详细接口说明

### 创建收支记录

**功能描述**：创建新的收支记录，支持支出和收入两种类型

**请求方法**：POST
**URL路径**：/api/v1/transactions
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |
| Authorization | string | 是 | Bearer JWT令牌 |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| book_id | integer | 是 | - | 账本ID | 正整数，账本必须存在且用户有访问权限 |
| type | string | 是 | - | 交易类型 | 枚举值：expense（支出）, income（收入） |
| amount | number | 是 | - | 交易金额 | 大于0，最多两位小数 |
| currency | string | 否 | "CNY" | 货币类型 | 有效的货币代码 |
| account_id | integer | 是 | - | 账户ID | 正整数，账户必须存在且属于该账本 |
| category_id | integer | 是 | - | 分类ID | 正整数，分类必须存在且属于该账本 |
| subcategory_id | integer | 否 | null | 子分类ID | 正整数或null，子分类必须存在且属于该分类 |
| date | string | 是 | - | 交易日期 | YYYY-MM-DD格式 |
| time | string | 否 | "00:00" | 交易时间 | HH:MM格式 |
| memo | string | 否 | "" | 交易备注 | 0-200个字符 |
| tags | array | 否 | [] | 标签ID列表 | 数组元素为正整数，标签必须存在且属于该账本 |
| member_id | integer | 否 | - | 记账成员ID | 正整数，成员必须存在且属于该账本 |
| attachments | array | 否 | [] | 附件列表 | 数组元素为base64编码的图片数据 |
| location | object | 否 | null | 地理位置信息 | 包含latitude, longitude, address字段 |
| location.latitude | number | 否 | - | 纬度 | -90到90之间的数值 |
| location.longitude | number | 否 | - | 经度 | -180到180之间的数值 |
| location.address | string | 否 | - | 详细地址 | 0-200个字符 |

**请求示例**
```http
POST /api/v1/transactions
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "type": "expense",
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

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| transaction_id | integer | 交易记录ID |
| type | string | 交易类型 |
| amount | number | 交易金额 |
| currency | string | 货币类型 |
| account_name | string | 账户名称 |
| category_name | string | 分类名称 |
| date | string | 交易日期 |
| time | string | 交易时间 |
| memo | string | 交易备注 |
| created_at | string | 创建时间 |

**成功响应**
```json
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

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误 | 检查请求参数是否符合要求，包括类型、格式、必填项等 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有创建收支记录的权限 |
| 404 | 账本不存在 | 检查book_id是否正确 |
| 404 | 账户不存在 | 检查account_id是否正确 |
| 404 | 分类不存在 | 检查category_id是否正确 |
| 800 | 账户余额不足 | 检查账户余额是否足够 |
| 801 | 无效的交易类型 | 检查type参数是否为有效值 |
| 802 | 无效的日期格式 | 检查date参数格式是否为YYYY-MM-DD |

**本地缓存策略**：创建成功后，本地缓存该收支记录，缓存时间24小时

### 获取收支记录列表

**功能描述**：获取指定条件的收支记录列表，支持分页和多条件筛选

**请求方法**：GET
**URL路径**：/api/v1/transactions
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| book_id | integer | 是 | - | 账本ID |
| type | string | 否 | - | 交易类型：expense（支出）, income（收入） |
| start_date | string | 否 | - | 开始日期，格式：YYYY-MM-DD |
| end_date | string | 否 | - | 结束日期，格式：YYYY-MM-DD |
| category_id | integer | 否 | - | 分类ID |
| page | integer | 否 | 1 | 页码 |
| limit | integer | 否 | 20 | 每页数量 |

**请求示例**
```http
GET /api/v1/transactions?book_id=1&type=expense&start_date=2023-01-01&end_date=2023-01-31&category_id=5&page=1&limit=20
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 收支记录列表数据 |
| data.total | integer | 总记录数 |
| data.items | array | 收支记录列表 |
| data.items[].transaction_id | integer | 交易ID |
| data.items[].type | string | 交易类型 |
| data.items[].amount | number | 交易金额 |
| data.items[].currency | string | 货币类型 |
| data.items[].account_id | integer | 账户ID |
| data.items[].account_name | string | 账户名称 |
| data.items[].category_id | integer | 分类ID |
| data.items[].category_name | string | 分类名称 |
| data.items[].date | string | 交易日期 |
| data.items[].memo | string | 交易备注 |
| data.items[].tags | array | 标签列表 |
| data.items[].is_locked | boolean | 是否锁定 |
| data.items[].created_at | string | 创建时间 |

**成功响应示例**
```json
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

### 创建周期记账

**功能描述**：创建新的周期记账记录，用于自动生成定期收支记录

**请求方法**：POST
**URL路径**：/api/v1/recurring-transactions
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |
| Authorization | string | 是 | Bearer JWT令牌 |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| book_id | integer | 是 | - | 账本ID | 正整数，账本必须存在且用户有访问权限 |
| type | string | 是 | - | 交易类型 | 枚举值：expense（支出）, income（收入） |
| amount | number | 是 | - | 交易金额 | 大于0，最多两位小数 |
| currency | string | 否 | "CNY" | 货币类型 | 有效的货币代码 |
| account_id | integer | 是 | - | 账户ID | 正整数，账户必须存在且属于该账本 |
| category_id | integer | 是 | - | 分类ID | 正整数，分类必须存在且属于该账本 |
| name | string | 是 | - | 周期记账名称 | 1-50个字符 |
| frequency | string | 是 | - | 频率 | 枚举值：daily, weekly, monthly, yearly |
| interval | integer | 否 | 1 | 间隔 | 正整数 |
| start_date | string | 是 | - | 开始日期 | YYYY-MM-DD格式 |
| end_date | string | 否 | null | 结束日期 | YYYY-MM-DD格式，永久周期为null |
| next_date | string | 是 | - | 下次执行日期 | YYYY-MM-DD格式 |
| memo | string | 否 | "" | 交易备注 | 0-200个字符 |
| tags | array | 否 | [] | 标签ID列表 | 数组元素为正整数 |
| auto_create | boolean | 否 | true | 是否自动创建记录 | 布尔值 |
| notify_before_days | integer | 否 | 0 | 提前通知天数 | 非负整数 |

**请求示例**
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

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 周期记账创建结果 |
| data.recurring_id | integer | 周期记账ID |
| data.name | string | 周期记账名称 |
| data.type | string | 交易类型 |
| data.amount | number | 交易金额 |
| data.frequency | string | 频率 |
| data.interval | integer | 间隔 |
| data.next_date | string | 下次执行日期 |
| data.status | string | 状态 |

**成功响应示例**
```json
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

### 批量创建收支记录

**功能描述**：批量创建多条收支记录，提高记账效率

**请求方法**：POST
**URL路径**：/api/v1/transactions/batch
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |
| Authorization | string | 是 | Bearer JWT令牌 |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| book_id | integer | 是 | - | 账本ID | 正整数，账本必须存在且用户有访问权限 |
| transactions | array | 是 | - | 收支记录列表 | 数组元素为收支记录对象，至少包含type, amount, account_id, category_id, date字段 |
| transactions[].type | string | 是 | - | 交易类型 | 枚举值：expense（支出）, income（收入） |
| transactions[].amount | number | 是 | - | 交易金额 | 大于0，最多两位小数 |
| transactions[].account_id | integer | 是 | - | 账户ID | 正整数，账户必须存在且属于该账本 |
| transactions[].category_id | integer | 是 | - | 分类ID | 正整数，分类必须存在且属于该账本 |
| transactions[].date | string | 是 | - | 交易日期 | YYYY-MM-DD格式 |
| transactions[].time | string | 否 | "00:00" | 交易时间 | HH:MM格式 |
| transactions[].memo | string | 否 | "" | 交易备注 | 0-200个字符 |
| transactions[].tags | array | 否 | [] | 标签ID列表 | 数组元素为正整数 |

**请求示例**
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

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 批量创建结果 |
| data.success_count | integer | 成功创建的记录数 |
| data.failed_count | integer | 创建失败的记录数 |
| data.transaction_ids | array | 成功创建的交易ID列表 |

**成功响应示例**
```json
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

### 导入外部账单

**功能描述**：从外部平台导入账单数据，支持支付宝、微信、银行等平台的账单格式

**请求方法**：POST
**URL路径**：/api/v1/transactions/import
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为multipart/form-data |
| Authorization | string | 是 | Bearer JWT令牌 |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| book_id | integer | 是 | - | 账本ID | 正整数，账本必须存在且用户有访问权限 |
| provider | string | 是 | - | 账单提供方 | 枚举值：alipay, wechat, bank |
| file | file | 是 | - | 账单文件 | 支持的文件格式：csv, xlsx, pdf |

**请求示例**
```http
POST /api/v1/transactions/import
Content-Type: multipart/form-data
Authorization: Bearer jwt_token_string

book_id=1
provider=alipay  // alipay, wechat, bank
file=...  // 上传的账单文件
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 导入结果 |
| data.import_id | string | 导入任务ID |
| data.total_count | integer | 总记录数 |
| data.success_count | integer | 成功导入的记录数 |
| data.failed_count | integer | 导入失败的记录数 |
| data.failed_records | array | 失败记录详情 |
| data.failed_records[].index | integer | 失败记录在文件中的索引 |
| data.failed_records[].error | string | 失败原因 |
| data.transaction_ids | array | 成功创建的交易ID列表 |

**成功响应示例**
```json
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

## 数据模型

### 收支记录模型

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

### 周期记账模型

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

### 记账模板模型

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

## 标签体系

### 标签模型

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

## 错误码说明

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