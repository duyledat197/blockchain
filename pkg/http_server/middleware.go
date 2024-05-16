package http_server

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	mtdt "openmyth/blockchain/pkg/metadata"
	"openmyth/blockchain/util"
)

type middlewareFunc func(http.Handler) http.Handler

var (
	Authorization = "Authorization"
	Bearer        = "Bearer"
	InfoKey       = "info_key"
	CookieKey     = "h5token"

	ignoredAPIs         []string
	invalidateCacheAPIs []string
	ignoredForLogAPIs   []string
)

func init() {

	ignoredAPIs = []string{}

	ignoredForLogAPIs = []string{}

	invalidateCacheAPIs = append(
		invalidateCacheAPIs,
	)
}

type payloadKeys struct{}

// GetClientIP get client IP from HTTP request
func GetClientIP(req *http.Request) string {
	md, ok := metadata.FromIncomingContext(req.Context())
	if !ok {
		return ""
	}
	clientIP := md.Get(mtdt.MDXForwardedFor)
	if len(clientIP) == 0 {
		return ""
	}

	return clientIP[0]
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, authorization")
		if r.Method != "OPTIONS" {
			h.ServeHTTP(w, r)
		}
	})
}

type mapMetaDataFunc func(context.Context, *http.Request) metadata.MD

// MapMetaDataWithBearerToken ...
func MapMetaDataWithBearerToken() mapMetaDataFunc {
	return func(ctx context.Context, r *http.Request) metadata.MD {
		md := metadata.MD{}
		authorization := r.Header.Get(Authorization)
		log.Println("authorization", authorization)
		if authorization != "" {
			schema, bearerToken, ok := strings.Cut(authorization, " ")
			if !ok || strings.ToLower(schema) != strings.ToLower(Bearer) {
				return md
			}
			log.Println("zzzz")
			payload, err := util.VerifyToken(bearerToken)
			if err == nil {
				md = metadata.Join(md, mtdt.ImportUserInfoToCtx(&mtdt.Payload{
					UserID: payload.Id,
					Ip:     GetClientIP(r),
					Token:  bearerToken,
				}))
			}
		}

		return md
	}
}

// func MapMetaDataWithBearerToken(authenticator authenticate.Authenticator) mapMetaDataFunc {
// 	return func(ctx context.Context, r *http.Request) metadata.MD {
// 		md := mtdt.ImportIpToCtx(GetClientIP(r))

// 		authorization := r.Header.Get(Authorization)

// 		if authorization != "" {
// 			bearerToken := strings.Split(authorization, Bearer+" ")
// 			if len(bearerToken) < 2 {
// 				return md
// 			}
// 			token := bearerToken[1]
// 			payload, err := authenticator.Verify(token)
// 			if err != nil {
// 				return md
// 			}
// 			payload.Token = token

// 			md = metadata.Join(md, mtdt.ImportUserInfoToCtx(payload))
// 		}

// 		return md
// 	}
// }

type Response struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
	Data    any      `json:"data"`
}

func ErrorResponse(w http.ResponseWriter, code int, err error) {
	resp := &Response{
		Code:    code,
		Message: err.Error(),
		Details: []string{},
	}

	jData, _ := json.Marshal(resp)

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func DataResponse(w http.ResponseWriter, data any) {
	resp := &Response{
		Data: data,
	}

	jData, _ := json.Marshal(resp)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func forwardErrorResponse(ctx context.Context, s *runtime.ServeMux, m runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	sta := status.Convert(err)
	errStr := sta.Message()
	firstColonPos := strings.Index(errStr, ":")

	if firstColonPos > 0 {
		errStr = errStr[:firstColonPos]
	}

	runtime.DefaultHTTPErrorHandler(ctx, s, m, w, r, errors.New(errStr))
}

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func NewWrapResponseWriter(w http.ResponseWriter) *responseWriterWrapper {
	return &responseWriterWrapper{w, http.StatusOK}
}

func (w *responseWriterWrapper) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}
