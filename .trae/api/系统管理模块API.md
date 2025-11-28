## 模块概述

系统管理模块是家庭记账系统的核心管理组件，负责提供用户个性化设置、分类管理、账本管理、系统日志、系统配置和通知管理等基础功能。本模块通过RESTful API接口，为用户提供完整的系统管理能力，确保系统的正常运行和个性化定制。

### 设计原则

1. **RESTful规范**：所有API接口严格遵循RESTful设计规范，使用标准HTTP方法和资源导向的URL路径
2. **本地缓存优先**：暂不集成远程缓存机制，如需使用缓存功能，严格限定为本地缓存方案
3. **完整的请求验证**：所有API接口包含严格的请求参数验证，确保数据完整性和安全性
4. **统一的错误处理**：采用统一的错误响应格式和错误码体系，便于客户端处理
5. **清晰的响应格式**：所有API响应采用标准JSON格式，包含明确的数据结构和状态信息
6. **可扩展性**：API设计考虑未来功能扩展，预留合理的扩展点

### 适用场景

- 用户需要个性化设置系统界面和功能
- 管理员需要管理账本和分类
- 系统需要记录和查询操作日志
- 用户需要管理通知设置和查看通知
- 管理员需要查看系统信息和配置

## 接口清单

<!-- tabs:start -->
<!-- tab:用户设置 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/user/settings`](#获取用户设置) | `GET` | 获取用户设置 |
| [`/api/v1/user/settings`](#更新用户设置) | `PUT` | 更新用户设置 |
| [`/api/v1/user/profile`](#获取用户资料) | `GET` | 获取用户资料 |
| [`/api/v1/user/profile`](#更新用户资料) | `PUT` | 更新用户资料 |
| [`/api/v1/user/avatar`](#上传用户头像) | `POST` | 上传用户头像 |
| [`/api/v1/user/password`](#修改密码) | `PUT` | 修改密码 |

<!-- tab:分类管理 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/categories`](#获取分类列表) | `GET` | 获取分类列表 |
| [`/api/v1/categories`](#创建分类) | `POST` | 创建分类 |
| [`/api/v1/categories/:id`](#更新分类) | `PUT` | 更新分类 |
| [`/api/v1/categories/:id`](#删除分类) | `DELETE` | 删除分类 |
| [`/api/v1/categories/import`](#导入分类) | `POST` | 导入分类 |
| [`/api/v1/categories/export`](#导出分类) | `GET` | 导出分类 |

<!-- tab:账本管理 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/books`](#获取账本列表) | `GET` | 获取账本列表 |
| [`/api/v1/books`](#创建账本) | `POST` | 创建账本 |
| [`/api/v1/books/:id`](#获取账本详情) | `GET` | 获取账本详情 |
| [`/api/v1/books/:id`](#更新账本) | `PUT` | 更新账本 |
| [`/api/v1/books/:id`](#删除账本) | `DELETE` | 删除账本 |
| [`/api/v1/books/:id/switch`](#切换当前账本) | `POST` | 切换当前账本 |

<!-- tab:系统日志 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/logs/operation`](#获取操作日志) | `GET` | 获取操作日志 |
| [`/api/v1/logs/error`](#获取错误日志) | `GET` | 获取错误日志 |
| [`/api/v1/logs/system`](#获取系统日志) | `GET` | 获取系统日志 |
| [`/api/v1/logs/user-activity`](#获取用户活动日志) | `GET` | 获取用户活动日志 |

<!-- tab:系统配置 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/system/config`](#获取系统配置) | `GET` | 获取系统配置 |
| [`/api/v1/system/config`](#更新系统配置) | `PUT` | 更新系统配置 |
| [`/api/v1/system/info`](#获取系统信息) | `GET` | 获取系统信息 |
| [`/api/v1/system/check-update`](#检查更新) | `GET` | 检查更新 |
| [`/api/v1/system/version`](#获取版本信息) | `GET` | 获取版本信息 |

<!-- tab:通知管理 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/notifications`](#获取通知列表) | `GET` | 获取通知列表 |
| [`/api/v1/notifications/:id`](#标记通知已读) | `PUT` | 标记通知已读 |
| [`/api/v1/notifications/read-all`](#标记所有通知已读) | `PUT` | 标记所有通知已读 |
| [`/api/v1/notifications/:id`](#删除通知) | `DELETE` | 删除通知 |
| [`/api/v1/notifications/settings`](#获取通知设置) | `GET` | 获取通知设置 |
| [`/api/v1/notifications/settings`](#更新通知设置) | `PUT` | 更新通知设置 |

<!-- tabs:end -->

## 详细接口说明

### 获取用户设置

**功能描述**：获取当前用户的个性化设置，包括主题、语言、日期格式等

**请求方法**：GET
**URL路径**：/api/v1/user/settings
**权限要求**：已登录用户
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**请求示例**
```http
GET /api/v1/user/settings
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| settings_id | integer | 设置ID |
| user_id | integer | 用户ID |
| theme | string | 主题：light, dark, system |
| language | string | 语言 |
| date_format | string | 日期格式 |
| time_format | string | 时间格式：12h, 24h |
| currency_format | string | 货币格式：symbol_before, symbol_after |
| decimal_places | integer | 小数位数 |
| auto_sync | boolean | 是否自动同步 |
| notifications | object | 通知设置 |
| notifications.email | boolean | 是否开启邮件通知 |
| notifications.push | boolean | 是否开启推送通知 |
| notifications.budget_alerts | boolean | 是否开启预算提醒 |
| notifications.reminders | boolean | 是否开启账单提醒 |
| notifications.updates | boolean | 是否开启更新通知 |
| created_at | string | 创建时间 |
| updated_at | string | 更新时间 |

**成功响应**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "settings_id": 1,
    "user_id": 1,
    "theme": "light",
    "language": "zh-CN",
    "date_format": "YYYY-MM-DD",
    "time_format": "24h",
    "currency_format": "symbol_before",
    "decimal_places": 2,
    "auto_sync": true,
    "notifications": {
      "email": true,
      "push": true,
      "budget_alerts": true,
      "reminders": true,
      "updates": true
    },
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-05T12:30:00Z"
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 500 | 服务器内部错误 | 请稍后重试，或联系系统管理员 |

**本地缓存策略**：获取成功后，本地缓存该设置信息，缓存时间1小时

### 更新用户设置

**功能描述**：更新当前用户的个性化设置，包括主题、语言、日期格式等

**请求方法**：PUT
**URL路径**：/api/v1/user/settings
**权限要求**：已登录用户
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |
| Authorization | string | 是 | Bearer JWT令牌 |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| theme | string | 否 | - | 主题 | 枚举值：light, dark, system |
| language | string | 否 | - | 语言 | 有效的语言代码 |
| date_format | string | 否 | - | 日期格式 | 有效的日期格式字符串 |
| time_format | string | 否 | - | 时间格式 | 枚举值：12h, 24h |
| currency_format | string | 否 | - | 货币格式 | 枚举值：symbol_before, symbol_after |
| decimal_places | integer | 否 | - | 小数位数 | 1-4之间的整数 |
| auto_sync | boolean | 否 | - | 是否自动同步 | 布尔值 |
| notifications | object | 否 | - | 通知设置 | - |
| notifications.email | boolean | 否 | - | 是否开启邮件通知 | 布尔值 |
| notifications.push | boolean | 否 | - | 是否开启推送通知 | 布尔值 |
| notifications.budget_alerts | boolean | 否 | - | 是否开启预算提醒 | 布尔值 |
| notifications.reminders | boolean | 否 | - | 是否开启账单提醒 | 布尔值 |
| notifications.updates | boolean | 否 | - | 是否开启更新通知 | 布尔值 |

**请求示例**
```http
PUT /api/v1/user/settings
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "theme": "dark",
  "language": "zh-CN",
  "date_format": "YYYY/MM/DD",
  "time_format": "24h",
  "currency_format": "symbol_before",
  "decimal_places": 2,
  "auto_sync": true,
  "notifications": {
    "email": true,
    "push": false,
    "budget_alerts": true,
    "reminders": true,
    "updates": false
  }
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| settings_id | integer | 设置ID |
| user_id | integer | 用户ID |
| theme | string | 主题：light, dark, system |
| language | string | 语言 |
| date_format | string | 日期格式 |
| time_format | string | 时间格式：12h, 24h |
| currency_format | string | 货币格式：symbol_before, symbol_after |
| decimal_places | integer | 小数位数 |
| auto_sync | boolean | 是否自动同步 |
| notifications | object | 通知设置 |
| notifications.email | boolean | 是否开启邮件通知 |
| notifications.push | boolean | 是否开启推送通知 |
| notifications.budget_alerts | boolean | 是否开启预算提醒 |
| notifications.reminders | boolean | 是否开启账单提醒 |
| notifications.updates | boolean | 是否开启更新通知 |
| updated_at | string | 更新时间 |

**成功响应**
```json
{
  "code": 200,
  "message": "更新成功",
  "data": {
    "settings_id": 1,
    "user_id": 1,
    "theme": "dark",
    "language": "zh-CN",
    "date_format": "YYYY/MM/DD",
    "time_format": "24h",
    "currency_format": "symbol_before",
    "decimal_places": 2,
    "auto_sync": true,
    "notifications": {
      "email": true,
      "push": false,
      "budget_alerts": true,
      "reminders": true,
      "updates": false
    },
    "updated_at": "2023-01-05T12:35:00Z"
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误 | 检查请求参数是否符合要求，包括类型、格式、必填项等 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 3000 | 设置更新失败 | 检查设置参数是否正确，特别是主题、语言等枚举值 |

**本地缓存策略**：更新成功后，更新本地缓存的设置信息，缓存时间1小时

### 创建分类

**功能描述**：创建新的收支分类，用于对收支记录进行分类管理

**请求方法**：POST
**URL路径**：/api/v1/categories
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
| name | string | 是 | - | 分类名称 | 1-50个字符，同一账本内同类型分类名称唯一 |
| type | string | 是 | - | 分类类型 | 枚举值：income, expense, transfer |
| parent_id | integer | 否 | null | 父分类ID，顶级分类为null | 正整数或null，父分类必须存在且类型相同 |
| icon | string | 否 | "default" | 分类图标 | 1-20个字符，必须是系统支持的图标名称 |
| color | string | 否 | "#4CAF50" | 分类颜色 | 有效的十六进制颜色代码 |
| sort_order | integer | 否 | 0 | 排序顺序 | 整数 |
| is_system | boolean | 否 | false | 是否系统分类 | 布尔值，普通用户只能创建非系统分类 |
| description | string | 否 | "" | 分类描述 | 0-100个字符 |

**请求示例**
```http
POST /api/v1/categories
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "name": "日常购物",
  "type": "expense",
  "parent_id": null,
  "icon": "shopping-bag",
  "color": "#4CAF50",
  "sort_order": 1,
  "is_system": false,
  "description": "日常购物消费"
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| category_id | integer | 分类ID |
| book_id | integer | 账本ID |
| name | string | 分类名称 |
| type | string | 分类类型：income, expense, transfer |
| parent_id | integer | 父分类ID，顶级分类为null |
| icon | string | 分类图标 |
| color | string | 分类颜色 |
| sort_order | integer | 排序顺序 |
| is_system | boolean | 是否系统分类 |
| description | string | 分类描述 |
| created_at | string | 创建时间 |
| updated_at | string | 更新时间 |

**成功响应**
```json
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "category_id": 1,
    "book_id": 1,
    "name": "日常购物",
    "type": "expense",
    "parent_id": null,
    "icon": "shopping-bag",
    "color": "#4CAF50",
    "sort_order": 1,
    "is_system": false,
    "description": "日常购物消费",
    "created_at": "2023-01-05T12:30:00Z",
    "updated_at": "2023-01-05T12:30:00Z"
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误 | 检查请求参数是否符合要求，包括类型、格式、必填项等 |
| 400 | 请求参数错误：同一账本内同类型分类名称已存在 | 更换分类名称 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有创建分类的权限 |
| 404 | 账本不存在 | 检查book_id是否正确 |
| 404 | 父分类不存在 | 检查parent_id是否正确 |
| 3001 | 分类创建失败 | 检查分类参数是否正确，特别是名称、类型等 |

**本地缓存策略**：创建成功后，本地缓存该分类信息，缓存时间24小时

### 创建账本

**功能描述**：创建新的账本，用于记录家庭或个人的收支情况

**请求方法**：POST
**URL路径**：/api/v1/books
**权限要求**：已登录用户
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Content-Type | string | 是 | 固定为application/json |
| Authorization | string | 是 | Bearer JWT令牌 |

**请求参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| name | string | 是 | - | 账本名称 | 1-50个字符，用户名下唯一 |
| description | string | 否 | "" | 账本描述 | 0-200个字符 |
| currency | string | 否 | "CNY" | 货币类型 | 有效的货币代码 |
| start_date | string | 是 | - | 开始日期 | YYYY-MM-DD格式 |
| is_default | boolean | 否 | false | 是否默认账本 | 布尔值 |
| auto_backup | boolean | 否 | false | 是否自动备份 | 布尔值 |
| template_id | integer | 否 | null | 模板ID | 正整数或null，模板必须存在 |

**请求示例**
```http
POST /api/v1/books
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "name": "我的家庭账本",
  "description": "记录家庭日常收支",
  "currency": "CNY",
  "start_date": "2023-01-01",
  "is_default": true,
  "auto_backup": true,
  "template_id": null
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| book_id | integer | 账本ID |
| name | string | 账本名称 |
| description | string | 账本描述 |
| currency | string | 货币类型 |
| start_date | string | 开始日期 |
| is_default | boolean | 是否默认账本 |
| auto_backup | boolean | 是否自动备份 |
| owner_id | integer | 所有者ID |
| created_at | string | 创建时间 |
| updated_at | string | 更新时间 |

**成功响应**
```json
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "book_id": 1,
    "name": "我的家庭账本",
    "description": "记录家庭日常收支",
    "currency": "CNY",
    "start_date": "2023-01-01",
    "is_default": true,
    "auto_backup": true,
    "owner_id": 1,
    "created_at": "2023-01-05T12:30:00Z",
    "updated_at": "2023-01-05T12:30:00Z"
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误 | 检查请求参数是否符合要求，包括类型、格式、必填项等 |
| 400 | 请求参数错误：用户名下账本名称已存在 | 更换账本名称 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 3002 | 账本创建失败 | 检查账本参数是否正确，特别是名称、日期等 |

**本地缓存策略**：创建成功后，本地缓存该账本信息，缓存时间24小时

### 获取操作日志

**功能描述**：获取系统操作日志，用于审计和监控用户操作

**请求方法**：GET
**URL路径**：/api/v1/logs/operation
**权限要求**：账本管理员或所有者
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 | 验证规则 |
|--------|------|------|--------|------|----------|
| book_id | integer | 否 | - | 账本ID | 正整数，账本必须存在且用户有访问权限 |
| start_date | string | 否 | - | 开始日期 | YYYY-MM-DD格式 |
| end_date | string | 否 | - | 结束日期 | YYYY-MM-DD格式，必须大于等于start_date |
| page | integer | 否 | 1 | 页码 | 正整数 |
| page_size | integer | 否 | 20 | 每页条数 | 1-100之间的整数 |
| action_type | string | 否 | - | 操作类型 | 枚举值：create, update, delete, login, logout |
| entity_type | string | 否 | - | 实体类型 | 有效的实体类型字符串 |
| user_id | integer | 否 | - | 用户ID | 正整数 |

**请求示例**
```http
GET /api/v1/logs/operation?book_id=1&start_date=2023-01-01&end_date=2023-01-31&page=1&page_size=20&action_type=create
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| total | integer | 总记录数 |
| page | integer | 当前页码 |
| page_size | integer | 每页条数 |
| logs | array | 日志列表 |
| logs[].log_id | integer | 日志ID |
| logs[].book_id | integer | 账本ID（可选） |
| logs[].user_id | integer | 用户ID |
| logs[].username | string | 用户名 |
| logs[].action_type | string | 操作类型：create, update, delete, login, logout |
| logs[].entity_type | string | 实体类型：transaction, account, category, etc. |
| logs[].entity_id | integer | 实体ID |
| logs[].details | string | 详细信息 |
| logs[].ip_address | string | IP地址 |
| logs[].user_agent | string | 用户代理 |
| logs[].created_at | string | 创建时间 |

**成功响应**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total": 150,
    "page": 1,
    "page_size": 20,
    "logs": [
      {
        "log_id": 1,
        "book_id": 1,
        "user_id": 1,
        "username": "张三",
        "action_type": "create",
        "entity_type": "transaction",
        "entity_id": 101,
        "details": "创建交易记录：日常购物 100元",
        "ip_address": "192.168.1.1",
        "user_agent": "Mozilla/5.0...",
        "created_at": "2023-01-05T12:30:00Z"
      }
    ]
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误 | 检查请求参数是否符合要求，包括类型、格式、必填项等 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 403 | 权限不足 | 检查当前用户是否有查看操作日志的权限 |
| 3005 | 日志查询失败 | 检查查询参数是否正确，特别是日期范围等 |

**本地缓存策略**：操作日志不缓存，每次请求重新查询

### 获取系统信息

**功能描述**：获取系统信息，包括版本、运行环境、数据库状态等

**请求方法**：GET
**URL路径**：/api/v1/system/info
**权限要求**：已登录用户
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**请求示例**
```http
GET /api/v1/system/info
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| system | object | 系统信息 |
| system.version | string | 系统版本 |
| system.build_date | string | 构建日期 |
| system.environment | string | 运行环境 |
| system.uptime | string | 运行时间 |
| system.database | object | 数据库信息 |
| system.database.type | string | 数据库类型 |
| system.database.version | string | 数据库版本 |
| system.database.size | string | 数据库大小 |
| system.database.status | string | 数据库状态 |
| system.storage | object | 存储信息 |
| system.storage.total | string | 总存储容量 |
| system.storage.used | string | 已使用存储容量 |
| system.storage.free | string | 剩余存储容量 |
| system.server | object | 服务器信息 |
| system.server.name | string | 服务器名称 |
| system.server.os | string | 操作系统 |
| system.server.memory | object | 内存信息 |
| system.server.memory.total | string | 总内存容量 |
| system.server.memory.used | string | 已使用内存容量 |
| system.server.memory.free | string | 剩余内存容量 |

**成功响应**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "system": {
      "version": "1.0.0",
      "build_date": "2023-01-01T00:00:00Z",
      "environment": "production",
      "uptime": "30d 5h 12m",
      "database": {
        "type": "SQLite",
        "version": "3.36.0",
        "size": "25.6MB",
        "status": "healthy"
      },
      "storage": {
        "total": "10GB",
        "used": "2.5GB",
        "free": "7.5GB"
      },
      "server": {
        "name": "family-bill-server",
        "os": "Linux Ubuntu 20.04",
        "memory": {
          "total": "8GB",
          "used": "2.4GB",
          "free": "5.6GB"
        }
      }
    }
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 500 | 服务器内部错误 | 请稍后重试，或联系系统管理员 |

**本地缓存策略**：系统信息本地缓存，缓存时间1小时

### 获取通知列表

**功能描述**：获取用户通知列表，包括系统通知、预算提醒等

**请求方法**：GET
**URL路径**：/api/v1/notifications
**权限要求**：已登录用户
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
| status | string | 否 | - | 通知状态 | 枚举值：read, unread |
| type | string | 否 | - | 通知类型 | 枚举值：budget_alert, reminder, system, shared |

**请求示例**
```http
GET /api/v1/notifications?page=1&page_size=20&status=unread
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| total | integer | 总记录数 |
| page | integer | 当前页码 |
| page_size | integer | 每页条数 |
| unread_count | integer | 未读通知数 |
| notifications | array | 通知列表 |
| notifications[].notification_id | integer | 通知ID |
| notifications[].user_id | integer | 用户ID |
| notifications[].title | string | 通知标题 |
| notifications[].content | string | 通知内容 |
| notifications[].type | string | 通知类型：budget_alert, reminder, system, shared |
| notifications[].is_read | boolean | 是否已读 |
| notifications[].is_pinned | boolean | 是否置顶 |
| notifications[].related_id | integer | 相关实体ID |
| notifications[].created_at | string | 创建时间 |

**成功响应**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total": 10,
    "page": 1,
    "page_size": 20,
    "unread_count": 3,
    "notifications": [
      {
        "notification_id": 1,
        "user_id": 1,
        "title": "预算提醒",
        "content": "您的月度餐饮预算已使用80%",
        "type": "budget_alert",
        "is_read": false,
        "is_pinned": false,
        "related_id": 1,
        "created_at": "2023-01-05T12:30:00Z"
      }
    ]
  }
}
```

**错误响应**
| 错误码 | 描述 | 解决方案建议 |
|-------|------|--------------|
| 400 | 请求参数错误 | 检查请求参数是否符合要求，包括类型、格式、必填项等 |
| 401 | 未授权，需要登录 | 请先登录获取有效的认证令牌 |
| 500 | 服务器内部错误 | 请稍后重试，或联系系统管理员 |

**本地缓存策略**：通知列表本地缓存，缓存时间5分钟

## 数据模型

### 用户设置模型

```javascript
{
  "settings_id": 1,          // 设置ID
  "user_id": 1,              // 用户ID
  "theme": "light",         // 主题：light, dark, system
  "language": "zh-CN",      // 语言
  "date_format": "YYYY-MM-DD", // 日期格式
  "time_format": "24h",     // 时间格式：12h, 24h
  "currency_format": "symbol_before", // 货币格式
  "decimal_places": 2,       // 小数位数
  "auto_sync": true,         // 是否自动同步
  "notifications": {         // 通知设置
    "email": true,           // 邮件通知
    "push": true,            // 推送通知
    "budget_alerts": true,   // 预算提醒
    "reminders": true,       // 账单提醒
    "updates": true          // 更新通知
  },
  "created_at": "2023-01-01T00:00:00Z", // 创建时间
  "updated_at": "2023-01-05T12:30:00Z"  // 更新时间
}
```

### 分类模型

```javascript
{
  "category_id": 1,          // 分类ID
  "book_id": 1,              // 账本ID
  "name": "日常购物",         // 分类名称
  "type": "expense",        // 类型：income, expense, transfer
  "parent_id": null,         // 父分类ID
  "icon": "shopping-bag",   // 图标
  "color": "#4CAF50",       // 颜色
  "sort_order": 1,           // 排序顺序
  "is_system": false,        // 是否系统分类
  "description": "日常购物消费", // 描述
  "created_at": "2023-01-05T12:30:00Z", // 创建时间
  "updated_at": "2023-01-05T12:30:00Z"  // 更新时间
}
```

### 账本模型

```javascript
{
  "book_id": 1,              // 账本ID
  "name": "我的家庭账本",     // 账本名称
  "description": "记录家庭日常收支", // 描述
  "currency": "CNY",        // 默认货币
  "start_date": "2023-01-01", // 开始日期
  "is_default": true,        // 是否默认账本
  "auto_backup": true,       // 是否自动备份
  "owner_id": 1,             // 所有者ID
  "created_at": "2023-01-05T12:30:00Z", // 创建时间
  "updated_at": "2023-01-05T12:30:00Z"  // 更新时间
}
```

### 系统日志模型

```javascript
{
  "log_id": 1,               // 日志ID
  "book_id": 1,              // 账本ID（可选）
  "user_id": 1,              // 用户ID
  "username": "张三",        // 用户名
  "action_type": "create",  // 操作类型
  "entity_type": "transaction", // 实体类型
  "entity_id": 101,          // 实体ID
  "details": "创建交易记录：日常购物 100元", // 详细信息
  "ip_address": "192.168.1.1", // IP地址
  "user_agent": "Mozilla/5.0...", // 用户代理
  "created_at": "2023-01-05T12:30:00Z" // 创建时间
}
```

### 通知模型

```javascript
{
  "notification_id": 1,      // 通知ID
  "user_id": 1,              // 用户ID
  "title": "预算提醒",       // 标题
  "content": "您的月度餐饮预算已使用80%", // 内容
  "type": "budget_alert",   // 类型
  "is_read": false,          // 是否已读
  "is_pinned": false,        // 是否置顶
  "related_id": 1,           // 相关实体ID
  "created_at": "2023-01-05T12:30:00Z" // 创建时间
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
| 3000 | 设置更新失败 | 检查设置参数是否正确，特别是主题、语言等枚举值 |
| 3001 | 分类创建失败 | 检查分类参数是否正确，特别是名称、类型等 |
| 3002 | 账本创建失败 | 检查账本参数是否正确，特别是名称、日期等 |
| 3003 | 头像上传失败 | 检查头像文件格式和大小是否符合要求 |
| 3004 | 密码修改失败 | 检查旧密码是否正确，新密码是否符合复杂度要求 |
| 3005 | 日志查询失败 | 检查查询参数是否正确，特别是日期范围等 |
| 3006 | 分类删除失败（已使用） | 该分类已被交易记录使用，无法删除 |
| 3007 | 账本删除失败（有成员） | 该账本还有其他成员，无法删除 |
| 3008 | 通知设置更新失败 | 检查通知设置参数是否正确 |
| 3009 | 系统配置更新失败 | 检查系统配置参数是否正确 |
| 3010 | 版本检查失败 | 检查网络连接，或稍后重试 |