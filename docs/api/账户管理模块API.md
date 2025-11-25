##  模块概述

账户管理模块负责管理用户的各类资产账户，包括现金、银行卡、支付宝、微信等，支持余额管理、转账和多币种。本模块提供了完整的账户生命周期管理和资金转移功能的API接口。

##  接口清单

<!-- tabs:start -->
<!-- tab:账户类型管理 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/account-types`](#获取账户类型列表) | `GET` | 获取账户类型列表 |
| [`/api/v1/account-types`](#创建自定义账户类型) | `POST` | 创建自定义账户类型 |
| [`/api/v1/account-types/:id`](#更新账户类型) | `PUT` | 更新账户类型 |
| [`/api/v1/account-types/:id`](#删除账户类型) | `DELETE` | 删除账户类型 |

<!-- tab:账户管理 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/accounts`](#获取账户列表) | `GET` | 获取账户列表 |
| [`/api/v1/accounts`](#创建账户) | `POST` | 创建账户 |
| [`/api/v1/accounts/:id`](#获取账户详情) | `GET` | 获取账户详情 |
| [`/api/v1/accounts/:id`](#更新账户信息) | `PUT` | 更新账户信息 |
| [`/api/v1/accounts/:id`](#删除账户) | `DELETE` | 删除账户 |
| [`/api/v1/accounts/:id/balance`](#调整账户余额) | `PUT` | 调整账户余额 |
| [`/api/v1/accounts/:id/status`](#修改账户状态) | `PUT` | 修改账户状态 |

<!-- tab:账户分组 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/account-groups`](#获取账户分组列表) | `GET` | 获取账户分组列表 |
| [`/api/v1/account-groups`](#创建账户分组) | `POST` | 创建账户分组 |
| [`/api/v1/account-groups/:id`](#更新账户分组) | `PUT` | 更新账户分组 |
| [`/api/v1/account-groups/:id`](#删除账户分组) | `DELETE` | 删除账户分组 |
| [`/api/v1/account-groups/:id/accounts`](#账户加入分组) | `POST` | 账户加入分组 |

<!-- tab:账户转账 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/transfers`](#创建转账记录) | `POST` | 创建转账记录 |
| [`/api/v1/transfers`](#获取转账记录列表) | `GET` | 获取转账记录列表 |
| [`/api/v1/transfers/:id`](#获取转账记录详情) | `GET` | 获取转账记录详情 |

<!-- tab:多币种支持 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/currencies`](#获取货币列表) | `GET` | 获取货币列表 |
| [`/api/v1/exchange-rates`](#获取汇率) | `GET` | 获取汇率 |
| [`/api/v1/exchange-rates`](#更新汇率) | `POST` | 更新汇率 |

<!-- tabs:end -->

##  详细接口说明

###  获取账户类型列表

**请求**

```http
GET /api/v1/account-types
Authorization: Bearer jwt_token_string
```

**响应**

```
{
  "code": 200,
  "message": "获取成功",
  "data": [
    {
      "type_id": 1,
      "name": "银行卡",
      "icon": "bank_icon",
      "color": "#0080FF",
      "is_system": true
    },
    {
      "type_id": 2,
      "name": "第三方支付",
      "icon": "payment_icon",
      "color": "#4CAF50",
      "is_system": true
    }
  ]
}
```

###  创建自定义账户类型

**请求**

```http
POST /api/v1/account-types
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "name": "投资账户",
  "icon": "investment_icon",
  "color": "#FF9800"
}
```

**响应**

```
# 成功
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "type_id": 10,
    "name": "投资账户",
    "icon": "investment_icon",
    "color": "#FF9800",
    "is_system": false
  }
}
```

###  更新账户类型

**请求**

```http
PUT /api/v1/account-types/:id
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "name": "投资账户(更新)",
  "icon": "new_investment_icon",
  "color": "#FF5722"
}
```

**响应**

```
# 成功
{
  "code": 200,
  "message": "更新成功",
  "data": {
    "type_id": 10,
    "name": "投资账户(更新)",
    "icon": "new_investment_icon",
    "color": "#FF5722",
    "is_system": false
  }
}
```

###  删除账户类型

**请求**

```http
DELETE /api/v1/account-types/:id
Authorization: Bearer jwt_token_string
```

**响应**

```
# 成功
{
  "code": 200,
  "message": "删除成功",
  "data": null
}
```

###  获取账户列表

**请求**

```http
GET /api/v1/accounts?book_id=1&status=active&page=1&limit=20
Authorization: Bearer jwt_token_string
```

**响应**

```
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total": 5,
    "items": [
      {
        "account_id": 1,
        "name": "我的工资卡",
        "account_type": "银行卡",
        "balance": 5000.00,
        "currency": "CNY",
        "status": "active",
        "created_at": "2023-01-01T12:00:00Z"
      },
      {
        "account_id": 2,
        "name": "支付宝",
        "account_type": "第三方支付",
        "balance": 2000.00,
        "currency": "CNY",
        "status": "active",
        "created_at": "2023-01-01T12:00:00Z"
      }
    ]
  }
}
```

###  创建账户

**请求**

```http
POST /api/v1/accounts
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "name": "我的工资卡",
  "account_type_id": 1,
  "currency": "CNY",
  "initial_balance": 5000.00,
  "account_group_id": null,
  "hidden_balance": false,
  "memo": "工商银行储蓄卡",
  "icon": "bank_icon",
  "color": "#0080FF"
}
```

**响应**

```
# 成功
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "account_id": 1,
    "name": "我的工资卡",
    "balance": 5000.00,
    "currency": "CNY",
    "status": "active",
    "created_at": "2023-01-01T12:00:00Z"
  }
}

