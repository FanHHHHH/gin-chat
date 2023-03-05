## 部署mysql

docker pull mysql:lastest

### 创建映射目录

mkdir /opt
mkdir /opt/mysql
mkdir /opt/mysql/conf/
mkdir /opt/mysql/logs/
mkdir /opt/mysql/data/

### 创建配置文件

touch /opt/mysql/my.cnf

### 配置文件内容

```
[mysqld]
user=mysql
character-set-server=utf8
default_authentication_plugin=mysql_native_password
secure_file_priv=/var/lib/mysql
expire_logs_days=7
sql_mode=STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION
max_connections=1000

[client]
default-character-set=utf8

[mysql]
default-character-set=utf8
```

### 启动mysql

docker run --restart=always --privileged=true  \
-v /opt/mysql/data/:/var/lib/mysql \
-v /opt/mysql/logs/:/var/log/mysql \
-v /opt/mysql/conf/:/etc/mysql \
-v /opt/mysql/my.cnf:/etc/mysql/my.cnf  \
-v /opt/mysql/conf/:/etc/mysql/conf.d \
-p 3306:3306 --name my-mysql \
-e MYSQL_ROOT_PASSWORD=123456 -d mysql


### 进入容器
docker exec -it my-mysql /bin/bash

### 进入mysql
mysql -u root -p


### 允许远程访问
```
GRANT ALL ON *.* TO 'root'@'%';
flush privileges;

ALTER USER 'root'@'localhost' IDENTIFIED BY 'password' PASSWORD EXPIRE NEVER;
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY '123456';
flush privileges;
```