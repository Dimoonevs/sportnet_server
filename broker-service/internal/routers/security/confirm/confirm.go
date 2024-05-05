package confirm

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Dimoonevs/SportsApp/broker-service/pkg"
	"github.com/Dimoonevs/SportsApp/broker-service/proto/security"
	"google.golang.org/grpc"
)

type confirmEmailRequest struct {
	Code string `json:"code"`
	Id   int32  `json:"id"`
}
type resetPasswordRequest struct {
	Email string `json:"email"`
}
type confirmResetPasswordRequest struct {
	Code     string `json:"code"`
	Password string `json:"password"`
	Id       int32  `json:"id"`
}

func ConfirmEmail(w http.ResponseWriter, r *http.Request) {
	var requestPayload confirmEmailRequest
	err := pkg.ReadJSON(w, r, &requestPayload)
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	conn, err := grpc.Dial("security-service:50001", grpc.WithInsecure())
	if err != nil {
		log.Println("Could not connect to gRPC server")
		pkg.ErrorJSON(w, err)
		return
	}
	defer conn.Close()

	client := security.NewSecurityServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	resp, err := client.ConfirmEmail(ctx, &security.ConfirmEmailRequest{
		Code: requestPayload.Code,
		Id:   requestPayload.Id,
	})
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	pkg.WriteJSON(w, http.StatusAccepted, resp)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var requestPayload resetPasswordRequest
	err := pkg.ReadJSON(w, r, &requestPayload)
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	conn, err := grpc.Dial("security-service:50001", grpc.WithInsecure())
	if err != nil {
		log.Println("Could not connect to gRPC server")
		pkg.ErrorJSON(w, err)
		return
	}
	defer conn.Close()

	client := security.NewSecurityServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	resp, err := client.ResetPassword(ctx, &security.ResetPasswordRequest{
		Email: requestPayload.Email,
	})
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	pkg.WriteJSON(w, http.StatusAccepted, resp)
}
func ConfirmResetPassword(w http.ResponseWriter, r *http.Request) {
	var requestPayload confirmResetPasswordRequest
	err := pkg.ReadJSON(w, r, &requestPayload)
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	conn, err := grpc.Dial("security-service:50001", grpc.WithInsecure())
	if err != nil {
		log.Println("Could not connect to gRPC server")
		pkg.ErrorJSON(w, err)
		return
	}
	defer conn.Close()

	client := security.NewSecurityServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	resp, err := client.ConfirmResetPassword(ctx, &security.ConfirmResetPasswordRequest{
		Code:     requestPayload.Code,
		Password: requestPayload.Password,
		Id:       requestPayload.Id,
	})
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	pkg.WriteJSON(w, http.StatusAccepted, resp)
}
