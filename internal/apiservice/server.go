package apiservice

import (
	"context"
	"fmt"
	"github.com/amit/twirp-to-rest/rpc/test"
)

type Server struct {}

func (s *Server) GetRpc(ctx context.Context, request *test.GetRpcRequest) (*test.GetRpcResponse, error) {
	responseMsg := fmt.Sprintf("get-rpc-called-with-request-id: %v", request.IdInRequest)
	return &test.GetRpcResponse{Response: responseMsg}, nil
}

func (s *Server) PostRpc(ctx context.Context, request *test.PostRpcRequest) (*test.PostRpcResponse, error) {
	responseMsg := fmt.Sprintf("post-rpc-called-with-request-data-1 :: %v :: request-data-2 :: %v", request.RequestData_1, request.RequestData_2)
	return &test.PostRpcResponse{Response: responseMsg}, nil
}