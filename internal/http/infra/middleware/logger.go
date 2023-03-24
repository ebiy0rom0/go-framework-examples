package middleware

import (
	"examples/internal/http/infra/log"
	"examples/internal/http/infra/web"
	"examples/internal/http/registry"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog"
)

type loggerMiddleware struct {
	logger zerolog.Logger
}

func NewLoggerMiddleware(w io.Writer) *loggerMiddleware {
	return &loggerMiddleware{logger: zerolog.New(w)}
}

func (m *loggerMiddleware) WithLogger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writer := web.NewResponseWriter(w, r)

		ctx := log.NewLogContext(r.Context())
		r = r.WithContext(ctx)

		next.ServeHTTP(writer, r)

		if err := registry.Logger.Send(ctx); err != nil {
			fmt.Println(err)
		}
		m.logger.Info().Object("accesslog", writer).Send()
	}
}
