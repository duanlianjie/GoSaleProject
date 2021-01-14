package services

import (
	"goproject/datamodels"
	"goproject/repositories"
)

type OrderService interface {
	InsertOrder(order *datamodels.Order) (int64, error)
	DeleteOrder(int642 int64) bool
	UpdateOrder(order *datamodels.Order) error
	GetOrderByID(int642 int64) (*datamodels.Order, error)
	GetAllOrder() ([]*datamodels.Order, error)
	GetAllOrderInfo() (map[int]map[string]string, error)
}

type OrderServiceManager struct {
	orderRepository repositories.OrderRepository
}

func (o *OrderServiceManager) InsertOrder(order *datamodels.Order) (int64, error) {
	//panic("implement me")
	return o.orderRepository.Insert(order)
}

func (o *OrderServiceManager) DeleteOrder(orderID int64) bool {
	//panic("implement me")
	return o.orderRepository.Delete(orderID)
}

func (o *OrderServiceManager) UpdateOrder(order *datamodels.Order) error {
	//panic("implement me")
	return o.orderRepository.Update(order)
}

func (o *OrderServiceManager) GetOrderByID(orderID int64) (*datamodels.Order, error) {
	//panic("implement me")
	return o.orderRepository.SelectByKey(orderID)
}

func (o *OrderServiceManager) GetAllOrder() ([]*datamodels.Order, error) {
	//panic("implement me")
	return o.orderRepository.SelectAll()
}

func (o *OrderServiceManager) GetAllOrderInfo() (map[int]map[string]string, error) {
	//panic("implement me")
	return o.orderRepository.SelectAllWithInfo()
}

func NewOrderService(repository repositories.OrderRepository) OrderService {
	return &OrderServiceManager{orderRepository: repository}
}
