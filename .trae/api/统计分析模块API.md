## 模块概述

统计分析模块提供多维度的财务数据分析功能，通过图表和报表帮助用户了解财务状况，辅助决策。本模块提供了完整的数据看板、收支分析、类别统计、趋势分析等功能的API接口。

### 设计原则

1. **RESTful规范**：所有API接口严格遵循RESTful设计规范，使用标准HTTP方法和资源导向的URL路径
2. **本地缓存优先**：暂不集成远程缓存机制，如需使用缓存功能，严格限定为本地缓存方案
3. **完整的请求验证**：所有API接口包含严格的请求参数验证，确保数据完整性和安全性
4. **统一的错误处理**：采用统一的错误响应格式和错误码体系，便于客户端处理
5. **清晰的响应格式**：所有API响应采用标准JSON格式，包含明确的数据结构和状态信息
6. **可扩展性**：API设计考虑未来功能扩展，预留合理的扩展点

### 适用场景

- 用户需要查看财务概览数据
- 用户需要分析收支对比和趋势
- 用户需要查看分类统计报表
- 用户需要了解资金流向
- 用户需要导出统计报表
- 用户需要创建自定义报表

## 接口清单

<!-- tabs:start -->
<!-- tab:数据看板 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/dashboard/summary`](#获取财务概览数据) | `GET` | 获取财务概览数据 |
| [`/api/v1/dashboard/quick-stats`](#获取快速统计数据) | `GET` | 获取快速统计数据 |
| [`/api/v1/dashboard/recent-transactions`](#获取最近收支记录) | `GET` | 获取最近收支记录 |
| [`/api/v1/dashboard/budget-progress`](#获取预算执行进度) | `GET` | 获取预算执行进度 |

<!-- tab:收支分析 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/analysis/income-expense`](#获取收支对比分析) | `GET` | 获取收支对比分析 |
| [`/api/v1/analysis/trend`](#获取收支趋势分析) | `GET` | 获取收支趋势分析 |
| [`/api/v1/analysis/flow`](#获取资金流向分析) | `GET` | 获取资金流向分析 |

<!-- tab:统计报表 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/reports/category`](#获取分类统计报表) | `GET` | 获取分类统计报表 |
| [`/api/v1/reports/category/trend`](#获取分类趋势报表) | `GET` | 获取分类趋势报表 |
| [`/api/v1/reports/account`](#获取账户收支报表) | `GET` | 获取账户收支报表 |
| [`/api/v1/reports/account/balance`](#获取账户余额报表) | `GET` | 获取账户余额报表 |
| [`/api/v1/reports/member`](#获取成员收支报表) | `GET` | 获取成员收支报表 |
| [`/api/v1/reports/member/contribution`](#获取成员贡献度分析) | `GET` | 获取成员贡献度分析 |
| [`/api/v1/reports/budget`](#获取预算执行报表) | `GET` | 获取预算执行报表 |
| [`/api/v1/reports/budget/alert`](#获取预算超支提醒) | `GET` | 获取预算超支提醒 |

<!-- tab:图表数据 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/charts/line`](#获取折线图数据) | `GET` | 获取折线图数据 |
| [`/api/v1/charts/pie`](#获取饼图数据) | `GET` | 获取饼图数据 |
| [`/api/v1/charts/bar`](#获取柱状图数据) | `GET` | 获取柱状图数据 |
| [`/api/v1/charts/radar`](#获取雷达图数据) | `GET` | 获取雷达图数据 |

<!-- tab:自定义报表 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/reports/custom`](#创建自定义报表) | `POST` | 创建自定义报表 |
| [`/api/v1/reports/custom`](#获取自定义报表列表) | `GET` | 获取自定义报表列表 |
| [`/api/v1/reports/custom/:id`](#获取自定义报表数据) | `GET` | 获取自定义报表数据 |
| [`/api/v1/reports/custom/:id`](#更新自定义报表) | `PUT` | 更新自定义报表 |
| [`/api/v1/reports/custom/:id`](#删除自定义报表) | `DELETE` | 删除自定义报表 |

<!-- tab:报表导出 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/export/excel`](#导出Excel报表) | `GET` | 导出Excel报表 |
| [`/api/v1/export/csv`](#导出CSV报表) | `GET` | 导出CSV报表 |
| [`/api/v1/export/pdf`](#导出PDF报表) | `GET` | 导出PDF报表 |

<!-- tabs:end -->

## 详细接口说明

### 获取财务概览数据

**功能描述**：获取指定账本的财务概览数据，包括总收入、总支出、余额、同比增长率等

**请求方法**：GET
**URL路径**：/api/v1/dashboard/summary
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
| period | string | 否 | month | 统计周期：day, week, month, year |
| date | string | 否 | 当前日期 | 统计日期，格式：YYYY-MM-DD或YYYY-MM |

**请求示例**
```http
GET /api/v1/dashboard/summary?book_id=1&period=month&date=2023-01
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 财务概览数据 |
| data.total_income | number | 总收入 |
| data.total_expense | number | 总支出 |
| data.balance | number | 结余 |
| data.income_rate | number | 收入同比增长率 |
| data.expense_rate | number | 支出同比增长率 |
| data.top_expense_category | object | 最高支出分类 |
| data.top_expense_category.id | integer | 分类ID |
| data.top_expense_category.name | string | 分类名称 |
| data.top_expense_category.amount | number | 支出金额 |
| data.top_expense_category.percentage | number | 占总支出比例 |
| data.budget_status | object | 预算执行情况 |
| data.budget_status.total | number | 总预算 |
| data.budget_status.used | number | 已使用预算 |
| data.budget_status.percentage | number | 预算使用比例 |
| data.budget_status.alert | boolean | 是否触发预算提醒 |
| data.account_summary | object | 账户汇总信息 |
| data.account_summary.total_balance | number | 总余额 |
| data.account_summary.account_count | integer | 账户数量 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total_income": 8000.00,
    "total_expense": 3500.00,
    "balance": 4500.00,
    "income_rate": 10.5,  // 同比增长率
    "expense_rate": -5.2,  // 同比增长率
    "top_expense_category": {
      "id": 5,
      "name": "餐饮",
      "amount": 1200.00,
      "percentage": 34.3
    },
    "budget_status": {
      "total": 5000.00,
      "used": 3500.00,
      "percentage": 70.0,
      "alert": false
    },
    "account_summary": {
      "total_balance": 25000.00,
      "account_count": 5
    }
  }
}
```

### 获取快速统计数据

**功能描述**：获取指定账本的快速统计数据，包括收支笔数、平均收支等

**请求方法**：GET
**URL路径**：/api/v1/dashboard/quick-stats
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
| period | string | 否 | month | 统计周期：day, week, month, year |
| date | string | 否 | 当前日期 | 统计日期，格式：YYYY-MM-DD或YYYY-MM |

**请求示例**
```http
GET /api/v1/dashboard/quick-stats?book_id=1&period=month&date=2023-01
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 快速统计数据 |
| data.income_count | integer | 收入笔数 |
| data.expense_count | integer | 支出笔数 |
| data.avg_income | number | 平均收入 |
| data.avg_expense | number | 平均支出 |
| data.total_transactions | integer | 总交易笔数 |
| data.highest_income | number | 最高单笔收入 |
| data.highest_expense | number | 最高单笔支出 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "income_count": 10,
    "expense_count": 30,
    "avg_income": 800.00,
    "avg_expense": 116.67,
    "total_transactions": 40,
    "highest_income": 5000.00,
    "highest_expense": 800.00
  }
}
```

### 获取最近收支记录

**功能描述**：获取指定账本的最近收支记录，支持分页和筛选

**请求方法**：GET
**URL路径**：/api/v1/dashboard/recent-transactions
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
| limit | integer | 否 | 10 | 每页记录数，最大50 |
| offset | integer | 否 | 0 | 偏移量 |
| type | string | 否 | - | 交易类型：income（收入）, expense（支出） |

**请求示例**
```http
GET /api/v1/dashboard/recent-transactions?book_id=1&limit=10&offset=0
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 最近收支记录数据 |
| data.list | array | 收支记录列表 |
| data.list[].id | integer | 记录ID |
| data.list[].amount | number | 金额 |
| data.list[].type | string | 交易类型 |
| data.list[].category_name | string | 分类名称 |
| data.list[].account_name | string | 账户名称 |
| data.list[].remark | string | 备注 |
| data.list[].created_at | string | 创建时间 |
| data.total | integer | 总记录数 |
| data.limit | integer | 每页记录数 |
| data.offset | integer | 偏移量 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "list": [
      {
        "id": 1001,
        "amount": 100.00,
        "type": "expense",
        "category_name": "餐饮",
        "account_name": "支付宝",
        "remark": "午餐",
        "created_at": "2023-01-31T12:00:00Z"
      },
      {
        "id": 1002,
        "amount": 5000.00,
        "type": "income",
        "category_name": "工资",
        "account_name": "我的工资卡",
        "remark": "1月工资",
        "created_at": "2023-01-30T10:00:00Z"
      }
    ],
    "total": 40,
    "limit": 10,
    "offset": 0
  }
}
```

### 获取预算执行进度

**功能描述**：获取指定账本的预算执行进度数据

**请求方法**：GET
**URL路径**：/api/v1/dashboard/budget-progress
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
| period | string | 否 | month | 统计周期：month, quarter, year |
| date | string | 否 | 当前日期 | 统计日期，格式：YYYY-MM |

**请求示例**
```http
GET /api/v1/dashboard/budget-progress?book_id=1&period=month&date=2023-01
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 预算执行进度数据 |
| data.total_budget | number | 总预算 |
| data.used_budget | number | 已使用预算 |
| data.remaining_budget | number | 剩余预算 |
| data.usage_rate | number | 预算使用率 |
| data.categories | array | 分类预算执行情况 |
| data.categories[].category_id | integer | 分类ID |
| data.categories[].category_name | string | 分类名称 |
| data.categories[].budget | number | 分类预算 |
| data.categories[].used | number | 已使用金额 |
| data.categories[].usage_rate | number | 分类预算使用率 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total_budget": 5000.00,
    "used_budget": 3500.00,
    "remaining_budget": 1500.00,
    "usage_rate": 70.0,
    "categories": [
      {
        "category_id": 5,
        "category_name": "餐饮",
        "budget": 1500.00,
        "used": 1200.00,
        "usage_rate": 80.0
      },
      {
        "category_id": 10,
        "category_name": "住房",
        "budget": 1000.00,
        "used": 1000.00,
        "usage_rate": 100.0
      }
    ]
  }
}
```

### 获取收支对比分析

**功能描述**：获取指定账本在特定时间范围内的收支对比数据，支持多维度对比分析

**请求方法**：GET
**URL路径**：/api/v1/analysis/income-expense
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
| period | string | 否 | "month" | 统计周期：day, week, month, year |
| start_date | string | 是 | - | 开始日期，格式：YYYY-MM-DD或YYYY-MM |
| end_date | string | 是 | - | 结束日期，格式：YYYY-MM-DD或YYYY-MM |
| compare_period | string | 否 | "previous" | 对比周期：previous（上一周期）, same（同比） |

**请求示例**
```http
GET /api/v1/analysis/income-expense?book_id=1&period=month&start_date=2023-01&end_date=2023-12&compare_period=previous
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 收支对比数据 |
| data.current | object | 当前周期数据 |
| data.current.income | number | 当前周期收入 |
| data.current.expense | number | 当前周期支出 |
| data.current.balance | number | 当前周期结余 |
| data.compare | object | 对比周期数据 |
| data.compare.income | number | 对比周期收入 |
| data.compare.expense | number | 对比周期支出 |
| data.compare.balance | number | 对比周期结余 |
| data.change | object | 变化率数据 |
| data.change.income_rate | number | 收入变化率 |
| data.change.expense_rate | number | 支出变化率 |
| data.change.balance_rate | number | 结余变化率 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "current": {
      "income": 8000.00,
      "expense": 3500.00,
      "balance": 4500.00
    },
    "compare": {
      "income": 7200.00,
      "expense": 3800.00,
      "balance": 3400.00
    },
    "change": {
      "income_rate": 11.1,
      "expense_rate": -7.9,
      "balance_rate": 32.4
    }
  }
}
```

### 获取收支趋势分析

**功能描述**：获取指定账本在特定时间范围内的收支趋势数据，支持按日、周、月、年统计

**请求方法**：GET
**URL路径**：/api/v1/analysis/trend
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
| type | string | 否 | "both" | 数据类型：income（收入）, expense（支出）, both（收支） |
| period | string | 否 | "month" | 统计周期：day, week, month, year |
| start_date | string | 是 | - | 开始日期，格式：YYYY-MM-DD或YYYY-MM |
| end_date | string | 是 | - | 结束日期，格式：YYYY-MM-DD或YYYY-MM |

**请求示例**
```http
GET /api/v1/analysis/trend?book_id=1&type=both&period=month&start_date=2023-01&end_date=2023-12
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 收支趋势数据 |
| data.period | string | 统计周期 |
| data.start_date | string | 开始日期 |
| data.end_date | string | 结束日期 |
| data.data | array | 趋势数据列表 |
| data.data[].date | string | 统计日期 |
| data.data[].income | number | 收入金额 |
| data.data[].expense | number | 支出金额 |
| data.data[].balance | number | 结余金额 |
| data.summary | object | 汇总数据 |
| data.summary.total_income | number | 总收入 |
| data.summary.total_expense | number | 总支出 |
| data.summary.total_balance | number | 总结余 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "period": "month",
    "start_date": "2023-01",
    "end_date": "2023-12",
    "data": [
      {
        "date": "2023-01",
        "income": 8000.00,
        "expense": 3500.00,
        "balance": 4500.00
      },
      {
        "date": "2023-02",
        "income": 8200.00,
        "expense": 3800.00,
        "balance": 4400.00
      },
      // 更多月份数据...
    ],
    "summary": {
      "total_income": 96000.00,
      "total_expense": 42000.00,
      "total_balance": 54000.00
    }
  }
}
```

### 获取资金流向分析

**功能描述**：获取指定账本在特定时间范围内的资金流向分析数据，展示资金的来源和去向

**请求方法**：GET
**URL路径**：/api/v1/analysis/flow
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
| period | string | 否 | "month" | 统计周期：day, week, month, year |
| start_date | string | 是 | - | 开始日期，格式：YYYY-MM-DD或YYYY-MM |
| end_date | string | 是 | - | 结束日期，格式：YYYY-MM-DD或YYYY-MM |
| min_amount | number | 否 | 0 | 最小金额过滤 |

**请求示例**
```http
GET /api/v1/analysis/flow?book_id=1&period=month&start_date=2023-01&end_date=2023-01&min_amount=100
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 资金流向数据 |
| data.income_flow | array | 收入流向列表 |
| data.income_flow[].category_id | integer | 分类ID |
| data.income_flow[].category_name | string | 分类名称 |
| data.income_flow[].amount | number | 金额 |
| data.income_flow[].percentage | number | 占比 |
| data.expense_flow | array | 支出流向列表 |
| data.expense_flow[].category_id | integer | 分类ID |
| data.expense_flow[].category_name | string | 分类名称 |
| data.expense_flow[].amount | number | 金额 |
| data.expense_flow[].percentage | number | 占比 |
| data.transfer_flow | array | 转账流向列表 |
| data.transfer_flow[].from_account_id | integer | 转出账户ID |
| data.transfer_flow[].from_account_name | string | 转出账户名称 |
| data.transfer_flow[].to_account_id | integer | 转入账户ID |
| data.transfer_flow[].to_account_name | string | 转入账户名称 |
| data.transfer_flow[].amount | number | 转账金额 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "income_flow": [
      {
        "category_id": 1,
        "category_name": "工资",
        "amount": 8000.00,
        "percentage": 100.0
      }
    ],
    "expense_flow": [
      {
        "category_id": 5,
        "category_name": "餐饮",
        "amount": 1200.00,
        "percentage": 34.3
      },
      {
        "category_id": 10,
        "category_name": "住房",
        "amount": 1000.00,
        "percentage": 28.6
      }
    ],
    "transfer_flow": [
      {
        "from_account_id": 1,
        "from_account_name": "工资卡",
        "to_account_id": 2,
        "to_account_name": "支付宝",
        "amount": 2000.00
      }
    ]
  }
}
```

### 获取分类统计报表

**功能描述**：获取指定账本在特定时间范围内的分类统计数据，支持按收入或支出类型统计

**请求方法**：GET
**URL路径**：/api/v1/reports/category
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
| type | string | 是 | - | 统计类型：income（收入）, expense（支出） |
| period | string | 否 | "month" | 统计周期：day, week, month, year |
| date | string | 是 | - | 统计日期，格式：YYYY-MM-DD或YYYY-MM |
| level | integer | 否 | 1 | 分类层级：1（一级分类）, 2（二级分类） |

**请求示例**
```http
GET /api/v1/reports/category?book_id=1&type=expense&period=month&date=2023-01&level=1
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 分类统计数据 |
| data.type | string | 统计类型 |
| data.period | string | 统计周期 |
| data.date | string | 统计日期 |
| data.total_amount | number | 总金额 |
| data.categories | array | 分类数据列表 |
| data.categories[].category_id | integer | 分类ID |
| data.categories[].category_name | string | 分类名称 |
| data.categories[].amount | number | 分类金额 |
| data.categories[].percentage | number | 占总金额比例 |
| data.categories[].count | integer | 交易笔数 |
| data.categories[].icon | string | 分类图标 |
| data.categories[].color | string | 分类颜色 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "type": "expense",
    "period": "month",
    "date": "2023-01",
    "total_amount": 3500.00,
    "categories": [
      {
        "category_id": 5,
        "category_name": "餐饮",
        "amount": 1200.00,
        "percentage": 34.3,
        "count": 30,
        "icon": "food_icon",
        "color": "#FF6B6B"
      },
      {
        "category_id": 10,
        "category_name": "住房",
        "amount": 1000.00,
        "percentage": 28.6,
        "count": 5,
        "icon": "home_icon",
        "color": "#4ECDC4"
      },
      {
        "category_id": 6,
        "category_name": "交通",
        "amount": 500.00,
        "percentage": 14.3,
        "count": 15,
        "icon": "transport_icon",
        "color": "#45B7D1"
      },
      // 更多分类数据...
    ]
  }
}
```

### 获取分类趋势报表

**功能描述**：获取指定账本在特定时间范围内的分类趋势数据，支持按收入或支出类型统计

**请求方法**：GET
**URL路径**：/api/v1/reports/category/trend
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
| type | string | 是 | - | 统计类型：income（收入）, expense（支出） |
| category_id | integer | 是 | - | 分类ID |
| period | string | 否 | "month" | 统计周期：day, week, month, year |
| start_date | string | 是 | - | 开始日期，格式：YYYY-MM-DD或YYYY-MM |
| end_date | string | 是 | - | 结束日期，格式：YYYY-MM-DD或YYYY-MM |

**请求示例**
```http
GET /api/v1/reports/category/trend?book_id=1&type=expense&category_id=5&period=month&start_date=2023-01&end_date=2023-12
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 分类趋势数据 |
| data.category_id | integer | 分类ID |
| data.category_name | string | 分类名称 |
| data.type | string | 统计类型 |
| data.period | string | 统计周期 |
| data.start_date | string | 开始日期 |
| data.end_date | string | 结束日期 |
| data.data | array | 趋势数据列表 |
| data.data[].date | string | 统计日期 |
| data.data[].amount | number | 金额 |
| data.data[].count | integer | 交易笔数 |
| data.summary | object | 汇总数据 |
| data.summary.total_amount | number | 总金额 |
| data.summary.total_count | integer | 总笔数 |
| data.summary.avg_amount | number | 平均金额 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "category_id": 5,
    "category_name": "餐饮",
    "type": "expense",
    "period": "month",
    "start_date": "2023-01",
    "end_date": "2023-12",
    "data": [
      {
        "date": "2023-01",
        "amount": 1200.00,
        "count": 30
      },
      {
        "date": "2023-02",
        "amount": 1300.00,
        "count": 32
      },
      // 更多月份数据...
    ],
    "summary": {
      "total_amount": 15000.00,
      "total_count": 365,
      "avg_amount": 1250.00
    }
  }
}
```

### 获取账户收支报表

**功能描述**：获取指定账本在特定时间范围内的账户收支数据，支持按账户类型统计

**请求方法**：GET
**URL路径**：/api/v1/reports/account
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
| type | string | 否 | "both" | 数据类型：income（收入）, expense（支出）, both（收支） |
| period | string | 否 | "month" | 统计周期：day, week, month, year |
| date | string | 是 | - | 统计日期，格式：YYYY-MM-DD或YYYY-MM |

**请求示例**
```http
GET /api/v1/reports/account?book_id=1&type=both&period=month&date=2023-01
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 账户收支数据 |
| data.type | string | 数据类型 |
| data.period | string | 统计周期 |
| data.date | string | 统计日期 |
| data.accounts | array | 账户收支列表 |
| data.accounts[].account_id | integer | 账户ID |
| data.accounts[].account_name | string | 账户名称 |
| data.accounts[].account_type | string | 账户类型 |
| data.accounts[].income | number | 收入金额 |
| data.accounts[].expense | number | 支出金额 |
| data.accounts[].balance_change | number | 余额变化 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "type": "both",
    "period": "month",
    "date": "2023-01",
    "accounts": [
      {
        "account_id": 1,
        "account_name": "我的工资卡",
        "account_type": "银行卡",
        "income": 8000.00,
        "expense": 2000.00,
        "balance_change": 6000.00
      },
      {
        "account_id": 2,
        "account_name": "支付宝",
        "account_type": "第三方支付",
        "income": 0.00,
        "expense": 1500.00,
        "balance_change": -1500.00
      },
      // 更多账户数据...
    ]
  }
}
```

### 获取账户余额报表

**功能描述**：获取指定账本在特定日期的账户余额报表，包括各账户余额及按账户类型统计的余额分布

**请求方法**：GET
**URL路径**：/api/v1/reports/account/balance
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
| date | string | 是 | - | 统计日期，格式：YYYY-MM-DD |

**请求示例**
```http
GET /api/v1/reports/account/balance?book_id=1&date=2023-01-31
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 账户余额报表数据 |
| data.date | string | 统计日期 |
| data.total_balance | number | 总余额 |
| data.accounts | array | 账户余额列表 |
| data.accounts[].account_id | integer | 账户ID |
| data.accounts[].account_name | string | 账户名称 |
| data.accounts[].account_type | string | 账户类型 |
| data.accounts[].balance | number | 账户余额 |
| data.accounts[].percentage | number | 占总余额比例 |
| data.accounts[].currency | string | 货币类型 |
| data.accounts[].icon | string | 账户图标 |
| data.accounts[].color | string | 账户颜色 |
| data.balance_by_type | array | 按账户类型统计的余额分布 |
| data.balance_by_type[].type | string | 账户类型 |
| data.balance_by_type[].amount | number | 该类型账户总余额 |
| data.balance_by_type[].percentage | number | 占总余额比例 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "date": "2023-01-31",
    "total_balance": 25000.00,
    "accounts": [
      {
        "account_id": 1,
        "account_name": "我的工资卡",
        "account_type": "银行卡",
        "balance": 15000.00,
        "percentage": 60.0,
        "currency": "CNY",
        "icon": "bank_icon",
        "color": "#0080FF"
      },
      {
        "account_id": 2,
        "account_name": "支付宝",
        "account_type": "第三方支付",
        "balance": 5000.00,
        "percentage": 20.0,
        "currency": "CNY",
        "icon": "alipay_icon",
        "color": "#1677FF"
      },
      // 更多账户数据...
    ],
    "balance_by_type": [
      {
        "type": "银行卡",
        "amount": 18000.00,
        "percentage": 72.0
      },
      {
        "type": "第三方支付",
        "amount": 7000.00,
        "percentage": 28.0
      }
    ]
  }
}
```

### 导出Excel报表

**功能描述**：导出指定账本的Excel报表，支持多种报表类型和时间范围

**请求方法**：GET
**URL路径**：/api/v1/export/excel
**权限要求**：账本成员
**限流策略**：每分钟30次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| book_id | integer | 是 | - | 账本ID |
| type | string | 是 | - | 报表类型：transaction（收支明细）, category（分类统计）, account（账户统计） |
| start_date | string | 是 | - | 开始日期，格式：YYYY-MM-DD |
| end_date | string | 是 | - | 结束日期，格式：YYYY-MM-DD |
| format | string | 否 | excel | 导出格式：excel |

**请求示例**
```http
GET /api/v1/export/excel?book_id=1&type=transaction&start_date=2023-01-01&end_date=2023-01-31&format=excel
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 导出结果数据 |
| data.file_url | string | 临时文件URL |
| data.file_name | string | 文件名 |
| data.expires_in | integer | 文件URL过期时间（秒） |

**响应示例**
```json
# 成功
{
  "code": 200,
  "message": "导出成功",
  "data": {
    "file_url": "临时文件URL",
    "file_name": "家庭记账_202301_收支明细.xlsx",
    "expires_in": 3600  // 文件URL过期时间（秒）
  }
}

# 失败
{
  "code": 500,
  "message": "导出失败：服务器错误",
  "data": null
}
```

### 导出CSV报表

**功能描述**：导出指定账本的CSV报表，支持多种报表类型和时间范围

**请求方法**：GET
**URL路径**：/api/v1/export/csv
**权限要求**：账本成员
**限流策略**：每分钟30次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| book_id | integer | 是 | - | 账本ID |
| type | string | 是 | - | 报表类型：transaction（收支明细）, category（分类统计）, account（账户统计） |
| start_date | string | 是 | - | 开始日期，格式：YYYY-MM-DD |
| end_date | string | 是 | - | 结束日期，格式：YYYY-MM-DD |
| format | string | 否 | csv | 导出格式：csv |

**请求示例**
```http
GET /api/v1/export/csv?book_id=1&type=transaction&start_date=2023-01-01&end_date=2023-01-31&format=csv
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 导出结果数据 |
| data.file_url | string | 临时文件URL |
| data.file_name | string | 文件名 |
| data.expires_in | integer | 文件URL过期时间（秒） |

**响应示例**
```json
# 成功
{
  "code": 200,
  "message": "导出成功",
  "data": {
    "file_url": "临时文件URL",
    "file_name": "家庭记账_202301_收支明细.csv",
    "expires_in": 3600  // 文件URL过期时间（秒）
  }
}

# 失败
{
  "code": 500,
  "message": "导出失败：服务器错误",
  "data": null
}
```

### 导出PDF报表

**功能描述**：导出指定账本的PDF报表，支持多种报表类型和时间范围

**请求方法**：GET
**URL路径**：/api/v1/export/pdf
**权限要求**：账本成员
**限流策略**：每分钟30次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| book_id | integer | 是 | - | 账本ID |
| type | string | 是 | - | 报表类型：transaction（收支明细）, category（分类统计）, account（账户统计） |
| start_date | string | 是 | - | 开始日期，格式：YYYY-MM-DD |
| end_date | string | 是 | - | 结束日期，格式：YYYY-MM-DD |
| format | string | 否 | pdf | 导出格式：pdf |
| orientation | string | 否 | portrait | 页面方向：portrait（纵向）, landscape（横向） |

**请求示例**
```http
GET /api/v1/export/pdf?book_id=1&type=category&start_date=2023-01-01&end_date=2023-01-31&format=pdf&orientation=landscape
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 导出结果数据 |
| data.file_url | string | 临时文件URL |
| data.file_name | string | 文件名 |
| data.expires_in | integer | 文件URL过期时间（秒） |

**响应示例**
```json
# 成功
{
  "code": 200,
  "message": "导出成功",
  "data": {
    "file_url": "临时文件URL",
    "file_name": "家庭记账_202301_分类统计.pdf",
    "expires_in": 3600  // 文件URL过期时间（秒）
  }
}

# 失败
{
  "code": 500,
  "message": "导出失败：服务器错误",
  "data": null
}
```

### 获取成员收支报表

**功能描述**：获取指定账本在特定时间范围内的成员收支数据，支持按成员统计

**请求方法**：GET
**URL路径**：/api/v1/reports/member
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
| period | string | 否 | "month" | 统计周期：day, week, month, year |
| date | string | 是 | - | 统计日期，格式：YYYY-MM-DD或YYYY-MM |
| member_id | integer | 否 | - | 成员ID，不填则统计所有成员 |

**请求示例**
```http
GET /api/v1/reports/member?book_id=1&period=month&date=2023-01
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 成员收支数据 |
| data.period | string | 统计周期 |
| data.date | string | 统计日期 |
| data.members | array | 成员收支列表 |
| data.members[].member_id | integer | 成员ID |
| data.members[].member_name | string | 成员名称 |
| data.members[].income | number | 收入金额 |
| data.members[].expense | number | 支出金额 |
| data.members[].balance | number | 结余金额 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "period": "month",
    "date": "2023-01",
    "members": [
      {
        "member_id": 1,
        "member_name": "张三",
        "income": 8000.00,
        "expense": 2000.00,
        "balance": 6000.00
      },
      {
        "member_id": 2,
        "member_name": "李四",
        "income": 0.00,
        "expense": 1500.00,
        "balance": -1500.00
      },
      // 更多成员数据...
    ]
  }
}
```

### 获取成员贡献度分析

**功能描述**：获取指定账本在特定时间范围内的成员贡献度数据，分析各成员的财务贡献

**请求方法**：GET
**URL路径**：/api/v1/reports/member/contribution
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
| period | string | 否 | "month" | 统计周期：day, week, month, year |
| date | string | 是 | - | 统计日期，格式：YYYY-MM-DD或YYYY-MM |

**请求示例**
```http
GET /api/v1/reports/member/contribution?book_id=1&period=month&date=2023-01
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 成员贡献度数据 |
| data.period | string | 统计周期 |
| data.date | string | 统计日期 |
| data.total_income | number | 总输入 |
| data.total_expense | number | 总支出 |
| data.members | array | 成员贡献度列表 |
| data.members[].member_id | integer | 成员ID |
| data.members[].member_name | string | 成员名称 |
| data.members[].income | number | 收入金额 |
| data.members[].income_percentage | number | 收入占比 |
| data.members[].expense | number | 支出金额 |
| data.members[].expense_percentage | number | 支出占比 |
| data.members[].net_contribution | number | 净贡献 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "period": "month",
    "date": "2023-01",
    "total_income": 8000.00,
    "total_expense": 3500.00,
    "members": [
      {
        "member_id": 1,
        "member_name": "张三",
        "income": 8000.00,
        "income_percentage": 100.0,
        "expense": 2000.00,
        "expense_percentage": 57.1,
        "net_contribution": 6000.00
      },
      {
        "member_id": 2,
        "member_name": "李四",
        "income": 0.00,
        "income_percentage": 0.0,
        "expense": 1500.00,
        "expense_percentage": 42.9,
        "net_contribution": -1500.00
      }
    ]
  }
}
```

### 获取预算执行报表

**功能描述**：获取指定账本在特定时间范围内的预算执行情况，支持按分类统计

**请求方法**：GET
**URL路径**：/api/v1/reports/budget
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
| period | string | 否 | "month" | 统计周期：month, quarter, year |
| date | string | 是 | - | 统计日期，格式：YYYY-MM |
| category_id | integer | 否 | - | 分类ID，不填则统计所有分类 |

**请求示例**
```http
GET /api/v1/reports/budget?book_id=1&period=month&date=2023-01
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 预算执行数据 |
| data.period | string | 统计周期 |
| data.date | string | 统计日期 |
| data.total_budget | number | 总预算 |
| data.total_used | number | 已使用金额 |
| data.total_remaining | number | 剩余金额 |
| data.overall_usage_rate | number | 整体使用率 |
| data.categories | array | 分类预算执行情况 |
| data.categories[].category_id | integer | 分类ID |
| data.categories[].category_name | string | 分类名称 |
| data.categories[].budget | number | 分类预算 |
| data.categories[].used | number | 已使用金额 |
| data.categories[].remaining | number | 剩余金额 |
| data.categories[].usage_rate | number | 分类使用率 |
| data.categories[].is_over_budget | boolean | 是否超预算 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "period": "month",
    "date": "2023-01",
    "total_budget": 5000.00,
    "total_used": 3500.00,
    "total_remaining": 1500.00,
    "overall_usage_rate": 70.0,
    "categories": [
      {
        "category_id": 5,
        "category_name": "餐饮",
        "budget": 1500.00,
        "used": 1200.00,
        "remaining": 300.00,
        "usage_rate": 80.0,
        "is_over_budget": false
      },
      {
        "category_id": 10,
        "category_name": "住房",
        "budget": 1000.00,
        "used": 1000.00,
        "remaining": 0.00,
        "usage_rate": 100.0,
        "is_over_budget": false
      },
      {
        "category_id": 6,
        "category_name": "交通",
        "budget": 500.00,
        "used": 600.00,
        "remaining": -100.00,
        "usage_rate": 120.0,
        "is_over_budget": true
      }
    ]
  }
}
```

### 获取预算超支提醒

**功能描述**：获取指定账本的预算超支提醒信息，支持按分类和成员筛选

**请求方法**：GET
**URL路径**：/api/v1/reports/budget/alert
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
| period | string | 否 | "month" | 统计周期：month, quarter, year |
| date | string | 否 | 当前日期 | 统计日期，格式：YYYY-MM |
| threshold | number | 否 | 80 | 提醒阈值，百分比（0-100） |

**请求示例**
```http
GET /api/v1/reports/budget/alert?book_id=1&period=month&date=2023-01&threshold=80
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 预算超支提醒数据 |
| data.period | string | 统计周期 |
| data.date | string | 统计日期 |
| data.threshold | number | 提醒阈值 |
| data.alerts | array | 超支提醒列表 |
| data.alerts[].category_id | integer | 分类ID |
| data.alerts[].category_name | string | 分类名称 |
| data.alerts[].budget | number | 分类预算 |
| data.alerts[].used | number | 已使用金额 |
| data.alerts[].usage_rate | number | 使用率 |
| data.alerts[].over_amount | number | 超支金额 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "period": "month",
    "date": "2023-01",
    "threshold": 80,
    "alerts": [
      {
        "category_id": 5,
        "category_name": "餐饮",
        "budget": 1500.00,
        "used": 1200.00,
        "usage_rate": 80.0,
        "over_amount": 0.00
      },
      {
        "category_id": 6,
        "category_name": "交通",
        "budget": 500.00,
        "used": 600.00,
        "usage_rate": 120.0,
        "over_amount": 100.00
      }
    ]
  }
}
```

### 获取折线图数据

**功能描述**：获取指定账本的折线图数据，支持多种数据维度和时间范围

**请求方法**：GET
**URL路径**：/api/v1/charts/line
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
| type | string | 是 | - | 图表类型：income-expense（收支对比）, category-trend（分类趋势）, account-trend（账户趋势） |
| period | string | 否 | "month" | 统计周期：day, week, month, year |
| start_date | string | 是 | - | 开始日期，格式：YYYY-MM-DD或YYYY-MM |
| end_date | string | 是 | - | 结束日期，格式：YYYY-MM-DD或YYYY-MM |
| category_ids | array | 否 | - | 分类ID列表，多个用逗号分隔 |

**请求示例**
```http
GET /api/v1/charts/line?book_id=1&type=income-expense&period=month&start_date=2023-01&end_date=2023-12
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 折线图数据 |
| data.title | string | 图表标题 |
| data.x_axis | array | X轴数据 |
| data.series | array | 系列数据 |
| data.series[].name | string | 系列名称 |
| data.series[].data | array | 系列数据点 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "title": "2023年收支趋势",
    "x_axis": ["2023-01", "2023-02", "2023-03", "2023-04", "2023-05", "2023-06", "2023-07", "2023-08", "2023-09", "2023-10", "2023-11", "2023-12"],
    "series": [
      {
        "name": "收入",
        "data": [8000.00, 8200.00, 8500.00, 8300.00, 8600.00, 8400.00, 8700.00, 8500.00, 8800.00, 8600.00, 8900.00, 9000.00]
      },
      {
        "name": "支出",
        "data": [3500.00, 3800.00, 3600.00, 3700.00, 3900.00, 3800.00, 4000.00, 3900.00, 4100.00, 4000.00, 4200.00, 4300.00]
      }
    ]
  }
}
```

### 获取饼图数据

**功能描述**：获取指定账本的饼图数据，支持多种数据维度

**请求方法**：GET
**URL路径**：/api/v1/charts/pie
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
| type | string | 是 | - | 图表类型：category（分类占比）, account（账户占比）, member（成员占比） |
| period | string | 否 | "month" | 统计周期：day, week, month, year |
| date | string | 是 | - | 统计日期，格式：YYYY-MM-DD或YYYY-MM |
| data_type | string | 否 | "expense" | 数据类型：income（收入）, expense（支出） |

**请求示例**
```http
GET /api/v1/charts/pie?book_id=1&type=category&period=month&date=2023-01&data_type=expense
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 饼图数据 |
| data.title | string | 图表标题 |
| data.series | array | 系列数据 |
| data.series[].name | string | 系列名称 |
| data.series[].data | array | 系列数据点 |
| data.series[].data[].name | string | 数据项名称 |
| data.series[].data[].value | number | 数据项值 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "title": "2023年1月支出分类占比",
    "series": [
      {
        "name": "支出分类",
        "data": [
          { "name": "餐饮", "value": 1200.00 },
          { "name": "住房", "value": 1000.00 },
          { "name": "交通", "value": 500.00 },
          { "name": "购物", "value": 400.00 },
          { "name": "娱乐", "value": 300.00 },
          { "name": "其他", "value": 100.00 }
        ]
      }
    ]
  }
}
```

### 获取柱状图数据

**功能描述**：获取指定账本的柱状图数据，支持多种数据维度和时间范围

**请求方法**：GET
**URL路径**：/api/v1/charts/bar
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
| type | string | 是 | - | 图表类型：income-expense（收支对比）, category（分类对比）, account（账户对比） |
| period | string | 否 | "month" | 统计周期：day, week, month, year |
| start_date | string | 是 | - | 开始日期，格式：YYYY-MM-DD或YYYY-MM |
| end_date | string | 是 | - | 结束日期，格式：YYYY-MM-DD或YYYY-MM |

**请求示例**
```http
GET /api/v1/charts/bar?book_id=1&type=category&period=month&start_date=2023-01&end_date=2023-03
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 柱状图数据 |
| data.title | string | 图表标题 |
| data.x_axis | array | X轴数据 |
| data.series | array | 系列数据 |
| data.series[].name | string | 系列名称 |
| data.series[].data | array | 系列数据点 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "title": "2023年1-3月支出分类对比",
    "x_axis": ["餐饮", "住房", "交通", "购物", "娱乐"],
    "series": [
      {
        "name": "1月",
        "data": [1200.00, 1000.00, 500.00, 400.00, 300.00]
      },
      {
        "name": "2月",
        "data": [1300.00, 1000.00, 550.00, 450.00, 350.00]
      },
      {
        "name": "3月",
        "data": [1250.00, 1000.00, 520.00, 420.00, 330.00]
      }
    ]
  }
}
```

### 获取雷达图数据

**功能描述**：获取指定账本的雷达图数据，支持多维度对比分析

**请求方法**：GET
**URL路径**：/api/v1/charts/radar
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
| period | string | 否 | "month" | 统计周期：month, quarter, year |
| start_date | string | 是 | - | 开始日期，格式：YYYY-MM |
| end_date | string | 是 | - | 结束日期，格式：YYYY-MM |
| compare_period | string | 否 | "previous" | 对比周期：previous（上一周期）, same（同比） |

**请求示例**
```http
GET /api/v1/charts/radar?book_id=1&period=quarter&start_date=2023-Q1&end_date=2023-Q2&compare_period=previous
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 雷达图数据 |
| data.title | string | 图表标题 |
| data.indicator | array | 指标数据 |
| data.indicator[].name | string | 指标名称 |
| data.indicator[].max | number | 指标最大值 |
| data.series | array | 系列数据 |
| data.series[].name | string | 系列名称 |
| data.series[].data | array | 系列数据点 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "title": "2023年Q1-Q2支出分类对比",
    "indicator": [
      { "name": "餐饮", "max": 2000 },
      { "name": "住房", "max": 1500 },
      { "name": "交通", "max": 1000 },
      { "name": "购物", "max": 1000 },
      { "name": "娱乐", "max": 800 },
      { "name": "其他", "max": 500 }
    ],
    "series": [
      {
        "name": "Q1",
        "data": [1800.00, 1500.00, 750.00, 600.00, 450.00, 150.00]
      },
      {
        "name": "Q2",
        "data": [1950.00, 1500.00, 825.00, 675.00, 525.00, 165.00]
      }
    ]
  }
}
```

### 创建自定义报表

**功能描述**：创建自定义报表配置，支持多种报表类型和参数设置

**请求方法**：POST
**URL路径**：/api/v1/reports/custom
**权限要求**：账本成员
**限流策略**：每分钟30次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |
| Content-Type | string | 是 | application/json |

**请求体**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| book_id | integer | 是 | - | 账本ID |
| name | string | 是 | - | 报表名称 |
| type | string | 是 | - | 报表类型：category（分类统计）, account（账户统计）, member（成员统计）, trend（趋势分析） |
| params | object | 是 | - | 报表参数配置 |
| params.period | string | 否 | "month" | 统计周期 |
| params.data_type | string | 否 | "expense" | 数据类型 |
| params.category_ids | array | 否 | - | 分类ID列表 |
| display_columns | array | 否 | - | 显示列配置 |
| chart_type | string | 否 | - | 图表类型：line（折线图）, bar（柱状图）, pie（饼图）, radar（雷达图） |

**请求示例**
```http
POST /api/v1/reports/custom
Authorization: Bearer jwt_token_string
Content-Type: application/json

