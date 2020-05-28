package convert

import "testing"

func TestToString(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
		wantError bool // エラーを期待するか
	}{
		{name: "String", args: args{v: "hoge"}, want: "hoge", wantError: false},
		{name: "Int", args: args{v: 1},want: "1", wantError: false},
		{name: "Bool", args: args{v: true}, want: "true", wantError: false},
		{name: "Other", args: args{v: 123.0}, want: "", wantError: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToString(tt.args.v)
			// エラーを期待しているにもかかわらず発生しない
			if err == nil && tt.wantError {
				t.Error("No error occurred.")
				return
			}
			// エラーを期待していないにもかかわらず発生した
			if err != nil && !tt.wantError {
				t.Error("An error has occurred.")
				t.Error(err.Error())
				return
			}
			if got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
