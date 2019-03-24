# goconfig
安装
go get github.com/Unknwon/goconfig

或
gopm get github.com/Unknwon/goconfig

API 文档
[Go Walker](http://gowalker.org/github.com/Unknwon/goconfig).

示例
请查看 [config.ini](conf/config.ini) 文件作为使用示例。

用例
函数 LoadConfigFile 加载一个或多个文件，然后返回一个类型为 ConfigFile 的变量。
GetValue 可以简单的获取某个值。
像 Bool、Int、Int64 这样的方法会直接返回指定类型的值。
以 Must 开头的方法不会返回错误，但当错误发生时会返回零值。
SetValue 可以设置某个值。
DeleteKey 可以删除某个键。
最后，SaveConfigFile 可以保持您的配置到本地文件系统。
使用方法 Reload 可以重载您的配置文件。