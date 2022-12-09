package main

import (
	memModel "lean-oauth/internal/repositories/members/mysql"
	regisCatModel "lean-oauth/internal/repositories/register_categories/mysql"
	"lean-oauth/pkg"
)

func main() {

	var conn, err = pkg.NewConnect("root", "root", "oauth", "0.0.0.0", "3306").Connect()
	if err != nil {
		panic(err)
	}
	pkg.AutoMigrate(conn, memModel.MembersModel{})
	pkg.AutoMigrate(conn, regisCatModel.RegisterCategoriesModel{})
}
