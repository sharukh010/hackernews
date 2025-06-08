package auth

import (
	"context"
	"net/http"
	"strconv"

	"github.com/sharukh010/hackernews/internal/pkg/jwt"
	"github.com/sharukh010/hackernews/internal/users"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

func Middleware() func(http.Handler) http.Handler{
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
			header := r.Header.Get("Authorization")
			if header == ""{
				next.ServeHTTP(w,r)
				return 
			}
			tokenStr := header 
			username,err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w,"Invalid Token",http.StatusForbidden)
				return 
			}

			user := users.User{Username:username}
			id,err := users.GetUserIdByUsername(username)
			if err != nil {
				next.ServeHTTP(w,r)
				return 
			}
			user.ID = strconv.Itoa(id)

			//adds users name to the context and create new context
			ctx := context.WithValue(r.Context(),userCtxKey,&user)

			//copy the context to request
			r = r.WithContext(ctx)
			next.ServeHTTP(w,r)

		})
	}
}

func ForContext(ctx context.Context) *users.User {
	raw,_ := ctx.Value(userCtxKey).(*users.User)
	return raw
}