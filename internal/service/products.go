package service

import (
	"context"
	"fmt"
	"time"

	"github.com/o-sokol-o/as/internal/transport/restclient"
	"github.com/o-sokol-o/as/pkg/domain"
)

type Repository interface {
	InsertLog(ctx context.Context, log domain.LogItem) error
	InsertProducts(ctx context.Context, products domain.Products) error
	UpdateProducts(ctx context.Context, products domain.Products) error
	GetProducts(ctx context.Context, paging domain.Paging, sorting domain.Sorting) (domain.Products, error)
}

type Services struct {
	repo Repository
}

func NewServices(repo Repository) *Services {
	return &Services{
		repo: repo,
	}
}

func (s *Services) InsertLog(ctx context.Context, log domain.LogItem) error {
	log.Time = time.Now().UTC()

	return s.repo.InsertLog(ctx, log)
}

func (s *Services) InsertProducts(ctx context.Context, products domain.Products) error {
	return s.repo.InsertProducts(ctx, products)
}

func (s *Services) UpdateProducts(ctx context.Context, products domain.Products) error {
	return s.repo.UpdateProducts(ctx, products)
}

func (s *Services) GetProducts(ctx context.Context) (domain.Products, error) {
	return s.repo.GetProducts(ctx, domain.Paging{}, domain.Sorting{})
}

func (s *Services) FetchExternProducts(ctx context.Context, url string) (int, int, error) {
	// Запрашиваем список продукции у внешнего сервиса по REST
	products, err := restclient.RestClientGet(url)
	if err != nil || len(products) == 0 {
		return 0, 0, err
	}
	fmt.Printf("Extern API products: %d\n", len(products))

	// Local store products
	pr, err := s.GetProducts(ctx)
	fmt.Printf("Local store products: %d\n", len(pr))
	storeProd := make(map[string]domain.Product)
	if len(pr) > 0 {
		for _, prd := range pr {
			storeProd[prd.Name] = prd
		}
	}

	// Product separation
	newProd := make([]domain.Product, 0)
	updateProd := make([]domain.Product, 0)
	for i := range products {
		prd, ok := storeProd[products[i].Name]
		if !ok {
			newProd = append(newProd, products[i])
		} else if prd.Cost != products[i].Cost {
			products[i].ChangeCount = prd.Cost + 1
			updateProd = append(updateProd, products[i])
		}
	}

	// Insert products
	if len(newProd) > 0 {
		fmt.Printf("Insert products: ")
		err = s.repo.InsertProducts(ctx, newProd)
		fmt.Printf("%d\n\n", len(newProd))
	}

	// Update products
	if len(updateProd) > 0 {
		fmt.Printf("Update products: ")
		err = s.repo.UpdateProducts(ctx, updateProd)
		fmt.Printf("%d\n\n", len(updateProd))
	}

	return len(newProd), len(updateProd), err
}

func (s *Services) ListExternProducts(ctx context.Context, paging domain.Paging, sorting domain.Sorting) (domain.Products, error) {
	return s.repo.GetProducts(ctx, paging, sorting)
}
