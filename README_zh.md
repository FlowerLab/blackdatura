# 黑曼陀罗

不可预知的黑暗，死亡和颠沛流离的爱，凡间的无爱与无仇，绝望的爱，不可预知的死亡和爱，被伤害的坚韧创痍的心灵，生的不归之路

[![GoDoc](https://pkg.go.dev/badge/pkg.go.dev/github.com/FlowerLab/blackdatura)](https://pkg.go.dev/pkg.go.dev/github.com/FlowerLab/blackdatura)
[![Sourcegraph](https://sourcegraph.com/github.com/FlowerLab/blackdatura/-/badge.svg)](https://sourcegraph.com/github.com/FlowerLab/blackdatura?badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/FlowerLab/blackdatura)](https://goreportcard.com/report/github.com/FlowerLab/blackdatura)
[![Release](https://img.shields.io/github/v/release/FlowerLab/blackdatura.svg)](https://github.com/FlowerLab/blackdatura/releases)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)


## 介绍

[zap][1] 和 [lumberjack][2] 的二次封装，集成了 [Gin][3] 的 logging 中间件


## 使用

```go
import (
	log "github.com/FlowerLab/blackdatura"
)
```

查看 [示例](example/main.go)


## 鸣谢
- [zap][1]
- [lumberjack][2]
- [Gin][3]
- [ginzap][4]


[1]:https://github.com/uber-go/zap
[2]:https://github.com/natefinch/lumberjack
[3]:https://github.com/gin-gonic/gin
[4]:https://github.com/gin-contrib/zap
