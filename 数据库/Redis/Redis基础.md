## redis

```bash
[root@localhost bin]# redis-cli -p 6379 #连接服务器设定服务器端口
127.0.0.1:6379> ping #测试连接
PONG
```

## Redis Key命令

#### 参看所有

keys *

#### 否存在某个key

EXISTS key

#### 设置过期时间

EXPIRE key expirationtime

#### 查看过期时间

ttl key

#### 参看数据类型

type key

> 具体用例

```bash
127.0.0.1:6379> set name weixi #String类型的设置方式。。
OK
127.0.0.1:6379> set value blog
OK
127.0.0.1:6379> keys *     参看所有的key
1) "value"
2) "name"
127.0.0.1:6379> EXISTS name #是否有名字为"name"的key
(integer) 1 #true 返回1
127.0.0.1:6379> EXISTS firstName
(integer) 0 #false 返回0
127.0.0.1:6379> EXPIRE name 5 #设置key的国企时间 基本单位为秒S
(integer) 1
127.0.0.1:6379> ttl name #参看name剩余时间
(integer) 2 #2S
127.0.0.1:6379> ttl name
(integer) 1
127.0.0.1:6379> ttl name
(integer) -2 #已过期
127.0.0.1:6379> get name #再次获取发下已经获取不到了
(nil)
127.0.0.1:6379> type value #查看value的类型
string
```

## String(字符串类型)

#### 设置值

set key value

```bash
  127.0.0.1:6379> set name1 zhangsan #设置值
  OK
```

#### 获得值

get key

```bash
  127.0.0.1:6379> get name1 #获取值
  "zhangsan"
```

#### 追加字符串

append key value

注：如果key不存在，则会先创建key

```bash
  127.0.0.1:6379> get name1 #获取值
  "zhangsan"
  127.0.0.1:6379> append name1 fawaikuangtu #追加字符串
  (integer) 20
  127.0.0.1:6379> get name1
  "zhangsanfawaikuangtu"
  127.0.0.1:6379> APPEND name2 lisi #追加字符串，如果不存在就添加
  (integer) 4
  127.0.0.1:6379> get name2
  "lisi"
```

#### 获取值的长度

strlen key

```bash
  127.0.0.1:6379> STRLEN name1 #获取值的长度
  (integer) 20
```

#### 自增1

INCR key

```bash
  127.0.0.1:6379> set i 1
  OK
  127.0.0.1:6379> type i
  string
  127.0.0.1:6379> INCR i #自增1
  (integer) 2 #返回结果
```

#### 自减1

DECR key

```bash
  127.0.0.1:6379> DECR i #自减1
  (integer) 1 
  127.0.0.1:6379> DECR i
  (integer) 0
```

#### 自增x

INCRBY i x

```bash
  127.0.0.1:6379> INCRBY i 7 #自增7  (integer) 7
```

#### 自减x

DECRBY i x

```bash
  127.0.0.1:6379> DECRBY i 10 #自减10  (integer) -3
```

#### 字符串范围Range

GETRANGE key start end

注：如果end为-1则获取到结束

```bash
  127.0.0.1:6379> GETRANGE name1 7 11 #获取索引7-11的值  "nfawa"  127.0.0.1:6379> GETRANGE name1 7 -1 #获取索引7-结束的值  "nfawaikuangtu"
```

#### 替换字符串

SETRANGE key offset value

```bash
  127.0.0.1:6379> get name1  "zhangsanfawaikuangtu"  127.0.0.1:6379> SETRANGE name1 4 si #替换4位置开始的字符串  (integer) 20  127.0.0.1:6379> get name1  "zhansianfawaikuangtu"
```

#### 设置值及过期时间

SETEX key seconds value

```bash
  127.0.0.1:6379> setex name3 15 wangwu #设置name3的值为wangwu 15秒后过期  OK  127.0.0.1:6379> ttl name3  (integer) 12  127.0.0.1:6379> ttl name3  (integer) 8  127.0.0.1:6379> ttl name3  (integer) 7  127.0.0.1:6379> get name3  "wangwu"  127.0.0.1:6379> ttl name3  (integer) 0  127.0.0.1:6379> ttl name3  (integer) -2
```

