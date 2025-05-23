package app

import (
	"CaliYa/core/domain/models"
	"CaliYa/core/domain/ports"
	"context"
)

type productsApp struct {
	repo ports.ProductsRepo
}

func NewProductsApp(repo ports.ProductsRepo) ports.ProductsApp {
	return &productsApp{repo}
}

func (p *productsApp) RegisterProduct(ctx context.Context) error {

	// if err := p.repo.RegisterProducts(ctx); err != nil {
	// 	return err
	// }

	return nil
}

func (p *productsApp) GetProducts(ctx context.Context) (*models.ProductsShops, error) {

	products, err := p.repo.GetProducts(ctx)
	if err != nil {
		return products, err
	}

	return products, nil
}
