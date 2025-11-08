###

- 登录认证
```shell
curl -XPOST 'http://127.0.0.1:8085/api/auth/login' -H "Content-Type: application/json" -d '{"email":"admin@test.com", "password":"123456"}'
```

- 新增用户
```shell
curl -XPOST 'http://127.0.0.1:8085/api/user' -H "Content-Type: application/json" -d '{
  "name": "admin",
  "email": "admin@test.com",
  "password": "123456",
  "roles": [
    {"name": "admin"},
    {"name": "editor"}
  ]
}
'
```

