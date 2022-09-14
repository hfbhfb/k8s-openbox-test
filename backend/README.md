- mysql 数据库
  - `CREATE USER 'userddns'@'%' IDENTIFIED BY 'sdfidsjf';`
  - `create database dbddns;`
  * `GRANT ALL ON dbddns.* TO 'userddns'@'%';`




- mysql 数据库
  - `CREATE USER 'userqsy'@'%' IDENTIFIED BY 'sdfidsjf';`
  - `create database dbqsy;`
  * `GRANT ALL ON dbqsy.* TO 'userqsy'@'%';`


## 恢复sql数据库


docker run --name mysqlqushuiyin --restart always -v /home/usera/nginxdockergz/forqsy/dbdata:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=mm9988 -p 15002:3306 -d mysql:5.7






gowatch -args=--'env=uat'

- docker run --name mysqlqushuiyin -v /home/ec2-user/mysqldata1:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=mm9988 -p 8367:3306 -d mysql:5.7.36

- 把 config-default.json -> config-dev.json
- 配置 mysql
- 配置 redis

- cd /Users/hfb/projects/go/mod-pro/gormt-master
- cp config-qsy.yml config.yml
- ./main -g=true

```json
# hfbdeMacBook-Pro:gormt-master hfb$ cat config.yml
base:
  is_dev: false
out_dir: /Users/hfb/projects/go/mod-pro/qu-sui-yin/backend/codeasset/automodel
#out_dir : ./modelgormt  # 输出目录
url_tag: json # web url tag(json,db(https://github.com/google/go-querystring))
language: # 语言(English,中 文)
db_tag: gorm # 数据库标签名(gorm,db)
simple: true # 简单输出(默认只输出gorm主键和字段标签)
is_db_tag: true # 是否输出 数据库标签(gorm,db)
is_out_sql: false # 是否输出 sql 原信息
is_out_func: true # 是否输出 快捷函数
is_web_tag: true # 是否打web标记(json标记前提条件)
is_web_tag_pk_hidden: true # web标记是否隐藏主键
is_foreign_key: true # 是否导出外键关联
is_gui: false # 是否ui模式显示
is_table_name: true # 是否直接生成表名
is_column_name: true # 是否直接生成列名
is_null_to_point: false # 数据库默认 'DEFAULT NULL' 时设置结构为指针类型
table_prefix: "" # 表前缀, 如果有则使用, 没有留空
table_names: "" # 指定表生成，多个表用,隔开

db_info:
  host: 127.0.0.1 # type=1的时候，host为yml文件全路径
  port: 3306
  username: go_rbac_adm
  password: 123456
  database: godbrbac
  type: 0 # 数据库类型:0:mysql , 1:sqlite , 2:mssql
self_type_define: # 自定义数据类型映射
  datetime: time.Time
  time: time.Time
out_file_name: "" # 自定义生成文件名
web_tag_type: 0 # json tag类型 0: 小驼峰 1: 下划线

# sqlite
# db_info:
#     host : /Users/xxj/Downloads/caoguo # type=1的时候，host为yml文件全路径
#     port :
#     username :
#     password :
#     database :
#     type: 1 # 数据库类型:0:mysql , 1:sqlite , 2:mssql

```

- 几个测试用户
  - 372292 123
