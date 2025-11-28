## 模块概述

高级功能模块是家庭记账系统的扩展功能集合，旨在增强用户体验和系统智能化水平。本模块提供预算管理、标签管理、共享账本、账单提醒、智能分类和汇率换算等核心扩展功能的RESTful API接口，为用户提供更全面、智能的财务管理能力。

### 设计原则

1. **RESTful规范**：所有API接口严格遵循RESTful设计规范，使用标准HTTP方法和资源导向的URL路径
2. **本地缓存优先**：暂不集成远程缓存机制，如需使用缓存功能，严格限定为本地缓存方案
3. **完整的请求验证**：所有API接口包含严格的请求参数验证，确保数据完整性和安全性
4. **统一的错误处理**：采用统一的错误响应格式和错误码体系，便于客户端处理
5. **清晰的响应格式**：所有API响应采用标准JSON格式，包含明确的数据结构和状态信息
6. **可扩展性**：API设计考虑未来功能扩展，预留合理的扩展点

### 适用场景

- 家庭用户需要进行预算规划和执行跟踪
- 多成员家庭需要共享账本并进行权限管理
- 用户需要智能分类和消费分析
- 跨国家庭需要进行多币种记账和汇率换算
- 用户需要设置账单提醒，避免遗漏重要支出

## 接口清单

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

## 详细接口说明

### 获取预算列表

**功能描述**：获取指定账本的预算列表，支持分页、筛选和排序

**请求方法**：GET
**URL路径**：/api/v1/budgets
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| book_id | integer | 是 | - | 账本ID | 正整数，必须存在且用户有访问权限 |
| status | string | 否 | - | 预算状态：active, paused, expired | 枚举值 |
| period | string | 否 | - | 预算周期：monthly, yearly, custom | 枚举值 |
| page | integer | 否 | 1 | 页码 | 正整数 |
| page_size | integer | 否 | 20 | 每页条数 | 1-100之间的整数 |
| sort_by | string | 否 | "created_at" | 排序字段 | 可选值：created_at, updated_at, name, amount |
| sort_order | string | 否 | "desc" | 排序顺序 | 可选值：asc, desc |

**请求示例**
```http
GET /api/v1/budgets?book_id=1&status=active&page=1&page_size=20
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| total | integer | 总记录数 |
| page | integer | 当前页码 |
| page_size | integer | 每页条数 |
| list | array | 预算列表 |
| list[].budget_id | integer | 预算ID |
| list[].name | string | 预算名称 |
| list[].amount | number | 预算金额 |
| list[].category_ids | array | 关联的分类IDs |
| list[].period | string | 预算周期 |
| list[].start_date | string | 开始日期 |
| list[].end_date | string | 结束日期 |
| list[].notify_threshold | integer | 通知阈值 |
| list[].is_recurring | boolean | 是否循环 |
| list[].status | string | 预算状态 |
| list[].created_at | string | 创建时间 |
| list[].updated_at | string | 更新时间 |

**成功响应**
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
        "created_at": "2023-01-01T00:00:00Z",
        "updated_at": "2023-01-01T00:00:00Z"
      }
    ]
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：页码必须为正整数 | 检查page参数是否为正整数 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有查看预算的权限 |
| 404 | 账本不存在 | 检查book_id是否正确 |

**本地缓存策略**：获取成功后，本地缓存该预算列表，缓存时间15分钟

### 创建预算

**功能描述**：创建新的预算记录，用于跟踪特定周期内的收支情况

**请求方法**：POST
**URL路径**：/api/v1/budgets
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
| book_id | integer | 是 | - | 账本ID | 正整数，必须存在且用户有访问权限 |
| name | string | 是 | - | 预算名称 | 1-50个字符，同一账本内名称唯一 |
| amount | number | 是 | - | 预算金额 | 大于0，最多两位小数 |
| category_ids | array | 否 | [] | 关联的分类IDs | 数组元素为正整数，分类必须存在 |
| period | string | 是 | - | 预算周期 | 枚举值：monthly, yearly, custom |
| start_date | string | 是 | - | 开始日期 | YYYY-MM-DD格式，必须小于等于end_date |
| end_date | string | 否 | - | 结束日期 | YYYY-MM-DD格式，仅当period为custom时必填，必须大于等于start_date |
| notify_threshold | integer | 否 | 80 | 通知阈值（百分比） | 1-100之间的整数 |
| is_recurring | boolean | 否 | false | 是否循环 | 布尔值 |

**请求示例**
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

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| budget_id | integer | 预算ID |
| name | string | 预算名称 |
| amount | number | 预算金额 |
| category_ids | array | 关联的分类IDs |
| period | string | 预算周期 |
| start_date | string | 开始日期 |
| end_date | string | 结束日期 |
| notify_threshold | integer | 通知阈值 |
| is_recurring | boolean | 是否循环 |
| status | string | 预算状态：active, paused, expired |
| created_at | string | 创建时间 |
| updated_at | string | 更新时间 |

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
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：预算金额必须大于0 | 检查amount参数是否大于0 |
| 400 | 请求参数错误：同一账本内预算名称已存在 | 更换预算名称 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有创建预算的权限 |
| 404 | 账本不存在 | 检查book_id是否正确 |
| 2000 | 预算创建失败 | 检查预算参数是否正确，特别是金额、日期范围等 |

**本地缓存策略**：创建成功后，本地缓存该预算信息，缓存时间1小时

### 更新预算

**功能描述**：更新指定预算的信息，包括名称、金额、周期等

**请求方法**：PUT
**URL路径**：/api/v1/budgets/:id
**权限要求**：账本管理员或所有者
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 描述 | 验证规则 |
|--------|------|------|------|----------|
| id | integer | 是 | 预算ID | 正整数，预算必须存在且用户有管理权限 |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| name | string | 否 | - | 预算名称 | 1-50个字符，同一账本内名称唯一 |
| amount | number | 否 | - | 预算金额 | 大于0，最多两位小数 |
| category_ids | array | 否 | - | 关联的分类IDs | 数组元素为正整数，分类必须存在 |
| period | string | 否 | - | 预算周期 | 枚举值：monthly, yearly, custom |
| start_date | string | 否 | - | 开始日期 | YYYY-MM-DD格式，必须小于等于end_date |
| end_date | string | 否 | - | 结束日期 | YYYY-MM-DD格式，必须大于等于start_date |
| notify_threshold | integer | 否 | - | 通知阈值（百分比） | 1-100之间的整数 |
| is_recurring | boolean | 否 | - | 是否循环 | 布尔值 |
| status | string | 否 | - | 预算状态：active, paused | 枚举值 |

**请求示例**
```http
PUT /api/v1/budgets/:id
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "name": "更新后的月度餐饮预算",
  "amount": 6000,
  "status": "active"
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| budget_id | integer | 预算ID |
| name | string | 预算名称 |
| amount | number | 预算金额 |
| category_ids | array | 关联的分类IDs |
| period | string | 预算周期 |
| start_date | string | 开始日期 |
| end_date | string | 结束日期 |
| notify_threshold | integer | 通知阈值 |
| is_recurring | boolean | 是否循环 |
| status | string | 预算状态 |
| updated_at | string | 更新时间 |

**成功响应**
```json
{
  "code": 200,
  "message": "更新成功",
  "data": {
    "budget_id": 1,
    "name": "更新后的月度餐饮预算",
    "amount": 6000,
    "category_ids": [5, 6],
    "period": "monthly",
    "start_date": "2023-01-01",
    "end_date": "2023-01-31",
    "notify_threshold": 80,
    "is_recurring": true,
    "status": "active",
    "updated_at": "2023-01-10T14:30:00Z"
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：预算金额必须大于0 | 检查amount参数是否大于0 |
| 400 | 请求参数错误：同一账本内预算名称已存在 | 更换预算名称 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有更新预算的权限 |
| 404 | 预算不存在 | 检查id参数是否正确 |
| 2000 | 预算更新失败 | 检查预算参数是否正确，特别是金额、日期范围等 |

**本地缓存策略**：更新成功后，更新本地缓存，缓存时间1小时

### 删除预算

**功能描述**：删除指定的预算记录

**请求方法**：DELETE
**URL路径**：/api/v1/budgets/:id
**权限要求**：账本管理员或所有者
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 描述 | 验证规则 |
|--------|------|------|------|----------|
| id | integer | 是 | 预算ID | 正整数，预算必须存在且用户有管理权限 |

**请求示例**
```http
DELETE /api/v1/budgets/:id
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| budget_id | integer | 已删除的预算ID |

