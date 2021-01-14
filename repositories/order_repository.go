package repositories

import (
	"database/sql"
	"goproject/commons"
	"goproject/datamodels"
	"strconv"
)

type OrderRepository interface {
	Conn() error
	Insert(order *datamodels.Order) (int64, error)
	Delete(id int64) bool
	Update(order *datamodels.Order) error
	SelectByKey(id int64) (*datamodels.Order, error)
	SelectAll() ([]*datamodels.Order, error)
	SelectAllWithInfo() (map[int]map[string]string, error)
}

type OrderRepositoryManager struct {
	table     string
	mysqlConn *sql.DB
}

func (o *OrderRepositoryManager) Conn() error {
	//panic("implement me")
	if o.mysqlConn == nil {
		mysql, err := commons.NewMysqlConn()
		if err != nil {
			return err
		}
		o.mysqlConn = mysql
	}
	if o.table == "" {
		o.table = "order1"
	}
	return nil
}

func (o *OrderRepositoryManager) Insert(order *datamodels.Order) (productID int64, err error) {
	//panic("implement me")
	if err = o.Conn(); err != nil {
		return
	}

	sql := "INSERT " + o.table + " SET userID=?,productID=?,orderStatus=?"
	stmt, err := o.mysqlConn.Prepare(sql)
	if err != nil {
		return
	}
	result, err := stmt.Exec(order.UserID, order.ProductID, order.OrderStatus)
	if err != nil {
		return
	}
	return result.LastInsertId()
}

func (o *OrderRepositoryManager) Delete(orderID int64) (deleteOK bool) {
	//panic("implement me")
	if err := o.Conn(); err != nil {
		return
	}

	sql := "delete from " + o.table + " where ID=?"
	stmt, err := o.mysqlConn.Prepare(sql)
	if err != nil {
		return
	}
	_, err = stmt.Exec(orderID)
	if err != nil {
		return
	}
	return true
}

func (o *OrderRepositoryManager) Update(order *datamodels.Order) (err error) {
	//panic("implement me")
	if err = o.Conn(); err != nil {
		return
	}

	sql := "update " + o.table + "set userID=?,productID=?,orderStatus=? where id=" + strconv.FormatInt(order.ID, 10)
	stmt, err := o.mysqlConn.Prepare(sql)
	if err != nil {
		return
	}
	_, err = stmt.Exec(order.UserID, order.ProductID, order.OrderStatus)
	if err != nil {
		return
	}
	return nil
}

func (o *OrderRepositoryManager) SelectAll() (orderArray []*datamodels.Order, err error) {
	//panic("implement me")
	if err = o.Conn(); err != nil {
		return
	}

	sql := "select * from " + o.table
	rows, err := o.mysqlConn.Query(sql)
	if err != nil {
		return
	}

	result := commons.GetResultRows(rows)
	if len(result) == 0 {
		return
	}
	for _, v := range result {
		order := &datamodels.Order{}
		commons.DataToStructByTagSql(v, order)
		orderArray = append(orderArray, order)
	}
	return
}

func (o *OrderRepositoryManager) SelectAllWithInfo() (orderMap map[int]map[string]string, err error) {
	//panic("implement me")
	if err = o.Conn(); err != nil {
		return
	}
	sql := "select order1.id,productName,orderStatus from " + o.table + " left join product on order1.productID=product.id" // TODO
	rows, err := o.mysqlConn.Query(sql)
	if err != nil {
		return
	}
	return commons.GetResultRows(rows), err
}

func (o *OrderRepositoryManager) SelectByKey(orderID int64) (order *datamodels.Order, err error) {
	//panic("implement me")
	if err = o.Conn(); err != nil {
		return &datamodels.Order{}, err
	}

	sql := "select * from" + o.table + " where id=" + strconv.FormatInt(orderID, 10)
	row, err := o.mysqlConn.Query(sql)
	if err != nil {
		return &datamodels.Order{}, err
	}

	result := commons.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.Order{}, err
	}

	order = &datamodels.Order{}
	commons.DataToStructByTagSql(result, order)
	return order, err
}

func NewOrderRepository(table string, conn *sql.DB) OrderRepository {
	return &OrderRepositoryManager{
		table:     table,
		mysqlConn: conn,
	}
}
