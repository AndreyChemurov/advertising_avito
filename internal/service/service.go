package service

import (
	"advertising_avito/internal/database"
	"advertising_avito/internal/types"
	"context"
	"fmt"
)

func Create(ctx context.Context, r types.CreateRequest) (*types.CreateResponse, error) {
	database, err := database.GetDatabase("postgres")
	if err != nil {
		return nil, err
	}

	id, err := database.Create(ctx, r.Name, r.Description, r.Links, r.Price)
	if err != nil {
		return nil, err
	}

	response := new(types.CreateResponse)
	response.ID = id

	return response, nil
}

func GetOne(ctx context.Context, r types.GetOneRequest) (*types.GetOneResponse, error) {
	database, err := database.GetDatabase("postgres")
	if err != nil {
		return nil, err
	}

	name, price, mainLink, desc, allLinks, err := database.GetOne(ctx, r.ID, r.Fields)
	if err != nil {
		return nil, err
	}

	if name == "" {
		return nil, fmt.Errorf("advertisement with id %d does not exist", r.ID)
	}

	response := new(types.GetOneResponse)
	response.Name = name
	response.Price = price
	response.MainLink = mainLink

	if r.Fields {
		response.Description = desc
		response.AllLinks = allLinks
	}

	return response, nil
}

func GetAll(ctx context.Context, r types.GetAllRequest) (*types.GetAllResponse, error) {
	_, err := database.GetDatabase("postgres")
	if err != nil {
		return nil, err
	}

	// name, mainLink, price := database.GetAll(ctx, r.Page, r.Sort)

	return nil, nil
}
