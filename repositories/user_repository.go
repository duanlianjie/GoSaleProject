package repositories

import (
	"database/sql"
	"errors"
	"goproject/commons"
	"goproject/datamodels"
	"strconv"
)

type UserRepository interface {
	Conn() (err error)
	Select(userName string) (user *datamodels.User, err error)
	Insert(user *datamodels.User) (userID int64, err error)
}

type UserRepositoryManager struct {
	table     string
	mysqlConn *sql.DB
}

func (u *UserRepositoryManager) Conn() (err error) {
	//panic("implement me")
	if u.mysqlConn == nil {
		mysql, err := commons.NewMysqlConn()
		if err != nil {
			return err
		}
		u.mysqlConn = mysql
	}
	if u.table == "" {
		u.table = "user"
	}
	return
}

func (u *UserRepositoryManager) Select(userName string) (user *datamodels.User, err error) {
	//panic("implement me")
	if userName == "" {
		return &datamodels.User{}, errors.New("用户名不能为空")
	}
	if err = u.Conn(); err != nil {
		return &datamodels.User{}, err
	}

	sql := "select * from " + u.table + " where userName=?"
	rows, err := u.mysqlConn.Query(sql, userName)
	defer rows.Close()
	if err != nil {
		return &datamodels.User{}, err
	}
	result := commons.GetResultRow(rows)
	if len(result) == 0 {
		return &datamodels.User{}, errors.New("用户不存在")
	}
	user = &datamodels.User{}
	commons.DataToStructByTagSql(result, user)
	return
}

func (u *UserRepositoryManager) Insert(user *datamodels.User) (userID int64, err error) {
	//panic("implement me")
	if err = u.Conn(); err != nil {
		return
	}

	sql := "insert " + u.table + " set nickName=?,userName=?,passWord=?"
	stmt, err := u.mysqlConn.Prepare(sql)
	if err != nil {
		return
	}
	result, err := stmt.Exec(user.NickName, user.UserName, user.HashPassword)
	if err != nil {
		return
	}
	return result.LastInsertId()
}

func (u *UserRepositoryManager) SelectByID(userID int64) (user *datamodels.User, err error) {
	if err = u.Conn(); err != nil {
		return
	}

	sql := "select * from " + u.table + " where ID=" + strconv.FormatInt(userID, 10)
	rows, err := u.mysqlConn.Query(sql)
	if err != nil {
		return &datamodels.User{}, err
	}

	result := commons.GetResultRow(rows)
	if len(result) == 0 {
		return &datamodels.User{}, errors.New("用户不存在")
	}

	user = &datamodels.User{}
	commons.DataToStructByTagSql(result, user)
	return
}

func NewUserRepository(table string, db *sql.DB) UserRepository {
	return &UserRepositoryManager{table: table, mysqlConn: db}
}
