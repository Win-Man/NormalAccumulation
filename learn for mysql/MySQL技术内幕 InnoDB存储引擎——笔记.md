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

InnoDB 存储引擎中，缓冲页的大小默认为 16KB 。

InnoDB 1.0.x 版本开始支持压缩页功能，即将原本 16KB 的页压缩为 1KB、2KB、4KB 和 8KB。对于非 16KB 的页通过 unzip_LRU 列表进行管理。在 unzip_LRU 列表中对不同压缩页大小的页进行分别管理。

- **1.缓冲池**

在内存中开辟一块空间存放：索引页、数据页、undo 页、插入缓冲、自适应哈希索引、InnoDB 存储的锁信息、数据字典信息。

`innodb_buffer_pool_size` 设置缓冲池的大小

`innodb_buffer_pool_instances` 设置缓冲池实例的个数（1.0.x斑版本开始）

- **2.LRU List、Free List 和 Flush List**

内存管理算法

1. LRU（最近最少使用算法）

LRU List 用来管理已经读取的页。最频繁使用的页在 LRU 列表的前端，而最少使用的页在 LRU 列表的尾端。InnoDB 引擎对 LRU 算法有优化，将新读取到的页插入到列表的 midpoint 位置，而不是尾端。

`innodb_old_blocks_pct` 插入 LRU 列表的位置

`innodb_old_blocks_time` 设置页读取到 mid 位置后需要等待多久才会被加入到 LRU 列表的热端

2. Free List

Free List 存放可用的空闲页

3. Flush List

在 LRU 列表中的页被修改之后，称该页为脏页，即缓冲池中的页和磁盘上的页的数据产生了不一致。这时数据库会通过 CHECKPOINT 机制将脏页刷新回磁盘，Flush 列表中存放的即为脏页列表。

- **3. 重做日志缓冲**

InnoDB 存储引擎首先将重做日志信息放入 redo log buffer ，然后按照一定的频率将其刷新到重做日志文件。 redo log buffer 一般不需要设置很大，通过参数 `innodb_log_buffer_size` 设置，默认为 8MB。

三种情况将 redo log buffer 中的内容刷新到磁盘

1. Master Thread 每一秒将重做日志缓冲刷新到重做日志文件
2. 每个事务提交时将重做日志缓冲刷新到重做日志文件
3. 当重做日志缓冲池剩余空间小于 1/2 时，重做日志缓冲刷新到重做日志文件。

- **4.额外的内存池**

在 InnoDB 存储引擎中，对内存的管理是通过一种称为内存堆的方式进行的。在对一些数据结构本身的内存进行分配是，需要从额外的内存池中进行申请，当该区域的内存不够是，会从缓冲池中进行申请。

## 2.4 Checkpoint 技术

当前事务数据库系统普遍采用了 Write Ahead Log 策略，即当事务提交时，先写重做日志，再修改页。

Checkpoint 技术的目的：

- 缩短数据库的恢复时间
- 缓冲池不够用是，刷新脏页
- 重做日志不可用时，刷新脏页

重做日志出现不可用的情况是因为当前事务数据库系统对重做日志的设计都是循环使用的，并不是让其无线增大的。

对 InnoDB 存储引擎而言，是通过 LSN（Log Sequence Number）来标记版本的。 LSN 实际上对应日志文件的偏移量。

在 InnoDB 存储引擎内部，氛分为两种 Checkpoint，分别是：**Sharp Checkpoint**、**Fuzzy Checkpoint**。

**Sharp Checkpoint**发生在数据库关闭时将所有的脏页都刷新回磁盘，这是默认的工作方式，即参数 `innodb_fast_shutdown=1`。

**Fuzzy Checkponit**是数据库在运行过程中采用的 Checkpint 方式，其包括四种情况：

- Master Thread Checkpoint

每秒或者每十秒的速度从缓冲池的脏页列表中刷新一定比例的页回磁盘。这个过程是异步的，即此时 InnoDB 存储引擎可以进行其他的操作，用户查询线程不会阻塞。

- FLUSH_LRU_LIST Checkpoint

为保证 LRU 列表中差不多存在 100 个空闲页可供使用
参数 `innodb_lru_scan_depth` 控制 LRU 列表中可用页的数量

- Async/Sync Flush Checkpoint

指的是重做日志文件不可用的情况，这时需要强制将一些页刷新回磁盘。

- Dirty Page too much Checkpoint

