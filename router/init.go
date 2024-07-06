package router

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	L "github.com/zzzgydi/thanks/common/logger"
)

func InitHttpServer() {
	r := gin.New()
	r.Use(gin.Recovery())

	// register routers
	RootRouter(r)
	HealthRouter(r)

	logger := slog.NewLogLogger(L.Handler, slog.LevelError)

	viper.SetDefault("PORT", 14090)
	port := viper.GetInt64("PORT")

	srv := &http.Server{
		Addr:     ":" + strconv.FormatInt(port, 10),
		Handler:  r,
		ErrorLog: logger,
	}

	slog.Info(fmt.Sprintf("http server listen on http://0.0.0.0:%d", port))

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
