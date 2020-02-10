## C-SDK发送二进制数据
git clone https://github.com/ucloud/ucloud-iot-device-sdk-c

修改：`./samples/mqtt/mqtt_sample.c    _publish_msg()`

```
    char topic_content[15] = {0x5A,0x5A,0xF5,0xE1,0xD2,0x3C,0x2A,0x02,0x7F,0x64,0x21,0x0A,0x0D};

    HAL_Printf("send message:");
    for (int i = 0; i < 10; i++)
    {
            HAL_Printf("%x ", topic_content[i]&0xff);
    }
	
	pub_params.payload_len = 13;
```


## 转发到HTTP

```golang
func main() {
        http.HandleFunc("/receive", HttpHandler)
        http.ListenAndServe(":8900", nil)
}
```

## 转发到Kafka

使用Kafka Client Library：`github.com/Shopify/sarama`


## 转发到MySQL

mysql -h 10.23.176.199 -u root -p uiotcore123456

create database ruleengintest;
use ruleengintest
CREATE TABLE IF NOT EXISTS `tabletest`(`id` INT UNSIGNED AUTO_INCREMENT,`number`  INT UNSIGNED NOT NULL,PRIMARY KEY ( `id` ))ENGINE=InnoDB;
select * from tabletest order by id desc limit 10;

## 转发到MongoDB

### 安装linux下MongoDB客户端
wget http://downloads.mongodb.org/linux/mongodb-linux-x86_64-rhel70-v4.2-latest.tgz
tar zxvf mongodb-linux-x86_64-rhel70-v4.2-latest.tgz
mkdir -p /usr/local/mongodb
cp mongodb-linux-x86_64-rhel70-4.2.2-51-g7a4e7bb/* /usr/local/mongodb
export PATH=/usr/local/mongodb/bin:$PATH

### 客户端执行命令：
mongo 10.23.174.5/admin -u root -p uiotcore123456
show dbs
use dbtest
db.createCollection("collectiontest")
db.createUser( {user: "usertest",pwd: "123456",roles: [ { role: "readWrite", db: "dbtest" } ]})
db.collectiontest.find({}).sort({"_id":-1}).limit(10)
db.collectiontest.find().count();

## 转发到influxdb

### 安装linux下infulxdb客户端
wget https://dl.influxdata.com/influxdb/releases/influxdb-1.7.9.x86_64.rpm
sudo yum localinstall influxdb-1.7.9.x86_64.rpm

### 客户端执行命令：
influx -host 10.23.238.77 -username testDataBase -password 123456  -precision rfc3339
show databases
use dbtest
show measurements
select * from measurementtest where time > now()-5m
