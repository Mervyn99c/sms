# TODO List  
## 1 服务
- [x] gateway   
- [ ] batch
- [ ] surl
- [ ] iam

## 2 项目结构
提取公共方法到pkg  
平铺式

## 3 架构图
参考已有文档

## 4 设计文档

## 对接短信供应商
阿里云
dummy fake

## 功能点
手机号白名单 

## 设计模式
### 创建供应商 
供应商枚举，根据类型创建对应的供应商实例  
供应商使用接口，一个impl  
多个接口方法：  
template： sms mms   
batch :sms mms  
getBackList  

provider分别实现template和batch接口  
使用泛型，由于有template和batch两种接口，请求参数和响应参数不一样   
所以用泛型 T R  

factory创建provider后，provider调用SMSTemplateProvider的delivery方法
  

### 规则校验

## 中间件
签名校验
admin jwt登录

## 流程
### 验证码
判断是否是验证码短信，是就发送
发送成功就缓存
根据类型发送到不同的队列
记录短信发送记录 异步到内存，内存到mq
- TODO 最后做 发送失败记录KPI告警
记录规则检验失败记录
使用worker 协程池消费
将KPI抽象为SDK
?如何使用redis集群 封装redis的key和类型
## 用户 Casbin RBAC




