package api

import (
	"context"
	"microtask/productservice/data"
	"microtask/productservice/service1pb"
)

type Server struct {
	service1pb.UnimplementedService1Server
}

func (s *Server) ProductList(ctx context.Context, req *service1pb.NoParams) (*service1pb.ProdulistResponse, error) {
	prod, err := data.GetProduct()
	if err != nil {
		return nil, err
	}
	return &service1pb.ProdulistResponse{
		Id:          int32(prod.Id),
		Name:        prod.Name,
		Description: prod.Description,
		Price:       float64(prod.Price),
	}, nil
}
