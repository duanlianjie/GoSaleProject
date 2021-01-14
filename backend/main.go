package main

import (
	ctx "context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"goproject/backend/web/controllers"
	"goproject/commons"
	"goproject/repositories"
	"goproject/services"
	"log"
)

func main() {
	app := iris.New()                                      // 创建 iris 实例
	app.Logger().SetLevel("debug")                         // 设置错误模式，在mvc模式下提示错误
	template := iris.HTML("./backend/web/views", ".html"). // 注册模板
								Layout("shared/layout.html").
								Reload(true)
	app.RegisterView(template)

	//app.StaticWeb("/assets", "./backend/web/assets")
	app.HandleDir("/assets", "./backend/web/assets") // 设置模板目标
	app.OnAnyErrorCode(func(context iris.Context) {  // 出现异常跳转到指定页面
		context.ViewData("message", context.Values().GetStringDefault("message", "访问的页面出错！"))
		context.ViewLayout("")
		context.View("shared/error.html")
	})

	// 连接数据库
	db, err := commons.NewMysqlConn()
	if err != nil {
		log.Println(err)
	}
	context, cancel := ctx.WithCancel(ctx.Background())
	defer cancel()

	// 注册控制器
	productRepository := repositories.NewProductRepository("product", db)
	productService := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(context, productService)
	product.Handle(new(controllers.ProductController))

	orderRepository := repositories.NewOrderRepository("order1", db)
	orderService := services.NewOrderService(orderRepository)
	orderParty := app.Party("/order")
	order := mvc.New(orderParty)
	order.Register(context, orderService)
	order.Handle(new(controllers.OrderController))

	// 6.启动服务
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