脏页的数量太多，导致 InnoDB 存储引擎强制进行 Checkpoint。其目的总的来说还是为了保证缓冲池中有足够可用的页。
参数 `innodb_max_dirty_pages_pct` 设置缓冲池中脏页的百分比，达到百分比后就进行 Checkpoint 。

## 2.5 Master Thread 工作方式

### 2.5.1 InnoDB 1.0.x 版本之前的 Master Thread

Master Thread 具有最高的线程优先级别。其内部由多个循环（loop）组成：主循环（loop）、后台循环（background loop）、刷新循环（flush loop）、暂停循环（suspend loop）。

主循环包括两大部分的操作：每秒钟的操作和每 10 秒的操作。

主循环每秒钟的操作包括：
- 日志缓冲刷新到磁盘，即使这个事务还没有提交（总是）
- 合并插入缓冲（可能）
- 至多刷新 100 个 InnoDB 的缓冲池中的脏页到磁盘（可能）
- 如果当前没有用户活动，则切换到 background loop （可能）

主循环每 10 秒的操作包括：
- 刷新 100 个脏页到磁盘（可能的情况下）
- 合并至多 5 个插入缓冲（总是）
- 将日志缓冲刷新到磁盘（总是）
- 删除无用的 Undo 页（总是）
- 刷新 100 个或者 10 个脏页到磁盘（总是）

后台循环（background loop）操作包括：
- 删除无用的 Undo 页（总是）
- 合并 20 个插入缓冲（总是）
- 跳回到主循环（总是）
- 不断刷新 100 个页知道符合条件（可能，跳转到 flush loop 中完成）

### 2.5.2 InnoDB 1.2.x 版本之前的Master Thread

从 InnoDB 1.0.x开始，提供参数 `innodb_io_capacity` ，用来表示磁盘 IO 的吞吐量，默认值200.对于刷新到磁盘也的数量，会按照 innodb_io_capacity 的百分比来进行控制。
- 在合并插入缓冲，合并插入缓冲的数量为 innodb_io_capacity 值为 5%
- 在从缓冲区刷新脏页时，刷新脏页的数量为innodb_io_capacity


InnoDB 1.0.x 版本带来的另一个参数是 `innodb_adaptive_flushing`（自适应地刷新），该值影响每秒刷新脏页的数量。原来的刷新规则是：脏页的缓冲池所占的比例小于 `innodb_max_dirty_pages_pct` 时，不刷新脏页；大于 `innodb_max_dirty_pages_pct` 时，刷新 100 个脏页。随着 `innodb_adaptive_flushing` 参数的引入，InnoDB 存储引擎会通过一个名为 `buf_flush_get_desired_flush_rate` 的函数来判断需要刷新脏页的最合适的数量。粗略地翻阅源代码后发现 `buf_flush_get_desired_flush_rate` 通过判断产生重做日志的速度来决定最合适的刷新脏页数量。因此，当脏页的比例小于　`innodb_max_dirty_pages_pct` 时，也会刷新一定量的脏页。

InnoDB 1.0.x版本引入参数 `innodb_purge_batch_size` 控制每次 flush purge 回收的 Undo 页的数量。

### 2.5.3 InnoDB 1.2.x 版本的 Master Thread

对于刷新脏页的操作，从 Master Thread 线程分离到一个单独的 Page Cleaner Thread，从而减轻了 Master Thread 的工作，同时进一步提高了系统的并发性。

## 2.6 InnoDB 关键特性

InnoDB 存储引擎的关键特性包括：
- 插入缓冲（Insert Buffer）
- 两次写（Double Write）
- 自适应哈希索引(Adaptive Hash Index)
- 异步 IO（Async IO）
- 刷新邻接页（Flush Neighbor Page）

### 2.6.1 插入缓冲

InnoDB 存储引擎开创性的设计了 Insert Buffer，对于非聚集索引的插入或更新操作，不是每一次直接插入到索引页中，而是先判断插入的非聚集索引页是否在缓冲池中，若在，则直接插入；若不在，则先放入到一个 Insert Buffer 对象中，好似欺骗数据库这个非聚集的索引已经插到叶子节点，而实际并没有，只是存放在另一个位置。然后再以一定频率和情况进行 Insert Buffer 和辅助索引叶子节点的 merge 操作，这时通常能将多个插入合并到一个操作中，这就大大提高了对于非聚集索引插入的性能。

Insert Buffer 使用需要满足条件:
- 索引是辅助索引
- 索引不是唯一的