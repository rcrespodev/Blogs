package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rcrespodev/Blogs/design/repository/api/v1/handlers"
	"net/http"
)

type Routes struct {
	Routes []route
}

type route struct {
	httpMethod   string
	relativePath string
	handler      gin.HandlerFunc
}

func (r route) RelativePath() string {
	return r.relativePath
}

func (r route) Handler() gin.HandlerFunc {
	return r.handler
}

func (r route) HttpMethod() string {
	return r.httpMethod
}

func NewRoutes() Routes {
	routes := []route{
		{
			httpMethod:   http.MethodGet,
			relativePath: "/bitcoin-price",
			handler:      handlers.HttpGetBitcoinPriceGinHandlerFunc(),
		},
	}
	return Routes{Routes: routes}
}
