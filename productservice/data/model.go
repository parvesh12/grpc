package data

import "microtask/productservice/config"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float32
}

var DB = config.SetupDB()

func GetProduct() (prod Product, err error) {
	if err := DB.Table("tbl_products").Find(&prod).Error; err != nil {
		return prod, err
	}
	return prod, nil
}
