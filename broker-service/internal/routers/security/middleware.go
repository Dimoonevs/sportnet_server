package security

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Dimoonevs/SportsApp/broker-service/pkg"
	"github.com/Dimoonevs/SportsApp/broker-service/proto/security"
	"google.golang.org/grpc"
)

func AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		if authorization == "" {
			pkg.WriteJSON(w, http.StatusUnauthorized, map[string]interface{}{"error": "Unauthorized"})
			return
		}

		token := strings.Split(authorization, "Bearer ")
		if len(token) < 2 {
			pkg.WriteJSON(w, http.StatusUnauthorized, map[string]interface{}{"error": "Unauthorized"})
			return
		}

		conn, err := grpc.Dial("security-service:50001", grpc.WithInsecure())
		if err != nil {
			log.Println("Could not connect to gRPC server")
			pkg.WriteJSON(w, http.StatusUnauthorized, map[string]interface{}{"error": "Unauthorized"})
			return
		}
		defer conn.Close()

		client := security.NewSecurityServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		resp, err := client.Validate(ctx, &security.ValidateRequest{
			Token: token[1],
		})
		if err != nil {
			pkg.WriteJSON(w, http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
			return
		}
		r.Header.Set("id", fmt.Sprintf("%d", resp.GetId()))
		next.ServeHTTP(w, r)
	})
}
