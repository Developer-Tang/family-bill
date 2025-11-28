## 模块概述

账户管理模块负责管理用户的各类资产账户，包括现金、银行卡、支付宝、微信等，支持余额管理、转账和多币种。本模块提供了完整的账户生命周期管理和资金转移功能的API接口。

### 设计原则

1. **RESTful规范**：所有API接口严格遵循RESTful设计规范，使用标准HTTP方法和资源导向的URL路径
2. **本地缓存优先**：暂不集成远程缓存机制，如需使用缓存功能，严格限定为本地缓存方案
3. **完整的请求验证**：所有API接口包含严格的请求参数验证，确保数据完整性和安全性
4. **统一的错误处理**：采用统一的错误响应格式和错误码体系，便于客户端处理
5. **清晰的响应格式**：所有API响应采用标准JSON格式，包含明确的数据结构和状态信息
6. **可扩展性**：API设计考虑未来功能扩展，预留合理的扩展点

### 适用场景

- 用户需要管理多种类型的资产账户
- 用户需要在不同账户间进行资金转账
- 用户需要管理多币种账户和汇率
- 用户需要创建和管理账户分组
- 用户需要查看账户余额和交易记录
- 用户需要设置账户初始余额

## 接口清单

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

## 详细接口说明

### 获取账户类型列表

**功能描述**：获取系统支持的所有账户类型，包括系统默认类型和自定义类型

**请求方法**：GET
**URL路径**：/api/v1/account-types
**权限要求**：已登录用户
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**请求示例**
```http
GET /api/v1/account-types
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | array | 账户类型列表 |
| data[].type_id | integer | 账户类型ID |
| data[].name | string | 账户类型名称 |
| data[].icon | string | 账户类型图标 |
| data[].color | string | 账户类型颜色 |
| data[].is_system | boolean | 是否为系统默认类型 |

**响应示例**
```json
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

### 创建自定义账户类型

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

### 更新账户类型

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

### 删除账户类型

**功能描述**：删除自定义账户类型

**请求方法**：DELETE
**URL路径**：/api/v1/account-types/:id
**权限要求**：已登录用户
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**URL参数**
| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| id | integer | 是 | 账户类型ID |

**请求示例**
```http
DELETE /api/v1/account-types/:id
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | null | 无数据返回 |

**响应示例**
```json
{
  "code": 200,
  "message": "删除成功",
  "data": null
}
```

### 获取账户列表

**功能描述**：获取指定账本下的账户列表，支持分页和状态筛选

**请求方法**：GET
**URL路径**：/api/v1/accounts
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
| status | string | 否 | - | 账户状态：active, inactive, archived |
| page | integer | 否 | 1 | 页码 |
| limit | integer | 否 | 20 | 每页数量 |

**请求示例**
```http
GET /api/v1/accounts?book_id=1&status=active&page=1&limit=20
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 账户列表数据 |
| data.total | integer | 总记录数 |
| data.items | array | 账户列表 |
| data.items[].account_id | integer | 账户ID |
| data.items[].name | string | 账户名称 |
| data.items[].account_type | string | 账户类型 |
| data.items[].balance | number | 账户余额 |
| data.items[].currency | string | 货币类型 |
| data.items[].status | string | 账户状态 |
| data.items[].created_at | string | 创建时间 |

**响应示例**
```json
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

### 创建账户

**功能描述**：创建新的账户记录

**请求方法**：POST
**URL路径**：/api/v1/accounts
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
| name | string | 是 | - | 账户名称 | 1-50个字符，同一账本内名称唯一 |
| account_type_id | integer | 是 | - | 账户类型ID | 正整数，账户类型必须存在 |
| currency | string | 否 | "CNY" | 货币类型 | 有效的货币代码 |
| initial_balance | number | 否 | 0.00 | 初始余额 | 最多两位小数 |
| account_group_id | integer | 否 | null | 账户分组ID | 正整数或null |
| hidden_balance | boolean | 否 | false | 是否隐藏余额 | 布尔值 |
| memo | string | 否 | "" | 账户备注 | 0-200个字符 |
| icon | string | 否 | - | 账户图标 | 有效的图标名称 |
| color | string | 否 | - | 账户颜色 | 有效的十六进制颜色值 |

**请求示例**
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

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 新创建的账户信息 |
| data.account_id | integer | 账户ID |
| data.name | string | 账户名称 |
| data.balance | number | 账户余额 |
| data.currency | string | 货币类型 |
| data.status | string | 账户状态 |
| data.created_at | string | 创建时间 |

**成功响应示例**
```json
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
```

**失败响应示例**
```json
{
  "code": 400,
  "message": "创建失败：账户名称已存在",
  "data": null
}
```

### 获取账户详情

**功能描述**：获取指定账户的详细信息

**请求方法**：GET
**URL路径**：/api/v1/accounts/:id
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**URL参数**
| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| id | integer | 是 | 账户ID |

**请求示例**
```http
GET /api/v1/accounts/:id
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 账户详细信息 |
| data.account_id | integer | 账户ID |
| data.name | string | 账户名称 |
| data.account_type_id | integer | 账户类型ID |
| data.account_type | string | 账户类型名称 |
| data.balance | number | 账户余额 |
| data.currency | string | 货币类型 |
| data.status | string | 账户状态 |
| data.hidden_balance | boolean | 是否隐藏余额 |
| data.memo | string | 账户备注 |
| data.icon | string | 账户图标 |
| data.color | string | 账户颜色 |
| data.created_at | string | 创建时间 |
| data.updated_at | string | 更新时间 |

