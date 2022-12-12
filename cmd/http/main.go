package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"lean-oauth/internal/core/usecases"
	"lean-oauth/internal/handlers/http/members"
	membersRepositories "lean-oauth/internal/repositories/members/mysql"
	categoriesRepositories "lean-oauth/internal/repositories/register_categories/mysql"
	"lean-oauth/pkg"
	"strconv"
)

func main() {
	file := pkg.NewFile()
	yml := pkg.NewYaml()
	crypto := pkg.NewCrypto()
	uuid := pkg.NewUuId()

	// prepare environments
	dbStr, _ := file.Read("env/db.yml")
	appStr, _ := file.Read("env/app.yml")
	dbEnv, _ := yml.ToMap(dbStr)
	appEnv, _ := yml.ToMap(appStr)

	mysqlEnv := dbEnv["MYSQL"].(map[string]interface{})
	serviceEnv := appEnv["SERVICE"].(map[string]interface{})
	mysqlPort := strconv.Itoa(mysqlEnv["PORT"].(int))
	servicePort := strconv.Itoa(serviceEnv["PORT"].(int))

	// connect db
	var conn, err = pkg.NewConnectMysql(mysqlEnv["USER"].(string), mysqlEnv["PASS"].(string), mysqlEnv["DB_NAME"].(string), mysqlEnv["HOST"].(string), mysqlPort)
	if err != nil {
		// err connection
		panic(err)
	}

	// repositories
	membersRepo := membersRepositories.NewMembersMysqlRepo(conn.GetInstance())
	categoriesRepo := categoriesRepositories.NewRegisterCategoriesMysqlRepo(conn.GetInstance())

	// usecases
	membersUsercase := usecases.NewMembersUseCase(membersRepo, categoriesRepo, uuid, crypto)

	// handlers
	handlers := handlers.NewHTTPHandler(membersUsercase)

	// create server
	app := fiber.New()
	app.Get("/", handlers.HelloWorld)
	app.Listen(fmt.Sprintf(":%s", servicePort))
}
