>SHOW DATABASES

显示当前连接的用户拥有访问权限的数据库

>DROP {DATABASE | SCHEMA} [IF EXISTS] db_name

删除数据库

- DROP:关键字
- DATABASE | SCHEMA : 
- IF EXISTS: 可选参数，表示如果要删除的库不存在，则记录警告信息
- db_name:指定要删除的库名

>SHOW WARNINGS

查询警告信息

>CREATE {DATABASE | SCHEMA} [IF NOT EXISTS] db_name
>[create_specification]......
>create_specification:
>[DEFAULT] CHARCTER SET [=] charset_name
>| [DEFAULT] COLLATE [=] collection_name

创建数据库

>SHOW CREATE DATABASE db_name

生成数据库创建时的脚本

>USE db_name

选定数据库

>SHOW TABLES

查看某个库中拥有的表对象

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

创建数据库表格

>{DESC | DESCRIBE} tbl_name

显示表结构

>ALTER [IGNORE] TABLE tbl_name
>[alter_specification [,alter_specification]...]
>[partition_options]

更改表结构

>DROP [TEMPORARY] TABLE [IF EXISTS]
>tbl_name [，tbl_name] ...
>[RESTRICT | CASCADE]

删除表

> RENAME TABLE tbl_name TO new_tbl_name

重命名表对象名称，可以用来移除表

## 权限