**响应示例**
```json
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

### 更新账户信息

**功能描述**：更新指定账户的信息

**请求方法**：PUT
**URL路径**：/api/v1/accounts/:id
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |
| Authorization | string | 是 | Bearer JWT令牌 |

**URL参数**
| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| id | integer | 是 | 账户ID |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| name | string | 否 | - | 账户名称 | 1-50个字符，同一账本内名称唯一 |
| memo | string | 否 | - | 账户备注 | 0-200个字符 |
| icon | string | 否 | - | 账户图标 | 有效的图标名称 |
| color | string | 否 | - | 账户颜色 | 有效的十六进制颜色值 |
| hidden_balance | boolean | 否 | - | 是否隐藏余额 | 布尔值 |

**请求示例**
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

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 更新后的账户信息 |
| data.account_id | integer | 账户ID |
| data.name | string | 账户名称 |
| data.memo | string | 账户备注 |
| data.icon | string | 账户图标 |
| data.color | string | 账户颜色 |
| data.hidden_balance | boolean | 是否隐藏余额 |
| data.updated_at | string | 更新时间 |

**成功响应示例**
```json
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

### 删除账户

**功能描述**：删除指定账户

**请求方法**：DELETE
**URL路径**：/api/v1/accounts/:id
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**URL参数**
| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| id | integer | 是 | 账户ID |

**请求示例**
```http
DELETE /api/v1/accounts/:id
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | null | 无数据返回 |

**成功响应示例**
```json
{
  "code": 200,
  "message": "删除成功",
  "data": null
}
```

### 调整账户余额

**功能描述**：调整指定账户的余额，支持增加或减少

**请求方法**：PUT
**URL路径**：/api/v1/accounts/:id/balance
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |
| Authorization | string | 是 | Bearer JWT令牌 |

**URL参数**
| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| id | integer | 是 | 账户ID |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| delta | number | 是 | - | 余额调整金额，正数表示增加，负数表示减少 | 最多两位小数 |
| reason | string | 否 | "" | 调整原因 | 0-200个字符 |
| date | string | 否 | 当前日期 | 调整日期，格式：YYYY-MM-DD | 有效的日期格式 |

**请求示例**
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

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 余额调整结果 |
| data.account_id | integer | 账户ID |
| data.old_balance | number | 调整前余额 |
| data.new_balance | number | 调整后余额 |
| data.delta | number | 调整金额 |
| data.adjusted_at | string | 调整时间 |

**成功响应示例**
```json
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

### 获取账户分组列表

**功能描述**：获取指定账本下的账户分组列表

**请求方法**：GET
**URL路径**：/api/v1/account-groups
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

**请求示例**
```http
GET /api/v1/account-groups?book_id=1
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | array | 账户分组列表 |
| data[].group_id | integer | 分组ID |
| data[].name | string | 分组名称 |
| data[].icon | string | 分组图标 |
| data[].color | string | 分组颜色 |
| data[].accounts_count | integer | 分组内账户数量 |
| data[].total_balance | number | 分组内账户总余额 |

**响应示例**
```json
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

### 创建账户分组

**功能描述**：创建新的账户分组

**请求方法**：POST
**URL路径**：/api/v1/account-groups
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
| name | string | 是 | - | 分组名称 | 1-50个字符，同一账本内名称唯一 |
| icon | string | 是 | - | 分组图标 | 有效的图标名称 |
| color | string | 是 | - | 分组颜色 | 有效的十六进制颜色值 |

**请求示例**
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

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 新创建的账户分组信息 |
| data.group_id | integer | 分组ID |
| data.name | string | 分组名称 |
| data.icon | string | 分组图标 |
| data.color | string | 分组颜色 |
| data.accounts_count | integer | 分组内账户数量 |
| data.total_balance | number | 分组内账户总余额 |

