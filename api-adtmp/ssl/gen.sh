#!/bin/bash

## 当前目录运行此脚本

## 客户端验证证书
# curl --cacert ssl/server.crt 'https://192.168.165.89:8085'

set -e

expired_time=3650

cat >openssl.cnf<<EOF
[req]
distinguished_name = req_distinguished_name
x509_extensions = v3_req
prompt = no

[req_distinguished_name]
C = CN   # 国家
ST = SiChuan  # 省份(州)
L = Chengdu  # 市
O = Freedom  # 组织(公司)
OU = own  # 组织单位(子部门或子公司)
CN = 192.168.165.89 # 对于 SSL/TLS 证书，通常这个字段会是域名

[v3_req]
keyUsage = keyEncipherment, dataEncipherment
extendedKeyUsage = serverAuth
subjectAltName = @alt_names

[alt_names]
IP.1 = 192.168.165.88
IP.2 = 192.168.165.89
EOF

# 生成证书
openssl req -x509 -nodes -days $expired_time -newkey rsa:2048 \
    -keyout server.key -out server.crt \
    -config openssl.cnf -extensions v3_req