{
  "book_id": 1,
  "name": "月度餐饮分析",
  "type": "category",
  "params": {
    "period": "month",
    "data_type": "expense",
    "category_ids": [5]
  },
  "display_columns": ["date", "amount", "percentage", "count"],
  "chart_type": "line"
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 自定义报表配置 |
| data.report_id | integer | 报表ID |
| data.name | string | 报表名称 |
| data.type | string | 报表类型 |
| data.params | object | 报表参数配置 |
| data.display_columns | array | 显示列配置 |
| data.chart_type | string | 图表类型 |
| data.created_at | string | 创建时间 |
| data.updated_at | string | 更新时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "report_id": 1,
    "name": "月度餐饮分析",
    "type": "category",
    "params": {
      "period": "month",
      "data_type": "expense",
      "category_ids": [5]
    },
    "display_columns": ["date", "amount", "percentage", "count"],
    "chart_type": "line",
    "created_at": "2023-01-01T12:00:00Z",
    "updated_at": "2023-01-01T12:00:00Z"
  }
}
```

### 获取自定义报表列表

**功能描述**：获取指定账本的自定义报表列表，支持分页和筛选

**请求方法**：GET
**URL路径**：/api/v1/reports/custom
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
| type | string | 否 | - | 报表类型 |
| limit | integer | 否 | 10 | 每页记录数，最大50 |
| offset | integer | 否 | 0 | 偏移量 |

**请求示例**
```http
GET /api/v1/reports/custom?book_id=1&type=category&limit=10&offset=0
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 自定义报表列表数据 |
| data.list | array | 报表列表 |
| data.list[].report_id | integer | 报表ID |
| data.list[].name | string | 报表名称 |
| data.list[].type | string | 报表类型 |
| data.list[].chart_type | string | 图表类型 |
| data.list[].created_at | string | 创建时间 |
| data.list[].updated_at | string | 更新时间 |
| data.total | integer | 总记录数 |
| data.limit | integer | 每页记录数 |
| data.offset | integer | 偏移量 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "list": [
      {
        "report_id": 1,
        "name": "月度餐饮分析",
        "type": "category",
        "chart_type": "line",
        "created_at": "2023-01-01T12:00:00Z",
        "updated_at": "2023-01-01T12:00:00Z"
      },
      {
        "report_id": 2,
        "name": "季度收支对比",
        "type": "trend",
        "chart_type": "bar",
        "created_at": "2023-01-02T10:30:00Z",
        "updated_at": "2023-01-02T10:30:00Z"
      }
    ],
    "total": 2,
    "limit": 10,
    "offset": 0
  }
}
```

### 获取自定义报表数据

**功能描述**：获取指定自定义报表的数据，支持动态参数调整

**请求方法**：GET
**URL路径**：/api/v1/reports/custom/:id
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| id | integer | 是 | 报表ID |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| book_id | integer | 是 | - | 账本ID |
| date | string | 否 | 当前日期 | 统计日期，格式：YYYY-MM-DD或YYYY-MM |
| start_date | string | 否 | - | 开始日期，格式：YYYY-MM-DD或YYYY-MM |
| end_date | string | 否 | - | 结束日期，格式：YYYY-MM-DD或YYYY-MM |

**请求示例**
```http
GET /api/v1/reports/custom/1?book_id=1&date=2023-01
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 自定义报表数据 |
| data.report_info | object | 报表信息 |
| data.report_info.report_id | integer | 报表ID |
| data.report_info.name | string | 报表名称 |
| data.report_info.type | string | 报表类型 |
| data.report_data | object | 报表数据 |
| data.chart_data | object | 图表数据 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "report_info": {
      "report_id": 1,
      "name": "月度餐饮分析",
      "type": "category"
    },
    "report_data": {
      "period": "month",
      "date": "2023-01",
      "total_amount": 1200.00,
      "categories": [
        {
          "category_id": 5,
          "category_name": "餐饮",
          "amount": 1200.00,
          "percentage": 100.0,
          "count": 30
        }
      ]
    },
    "chart_data": {
      "title": "月度餐饮分析",
      "x_axis": ["2023-01-01", "2023-01-02", ..., "2023-01-31"],
      "series": [
        {
          "name": "餐饮支出",
          "data": [40.00, 0.00, 50.00, ..., 30.00]
        }
      ]
    }
  }
}
```

### 更新自定义报表

**功能描述**：更新指定自定义报表的配置信息

**请求方法**：PUT
**URL路径**：/api/v1/reports/custom/:id
**权限要求**：账本成员
**限流策略**：每分钟30次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |
| Content-Type | string | 是 | application/json |

**路径参数**
| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| id | integer | 是 | 报表ID |

**请求体**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| book_id | integer | 是 | - | 账本ID |
| name | string | 否 | - | 报表名称 |
| params | object | 否 | - | 报表参数配置 |
| display_columns | array | 否 | - | 显示列配置 |
| chart_type | string | 否 | - | 图表类型 |

**请求示例**
```http
PUT /api/v1/reports/custom/1
Authorization: Bearer jwt_token_string
Content-Type: application/json