**成功响应示例**
```json
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

### 更新账户分组

**功能描述**：更新指定账户分组的信息

**请求方法**：PUT
**URL路径**：/api/v1/account-groups/:id
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |
| Authorization | string | 是 | Bearer JWT令牌 |

**URL参数**
| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| id | integer | 是 | 账户分组ID |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| name | string | 否 | - | 分组名称 | 1-50个字符，同一账本内名称唯一 |
| icon | string | 否 | - | 分组图标 | 有效的图标名称 |
| color | string | 否 | - | 分组颜色 | 有效的十六进制颜色值 |

**请求示例**
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

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 更新后的账户分组信息 |
| data.group_id | integer | 分组ID |
| data.name | string | 分组名称 |
| data.icon | string | 分组图标 |
| data.color | string | 分组颜色 |
| data.accounts_count | integer | 分组内账户数量 |
| data.total_balance | number | 分组内账户总余额 |

**成功响应示例**
```json
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

### 删除账户分组

**功能描述**：删除指定账户分组

**请求方法**：DELETE
**URL路径**：/api/v1/account-groups/:id
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**URL参数**
| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| id | integer | 是 | 账户分组ID |

**请求示例**
```http
DELETE /api/v1/account-groups/:id
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | null | 无数据返回 |

**成功响应示例**
```json
{
  "code": 200,
  "message": "删除成功",
  "data": null
}
```

### 账户加入分组

**功能描述**：将多个账户加入到指定分组

**请求方法**：POST
**URL路径**：/api/v1/account-groups/:id/accounts
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |
| Authorization | string | 是 | Bearer JWT令牌 |

**URL参数**
| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| id | integer | 是 | 账户分组ID |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| account_ids | array | 是 | - | 账户ID列表 | 数组元素为正整数，账户必须存在且属于同一账本 |

**请求示例**
```http
POST /api/v1/account-groups/:id/accounts
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "account_ids": [1, 2, 3]
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 分组更新后的信息 |
| data.group_id | integer | 分组ID |
| data.accounts_count | integer | 分组内账户数量 |
| data.total_balance | number | 分组内账户总余额 |

**成功响应示例**
```json
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

### 创建转账记录

**功能描述**：创建新的转账记录，支持同一账本内不同账户间的资金转移

**请求方法**：POST
**URL路径**：/api/v1/transfers
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
| from_account_id | integer | 是 | - | 转出账户ID | 正整数，账户必须存在且属于该账本 |
| to_account_id | integer | 是 | - | 转入账户ID | 正整数，账户必须存在且属于该账本 |
| amount | number | 是 | - | 转账金额 | 大于0，最多两位小数 |
| currency | string | 否 | "CNY" | 货币类型 | 有效的货币代码 |
| transfer_date | string | 否 | 当前日期 | 转账日期，格式：YYYY-MM-DD | 有效的日期格式 |
| memo | string | 否 | "" | 转账备注 | 0-200个字符 |
| exchange_rate | number | 否 | 1.0 | 汇率 | 大于0，最多六位小数 |

**请求示例**
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

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 转账记录信息 |
| data.transfer_id | integer | 转账记录ID |
| data.from_account | object | 转出账户信息 |
| data.from_account.id | integer | 转出账户ID |
| data.from_account.name | string | 转出账户名称 |
| data.from_account.balance | number | 转出账户余额 |
| data.to_account | object | 转入账户信息 |
| data.to_account.id | integer | 转入账户ID |
| data.to_account.name | string | 转入账户名称 |
| data.to_account.balance | number | 转入账户余额 |
| data.amount | number | 转账金额 |
| data.currency | string | 货币类型 |
| data.transfer_date | string | 转账日期 |
| data.memo | string | 转账备注 |

**成功响应示例**
```json
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
```

**失败响应示例**
```json
{
  "code": 400,
  "message": "转账失败：转出账户余额不足",
  "data": null
}
```

### 获取转账记录列表

**功能描述**：获取指定账本下的转账记录列表，支持分页和日期范围筛选

**请求方法**：GET
**URL路径**：/api/v1/transfers
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
| page | integer | 否 | 1 | 页码 |
| limit | integer | 否 | 20 | 每页数量 |
| start_date | string | 否 | - | 开始日期，格式：YYYY-MM-DD |
| end_date | string | 否 | - | 结束日期，格式：YYYY-MM-DD |

**请求示例**
```http
GET /api/v1/transfers?book_id=1&page=1&limit=20&start_date=2023-01-01&end_date=2023-01-31
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 转账记录列表数据 |
| data.total | integer | 总记录数 |
| data.items | array | 转账记录列表 |
| data.items[].transfer_id | integer | 转账记录ID |
| data.items[].from_account_id | integer | 转出账户ID |
| data.items[].from_account_name | string | 转出账户名称 |
| data.items[].to_account_id | integer | 转入账户ID |
| data.items[].to_account_name | string | 转入账户名称 |
| data.items[].amount | number | 转账金额 |
| data.items[].currency | string | 货币类型 |
| data.items[].transfer_date | string | 转账日期 |
| data.items[].memo | string | 转账备注 |