#### 如果不存在key则设置值

SETNX key value

```bash
  127.0.0.1:6379> setnx name4 zhaoliu #如果不存在name4，则设置name4  (integer) 1  127.0.0.1:6379> keys *  1) "name4"  2) "i"  3) "name2"  4) "name1"  127.0.0.1:6379> setnx name4 zhangwuji  (integer) 0 #返回0，设置失败，因为已经设置了name4
```

#### 设置多个值

mset key value [key value …]

```bash
  127.0.0.1:6379> FLUSHdb #清空数据库  OK  127.0.0.1:6379> mset name1 zhangsan name2 lisi name3 wangwu #同时设置多个值  OK  127.0.0.1:6379> keys *  1) "name2"  2) "name3"  3) "name1"
```

#### 获取多个值

MGET key [key …]

```bash
  127.0.0.1:6379> mget name1 name2 name3 #同时获取多个值  1) "zhangsan"  2) "lisi"  3) "wangwu"
```

#### 如果不存在同时设置多个值

MSETNX key value [key value …]

```bash
  127.0.0.1:6379> msetnx name1 zhangsan name4 zhaoliu #设置失败，msetnx是原子性操作有失败，则全失败  (integer) 0  127.0.0.1:6379> keys *  1) "name2"  2) "name3"  3) "name1"
```

注：msetnx是原子性操作有失败，则全失败

#### 先获取再赋值

GETSET key value

```bash
  127.0.0.1:6379> FLUSHALL #清空数据库  OK  127.0.0.1:6379> keys *        (empty list or set)  127.0.0.1:6379> getset name1 zhangsan #获取name1，并将张三赋值给name1  (nil) #name1不存在，使用获取到的是nil  127.0.0.1:6379> get name1   "zhangsan"  127.0.0.1:6379> getset name1 lisi  "zhangsan"   127.0.0.1:6379> get name1  "lisi"
```

> 具体用例

```bash
127.0.0.1:6379> set name1 zhangsan #设置值OK127.0.0.1:6379> get name1 #获取值"zhangsan"127.0.0.1:6379> append name1 fawaikuangtu #追加字符串(integer) 20127.0.0.1:6379> get name1"zhangsanfawaikuangtu"127.0.0.1:6379> STRLEN name1 #获取值的长度(integer) 20127.0.0.1:6379> keys *1) "name1"127.0.0.1:6379> APPEND name2 lisi #追加字符串，如果不存在就添加(integer) 4127.0.0.1:6379> get name2"lisi"127.0.0.1:6379> set i 1OK127.0.0.1:6379> type istring127.0.0.1:6379> INCR i #自增1(integer) 2 #返回结果127.0.0.1:6379> DECR i #自减1(integer) 1 127.0.0.1:6379> DECR i(integer) 0127.0.0.1:6379> INCRBY i 7 #自增7(integer) 7127.0.0.1:6379> DECRBY i 10 #自减10(integer) -3127.0.0.1:6379> GETRANGE name1 7 11 #获取索引7-11的值"nfawa"127.0.0.1:6379> GETRANGE name1 7 -1 #获取索引7-结束的值"nfawaikuangtu"127.0.0.1:6379> get name1"zhangsanfawaikuangtu"127.0.0.1:6379> SETRANGE name1 4 si #替换4位置开始的字符串(integer) 20127.0.0.1:6379> get name1"zhansianfawaikuangtu"127.0.0.1:6379> setex name3 15 wangwu #设置name3的值为wangwu 15秒后过期OK127.0.0.1:6379> ttl name3(integer) 12127.0.0.1:6379> ttl name3(integer) 8127.0.0.1:6379> ttl name3(integer) 7127.0.0.1:6379> get name3"wangwu"127.0.0.1:6379> ttl name3(integer) 0127.0.0.1:6379> ttl name3(integer) -2127.0.0.1:6379> setnx name4 zhaoliu #如果不存在name4，则设置name4(integer) 1127.0.0.1:6379> keys *1) "name4"2) "i"3) "name2"4) "name1"127.0.0.1:6379> setnx name4 zhangwuji(integer) 0 #返回0，设置失败，因为已经设置了name4127.0.0.1:6379> FLUSHdb #清空数据库OK127.0.0.1:6379> mset name1 zhangsan name2 lisi name3 wangwu #同时设置多个值OK127.0.0.1:6379> keys *1) "name2"2) "name3"3) "name1"127.0.0.1:6379> mget name1 name2 name3 #同时获取多个值1) "zhangsan"2) "lisi"3) "wangwu"127.0.0.1:6379> msetnx name1 zhangsan name4 zhaoliu #设置失败，msetnx是原子性操作有失败，则全失败(integer) 0127.0.0.1:6379> keys *1) "name2"2) "name3"3) "name1"127.0.0.1:6379> FLUSHALL #清空数据库OK127.0.0.1:6379> keys *      (empty list or set)127.0.0.1:6379> getset name1 zhangsan #获取name1，并将张三赋值给name1(nil) #name1不存在，使用获取到的是nil127.0.0.1:6379> get name1 "zhangsan"127.0.0.1:6379> getset name1 lisi"zhangsan" 127.0.0.1:6379> get name1"lisi"
```