# 失败
{
  "code": 400,
  "message": "创建失败：账户名称已存在",
  "data": null
}
```

###  获取账户详情

**请求**

```http
GET /api/v1/accounts/:id
Authorization: Bearer jwt_token_string
```

**响应**

```
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "account_id": 1,
    "name": "我的工资卡",
    "account_type_id": 1,
    "account_type": "银行卡",
    "balance": 5000.00,
    "currency": "CNY",
    "status": "active",
    "hidden_balance": false,
    "memo": "工商银行储蓄卡",
    "icon": "bank_icon",
    "color": "#0080FF",
    "created_at": "2023-01-01T12:00:00Z",
    "updated_at": "2023-01-01T12:00:00Z"
  }
}
```

###  更新账户信息

**请求**

```http
PUT /api/v1/accounts/:id
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "name": "我的新工资卡",
  "memo": "建设银行储蓄卡",
  "icon": "new_bank_icon",
  "color": "#2196F3",
  "hidden_balance": true
}
```

**响应**

```
# 成功
{
  "code": 200,
  "message": "更新成功",
  "data": {
    "account_id": 1,
    "name": "我的新工资卡",
    "memo": "建设银行储蓄卡",
    "icon": "new_bank_icon",
    "color": "#2196F3",
    "hidden_balance": true,
    "updated_at": "2023-01-01T12:00:00Z"
  }
}
```

###  删除账户

**请求**

```http
DELETE /api/v1/accounts/:id
Authorization: Bearer jwt_token_string
```

**响应**

```
# 成功
{
  "code": 200,
  "message": "删除成功",
  "data": null
}
```

###  调整账户余额

**请求**

```http
PUT /api/v1/accounts/:id/balance
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "delta": 1000.00,  // 正数表示增加，负数表示减少
  "reason": "调整初始余额",
  "date": "2023-01-01"
}
```

**响应**

```
# 成功
{
  "code": 200,
  "message": "余额调整成功",
  "data": {
    "account_id": 1,
    "old_balance": 5000.00,
    "new_balance": 6000.00,
    "delta": 1000.00,
    "adjusted_at": "2023-01-01T12:00:00Z"
  }
}
```

###  获取账户分组列表

**请求**

```http
GET /api/v1/account-groups?book_id=1
Authorization: Bearer jwt_token_string
```

**响应**

```
{
  "code": 200,
  "message": "获取成功",
  "data": [
    {
      "group_id": 1,
      "name": "日常消费",
      "icon": "shopping_icon",
      "color": "#FF6B6B",
      "accounts_count": 3,
      "total_balance": 8000.00
    },
    {
      "group_id": 2,
      "name": "投资账户",
      "icon": "investment_icon",
      "color": "#4CAF50",
      "accounts_count": 2,
      "total_balance": 20000.00
    }
  ]
}
```

###  创建账户分组

**请求**

```http
POST /api/v1/account-groups
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "name": "旅行基金",
  "icon": "travel_icon",
  "color": "#2196F3"
}
```

**响应**

```
# 成功
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "group_id": 3,
    "name": "旅行基金",
    "icon": "travel_icon",
    "color": "#2196F3",
    "accounts_count": 0,
    "total_balance": 0.00
  }
}
```

###  更新账户分组

**请求**

```http
PUT /api/v1/account-groups/:id
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "name": "梦想旅行基金",
  "icon": "dream_travel_icon",
  "color": "#9C27B0"
}
```

**响应**

```
# 成功
{
  "code": 200,
  "message": "更新成功",
  "data": {
    "group_id": 3,
    "name": "梦想旅行基金",
    "icon": "dream_travel_icon",
    "color": "#9C27B0",
    "accounts_count": 0,
    "total_balance": 0.00
  }
}
```

###  删除账户分组

**请求**

```http
DELETE /api/v1/account-groups/:id
Authorization: Bearer jwt_token_string
```

**响应**

```
# 成功
{
  "code": 200,
  "message": "删除成功",
  "data": null
}
```

###  账户加入分组

**请求**

```http
POST /api/v1/account-groups/:id/accounts
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "account_ids": [1, 2, 3]
}
```

**响应**

```
# 成功
{
  "code": 200,
  "message": "操作成功",
  "data": {
    "group_id": 1,
    "accounts_count": 3,
    "total_balance": 15000.00
  }
}
```

###  创建转账记录

**请求**

```http
POST /api/v1/transfers
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "from_account_id": 1,
  "to_account_id": 2,
  "amount": 1000.00,
  "currency": "CNY",
  "transfer_date": "2023-01-05",
  "memo": "转账到支付宝",
  "exchange_rate": 1.0  // 如果涉及不同币种转账，需要指定汇率
}
```

**响应**

```
# 成功
{
  "code": 200,
  "message": "转账成功",
  "data": {
    "transfer_id": 1,
    "from_account": {
      "id": 1,
      "name": "我的工资卡",
      "balance": 5000.00
    },
    "to_account": {
      "id": 2,
      "name": "支付宝",
      "balance": 3000.00
    },
    "amount": 1000.00,
    "currency": "CNY",
    "transfer_date": "2023-01-05T00:00:00Z",
    "memo": "转账到支付宝"
  }
}

