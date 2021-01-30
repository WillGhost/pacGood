# 好用PAC列表
gfwlist不好用，我基于网络上的列表修改了一个适用自己的规则。
# 脚本介绍
域名文件存于`data`目录 有以下2种生成方式
- `v2ray_domain_list.go`生成`dlc.dat`用于v2ray（从 https://github.com/v2fly/domain-list-community 复制）
- `make_pac_text.go`生成`pac.txt`
# 分流策略
### 策略
1. 命中`bypass`的域名走代理
2. 命中`mainland`的域名走直连
3. 命中这些域名后缀`cn`,`com`,`net`,`fm`,`gs`,`[0-9]`(IP访问)，走直连
4. 以上都未命中走代理
### 特色
无论新增还是存量的`org`或`io`后缀都会走代理，但是遇到站点在境内的话需要加入`mainlan`
# 使用方法
### Windows10
`开始` `设置` `网络和Internet 代理`

启用`使用设置脚本`添加`脚本地址`并`保存`

确认本地SOCKS监听的端口与pac一致

重启浏览器测试 [http://ip111.cn/](http://ip111.cn/)

注意：有些代理工具可能自动修改系统代理，比如v2ray需要将代理模式改为“仅开启http代理，不改变系统代理”

### 其他系统参考Win10
