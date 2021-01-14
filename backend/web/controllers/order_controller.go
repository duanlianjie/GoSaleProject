package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"goproject/services"
)

type OrderController struct {
	Context      iris.Context
	OrderService services.OrderService
}

func (o *OrderController) Get() mvc.View {
	orderArray, err := o.OrderService.GetAllOrderInfo()
	if err != nil {
		o.Context.Application().Logger().Debug(err)
	}
	return mvc.View{
		Name: "order/view.html",
		Data: iris.Map{
			"order": orderArray,
		},
	}
}
