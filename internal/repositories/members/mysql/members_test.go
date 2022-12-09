package mysql

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	domain "lean-oauth/internal/core/domains"
	"lean-oauth/pkg"
	"reflect"
	"testing"
	"time"
)

// connect to real database
var conn, err = pkg.NewConnect("root", "root", "oauth", "0.0.0.0", "3306").Connect()

func Test_membersMysqlRepo_Create(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		todo *domain.Members
	}

	newMember := domain.NewMember(uuid.New(), "root", "root", "root", "root", time.Now(), 1, time.Now())

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Members
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test insert member should be success",
			fields{
				conn,
			},
			args{
				newMember,
			},
			newMember,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := membersMysqlRepo{
				db: tt.fields.db,
			}
			got, err := m.Create(tt.args.todo)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_membersMysqlRepo_Get(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"Test get member by id should success",
			fields{conn},
			args{"9cb6eb62-0b83-4034-92de-61d0e6256699"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := membersMysqlRepo{
				db: tt.fields.db,
			}
			_, err := m.Get(uuid.MustParse(tt.args.id))
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Get() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
