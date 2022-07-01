# count(*) 为什么慢

count(*) 为什么慢，扫描行遇到可见行就记数

MySQL可以优化会选取普通索引计数。

`show table status` 获取到的行数是估计数有 `误差`

MyISAM表count(*)很快，但是不支持事务；

show table status 虽快但是不准确

InnoDB count(*) 会遍历全表，虽然准确但是慢

## **count(*)** **、** **count(**主键**id)** **、** **count(** **字段****) 和  **count(1)** **区别**

对于count(主键id)来说，InnoDB会遍历整表，把id值都取出来返回到server层，server层拿到id进行判断再累加。

对于count(1)来说，InnoDB遍历表但是不取值，server层对于返回的每一行放一个数字1进去，判断再累加

对于count(字段)来说，定义为not nill，需要一行行的读出这个地段，判断不能为null，按行累加。

如果允许为null还需要取值判断是否为null再累加。

对于count(*)是例外，它不会把全部字段取出来，二手专门做了优化不取值 `count(*)`肯定不是null，按行累加。


## 分页怎么做

大数据下 `SELECT * FROM T LIMIT N,M` 是很慢的，MySQL处理limit N,M的做法就是按顺序一个一个地读出来数据集，limit只做过滤，然后丢掉前N个，剩下M个记录作为返回结果，因此这一步需要扫描N+M行;N越大扫的行数越多;值越大越需要遍历大量数据页，耗费的时间也越久。

* 使用断点记录：

`select * from xes_order id>xxx limit new_N,M;`每个一定距离记录一个值。

* 带where条件的使用覆盖索引：

`SELECT * FROM t WHERE K=X LIMIT N,M`，是根据k索引存储的主键去查找对应的行。K索引和主键索引数据不在同一个物理块上，N值越大越需要遍历大量索引页和数据叶，耗费的时间就越久。

使用覆盖索引进行解决，减小无效回表 `select * from tplogic a inner join (select post_id from tplogic where thread_id = 1 limit 10000000, 100) b where a.post_id = b.post_id`。

利用覆盖索引:平均遍历5w的偏移需要20ms，10w的偏移需要30ms，在小于100w的情况 下每增加10w的偏移，耗时增加30ms，到达90w时耗时270ms，响应耗时控制在270ms的情况下，可支持在180w内的精准查看;
