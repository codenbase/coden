package token

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Config struct {
	key         string
	identityKey string
	expiration  time.Duration
}

var (
	config = Config{"Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5", "identityKey", 2 * time.Hour}
	once   sync.Once
)

func Init(key string, identityKey string, expiration time.Duration) {
	once.Do(func() {
		if key != "" {
			config.key = key
		}

		if identityKey != "" {
			config.identityKey = identityKey
		}

		if expiration > 0 {
			config.expiration = expiration
		}
	})
}

func Parse(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(config.key), nil
	})

	if err != nil {
		return "", err
	}

	var identityKey string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if key, exists := claims[config.identityKey]; exists {
			if identity, valid := key.(string); valid {
				identityKey = identity
			}
		}
	}

	if identityKey == "" {
		return "", jwt.ErrSignatureInvalid
	}

	return identityKey, nil
}

func ParseRequest(ctx context.Context) (string, error) {
	var (
		token string
		err   error
	)

	switch typed := ctx.(type) {
	// 使用 Gin 框架开发的 HTTP 服务
	case *gin.Context:
		header := typed.Request.Header.Get("Authorization")
		if len(header) == 0 {
			//nolint: err113
			return "", errors.New("the length of the `Authorization` header is zero")
		}

		// 从请求头中取出 token
		_, _ = fmt.Sscanf(header, "Bearer %s", &token)
	// 使用 google.golang.org/grpc 框架开发的 gRPC 服务
	default:
		token, err = auth.AuthFromMD(typed, "Bearer")
		if err != nil {
			return "", status.Errorf(codes.Unauthenticated, "invalid auth token")
		}
	}

	return Parse(token)
}

func Sign(identityKey string) (string, time.Time, error) {
	now := time.Now()
	// 计算过期时间
	expireAt := now.Add(config.expiration)

	// Token 的内容
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		config.identityKey: identityKey,     // 存放用户身份
		"nbf":              now.Unix(),      // token 生效时间
		"iat":              now.Unix(),      // token 签发时间
		"exp":              expireAt.Unix(), // token 过期时间
	})

	if config.key == "" {
		return "", time.Time{}, jwt.ErrInvalidKey
	}

	// 签发 token
	tokenString, err := token.SignedString([]byte(config.key))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expireAt, nil // 返回 token 字符串、过期时间和错误
}
