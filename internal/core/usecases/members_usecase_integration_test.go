//go:build integration
// +build integration

package usecases

//conn , err := pkg.NewConnectMysql()
//
//func Test_membersUseCase_NewMember_Integration(t *testing.T) {
//	file := pkg.NewFile()
//
//	c, _ := file.Read("../../../env/db.yml")
//	yaml := pkg.NewYaml()
//	strToMap, _ := yaml.ToMap(c)
//	config := strToMap["MYSQL"].(map[string]interface{})
//	s2 := strconv.Itoa(config["PORT"].(int))
//	conn, err := pkg.NewConnectMysql(config["USER"].(string), config["PASS"].(string), config["DB_NAME"].(string), config["HOST"].(string), s2)
//	log.Println(conn, err)
//
//	type fields struct {
//		membersRepo        ports.MembersRepository
//		RegisterCategories ports.RegisterCategories
//		uidService         ports.IUuidService
//		cryptoService      ports.ICryptoService
//	}
//
//	type args struct {
//		user     string
//		pass     string
//		fistName string
//		lastName string
//		dob      time.Time
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *domain.Members
//		wantErr bool
//	}{
//		//{ "test create member should be success",
//		//	fields{
//		//		mockMembersRepo,
//		//		mockCatepgoriesRepo,
//		//		mockUidService,
//		//		mockCryptoService,
//		//	},
//		//	args{
//		//		"root",
//		//		"root",
//		//		"root",
//		//		"root",
//		//		time.Now(),
//		//	},
//		//	mockRequest,
//		//	false,
//		//},
//		//{
//		//	"test create member should be error because username is used",
//		//	fields{
//		//		mockMembersRepoCaseTwo,
//		//		mockCatepgoriesRepo,
//		//		mockUidService,
//		//		mockCryptoService,
//		//	},
//		//	args{
//		//		"root",
//		//		"root",
//		//		"root",
//		//		"root",
//		//		time.Now(),
//		//	},
//		//	&domain.Members{},
//		//	true,
//		//},
//	}
//	for _, tt := range tests {
//
//		t.Run(tt.name, func(t *testing.T) {
//			m := membersUseCase{
//				membersRepo:        tt.fields.membersRepo,
//				RegisterCategories: tt.fields.RegisterCategories,
//				UidService:         tt.fields.uidService,
//				CryptoService:      tt.fields.cryptoService,
//			}
//			got, err := m.NewMember(tt.args.user, tt.args.pass, tt.args.fistName, tt.args.lastName, tt.args.dob)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("NewMember() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewMember() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//
//func Test_membersUseCase_FindMemberById_Integration(t *testing.T) {
//	id := "9cb6eb12-0b83-4034-92de-61d0e6256699"
//	mockResponse := &domain.Members{Mid: id, Username: "root", Password: "root", FirstName: "root", LastName: "root", DateOfBird: time.Now(), RegisterType: 1, CreatedAt: time.Now()}
//	mockMembersRepoCaseOne := new(_membersMock.MembersRepository)
//	mockMembersRepoCaseTwo := new(_membersMock.MembersRepository)
//
//	mockCatepgoriesRepo := new(_membersMock.RegisterCategories)
//	mockMembersRepoCaseOne.On("Get", mock.AnythingOfType("string")).Return(mockResponse, nil)
//	mockMembersRepoCaseTwo.On("Get", "random").Return(&domain.Members{}, nil)
//
//	type fields struct {
//		membersRepo        ports.MembersRepository
//		RegisterCategories ports.RegisterCategories
//	}
//	type args struct {
//		id string
//	}
//
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *domain.Members
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//		{
//			"test get member by id should success",
//			fields{mockMembersRepoCaseOne, mockCatepgoriesRepo},
//			args{id: id},
//			mockResponse,
//			false,
//		},
//		{
//			"test get member by id should error because not found member",
//			fields{mockMembersRepoCaseTwo, mockCatepgoriesRepo},
//			args{id: "random"},
//			&domain.Members{},
//			true,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			m := membersUseCase{
//				membersRepo:        tt.fields.membersRepo,
//				RegisterCategories: tt.fields.RegisterCategories,
//			}
//			got, err := m.FindMemberById(tt.args.id)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("FindMemberById() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("FindMemberById() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
