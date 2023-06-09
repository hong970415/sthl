package authentication

import (
	"context"
	"net/http"
	"sthl/constants"
	"sthl/utils"
	"strings"

	"go.uber.org/zap"
)

type Authenticator interface {
	Authenticate(ctx context.Context, token string) (string, error)
}

func tokenFromHeader(r *http.Request) string {
	// get token from authorization header.
	bearer := r.Header.Get("authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}

func AuthInterceptor(l *zap.Logger, authentor Authenticator) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// extract token
			tokenString := tokenFromHeader(r)
			ctx := r.Context()
			result, err := authentor.Authenticate(ctx, tokenString)
			if err != nil {
				utils.ResponseSend[any](w, http.StatusUnauthorized, "", nil)
				l.Info("auth not pass")
				return
			}
			l.Info("auth pass", zap.Any("result", result))

			r = r.WithContext(context.WithValue(ctx, constants.AccessTokenKey, tokenString))
			r = r.WithContext(context.WithValue(ctx, constants.AccessTokenInfoKey, result))
			next.ServeHTTP(w, r)
		})
	}
}

// import (
// 	"context"
// 	"eee/shared/constants"
// 	"eee/shared/utils"

// 	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
// 	"go.uber.org/zap"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/codes"
// )

// type Authenticator interface {
// 	Authenticate(ctx context.Context, token string) (string, error)
// }

// func GetTokenInfoFromCtx(ctx context.Context) (string, bool) {
// 	tokenInfo, ok := ctx.Value(constants.AccessTokenInfoKey).(string)
// 	return tokenInfo, ok
// }
// func GetAccessTokenFromCtx(ctx context.Context) (string, bool) {
// 	token, ok := ctx.Value(constants.AccessTokenKey).(string)
// 	return token, ok
// }

// // AuthInterceptor
// // get grpc Method name, check whether in whitelist
// // yes: return original ctx, skip authentication
// // no:
// // return new ctx with accessToken and userId if pass authentication;
// // return err response if not pass authentication
// // Authenticator: depends on which microservice, call local if account service; call remotely if others
// func AuthInterceptor(logger *zap.Logger, whitelist map[string]bool, authtor Authenticator) grpc_auth.AuthFunc {
// 	return func(ctx context.Context) (context.Context, error) {
// 		fullMethod, ok := grpc.Method(ctx)
// 		if !ok {
// 			logger.Info("Fail to get fullMethod")
// 			return nil, utils.ErrResponse(constants.ErrInternalServer, codes.Internal)
// 		}

// 		logger.Info("fullMethod", zap.String("val", fullMethod))
// 		if whitelist[fullMethod] {
// 			return ctx, nil
// 		}

// 		accessToken, err := grpc_auth.AuthFromMD(ctx, "bearer")
// 		if err != nil {
// 			logger.Info("Fail to get token from MD", zap.Error(err))
// 			return nil, utils.ErrResponse(constants.ErrUnauthorized, codes.Unauthenticated)
// 		}

// 		userId, err := authtor.Authenticate(ctx, accessToken)
// 		if err != nil {
// 			logger.Info("Fail to authenticate", zap.Error(err))
// 			return nil, utils.ErrResponse(constants.ErrUnauthorized, codes.Unauthenticated)
// 		}

// 		logger.Info("Pass", zap.Any("userId", userId))
// 		newCtx1 := context.WithValue(ctx, constants.AccessTokenKey, accessToken)
// 		newCtx2 := context.WithValue(newCtx1, constants.AccessTokenInfoKey, userId)
// 		return newCtx2, nil
// 	}
// }
