# Autonomy

Autonomy 是一个尊重个人主权的个人数据自动化框架。

## 核心原则

1. 不依赖云服务
2. 不处理大文件同步

## 目标用户

1. 具有一定技术能力的个人

##  功能清单

### autonomy-agentd

- [ ] 监听并执行定时任务
- [ ] 数据上报

### autonomy-core

- [ ] 接受上报数据存入数据库

## 使用

### autonomy-agentd

```
autonomy-agentd --config config.yaml --workflows workflows_dir
```