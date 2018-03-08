# a-bargain-hunting

中证500成分股，近30日最低价top50。

实时数据；动态刷新；抄底利器。
## 构建说明
日k历史数据来源：http://api.finance.ifeng.com/akdaily/?code=sh600848&type=last （部分数据更新不及时）

实时数据来源：http://hq.sinajs.cn/list=sh600848 （相对稳定）

历史数据获取默认重试3次如果还获取不到，30日最低价会显示为10000000。使用中如出现这个情况，可以重启一次解决。

## 编译&运行

    godep get
    godep go build abhunting.go
    ./abhunting

## 使用说明
界面截图：

 ![image](https://github.com/whomm/a-bargain-hunting/raw/master/screenshot.png)