## List（列表）

List列表L开头代表从列表头部对List进行操作，R开头代表从List尾部进行操作

#### 查看List中值

LRANGE key start stop

```bash
  127.0.0.1:6379> FLUSHDB  127.0.0.1:6379> LPUSH list one  (integer) 1  127.0.0.1:6379> LPUSH list two  (integer) 2  127.0.0.1:6379> LPUSH list three  (integer) 3  127.0.0.1:6379> LRANGE list 0 -1  1) "three"  2) "two"  3) "one"
```

#### 将一个值多个值添加到列表

##### 添加到头部

LPUSH key value [value …]

```
127.0.0.1:6379> FLUSHDB127.0.0.1:6379> LPUSH list one(integer) 1127.0.0.1:6379> LPUSH list two(integer) 2127.0.0.1:6379> LPUSH list three(integer) 3127.0.0.1:6379> LRANGE list 0 -11) "three"2) "two"3) "one"
```

##### 添加到尾部

RPUSH key value [value …]

```
127.0.0.1:6379> RPUSH list four #将four添加到list尾部(integer) 4127.0.0.1:6379> LRANGE list 0 -11) "three"2) "two"3) "one"4) "four"
```

#### 移除List中值

LPOP key

```bash
  127.0.0.1:6379> LPOP list #移除头元素  "three"  127.0.0.1:6379> LRANGE list 0 -1  1) "two"  2) "one"  3) "four"
```

#### 查询list中某个索引下的值

LINDEX KEY INDEX

```bash
  127.0.0.1:6379> LINDEX list 1  "one"  127.0.0.1:6379> LINDEX list 0  "two"
```

#### 查询list中的长度

LLEN KEY

```bash
  127.0.0.1:6379> LLEN list  (integer) 3
```

#### 移除任意个指定的值

LREM key count value

```bash
  127.0.0.1:6379> LRANGE list 0 -1  1) "two"  2) "one"  3) "four"  127.0.0.1:6379> LREM list 1 one #移除1个one，从头开始  (integer) 1  127.0.0.1:6379> LRANGE list 0 -1  1) "two"  2) "four"
```

#### 修建截断

LTRIM key start stop

```bash
  127.0.0.1:6379> FLUSHALL  OK  127.0.0.1:6379> LPUSH list one  (integer) 1  127.0.0.1:6379> LPUSH list two  (integer) 2  127.0.0.1:6379> LPUSH list three  (integer) 3  127.0.0.1:6379> LTRIM list 1 2 #取出list中1-2的元素并付给list  OK   127.0.0.1:6379> LRANGE  list 0 -1  1) "two"  2) "one"
```

