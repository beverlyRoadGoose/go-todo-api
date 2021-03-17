package transport

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"todo-api/internal"
)

type Endpoints struct {
	GetItemEndpoint    endpoint.Endpoint
	CreateItemEndpoint endpoint.Endpoint
	UpdateItemEndpoint endpoint.Endpoint
	DeleteItemEndpoint endpoint.Endpoint
}

func MakeGetItemEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, validReq := request.(GetItemRequest)
		if !validReq {
			return nil, errors.New("invalid request")
		}
		item, err := svc.GetItem(req.Id)
		if err != nil {
			return nil, err
		}
		return GetItemResponse{Item: *item}, nil
	}
}

func MakeCreateItemEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, validReq := request.(CreateItemRequest)
		if !validReq {
			return nil, errors.New("invalid request")
		}
		item, err := svc.CreateItem(req.Text)
		if err != nil {
			return nil, err
		}
		return CreateItemResponse{Item: *item}, nil
	}
}

func MakeUpdateItemEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, validReq := request.(UpdateItemRequest)
		if !validReq {
			return nil, errors.New("invalid request")
		}
		item, err := svc.UpdateItem(req.Id, req.Text, req.Done)
		if err != nil {
			return nil, err
		}
		return UpdateItemResponse{Item: *item}, nil
	}
}

func MakeDeleteItemEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, validReq := request.(DeleteItemRequest)
		if !validReq {
			return nil, errors.New("invalid request")
		}
		deleted, err := svc.DeleteItem(req.Id)
		if err != nil {
			return nil, err
		}
		return DeleteItemResponse{Deleted: deleted}, nil
	}
}

func MakeEndpoints(svc internal.Service) Endpoints {
	return Endpoints{
		GetItemEndpoint:    MakeGetItemEndpoint(svc),
		CreateItemEndpoint: MakeCreateItemEndpoint(svc),
		UpdateItemEndpoint: MakeUpdateItemEndpoint(svc),
		DeleteItemEndpoint: MakeDeleteItemEndpoint(svc),
	}
}
