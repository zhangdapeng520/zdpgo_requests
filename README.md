# zdpgo_requests

Golang中用于发送HTTP请求的库

## 版本历史

- v0.1.0 2022/04/09 新增GET和POST请求
- v0.1.1 2022/04/11 POST的map默认当表单数据
- v0.1.2 2022/04/11 添加忽略URL解析错误的请求方法
- v0.1.3 2022/04/12 支持POST纯文本数据
- v0.1.4 2022/04/12 代码重构
- v0.1.5 2022/04/13 支持任意类型HTTP请求
- v0.1.6 2022/04/13 支持设置代理
- v0.1.7 2022/04/13 支持发送JSON数据
- v0.1.8 2022/04/16 解决部分URL无法正常请求的BUG
- v0.1.9 2022/04/18 BUG修复：header请求头重复
- v0.2.0 2022/04/18 新增：获取请求和响应详情
- v0.2.1 2022/04/20 新增：获取响应状态码
- v0.2.2 2022/04/20 新增：下载文件
- v0.2.3 2022/04/21 新增：文件上传
- v0.2.4 2022/04/22 新增：支持上传FS文件系统文件
- v0.2.5 2022/04/28 新增：检查重定向和请求消耗时间
- v0.2.6 2022/05/06 新增：根据字节数组上传文件
- v0.2.7 2022/05/08 新增：根据超时时间发送POST请求并携带JSON数据
- v0.2.8 2022/05/09 BUG修复：修复POST超时单位不为秒的BUG
- v0.2.9 2022/05/17 升级：日志组件升级
- v0.3.1 2022/05/17 新增：忽略HTTPS证书校验
- v0.3.2 2022/05/17 升级：升级random组件
- v0.3.3 2022/05/18 优化：整体架构优化
- v0.3.4 2022/05/18 新增：初始化数据的方法
- v0.3.5 2022/05/19 BUG修复：修复UserAgent不正确
- v0.3.6 2022/05/19 优化：整体架构优化
- v0.3.7 2022/05/20 新增：设置请求超时时间
- v0.3.8 2022/05/25 新增：根据字节数组上传文件
- v0.3.9 2022/05/26 优化：优化字节数组上传方法
- v0.4.0 2022/05/27 新增：任意方法的JSON请求
- v0.4.1 2022/05/27 优化：精简代码
- v0.4.2 2022/05/28 新增：任意方法的Text请求
- v0.4.3 2022/06/01 新增：Response不为nil的方法
- v0.4.4 2022/06/06 新增：POST提交ECC加密数据
- v0.4.5 2022/06/11 新增：POST提交ECC加密文本数据
- v0.4.6 2022/06/11 新增：POST提交AES加密数据
- v0.4.7 2022/06/14 新增：支持两次请求状态码校验
- v0.4.8 2022/06/14 新增：自定义X-Author请求头
- v0.4.9 2022/06/15 优化：优化日志和Author请求头
- v0.5.3 2022/06/20 优化：优化日志
- v0.5.4 2022/06/22 新增：请求限流
- v0.5.5 2022/06/27 优化：请求和响应详情没有不报错
- v0.5.6 2022/06/28 优化：移除日志组件
- v0.5.7 2022/06/28 BUG修复：修复日志移除后的相关BUG
- v0.5.8 2022/07/08 升级：升级password组件

## 使用示例

请查看 examples 目录