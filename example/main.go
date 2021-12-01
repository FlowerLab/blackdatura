//go:build bd_all

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
	"go.uber.org/zap"
	"go.x2ox.com/blackdatura"
)

// go build -tags "ba_all" go.x2ox.com/blackdatura/example

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

	app := fiber.New()

	app.Use(
		requestid.New(),
		blackdatura.FiberLogger(i, blackdatura.FiberLoggerConfig{
			Next:          nil,
			Level:         "debug",
			OutputHeader:  []string{"Cookie"},
			OutputCookies: []string{"grafana_session"},
			OutputLocals:  []string{"a"},
			OutputBody:    true,
			OutputResp:    true,
		}),
		requestid.New(requestid.Config{
			Header: "X-Request-ID",
			Generator: func() string {
				return utils.UUID()
			},
		}),
	)

	app.Get("/", func(c *fiber.Ctx) error {
		c.Locals("AppID", 2)
		return nil
	})

	if err := app.Listen(":13079"); err != nil {
		i.Fatal("listen error", zap.Error(err))
	}
}
