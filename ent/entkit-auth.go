// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entkit/entkit/v2"
	"github.com/gin-gonic/gin"
)

// _entkitDefaultValue to provide default values from entc definition and also get ability to provide from caller
func _entkitDefaultValue(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

func JeppAuthMiddleware(
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Entkit: Authentication/Authorization is currently disabled. To utilize JeppAuthMiddleware, please enable this feature.")
		next.ServeHTTP(w, r)
	})
}

func JeppAuthGinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Entkit: Authentication/Authorization is currently disabled. To utilize JeppAuthGinMiddleware, please enable this feature.")
		c.Next()
	}
}

type JeppResource int

const (
	JeppGameResource JeppResource = iota
	JeppSeasonResource
)

func (e JeppResource) String() string {
	switch e {
	case JeppGameResource:
		return "JeppGame"
	case JeppSeasonResource:
		return "JeppSeason"
	default:
		return "unknown"
	}
}

type JeppScope int

const (
	JeppReadScope JeppScope = iota
	JeppCreateScope
	JeppUpdateScope
	JeppDeleteScope
)

func (e JeppScope) String() string {
	switch e {
	case JeppReadScope:
		return "JeppRead"
	case JeppCreateScope:
		return "JeppCreate"
	case JeppUpdateScope:
		return "JeppUpdate"
	case JeppDeleteScope:
		return "JeppDelete"
	default:
		return "unknown"
	}
}

func JeppAuthorizeByResource(ctx context.Context, resource JeppResource, scope JeppScope) error {
	fmt.Println("Entkit: Authentication/Authorization is currently disabled. To utilize JeppAuthorizeByResource, please enable this feature.")
	return nil
}

func JeppEnforceGameRead(ctx context.Context) error {
	return JeppAuthorizeByResource(ctx, JeppGameResource, JeppReadScope)
}

func JeppEnforceGameCreate(ctx context.Context) error {
	return JeppAuthorizeByResource(ctx, JeppGameResource, JeppCreateScope)
}

func JeppEnforceGameUpdate(ctx context.Context) error {
	return JeppAuthorizeByResource(ctx, JeppGameResource, JeppUpdateScope)
}

func JeppEnforceGameDelete(ctx context.Context) error {
	return JeppAuthorizeByResource(ctx, JeppGameResource, JeppDeleteScope)
}

func JeppEnforceSeasonRead(ctx context.Context) error {
	return JeppAuthorizeByResource(ctx, JeppSeasonResource, JeppReadScope)
}

func JeppEnforceSeasonCreate(ctx context.Context) error {
	return JeppAuthorizeByResource(ctx, JeppSeasonResource, JeppCreateScope)
}

func JeppEnforceSeasonUpdate(ctx context.Context) error {
	return JeppAuthorizeByResource(ctx, JeppSeasonResource, JeppUpdateScope)
}

func JeppEnforceSeasonDelete(ctx context.Context) error {
	return JeppAuthorizeByResource(ctx, JeppSeasonResource, JeppDeleteScope)
}

func JeppAuthContextFromRequestContext(ctx context.Context) (*entkit.AuthContext, error) {
	fmt.Println("Entkit: Authentication/Authorization is currently disabled. To utilize JeppAuthContextFromRequestContext, please enable this feature.")
	return nil, nil
}
