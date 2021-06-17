package service

import (
	"advertising_avito/internal/database"
	"advertising_avito/internal/types"
	"context"
)

func Create(ctx context.Context, r types.CreateRequest) (types.CreateResponse, error) {
	database, err := database.GetDatabase("postgres")
	if err != nil {
		return types.CreateResponse{}, err
	}

	id, err := database.Create(r.Name, r.Description, r.Links, r.Price)
	if err != nil {
		return types.CreateResponse{}, err
	}

	return types.CreateResponse{ID: id}, nil
}

func GetAll(ctx context.Context, r types.GetAllRequest) (types.GetAllResponse, error) {
	_, err := database.GetDatabase("postgres")
	if err != nil {
		return types.GetAllResponse{}, err
	}

	return types.GetAllResponse{}, nil
}

func GetOne(ctx context.Context, r types.GetOneRequest) (types.GetOneResponse, error) {
	_, err := database.GetDatabase("postgres")
	if err != nil {
		return types.GetOneResponse{}, err
	}

	return types.GetOneResponse{}, nil
}
