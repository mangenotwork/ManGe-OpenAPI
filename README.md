## extras 临时演员
> 附加功能微服务,开箱即用。目的是扩展整个项目的功能性，提供api给前后端使用；

## 通讯协议支持
- http/s: 使用jsonp输出主要客户端使用
- grpc: 服务端远程调用
- tcp: 服务端使用
- udp: 服务端使用

## 服务

#### BlockWord 屏蔽词服务
> 屏蔽词增删该查，词语白名单等; 提供 http/s, grpc api

>  [点击查看](https://github.com/mangenotwork/extras/tree/master/apps/BlockWord)

#### ShortLink 短链接服务
>  短链接生成,管理; 提供 http/s,grpc api

>  [点击查看](https://github.com/mangenotwork/extras/tree/master/apps/ShortLink)

#### Push 推送服务
> 提供 websocket, tcp, udp的推送, 支持 发布订阅模式,广播模式,组播模式等

>  [点击查看](https://github.com/mangenotwork/extras/tree/master/apps/Push)

#### ImgHelper 图片功能服务
> 生成二维码, 图片压缩， 图片水印， 生成文字图片， 图片固定剪切， 生成gif， 图片固定拼接， 图片基础信息获取，固定旋转， 格式转换;
> 提供 http/s, grpc api

>  [点击查看](https://github.com/mangenotwork/extras/tree/master/apps/ImgHelper)

#### WordHelper 文字处理服务
> 分词， OCR, 翻译, 加密解密, 文本内容的领域信息, 文本相似度, 彩票开奖, 拼音, 标签提取,
> 提供 http/s, grpc api

>  [点击查看](https://github.com/mangenotwork/extras/tree/master/apps/WordHelper)

#### ServiceTable 分布式数据表主要用于配置中心, 发现注册服务 
> 轻量级分布式数据表, 主要用于配置中心, 发现注册服务; 参考raft算法 

>  [点击查看](https://github.com/mangenotwork/extras/tree/master/apps/ServiceTable)

#### IM 即时聊天功能服务
> 即时聊天功能; 提供 websocket, tcp, udp

- IM-Conn 提供连接与网络通讯 >  [点击查看](https://github.com/mangenotwork/extras/tree/master/apps/IM-Conn)
- IM-Msg  提供消息业务 >  [点击查看](https://github.com/mangenotwork/extras/tree/master/apps/IM-Msg)
- IM-User 提供用户业务 >  [点击查看](https://github.com/mangenotwork/extras/tree/master/apps/IM-User)
- IM-Test 测试 >  [点击查看](https://github.com/mangenotwork/extras/tree/master/apps/IM-Test)

#### LogCenter  日志中心,收集微服务日志
> 日志中心,收集微服务日志,可以监控微服务的http请求日志,grpc请求日志与链路日志等等; 简化运维工作

>  [点击查看](https://github.com/mangenotwork/extras/tree/master/apps/LogCenter)

#### 


## LICENSE : MIT License



