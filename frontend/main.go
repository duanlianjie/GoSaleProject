package main

import (
	ctx "context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"goproject/commons"
	"goproject/frontend/middleware"
	"goproject/frontend/web/controllers"
	"goproject/repositories"
	"goproject/services"
	"log"
)

func main() {
	app := iris.New()                                       // 创建 iris 实例
	app.Logger().SetLevel("debug")                          // 设置错误模式，在mvc模式下提示错误
	template := iris.HTML("./frontend/web/views", ".html"). // 注册模板
								Layout("shared/layout.html").
								Reload(true)
	app.RegisterView(template)

	//app.StaticWeb("/assets", "./backend/web/assets")
	app.HandleDir("/public", "./frontend/web/public")        // 设置模板目标
	app.HandleDir("/html", "./frontend/web/htmlProductShow") // 访问生成好的 html 静态文件
	app.OnAnyErrorCode(func(context iris.Context) {          // 出现异常跳转到指定页面
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

	//session := sessions.New(sessions.Config{
	//	Cookie:  "helloworld",
	//	Expires: 60 * time.Minute,
	//})

	// 注册用户控制器
	userRepository := repositories.NewUserRepository("user", db)
	userService := services.NewUserService(userRepository)

	userController := mvc.New(app.Party("/user"))
	//userController.Register(userService, context, session.Start)
	userController.Register(userService, context)
	userController.Handle(new(controllers.UserController))


	// 注册商品订单控制器
	productRepository := repositories.NewProductRepository("product", db)
	productService := services.NewProductService(productRepository)
	orderRepository := repositories.NewOrderRepository("order1", db)
	orderService := services.NewOrderService(orderRepository)

	productParty := app.Party("/product")
	productParty.Use(middleware.AuthConProduct)		// Cookie 验证中间件
	productController := mvc.New(productParty)
	//productController.Register(productService, orderService, session.Start)
	productController.Register(productService, orderService)
	productController.Handle(new(controllers.ProductController))

	app.Run(
		iris.Addr("localhost:8082"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
