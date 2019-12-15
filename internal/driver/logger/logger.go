package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Writer specifies output of logger.
var writer zapcore.WriteSyncer = os.Stdout

func New() *zap.Logger {
	zapCfg := zap.NewProductionConfig()
	encCfg := zapCfg.EncoderConfig
	encCfg.EncodeTime = jstTimeEncoder

	return zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encCfg),
			zapcore.Lock(writer),
			zap.NewAtomicLevel(),
		),
	)
}

func jstTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	const layout = "2006-01-02 15:04:05"
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	enc.AppendString(t.In(jst).Format(layout))
}
