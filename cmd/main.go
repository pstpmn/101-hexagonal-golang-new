package main

import (
	"lean-oauth/pkg"
	"log"
)

func main() {
	var conn, err = pkg.NewConnect("root", "root", "oauth", "0.0.0.0", "3306")
	log.Println(conn, err)

	file := pkg.NewFile()
	content, _ := file.Read("env/db.yml")

	yml := pkg.NewYaml()
	data, err := yml.ToMap(content)
	log.Println(data, err)
	//if err != nil {
	//	panic(err)
	//}
	//pkg.AutoMigrate(conn, memModel.MembersModel{})
	//pkg.AutoMigrate(conn, regisCatModel.RegisterCategoriesModel{})
}
