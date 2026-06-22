package httpx

import "context"

type ContextKey string

const UserContextKey ContextKey = "user"

type AuthUser struct {
	PublicID string
	Role     string
}

func SetUser(ctx context.Context, user any) context.Context {
	return context.WithValue(ctx, UserContextKey, user)
}

func GetAuthUser(ctx context.Context) (AuthUser, bool) {
	u, ok := ctx.Value(UserContextKey).(AuthUser)
	return u, ok
}