# 失败 - 余额不足
{
  "code": 400,
  "message": "转账失败：转出账户余额不足",
  "data": null
}
```

###  获取转账记录列表

**请求**

```http
GET /api/v1/transfers?book_id=1&page=1&limit=20&start_date=2023-01-01&end_date=2023-01-31
Authorization: Bearer jwt_token_string
```

**响应**

```
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total": 5,
    "items": [
      {
        "transfer_id": 1,
        "from_account_id": 1,
        "from_account_name": "我的工资卡",
        "to_account_id": 2,
        "to_account_name": "支付宝",
        "amount": 1000.00,
        "currency": "CNY",
        "transfer_date": "2023-01-05T00:00:00Z",
        "memo": "转账到支付宝"
      },
      {
        "transfer_id": 2,
        "from_account_id": 2,
        "from_account_name": "支付宝",
        "to_account_id": 3,
        "to_account_name": "微信钱包",
        "amount": 500.00,
        "currency": "CNY",
        "transfer_date": "2023-01-10T00:00:00Z",
        "memo": "转账到微信"
      }
    ]
  }
}
```

###  获取转账记录详情

**请求**

```http
GET /api/v1/transfers/:id
Authorization: Bearer jwt_token_string
```

**响应**

```
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "transfer_id": 1,
    "book_id": 1,
    "from_account_id": 1,
    "from_account_name": "我的工资卡",
    "from_account_type": "银行卡",
    "to_account_id": 2,
    "to_account_name": "支付宝",
    "to_account_type": "第三方支付",
    "amount": 1000.00,
    "currency": "CNY",
    "transfer_date": "2023-01-05T00:00:00Z",
    "memo": "转账到支付宝",
    "created_at": "2023-01-05T12:00:00Z"
  }
}
```

###  获取货币列表

**请求**

```http
GET /api/v1/currencies
Authorization: Bearer jwt_token_string
```

**响应**

```
{
  "code": 200,
  "message": "获取成功",
  "data": [
    {
      "code": "CNY",
      "name": "人民币",
      "symbol": "¥"
    },
    {
      "code": "USD",
      "name": "美元",
      "symbol": "$"
    },
    {
      "code": "EUR",
      "name": "欧元",
      "symbol": "€"
    },
    {
      "code": "GBP",
      "name": "英镑",
      "symbol": "£"
    },
    {
      "code": "JPY",
      "name": "日元",
      "symbol": "¥"
    }
  ]
}
```

###  获取汇率

**请求**

```http
GET /api/v1/exchange-rates?from=CNY&to=USD,EUR,GBP
Authorization: Bearer jwt_token_string
```

**响应**

```
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "base": "CNY",
    "rates": {
      "USD": 0.14,
      "EUR": 0.13,
      "GBP": 0.11
    },
    "updated_at": "2023-01-01T12:00:00Z"
  }
}
```

###  更新汇率

**请求**

```http
POST /api/v1/exchange-rates
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "base": "CNY",
  "rates": {
    "USD": 0.145,
    "EUR": 0.135,
    "GBP": 0.115
  }
}
```

**响应**

```
# 成功
{
  "code": 200,
  "message": "更新成功",
  "data": {
    "base": "CNY",
    "updated_at": "2023-01-01T12:00:00Z"
  }
}
```

###  获取账户列表

**请求**

```http
GET /api/v1/accounts?book_id=1&status=active&page=1&limit=20
Authorization: Bearer jwt_token_string
```

**响应**

```
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total": 5,
    "items": [
      {
        "account_id": 1,
        "name": "我的工资卡",
        "account_type": "银行卡",
        "balance": 5000.00,
        "currency": "CNY",
        "status": "active",
        "created_at": "2023-01-01T12:00:00Z"
      },
      {
        "account_id": 2,
        "name": "支付宝",
        "account_type": "第三方支付",
        "balance": 2000.00,
        "currency": "CNY",
        "status": "active",
        "created_at": "2023-01-01T12:00:00Z"
      }
    ]
  }
}
```

###  修改账户状态

**请求**

```http
PUT /api/v1/accounts/:id/status
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "status": "inactive",  // active, inactive, archived
  "reason": "账户不再使用"
}
```

**响应**

```
# 成功
{
  "code": 200,
  "message": "状态更新成功",
  "data": {
    "account_id": 1,
    "name": "旧银行卡",
    "status": "inactive",
    "updated_at": "2023-01-01T12:00:00Z"
  }
}
```

##  数据模型

###  账户模型

```javascript
{
  "account_id": 1,         // 账户ID
  "book_id": 1,           // 所属账本ID
  "name": "我的工资卡",    // 账户名称
  "account_type_id": 1,   // 账户类型ID
  "account_type": "银行卡", // 账户类型名称
  "balance": 5000.00,     // 当前余额
  "currency": "CNY",     // 货币类型
  "status": "active",    // 状态：active, inactive, archived
  "hidden_balance": false, // 是否隐藏余额
  "memo": "工商银行储蓄卡", // 备注
  "icon": "bank_icon",    // 图标
  "color": "#0080FF",    // 颜色
  "created_at": "2023-01-01T12:00:00Z", // 创建时间
  "updated_at": "2023-01-01T12:00:00Z"  // 更新时间
}
```

###  账户类型模型

```javascript
{
  "type_id": 1,           // 类型ID
  "name": "银行卡",       // 类型名称
  "icon": "bank_icon",    // 图标
  "color": "#0080FF",    // 颜色
  "is_system": true,      // 是否系统预设
  "created_at": "2023-01-01T12:00:00Z" // 创建时间
}
```

###  转账记录模型

```javascript
{
  "transfer_id": 1,       // 转账ID
  "book_id": 1,           // 所属账本ID
  "from_account_id": 1,   // 转出账户ID
  "from_account_name": "我的工资卡", // 转出账户名称
  "to_account_id": 2,     // 转入账户ID
  "to_account_name": "支付宝", // 转入账户名称
  "amount": 1000.00,      // 转账金额
  "currency": "CNY",     // 货币类型
  "transfer_date": "2023-01-05T00:00:00Z", // 转账日期
  "memo": "转账到支付宝",  // 备注
  "created_at": "2023-01-05T12:00:00Z" // 创建时间
}
```

###  账户分组模型

```javascript
{
  "group_id": 1,          // 分组ID
  "name": "日常消费",      // 分组名称
  "icon": "shopping_icon", // 图标
  "color": "#FF6B6B",    // 颜色
  "accounts_count": 3,    // 包含账户数
  "total_balance": 8000.00, // 总余额
  "created_at": "2023-01-01T12:00:00Z" // 创建时间
}
```

##  货币与汇率说明

###  支持的货币类型

- CNY (人民币)
- USD (美元)
- EUR (欧元)
- GBP (英镑)
- JPY (日元)
- KRW (韩元)
- HKD (港币)
- TWD (新台币)

###  汇率更新机制

- 系统支持手动更新汇率
- 可配置自动从第三方API获取最新汇率
- 汇率精度为6位小数

##  错误码说明

| 错误码 | 描述 |
|-------|------|
| 400 | 请求参数错误 |
| 401 | 未授权，需要登录 |
| 403 | 权限不足 |
| 404 | 账户不存在 |
| 500 | 服务器内部错误 |
| 700 | 账户名称已存在 |
| 701 | 余额不足 |
| 702 | 无效的账户状态 |
| 703 | 不支持的货币类型 |
| 704 | 不能对系统预设账户类型进行修改 |