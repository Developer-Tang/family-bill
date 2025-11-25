##  模块概述

统计分析模块提供多维度的财务数据分析功能，通过图表和报表帮助用户了解财务状况，辅助决策。本模块提供了完整的数据看板、收支分析、类别统计、趋势分析等功能的API接口。

##  接口清单

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

##  详细接口说明

###  获取财务概览数据

**请求**

```http
GET /api/v1/dashboard/summary?book_id=1&period=month&date=2023-01
Authorization: Bearer jwt_token_string
```

**响应**

```
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

###  获取收支趋势分析

**请求**

```http
GET /api/v1/analysis/trend?book_id=1&type=both&period=month&start_date=2023-01&end_date=2023-12
Authorization: Bearer jwt_token_string
```

**响应**

```
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

###  获取分类统计报表

**请求**

```http
GET /api/v1/reports/category?book_id=1&type=expense&period=month&date=2023-01&level=1
Authorization: Bearer jwt_token_string
```

**响应**

```
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

###  获取账户余额报表

**请求**

```http
GET /api/v1/reports/account/balance?book_id=1&date=2023-01-31
Authorization: Bearer jwt_token_string
```

**响应**

```
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

###  导出Excel报表

**请求**

```http
GET /api/v1/export/excel?book_id=1&type=transaction&start_date=2023-01-01&end_date=2023-01-31&format=excel
Authorization: Bearer jwt_token_string
```

**响应**

```
# 成功 - 返回文件流
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

##  数据模型

###  财务概览模型

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

###  趋势数据模型

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

###  分类统计模型

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

###  自定义报表配置模型

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

##  统计维度说明

###  时间维度

- **日**: 按天统计数据
- **周**: 按周统计数据
- **月**: 按月统计数据
- **季**: 按季度统计数据
- **年**: 按年统计数据
- **自定义**: 自定义时间范围

###  统计类型

- **收支对比**: 收入与支出对比
- **分类统计**: 按分类统计收支
- **账户统计**: 按账户统计收支
- **成员统计**: 按成员统计收支
- **标签统计**: 按标签统计收支
- **项目统计**: 按项目统计收支

##  图表类型说明

###  支持的图表类型

- **折线图**: 适合展示趋势变化
- **饼图**: 适合展示占比分析
- **柱状图**: 适合展示对比分析
- **雷达图**: 适合展示多维度对比
- **热力图**: 适合展示活跃度分析
- **散点图**: 适合展示相关性分析

##  错误码说明

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