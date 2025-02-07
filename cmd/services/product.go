package services

import (
	"context"

	"github.com/verlinof/go-grpc/cmd/helpers"
	paginationPb "github.com/verlinof/go-grpc/pb/pagination"
	productPb "github.com/verlinof/go-grpc/pb/products"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ProductService struct {
	//Interface dari gRPC
	productPb.UnimplementedProductServiceServer
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{
		db: db,
	}
}

func (p *ProductService) GetProducts(ctx context.Context, pageParam *productPb.Page) (*productPb.Products, error) {
	var page int64
	var pagination paginationPb.Pagination
	var products []*productPb.Product

	// Get page from request
	if pageParam.GetPage() != 0 {
		page = pageParam.GetPage()
	}

	sql := p.db.Table("products as p").
		Joins("LEFT JOIN categories as c on c.id = p.category_id").
		Select("p.id, p.name, p.price, p.stock, c.id as category_id, c.name as category_name")

	offset, limit := helpers.Pagination(sql, page, &pagination)

	rows, err := sql.Offset(int(offset)).Limit(int(limit)).Rows()

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var product productPb.Product
		var category productPb.Category

		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &category.Id, &category.Name)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		product.Category = &category
		products = append(products, &product)
	}

	response := &productPb.Products{
		Pagination: &pagination,
		Data:       products,
	}

	return response, nil
}
