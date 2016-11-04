## 2016.10.31

### mysqlimport 出错

执行
```c
mysqlimport -uroot -p'admin' -S /var/lib/mysql/mysql.sock jssdb --default-character-set=gbk --fields-terminated-by=',' /tmp/ld_cmd.txt;
```
出现错误：
```
mysqlimport: Error: 1290, The MySQL server is running with the --secure-file-priv option so it cannot execute this statement, when using table: ld_cmd
```

解决方案：
在 /etc/my.cnf 文件中添加一行 secure-file-priv=/tmp 指明导入文件的路径，并重启MySQL服务


## 2016.11.4

### 配置主从环境 slave节点连接不上master节点

```
Last_IO_Errno: 2003
Last_IO_Error: error connecting to master 'repl@192.168.222.131:3306' - retry-time: 60  retries: 3

```

原因：master 节点的防火墙没有关闭，导致无法访问端口

解决方案：

master 节点执行 `service iptables stop`,重启 slave 节点即可

