package main

import (
	"fmt"
	"learn-oauth2/internal/core/usecases"
	handlers "learn-oauth2/internal/protocal/http/handlers/members"
	"learn-oauth2/internal/protocal/http/middlewares"
	membersRepositories "learn-oauth2/internal/repositories/members/mysql"
	categoriesRepositories "learn-oauth2/internal/repositories/register_categories/mysql"
	"learn-oauth2/internal/server"
	"learn-oauth2/pkg"
)

func main() {
	file := pkg.NewFile()
	yml := pkg.NewYaml()
	crypto := pkg.NewCrypto()
	uuid := pkg.NewUuId()
	jwt := pkg.NewJsonWebToken()
	request := pkg.NewRequests()

	// prepare environments
	dbStr, _ := file.Read("env/db.yml")
	appStr, _ := file.Read("env/app.yml")
	oauthStr, _ := file.Read("env/oauth.yml")
	dbEnv, _ := yml.ToMap(dbStr)
	appEnv, _ := yml.ToMap(appStr)
	oauthEnv, _ := yml.ToMap(oauthStr)

	mysqlEnv := dbEnv["MYSQL"].(map[string]interface{})
	serviceEnv := appEnv["SERVICE"].(map[string]interface{})
	facebookEnv := oauthEnv["facebook"].(map[string]interface{})
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
	membersUsercase := usecases.NewMembersUseCase(membersRepo, categoriesRepo, uuid, crypto, jwt, request)
	oauth2UseCase := usecases.NewOauth2UseCase(request)

	// protocal
	handlers := handlers.NewHTTPHandler(membersUsercase, oauth2UseCase, server.NewResponse(), authKey, facebookEnv)

	// middlewares
	middlewares := middlewares.NewHTTPMiddleware(membersUsercase, server.NewResponse(), authKey)

	server := server.NewServer(handlers, middlewares, serviceEnv)
	server.Initialize()
}

//
////https://www.facebook.com/v15.0/dialog/oauth?client_id=817832142661743&redirect_uri=https://api.autoplatf.com/&state=
//
//https://graph.facebook.com/v15.0/oauth/access_token?
//client_id=817832142661743
//&redirect_uri=https://api.autoplatf.com/
//&client_secret=ec98a91f0a9c4f1aba82672347b9f7f1
//&code=AQCHbxStIMHKixiRPuQ5gItypResYAIyaB4LoTJSHj5uRE5udseEmJUZjECvChVlF7J8mdMP0roHFMQeA3mDXi-P-niiXUzVzwBBeM3-atBqKfFIR2x-hrKSY0p5Tmdi88nOD0RC9mrhaRKV3N9tSOc-bPSBuKRctFwI4SL2XXbLr0H1fLCcxyXUMOEZQ-BpwPJ_0ffP2ZzL4uuCqW0GnHQhTv2bl94tT4qyHGd4sjPvdKZQuXmMZKVWVKVsk9QwkyG4savkukImKgqxtbmYtEN-8zUhsStHDHjqD28d_UUQ4RnEZidrlM8FlByWEl8VHei65lPNCVQU8zlOcGglOenxwIVv4aswtSFAfbQuAqf1vN3tAyBy2Wem_cXhNXixcopEt5n54ZRy7W53QX-_8xYlR0Z8NL-ftOofh9H-h4oekw#_=_
//
////
////{"access_token":"EAALn0GJZAaG8BAHZBEY1ZB4x38tGg7mZAHZB2JUtpxer2AkRxDdDhPXbzgz8MGSyztZCveNFo2wbGTvBjFMZC54OEZAFeT3g5pjSzKu2D2ZCQ0FA5mPI1Fco0MNFZC0ELgW17a0nAY7n1iZA5lnL77pqZCxNR3wIXfAwiplQqcZARJZCG4NVcHzfLkZCkz89nlzzHNUv42Gw2Qo6XCU3EnJrjH0ecweGjbgvVs9jlZAtM1ZBzCGqe6QZDZD","token_type":"bearer","expires_in":5181110}
////{"access_token":"EAALn0GJZAaG8BAEyHN4tqgJmFN7uiCS7LBQmBaieJxqBEzXnA9JYDA6D8V1yC4I0PgjGywZC0NYNihF2jmKL1e8sc64WZA20Lyt8kZAU86BM8M5GKZCZB7UZAjE2GLADr71y312lppu6WcH0eZBfEWd9LMMUNCIBXekrO0LAxLbxaqNAgUTQwyTZBrk2X0QZAE05IQRt3GdK23vmfpBABTInWCaTOrChXKq0cWpCUyZCHcqVwZDZD","token_type":"bearer","expires_in":5181240}
////EAALn0GJZAaG8BAOHpLqLFYOFdLfLuRzETzmWrgOAEycnNdkOqwmfnViznxQkg7AOBuvdZCyxJmUWjcVEEHfXyZAoDHyFQyNUznEYsQfzozGUwTaZBiiRCgcKvmSVU466g1ZBbl4vkkabw0lrYgHqLPJTZBzrpf51ZA0e6pUcda1uCDaLhmMpIN3GZBNvIDO3oo3Wbqjk3ZB38upxJTyrnYiC91SZCN6YAExfUwo7ZCVO1WOIAZDZD
//
//graph.facebook.com/debug_token?
//input_token=EAALn0GJZAaG8BACKouZCfM5nzflnwsTy2PJWrOZAql0JmYPBYQc4vRKbl2y5flJzNhMJCRm6QZBOnFyZANGDRLifemxrBo9YJ0EAhfiVJwzN5wG51zJDxH1xYGiKmGCBG2HVkNHUrZCqBEIraCadQPtcpUJdNBA7EVIXBPc8VfLq5ZBR81xaLGTphqdJZAdufHgjliFy2rtE1gWFWUzl8fFm
//&access_token=817832142661743|_Vx1271myLT4XIfOBUzSiJWYwRU
//

//https://graph.facebook.com/oauth/access_token?client_id=817832142661743&client_secret=ec98a91f0a9c4f1aba82672347b9f7f1&grant_type=client_credentials
//
////817832142661743|_Vx1271myLT4XIfOBUzSiJWYwRU
