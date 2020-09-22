package middlewares

import (
	"log"
	"net/http"
)

func CorrelationIDMiddleware(next http.Handler) http.Handler  {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		correlationId := req.Header.Get("correlationId")

		log.Println("Incoming correlationId", correlationId)
		next.ServeHTTP(res, req)
		log.Println("Out coming correlationId", correlationId)
	})
}
