package entity

type OrderRepositoryInteface interface {
	Save(order *Order) error
	GetTotal() (int, error)
}