{
  "book_id": 1,
  "name": "更新后的月度餐饮分析",
  "params": {
    "period": "week",
    "data_type": "expense"
  }
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 更新后的报表配置 |
| data.report_id | integer | 报表ID |
| data.name | string | 报表名称 |
| data.type | string | 报表类型 |
| data.params | object | 报表参数配置 |
| data.display_columns | array | 显示列配置 |
| data.chart_type | string | 图表类型 |
| data.updated_at | string | 更新时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "更新成功",
  "data": {
    "report_id": 1,
    "name": "更新后的月度餐饮分析",
    "type": "category",
    "params": {
      "period": "week",
      "data_type": "expense",
      "category_ids": [5]
    },
    "display_columns": ["date", "amount", "percentage", "count"],
    "chart_type": "line",
    "updated_at": "2023-01-15T14:30:00Z"
  }
}
```

### 删除自定义报表

**功能描述**：删除指定的自定义报表配置

**请求方法**：DELETE
**URL路径**：/api/v1/reports/custom/:id
**权限要求**：账本成员
**限流策略**：每分钟30次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| id | integer | 是 | 报表ID |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| book_id | integer | 是 | - | 账本ID |

**请求示例**
```http
DELETE /api/v1/reports/custom/1?book_id=1
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | null | 无数据 |

