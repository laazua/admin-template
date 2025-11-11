###

- [接口使用示例](./api.md)

- 生成证书
```shell
# 生成私钥
openssl genrsa -out server.key 2048

# 生成证书签名请求
openssl req -new -key server.key -out server.csr

# 生成自签名证书
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt

# 设置严格的文件权限
chmod 600 server.key      # 只有所有者可读写
chmod 644 server.crt      # 证书可公开读取

```