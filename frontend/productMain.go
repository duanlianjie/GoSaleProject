package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New() // 创建 iris 实例

	//app.StaticWeb("/assets", "./backend/web/assets")
	app.HandleDir("/public", "./frontend/web/public")        // 设置模板目标
	app.HandleDir("/html", "./frontend/web/htmlProductShow") // 访问生成好的 html 静态文件

	app.Run(
		iris.Addr("0.0.0.0:80"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