**成功响应**
```json
{
  "code": 200,
  "message": "删除成功",
  "data": {
    "budget_id": 1
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有删除预算的权限 |
| 404 | 预算不存在 | 检查id参数是否正确 |
| 2000 | 预算删除失败 | 预算可能已被使用或存在其他关联数据 |

**本地缓存策略**：删除成功后，清除本地缓存中该预算的相关信息

### 获取预算汇总信息

**功能描述**：获取指定账本的预算汇总信息，包括总预算、已使用金额、剩余金额等

**请求方法**：GET
**URL路径**：/api/v1/budgets/summary
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| book_id | integer | 是 | - | 账本ID | 正整数，必须存在且用户有访问权限 |
| period | string | 否 | - | 预算周期：monthly, yearly, custom | 枚举值 |
| status | string | 否 | - | 预算状态：active, paused, expired | 枚举值 |

**请求示例**
```http
GET /api/v1/budgets/summary?book_id=1&period=monthly&status=active
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| total_budget | number | 总预算金额 |
| used_amount | number | 已使用金额 |
| remaining_amount | number | 剩余金额 |
| budget_count | integer | 预算数量 |
| over_budget_count | integer | 超支预算数量 |
| approaching_budget_count | integer | 即将超支预算数量（使用率≥80%） |
| status_distribution | object | 预算状态分布 |
| status_distribution.active | integer | 活跃预算数量 |
| status_distribution.paused | integer | 暂停预算数量 |
| status_distribution.expired | integer | 过期预算数量 |

**成功响应**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total_budget": 20000,
    "used_amount": 8000,
    "remaining_amount": 12000,
    "budget_count": 5,
    "over_budget_count": 1,
    "approaching_budget_count": 2,
    "status_distribution": {
      "active": 3,
      "paused": 1,
      "expired": 1
    }
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有查看预算的权限 |
| 404 | 账本不存在 | 检查book_id是否正确 |

**本地缓存策略**：获取成功后，本地缓存该预算汇总信息，缓存时间15分钟

### 获取预算执行状态

**功能描述**：获取指定预算的执行状态，包括已支出金额、剩余金额、执行进度等

**请求方法**：GET
**URL路径**：/api/v1/budgets/:id/status
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 描述 | 验证规则 |
|--------|------|------|------|----------|
| id | integer | 是 | 预算ID | 正整数，预算必须存在且用户有访问权限 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| date | string | 否 | 当前日期 | 查询日期 | YYYY-MM-DD格式 |

**请求示例**
```http
GET /api/v1/budgets/:id/status?date=2023-01-15
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| budget_id | integer | 预算ID |
| total_budget | number | 总预算金额 |
| spent_amount | number | 已支出金额 |
| remaining_amount | number | 剩余金额 |
| percentage | integer | 执行进度百分比 |
| is_over_budget | boolean | 是否已超预算 |
| daily_average | number | 日均支出 |
| trend | string | 消费趋势：normal, warning, danger |
| transactions_count | integer | 相关交易数量 |
| forecast | object | 预测信息 |
| forecast.estimated_spend | number | 预计总支出 |
| forecast.estimated_remaining | number | 预计剩余金额 |
| forecast.will_exceed | boolean | 是否预计超支 |

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
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：日期格式不正确 | 检查date参数格式是否为YYYY-MM-DD |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有查看该预算的权限 |
| 404 | 预算不存在 | 检查id参数是否正确 |

**本地缓存策略**：获取成功后，本地缓存该预算执行状态，缓存时间15分钟

### 获取标签列表

**功能描述**：获取指定账本的标签列表，支持分页、筛选和排序

**请求方法**：GET
**URL路径**：/api/v1/tags
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| book_id | integer | 是 | - | 账本ID | 正整数，必须存在且用户有访问权限 |
| name | string | 否 | - | 标签名称模糊搜索 | 1-30个字符 |
| page | integer | 否 | 1 | 页码 | 正整数 |
| page_size | integer | 否 | 20 | 每页条数 | 1-100之间的整数 |
| sort_by | string | 否 | "created_at" | 排序字段 | 可选值：created_at, updated_at, name, usage_count |
| sort_order | string | 否 | "desc" | 排序顺序 | 可选值：asc, desc |

**请求示例**
```http
GET /api/v1/tags?book_id=1&name=生日&page=1&page_size=20
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| total | integer | 总记录数 |
| page | integer | 当前页码 |
| page_size | integer | 每页条数 |
| list | array | 标签列表 |
| list[].tag_id | integer | 标签ID |
| list[].book_id | integer | 账本ID |
| list[].name | string | 标签名称 |
| list[].color | string | 标签颜色 |
| list[].icon | string | 标签图标 |
| list[].description | string | 标签描述 |
| list[].usage_count | integer | 使用次数 |
| list[].created_at | string | 创建时间 |
| list[].updated_at | string | 更新时间 |

**成功响应**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total": 10,
    "page": 1,
    "page_size": 20,
    "list": [
      {
        "tag_id": 1,
        "book_id": 1,
        "name": "生日礼物",
        "color": "#FF5733",
        "icon": "gift",
        "description": "生日礼物相关消费",
        "usage_count": 5,
        "created_at": "2023-01-05T12:30:00Z",
        "updated_at": "2023-01-05T12:30:00Z"
      }
    ]
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：页码必须为正整数 | 检查page参数是否为正整数 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有查看标签的权限 |
| 404 | 账本不存在 | 检查book_id是否正确 |

**本地缓存策略**：获取成功后，本地缓存该标签列表，缓存时间1小时

### 创建标签

**功能描述**：创建新的标签，用于对收支记录进行分类和标记

**请求方法**：POST
**URL路径**：/api/v1/tags
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
| book_id | integer | 是 | - | 账本ID | 正整数，必须存在且用户有访问权限 |
| name | string | 是 | - | 标签名称 | 1-30个字符，同一账本内名称唯一 |
| color | string | 否 | "#3366FF" | 标签颜色 | 有效的十六进制颜色代码 |
| icon | string | 否 | "tag" | 标签图标 | 1-20个字符，必须是系统支持的图标名称 |
| description | string | 否 | "" | 标签描述 | 0-100个字符 |

**请求示例**
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

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| tag_id | integer | 标签ID |
| book_id | integer | 账本ID |
| name | string | 标签名称 |
| color | string | 标签颜色 |
| icon | string | 标签图标 |
| description | string | 标签描述 |
| usage_count | integer | 使用次数 |
| created_at | string | 创建时间 |
| updated_at | string | 更新时间 |

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
    "created_at": "2023-01-05T12:30:00Z",
    "updated_at": "2023-01-05T12:30:00Z"
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：标签名称已存在 | 更换标签名称 |
| 400 | 请求参数错误：颜色格式不正确 | 检查color参数是否为有效的十六进制颜色代码 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有创建标签的权限 |
| 404 | 账本不存在 | 检查book_id是否正确 |
| 2001 | 标签创建失败 | 检查标签参数是否正确，特别是名称、颜色等 |

**本地缓存策略**：创建成功后，本地缓存该标签信息，缓存时间24小时

### 更新标签

**功能描述**：更新指定标签的信息，包括名称、颜色、图标等

