package main

import (
	"fmt"
	"learn-oauth2/internal/core/usecases"
	handlers "learn-oauth2/internal/protocal/http/handlers/members"
	"learn-oauth2/internal/protocal/http/middlewares"
	membersRepositories "learn-oauth2/internal/repositories/members/mysql"
	categoriesRepositories "learn-oauth2/internal/repositories/register_categories/mysql"
	server2 "learn-oauth2/internal/server"
	pkg2 "learn-oauth2/pkg"
)

func main() {
	file := pkg2.NewFile()
	yml := pkg2.NewYaml()
	crypto := pkg2.NewCrypto()
	uuid := pkg2.NewUuId()
	jwt := pkg2.NewJsonWebToken()

	// prepare environments
	dbStr, _ := file.Read("env/db.yml")
	appStr, _ := file.Read("env/app.yml")
	dbEnv, _ := yml.ToMap(dbStr)
	appEnv, _ := yml.ToMap(appStr)

	mysqlEnv := dbEnv["MYSQL"].(map[string]interface{})
	serviceEnv := appEnv["SERVICE"].(map[string]interface{})
	authKey := fmt.Sprint(appEnv["AUTH_KEY"])

	// connect db
	var conn, err = pkg2.NewConnectMysql(
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
	handlers := handlers.NewHTTPHandler(membersUsercase, server2.NewResponse(), authKey)

	// middlewares
	middlewares := middlewares.NewHTTPMiddleware(membersUsercase, server2.NewResponse(), authKey)

	server := server2.NewServer(handlers, middlewares, serviceEnv)
	server.Initialize()
}
