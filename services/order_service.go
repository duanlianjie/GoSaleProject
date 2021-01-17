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
	InsertOrderByMessage(message *datamodels.Message) (orderID int64, err error)
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

//根据消息创建订单
func (o *OrderServiceManager) InsertOrderByMessage(message *datamodels.Message) (orderID int64, err error) {
	order := &datamodels.Order{
		UserID:      message.UserID,
		ProductID:   message.ProductID,
		OrderStatus: datamodels.OrderSuccess,
	}
	return o.InsertOrder(order)
}

func NewOrderService(repository repositories.OrderRepository) OrderService {
	return &OrderServiceManager{orderRepository: repository}
}
