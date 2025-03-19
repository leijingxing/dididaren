# 滴滴打人 API 文档

## 基础信息

- 基础URL: `http://localhost:8080/api/v1`
- 所有需要认证的接口都需要在请求头中携带 `Authorization: Bearer <token>`
- 响应格式统一为 JSON

## 认证相关

### 用户注册

- 请求方法：`POST`
- 路径：`/users/register`
- 请求体：
```json
{
    "phone": "13800138000",
    "password": "password123",
    "name": "张三"
}
```
- 响应：
```json
{
    "message": "注册成功",
    "data": {
        "id": 1,
        "phone": "13800138000",
        "name": "张三"
    }
}
```

### 用户登录

- 请求方法：`POST`
- 路径：`/users/login`
- 请求体：
```json
{
    "phone": "13800138000",
    "password": "password123"
}
```
- 响应：
```json
{
    "message": "登录成功",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIs...",
        "user": {
            "id": 1,
            "phone": "13800138000",
            "name": "张三"
        }
    }
}
```

### 验证码验证

- 请求方法：`POST`
- 路径：`/users/verify-code`
- 请求体：
```json
{
    "phone": "13800138000",
    "code": "123456"
}
```
- 响应：
```json
{
    "message": "验证成功"
}
```

## 用户相关

### 获取用户信息

- 请求方法：`GET`
- 路径：`/users/profile`
- 需要认证：是
- 响应：
```json
{
    "data": {
        "id": 1,
        "phone": "13800138000",
        "name": "张三",
        "avatar": "http://example.com/avatar.jpg"
    }
}
```

### 更新用户信息

- 请求方法：`PUT`
- 路径：`/users/profile`
- 需要认证：是
- 请求体：
```json
{
    "name": "张三",
    "avatar": "http://example.com/avatar.jpg"
}
```
- 响应：
```json
{
    "message": "更新成功",
    "data": {
        "id": 1,
        "phone": "13800138000",
        "name": "张三",
        "avatar": "http://example.com/avatar.jpg"
    }
}
```

### 添加紧急联系人

- 请求方法：`POST`
- 路径：`/users/emergency-contacts`
- 需要认证：是
- 请求体：
```json
{
    "name": "李四",
    "phone": "13900139000",
    "relation": "朋友"
}
```
- 响应：
```json
{
    "message": "添加成功",
    "data": {
        "id": 1,
        "name": "李四",
        "phone": "13900139000",
        "relation": "朋友"
    }
}
```

### 获取紧急联系人列表

- 请求方法：`GET`
- 路径：`/users/emergency-contacts`
- 需要认证：是
- 响应：
```json
{
    "data": [
        {
            "id": 1,
            "name": "李四",
            "phone": "13900139000",
            "relation": "朋友"
        }
    ]
}
```

### 删除紧急联系人

- 请求方法：`DELETE`
- 路径：`/users/emergency-contacts/:id`
- 需要认证：是
- 响应：
```json
{
    "message": "删除成功"
}
```

## 紧急事件相关

### 创建紧急事件

- 请求方法：`POST`
- 路径：`/emergencies`
- 需要认证：是
- 请求体：
```json
{
    "type": "抢劫",
    "description": "在XX路发生抢劫事件",
    "latitude": 39.9042,
    "longitude": 116.4074,
    "address": "北京市东城区XX路"
}
```
- 响应：
```json
{
    "message": "创建成功",
    "data": {
        "id": 1,
        "type": "抢劫",
        "description": "在XX路发生抢劫事件",
        "latitude": 39.9042,
        "longitude": 116.4074,
        "address": "北京市东城区XX路",
        "status": "pending"
    }
}
```

### 获取事件详情

- 请求方法：`GET`
- 路径：`/emergencies/:id`
- 需要认证：是
- 响应：
```json
{
    "data": {
        "id": 1,
        "type": "抢劫",
        "description": "在XX路发生抢劫事件",
        "latitude": 39.9042,
        "longitude": 116.4074,
        "address": "北京市东城区XX路",
        "status": "processing",
        "created_at": "2024-01-01T12:00:00Z",
        "updated_at": "2024-01-01T12:05:00Z"
    }
}
```

### 更新事件状态

- 请求方法：`PUT`
- 路径：`/emergencies/:id/status`
- 需要认证：是
- 请求体：
```json
{
    "status": "completed"
}
```
- 响应：
```json
{
    "message": "更新成功",
    "data": {
        "id": 1,
        "status": "completed"
    }
}
```

### 获取事件历史

- 请求方法：`GET`
- 路径：`/emergencies/history`
- 需要认证：是
- 响应：
```json
{
    "data": [
        {
            "id": 1,
            "type": "抢劫",
            "status": "completed",
            "created_at": "2024-01-01T12:00:00Z"
        }
    ]
}
```

### 创建处理记录

- 请求方法：`POST`
- 路径：`/emergencies/:id/handling-records`
- 需要认证：是
- 请求体：
```json
{
    "content": "已到达现场，正在处理中",
    "status": "processing"
}
```
- 响应：
```json
{
    "message": "创建成功",
    "data": {
        "id": 1,
        "content": "已到达现场，正在处理中",
        "status": "processing",
        "created_at": "2024-01-01T12:05:00Z"
    }
}
```

