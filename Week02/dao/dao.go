package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

type Dao struct {

}

func NewDao()  *Dao{
	return &Dao{}
}

type Order struct {
	Id  int
	Name string
}

var db *sql.DB

func (dao *Dao)GetOrderById(id int) (o *Order ,err error) {
	err=db.QueryRow("select GoodsName from orders where id = ?",id).Scan(&o.Name)
	err = sql.ErrNoRows
	return o,errors.Wrap(err,"Dao cannot get")

}



