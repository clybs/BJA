package utils

import "testing"

func TestAuth_CheckPasswordHash(t *testing.T) {
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name string
		at   *Auth
		args args
		want bool
	}{
		{
			"hash don't match",
			&Auth{},
			args{"123456", "abc"},
			false,
		},
		{
			"hash match",
			&Auth{},
			args{"123456", "$2a$04$CaqFgizWM2cgRcUsSprk8.H6roecJ0DUIXWft5/w4pEvos83NTFrS"},
			true,
		},
		{
			"different hash match",
			&Auth{},
			args{"123456", "$2a$04$XEfKajL7v7VKUNSixVRIj.bW4bs//SyPHQtoy7cOrnfy73vDKJawC"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.at.CheckPasswordHash(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("Auth.CheckPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuth_CreateToken(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		at   *Auth
		args args
		want int
	}{
		{
			"length greater than 10",
			&Auth{},
			args{"1"},
			10,
		},
		{
			"length greater than 20",
			&Auth{},
			args{"123456"},
			20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.at.CreateToken(tt.args.id); len(got) < tt.want {
				t.Errorf("Auth.CreateToken() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestAuth_HashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		at      *Auth
		args    args
		want    int
		wantErr bool
	}{
		{
			"length greater than 10",
			&Auth{},
			args{""},
			10,
			false,
		},
		{
			"length greater than 20",
			&Auth{},
			args{"1"},
			20,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.at.HashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Auth.HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < tt.want {
				t.Errorf("Auth.HashPassword() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
