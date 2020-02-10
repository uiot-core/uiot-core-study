# C-SDK的使用

## 参考文档

参考[文档](https://docs.ucloud.cn/iot/uiot-core/device_develop_guide/c_sdk_example/csdkquickstart)


## C-SDK下载

```
git clone https://github.com/ucloud/ucloud-iot-device-sdk-c
```


## 设备数据上下行

### 静态注册

示例代码：
```
./samples/mqtt/mqtt_sample.c
```

修改：
```
#define UIOT_MY_PRODUCT_SN            "PRODUCT_SN"

#define UIOT_MY_DEVICE_SN             "DEVICE_SN"

#define UIOT_MY_DEVICE_SECRET         "DEVICE_SECRET"
```

### 动态注册

示例代码：
```
./samples/dynamic_auth/dynamic_auth_sample.c
```

修改：
```
#define UIOT_MY_PRODUCT_SN            "PRODUCT_SN"

#define UIOT_MY_DEVICE_SN             "DEVICE_SN"

#define UIOT_MY_PRODUCT_SECRET        "PRODUCT_SECRET"

```


