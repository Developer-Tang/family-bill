# 数据同步与备份模块 API 文档

##  模块概述

数据同步与备份模块确保多端数据一致性和数据安全，提供数据同步、备份、恢复和导出功能。本模块提供了完整的数据同步机制、数据备份与恢复、数据导入导出等功能的API接口。

##  接口清单

| 功能模块 | 接口路径 | 方法 | 功能描述 |
|---------|---------|------|--------|
| **数据同步** | `/api/v1/sync/status` | `GET` | 获取同步状态 |
| | `/api/v1/sync/data` | `POST` | 同步数据到服务器 |
| | `/api/v1/sync/pull` | `GET` | 从服务器拉取最新数据 |
| | `/api/v1/sync/conflict` | `GET` | 获取数据冲突列表 |
| | `/api/v1/sync/conflict/:id` | `PUT` | 解决数据冲突 |
| **数据导出** | `/api/v1/export/data` | `GET` | 导出记账数据 |
| | `/api/v1/export/transactions` | `GET` | 导出收支记录 |
| | `/api/v1/export/accounts` | `GET` | 导出账户数据 |
| | `/api/v1/export/categories` | `GET` | 导出分类数据 |
| **数据备份** | `/api/v1/backup` | `GET` | 获取备份列表 |
| | `/api/v1/backup` | `POST` | 创建手动备份 |
| | `/api/v1/backup/:id` | `GET` | 获取备份详情 |
| | `/api/v1/backup/:id` | `DELETE` | 删除备份 |
| | `/api/v1/backup/auto-config` | `GET` | 获取自动备份配置 |
| | `/api/v1/backup/auto-config` | `PUT` | 设置自动备份配置 |
| **数据恢复** | `/api/v1/restore/:backupId` | `POST` | 从备份恢复数据 |
| | `/api/v1/restore/status/:taskId` | `GET` | 获取恢复任务状态 |
| **数据迁移** | `/api/v1/migrate/import` | `POST` | 导入外部记账数据 |
| | `/api/v1/migrate/templates` | `GET` | 获取导入模板列表 |
| | `/api/v1/migrate/preview` | `POST` | 预览导入数据 |
| **同步配置** | `/api/v1/sync/config` | `GET` | 获取同步配置 |
| | `/api/v1/sync/config` | `PUT` | 更新同步配置 |

##  详细接口说明

###  获取同步状态

#### 请求

```http
GET /api/v1/sync/status?book_id=1
Authorization: Bearer jwt_token_string
```

#### 响应

```
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

###  同步数据到服务器

#### 请求

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

#### 响应

```
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

###  创建手动备份

#### 请求

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

#### 响应

```
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

###  从备份恢复数据

#### 请求

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

#### 响应

```
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

###  导出记账数据

#### 请求

```http
GET /api/v1/export/data?book_id=1&format=excel&start_date=2023-01-01&end_date=2023-01-31&include_transactions=true&include_accounts=true&include_categories=true
Authorization: Bearer jwt_token_string
```

#### 响应

```
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

###  导入外部记账数据

#### 请求

```http
POST /api/v1/migrate/import
Content-Type: multipart/form-data
Authorization: Bearer jwt_token_string

book_id=1
provider=mint  // 外部记账软件标识
file=...  // 上传的导入文件
mapping=...  // 字段映射JSON字符串
```

#### 响应

```
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

###  设置自动备份配置

#### 请求

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

#### 响应

```
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

##  数据模型

###  同步状态模型

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

###  备份记录模型

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

###  同步配置模型

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

###  数据冲突模型

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

##  同步机制说明

###  增量同步

- 系统采用增量同步机制，只同步发生变化的数据
- 每个实体都有版本号，用于检测冲突
- 支持离线操作，网络恢复后自动同步

###  冲突解决策略

- **服务器优先**: 以服务器数据为准
- **客户端优先**: 以客户端数据为准
- **手动解决**: 提示用户选择保留的数据

###  同步范围

- 收支记录
- 账户信息
- 分类信息
- 标签信息
- 附件文件
- 预算数据

##  备份机制说明

###  备份类型

- **完整备份**: 包含所有数据的完整副本
- **增量备份**: 仅包含上次备份后变更的数据

###  备份存储位置

- 云端备份
- 本地备份（可选）
- 外部存储（可选）

###  备份保留策略

- 可配置保留天数
- 自动清理过期备份
- 重要备份可设为永久保留

##  错误码说明

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