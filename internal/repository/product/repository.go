package product

import (
	"context"
	"log"
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


func (p *productRepository) GetProduct(Id int) model.Product{
	row := p.db.QueryRow(context.Background(), "SELECT p.product_name, p.product_price, c.category_id, c.category_name FROM products p JOIN categories c ON p.category_id = c.category_id WHERE p.product_id=$1;", Id)
	product := model.Product{}
	err := row.Scan(&product.Name, &product.Price, &product.Category.Id, &product.Category.Name)
	if err != nil{
		log.Println(err)
	}
	product.Attributes = p.getAttributes(Id)
	return product
}


func (p *productRepository) GetProducts() []model.Product{
	rows, err := p.db.Query(context.Background(),"SELECT p.product_id, p.product_name, p.product_price, c.category_id, c.category_name FROM products p JOIN categories c ON p.category_id = c.category_id;")
	if err != nil{
		log.Println(err)
	}
	users, err := pgx.CollectRows(rows, func (row pgx.CollectableRow) (model.Product, error){
		var product model.Product
		var id int
		err := row.Scan(&id, &product.Name, &product.Price, &product.Category.Id, &product.Category.Name)
		product.Attributes = p.getAttributes(id)
		return product, err
	})
	if err != nil{
		log.Println(err)
	}
	return users
}

func (p *productRepository) AddProduct(product model.Product) error{
	row:= p.db.QueryRow(context.Background(), "insert into products (product_name, product_price, category_id) values ($1, $2, $3) RETURNING product_id", product.Name, product.Price, product.Category.Id)
	var id int
	row.Scan(&id)
	log.Println(id)
	for _, attribute := range product.Attributes{
		_, err := p.db.Exec(context.Background(), "insert into attributes (attribute_name, attribute_value, product_id) values ($1, $2, $3)", attribute.Name, attribute.Value, id)
		if err != nil{
			log.Println(err)
		}
	} 
	return nil
}

func (p *productRepository) getAttributes(Id int) []model.Attribute{
	rows, err := p.db.Query(context.Background(), "select attribute_name, attribute_value from attributes where product_id=$1", Id)
	if err != nil{
		log.Println(err)
	}
	attributes, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Attribute])
	if err != nil{
		log.Println(err)
	}
	return attributes
}

func (p *productRepository) AddAttribute(attribute model.Attribute, product_id int) error{
	_, err := p.db.Exec(context.Background(), "insert into attributes (attribute_name, attribute_value, product_id) values ($1, $2, $3)", attribute.Name, attribute.Value, product_id)
	return err
}


func (p *productRepository) AddCategory(category model.Category) error{
	_, err := p.db.Exec(context.Background(), "insert into categories (category_name) values($1)", category.Name)
	return err
}