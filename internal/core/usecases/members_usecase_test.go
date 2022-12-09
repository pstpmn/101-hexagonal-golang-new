package usecases

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	domain "lean-oauth/internal/core/domains"
	"lean-oauth/internal/core/ports"
	_membersMock "lean-oauth/internal/core/ports/mocks"
	"log"
	"testing"
	"time"
)

func Test_membersUseCase_NewMember(t *testing.T) {
	mockRequest := &domain.Members{uuid.New(), "1", "root", "root", "root", time.Now(), 1, time.Now()}
	mockMembersRepo := new(_membersMock.MembersRepository)
	mockCatepgoriesRepo := new(_membersMock.RegisterCategories)
	// บอกว่าเราเรียกใช้ function ไหนของ Repository
	mockMembersRepo.On("Create", mock.AnythingOfType("*domains.Members")).Return(mockRequest, nil)

	type fields struct {
		membersRepo        ports.MembersRepository
		RegisterCategories ports.RegisterCategories
	}
	type args struct {
		user     string
		pass     string
		fistName string
		lastName string
		dob      time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		//want    *domain.Members
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"",
			fields{
				mockMembersRepo,
				mockCatepgoriesRepo,
			},
			args{
				"root2",
				"root",
				"root",
				"root",
				time.Now(),
			},
			//mockRequest,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := membersUseCase{
				membersRepo:        tt.fields.membersRepo,
				RegisterCategories: tt.fields.RegisterCategories,
			}
			d, err := m.NewMember(tt.args.user, tt.args.pass, tt.args.fistName, tt.args.lastName, tt.args.dob)
			log.Println("1 : ", d)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("NewMember() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func Test_membersUseCase_FindMemberById(t *testing.T) {
	id := uuid.MustParse("9cb6eb12-0b83-4034-92de-61d0e6256699")
	mockResponse := &domain.Members{Mid: id, Username: "root", Password: "root", FirstName: "root", LastName: "root", DateOfBird: time.Now(), RegisterType: 1, CreatedAt: time.Now()}

	mockMembersRepo := new(_membersMock.MembersRepository)
	mockCatepgoriesRepo := new(_membersMock.RegisterCategories)
	mockMembersRepo.On("Get", mock.AnythingOfType("UUID")).Return(mockResponse, nil)

	type fields struct {
		membersRepo        ports.MembersRepository
		RegisterCategories ports.RegisterCategories
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		//want    *domain.Members
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test get member by id should success",
			fields{mockMembersRepo, mockCatepgoriesRepo},
			args{id: uuid.New()},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := membersUseCase{
				membersRepo:        tt.fields.membersRepo,
				RegisterCategories: tt.fields.RegisterCategories,
			}
			got, err := m.FindMemberById(tt.args.id)
			log.Println(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindMemberById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("FindMemberById() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
