package main

import (
	"fmt"
	"goproject/commons"
	"goproject/rabbitmq"
	"goproject/repositories"
	"goproject/services"
)

func main() {
	db, err := commons.NewMysqlConn()
	if err != nil {
		fmt.Println(err)
	}
	//创建product数据库操作实例
	product := repositories.NewProductRepository("product", db)
	//创建product service
	productService := services.NewProductService(product)
	//创建Order数据库实例
	order := repositories.NewOrderRepository("order1", db)
	//创建order Service
	orderService := services.NewOrderService(order)

	rabbitmqConsumeSimple := rabbitmq.NewRabbitMQSimple("product")
	rabbitmqConsumeSimple.ConsumeSimple(orderService, productService)
}
