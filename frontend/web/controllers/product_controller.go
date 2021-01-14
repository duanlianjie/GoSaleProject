package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"goproject/datamodels"
	"goproject/services"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
)

type ProductController struct {
	Context        iris.Context
	ProductService services.ProductService
	OrderService   services.OrderService
	Session        *sessions.Session
}

var (
	//生成的Html保存目录
	htmlOutPath = "./frontend/web/htmlProductShow/"
	//静态文件模版目录
	templatePath = "./frontend/web/views/template/"
)

//生成html静态文件
func generateStaticHtml(context iris.Context, template *template.Template, fileName string, product *datamodels.Product) {
	//1.判断静态文件是否存在
	if exist(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			context.Application().Logger().Error(err)
		}
	}
	//2.生成静态文件
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		context.Application().Logger().Error(err)
	}
	defer file.Close()
	template.Execute(file, &product)
}

//判断文件是否存在
func exist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

func (p *ProductController) GetGenerateHtml() {
	productString := p.Context.URLParam("productID")
	productID, err := strconv.Atoi(productString)
	if err != nil {
		p.Context.Application().Logger().Debug(err)
	}

	//1.获取模版
	contentTemplate, err := template.ParseFiles(filepath.Join(templatePath, "product.html"))
	if err != nil {
		p.Context.Application().Logger().Debug(err)
	}
	//2.获取html生成路径
	fileName := filepath.Join(htmlOutPath, "htmlProduct.html")
	//3.获取模版渲染数据
	product, err := p.ProductService.GetProductByID(int64(productID))
	if err != nil {
		p.Context.Application().Logger().Debug(err)
	}
	//4.生成静态文件
	generateStaticHtml(p.Context, contentTemplate, fileName, product)
}

func (p *ProductController) GetDetail() mvc.View {
	product, err := p.ProductService.GetProductByID(1)
	if err != nil {
		p.Context.Application().Logger().Error(err)
	}

	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/view.html",
		Data: iris.Map{
			"product": product,
		},
	}
}

func (p *ProductController) GetOrder() mvc.View {
	productString := p.Context.URLParam("productID")
	userString := p.Context.GetCookie("uid")

	productID, err := strconv.Atoi(productString)
	if err != nil {
		p.Context.Application().Logger().Debug(err)
	}

	product, err := p.ProductService.GetProductByID(int64(productID))
	if err != nil {
		p.Context.Application().Logger().Debug(err)
	}

	var orderID int64
	showMessage := "抢购失败"
	if product.ProductNum > 0 {
		product.ProductNum -= 1
		err := p.ProductService.UpdateProduct(product)
		if err != nil {
			p.Context.Application().Logger().Debug(err)
		}

		userID, err := strconv.Atoi(userString)
		if err != nil {
			p.Context.Application().Logger().Debug(err)
		}

		order := &datamodels.Order{
			UserID:      int64(userID),
			ProductID:   int64(productID),
			OrderStatus: datamodels.OrderSuccess,
		}
		orderID, err = p.OrderService.InsertOrder(order)
		if err != nil {
			p.Context.Application().Logger().Debug(err)
		} else {
			showMessage = "抢购成功"
		}
	}
	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/result.html",
		Data: iris.Map{
			"orderID":     orderID,
			"showMessage": showMessage,
		},
	}
}
