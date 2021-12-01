//go:build bd_all || bd_fiber || fiber

package blackdatura

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type FiberLoggerConfig struct {
	Next          func(c *fiber.Ctx) bool
	Level         string
	OutputHeader  []string
	OutputLocals  []string
	OutputCookies []string
	OutputBody    bool
	OutputResp    bool
}

func FiberLogger(log *zap.Logger, cfg FiberLoggerConfig) fiber.Handler {
	var (
		pid = strconv.Itoa(os.Getpid())
		l   = log.With(zap.String("parent", "Fiber"))
		zl  = zapLevel(cfg.Level)
	)

	return func(c *fiber.Ctx) (err error) {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		startTime := time.Now()

		// Handle request, store err for logging
		chainErr := c.Next()

		// Manually call error handler
		if chainErr != nil {
			if err = c.App().ErrorHandler(c, chainErr); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		stopTime := time.Now()

		arr := make([]zap.Field, 0, 20+len(cfg.OutputHeader)+len(cfg.OutputCookies)+len(cfg.OutputLocals))

		arr = append(arr,
			zap.Time("Time", startTime),
			zap.String("Referer", c.Get(fiber.HeaderReferer)),
			zap.String("Protocol", c.Protocol()),
			zap.String("Method", c.Method()),
			zap.String("PID", pid),
			zap.String("Port", c.Port()),
			zap.String("IP", c.IP()),
			zap.String("IPs", c.Get(fiber.HeaderXForwardedFor)),
			zap.String("Host", c.Hostname()),
			zap.String("Path", c.Path()),
			zap.String("URL", c.OriginalURL()),
			zap.String("UserAgent", c.Get(fiber.HeaderUserAgent)),
			zap.Duration("Latency", stopTime.Sub(startTime)),
			zap.String("QueryStringParams", c.Request().URI().QueryArgs().String()),
			zap.String("Route", c.Route().Path),
			zap.Int("BytesReceived", len(c.Request().Body())),
			zap.Int("BytesSent", len(c.Response().Body())),
			zap.Int("Status", c.Response().StatusCode()),
			zap.ByteString("Body", c.Body()),
			zap.ByteString("ResBody", c.Response().Body()),
		)

		if cfg.OutputResp {
			arr = append(arr, zap.ByteString("ResBody", c.Response().Body()))
		}
		if cfg.OutputBody {
			arr = append(arr, zap.ByteString("Body", c.Body()))
		}

		for _, v := range cfg.OutputHeader {
			arr = append(arr, zap.String("Header: "+v, c.Get(v)))
		}
		for _, v := range cfg.OutputCookies {
			arr = append(arr, zap.String("Cookie: "+v, c.Cookies(v)))
		}
		for _, v := range cfg.OutputLocals {
			switch val := c.Locals(v).(type) {
			case []byte:
				arr = append(arr, zap.ByteString("Locals: "+v, val))
			case string:
				arr = append(arr, zap.String("Header: "+v, val))
			case nil:
			default:
				arr = append(arr, zap.String("Header: "+v, fmt.Sprintf("%v", val)))
			}
		}

		switch zl {
		case zapcore.DebugLevel:
			l.Debug(c.Path(), arr...)
		case zapcore.InfoLevel:
			l.Info(c.Path(), arr...)
		case zapcore.WarnLevel:
			l.Warn(c.Path(), arr...)
		case zapcore.ErrorLevel:
			l.Error(c.Path(), arr...)
		case zapcore.FatalLevel:
			l.Fatal(c.Path(), arr...)
		}

		return nil
	}
}
