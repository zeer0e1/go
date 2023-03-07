package database

import (
	"database/sql"

	"github.com/zeer0e1/go/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository{
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error{
 _,erro := r.Db.Exec("INSERT INTO ORDERS (id,price,tax,final_Price) VALUES (?,?,?,?)", order.ID,order.Price,order.Tax,order.FinalPrice)
 if erro != nil{
	return erro
 }
 return nil
}

func (r * OrderRepository) GetTotal() (int,error){
	var total int
	err := r.Db.QueryRow("select count(*) from orders").Scan(&total)
	if err != nil {
		return 0,err
	}
	return total,nil
}