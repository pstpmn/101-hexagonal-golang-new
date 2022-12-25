package main

import (
	"fmt"
	"lean-oauth/internal/core/usecases"
	"lean-oauth/internal/protocal/http/handlers/members"
	"lean-oauth/internal/protocal/http/middlewares"
	membersRepositories "lean-oauth/internal/repositories/members/mysql"
	categoriesRepositories "lean-oauth/internal/repositories/register_categories/mysql"
	"lean-oauth/internal/server"
	"lean-oauth/pkg"
)

func main() {
	file := pkg.NewFile()
	yml := pkg.NewYaml()
	crypto := pkg.NewCrypto()
	uuid := pkg.NewUuId()
	jwt := pkg.NewJsonWebToken()

	// prepare environments
	dbStr, _ := file.Read("env/db.yml")
	appStr, _ := file.Read("env/app.yml")
	dbEnv, _ := yml.ToMap(dbStr)
	appEnv, _ := yml.ToMap(appStr)

	mysqlEnv := dbEnv["MYSQL"].(map[string]interface{})
	serviceEnv := appEnv["SERVICE"].(map[string]interface{})
	authKey := fmt.Sprint(appEnv["AUTH_KEY"])

	// connect db
	var conn, err = pkg.NewConnectMysql(
		mysqlEnv["USER"].(string),
		mysqlEnv["PASS"].(string),
		mysqlEnv["DB_NAME"].(string),
		mysqlEnv["HOST"].(string),
		fmt.Sprintf("%d", mysqlEnv["PORT"].(int)),
	)

	if err != nil {
		panic(err)
	}

	// repositories
	membersRepo := membersRepositories.NewMembersMysqlRepo(conn.GetInstance())
	categoriesRepo := categoriesRepositories.NewRegisterCategoriesMysqlRepo(conn.GetInstance())

	// usecases
	membersUsercase := usecases.NewMembersUseCase(membersRepo, categoriesRepo, uuid, crypto, jwt)

	// protocal
	handlers := handlers.NewHTTPHandler(membersUsercase, server.NewResponse(), authKey)

	// middlewares
	middlewares := middlewares.NewHTTPMiddleware(membersUsercase, server.NewResponse(), authKey)

	server := server.NewServer(handlers, middlewares, serviceEnv)
	server.Initialize()
}
