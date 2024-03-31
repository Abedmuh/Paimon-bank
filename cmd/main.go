package main

import (
	"net/http"
	"time"

	"github.com/Abedmuh/Paimon-bank/internal/routes"
	"github.com/Abedmuh/Paimon-bank/pkg/dbconfig"
	"github.com/Abedmuh/Paimon-bank/pkg/middleware"
	"github.com/Abedmuh/Paimon-bank/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var httpRequestProm = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "http_request_histogram",
	Help:    "Histogram of the http request duration.",
	Buckets: prometheus.LinearBuckets(1, 1, 10),
}, []string{"path", "method", "status"})

func main() {

	db, err := dbconfig.GetDBConnection()
	if err!= nil {
    panic(err)
  }
	defer db.Close()

	app := gin.Default()
	app.Use(middleware.RecoveryMiddleware())
	app.MaxMultipartMemory = 2 << 20 

	validate:= validator.New()
	validate.RegisterValidation("myUrl", utils.ValidateTransferProof)

	app.Use(GinPrometheusMiddleware())
	app.GET("/metrics", gin.WrapH(promhttp.Handler()))
	v1 := app.Group("v1")
	{
		routes.UserRoutes(v1, db, validate)
		routes.BalanceRoutes(v1, db, validate)
		routes.ImageRoutes(v1)
	}


	app.Run(":8080")
}

func GinPrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
			start := time.Now()
			c.Next() // Process request

			status := c.Writer.Status()
			httpRequestProm.WithLabelValues(c.Request.URL.Path, c.Request.Method, http.StatusText(status)).Observe(float64(time.Since(start).Milliseconds()))
	}
}