#### 将list最后一个值弹出并赋值到otherlist中

RPOPLPUSH source destination

```bash
  127.0.0.1:6379> LRANGE  list 0 -1
  1) "two"
  2) "one"
  127.0.0.1:6379> keys *
  1) "list"
  127.0.0.1:6379> RPOPLPUSH list otherlist #将list的最后一个值弹出并添加到otherlist
  "one"
  127.0.0.1:6379> LRANGE list 0 -1
  1) "two"
  127.0.0.1:6379> LRANGE otherlist 0 -1
  1) "one"
```

#### 列表中指定位置更新值

LSET key index value

注：如果index不存在会报错

```bash
  127.0.0.1:6379> LRANGE list 0 -1
  1) "two"
  127.0.0.1:6379> lset list 0 one #设置list的第0个值为one
  OK
  127.0.0.1:6379> LRANGE list 0 -1
  1) "one"
  127.0.0.1:6379> lset list 1 two
  (error) ERR index out of range  #下标1不存在，报错！
```

#### 列表中插入值

LINSERT key BEFORE|AFTER pivot value

```bash
  127.0.0.1:6379> LRANGE list 0 -1
  1) "one"
  127.0.0.1:6379> LINSERT list before one two #再one前面添加two
  (integer) 2
  127.0.0.1:6379> LINSERT list after one three #再one后面添加three
  (integer) 3
  127.0.0.1:6379> LINSERT list before one two
  (integer) 4
  127.0.0.1:6379> LRANGE list 0 -1
  1) "two"
  2) "two"
  3) "one"
  4) "three"
```

> 具体用例

```bash
127.0.0.1:6379> FLUSHDB
127.0.0.1:6379> LPUSH list one
(integer) 1
127.0.0.1:6379> LPUSH list two
(integer) 2
127.0.0.1:6379> LPUSH list three
(integer) 3
127.0.0.1:6379> LRANGE list 0 -1
1) "three"
2) "two"
3) "one"
127.0.0.1:6379> FLUSHDB
127.0.0.1:6379> LPUSH list one
(integer) 1
127.0.0.1:6379> LPUSH list two
(integer) 2
127.0.0.1:6379> LPUSH list three
(integer) 3
127.0.0.1:6379> LRANGE list 0 -1
1) "three"
2) "two"
3) "one"
127.0.0.1:6379> RPUSH list four #将four添加到list尾部
(integer) 4
127.0.0.1:6379> LRANGE list 0 -1
1) "three"
2) "two"
3) "one"
4) "four"
127.0.0.1:6379> LPOP list #移除头元素
"three"
127.0.0.1:6379> LRANGE list 0 -1
1) "two"
2) "one"
3) "four"
127.0.0.1:6379> LINDEX list 1
"one"
127.0.0.1:6379> LINDEX list 0
"two"
127.0.0.1:6379> LLEN list
(integer) 3
127.0.0.1:6379> LRANGE list 0 -1
1) "two"
2) "one"
3) "four"
127.0.0.1:6379> LREM list 1 one #移除1个one，从头开始
(integer) 1
127.0.0.1:6379> LRANGE list 0 -1
1) "two"
2) "four"
127.0.0.1:6379> FLUSHALL
OK
127.0.0.1:6379> LPUSH list one
(integer) 1
127.0.0.1:6379> LPUSH list two
(integer) 2
127.0.0.1:6379> LPUSH list three
(integer) 3
127.0.0.1:6379> LTRIM list 1 2 #取出list中1-2的元素并付给list
OK 
127.0.0.1:6379> LRANGE  list 0 -1
1) "two"
2) "one"
127.0.0.1:6379> LRANGE  list 0 -1
1) "two"
2) "one"
127.0.0.1:6379> keys *
1) "list"
127.0.0.1:6379> RPOPLPUSH list otherlist #将list的最后一个值弹出并添加到otherlist
"one"
127.0.0.1:6379> LRANGE list 0 -1
1) "two"
127.0.0.1:6379> LRANGE otherlist 0 -1
1) "one"
127.0.0.1:6379> LRANGE list 0 -1
1) "two"
127.0.0.1:6379> lset list 0 one #设置list的第0个值为one
OK
127.0.0.1:6379> LRANGE list 0 -1
1) "one"
127.0.0.1:6379> lset list 1 two
(error) ERR index out of range  #下标1不存在，报错！
127.0.0.1:6379> LRANGE list 0 -1
1) "one"
127.0.0.1:6379> LINSERT list before one two #再one前面添加two
(integer) 2
127.0.0.1:6379> LINSERT list after one three #再one后面添加three
(integer) 3
127.0.0.1:6379> LINSERT list before one two
(integer) 4
127.0.0.1:6379> LRANGE list 0 -1
1) "two"
2) "two"
3) "one"
4) "three"
```

