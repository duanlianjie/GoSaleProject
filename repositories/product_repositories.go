package repositories

import (
	"database/sql"
	"goproject/commons"
	"goproject/datamodels"
	"strconv"
)

// 1.定义对应的接口
// 2.实现定义的接口

type ProductRepository interface {
	Conn() error
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
	SubProductNum(productID int64) error
}

type ProductRepositoryManager struct {
	table     string
	mysqlConn *sql.DB
}

func NewProductRepository(table string, db *sql.DB) ProductRepository {
	return &ProductRepositoryManager{table: table, mysqlConn: db}
}

func (p *ProductRepositoryManager) Conn() (err error) {
	//panic("implement me")
	if p.mysqlConn == nil {
		mysqlConn, err := commons.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = mysqlConn
	}
	if p.table == "" {
		p.table = "product"
	}
	return err
}

// Insert 插入
func (p *ProductRepositoryManager) Insert(product *datamodels.Product) (productId int64, err error) {
	//panic("implement me")
	//1.判断连接是否成功
	if err = p.Conn(); err != nil {
		return
	}

	//2.准备sql
	sql := "INSERT product SET productName=?,productNum=?,productImage=?,productUrl=?"
	stmt, errSql := p.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if errSql != nil {
		return 0, errSql
	}

	//3.传入参数
	result, errStmt := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if errStmt != nil {
		return 0, errStmt
	}
	return result.LastInsertId()
}

// Delete 删除
func (p *ProductRepositoryManager) Delete(productId int64) bool {
	//panic("implement me")
	if err := p.Conn(); err != nil {
		return false
	}

	sql := "delete from product where id=?"
	stmt, err := p.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return false
	}

	_, err = stmt.Exec(productId)
	if err != nil {
		return false
	}
	return true
}

// Update 更新
func (p *ProductRepositoryManager) Update(product *datamodels.Product) (err error) {
	//panic("implement me")
	if err = p.Conn(); err != nil {
		return
	}

	sql := "update product set productName=?,productNum=?,productImage=?,productUrl=? where id=" + strconv.FormatInt(product.ID, 10)
	stmt, err := p.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}

	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return err
	}
	return nil
}

// SelectByKey 查询一条记录
func (p *ProductRepositoryManager) SelectByKey(productId int64) (product *datamodels.Product, err error) {
	//panic("implement me")
	if err = p.Conn(); err != nil {
		return &datamodels.Product{}, err
	}

	sql := "select * from " + p.table + " where id=" + strconv.FormatInt(productId, 10)
	row, err := p.mysqlConn.Query(sql)
	defer row.Close()
	if err != nil {
		return &datamodels.Product{}, err
	}

	result := commons.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.Product{}, nil
	}
	product = &datamodels.Product{}
	commons.DataToStructByTagSql(result, product)
	return
}

// SelectAll 查询所有记录
func (p *ProductRepositoryManager) SelectAll() (products []*datamodels.Product, err error) {
	//panic("implement me")
	if err = p.Conn(); err != nil {
		return nil, err
	}

	sql := "select * from " + p.table
	rows, err := p.mysqlConn.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	results := commons.GetResultRows(rows)
	if len(results) == 0 {
		return nil, err
	}
	for _, v := range results {
		product := &datamodels.Product{}
		commons.DataToStructByTagSql(v, product)
		products = append(products, product)
	}
	return
}

func (p *ProductRepositoryManager) SubProductNum(productID int64) error {
	if err := p.Conn(); err != nil {
		return err
	}
	sql := "update " + p.table + " set " + " productNum=productNum-1 where ID=" + strconv.FormatInt(productID, 10)
	stmt, err := p.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	return err
}
