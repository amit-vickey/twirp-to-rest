package apiservice

import (
	"fmt"
	"github.com/amit/twirp-to-rest/rpc/test"
	"github.com/gin-gonic/gin"
)

func StatusHandler(ctx *gin.Context) {

}

func GetRpcHandler(ctx *gin.Context) {
	var requestInput test.GetRpcRequest
	err := ctx.BindJSON(&requestInput)
	if err != nil {
		fmt.Println(err)
	}
	response, err := (&Server{}).GetRpc(ctx, &requestInput)
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(200, gin.H{
		"response": response.Response,
	})
}

func PostRpcHandler(ctx *gin.Context) {
	var requestInput test.PostRpcRequest
	err := ctx.BindJSON(&requestInput)
	if err != nil {
		fmt.Println(err)
	}
	response, err := (&Server{}).PostRpc(ctx, &requestInput)
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(200, gin.H{
		"response": response.Response,
	})
}