### 获取处理记录

- 请求方法：`GET`
- 路径：`/emergencies/:id/handling-records`
- 需要认证：是
- 响应：
```json
{
    "data": [
        {
            "id": 1,
            "content": "已到达现场，正在处理中",
            "status": "processing",
            "created_at": "2024-01-01T12:05:00Z"
        }
    ]
}
```

## 安保人员相关

### 申请成为安保人员

- 请求方法：`POST`
- 路径：`/security/apply`
- 需要认证：是
- 请求体：
```json
{
    "id_card": "110101199001011234",
    "real_name": "张三",
    "experience": "5年安保经验"
}
```
- 响应：
```json
{
    "message": "申请成功",
    "data": {
        "id": 1,
        "user_id": 1,
        "id_card": "110101199001011234",
        "real_name": "张三",
        "experience": "5年安保经验",
        "status": "pending"
    }
}
```

### 更新位置

- 请求方法：`PUT`
- 路径：`/security/location`
- 需要认证：是
- 请求体：
```json
{
    "latitude": 39.9042,
    "longitude": 116.4074
}
```
- 响应：
```json
{
    "message": "更新成功"
}
```

### 接单

- 请求方法：`POST`
- 路径：`/security/events/:id/accept`
- 需要认证：是
- 响应：
```json
{
    "message": "接单成功",
    "data": {
        "event_id": 1,
        "staff_id": 1,
        "status": "processing"
    }
}
```

### 完成订单

- 请求方法：`PUT`
- 路径：`/security/events/:id/complete`
- 需要认证：是
- 响应：
```json
{
    "message": "完成成功"
}
```

### 获取安保人员信息

- 请求方法：`GET`
- 路径：`/security/profile`
- 需要认证：是
- 响应：
```json
{
    "data": {
        "id": 1,
        "user_id": 1,
        "id_card": "110101199001011234",
        "real_name": "张三",
        "experience": "5年安保经验",
        "status": "active",
        "order_count": 10,
        "rating": 4.8
    }
}
```

## 危险区域相关

### 创建危险区域

- 请求方法：`POST`
- 路径：`/danger-zones`
- 需要认证：是
- 请求体：
```json
{
    "name": "XX路危险区域",
    "description": "该区域经常发生抢劫事件",
    "latitude": 39.9042,
    "longitude": 116.4074,
    "radius": 1000,
    "heat_level": 3
}
```
- 响应：
```json
{
    "message": "创建成功",
    "data": {
        "id": 1,
        "name": "XX路危险区域",
        "description": "该区域经常发生抢劫事件",
        "latitude": 39.9042,
        "longitude": 116.4074,
        "radius": 1000,
        "heat_level": 3,
        "status": 1
    }
}
```

### 获取危险区域详情

- 请求方法：`GET`
- 路径：`/danger-zones/:id`
- 需要认证：是
- 响应：
```json
{
    "data": {
        "id": 1,
        "name": "XX路危险区域",
        "description": "该区域经常发生抢劫事件",
        "latitude": 39.9042,
        "longitude": 116.4074,
        "radius": 1000,
        "heat_level": 3,
        "status": 1
    }
}
```

### 更新危险区域

- 请求方法：`PUT`
- 路径：`/danger-zones/:id`
- 需要认证：是
- 请求体：
```json
{
    "name": "XX路危险区域",
    "description": "该区域经常发生抢劫事件",
    "latitude": 39.9042,
    "longitude": 116.4074,
    "radius": 1000,
    "heat_level": 3,
    "status": 1
}
```
- 响应：
```json
{
    "message": "更新成功",
    "data": {
        "id": 1,
        "name": "XX路危险区域",
        "description": "该区域经常发生抢劫事件",
        "latitude": 39.9042,
        "longitude": 116.4074,
        "radius": 1000,
        "heat_level": 3,
        "status": 1
    }
}
```

### 删除危险区域

- 请求方法：`DELETE`
- 路径：`/danger-zones/:id`
- 需要认证：是
- 响应：
```json
{
    "message": "删除成功"
}
```

### 获取附近的危险区域

- 请求方法：`GET`
- 路径：`/danger-zones/nearby`
- 需要认证：是
- 查询参数：
  - `latitude`: 纬度
  - `longitude`: 经度
  - `radius`: 搜索半径（米）
- 响应：
```json
{
    "data": [
        {
            "id": 1,
            "name": "XX路危险区域",
            "description": "该区域经常发生抢劫事件",
            "latitude": 39.9042,
            "longitude": 116.4074,
            "radius": 1000,
            "heat_level": 3,
            "status": 1
        }
    ]
}
```

### 获取所有活跃的危险区域

- 请求方法：`GET`
- 路径：`/danger-zones/active`
- 需要认证：是
- 响应：
```json
{
    "data": [
        {
            "id": 1,
            "name": "XX路危险区域",
            "description": "该区域经常发生抢劫事件",
            "latitude": 39.9042,
            "longitude": 116.4074,
            "radius": 1000,
            "heat_level": 3,
            "status": 1
        }
    ]
}
```

