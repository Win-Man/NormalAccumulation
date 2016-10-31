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