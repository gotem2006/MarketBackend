package product

import (
	"context"
	model "shop/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)



type productRepository struct{
	db *pgxpool.Pool
}



func NewProductRepository(db *pgxpool.Pool) *productRepository{
	return &productRepository{db}
}


func (p *productRepository) GetProduct(Id int) (product model.Product, err error) {
	query := `
		SELECT p.product_name, p.product_price, c.category_id, c.category_name 
		FROM products p 
		JOIN categories c ON p.category_id = c.category_id 
		WHERE p.product_id=$1;
	`
	err = p.db.QueryRow(context.Background(), query, Id).Scan(
		&product.Name, &product.Price, &product.Category.Id, &product.Category.Name,
	)
	if err != nil {
		return product, err
	}
	product.Attributes, err = p.getAttributes(Id)
	if err != nil{
		return product, err
	}
	return product, nil
}

func (p *productRepository) GetProducts() (Products []model.Product, err error){
	query := `
		SELECT p.product_id, p.product_name, p.product_price, c.category_id, c.category_name 
		FROM products p 
		JOIN categories c 
		ON p.category_id = c.category_id;
	`
	rows, err := p.db.Query(context.Background(), query)
	if err != nil{
		return nil, err
	}
	Products, err = pgx.CollectRows(rows, func (row pgx.CollectableRow) (model.Product, error){
		var product model.Product
		var id int
		err := row.Scan(&id, &product.Name, &product.Price, &product.Category.Id, &product.Category.Name)
		if err != nil{
			return product, err
		}
		product.Attributes, err = p.getAttributes(id)
		if err != nil{
			return product, err
		}
		return product, nil
	})
	if err != nil{
		return nil, err
	}
	return Products, err
}

func (p *productRepository) AddProduct(product model.Product) error{
	query := `
		INSERT INTO products (product_name, product_price, category_id) 
		VALUES ($1, $2, $3) RETURNING 
		product_id
	`
	var id int
	p.db.QueryRow(context.Background(), query, product.Name, product.Price, product.Category.Id).Scan(&id)
	for _, attribute := range product.Attributes{
		query = `
			INSERT INTO attributes (attribute_name, attribute_value, product_id) VALUES ($1, $2, $3)
		`
		_, err := p.db.Exec(context.Background(), query, attribute.Name, attribute.Value, id)
		if err != nil{
			return err
		}
	}
	return nil
}

func (p *productRepository) getAttributes(Id int) (Attributes []model.Attribute, err error){
	rows, err := p.db.Query(context.Background(), "select attribute_name, attribute_value from attributes where product_id=$1", Id)
	if err != nil{
		return nil, err
	}
	Attributes, err = pgx.CollectRows(rows, pgx.RowToStructByName[model.Attribute])
	if err != nil{
		return nil, err
	}
	return Attributes, nil
}

func (p *productRepository) AddAttribute(attribute model.Attribute, product_id int) error{
	query := `
		INSERT INTO attributes (attribute_name, attribute_value, product_id) VALUES ($1, $2, $3)
	`
	_, err := p.db.Exec(context.Background(), query, attribute.Name, attribute.Value, product_id)
	return err
}


func (p *productRepository) AddCategory(category model.Category) error{
	query := `
		INSERT INTO categories (category_name) VALUES ($1)
	`
	_, err := p.db.Exec(context.Background(), query, category.Name)
	return err
}
