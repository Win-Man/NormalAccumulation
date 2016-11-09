[mysqld]

`log_bin = mysql_bin`  //开启二进制日志，指定二进制日志文件的名字，不指定的话采用主机名作为默认名字

`binlog-ignore-db = master_slave_new` //过滤该db，不写到binlog中

`binlog-do-db = master_slave` //允许该db相关操作写到binlog中

`basedir = /usr/local/mysql` //mysql所在目录

`datadir =/usr/local/mysql/data` // mysql的data目录

`port = 3306` //mysql启动端口

`server_id = 2433306` //mysql实例的id，主从结构时不能重复

`socket = .....` // mysql sock文件所在地址

`read_only` //设置数据库为只读状态

`sync_binlog = 1` //每经过 n 次日志写就把日志文件写入硬盘一次

`binlog-format=ROW` // 设置 binlog 记录形式，是基于行的还是基于语句的还是混合的




sql_mode=NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES


[参考博客](http://www.cnblogs.com/toby/articles/2198697.html)
