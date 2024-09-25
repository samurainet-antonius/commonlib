package validation

import "testing"

func TestIsValidPhone(t *testing.T) {
	type args struct {
		phone string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should accept number with + prefix ",
			args: args{
				phone: "+628999239159",
			},
			want: true,
		},
		{
			name: "should accept number length 15 ",
			args: args{
				phone: "003729847289347",
			},
			want: true,
		},
		{
			name: "should accept number length 12 ",
			args: args{
				phone: "372984728934",
			},
			want: true,
		},
		{
			name: "should accept number length 11 ",
			args: args{
				phone: "72984728934",
			},
			want: true,
		},
		{
			name: "should accept number length 10 ",
			args: args{
				phone: "7298472893",
			},
			want: true,
		},
		{
			name: "should accept number length 9 ",
			args: args{
				phone: "729847289",
			},
			want: true,
		},
		{
			name: "should accept number length 7 ",
			args: args{
				phone: "7298472",
			},
			want: true,
		},
		{
			name: "should accept number without + prefix ",
			args: args{
				phone: "7298472",
			},
			want: true,
		},
		{
			name: "should not accept number with 5 length ",
			args: args{
				phone: "98472",
			},
			want: false,
		},
		{
			name: "should not accept number with 20 length ",
			args: args{
				phone: "00372984728934789347",
			},
			want: false,
		},
		{
			name: "should not accept number with wrong prefix ",
			args: args{
				phone: "*628999239159",
			},
			want: false,
		},
		{
			name: "should not accept number with length 16 ",
			args: args{
				phone: "*0037298472893479",
			},
			want: false,
		},
		{
			name: "should not accept number with length 6 ",
			args: args{
				phone: "*984721",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidPhone(tt.args.phone); got != tt.want {
				t.Errorf("isValidPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}
