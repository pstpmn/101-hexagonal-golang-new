package pkg

//func TestRequestJson(t *testing.T) {
//	type args struct {
//		method  string
//		url     string
//		header  map[string]string
//		mapBody map[string]interface{}
//	}
//	tests := []struct {
//		name string
//		args args
//		//want    map[string]interface{}
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//		{
//			"Test Intergration Send Request",
//			args{"GET", "https://covid19.ddc.moph.go.th/api/Cases/round-4-line-lists", map[string]string{"Content-Type": "application/json"}, nil},
//			//nil,
//			false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := Json(tt.args.method, tt.args.url, tt.args.header, tt.args.mapBody)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("RequestJson() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			log.Println("got :", got)
//			//if !reflect.DeepEqual(got, tt.want) {
//			//	t.Errorf("RequestJson() got = %v, want %v", got, tt.want)
//			//}
//		})
//	}
//}
