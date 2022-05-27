package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rcrespodev/Blogs/design/repository/pkg/domain"
	"github.com/rcrespodev/Blogs/design/repository/pkg/server/globalObjects"
)

func HttpGetBitcoinPriceGinHandlerFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(HttpGetBitcoinPrice())
	}
}

func HttpGetBitcoinPrice() (httpStatusCode int, responseData interface{}) {
	bitcoinSrv := domain.NewBitcoinSrv(globalObjects.Factory.Repository())
	return 200, bitcoinSrv.GetBitcoinPrice()
}
