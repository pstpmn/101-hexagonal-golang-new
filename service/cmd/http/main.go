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

//https://www.facebook.com/v15.0/dialog/oauth?
//client_id=817832142661743
//&redirect_uri=http://localhost:3000/
//&state=

//GET https://graph.facebook.com/v15.0/oauth/access_token?
//client_id=817832142661743
//&redirect_uri=https://api.autoplatf.com/
//&client_secret=ec98a91f0a9c4f1aba82672347b9f7f1
//&code=AQBqtnaJRyxfTF85I1m7brKy3oBlzAe9cCCjTNW91-n9KzozBw9_p7iYI4Vkw1XBG1SxWlOf7XkFkzVyJpZqom_WGRSgU6kVPaRQslX_1Gx1x0aofq20BtJS5PLKj1qPZMfoYA3g3-Ps0nVkblmfvJPDZ6D93wUU3OBWkB6yQt91O51ZQSznBBpRZGX532fkpUu4NO9C9AVUpjEnZaTJOOxR1AQ6N6iReINxjuE6H1hNjfwmNaa12assYUTGIM8pyOxag0ZuJlpPxaVdKyYVmjyw4tHnCssvGXDeAZoQf-edLoseBxbsf7HJiYrx1P9kL6b4hWsnTtji1KBEll4MO27zFM86CsL9Mkh7kypeUrC-0MKuGXPCfAf2LD88jnCvL6iVagiHccMcLMy3k1GUsKcxO6klfJ-0z2p4jOQZeyJe5g#_=_
//
//
//{"access_token":"EAALn0GJZAaG8BAHZBEY1ZB4x38tGg7mZAHZB2JUtpxer2AkRxDdDhPXbzgz8MGSyztZCveNFo2wbGTvBjFMZC54OEZAFeT3g5pjSzKu2D2ZCQ0FA5mPI1Fco0MNFZC0ELgW17a0nAY7n1iZA5lnL77pqZCxNR3wIXfAwiplQqcZARJZCG4NVcHzfLkZCkz89nlzzHNUv42Gw2Qo6XCU3EnJrjH0ecweGjbgvVs9jlZAtM1ZBzCGqe6QZDZD","token_type":"bearer","expires_in":5181110}
//{"access_token":"EAALn0GJZAaG8BAEyHN4tqgJmFN7uiCS7LBQmBaieJxqBEzXnA9JYDA6D8V1yC4I0PgjGywZC0NYNihF2jmKL1e8sc64WZA20Lyt8kZAU86BM8M5GKZCZB7UZAjE2GLADr71y312lppu6WcH0eZBfEWd9LMMUNCIBXekrO0LAxLbxaqNAgUTQwyTZBrk2X0QZAE05IQRt3GdK23vmfpBABTInWCaTOrChXKq0cWpCUyZCHcqVwZDZD","token_type":"bearer","expires_in":5181240}
//
//GET graph.facebook.com/debug_token?
//input_token=2040cbf8d424051ead7c6179618ec22f
//&access_token=EAALn0GJZAaG8BAEyHN4tqgJmFN7uiCS7LBQmBaieJxqBEzXnA9JYDA6D8V1yC4I0PgjGywZC0NYNihF2jmKL1e8sc64WZA20Lyt8kZAU86BM8M5GKZCZB7UZAjE2GLADr71y312lppu6WcH0eZBfEWd9LMMUNCIBXekrO0LAxLbxaqNAgUTQwyTZBrk2X0QZAE05IQRt3GdK23vmfpBABTInWCaTOrChXKq0cWpCUyZCHcqVwZDZD
