package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"goproject/commons"
	"goproject/datamodels"
	"goproject/services"
	"strconv"
)

type ProductController struct {
	Context        iris.Context
	ProductService services.ProductService
}

func (p *ProductController) GetAll() mvc.View {
	productArray, err := p.ProductService.GetAllProduct()
	if err != nil {
		p.Context.Application().Logger().Debug(err)
	}
	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			"productArray": productArray,
		},
	}
}

func (p *ProductController) GetAdd() mvc.View {
	return mvc.View{
		Name: "product/add.html",
	}
}

func (p *ProductController) GetManager() mvc.View {
	idString := p.Context.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 16)
	if err != nil {
		p.Context.Application().Logger().Debug(err)
	}
	product, err := p.ProductService.GetProductByID(id)
	if err != nil {
		p.Context.Application().Logger().Debug(err)
	}
	return mvc.View{
		Name: "product/manager.html",
		Data: iris.Map{
			"product": product,
		},
	}
}

func (p *ProductController) GetDelete() {
	idString := p.Context.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		p.Context.Application().Logger().Debug(err)
	}

	deleteOk := p.ProductService.DeleteProduct(id)

	if deleteOk {
		p.Context.Application().Logger().Debug("删除商品成功，ID：" + idString)
	} else {
		p.Context.Application().Logger().Debug("删除商品失败，ID：" + idString)
	}
	p.Context.Redirect("/product/all")
}

func (p *ProductController) PostAdd() {
	product := &datamodels.Product{}
	p.Context.Request().ParseForm()
	dec := commons.NewDecoder(&commons.DecoderOptions{TagName: "imooc"})
	if err := dec.Decode(p.Context.Request().Form, product); err != nil {
		p.Context.Application().Logger().Debug(err)
	}

	if _, err := p.ProductService.InsertProduct(product); err != nil {
		p.Context.Application().Logger().Debug(err)
	}
	p.Context.Redirect("/product/all")
}

func (p *ProductController) PostUpdate() {
	product := &datamodels.Product{}
	p.Context.Request().ParseForm()
	dec := commons.NewDecoder(&commons.DecoderOptions{TagName: "imooc"})
	if err := dec.Decode(p.Context.Request().Form, product); err != nil {
		p.Context.Application().Logger().Debug(err)
	}

	if err := p.ProductService.UpdateProduct(product); err != nil {
		p.Context.Application().Logger().Debug(err)
	}
	p.Context.Redirect("/product/all")
}
