package service

import (
	"Week02/dao"
	"github.com/pkg/errors"
)

type Service struct {
	dao *dao.Dao
}
func NewService() *Service{
	return &Service{dao.NewDao()}
}

func (s *Service) GetOrderNameByOrderId(id int) (order *dao.Order,err error) {
	s=NewService()
	order,err=s.dao.GetOrderById(id)
	return order,errors.Wrapf(err,"service: order %d cannot get",id)
}