## Set（集合）

set中的值不能重复

#### set中添加值

SADD key member [member …]

```bash
  127.0.0.1:6379> SADD set one #添加one到set
  (integer) 1
  127.0.0.1:6379> SADD set two
  (integer) 1
  127.0.0.1:6379> SADD set three
  (integer) 1
```

#### 查看set中的值

SMEMBERS key

```bash
  127.0.0.1:6379> SMEMBERS set #查看set的值
  1) "three"
  2) "one"
  3) "two"
```

#### 判断set中是否存在该value

SISMEMBER key member

```bash
  127.0.0.1:6379> SISMEMBER set one #查看set中是否有one
  (integer) 1
  127.0.0.1:6379> SISMEMBER set four
  (integer) 0
```

#### 获取set中元素个数

SCARD key

```bash
  127.0.0.1:6379> SCARD set #获取set中元素个数
  (integer) 3
```

#### 移除set中某个元素

SREM key member [member …]

```bash
  127.0.0.1:6379> SREM set one #从set中移除one
  (integer) 1
  127.0.0.1:6379> SMEMBERS set
  1) "three"
  2) "two"
```

#### 随机从set中抽取指定个数元素

SRANDMEMBER key [count]

```bash
  127.0.0.1:6379> SRANDMEMBER set #随机从set中去一个值
  "two"
  127.0.0.1:6379> SRANDMEMBER set
  "three"
  127.0.0.1:6379> SRANDMEMBER set
  "three"
```

#### 随机弹出值

SPOP key [count]

```bash
  127.0.0.1:6379> SMEMBERS set
  1) "three"
  2) "two"
  127.0.0.1:6379> SPOP set #随机删除set中一个数
  "three"
  127.0.0.1:6379> SMEMBERS set
  1) "two"
```

#### 将set中一个指定的值，移动otherset中

SMOVE source destination member

```bash
  127.0.0.1:6379> sadd set2 zhangsan lisi wangwu
  (integer) 3
  127.0.0.1:6379> SMEMBERS set2
  1) "zhangsan"
  2) "wangwu"
  3) "lisi"
  127.0.0.1:6379> SMEMBERS set
  1) "two"
  127.0.0.1:6379> SMOVE set2 set zhangsan #将set2中的zhangsan移到set
  (integer) 1
  127.0.0.1:6379> SMEMBERS set2
  1) "wangwu"
  2) "lisi"
  127.0.0.1:6379> SMEMBERS set
  1) "zhangsan"
  2) "two"
```

#### set的差并交

```bash
  127.0.0.1:6379> FLUSHALL
  OK
  127.0.0.1:6379> sadd set1 a b c
  (integer) 3
  127.0.0.1:6379> sadd set2 c d e
  (integer) 3
```

#### 差集

SDIFF key [key …]

```
127.0.0.1:6379> SDIFF set1 set2  #set1相对于set2的差集
1) "a"
2) "b"
```

#### 交集

SINTER key [key …]

```
127.0.0.1:6379> SINTER set1 set2 #set1与set2的交集
1) "c"
```

#### 并集

SUNION key [key …]

```
127.0.0.1:6379> SUNION set1 set2 #set1和set2的并集
1) "a"
2) "b"
3) "c"
4) "d"
5) "e"
```

