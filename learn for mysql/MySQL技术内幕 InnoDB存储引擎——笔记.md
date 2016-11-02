# 第 1 章 MySQL 体系结构和存储引擎

## 1.1 定义数据库和实例

Linux 下通过命令 `ps -ef | grep mysqld` 查看 MySQL 实例进程是否启动。

通过 `mysql --help | grep my.cnf` 查看 MySQL 数据库实例启动时，会在哪些位置查找配置文件。（按顺序读取，以最后一个文件的配置为准）

在 `mysql> SHOW VARIABLES LIKE 'datadir'\G` 查看数据库所在路径。

## 1.2 MySQL 体系结构

MySQL 数据库体系结构如图所示：
![](http://img.blog.csdn.net/20130517000945865)

MySQL 数据库区别于其他数据库的最重要的一个特点就是其插件式的表存储引擎。**注意点：存储引擎是基于表的，而不是数据库**。

## 1.3 MySQL 存储引擎

### 1.3.1 InnoDB 存储引擎

MySQL5.5.8版本开始，成为默认引擎

支持事务、行锁设计、支持外键

InnoDB 存储引擎采用聚集的方式，因此每张表的存储都是按主键的顺序进行存放。如果没有显示地在表定义时指定主键，InnoDB 存储引擎会为每一行生成一个 6 字节的 ROWID ，并以此作为主键。

### 1.3.2 MyISAM 存储引擎

不支持事务、表锁设计、支持全文搜索

### 1.3.3 NDB 存储引擎

数据全部存放在内存中

### 1.3.4 Memory 存储引擎

将表中的数据存放在内存中，如果数据重启或发生崩溃，表中的数据都将消失

### 1.3.5 Archive 存储引擎

只支持 INSERT 和 SELECT 操作，Archive 存储引擎非常适合存储归档数据

### 1.3.6 Federated 存储引擎

### 1.3.7 Maria 存储引擎

支持缓存数据和索引文件、行锁设计、提供 MVCC 功能，支持事务和非事务安全的选项、更好的 BLOB 字符类型的处理性能

### 1.3.8 其他存储引擎

## 1.4 存储引擎之间的比较

通过 `SHOW ENGINES` 查看当前使用的 MySQL数据所支持的存储引擎

各存储引擎之间比较：

![](http://img4.07net01.com/upload/images/2016/05/16/176201161020079.jpg)

## 1.5 连接 MySQL

### 1.5.1 TCP/IP

### 1.5.2 命名管道和共享内存

### 1.5.3 UNIX 域套接字

通过命令 `SHOW VARIABLES LIKE 'socket'` 进行 UNIX 域套接字文件的查找

# 第 2 章 InnoDB 存储引擎

## 2.1 InnoDB 存储引擎概述

## 2.2 InnoDB 存储引擎的版本

| 版本          | 功能                  |
| -------------- | ----------------------|
| 老版本 InnoDB  | 支持ACID、行锁设计、MVCC |
| InnoDB 1.0.x | 继承了上述版本所有功能，增加了 compress 和dynamic 页格式 |
| InnoDB 1.1.x | 继承了上述版本所有功能，增加了 Linux AIO、多回滚段 |
| InnoDB 1.2.x | 继承了上述版本所有功能，增加了全文索引支持、在线索引添加 |

## 2.3 InnoDB 体系架构

![](http://jockchou.com/blog/img/innodb-1.png)

后台线程的主要作用是负责刷新内存池的数据，保证缓冲池中的内存缓存的是最近的数据。此外将已修改的数据文件刷新到磁盘文件，同时保证在数据库发生异常的情况下 InnoDB 会恢复到正常运行状态。

### 2.3.1 后台线程

- **1.Master Thread**

将缓冲池中的数据**异步**刷新到磁盘，保证数据的一致性，包括脏页的刷新、合并插入缓冲、UNDO页的回收。

- **2.IO Thread**

处理 IO 请求
包括 write、read、insert buffer和log IO thread

- **3.Purge Thread**

回收已经使用并分配的 undo 页。1.2版本之前只支持一个 Purge Thread 线程

- **4.Page Cleaner Thread**

在 1.2.x 版本中引入，作用是讲之前版本中脏页的刷新操作都放入到单独的线程中来完成。其目的是减轻原 Master Thread 的工作及对于用户查询线程的阻塞。

### 2.3.2 内存

- **1.缓冲池**

在内存中开辟一块空间存放：索引页、数据页、undo 页、插入缓冲、自适应哈希索引、InnoDB 存储的锁信息、数据字典信息。

`innodb_buffer_pool_size` 设置缓冲池的大小

`innodb_buffer_pool_instances` 设置缓冲池实例的个数（1.0.x斑版本开始）

- **LRU List、Free List 和 Flush List**

内存管理算法

1. LRU（最近最少使用算法）

LRU List 用来管理已经读取的页。最频繁使用的页在 LRU 列表的前端，而最少使用的页在 LRU 列表的尾端。InnoDB 引擎对 LRU 算法有优化，将新读取到的页插入到列表的 midpoint 位置，而不是尾端。

`innodb_old_blocks_pct` 插入 LRU 列表的位置

`innodb_old_blocks_time` 设置页读取到 mid 位置后需要等待多久才会被加入到 LRU 列表的热端

2. Free