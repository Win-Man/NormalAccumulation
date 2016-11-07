[mysqld]

`log_bin = mysql_bin`  //开启二进制日志，指定二进制日志文件的名字，不指定的话采用主机名作为默认名字

`binlog-ignore-db = master_slave_new` //过滤该db，不写到binlog中

`binlog-do-db = master_slave` //允许该db相关操作写到binlog中

`basedir = /usr/local/mysql` //mysql所在目录
`datadir =/usr/local/mysql/data` // mysql的data目录
`port = 3306` //mysql启动端口
`server_id = 2433306` //mysql实例的id，主从结构时不能重复
`socket = .....` // mysql sock文件所在地址


sql_mode=NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES


[参考博客](http://www.cnblogs.com/toby/articles/2198697.html)
