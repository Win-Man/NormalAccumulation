### 1.显示当前连接的用户拥有访问权限的数据库

>SHOW DATABASES

### 2.删除数据库

>DROP {DATABASE | SCHEMA} [IF EXISTS] db_name

- DROP:关键字
- DATABASE | SCHEMA : 
- IF EXISTS: 可选参数，表示如果要删除的库不存在，则记录警告信息
- db_name:指定要删除的库名

### 3.查询警告信息

>SHOW WARNINGS

### 4.创建数据库

>CREATE {DATABASE | SCHEMA} [IF NOT EXISTS] db_name
>[create_specification]......
>create_specification:
>[DEFAULT] CHARCTER SET [=] charset_name
>| [DEFAULT] COLLATE [=] collection_name

### 5.生成数据库创建时的脚本

>SHOW CREATE DATABASE db_name

### 6.选定数据库

>USE db_name

### 7.查看某个库中拥有的表对象

>SHOW TABLES

### 8.创建数据库表格

>CREATE [TEMPORARY] TABLE [IF NOT EXISTS] tbl_name
>(create_definition,...)
>[table_options]
>[partition_options]

Or

>CREATE [TEMPORARY] TABLE [IF NOT EXISTS] tbl_name
>[(create_definition,...)]
>[table_options]
>[partition_options]
>select_statement

Or

>CREATE [TEMPORARY] TABLE [IF NOT EXISTS] tbl_name
>{LIKE old_tbl_name | (LIKE old_tbl_name)}

### 9.显示表结构

>{DESC | DESCRIBE} tbl_name

### 10.更改表结构

>ALTER [IGNORE] TABLE tbl_name
>[alter_specification [,alter_specification]...]
>[partition_options]

### 11.删除表

>DROP [TEMPORARY] TABLE [IF EXISTS]
>tbl_name [，tbl_name] ...
>[RESTRICT | CASCADE]

### 12.重命名表对象名称，可以用来移除表

> RENAME TABLE tbl_name TO new_tbl_name

### 13.创建用户

> CREATE USER user_specification [, user_specification] ...

### 14.修改密码

> SET PASSWORD [FOR user] = password_option

### 15.授予权限

> GRANT
priv_type [(column_list)]
  [, priv_type [(column_list)]] ...
ON [object_type] priv_level
TO user_specification [, user_specification] ...
[REQUIRE {NONE | tls_option [[AND] tls_option] ...}]
[WITH {GRANT OPTION | resource_option} ...]

### 16.查看权限

> SHOW GRANTS [FOR user]

### 17.收回权限

> REVOKE
priv_type [(column_list)]
  [, priv_type [(column_list)]] ...
ON [object_type] priv_level
FROM user [, user] ...

>REVOKE ALL PRIVILEGES,GRANT OPTION FROM user
>收回所有权限

### 18.删除用户

> DROP USER user[,user]...

### 19.指定客户端当前会话使用的字符集

>SET NAMES {'charset_name'
[COLLATE 'collation_name'] | DEFAULT}

### 20.显示数据库当前所支持的字符集

> show character set;

### 21.显示当前 MySQL 数据所支持的存储引擎

> SHOW ENGINES;