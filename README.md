### Go 短链接生成器

Short link 短链接服务器
### 什么是短链接
就是把普通网址，转换成比较短的网址。比如：https://liqian.top/XzhYJMkZ

### 原理解析
当我们在浏览器里输入 https://liqian.top/XzhYJMkZ<br>
DNS首先解析获得 https://liqian.top/ 的 IP 地址<br>
当 DNS 获得 IP 地址以后，会向这个地址发送 HTTP请求，查询短码 XzhYJMkZ<br>
https://liqian.top/ 服务器会通过短码 XzhYJMkZ 获取对应的长 URL<br>
请求通过 HTTP 301 转到对应的长 URL http://www.baidu.com<br>

### 本文采用  自增序列算法 + 用户自定义短码

设置 id 自增，一个 10进制 id 对应一个 62进制的数值，1对1，也就不会出现重复的情况。<br>这个利用的就是低进制转化为高进制时，字符数会减少的特性。<br>
可使用redis Incr 实现id自增。<br>  