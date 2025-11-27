## 模块概述

数据同步与备份模块确保多端数据一致性和数据安全，提供数据同步、备份、恢复和导出功能。本模块提供了完整的数据同步机制、数据备份与恢复、数据导入导出等功能的API接口。

### 设计原则

1. **RESTful规范**：所有API接口严格遵循RESTful设计规范，使用标准HTTP方法和资源导向的URL路径
2. **本地缓存优先**：暂不集成远程缓存机制，如需使用缓存功能，严格限定为本地缓存方案
3. **完整的请求验证**：所有API接口包含严格的请求参数验证，确保数据完整性和安全性
4. **统一的错误处理**：采用统一的错误响应格式和错误码体系，便于客户端处理
5. **清晰的响应格式**：所有API响应采用标准JSON格式，包含明确的数据结构和状态信息
6. **可扩展性**：API设计考虑未来功能扩展，预留合理的扩展点

### 适用场景

- 多设备用户需要同步记账数据
- 用户需要定期备份重要财务数据
- 用户需要从备份中恢复数据
- 用户需要导入外部记账数据
- 用户需要导出数据进行离线分析
- 用户需要解决数据冲突问题

## 接口清单

<!-- tabs:start -->
<!-- tab:数据同步 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/sync/status`](#获取同步状态) | `GET` | 获取同步状态 |
| [`/api/v1/sync/data`](#同步数据到服务器) | `POST` | 同步数据到服务器 |
| [`/api/v1/sync/pull`](#从服务器拉取最新数据) | `GET` | 从服务器拉取最新数据 |
| [`/api/v1/sync/conflict`](#获取数据冲突列表) | `GET` | 获取数据冲突列表 |
| [`/api/v1/sync/conflict/:id`](#解决数据冲突) | `PUT` | 解决数据冲突 |
| [`/api/v1/sync/config`](#获取同步配置) | `GET` | 获取同步配置 |
| [`/api/v1/sync/config`](#更新同步配置) | `PUT` | 更新同步配置 |

<!-- tab:数据备份 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/backup`](#获取备份列表) | `GET` | 获取备份列表 |
| [`/api/v1/backup`](#创建手动备份) | `POST` | 创建手动备份 |
| [`/api/v1/backup/:id`](#获取备份详情) | `GET` | 获取备份详情 |
| [`/api/v1/backup/:id`](#删除备份) | `DELETE` | 删除备份 |
| [`/api/v1/backup/auto-config`](#获取自动备份配置) | `GET` | 获取自动备份配置 |
| [`/api/v1/backup/auto-config`](#设置自动备份配置) | `PUT` | 设置自动备份配置 |

<!-- tab:数据恢复 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/restore/:backupId`](#从备份恢复数据) | `POST` | 从备份恢复数据 |
| [`/api/v1/restore/status/:taskId`](#获取恢复任务状态) | `GET` | 获取恢复任务状态 |

<!-- tab:数据导入导出 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/export/data`](#导出记账数据) | `GET` | 导出记账数据 |
| [`/api/v1/export/transactions`](#导出收支记录) | `GET` | 导出收支记录 |
| [`/api/v1/export/accounts`](#导出账户数据) | `GET` | 导出账户数据 |
| [`/api/v1/export/categories`](#导出分类数据) | `GET` | 导出分类数据 |

<!-- tab:数据迁移 -->
| 接口路径 | 方法 | 功能描述 |
|---------|------|--------|
| [`/api/v1/migrate/import`](#导入外部记账数据) | `POST` | 导入外部记账数据 |
| [`/api/v1/migrate/templates`](#获取导入模板列表) | `GET` | 获取导入模板列表 |
| [`/api/v1/migrate/preview`](#预览导入数据) | `POST` | 预览导入数据 |

<!-- tabs:end -->

## 详细接口说明

### 获取同步状态

**功能描述**：获取指定账本的同步状态信息

**请求方法**：GET
**URL路径**：/api/v1/sync/status
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
GET /api/v1/sync/status?book_id=1
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 同步状态信息 |
| data.last_sync_time | string | 最后同步时间 |
| data.sync_status | string | 同步状态：synced, syncing, conflict, error |
| data.pending_changes | integer | 待同步变更数量 |
| data.conflicts_count | integer | 冲突数量 |
| data.client_version | string | 客户端版本 |
| data.server_version | string | 服务器版本 |
| data.last_full_sync_time | string | 最后全量同步时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "last_sync_time": "2023-01-05T12:30:00Z",
    "sync_status": "synced",  // synced, syncing, conflict, error
    "pending_changes": 0,
    "conflicts_count": 0,
    "client_version": "1.0.0",
    "server_version": "1.0.0",
    "last_full_sync_time": "2023-01-01T00:00:00Z"
  }
}
```

### 同步数据到服务器

**功能描述**：将客户端本地变更数据同步到服务器

**请求方法**：POST
**URL路径**：/api/v1/sync/data
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
| client_timestamp | string | 是 | - | 客户端时间戳 |
| client_version | string | 是 | - | 客户端版本 |
| changes | object | 是 | - | 待同步的变更数据 |
| changes.transactions | object | 否 | - | 收支记录变更 |
| changes.transactions.created | array | 否 | [] | 新增的收支记录 |
| changes.transactions.updated | array | 否 | [] | 更新的收支记录 |
| changes.transactions.deleted | array | 否 | [] | 删除的收支记录ID列表 |
| changes.accounts | object | 否 | - | 账户变更 |
| changes.accounts.created | array | 否 | [] | 新增的账户 |
| changes.accounts.updated | array | 否 | [] | 更新的账户 |
| changes.accounts.deleted | array | 否 | [] | 删除的账户ID列表 |
| changes.categories | object | 否 | - | 分类变更 |
| changes.categories.created | array | 否 | [] | 新增的分类 |
| changes.categories.updated | array | 否 | [] | 更新的分类 |
| changes.categories.deleted | array | 否 | [] | 删除的分类ID列表 |

**请求示例**
```http
POST /api/v1/sync/data
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "client_timestamp": "2023-01-05T12:30:00Z",
  "client_version": "1.0.0",
  "changes": {
    "transactions": {
      "created": [
        {
          "transaction_id": "local_123",
          "type": "expense",
          "amount": 100.50,
          "account_id": 1,
          "category_id": 5,
          "date": "2023-01-05",
          "created_at": "2023-01-05T12:30:00Z",
          "sync_version": 1
        }
      ],
      "updated": [
        // 更新的记录...
      ],
      "deleted": [
        // 删除的记录ID...
      ]
    },
    "accounts": {
      "created": [
        // 新增的账户...
      ],
      "updated": [
        // 更新的账户...
      ]
    }
  }
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 同步结果信息 |
| data.sync_id | string | 同步任务ID |
| data.status | string | 同步状态：success, conflict, error |
| data.server_timestamp | string | 服务器时间戳 |
| data.applied_changes | object | 已应用的变更 |
| data.applied_changes.created | array | 已创建的记录映射 |
| data.applied_changes.updated | array | 已更新的记录映射 |
| data.applied_changes.deleted | array | 已删除的记录ID列表 |
| data.conflicts | array | 冲突列表 |
| data.pending_changes_count | integer | 待处理变更数量 |

**响应示例**
```json
{
  "code": 200,
  "message": "同步成功",
  "data": {
    "sync_id": "sync_123",
    "status": "success",
    "server_timestamp": "2023-01-05T12:30:05Z",
    "applied_changes": {
      "created": [
        {
          "local_id": "local_123",
          "server_id": 101,
          "sync_version": 1
        }
      ],
      "updated": [],
      "deleted": []
    },
    "conflicts": [],
    "pending_changes_count": 0
  }
}
```

**错误响应**
```json
{
  "code": 1001,
  "message": "数据冲突",
  "data": {
    "conflicts": [
      {
        "entity_type": "transaction",
        "entity_id": 101,
        "local_version": 2,
        "server_version": 3,
        "conflicting_fields": ["amount", "memo"]
      }
    ]
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 创建手动备份

**功能描述**：创建手动备份任务

**请求方法**：POST
**URL路径**：/api/v1/backup
**权限要求**：账本管理员
**限流策略**：每分钟10次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |
| Content-Type | string | 是 | application/json |

**请求体**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| book_id | integer | 是 | - | 账本ID |
| name | string | 是 | - | 备份名称 |
| description | string | 否 | - | 备份描述 |
| type | string | 是 | full | 备份类型：full, incremental |
| include_attachments | boolean | 否 | false | 是否包含附件 |

**请求示例**
```http
POST /api/v1/backup
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "name": "2023年1月手动备份",
  "description": "包含1月份所有数据",
  "type": "full",  // full, incremental
  "include_attachments": true
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 备份任务信息 |
| data.backup_id | string | 备份ID |
| data.task_id | string | 任务ID |
| data.name | string | 备份名称 |
| data.status | string | 状态：pending, processing, completed, failed |
| data.created_at | string | 创建时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "备份任务已创建",
  "data": {
    "backup_id": "backup_123",
    "task_id": "task_456",
    "name": "2023年1月手动备份",
    "status": "pending",  // pending, processing, completed, failed
    "created_at": "2023-01-05T12:30:00Z"
  }
}
```

**错误响应**
```json
{
  "code": 1002,
  "message": "备份创建失败",
  "data": {
    "error_details": "存储空间不足"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 从备份恢复数据

**功能描述**：从指定备份恢复数据到账本

**请求方法**：POST
**URL路径**：/api/v1/restore/:backupId
**权限要求**：账本管理员
**限流策略**：每分钟5次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |
| Content-Type | string | 是 | application/json |

**路径参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| backupId | string | 是 | - | 备份ID |

**请求体**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| book_id | integer | 是 | - | 账本ID |
| restore_type | string | 是 | overwrite | 恢复类型：overwrite, merge |
| include_attachments | boolean | 否 | false | 是否包含附件 |
| confirm | boolean | 是 | false | 是否确认恢复操作 |

**请求示例**
```http
POST /api/v1/restore/:backupId
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "restore_type": "overwrite",  // overwrite, merge
  "include_attachments": true,
  "confirm": true
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 恢复任务信息 |
| data.task_id | string | 恢复任务ID |
| data.backup_id | string | 备份ID |
| data.status | string | 状态：pending, processing, completed, failed |
| data.created_at | string | 创建时间 |
| data.estimated_time | string | 预计完成时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "恢复任务已创建",
  "data": {
    "task_id": "restore_123",
    "backup_id": "backup_123",
    "status": "pending",  // pending, processing, completed, failed
    "created_at": "2023-01-05T12:30:00Z",
    "estimated_time": "约5分钟"
  }
}
```

**错误响应**
```json
{
  "code": 1003,
  "message": "恢复失败",
  "data": {
    "error_details": "备份文件损坏"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 导出记账数据

**功能描述**：导出指定账本的记账数据

**请求方法**：GET
**URL路径**：/api/v1/export/data
**权限要求**：账本成员
**限流策略**：每分钟20次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| book_id | integer | 是 | - | 账本ID |
| format | string | 是 | excel | 导出格式：excel, csv, pdf |
| start_date | string | 否 | - | 开始日期，格式：YYYY-MM-DD |
| end_date | string | 否 | - | 结束日期，格式：YYYY-MM-DD |
| include_transactions | boolean | 否 | true | 是否包含收支记录 |
| include_accounts | boolean | 否 | true | 是否包含账户数据 |
| include_categories | boolean | 否 | true | 是否包含分类数据 |
| include_attachments | boolean | 否 | false | 是否包含附件 |

**请求示例**
```http
GET /api/v1/export/data?book_id=1&format=excel&start_date=2023-01-01&end_date=2023-01-31&include_transactions=true&include_accounts=true&include_categories=true
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 导出结果信息 |
| data.file_url | string | 临时文件URL |
| data.file_name | string | 文件名 |
| data.file_size | string | 文件大小 |
| data.expires_in | integer | 文件URL过期时间（秒） |
| data.export_time | string | 导出时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "导出成功",
  "data": {
    "file_url": "临时文件URL",
    "file_name": "家庭记账_202301_完整数据.xlsx",
    "file_size": "2.5MB",
    "expires_in": 3600,  // 文件URL过期时间（秒）
    "export_time": "2023-01-05T12:30:00Z"
  }
}
```

**错误响应**
```json
{
  "code": 1004,
  "message": "导出失败",
  "data": {
    "error_details": "导出格式不支持"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 导入外部记账数据

**功能描述**：从外部记账软件导入数据到账本

**请求方法**：POST
**URL路径**：/api/v1/migrate/import
**权限要求**：账本管理员
**限流策略**：每分钟10次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |
| Content-Type | string | 是 | multipart/form-data |

**表单数据**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| book_id | integer | 是 | - | 账本ID |
| provider | string | 是 | - | 外部记账软件标识 |
| file | file | 是 | - | 上传的导入文件 |
| mapping | string | 否 | - | 字段映射JSON字符串 |

**请求示例**
```http
POST /api/v1/migrate/import
Content-Type: multipart/form-data
Authorization: Bearer jwt_token_string

book_id=1
provider=mint  // 外部记账软件标识
file=...  // 上传的导入文件
mapping=...  // 字段映射JSON字符串
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 导入任务信息 |
| data.task_id | string | 导入任务ID |
| data.status | string | 状态：pending, processing, completed, failed |
| data.created_at | string | 创建时间 |
| data.estimated_time | string | 预计完成时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "导入任务已创建",
  "data": {
    "task_id": "import_123",
    "status": "pending",  // pending, processing, completed, failed
    "created_at": "2023-01-05T12:30:00Z",
    "estimated_time": "约10分钟"
  }
}
```

**错误响应**
```json
{
  "code": 1005,
  "message": "导入失败",
  "data": {
    "error_details": "文件格式不支持"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 设置自动备份配置

**功能描述**：设置或更新自动备份配置

**请求方法**：PUT
**URL路径**：/api/v1/backup/auto-config
**权限要求**：账本管理员
**限流策略**：每分钟10次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |
| Content-Type | string | 是 | application/json |

**请求体**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| book_id | integer | 是 | - | 账本ID |
| enabled | boolean | 是 | false | 是否启用自动备份 |
| frequency | string | 是 | daily | 备份频率：daily, weekly, monthly |
| time | string | 是 | 02:00 | 备份执行时间，格式：HH:MM |
| retention_days | integer | 否 | 30 | 备份保留天数 |
| include_attachments | boolean | 否 | false | 是否包含附件 |
| notification | boolean | 否 | true | 是否发送备份通知 |

**请求示例**
```http
PUT /api/v1/backup/auto-config
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "book_id": 1,
  "enabled": true,
  "frequency": "daily",  // daily, weekly, monthly
  "time": "02:00",  // 备份执行时间
  "retention_days": 30,
  "include_attachments": true,
  "notification": true
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 备份配置信息 |
| data.config_id | string | 配置ID |
| data.enabled | boolean | 是否启用自动备份 |
| data.frequency | string | 备份频率 |
| data.time | string | 备份执行时间 |
| data.retention_days | integer | 备份保留天数 |
| data.next_backup_time | string | 下次备份时间 |
| data.updated_at | string | 更新时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "配置更新成功",
  "data": {
    "config_id": "config_123",
    "enabled": true,
    "frequency": "daily",
    "time": "02:00",
    "retention_days": 30,
    "next_backup_time": "2023-01-06T02:00:00Z",
    "updated_at": "2023-01-05T12:30:00Z"
  }
}
```

**错误响应**
```json
{
  "code": 400,
  "message": "请求参数错误",
  "data": {
    "error_details": "备份频率无效"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 从服务器拉取最新数据

**功能描述**：从服务器拉取最新数据到客户端

**请求方法**：GET
**URL路径**：/api/v1/sync/pull
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
| last_sync_time | string | 否 | - | 最后同步时间 |
| client_version | string | 是 | - | 客户端版本 |

**请求示例**
```http
GET /api/v1/sync/pull?book_id=1&last_sync_time=2023-01-05T12:30:00Z&client_version=1.0.0
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 同步数据信息 |
| data.server_timestamp | string | 服务器时间戳 |
| data.changes | object | 待同步的变更数据 |
| data.sync_version | integer | 同步版本号 |
| data.conflicts | array | 冲突列表 |

**响应示例**
```json
{
  "code": 200,
  "message": "拉取成功",
  "data": {
    "server_timestamp": "2023-01-05T12:30:05Z",
    "changes": {
      "transactions": {
        "created": [],
        "updated": [],
        "deleted": []
      },
      "accounts": {
        "created": [],
        "updated": [],
        "deleted": []
      }
    },
    "sync_version": 5,
    "conflicts": []
  }
}
```

**错误响应**
```json
{
  "code": 1000,
  "message": "同步失败",
  "data": {
    "error_details": "网络连接错误"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 获取数据冲突列表

**功能描述**：获取指定账本的数据冲突列表

**请求方法**：GET
**URL路径**：/api/v1/sync/conflict
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
| page_size | integer | 否 | 20 | 每页条数 |

**请求示例**
```http
GET /api/v1/sync/conflict?book_id=1&page=1&page_size=20
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 冲突列表信息 |
| data.total | integer | 总冲突数 |
| data.list | array | 冲突列表 |
| data.list[].conflict_id | string | 冲突ID |
| data.list[].entity_type | string | 实体类型 |
| data.list[].entity_id | integer | 实体ID |
| data.list[].local_version | integer | 本地版本 |
| data.list[].server_version | integer | 服务器版本 |
| data.list[].conflicting_fields | array | 冲突字段列表 |
| data.list[].created_at | string | 创建时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total": 1,
    "list": [
      {
        "conflict_id": "conflict_123",
        "entity_type": "transaction",
        "entity_id": 101,
        "local_version": 2,
        "server_version": 3,
        "conflicting_fields": ["amount", "memo"],
        "created_at": "2023-01-05T12:30:00Z"
      }
    ]
  }
}
```

**错误响应**
```json
{
  "code": 400,
  "message": "请求参数错误",
  "data": {
    "error_details": "账本ID不能为空"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 解决数据冲突

**功能描述**：解决指定的数据冲突

**请求方法**：PUT
**URL路径**：/api/v1/sync/conflict/:id
**权限要求**：账本成员
**限流策略**：每分钟30次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |
| Content-Type | string | 是 | application/json |

**路径参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| id | string | 是 | - | 冲突ID |

**请求体**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| resolution_strategy | string | 是 | - | 解决策略：server_wins, client_wins, manual |
| manual_data | object | 否 | - | 手动解决的数据（当resolution_strategy为manual时必填） |

**请求示例**
```http
PUT /api/v1/sync/conflict/conflict_123
Content-Type: application/json
Authorization: Bearer jwt_token_string

{
  "resolution_strategy": "server_wins"
}
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 解决结果 |
| data.conflict_id | string | 冲突ID |
| data.status | string | 解决状态：resolved, failed |
| data.resolved_at | string | 解决时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "冲突已解决",
  "data": {
    "conflict_id": "conflict_123",
    "status": "resolved",
    "resolved_at": "2023-01-05T12:35:00Z"
  }
}
```

**错误响应**
```json
{
  "code": 1001,
  "message": "冲突解决失败",
  "data": {
    "error_details": "冲突已被解决"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 获取同步配置

**功能描述**：获取指定账本的同步配置

**请求方法**：GET
**URL路径**：/api/v1/sync/config
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
GET /api/v1/sync/config?book_id=1
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 同步配置信息 |
| data.sync_config_id | string | 配置ID |
| data.book_id | integer | 账本ID |
| data.auto_sync | boolean | 是否自动同步 |
| data.sync_frequency | string | 同步频率：realtime, hourly, daily |
| data.sync_when_offline | boolean | 离线时是否缓存 |
| data.conflict_strategy | string | 冲突解决策略：server_wins, client_wins, ask_user |
| data.sync_scope | object | 同步范围 |
| data.last_updated | string | 最后更新时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "sync_config_id": "config_123",
    "book_id": 1,
    "auto_sync": true,
    "sync_frequency": "realtime",
    "sync_when_offline": true,
    "conflict_strategy": "server_wins",
    "sync_scope": {
      "transactions": true,
      "accounts": true,
      "categories": true,
      "attachments": true
    },
    "last_updated": "2023-01-05T12:30:00Z"
  }
}
```

**错误响应**
```json
{
  "code": 404,
  "message": "配置不存在",
  "data": {
    "error_details": "未找到该账本的同步配置"
  }
}
```

**本地缓存策略**：缓存10分钟

### 获取备份列表

**功能描述**：获取指定账本的备份列表

**请求方法**：GET
**URL路径**：/api/v1/backup
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
| page_size | integer | 否 | 20 | 每页条数 |
| type | string | 否 | - | 备份类型：full, incremental |
| is_auto | boolean | 否 | - | 是否自动备份 |

**请求示例**
```http
GET /api/v1/backup?book_id=1&page=1&page_size=20&type=full
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 备份列表信息 |
| data.total | integer | 总备份数 |
| data.list | array | 备份列表 |
| data.list[].backup_id | string | 备份ID |
| data.list[].book_id | integer | 账本ID |
| data.list[].name | string | 备份名称 |
| data.list[].description | string | 备份描述 |
| data.list[].type | string | 备份类型 |
| data.list[].size | string | 备份文件大小 |
| data.list[].status | string | 状态 |
| data.list[].is_auto | boolean | 是否自动备份 |
| data.list[].created_at | string | 创建时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "total": 5,
    "list": [
      {
        "backup_id": "backup_123",
        "book_id": 1,
        "name": "2023年1月手动备份",
        "description": "包含1月份所有数据",
        "type": "full",
        "size": "5.2MB",
        "status": "completed",
        "is_auto": false,
        "created_at": "2023-01-05T12:30:00Z"
      }
    ]
  }
}
```

**错误响应**
```json
{
  "code": 400,
  "message": "请求参数错误",
  "data": {
    "error_details": "账本ID不能为空"
  }
}
```

**本地缓存策略**：缓存5分钟

### 获取备份详情

**功能描述**：获取指定备份的详细信息

**请求方法**：GET
**URL路径**：/api/v1/backup/:id
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| id | string | 是 | - | 备份ID |

**请求示例**
```http
GET /api/v1/backup/backup_123
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 备份详情信息 |
| data.backup_id | string | 备份ID |
| data.book_id | integer | 账本ID |
| data.name | string | 备份名称 |
| data.description | string | 备份描述 |
| data.type | string | 备份类型 |
| data.size | string | 备份文件大小 |
| data.data_count | object | 数据统计 |
| data.created_at | string | 创建时间 |
| data.created_by | integer | 创建用户ID |
| data.status | string | 状态 |
| data.is_auto | boolean | 是否自动备份 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "backup_id": "backup_123",
    "book_id": 1,
    "name": "2023年1月手动备份",
    "description": "包含1月份所有数据",
    "type": "full",
    "size": "5.2MB",
    "data_count": {
      "transactions": 1500,
      "accounts": 10,
      "categories": 50,
      "attachments": 20
    },
    "created_at": "2023-01-05T12:30:00Z",
    "created_by": 1,
    "status": "completed",
    "is_auto": false
  }
}
```

**错误响应**
```json
{
  "code": 404,
  "message": "备份不存在",
  "data": {
    "error_details": "未找到该备份记录"
  }
}
```

**本地缓存策略**：缓存10分钟

### 删除备份

**功能描述**：删除指定的备份

**请求方法**：DELETE
**URL路径**：/api/v1/backup/:id
**权限要求**：账本管理员
**限流策略**：每分钟30次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| id | string | 是 | - | 备份ID |

**请求示例**
```http
DELETE /api/v1/backup/backup_123
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 删除结果 |
| data.backup_id | string | 备份ID |
| data.deleted_at | string | 删除时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "删除成功",
  "data": {
    "backup_id": "backup_123",
    "deleted_at": "2023-01-05T12:35:00Z"
  }
}
```

**错误响应**
```json
{
  "code": 403,
  "message": "权限不足",
  "data": {
    "error_details": "只有账本管理员可以删除备份"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 获取自动备份配置

**功能描述**：获取指定账本的自动备份配置

**请求方法**：GET
**URL路径**：/api/v1/backup/auto-config
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
GET /api/v1/backup/auto-config?book_id=1
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 自动备份配置信息 |
| data.config_id | string | 配置ID |
| data.book_id | integer | 账本ID |
| data.enabled | boolean | 是否启用自动备份 |
| data.frequency | string | 备份频率 |
| data.time | string | 备份执行时间 |
| data.retention_days | integer | 备份保留天数 |
| data.include_attachments | boolean | 是否包含附件 |
| data.notification | boolean | 是否发送备份通知 |
| data.next_backup_time | string | 下次备份时间 |
| data.updated_at | string | 更新时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "config_id": "config_123",
    "book_id": 1,
    "enabled": true,
    "frequency": "daily",
    "time": "02:00",
    "retention_days": 30,
    "include_attachments": true,
    "notification": true,
    "next_backup_time": "2023-01-06T02:00:00Z",
    "updated_at": "2023-01-05T12:30:00Z"
  }
}
```

**错误响应**
```json
{
  "code": 404,
  "message": "配置不存在",
  "data": {
    "error_details": "未找到该账本的自动备份配置"
  }
}
```

**本地缓存策略**：缓存10分钟

### 获取恢复任务状态

**功能描述**：获取指定恢复任务的状态

**请求方法**：GET
**URL路径**：/api/v1/restore/status/:taskId
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**路径参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| taskId | string | 是 | - | 恢复任务ID |

**请求示例**
```http
GET /api/v1/restore/status/restore_123
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 恢复任务状态信息 |
| data.task_id | string | 恢复任务ID |
| data.backup_id | string | 备份ID |
| data.status | string | 状态：pending, processing, completed, failed |
| data.progress | integer | 进度百分比 |
| data.created_at | string | 创建时间 |
| data.updated_at | string | 更新时间 |
| data.error_message | string | 错误信息（失败时返回） |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "task_id": "restore_123",
    "backup_id": "backup_123",
    "status": "processing",
    "progress": 50,
    "created_at": "2023-01-05T12:30:00Z",
    "updated_at": "2023-01-05T12:32:00Z"
  }
}
```

**错误响应**
```json
{
  "code": 404,
  "message": "任务不存在",
  "data": {
    "error_details": "未找到该恢复任务"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 导出收支记录

**功能描述**：导出指定账本的收支记录

**请求方法**：GET
**URL路径**：/api/v1/export/transactions
**权限要求**：账本成员
**限流策略**：每分钟20次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| book_id | integer | 是 | - | 账本ID |
| format | string | 是 | excel | 导出格式：excel, csv, pdf |
| start_date | string | 否 | - | 开始日期，格式：YYYY-MM-DD |
| end_date | string | 否 | - | 结束日期，格式：YYYY-MM-DD |
| type | string | 否 | - | 交易类型：income, expense |

**请求示例**
```http
GET /api/v1/export/transactions?book_id=1&format=excel&start_date=2023-01-01&end_date=2023-01-31&type=expense
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 导出结果信息 |
| data.file_url | string | 临时文件URL |
| data.file_name | string | 文件名 |
| data.file_size | string | 文件大小 |
| data.expires_in | integer | 文件URL过期时间（秒） |
| data.export_time | string | 导出时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "导出成功",
  "data": {
    "file_url": "临时文件URL",
    "file_name": "家庭记账_202301_支出记录.xlsx",
    "file_size": "1.2MB",
    "expires_in": 3600,
    "export_time": "2023-01-05T12:30:00Z"
  }
}
```

**错误响应**
```json
{
  "code": 1004,
  "message": "导出失败",
  "data": {
    "error_details": "导出格式不支持"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 导出账户数据

**功能描述**：导出指定账本的账户数据

**请求方法**：GET
**URL路径**：/api/v1/export/accounts
**权限要求**：账本成员
**限流策略**：每分钟20次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| book_id | integer | 是 | - | 账本ID |
| format | string | 是 | excel | 导出格式：excel, csv, pdf |

**请求示例**
```http
GET /api/v1/export/accounts?book_id=1&format=excel
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 导出结果信息 |
| data.file_url | string | 临时文件URL |
| data.file_name | string | 文件名 |
| data.file_size | string | 文件大小 |
| data.expires_in | integer | 文件URL过期时间（秒） |
| data.export_time | string | 导出时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "导出成功",
  "data": {
    "file_url": "临时文件URL",
    "file_name": "家庭记账_账户数据.xlsx",
    "file_size": "0.5MB",
    "expires_in": 3600,
    "export_time": "2023-01-05T12:30:00Z"
  }
}
```

**错误响应**
```json
{
  "code": 1004,
  "message": "导出失败",
  "data": {
    "error_details": "导出格式不支持"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 导出分类数据

**功能描述**：导出指定账本的分类数据

**请求方法**：GET
**URL路径**：/api/v1/export/categories
**权限要求**：账本成员
**限流策略**：每分钟20次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| book_id | integer | 是 | - | 账本ID |
| format | string | 是 | excel | 导出格式：excel, csv, pdf |

**请求示例**
```http
GET /api/v1/export/categories?book_id=1&format=excel
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 导出结果信息 |
| data.file_url | string | 临时文件URL |
| data.file_name | string | 文件名 |
| data.file_size | string | 文件大小 |
| data.expires_in | integer | 文件URL过期时间（秒） |
| data.export_time | string | 导出时间 |

**响应示例**
```json
{
  "code": 200,
  "message": "导出成功",
  "data": {
    "file_url": "临时文件URL",
    "file_name": "家庭记账_分类数据.xlsx",
    "file_size": "0.3MB",
    "expires_in": 3600,
    "export_time": "2023-01-05T12:30:00Z"
  }
}
```

**错误响应**
```json
{
  "code": 1004,
  "message": "导出失败",
  "data": {
    "error_details": "导出格式不支持"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

### 获取导入模板列表

**功能描述**：获取支持的导入模板列表

**请求方法**：GET
**URL路径**：/api/v1/migrate/templates
**权限要求**：账本成员
**限流策略**：每分钟60次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |

**查询参数**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| provider | string | 否 | - | 外部记账软件标识 |

**请求示例**
```http
GET /api/v1/migrate/templates?provider=mint
Authorization: Bearer jwt_token_string
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | array | 模板列表 |
| data[].template_id | string | 模板ID |
| data[].provider | string | 外部记账软件标识 |
| data[].name | string | 模板名称 |
| data[].description | string | 模板描述 |
| data[].supported_formats | array | 支持的文件格式 |
| data[].download_url | string | 模板下载URL |

**响应示例**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": [
    {
      "template_id": "template_123",
      "provider": "mint",
      "name": "Mint导入模板",
      "description": "适用于Mint记账软件的导入模板",
      "supported_formats": ["csv", "excel"],
      "download_url": "模板下载URL"
    }
  ]
}
```

**错误响应**
```json
{
  "code": 400,
  "message": "请求参数错误",
  "data": {
    "error_details": "不支持该外部记账软件"
  }
}
```

**本地缓存策略**：缓存30分钟

### 预览导入数据

**功能描述**：预览导入的数据

**请求方法**：POST
**URL路径**：/api/v1/migrate/preview
**权限要求**：账本成员
**限流策略**：每分钟20次

**请求头**
| 头部名称 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| Authorization | string | 是 | Bearer JWT令牌 |
| Content-Type | string | 是 | multipart/form-data |

**表单数据**
| 参数名 | 类型 | 必填 | 默认值 | 描述 |
|--------|------|------|--------|------|
| book_id | integer | 是 | - | 账本ID |
| provider | string | 是 | - | 外部记账软件标识 |
| file | file | 是 | - | 上传的导入文件 |
| mapping | string | 否 | - | 字段映射JSON字符串 |

**请求示例**
```http
POST /api/v1/migrate/preview
Content-Type: multipart/form-data
Authorization: Bearer jwt_token_string

book_id=1
provider=mint
file=...  // 上传的导入文件
mapping=...  // 字段映射JSON字符串
```

**响应数据结构**
| 字段名 | 类型 | 描述 |
|--------|------|------|
| code | integer | 响应状态码 |
| message | string | 响应消息 |
| data | object | 预览结果信息 |
| data.total_count | integer | 总记录数 |
| data.valid_count | integer | 有效记录数 |
| data.invalid_count | integer | 无效记录数 |
| data.sample_data | array | 示例数据（前10条） |
| data.errors | array | 错误信息 |

**响应示例**
```json
{
  "code": 200,
  "message": "预览成功",
  "data": {
    "total_count": 100,
    "valid_count": 95,
    "invalid_count": 5,
    "sample_data": [
      {
        "date": "2023-01-01",
        "amount": 100.50,
        "category": "餐饮",
        "description": "午餐"
      }
    ],
    "errors": [
      {
        "line": 5,
        "message": "金额格式不正确"
      }
    ]
  }
}
```

**错误响应**
```json
{
  "code": 1005,
  "message": "预览失败",
  "data": {
    "error_details": "文件格式不支持"
  }
}
```

**本地缓存策略**：不缓存，每次请求都直接发送到服务器

## 数据模型

### 同步状态模型

```javascript
{
  "last_sync_time": "2023-01-05T12:30:00Z",  // 最后同步时间
  "sync_status": "synced",  // 同步状态：synced, syncing, conflict, error
  "pending_changes": 0,      // 待同步变更数
  "conflicts_count": 0,      // 冲突数
  "client_version": "1.0.0", // 客户端版本
  "server_version": "1.0.0", // 服务器版本
  "last_full_sync_time": "2023-01-01T00:00:00Z" // 最后完整同步时间
}
```

### 备份记录模型

```javascript
{
  "backup_id": "backup_123",        // 备份ID
  "book_id": 1,                    // 账本ID
  "name": "2023年1月手动备份",       // 备份名称
  "description": "包含1月份所有数据",  // 备份描述
  "type": "full",                 // 备份类型：full, incremental
  "size": "5.2MB",                // 备份文件大小
  "data_count": {
    "transactions": 1500,
    "accounts": 10,
    "categories": 50,
    "attachments": 20
  },
  "created_at": "2023-01-05T12:30:00Z", // 创建时间
  "created_by": 1,                 // 创建用户ID
  "status": "completed",          // 状态：pending, processing, completed, failed
  "is_auto": false                 // 是否自动备份
}
```

### 同步配置模型

```javascript
{
  "sync_config_id": "config_123",  // 配置ID
  "book_id": 1,                    // 账本ID
  "auto_sync": true,               // 是否自动同步
  "sync_frequency": "realtime",   // 同步频率：realtime, hourly, daily
  "sync_when_offline": true,       // 离线时是否缓存
  "conflict_strategy": "server_wins", // 冲突解决策略：server_wins, client_wins, ask_user
  "sync_scope": {
    "transactions": true,
    "accounts": true,
    "categories": true,
    "attachments": true
  },
  "last_updated": "2023-01-05T12:30:00Z" // 最后更新时间
}
```

### 数据冲突模型

```javascript
{
  "conflict_id": "conflict_123",   // 冲突ID
  "entity_type": "transaction",    // 实体类型
  "entity_id": 101,                // 实体ID
  "local_version": 2,              // 本地版本
  "server_version": 3,             // 服务器版本
  "local_data": {
    // 本地数据...
  },
  "server_data": {
    // 服务器数据...
  },
  "conflicting_fields": ["amount", "memo"], // 冲突字段
  "created_at": "2023-01-05T12:30:00Z"  // 创建时间
}
```

## 同步机制说明

### 增量同步

- 系统采用增量同步机制，只同步发生变化的数据
- 每个实体都有版本号，用于检测冲突
- 支持离线操作，网络恢复后自动同步

### 冲突解决策略

- **服务器优先**: 以服务器数据为准
- **客户端优先**: 以客户端数据为准
- **手动解决**: 提示用户选择保留的数据

### 同步范围

- 收支记录
- 账户信息
- 分类信息
- 标签信息
- 附件文件
- 预算数据

## 备份机制说明

### 备份类型

- **完整备份**: 包含所有数据的完整副本
- **增量备份**: 仅包含上次备份后变更的数据

### 备份存储位置

- 云端备份
- 本地备份（可选）
- 外部存储（可选）

### 备份保留策略

- 可配置保留天数
- 自动清理过期备份
- 重要备份可设为永久保留

## 错误码说明

| 错误码 | 描述 |
|-------|------|
| 400 | 请求参数错误 |
| 401 | 未授权，需要登录 |
| 403 | 权限不足 |
| 404 | 备份不存在 |
| 500 | 服务器内部错误 |
| 1000 | 同步失败 |
| 1001 | 数据冲突 |
| 1002 | 备份失败 |
| 1003 | 恢复失败 |
| 1004 | 导出失败 |
| 1005 | 导入失败 |
| 1006 | 存储空间不足 |
| 1007 | 文件格式不支持 |
| 1008 | 同步版本不兼容 |