> 具体用例

```bash
127.0.0.1:6379> SADD set one #添加one到set
(integer) 1
127.0.0.1:6379> SADD set two
(integer) 1
127.0.0.1:6379> SADD set three
(integer) 1
127.0.0.1:6379> SMEMBERS set #查看set的值
1) "three"
2) "one"
3) "two"
127.0.0.1:6379> SISMEMBER set one #查看set中是否有one
(integer) 1
127.0.0.1:6379> SISMEMBER set four
(integer) 0
127.0.0.1:6379> SCARD set #获取set中元素个数
(integer) 3
127.0.0.1:6379> SREM set one #从set中移除one
(integer) 1
127.0.0.1:6379> SMEMBERS set
1) "three"
2) "two"
127.0.0.1:6379> SRANDMEMBER set #随机从set中去一个值
"two"
127.0.0.1:6379> SRANDMEMBER set
"three"
127.0.0.1:6379> SRANDMEMBER set
"three"
127.0.0.1:6379> SMEMBERS set
1) "three"
2) "two"
127.0.0.1:6379> SPOP set #随机删除set中一个数
"three"
127.0.0.1:6379> SMEMBERS set
1) "two"
127.0.0.1:6379> sadd set2 zhangsan lisi wangwu
(integer) 3
127.0.0.1:6379> SMEMBERS set2
1) "zhangsan"
2) "wangwu"
3) "lisi"
127.0.0.1:6379> SMEMBERS set
1) "two"
127.0.0.1:6379> SMOVE set2 set zhangsan #将set2中的zhangsan移到set
(integer) 1
127.0.0.1:6379> SMEMBERS set2
1) "wangwu"
2) "lisi"
127.0.0.1:6379> SMEMBERS set
1) "zhangsan"
2) "two"
127.0.0.1:6379> FLUSHALL
OK
127.0.0.1:6379> sadd set1 a b c
(integer) 3
127.0.0.1:6379> sadd set2 c d e
(integer) 3
127.0.0.1:6379> SDIFF set1 set2  #set1相对于set2的差集
1) "a"
2) "b"
127.0.0.1:6379> SINTER set1 set2 #set1与set2的交集
1) "c"
127.0.0.1:6379> SUNION set1 set2 #set1和set2的并集
1) "a"
2) "b"
3) "c"
4) "d"
5) "e"
```

## Hash(哈希）

Hash本质就是Map集合

#### 存值

HSET key field value

```bash
  127.0.0.1:6379> hset hash one zhangsan #将key为one，value为zhangsan存入hash中
  (integer) 1
  127.0.0.1:6379> hset hash two lisi
  (integer) 1
```

#### 取值

HGET key field

```bash
  127.0.0.1:6379> hget hash one
  "zhangsan"
  127.0.0.1:6379> hget hash two
  "lisi"
```

#### 批量存值

MSET key field value [field value …]

```bash
  127.0.0.1:6379> HMset hash three wangwu four zhaoliu #同时存入多个key-value至hash
  OK
```

#### 批量取值

HMGET key field [field …]

```bash
  127.0.0.1:6379> HMGET hash three four #根据key同时取出hash中多个值
  1) "wangwu"
  2) "zhaoliu"
```

#### 获取hash中所有键值对

HGETALL key

```bash
  127.0.0.1:6379> HGETALL hash
  1) "one"
  2) "zhangsan"
  3) "two"
  4) "lisi"
  5) "three"
  6) "wangwu"
  7) "four"
  8) "zhaoliu"
```

#### 根据key删除hash中指定的键值对

HDEL key field [field …]

```bash
  127.0.0.1:6379> HDEL hash two #删除key值为2的hash键值对
  (integer) 1
  127.0.0.1:6379> HGETALL hash
  1) "one"
  2) "zhangsan"
  3) "three"
  4) "wangwu"
  5) "four"
  6) "zhaoliu"
```

#### 判断hash键是否存在

HEXISTS key field

