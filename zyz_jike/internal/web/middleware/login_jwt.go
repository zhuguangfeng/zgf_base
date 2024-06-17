package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	ijwt "zyz_jike/internal/web/jwt"
)

type LoginJwtMiddlewareBuilder struct {
	ijwt.Handler
}

func NewLoginJwtMiddlewareBuilder(hdl ijwt.Handler) *LoginJwtMiddlewareBuilder {
	return &LoginJwtMiddlewareBuilder{
		Handler: hdl,
	}
}

func (m *LoginJwtMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if path == "/app/user/pwd/login" ||
			path == "/app/user/register" {
			return
		}
		tokenStr := m.ExtractToken(ctx)
		var uc ijwt.UserClaims
		token, err := jwt.ParseWithClaims(tokenStr, &uc, func(token *jwt.Token) (interface{}, error) {
			return ijwt.JwtKey, nil
		})
		if err != nil {
			fmt.Println("1111")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if token == nil || !token.Valid {
			fmt.Println("2222")

			// token 解析出来了，但是 token 可能是非法的，或者过期了的
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		err = m.CheckSession(ctx, uc.Ssid)
		if err != nil {
			fmt.Println("3333")

			// token 无效或者 redis 有问题
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user", uc)
		fmt.Println(uc)
	}
}
