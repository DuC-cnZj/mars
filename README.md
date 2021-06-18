<h1 align="center">Mars</h1>
<p align="center">专为devops而生，10秒内部署一个应用。</p>
<br><br>

## 💡 简介

[Mars](https://github.com/DuC-cnZj/mars) 是一款专门为devops服务的一款应用，基于 kubernetes 之上，可以在短短几秒内部署一个和生产环境一模一样的应用。它打通了 gitlab、kubernetes、helm，通过 gitlab ci 构建镜像，然后通过kubernetes 部署高可用应用，一气呵成。

## 🗺️ 背景

随着 devops 概念的兴起，现在软件开发不仅要求开发效率高，而且还要求部署便捷，最好能做到流水线开发打包测试上线一条龙服务。
[Mars](https://github.com/DuC-cnZj/mars) 由此而生，它打通了打包、测试、部署，基于 gitlab ci/cd 做到任何人不管是开发大牛，还是不懂代码的产品小白，都能在10秒部署一个生产级别的应用。真真做到一教即会，高效生产。

## ✨  特性

* 支持基于 helm charts 开发的任何应用。
* 支持自动配置 https 域名。
* 支持高可用，弹性部署。
* 支持命令行操作。
* 支持查看容器日志。
* 支持查看容器cpu和内存使用情况。

## 🛠️ 使用文档

1. 如果本地有 (golang > 1.16) 环境，则可以使用
```bash
go get github.com/duc-cnZj/mars
```
初始化配置
```bash
mars init
```

2. 在 kubernetes 内部署（推荐）

```bash
helm repo add mars-charts https://duc-cnzj.github.io/mars-charts/
# 这里需要自行配置相关参数
helm show values mars-charts/mars > mars-values.yaml
helm upgrade --install mars mars-charts/mars -f mars-values.yaml
```

## 🏗 preview

> xuanji golang 版本。

https://github.com/Lick-Dog-Club/xuanji-k8s-all-in-one


## TODO

- ui 美化
- 配置管理界面可以直接查看 `.mars.yaml` 配置文件