```bash
127.0.0.1:6379> HGETALL hash
1) "one"
2) "zhangsan"
3) "three"
4) "wangwu"
5) "four"
6) "zhaoliu"
127.0.0.1:6379> HEXISTS hash one #判断键为one的键值对是否存在
(integer) 1 #存在
127.0.0.1:6379> HEXISTS hash two
(integer) 0 #不存在
```

#### 获取所有的key或者所有的value

##### 获取所有的key(键)

HKEYS key

```bash
127.0.0.1:6379> hkeys hash
1) "one"
2) "three"
3) "four"
```

##### 获取所有的value(值)

HVALS key

```bash
127.0.0.1:6379> HVALS hash
1) "zhangsan"
2) "wangwu"
3) "zhaoliu"
```

#### 自增

HINCRBY key field increment

```bash
127.0.0.1:6379> hset hash one 7 
(integer) 1
127.0.0.1:6379> HINCRBY hash one 1 #使hash中one的值加1
(integer) 8
127.0.0.1:6379> HGET hash one 
"8"
```

#### 不存在则设置值

HSETNX key field value

```bash
127.0.0.1:6379> HSETNX hash two zhangsan #如歌不存在键two这添加键值对<two,zhangsan>
(integer) 1
127.0.0.1:6379> HSETNX hash two lisi
(integer) 0 #已近存在所有错误，返回0
```

> 具体用例

```bash
  127.0.0.1:6379> hset hash one zhangsan #将key为one，value为zhangsan存入hash中
  (integer) 1
  127.0.0.1:6379> hset hash two lisi
  (integer) 1
  127.0.0.1:6379> hget hash one
  "zhangsan"
  127.0.0.1:6379> hget hash two
  "lisi"
  127.0.0.1:6379> HMset hash three wangwu four zhaoliu #同时存入多个key-value至hash
  OK
127.0.0.1:6379> HMGET hash three four #根据key同时取出hash中多个值
  1) "wangwu"
  2) "zhaoliu"
127.0.0.1:6379> HGETALL hash
  1) "one"
  2) "zhangsan"
  3) "two"
  4) "lisi"
  5) "three"
  6) "wangwu"
  7) "four"
  8) "zhaoliu"
127.0.0.1:6379> HDEL hash two #删除key值为2的hash键值对
  (integer) 1
  127.0.0.1:6379> HGETALL hash
  1) "one"
  2) "zhangsan"
  3) "three"
  4) "wangwu"
  5) "four"
  6) "zhaoliu"
127.0.0.1:6379> HGETALL hash
1) "one"
2) "zhangsan"
3) "three"
4) "wangwu"
5) "four"
6) "zhaoliu"
127.0.0.1:6379> HEXISTS hash one #判断键为one的键值对是否存在
(integer) 1 #存在
127.0.0.1:6379> HEXISTS hash two
(integer) 0 #不存在
127.0.0.1:6379> hkeys hash
1) "one"
2) "three"
3) "four"
127.0.0.1:6379> HVALS hash
1) "zhangsan"
2) "wangwu"
3) "zhaoliu"
127.0.0.1:6379> hset hash one 7 
(integer) 1
127.0.0.1:6379> HINCRBY hash one 1 #使hash中one的值加1
(integer) 8
127.0.0.1:6379> HGET hash one 
"8"
127.0.0.1:6379> HSETNX hash two zhangsan #如歌不存在键two这添加键值对<two,zhangsan>
(integer) 1
127.0.0.1:6379> HSETNX hash two lisi
(integer) 0 #已近存在所有错误，返回0
```

## Zset(有序集合)

在set的基础上增加排序

#### 添加数据

zadd key [NX|XX] [CH] [INCR] score member [score member …]

```bash
127.0.0.1:6379> zadd zset 1 one
(integer) 1
127.0.0.1:6379> zadd zset 2 two
(integer) 1
127.0.0.1:6379> zadd zset 3 three
(integer) 1
```

#### 获取数据

ZRANGE key start stop [WITHSCORES]

