### 接口使用示例

- 登录认证
```shell
curl -XPOST 'https://192.168.165.89:8085/api/auth/login' --cacert admin-template/api-adtmp/ssl/server.crt -H "Content-Type: application/json" -d '{"email":"admin@test.com", "password":"123456"}' 
```

- 用户信息
```shell
curl -XGET 'https://192.168.165.89:8085/api/auth/userinfo' --cacert admin-template/api-adtmp/ssl/server.crt -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQHRlc3QuY29tIiwibmFtZSI6ImFkbWluIiwiZXhwIjoxNzYyODUwMTQyLCJpYXQiOjE3NjI4NTAwODJ9.aF50xYwPhFgNUJPKpdGgyx7A_Z26Ns1zjMf02CNVUVE"
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