### 更新危险区域热度等级

- 请求方法：`PUT`
- 路径：`/danger-zones/:id/heat-level`
- 需要认证：是
- 查询参数：
  - `heat_level`: 热度等级（0-5）
- 响应：
```json
{
    "message": "更新成功"
}
```

## 评价相关

### 创建评价

- 请求方法：`POST`
- 路径：`/ratings`
- 需要认证：是
- 请求体：
```json
{
    "event_id": 1,
    "staff_id": 1,
    "score": 5,
    "comment": "服务很好，处理及时"
}
```
- 响应：
```json
{
    "message": "创建成功",
    "data": {
        "id": 1,
        "event_id": 1,
        "staff_id": 1,
        "score": 5,
        "comment": "服务很好，处理及时",
        "created_at": "2024-01-01T12:00:00Z"
    }
}
```

### 获取评价详情

- 请求方法：`GET`
- 路径：`/ratings/:id`
- 需要认证：是
- 响应：
```json
{
    "data": {
        "id": 1,
        "event_id": 1,
        "staff_id": 1,
        "score": 5,
        "comment": "服务很好，处理及时",
        "created_at": "2024-01-01T12:00:00Z"
    }
}
```

### 获取事件相关评价

- 请求方法：`GET`
- 路径：`/ratings/event/:event_id`
- 需要认证：是
- 响应：
```json
{
    "data": {
        "id": 1,
        "event_id": 1,
        "staff_id": 1,
        "score": 5,
        "comment": "服务很好，处理及时",
        "created_at": "2024-01-01T12:00:00Z"
    }
}
```

### 获取安保人员的所有评价

- 请求方法：`GET`
- 路径：`/ratings/staff/:staff_id`
- 需要认证：是
- 响应：
```json
{
    "data": [
        {
            "id": 1,
            "event_id": 1,
            "staff_id": 1,
            "score": 5,
            "comment": "服务很好，处理及时",
            "created_at": "2024-01-01T12:00:00Z"
        }
    ]
}
```

### 获取安保人员的平均评分

- 请求方法：`GET`
- 路径：`/ratings/staff/:staff_id/average`
- 需要认证：是
- 响应：
```json
{
    "data": {
        "average_rating": 4.8
    }
}
```

### 更新评价

- 请求方法：`PUT`
- 路径：`/ratings/:id`
- 需要认证：是
- 请求体：
```json
{
    "score": 4,
    "comment": "服务一般"
}
```
- 响应：
```json
{
    "message": "更新成功",
    "data": {
        "id": 1,
        "event_id": 1,
        "staff_id": 1,
        "score": 4,
        "comment": "服务一般",
        "created_at": "2024-01-01T12:00:00Z"
    }
}
```

### 删除评价

- 请求方法：`DELETE`
- 路径：`/ratings/:id`
- 需要认证：是
- 响应：
```json
{
    "message": "删除成功"
}
```

## 系统配置相关

### 获取所有配置

- 请求方法：`GET`
- 路径：`/configs`
- 响应：
```json
{
    "data": [
        {
            "config_key": "max_emergency_distance",
            "config_value": "5000",
            "description": "最大紧急事件响应距离（米）"
        }
    ]
}
```

### 获取配置值

- 请求方法：`GET`
- 路径：`/configs/:key`
- 响应：
```json
{
    "data": {
        "value": "5000"
    }
}
```

### 创建配置（管理员）

- 请求方法：`POST`
- 路径：`/admin/configs`
- 需要认证：是
- 需要管理员权限：是
- 请求体：
```json
{
    "config_key": "max_emergency_distance",
    "config_value": "5000",
    "description": "最大紧急事件响应距离（米）"
}
```
- 响应：
```json
{
    "message": "创建成功",
    "data": {
        "config_key": "max_emergency_distance",
        "config_value": "5000",
        "description": "最大紧急事件响应距离（米）"
    }
}
```

### 更新配置（管理员）

- 请求方法：`PUT`
- 路径：`/admin/configs`
- 需要认证：是
- 需要管理员权限：是
- 请求体：
```json
{
    "config_key": "max_emergency_distance",
    "config_value": "6000",
    "description": "最大紧急事件响应距离（米）"
}
```
- 响应：
```json
{
    "message": "更新成功",
    "data": {
        "config_key": "max_emergency_distance",
        "config_value": "6000",
        "description": "最大紧急事件响应距离（米）"
    }
}
```

### 删除配置（管理员）

- 请求方法：`DELETE`
- 路径：`/admin/configs/:key`
- 需要认证：是
- 需要管理员权限：是
- 响应：
```json
{
    "message": "删除成功"
}
```

### 更新配置值（管理员）

- 请求方法：`PUT`
- 路径：`/admin/configs/:key/value`
- 需要认证：是
- 需要管理员权限：是
- 请求体：
```json
{
    "value": "6000"
}
```
- 响应：
```json
{
    "message": "更新成功"
}
``` 