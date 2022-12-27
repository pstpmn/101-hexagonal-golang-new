//go:build integration
// +build integration

package pkg

// func Test_conn_Connect(t *testing.T) {
//	type fields struct {
//		User   string
//		Pass   string
//		DbName string
//		Host   string
//		Port   string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//		{
//			"test connect should be success",
//			fields{User: "root", Pass: "root", DbName: "oauth", Host: "0.0.0.0", Port: "3306"},
//			false,
//		},
//		{
//			"test connect should be error because invalid username or password ",
//			fields{User: "root1", Pass: "root1", DbName: "oauth", Host: "0.0.0.0", Port: "3306"},
//			true,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := config{
//				User:   tt.fields.User,
//				Pass:   tt.fields.Pass,
//				DbName: tt.fields.DbName,
//				Host:   tt.fields.Host,
//				Port:   tt.fields.Port,
//			}
//			_, err := c.Connect()
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//		})
//	}
//}
