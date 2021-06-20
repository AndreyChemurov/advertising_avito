package service

import (
	"advertising_avito/internal/database"
	"advertising_avito/internal/types"
	"context"

	"errors"
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

	name, price, desc, allLinks, err := database.GetOne(ctx, r.ID, r.Fields)
	if err != nil {
		return nil, err
	}

	// Так как БД не заботится о том, есть запись с таким ID или нет,
	// то она просто возвращает пустые значения, если записи по ID не существует.
	if name == "" {
		return nil, fmt.Errorf("advertisement with id %d does not exist", r.ID)
	}

	response := new(types.GetOneResponse)
	response.Name = name
	response.Price = price
	response.MainLink = allLinks[0]

	if r.Fields {
		response.Description = desc
		response.AllLinks = allLinks
	}

	return response, nil
}

func GetAll(ctx context.Context, r types.GetAllForService) (*types.GetAllResponse, error) {
	database, err := database.GetDatabase("postgres")
	if err != nil {
		return nil, err
	}

	advertisements, err := database.GetAll(ctx, r.Page, r.Sort)
	if err != nil {
		return nil, err
	}

	// Исключение ситуации, при которой можно запросить информацию об объявлениях,
	// не создав при этом ни одного
	if len(advertisements) == 0 {
		return nil, errors.New("no advertisements created")
	}

	response := new(types.GetAllResponse)
	response.Advertisements = advertisements

	return response, nil
}