**请求方法**：PUT
**URL路径**：/api/v1/tags/:id
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 描述 | 验证规则 |
|--------|------|------|------|----------|
| id | integer | 是 | 标签ID | 正整数，标签必须存在且用户有访问权限 |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| name | string | 否 | - | 标签名称 | 1-30个字符，同一账本内名称唯一 |
| color | string | 否 | - | 标签颜色 | 有效的十六进制颜色代码 |
| icon | string | 否 | - | 标签图标 | 1-20个字符，必须是系统支持的图标名称 |
| description | string | 否 | - | 标签描述 | 0-100个字符 |

**请求示例**
```http
PUT /api/v1/tags/:id
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "name": "更新后的生日礼物",
  "color": "#FF0000",
  "icon": "birthday-cake"
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| tag_id | integer | 标签ID |
| book_id | integer | 账本ID |
| name | string | 标签名称 |
| color | string | 标签颜色 |
| icon | string | 标签图标 |
| description | string | 标签描述 |
| usage_count | integer | 使用次数 |
| updated_at | string | 更新时间 |

**成功响应**
```json
{
  "code": 200,
  "message": "更新成功",
  "data": {
    "tag_id": 1,
    "book_id": 1,
    "name": "更新后的生日礼物",
    "color": "#FF0000",
    "icon": "birthday-cake",
    "description": "生日礼物相关消费",
    "usage_count": 5,
    "updated_at": "2023-01-10T14:30:00Z"
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：标签名称已存在 | 更换标签名称 |
| 400 | 请求参数错误：颜色格式不正确 | 检查color参数是否为有效的十六进制颜色代码 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有更新标签的权限 |
| 404 | 标签不存在 | 检查id参数是否正确 |
| 2001 | 标签更新失败 | 检查标签参数是否正确，特别是名称、颜色等 |

**本地缓存策略**：更新成功后，更新本地缓存，缓存时间24小时

### 删除标签

**功能描述**：删除指定的标签

**请求方法**：DELETE
**URL路径**：/api/v1/tags/:id
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 描述 | 验证规则 |
|--------|------|------|------|----------|
| id | integer | 是 | 标签ID | 正整数，标签必须存在且用户有访问权限 |

**请求示例**
```http
DELETE /api/v1/tags/:id
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| tag_id | integer | 已删除的标签ID |

**成功响应**
```json
{
  "code": 200,
  "message": "删除成功",
  "data": {
    "tag_id": 1
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有删除标签的权限 |
| 404 | 标签不存在 | 检查id参数是否正确 |
| 2001 | 标签删除失败 | 标签可能已被使用或存在其他关联数据 |

**本地缓存策略**：删除成功后，清除本地缓存中该标签的相关信息

### 批量创建标签

**功能描述**：批量创建多个标签，提高标签创建效率

**请求方法**：POST
**URL路径**：/api/v1/tags/batch
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
| book_id | integer | 是 | - | 账本ID | 正整数，必须存在且用户有访问权限 |
| tags | array | 是 | - | 标签列表 | 数组长度1-50，每个元素为标签对象 |
| tags[].name | string | 是 | - | 标签名称 | 1-30个字符，同一账本内名称唯一 |
| tags[].color | string | 否 | "#3366FF" | 标签颜色 | 有效的十六进制颜色代码 |
| tags[].icon | string | 否 | "tag" | 标签图标 | 1-20个字符，必须是系统支持的图标名称 |
| tags[].description | string | 否 | "" | 标签描述 | 0-100个字符 |

**请求示例**
```http
POST /api/v1/tags/batch
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "tags": [
    {
      "name": "标签1",
      "color": "#FF0000",
      "icon": "tag"
    },
    {
      "name": "标签2",
      "color": "#00FF00",
      "icon": "tag"
    }
  ]
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| success_count | integer | 成功创建的标签数量 |
| failed_count | integer | 创建失败的标签数量 |
| created_tags | array | 成功创建的标签列表 |
| created_tags[].tag_id | integer | 标签ID |
| created_tags[].name | string | 标签名称 |
| created_tags[].color | string | 标签颜色 |
| created_tags[].icon | string | 标签图标 |
| created_tags[].description | string | 标签描述 |
| created_tags[].usage_count | integer | 使用次数 |
| created_tags[].created_at | string | 创建时间 |
| created_tags[].updated_at | string | 更新时间 |

**成功响应**
```json
{
  "code": 200,
  "message": "批量创建成功",
  "data": {
    "success_count": 2,
    "failed_count": 0,
    "created_tags": [
      {
        "tag_id": 2,
        "name": "标签1",
        "color": "#FF0000",
        "icon": "tag",
        "description": "",
        "usage_count": 0,
        "created_at": "2023-01-10T14:30:00Z",
        "updated_at": "2023-01-10T14:30:00Z"
      },
      {
        "tag_id": 3,
        "name": "标签2",
        "color": "#00FF00",
        "icon": "tag",
        "description": "",
        "usage_count": 0,
        "created_at": "2023-01-10T14:30:00Z",
        "updated_at": "2023-01-10T14:30:00Z"
      }
    ]
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：标签列表不能为空 | 检查tags参数是否为空数组 |
| 400 | 请求参数错误：标签名称已存在 | 检查标签名称是否已存在 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有创建标签的权限 |
| 404 | 账本不存在 | 检查book_id是否正确 |
| 2001 | 标签创建失败 | 检查标签参数是否正确，特别是名称、颜色等 |

**本地缓存策略**：创建成功后，本地缓存创建的标签信息，缓存时间24小时

### 获取推荐标签

**功能描述**：根据用户的记账习惯，智能推荐标签

**请求方法**：GET
**URL路径**：/api/v1/tags/recommend
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| book_id | integer | 是 | - | 账本ID | 正整数，必须存在且用户有访问权限 |
| count | integer | 否 | 10 | 推荐数量 | 1-50之间的整数 |
| transaction_description | string | 否 | - | 交易描述 | 用于更精准的推荐 |

**请求示例**
```http
GET /api/v1/tags/recommend?book_id=1&count=5&transaction_description=超市购物
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| recommended_tags | array | 推荐标签列表 |
| recommended_tags[].tag_id | integer | 标签ID |
| recommended_tags[].name | string | 标签名称 |
| recommended_tags[].color | string | 标签颜色 |
| recommended_tags[].icon | string | 标签图标 |
| recommended_tags[].description | string | 标签描述 |
| recommended_tags[].usage_count | integer | 使用次数 |
| recommended_tags[].relevance_score | number | 相关度分数 |

**成功响应**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "recommended_tags": [
      {
        "tag_id": 1,
        "name": "日常购物",
        "color": "#3366FF",
        "icon": "shopping-cart",
        "description": "日常购物相关消费",
        "usage_count": 10,
        "relevance_score": 0.95
      },
      {
        "tag_id": 2,
        "name": "杂货",
        "color": "#3366FF",
        "icon": "tag",
        "description": "杂货相关消费",
        "usage_count": 5,
        "relevance_score": 0.85
      }
    ]
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：推荐数量必须在1-50之间 | 检查count参数是否在1-50之间 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有获取推荐标签的权限 |
| 404 | 账本不存在 | 检查book_id是否正确 |

**本地缓存策略**：获取成功后，本地缓存推荐标签列表，缓存时间1小时

### 获取共享账本列表

**功能描述**：获取当前用户参与的共享账本列表

**请求方法**：GET
**URL路径**：/api/v1/shared/books
**权限要求**：登录用户
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| page | integer | 否 | 1 | 页码 | 正整数 |
| page_size | integer | 否 | 20 | 每页条数 | 1-100之间的整数 |
| sort_by | string | 否 | "created_at" | 排序字段 | 可选值：created_at, updated_at, name |
| sort_order | string | 否 | "desc" | 排序顺序 | 可选值：asc, desc |

**请求示例**
```http
GET /api/v1/shared/books?page=1&page_size=20
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| total | integer | 总记录数 |
| page | integer | 当前页码 |
| page_size | integer | 每页条数 |
| list | array | 共享账本列表 |
| list[].book_id | integer | 账本ID |
| list[].name | string | 账本名称 |
| list[].description | string | 账本描述 |
| list[].owner_id | integer | 所有者ID |
| list[].owner_name | string | 所有者名称 |
| list[].permission | string | 当前用户权限 | 枚举值：viewer, editor, manager, owner |
| list[].member_count | integer | 成员数量 |
| list[].created_at | string | 创建时间 |
| list[].updated_at | string | 更新时间 |

**成功响应**
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
        "book_id": 1,
        "name": "家庭账本",
        "description": "家庭共同记账",
        "owner_id": 1,
        "owner_name": "张三",
        "permission": "owner",
        "member_count": 3,
        "created_at": "2023-01-01T00:00:00Z",
        "updated_at": "2023-01-01T00:00:00Z"
      }
    ]
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |

**本地缓存策略**：获取成功后，本地缓存共享账本列表，缓存时间1小时

### 获取账本成员列表

**功能描述**：获取指定共享账本的成员列表

**请求方法**：GET
**URL路径**：/api/v1/shared/books/:bookId/members
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 描述 | 验证规则 |
|--------|------|------|------|----------|
| bookId | integer | 是 | 账本ID | 正整数，账本必须存在且用户有访问权限 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| page | integer | 否 | 1 | 页码 | 正整数 |
| page_size | integer | 否 | 20 | 每页条数 | 1-100之间的整数 |
| sort_by | string | 否 | "joined_at" | 排序字段 | 可选值：joined_at, username, permission |
| sort_order | string | 否 | "asc" | 排序顺序 | 可选值：asc, desc |

**请求示例**
```http
GET /api/v1/shared/books/:bookId/members?page=1&page_size=20
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| total | integer | 总记录数 |
| page | integer | 当前页码 |
| page_size | integer | 每页条数 |
| list | array | 成员列表 |
| list[].member_id | integer | 成员ID |
| list[].user_id | integer | 用户ID |
| list[].username | string | 用户名 |
| list[].email | string | 邮箱 |
| list[].permission | string | 权限级别 | 枚举值：viewer, editor, manager, owner |
| list[].joined_at | string | 加入时间 |
| list[].last_accessed | string | 最后访问时间 |
| list[].status | string | 成员状态 | 枚举值：active, suspended |

**成功响应**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total": 3,
    "page": 1,
    "page_size": 20,
    "list": [
      {
        "member_id": 1,
        "user_id": 1,
        "username": "张三",
        "email": "zhangsan@example.com",
        "permission": "owner",
        "joined_at": "2023-01-01T00:00:00Z",
        "last_accessed": "2023-01-15T14:30:00Z",
        "status": "active"
      }
    ]
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有查看成员列表的权限 |
| 404 | 账本不存在 | 检查bookId是否正确 |

**本地缓存策略**：获取成功后，本地缓存成员列表，缓存时间30分钟

### 添加共享成员

**功能描述**：邀请成员加入共享账本，设置其权限级别

**请求方法**：POST
**URL路径**：/api/v1/shared/books/:bookId/members
**权限要求**：账本管理员或所有者
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 描述 | 验证规则 |
|--------|------|------|------|----------|
| bookId | integer | 是 | 账本ID | 正整数，账本必须存在且用户有管理权限 |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| email | string | 是 | - | 被邀请人邮箱 | 有效的邮箱格式 |
| permission | string | 是 | - | 权限级别 | 枚举值：viewer, editor, manager, owner |
| message | string | 否 | "" | 邀请消息 | 0-200个字符 |

**请求示例**
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

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| invitation_id | string | 邀请ID |
| book_id | integer | 账本ID |
| recipient_email | string | 被邀请人邮箱 |
| permission | string | 权限级别 |
| status | string | 邀请状态：pending, accepted, rejected, expired |
| expires_at | string | 邀请过期时间 |
| created_at | string | 邀请创建时间 |

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
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：邮箱格式不正确 | 检查email参数格式是否有效 |
| 400 | 请求参数错误：无效的权限级别 | 检查permission参数是否为有效值 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足：您没有邀请成员的权限 | 检查当前用户是否有邀请成员的权限 |
| 403 | 共享权限不足 | 只有管理员或账本所有者可以邀请成员 |
| 404 | 账本不存在 | 检查bookId是否正确 |
| 2008 | 邀请已过期 | 邀请链接已过期，请重新发送邀请 |

**本地缓存策略**：邀请发送成功后，本地缓存邀请信息，缓存时间1小时

### 修改成员权限

**功能描述**：修改共享账本中指定成员的权限级别

**请求方法**：PUT
**URL路径**：/api/v1/shared/books/:bookId/members/:userId
**权限要求**：账本管理员或所有者
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 描述 | 验证规则 |
|--------|------|------|------|----------|
| bookId | integer | 是 | 账本ID | 正整数，账本必须存在且用户有管理权限 |
| userId | integer | 是 | 用户ID | 正整数，用户必须是账本成员 |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| permission | string | 是 | - | 新的权限级别 | 枚举值：viewer, editor, manager, owner |

**请求示例**
```http
PUT /api/v1/shared/books/:bookId/members/:userId
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "permission": "manager"
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| member_id | integer | 成员ID |
| user_id | integer | 用户ID |
| username | string | 用户名 |
| permission | string | 新的权限级别 |
| updated_at | string | 更新时间 |

**成功响应**
```json
{
  "code": 200,
  "message": "权限修改成功",
  "data": {
    "member_id": 2,
    "user_id": 3,
    "username": "李四",
    "permission": "manager",
    "updated_at": "2023-01-10T14:30:00Z"
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：无效的权限级别 | 检查permission参数是否为有效值 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足：您没有修改成员权限的权限 | 检查当前用户是否有修改成员权限的权限 |
| 404 | 账本不存在 | 检查bookId是否正确 |
| 404 | 用户不是该账本的成员 | 检查userId是否正确，或用户是否是账本成员 |

**本地缓存策略**：修改成功后，更新本地缓存中的成员权限信息，缓存时间30分钟

### 移除共享成员

**功能描述**：将成员从共享账本中移除

**请求方法**：DELETE
**URL路径**：/api/v1/shared/books/:bookId/members/:userId
**权限要求**：账本管理员或所有者
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 描述 | 验证规则 |
|--------|------|------|------|----------|
| bookId | integer | 是 | 账本ID | 正整数，账本必须存在且用户有管理权限 |
| userId | integer | 是 | 用户ID | 正整数，用户必须是账本成员 |

**请求示例**
```http
DELETE /api/v1/shared/books/:bookId/members/:userId
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| book_id | integer | 账本ID |
| user_id | integer | 被移除的用户ID |

**成功响应**
```json
{
  "code": 200,
  "message": "成员移除成功",
  "data": {
    "book_id": 1,
    "user_id": 3
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足：您没有移除成员的权限 | 检查当前用户是否有移除成员的权限 |
| 404 | 账本不存在 | 检查bookId是否正确 |
| 404 | 用户不是该账本的成员 | 检查userId是否正确，或用户是否是账本成员 |

**本地缓存策略**：移除成功后，清除本地缓存中该成员的相关信息

### 获取邀请列表

**功能描述**：获取当前用户发送或收到的邀请列表

**请求方法**：GET
**URL路径**：/api/v1/shared/invitations
**权限要求**：登录用户
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| type | string | 否 | - | 邀请类型：sent（发送的邀请）, received（收到的邀请） | 枚举值 |
| status | string | 否 | - | 邀请状态：pending, accepted, rejected, expired | 枚举值 |
| page | integer | 否 | 1 | 页码 | 正整数 |
| page_size | integer | 否 | 20 | 每页条数 | 1-100之间的整数 |

**请求示例**
```http
GET /api/v1/shared/invitations?type=sent&status=pending&page=1&page_size=20
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| total | integer | 总记录数 |
| page | integer | 当前页码 |
| page_size | integer | 每页条数 |
| list | array | 邀请列表 |
| list[].invitation_id | string | 邀请ID |
| list[].book_id | integer | 账本ID |
| list[].book_name | string | 账本名称 |
| list[].sender_id | integer | 发送者ID |
| list[].sender_name | string | 发送者名称 |
| list[].recipient_email | string | 接收者邮箱 |
| list[].permission | string | 权限级别 | 枚举值：viewer, editor, manager, owner |
| list[].status | string | 邀请状态 | 枚举值：pending, accepted, rejected, expired |
| list[].message | string | 邀请消息 |
| list[].expires_at | string | 过期时间 |
| list[].created_at | string | 创建时间 |
| list[].updated_at | string | 更新时间 |

**成功响应**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total": 2,
    "page": 1,
    "page_size": 20,
    "list": [
      {
        "invitation_id": "inv_123",
        "book_id": 1,
        "book_name": "家庭账本",
        "sender_id": 1,
        "sender_name": "张三",
        "recipient_email": "family_member@example.com",
        "permission": "editor",
        "status": "pending",
        "message": "邀请您加入我们的家庭账本",
        "expires_at": "2023-01-12T12:30:00Z",
        "created_at": "2023-01-05T12:30:00Z",
        "updated_at": "2023-01-05T12:30:00Z"
      }
    ]
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |

**本地缓存策略**：获取成功后，本地缓存邀请列表，缓存时间30分钟

### 创建共享邀请

**功能描述**：创建共享账本邀请，生成邀请链接

**请求方法**：POST
**URL路径**：/api/v1/shared/invitations
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
| book_id | integer | 是 | - | 账本ID | 正整数，账本必须存在且用户有管理权限 |
| email | string | 是 | - | 被邀请人邮箱 | 有效的邮箱格式 |
| permission | string | 是 | - | 权限级别 | 枚举值：viewer, editor, manager, owner |
| message | string | 否 | "" | 邀请消息 | 0-200个字符 |
| expires_in | integer | 否 | 7 | 邀请有效期（天） | 1-30之间的整数 |

**请求示例**
```http
POST /api/v1/shared/invitations
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "email": "family_member@example.com",
  "permission": "editor",
  "message": "邀请您加入我们的家庭账本",
  "expires_in": 7
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| invitation_id | string | 邀请ID |
| invitation_url | string | 邀请链接 |
| book_id | integer | 账本ID |
| recipient_email | string | 被邀请人邮箱 |
| permission | string | 权限级别 |
| status | string | 邀请状态 | 枚举值：pending, accepted, rejected, expired |
| expires_at | string | 过期时间 |
| created_at | string | 创建时间 |

**成功响应**
```json
{
  "code": 200,
  "message": "邀请创建成功",
  "data": {
    "invitation_id": "inv_123",
    "invitation_url": "https://example.com/invite/inv_123",
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
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：邮箱格式不正确 | 检查email参数格式是否有效 |
| 400 | 请求参数错误：无效的权限级别 | 检查permission参数是否为有效值 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足：您没有创建邀请的权限 | 检查当前用户是否有创建邀请的权限 |
| 404 | 账本不存在 | 检查book_id是否正确 |

**本地缓存策略**：创建成功后，本地缓存邀请信息，缓存时间1小时

### 处理邀请

**功能描述**：接受或拒绝共享账本邀请

**请求方法**：PUT
**URL路径**：/api/v1/shared/invitations/:id
**权限要求**：登录用户
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 描述 | 验证规则 |
|--------|------|------|------|----------|
| id | string | 是 | 邀请ID | 字符串，邀请必须存在且用户是接收者 |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| action | string | 是 | - | 处理动作 | 枚举值：accept, reject |

**请求示例**
```http
PUT /api/v1/shared/invitations/:id
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "action": "accept"
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| invitation_id | string | 邀请ID |
| status | string | 处理后的邀请状态 | 枚举值：accepted, rejected |
| book_id | integer | 账本ID |
| book_name | string | 账本名称 |

**成功响应**
```json
{
  "code": 200,
  "message": "邀请处理成功",
  "data": {
    "invitation_id": "inv_123",
    "status": "accepted",
    "book_id": 1,
    "book_name": "家庭账本"
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：无效的处理动作 | 检查action参数是否为有效值 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足：您不是该邀请的接收者 | 检查当前用户是否是邀请的接收者 |
| 404 | 邀请不存在 | 检查id参数是否正确 |
| 2008 | 邀请已过期 | 邀请已过期，无法处理 |

**本地缓存策略**：处理成功后，更新本地缓存中的邀请状态，缓存时间30分钟

### 获取提醒列表

**功能描述**：获取指定账本的提醒列表，支持分页、筛选和排序

**请求方法**：GET
**URL路径**：/api/v1/reminders
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| book_id | integer | 是 | - | 账本ID | 正整数，必须存在且用户有访问权限 |
| type | string | 否 | - | 提醒类型：income, expense, fixed_income, fixed_expense | 枚举值 |
| status | string | 否 | - | 提醒状态：active, inactive | 枚举值 |
| recurrence | string | 否 | - | 重复频率：one_time, weekly, biweekly, monthly, yearly | 枚举值 |
| page | integer | 否 | 1 | 页码 | 正整数 |
| page_size | integer | 否 | 20 | 每页条数 | 1-100之间的整数 |
| sort_by | string | 否 | "created_at" | 排序字段 | 可选值：created_at, updated_at, title, next_occurrence |
| sort_order | string | 否 | "desc" | 排序顺序 | 可选值：asc, desc |

**请求示例**
```http
GET /api/v1/reminders?book_id=1&type=expense&status=active&page=1&page_size=20
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| total | integer | 总记录数 |
| page | integer | 当前页码 |
| page_size | integer | 每页条数 |
| list | array | 提醒列表 |
| list[].reminder_id | integer | 提醒ID |
| list[].book_id | integer | 账本ID |
| list[].title | string | 提醒标题 |
| list[].amount | number | 提醒金额 |
| list[].currency | string | 货币类型 |
| list[].type | string | 提醒类型 |
| list[].recurrence | string | 重复频率 |
| list[].next_occurrence | string | 下次提醒时间 |
| list[].is_active | boolean | 是否激活 |
| list[].created_at | string | 创建时间 |
| list[].updated_at | string | 更新时间 |

**成功响应**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total": 10,
    "page": 1,
    "page_size": 20,
    "list": [
      {
        "reminder_id": 1,
        "book_id": 1,
        "title": "房贷还款提醒",
        "amount": 8000,
        "currency": "CNY",
        "type": "fixed_expense",
        "recurrence": "monthly",
        "next_occurrence": "2023-02-15",
        "is_active": true,
        "created_at": "2023-01-05T12:30:00Z",
        "updated_at": "2023-01-05T12:30:00Z"
      }
    ]
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：页码必须为正整数 | 检查page参数是否为正整数 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有查看提醒的权限 |
| 404 | 账本不存在 | 检查book_id是否正确 |

**本地缓存策略**：获取成功后，本地缓存提醒列表，缓存时间1小时

### 创建提醒

**功能描述**：创建新的账单提醒，用于提醒用户即将到来的收支事项

**请求方法**：POST
**URL路径**：/api/v1/reminders
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
| book_id | integer | 是 | - | 账本ID | 正整数，必须存在且用户有访问权限 |
| title | string | 是 | - | 提醒标题 | 1-50个字符 |
| amount | number | 是 | - | 提醒金额 | 大于0，最多两位小数 |
| currency | string | 否 | "CNY" | 货币类型 | 有效的货币代码 |
| type | string | 是 | - | 提醒类型 | 枚举值：income, expense, fixed_income, fixed_expense |
| recurrence | string | 是 | - | 重复频率 | 枚举值：one_time, weekly, biweekly, monthly, yearly |
| recurrence_day | integer | 否 | - | 重复日 | 1-31之间的整数，仅当recurrence为monthly时必填 |
| start_date | string | 是 | - | 开始日期 | YYYY-MM-DD格式 |
| end_date | string | 否 | - | 结束日期 | YYYY-MM-DD格式，必须大于等于start_date |
| reminder_days | array | 否 | [1] | 提前提醒天数 | 数组元素为0-30之间的整数 |
| account_id | integer | 是 | - | 账户ID | 正整数，账户必须存在且属于该账本 |
| category_id | integer | 是 | - | 分类ID | 正整数，分类必须存在且属于该账本 |
| auto_create | boolean | 否 | false | 是否自动创建交易记录 | 布尔值 |
| is_active | boolean | 否 | true | 是否激活 | 布尔值 |

**请求示例**
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

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| reminder_id | integer | 提醒ID |
| book_id | integer | 账本ID |
| title | string | 提醒标题 |
| amount | number | 提醒金额 |
| currency | string | 货币类型 |
| type | string | 提醒类型 |
| recurrence | string | 重复频率 |
| recurrence_day | integer | 重复日 |
| start_date | string | 开始日期 |
| end_date | string | 结束日期 |
| reminder_days | array | 提前提醒天数 |
| account_id | integer | 账户ID |
| category_id | integer | 分类ID |
| auto_create | boolean | 是否自动创建交易记录 |
| is_active | boolean | 是否激活 |
| next_occurrence | string | 下次提醒时间 |
| created_at | string | 创建时间 |
| updated_at | string | 更新时间 |

**成功响应**
```json
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "reminder_id": 1,
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
    "is_active": true,
    "next_occurrence": "2023-02-15",
    "created_at": "2023-01-05T12:30:00Z",
    "updated_at": "2023-01-05T12:30:00Z"
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：结束日期必须大于等于开始日期 | 检查end_date是否大于等于start_date |
| 400 | 请求参数错误：金额必须大于0 | 检查amount参数是否大于0 |
| 400 | 请求参数错误：无效的重复频率 | 检查recurrence参数是否为有效值 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有创建提醒的权限 |
| 404 | 账本不存在 | 检查book_id是否正确 |
| 404 | 账户不存在 | 检查account_id是否正确 |
| 404 | 分类不存在 | 检查category_id是否正确 |
| 2003 | 提醒创建失败 | 检查提醒参数是否正确，特别是日期范围、重复规则等 |
| 2009 | 提醒频率过高 | 提醒频率设置过高，请调整提醒规则 |

**本地缓存策略**：创建成功后，本地缓存该提醒信息，缓存时间24小时

### 更新提醒

**功能描述**：更新指定提醒的信息，包括标题、金额、重复规则等

**请求方法**：PUT
**URL路径**：/api/v1/reminders/:id
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 描述 | 验证规则 |
|--------|------|------|------|----------|
| id | integer | 是 | 提醒ID | 正整数，提醒必须存在且用户有访问权限 |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| title | string | 否 | - | 提醒标题 | 1-50个字符 |
| amount | number | 否 | - | 提醒金额 | 大于0，最多两位小数 |
| currency | string | 否 | - | 货币类型 | 有效的货币代码 |
| type | string | 否 | - | 提醒类型 | 枚举值：income, expense, fixed_income, fixed_expense |
| recurrence | string | 否 | - | 重复频率 | 枚举值：one_time, weekly, biweekly, monthly, yearly |
| recurrence_day | integer | 否 | - | 重复日 | 1-31之间的整数，仅当recurrence为monthly时必填 |
| start_date | string | 否 | - | 开始日期 | YYYY-MM-DD格式 |
| end_date | string | 否 | - | 结束日期 | YYYY-MM-DD格式，必须大于等于start_date |
| reminder_days | array | 否 | - | 提前提醒天数 | 数组元素为0-30之间的整数 |
| account_id | integer | 否 | - | 账户ID | 正整数，账户必须存在且属于该账本 |
| category_id | integer | 否 | - | 分类ID | 正整数，分类必须存在且属于该账本 |
| auto_create | boolean | 否 | - | 是否自动创建交易记录 | 布尔值 |
| is_active | boolean | 否 | - | 是否激活 | 布尔值 |

**请求示例**
```http
PUT /api/v1/reminders/:id
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "title": "更新后的房贷还款提醒",
  "amount": 8500,
  "reminder_days": [7, 3, 1, 0]
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| reminder_id | integer | 提醒ID |
| book_id | integer | 账本ID |
| title | string | 提醒标题 |
| amount | number | 提醒金额 |
| currency | string | 货币类型 |
| type | string | 提醒类型 |
| recurrence | string | 重复频率 |
| recurrence_day | integer | 重复日 |
| start_date | string | 开始日期 |
| end_date | string | 结束日期 |
| reminder_days | array | 提前提醒天数 |
| account_id | integer | 账户ID |
| category_id | integer | 分类ID |
| auto_create | boolean | 是否自动创建交易记录 |
| is_active | boolean | 是否激活 |
| next_occurrence | string | 下次提醒时间 |
| updated_at | string | 更新时间 |

**成功响应**
```json
{
  "code": 200,
  "message": "更新成功",
  "data": {
    "reminder_id": 1,
    "book_id": 1,
    "title": "更新后的房贷还款提醒",
    "amount": 8500,
    "currency": "CNY",
    "type": "fixed_expense",
    "recurrence": "monthly",
    "recurrence_day": 15,
    "start_date": "2023-01-15",
    "end_date": "2030-12-15",
    "reminder_days": [7, 3, 1, 0],
    "account_id": 1,
    "category_id": 2,
    "auto_create": true,
    "is_active": true,
    "next_occurrence": "2023-02-15",
    "updated_at": "2023-01-10T14:30:00Z"
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：结束日期必须大于等于开始日期 | 检查end_date是否大于等于start_date |
| 400 | 请求参数错误：金额必须大于0 | 检查amount参数是否大于0 |
| 400 | 请求参数错误：无效的重复频率 | 检查recurrence参数是否为有效值 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有更新提醒的权限 |
| 404 | 提醒不存在 | 检查id参数是否正确 |
| 404 | 账户不存在 | 检查account_id是否正确 |
| 404 | 分类不存在 | 检查category_id是否正确 |
| 2003 | 提醒更新失败 | 检查提醒参数是否正确，特别是日期范围、重复规则等 |

**本地缓存策略**：更新成功后，更新本地缓存，缓存时间24小时

### 删除提醒

**功能描述**：删除指定的提醒

**请求方法**：DELETE
**URL路径**：/api/v1/reminders/:id
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 描述 | 验证规则 |
|--------|------|------|------|----------|
| id | integer | 是 | 提醒ID | 正整数，提醒必须存在且用户有访问权限 |

**请求示例**
```http
DELETE /api/v1/reminders/:id
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| reminder_id | integer | 已删除的提醒ID |

**成功响应**
```json
{
  "code": 200,
  "message": "删除成功",
  "data": {
    "reminder_id": 1
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有删除提醒的权限 |
| 404 | 提醒不存在 | 检查id参数是否正确 |

**本地缓存策略**：删除成功后，清除本地缓存中该提醒的相关信息

### 激活停用提醒

**功能描述**：激活或停用指定的提醒

**请求方法**：PUT
**URL路径**：/api/v1/reminders/:id/activate
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 描述 | 验证规则 |
|--------|------|------|------|----------|
| id | integer | 是 | 提醒ID | 正整数，提醒必须存在且用户有访问权限 |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| is_active | boolean | 是 | - | 是否激活 | 布尔值 |

**请求示例**
```http
PUT /api/v1/reminders/:id/activate
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "is_active": false
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| reminder_id | integer | 提醒ID |
| is_active | boolean | 提醒状态 |
| updated_at | string | 更新时间 |

**成功响应**
```json
{
  "code": 200,
  "message": "操作成功",
  "data": {
    "reminder_id": 1,
    "is_active": false,
    "updated_at": "2023-01-10T14:30:00Z"
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：is_active必须为布尔值 | 检查is_active参数是否为布尔值 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有操作该提醒的权限 |
| 404 | 提醒不存在 | 检查id参数是否正确 |

**本地缓存策略**：操作成功后，更新本地缓存中的提醒状态，缓存时间24小时

### 获取即将到来的提醒

**功能描述**：获取即将到来的提醒列表，支持按时间范围筛选

**请求方法**：GET
**URL路径**：/api/v1/reminders/upcoming
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| book_id | integer | 是 | - | 账本ID | 正整数，必须存在且用户有访问权限 |
| days | integer | 否 | 7 | 未来天数 | 1-30之间的整数 |
| type | string | 否 | - | 提醒类型：income, expense, fixed_income, fixed_expense | 枚举值 |

**请求示例**
```http
GET /api/v1/reminders/upcoming?book_id=1&days=7&type=expense
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| total | integer | 总记录数 |
| list | array | 即将到来的提醒列表 |
| list[].reminder_id | integer | 提醒ID |
| list[].book_id | integer | 账本ID |
| list[].title | string | 提醒标题 |
| list[].amount | number | 提醒金额 |
| list[].currency | string | 货币类型 |
| list[].type | string | 提醒类型 |
| list[].occurrence_date | string | 提醒日期 |
| list[].is_active | boolean | 是否激活 |
| list[].days_until | integer | 距离提醒的天数 |

**成功响应**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total": 3,
    "list": [
      {
        "reminder_id": 1,
        "book_id": 1,
        "title": "房贷还款提醒",
        "amount": 8000,
        "currency": "CNY",
        "type": "fixed_expense",
        "occurrence_date": "2023-02-15",
        "is_active": true,
        "days_until": 5
      }
    ]
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：days必须在1-30之间 | 检查days参数是否在1-30之间 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有查看提醒的权限 |
| 404 | 账本不存在 | 检查book_id是否正确 |

**本地缓存策略**：获取成功后，本地缓存即将到来的提醒列表，缓存时间30分钟

### 预测交易分类

**功能描述**：根据交易信息智能预测分类，提高记账效率

**请求方法**：POST
**URL路径**：/api/v1/smart/category-predict
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
| book_id | integer | 是 | - | 账本ID | 正整数，必须存在且用户有访问权限 |
| transaction | object | 是 | - | 交易信息 | 包含amount, description, account_id, date, currency字段 |
| transaction.amount | number | 是 | - | 交易金额 | 大于0，最多两位小数 |
| transaction.description | string | 是 | - | 交易描述 | 1-200个字符 |
| transaction.account_id | integer | 是 | - | 账户ID | 正整数，账户必须存在且属于该账本 |
| transaction.date | string | 是 | - | 交易日期 | YYYY-MM-DD格式 |
| transaction.currency | string | 否 | "CNY" | 货币类型 | 有效的货币代码 |
| include_probability | boolean | 否 | false | 是否包含概率 | 布尔值 |

**请求示例**
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

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| predicted_categories | array | 预测的分类列表，按概率降序排列 |
| predicted_categories[].category_id | integer | 分类ID |
| predicted_categories[].name | string | 分类名称 |
| predicted_categories[].probability | number | 预测概率（仅当include_probability为true时返回） |
| confidence_level | string | 置信度：high, medium, low |
| model_version | string | 模型版本 |

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
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：交易信息不完整 | 检查transaction参数是否包含所有必填字段 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有使用智能分类功能的权限 |
| 404 | 账本不存在 | 检查book_id是否正确 |
| 404 | 账户不存在 | 检查transaction.account_id是否正确 |
| 2005 | 智能分类模型错误：模型加载失败 | 模型加载或预测失败，请稍后重试 |

**本地缓存策略**：预测结果不缓存，每次请求重新计算

### 检测周期性交易

**功能描述**：自动检测用户的周期性交易，帮助用户发现固定支出和收入

**请求方法**：GET
**URL路径**：/api/v1/smart/recurring-detection
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| book_id | integer | 是 | - | 账本ID | 正整数，必须存在且用户有访问权限 |
| period | string | 否 | "3m" | 检测周期 | 格式：数字+m（月），如3m, 6m, 12m |
| min_occurrences | integer | 否 | 3 | 最小出现次数 | 2-12之间的整数 |
| type | string | 否 | - | 交易类型：income, expense | 枚举值 |

**请求示例**
```http
GET /api/v1/smart/recurring-detection?book_id=1&period=6m&min_occurrences=3&type=expense
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| recurring_transactions | array | 检测到的周期性交易列表 |
| recurring_transactions[].id | string | 周期性交易ID |
| recurring_transactions[].description | string | 交易描述 |
| recurring_transactions[].amount | number | 交易金额 |
| recurring_transactions[].type | string | 交易类型 |
| recurring_transactions[].frequency | string | 频率：daily, weekly, monthly, yearly |
| recurring_transactions[].next_occurrence | string | 预计下次发生时间 |
| recurring_transactions[].occurrences | array | 历史发生记录 |
| recurring_transactions[].occurrences[].transaction_id | integer | 交易ID |
| recurring_transactions[].occurrences[].date | string | 交易日期 |
| recurring_transactions[].confidence_score | number | 置信度分数 |

**成功响应**
```json
{
  "code": 200,
  "message": "检测成功",
  "data": {
    "recurring_transactions": [
      {
        "id": "recur_123",
        "description": "房贷还款",
        "amount": 8000.00,
        "type": "expense",
        "frequency": "monthly",
        "next_occurrence": "2023-02-15",
        "occurrences": [
          {
            "transaction_id": 1001,
            "date": "2023-01-15"
          },
          {
            "transaction_id": 987,
            "date": "2022-12-15"
          },
          {
            "transaction_id": 956,
            "date": "2022-11-15"
          }
        ],
        "confidence_score": 0.98
      }
    ]
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：无效的检测周期 | 检查period参数格式是否正确 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有使用智能分类功能的权限 |
| 404 | 账本不存在 | 检查book_id是否正确 |
| 2005 | 智能检测模型错误 | 模型检测失败，请稍后重试 |

**本地缓存策略**：检测结果本地缓存，缓存时间1小时

### 获取消费模式分析

**功能描述**：分析用户的消费模式，提供消费习惯洞察

**请求方法**：GET
**URL路径**：/api/v1/smart/spending-pattern
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| book_id | integer | 是 | - | 账本ID | 正整数，必须存在且用户有访问权限 |
| period | string | 否 | "3m" | 分析周期 | 格式：数字+m（月），如3m, 6m, 12m |
| category_id | integer | 否 | - | 分类ID | 正整数，分类必须存在 |

**请求示例**
```http
GET /api/v1/smart/spending-pattern?book_id=1&period=6m&category_id=5
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| spending_patterns | array | 消费模式列表 |
| spending_patterns[].type | string | 模式类型：daily, weekly, monthly, category |
| spending_patterns[].description | string | 模式描述 |
| spending_patterns[].data | object | 模式数据 |
| spending_patterns[].data.avg_amount | number | 平均金额 |
| spending_patterns[].data.max_amount | number | 最大金额 |
| spending_patterns[].data.min_amount | number | 最小金额 |
| spending_patterns[].data.total_amount | number | 总金额 |
| spending_patterns[].data.frequency | object | 频率分布 |
| insights | array | 消费洞察建议 |
| insights[].type | string | 洞察类型 |
| insights[].description | string | 洞察描述 |
| insights[].suggestion | string | 改进建议 |

**成功响应**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "spending_patterns": [
      {
        "type": "monthly",
        "description": "月度消费趋势",
        "data": {
          "avg_amount": 8500.00,
          "max_amount": 10000.00,
          "min_amount": 7000.00,
          "total_amount": 51000.00,
          "frequency": {
            "weekday": 6500.00,
            "weekend": 2000.00
          }
        }
      }
    ],
    "insights": [
      {
        "type": "high_spending",
        "description": "您的餐饮支出占总支出的30%，高于平均水平",
        "suggestion": "建议设置餐饮预算，控制餐饮消费"
      }
    ]
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：无效的分析周期 | 检查period参数格式是否正确 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有使用智能分类功能的权限 |
| 404 | 账本不存在 | 检查book_id是否正确 |
| 404 | 分类不存在 | 检查category_id是否正确 |
| 2005 | 消费模式分析失败 | 分析失败，请稍后重试 |

**本地缓存策略**：分析结果本地缓存，缓存时间24小时

### 异常消费检测

**功能描述**：检测用户的异常消费行为，提醒用户可能的异常支出

**请求方法**：GET
**URL路径**：/api/v1/smart/anomaly-detection
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| book_id | integer | 是 | - | 账本ID | 正整数，必须存在且用户有访问权限 |
| period | string | 否 | "3m" | 检测周期 | 格式：数字+m（月），如3m, 6m, 12m |
| sensitivity | string | 否 | "medium" | 敏感度：low, medium, high | 枚举值 |

**请求示例**
```http
GET /api/v1/smart/anomaly-detection?book_id=1&period=3m&sensitivity=high
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| anomalies | array | 检测到的异常消费列表 |
| anomalies[].transaction_id | integer | 交易ID |
| anomalies[].description | string | 交易描述 |
| anomalies[].amount | number | 交易金额 |
| anomalies[].date | string | 交易日期 |
| anomalies[].category_name | string | 分类名称 |
| anomalies[].anomaly_score | number | 异常分数 |
| anomalies[].reason | string | 异常原因 |
| anomalies[].suggestion | string | 建议 |

**成功响应**
```json
{
  "code": 200,
  "message": "检测成功",
  "data": {
    "anomalies": [
      {
        "transaction_id": 1001,
        "description": "奢侈品消费",
        "amount": 5000.00,
        "date": "2023-01-15",
        "category_name": "购物",
        "anomaly_score": 0.95,
        "reason": "金额远高于该分类平均支出",
        "suggestion": "确认是否为正常消费"
      }
    ]
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：无效的检测周期 | 检查period参数格式是否正确 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有使用智能分类功能的权限 |
| 404 | 账本不存在 | 检查book_id是否正确 |
| 2005 | 异常检测失败 | 检测失败，请稍后重试 |

**本地缓存策略**：检测结果本地缓存，缓存时间12小时

### 获取汇率列表

**功能描述**：获取当前支持的所有货币汇率列表

**请求方法**：GET
**URL路径**：/api/v1/currency/rates
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| base_currency | string | 否 | "CNY" | 基准货币 | 有效的货币代码 |
| date | string | 否 | 当前日期 | 汇率日期 | YYYY-MM-DD格式 |

**请求示例**
```http
GET /api/v1/currency/rates?base_currency=CNY&date=2023-01-05
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| base_currency | string | 基准货币 |
| date | string | 汇率日期 |
| last_updated | string | 最后更新时间 |
| rates | object | 汇率列表，key为货币代码，value为汇率 |

**成功响应**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "base_currency": "CNY",
    "date": "2023-01-05",
    "last_updated": "2023-01-05T08:00:00Z",
    "rates": {
      "USD": 0.1469,
      "EUR": 0.1325,
      "JPY": 15.789,
      "GBP": 0.1156
    }
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：无效的货币代码 | 检查base_currency是否为有效的货币代码 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有使用汇率功能的权限 |
| 2004 | 汇率获取失败 | 检查网络连接，或稍后重试 |
| 2010 | 汇率数据过期 | 汇率数据已过期，请调用刷新汇率接口更新数据 |

**本地缓存策略**：汇率数据本地缓存，缓存时间1小时

### 货币转换

**功能描述**：进行不同货币之间的金额转换，支持历史汇率查询

**请求方法**：POST
**URL路径**：/api/v1/currency/convert
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
| amount | number | 是 | - | 转换金额 | 大于0，最多两位小数 |
| from_currency | string | 是 | - | 源货币 | 有效的货币代码 |
| to_currency | string | 是 | - | 目标货币 | 有效的货币代码 |
| date | string | 否 | 当前日期 | 汇率日期 | YYYY-MM-DD格式 |

**请求示例**
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

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| amount | number | 原始金额 |
| from_currency | string | 源货币 |
| to_currency | string | 目标货币 |
| converted_amount | number | 转换后金额 |
| exchange_rate | number | 汇率 |
| rate_date | string | 汇率日期 |
| rate_source | string | 汇率来源 |
| last_updated | string | 汇率最后更新时间 |

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
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：金额必须大于0 | 检查amount参数是否大于0 |
| 400 | 请求参数错误：无效的货币代码 | 检查from_currency和to_currency是否为有效的货币代码 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有使用货币转换功能的权限 |
| 2004 | 汇率获取失败 | 检查网络连接，或稍后重试 |
| 2010 | 汇率数据过期：请刷新汇率数据 | 汇率数据已过期，请调用刷新汇率接口更新数据 |

**本地缓存策略**：汇率数据本地缓存，缓存时间1小时

### 获取支持的货币列表

**功能描述**：获取系统支持的所有货币列表

**请求方法**：GET
**URL路径**：/api/v1/currency/supported
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**请求示例**
```http
GET /api/v1/currency/supported
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| currencies | array | 支持的货币列表 |
| currencies[].code | string | 货币代码 |
| currencies[].name | string | 货币名称 |
| currencies[].symbol | string | 货币符号 |
| currencies[].decimal_places | integer | 小数位数 |

**成功响应**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "currencies": [
      {
        "code": "CNY",
        "name": "人民币",
        "symbol": "¥",
        "decimal_places": 2
      },
      {
        "code": "USD",
        "name": "美元",
        "symbol": "$",
        "decimal_places": 2
      },
      {
        "code": "EUR",
        "name": "欧元",
        "symbol": "€",
        "decimal_places": 2
      }
    ]
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有使用汇率功能的权限 |

**本地缓存策略**：货币列表本地缓存，缓存时间7天

### 刷新汇率

**功能描述**：手动刷新汇率数据，更新本地缓存

**请求方法**：POST
**URL路径**：/api/v1/currency/refresh
**权限要求**：账本成员
**限流策略**：每分钟10次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| base_currency | string | 否 | "CNY" | 基准货币 | 有效的货币代码 |

**请求示例**
```http
POST /api/v1/currency/refresh
Authorization: Bearer jwt_token_string

{
  "base_currency": "CNY"
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| base_currency | string | 基准货币 |
| last_updated | string | 刷新时间 |
| updated_currencies | integer | 更新的货币数量 |

**成功响应**
```json
{
  "code": 200,
  "message": "刷新成功",
  "data": {
    "base_currency": "CNY",
    "last_updated": "2023-01-05T14:30:00Z",
    "updated_currencies": 156
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误：无效的货币代码 | 检查base_currency是否为有效的货币代码 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有使用汇率功能的权限 |
| 2004 | 汇率刷新失败 | 检查网络连接，或稍后重试 |

**本地缓存策略**：刷新成功后，更新本地缓存，缓存时间1小时

## 数据模型

### 预算模型

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

### 标签模型

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

### 共享成员模型

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

### 提醒模型

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

### 汇率模型

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

## 错误码说明

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
7. **本地缓存策略**：所有API响应可根据需要进行本地缓存，缓存时间根据数据类型和更新频率确定
8. **数据验证**：客户端应在发送请求前对数据进行验证，减少无效请求
9. **错误重试**：对于500错误，客户端可进行适当的重试机制
10. **日志记录**：客户端应记录API请求和响应日志，便于问题排查