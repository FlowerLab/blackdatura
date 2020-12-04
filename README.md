# Black Datura

Unpredictable darkness, love of death and turbulence, loveless and grudgless in the world, desperate love, unpredictable death and love, hurt but strong mind, unreturnable path and road.

[![GoDoc](https://pkg.go.dev/badge/pkg.go.dev/github.com/FlowerLab/blackdatura)](https://pkg.go.dev/github.com/FlowerLab/blackdatura)
[![Sourcegraph](https://sourcegraph.com/github.com/FlowerLab/blackdatura/-/badge.svg)](https://sourcegraph.com/github.com/FlowerLab/blackdatura?badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/FlowerLab/blackdatura)](https://goreportcard.com/report/github.com/FlowerLab/blackdatura)
[![Release](https://img.shields.io/github/v/release/FlowerLab/blackdatura.svg)](https://github.com/FlowerLab/blackdatura/releases)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)


> [中文文档](README_zh.md)

## Introduction

The secondary packaging of [zap][1] and [lumberjack][2], integrate [Gin][3] and [GORM][5] logging middleware.
Added [Redis][6] and [Kafaka][7] output mode.

## Usage

```go
import (
	log "github.com/FlowerLab/blackdatura"
)
```

See the [example](example/main.go).


## Thanks
- [zap][1]
- [lumberjack][2]
- [Gin][3]
- [ginzap][4]
- [GORM][5]
- [Go Redis][6]
- [sarama][7]


[1]:https://github.com/uber-go/zap
[2]:https://github.com/natefinch/lumberjack
[3]:https://github.com/gin-gonic/gin
[4]:https://github.com/gin-contrib/zap
[5]:https://github.com/go-gorm/gorm
[6]:https://github.com/go-redis/redis
[7]:https://github.com/Shopify/sarama