**响应示例**
```json
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

### 获取转账记录详情

**功能描述**：获取指定转账记录的详细信息

**请求方法**：GET
**URL路径**：/api/v1/transfers/:id
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**URL参数**
| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| id | integer | 是 | 转账记录ID |

**请求示例**
```http
GET /api/v1/transfers/:id
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 转账记录详细信息 |
| data.transfer_id | integer | 转账记录ID |
| data.book_id | integer | 账本ID |
| data.from_account_id | integer | 转出账户ID |
| data.from_account_name | string | 转出账户名称 |
| data.from_account_type | string | 转出账户类型 |
| data.to_account_id | integer | 转入账户ID |
| data.to_account_name | string | 转入账户名称 |
| data.to_account_type | string | 转入账户类型 |
| data.amount | number | 转账金额 |
| data.currency | string | 货币类型 |
| data.transfer_date | string | 转账日期 |
| data.memo | string | 转账备注 |
| data.created_at | string | 创建时间 |

**响应示例**
```json
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

### 获取货币列表

**功能描述**：获取系统支持的所有货币列表

**请求方法**：GET
**URL路径**：/api/v1/currencies
**权限要求**：已登录用户
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**请求示例**
```http
GET /api/v1/currencies
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | array | 货币列表 |
| data[].code | string | 货币代码 |
| data[].name | string | 货币名称 |
| data[].symbol | string | 货币符号 |

**响应示例**
```json
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

### 获取汇率

**功能描述**：获取指定货币对的汇率信息

**请求方法**：GET
**URL路径**：/api/v1/exchange-rates
**权限要求**：已登录用户
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| from | string | 是 | - | 基准货币代码 |
| to | string | 是 | - | 目标货币代码，多个用逗号分隔 |

**请求示例**
```http
GET /api/v1/exchange-rates?from=CNY&to=USD,EUR,GBP
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 汇率信息 |
| data.base | string | 基准货币代码 |
| data.rates | object | 汇率数据，键为目标货币代码，值为汇率 |
| data.updated_at | string | 汇率更新时间 |

**响应示例**
```json
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

### 更新汇率

**功能描述**：更新指定货币的汇率信息

**请求方法**：POST
**URL路径**：/api/v1/exchange-rates
**权限要求**：账本管理员或所有者
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |
| Authorization | string | 是 | Bearer JWT令牌 |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| base | string | 是 | - | 基准货币代码 | 有效的货币代码 |
| rates | object | 是 | - | 汇率数据，键为目标货币代码，值为汇率 | 汇率值大于0，最多六位小数 |

**请求示例**
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

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 汇率更新结果 |
| data.base | string | 基准货币代码 |
| data.updated_at | string | 汇率更新时间 |

**成功响应示例**
```json
{
  "code": 200,
  "message": "更新成功",
  "data": {
    "base": "CNY",
    "updated_at": "2023-01-01T12:00:00Z"
  }
}
```

### 修改账户状态

**功能描述**：修改指定账户的状态，支持激活、停用和归档操作

**请求方法**：PUT
**URL路径**：/api/v1/accounts/:id/status
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |
| Authorization | string | 是 | Bearer JWT令牌 |

**URL参数**
| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| id | integer | 是 | 账户ID |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| status | string | 是 | - | 账户状态：active, inactive, archived | 必须是有效值之一 |
| reason | string | 否 | "" | 状态变更原因 | 0-200个字符 |

**请求示例**
```http
PUT /api/v1/accounts/:id/status
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "status": "inactive",  // active, inactive, archived
  "reason": "账户不再使用"
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 账户状态更新结果 |
| data.account_id | integer | 账户ID |
| data.name | string | 账户名称 |
| data.status | string | 更新后的账户状态 |
| data.updated_at | string | 更新时间 |

**成功响应示例**
```json
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

## 数据模型

### 账户模型

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

### 账户类型模型

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

### 转账记录模型

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

### 账户分组模型

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

## 货币与汇率说明

### 支持的货币类型

- CNY (人民币)
- USD (美元)
- EUR (欧元)
- GBP (英镑)
- JPY (日元)
- KRW (韩元)
- HKD (港币)
- TWD (新台币)

### 汇率更新机制

- 系统支持手动更新汇率
- 可配置自动从第三方API获取最新汇率
- 汇率精度为6位小数

## 错误码说明

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