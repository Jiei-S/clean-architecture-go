package controller_test

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	rest "github.com/Jiei-S/boilerplate-clean-architecture/go-rest/internal/infrastructure/openapi"

	"github.com/Jiei-S/boilerplate-clean-architecture/go-rest/internal/adapter/gateway"
	"github.com/Jiei-S/boilerplate-clean-architecture/go-rest/internal/infrastructure/bun"
	"github.com/stretchr/testify/assert"
)

const BASE_API_URL = "http://localhost:8081/users"

func TestAddUser(t *testing.T) {
	defer func() {
		db := bun.NewDB()
		db.NewTruncateTable().Model(&gateway.User{}).Exec(context.Background())
	}()

	type args struct {
		body string
	}
	type test struct {
		name string
		args
		want rest.User
		code int
	}

	tests := []test{
		{
			name: "success",
			args: args{
				body: `{"firstName":"test","lastName":"user","age":20}`,
			},
			want: rest.User{
				FirstName: "test",
				LastName:  "user",
				Age:       20,
			},
			code: http.StatusOK,
		},
		{
			name: "bad request",
			args: args{
				body: `{"firstName":"test","lastName":"user","age":"20"}`,
			},
			code: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.Post(BASE_API_URL, "application/json", strings.NewReader(tt.args.body))
			if err != nil {
				t.Fatal(err)
			}
			defer r.Body.Close()

			var act rest.User
			if err := json.NewDecoder(r.Body).Decode(&act); err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.code, r.StatusCode)
			if tt.code == http.StatusOK {
				assert.Equal(t, tt.want.FirstName, act.FirstName)
			}
		})
	}
}

func TestFindUser(t *testing.T) {
	db := bun.NewDB()
	user := &gateway.User{
		FirstName: "test",
		LastName:  "user",
		Age:       20,
	}
	db.NewInsert().Model(user).Exec(context.Background())
	defer func() {
		db.NewTruncateTable().Model(&gateway.User{}).Exec(context.Background())
	}()

	type args struct {
		id string
	}
	tests := []struct {
		name string
		args
		want rest.User
		code int
	}{
		{
			name: "success",
			args: args{
				id: user.ID,
			},
			want: rest.User{
				Id:        user.ID,
				FirstName: "test",
				LastName:  "user",
				Age:       20,
			}, code: http.StatusOK,
		},
		{
			name: "not found",
			args: args{
				id: "1",
			},
			code: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.Get(BASE_API_URL + "/" + tt.args.id)
			if err != nil {
				t.Fatal(err)
			}
			defer r.Body.Close()

			var act rest.User
			if err := json.NewDecoder(r.Body).Decode(&act); err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.code, r.StatusCode)
			if tt.code == http.StatusOK {
				assert.Equal(t, tt.want.FirstName, act.FirstName)
			}
		})
	}
}
