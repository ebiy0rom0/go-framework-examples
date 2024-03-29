package middleware

import (
	"examples/pkg/infra/framework/http/web"
	"io"
	"net/http"
	"strings"

	"github.com/rs/zerolog"
)

var Logger *loggerMiddleware

type loggerMiddleware struct {
	logger zerolog.Logger
}

func InitLoggerMiddleware(w io.Writer) {
	Logger = &loggerMiddleware{logger: zerolog.New(w)}
}

func (m *loggerMiddleware) WithLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writer := web.NewResponseWriter(w, r)

		next.ServeHTTP(writer, r)

		if !strings.Contains(r.URL.Path, "swagger") {
			m.logger.Info().Object("accesslog", writer).Send()
		}
	})
}