**响应示例**
```json
{
  "code": 200,
  "message": "删除成功",
  "data": null
}
```

## 数据模型

### 财务概览模型

```javascript
{
  "total_income": 8000.00,     // 总收入
  "total_expense": 3500.00,    // 总支出
  "balance": 4500.00,          // 结余
  "income_rate": 10.5,         // 同比增长率
  "expense_rate": -5.2,        // 同比增长率
  "date_range": {
    "start": "2023-01-01",
    "end": "2023-01-31"
  }
}
```

### 趋势数据模型

```javascript
{
  "period": "month",           // 时间周期：day, week, month, year
  "data": [
    {
      "date": "2023-01",      // 日期
      "income": 8000.00,       // 收入
      "expense": 3500.00,      // 支出
      "balance": 4500.00,      // 结余
      "income_change": 10.5,   // 环比变化率
      "expense_change": -5.2   // 环比变化率
    }
  ]
}
```

### 分类统计模型

```javascript
{
  "category_id": 5,            // 分类ID
  "category_name": "餐饮",     // 分类名称
  "amount": 1200.00,           // 金额
  "percentage": 34.3,          // 占比
  "count": 30,                 // 记录数
  "icon": "food_icon",        // 图标
  "color": "#FF6B6B",         // 颜色
  "subcategories": [           // 子分类数据
    {
      "category_id": 51,
      "category_name": "午餐",
      "amount": 800.00,
      "percentage": 66.7
    }
  ]
}
```

