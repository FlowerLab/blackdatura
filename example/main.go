package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.x2ox.com/blackdatura"
)

// go build -tags "gin" go.x2ox.com/blackdatura/example

func main() {
	blackdatura.Init("debug", true, blackdatura.DefaultLumberjack())

	i := blackdatura.New()
	i.Info("Black Datura")

	j := blackdatura.With("flowers meaning")
	j.Debug("black datura",
		zap.Any("Unpredictable darkness", "不可预知的黑暗"),
		zap.Any("love of death and turbulence", "死亡和颠沛流离的爱"),
		zap.Any("loveless and grudgless in the world", "凡间的无爱与无仇"),
		zap.Any("desperate love", "绝望的爱"),
		zap.Any("unpredictable death and love", "不可预知的死亡和爱"),
		zap.Any("hurt but strong mind", "被伤害的坚韧创痍的心灵"),
		zap.Any("unreturnable path and road", "生的不归之路"),
	)

	r := gin.New()

	r.Use(blackdatura.Ginzap(blackdatura.With("gin zap")),
		blackdatura.RecoveryWithZap(blackdatura.With("recovery with zap")))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/panic", func(c *gin.Context) {
		panic("An unexpected error happen!")
	})

	_ = r.Run(":13079")
}
