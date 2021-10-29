package health

import (
	"context"
	"encoding/json"
	"github.com/twitchtv/twirp/ctxsetters"
	"net/http"
	"strconv"
)

type StatusHandler interface {
	Status (ctx context.Context, resp http.ResponseWriter, req *http.Request)
}

type StatusHandlerServer struct {
}

func (s *StatusHandlerServer) Status(ctx context.Context, resp http.ResponseWriter, req *http.Request){

	respContent := make(map[string]string)
	respContent["status"] = "ok"
	resp.WriteHeader(http.StatusOK)

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes, _ := json.Marshal(respContent)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))

	resp.Write(respBytes)
}

type BasicServer interface {
	http.Handler
}

func (s *StatusHandlerServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/status":
		s.Status(req.Context(), resp, req)
		return
	}
}

func NewStatusHandlerServer() BasicServer {
	return &StatusHandlerServer{}
}
