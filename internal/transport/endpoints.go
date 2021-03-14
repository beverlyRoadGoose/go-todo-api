package transport

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"todo-api/internal"
)

type Endpoints struct {
	CreateItemEndpoint endpoint.Endpoint
	UpdateItemEndpoint endpoint.Endpoint
	DeleteItemEndpoint endpoint.Endpoint
}

func MakeCreateItemEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, validReq := request.(CreateItemRequest)
		if !validReq {
			return nil, errors.New("invalid request")
		}
		item := svc.CreateItem(req.Text)
		return CreateItemResponse{Item: item}, err
	}
}

func MakeUpdateItemEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, validReq := request.(UpdateItemRequest)
		if !validReq {
			return nil, errors.New("invalid request")
		}
		item := svc.UpdateItem(req.Text, req.Done)
		return UpdateItemResponse{Item: item}, err
	}
}

func MakeDeleteItemEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, validReq := request.(DeleteItemRequest)
		if !validReq {
			return nil, errors.New("invalid request")
		}
		svc.DeleteItem(req.Id)
		return DeleteItemResponse{}, err
	}
}

func MakeEndpoints(svc internal.Service) Endpoints {
	return Endpoints{
		CreateItemEndpoint: MakeCreateItemEndpoint(svc),
		UpdateItemEndpoint: MakeUpdateItemEndpoint(svc),
		DeleteItemEndpoint: MakeDeleteItemEndpoint(svc),
	}
}