```bash
127.0.0.1:6379> ZRANGE zset 0 -1 #获取所有数据
1) "one"
2) "two"
3) "three"
```

#### 排序查询

ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]

```bash
127.0.0.1:6379> zadd salary 1500 zhangsan 2500 lisi 3500 wangwu 
(integer) 3
127.0.0.1:6379> ZRANGE salary 0 -1
1) "zhangsan"
2) "lisi"
3) "wangwu"
127.0.0.1:6379> ZRANGEBYSCORE salary -inf +inf #升序查询所有
1) "zhangsan"
2) "lisi"
3) "wangwu"
127.0.0.1:6379> ZRANGEBYSCORE salary -inf +inf withscores #升序查询所有并带上value
1) "zhangsan"
2) "1500"
3) "lisi"
4) "2500"
5) "wangwu"
6) "3500"
127.0.0.1:6379> ZRANGEBYSCORE salary 1000 3000 #升序查询key在1000-3000的value
1) "zhangsan"
2) "lisi"
127.0.0.1:6379> ZREVRANGE salary 0 -1 #降序查询所有元素并
1) "wangwu"
2) "lisi"
3）"zhangsan"
```

#### 移除指定元素

ZREM key member [member …]

```bash
127.0.0.1:6379> ZRANGE salary 0 -1
1) "zhangsan"
2) "lisi"
3) "wangwu"
127.0.0.1:6379> ZREM salary zhangsan #移除zhangsan
(integer) 1
127.0.0.1:6379> ZRANGE salary 0 -1
1) "lisi"
2) "wangwu"
```

#### 获取集合中的个数

ZCARD key

```bash
127.0.0.1:6379> ZCARD salary
(integer) 2
```

#### 获取指定区间的数量

ZCOUNT key min max

```bash
127.0.0.1:6379> FLUSHALL
OK
127.0.0.1:6379> zadd zset 1 one 2 two 3 three 4 four
(integer) 4
127.0.0.1:6379> ZCOUNT zset 1 3
(integer) 3
```

> 具体用例

```bash
127.0.0.1:6379> zadd zset 1 one
(integer) 1
127.0.0.1:6379> zadd zset 2 two
(integer) 1
127.0.0.1:6379> zadd zset 3 three
(integer) 1
127.0.0.1:6379> ZRANGE zset 0 -1 #获取所有数据
1) "one"
2) "two"
3) "three"
127.0.0.1:6379> zadd salary 1500 zhangsan 2500 lisi 3500 wangwu 
(integer) 3
127.0.0.1:6379> ZRANGE salary 0 -1
1) "zhangsan"
2) "lisi"
3) "wangwu"
127.0.0.1:6379> ZRANGEBYSCORE salary -inf +inf #升序查询所有
1) "zhangsan"
2) "lisi"
3) "wangwu"
127.0.0.1:6379> ZRANGEBYSCORE salary -inf +inf withscores #升序查询所有并带上value
1) "zhangsan"
2) "1500"
3) "lisi"
4) "2500"
5) "wangwu"
6) "3500"
127.0.0.1:6379> ZRANGEBYSCORE salary 1000 3000 #升序查询key在1000-3000的value
1) "zhangsan"
2) "lisi"
127.0.0.1:6379> ZREVRANGE salary 0 -1 #降序查询所有元素并
1) "wangwu"
2) "lisi"
3）"zhangsan"
127.0.0.1:6379> ZRANGE salary 0 -1
1) "zhangsan"
2) "lisi"
3) "wangwu"
127.0.0.1:6379> ZREM salary zhangsan #移除zhangsan
(integer) 1
127.0.0.1:6379> ZRANGE salary 0 -1
1) "lisi"
2) "wangwu"
127.0.0.1:6379> ZCARD salary
(integer) 2
127.0.0.1:6379> FLUSHALL
OK
127.0.0.1:6379> zadd zset 1 one 2 two 3 three 4 four
(integer) 4
127.0.0.1:6379> ZCOUNT zset 1 3
(integer) 3
```