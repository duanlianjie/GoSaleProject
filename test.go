package main

import (
	"goproject/commons"
	"goproject/datamodels"
)

func main() {
	data := map[string]string{"id": "1", "productName": "11", "productNum": "111", "productImage": "1111", "productUrl": "11111"}
	product := &datamodels.Product{}
	commons.DataToStructByTagSql(data, product)
}