### 自定义报表配置模型

```javascript
{
  "report_id": 1,              // 报表ID
  "name": "月度餐饮分析",       // 报表名称
  "type": "category",         // 报表类型
  "params": {
    "book_id": 1,
    "category_id": 5,
    "period": "month",
    "date_range": "last_6_month"
  },
  "display_columns": ["date", "amount", "percentage", "count"],
  "chart_type": "bar",        // 图表类型
  "created_at": "2023-01-01T12:00:00Z", // 创建时间
  "updated_at": "2023-01-01T12:00:00Z"  // 更新时间
}
```

## 统计维度说明

### 时间维度

- **日**: 按天统计数据
- **周**: 按周统计数据
- **月**: 按月统计数据
- **季**: 按季度统计数据
- **年**: 按年统计数据
- **自定义**: 自定义时间范围

### 统计类型

- **收支对比**: 收入与支出对比
- **分类统计**: 按分类统计收支
- **账户统计**: 按账户统计收支
- **成员统计**: 按成员统计收支
- **标签统计**: 按标签统计收支
- **项目统计**: 按项目统计收支

## 图表类型说明

### 支持的图表类型

- **折线图**: 适合展示趋势变化
- **饼图**: 适合展示占比分析
- **柱状图**: 适合展示对比分析
- **雷达图**: 适合展示多维度对比
- **热力图**: 适合展示活跃度分析
- **散点图**: 适合展示相关性分析

## 错误码说明

| 错误码 | 描述 |
|-------|------|
| 400 | 请求参数错误 |
| 401 | 未授权，需要登录 |
| 403 | 权限不足 |
| 404 | 报表不存在 |
| 500 | 服务器内部错误 |
| 900 | 无效的统计维度 |
| 901 | 无效的时间范围 |
| 902 | 数据量过大，请缩小时间范围 |
| 903 | 不支持的导出格式 |
| 904 | 导出失败 |