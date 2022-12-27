package mysql

import (
	"learn-oauth2/pkg"
)

var conn, err = pkg.NewConnectMysql("root", "root", "oauth", "0.0.0.0", "3306")

//
//func Test_registerCategoriesMysqlRepo_List(t *testing.T) {
//	type fields struct {
//		db *gorm.DB
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		//want    []domain.RegisterCategories
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//		{
//			"test select all value",
//			fields{db: conn.GetInstance()},
//			//[]domain.RegisterCategories{},
//			false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := registerCategoriesMysqlRepo{
//				db: tt.fields.db,
//			}
//			_, err := r.List()
//			if (err != nil) != tt.wantErr {
//				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			//if !reflect.DeepEqual(got, tt.want) {
//			//	t.Errorf("List() got = %v, want %v", got, tt.want)
//			//}
//		})
//	}
//}
