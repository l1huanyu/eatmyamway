package database

import "github.com/l1huanyu/eatmyamway/model"

func QueryAmwayRand() (*model.Amway, error) {
	a := new(model.Amway)
	err := Conn().Where(a.ValidColumnName()+" = ?", true).First(a, "id >= ((SELECT MAX(id) FROM amways)-(SELECT MIN(id) FROM amways)) * RAND() + (SELECT MIN(id) FROM amways)").Error
	return a, err
}