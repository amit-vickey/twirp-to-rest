package main

import (
	"context"
	"fmt"
	"github.com/amit/twirp-to-rest/internal/apiservice"
	"github.com/amit/twirp-to-rest/internal/health"
	"github.com/amit/twirp-to-rest/rpc/test"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/twitchtv/twirp"
	"net"
	"net/http"
)

const (
	authUserCtxKey = "authUser"
	authPassCtxKey = "authPass"
	requestIdHeader = "X-Task-ID"
)

func main() {

	router := gin.New()
	router.Use(gin.Recovery())
	authorized := router.Group("/")
	authorized.Use(AuthMiddleware)

	router.Handle("GET", "/getRpc", apiservice.GetRpcHandler)
	router.Handle("GET", "/postRpc", apiservice.PostRpcHandler)
	router.Handle("GET", "/status", apiservice.StatusHandler)

	server := &apiservice.Server{}
	authServerHook := AuthServerHook()

	twirpHandler := test.NewTestServer(server, authServerHook)
	httpHandler := WithBasicHeaders(twirpHandler)
	statusHandlerServer := health.NewStatusHandlerServer()
	statusHandler := HandleStatus(statusHandlerServer)

	mux := http.NewServeMux()
	mux.Handle(twirpHandler.PathPrefix(), httpHandler)
	mux.Handle("/status", statusHandler)

	listener, err := net.Listen("tcp4", ":8080")
	if err != nil {
		panic(fmt.Sprintf("failed-to-listen: %v", err))
	}
	httpServer := http.Server{
		Handler: router,
	}

	if err := httpServer.Serve(listener); err != nil {
		panic(fmt.Sprintf("failed-to-start-http-server: %v", err))
	}
}

func AuthServerHook() *twirp.ServerHooks {
	hooks := &twirp.ServerHooks{}
	hooks.RequestRouted = func(ctx context.Context) (context.Context, error) {
		requestUser := ctx.Value(authUserCtxKey).(string)
		requestPassword := ctx.Value(authPassCtxKey).(string)
		if requestUser != "admin" && requestPassword != "admin" {
			return ctx, errors.New("unauthorised")
		}
		return ctx, nil
	}
	return hooks
}

func AuthMiddleware(ctx *gin.Context) {
	ctx.Next()
}

func WithBasicHeaders(base http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		authUser, authPass, _ := request.BasicAuth()
		ctx = context.WithValue(ctx, authUserCtxKey, authUser)
		ctx = context.WithValue(ctx, authPassCtxKey, authPass)
		ctx = context.WithValue(ctx, "taskId", request.Header.Get(requestIdHeader))
		request = request.WithContext(ctx)

		base.ServeHTTP(responseWriter, request)
	})
}

func HandleStatus(base http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		base.ServeHTTP(responseWriter, request)
	})
}
