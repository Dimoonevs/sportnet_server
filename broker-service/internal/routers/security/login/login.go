package login

import (
	"context"
	"net/http"
	"time"

	"github.com/Dimoonevs/SportsApp/broker-service/pkg"
	"github.com/Dimoonevs/SportsApp/broker-service/proto/security"
	"google.golang.org/grpc"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
type requestPayload struct {
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var requestPayload requestPayload
	err := pkg.ReadJSON(w, r, &requestPayload)
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}

	conn, err := grpc.Dial("security-service:50001", grpc.WithInsecure())
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	defer conn.Close()

	client := security.NewSecurityServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	resp, err := client.Login(ctx, &security.LoginRequest{
		Email:    requestPayload.Email,
		Username: requestPayload.Username,
		Password: requestPayload.Password,
	})
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    resp.Token,
		Expires:  time.Now().Local().Add(time.Hour * time.Duration(24*100)),
		HttpOnly: true,
	})
	var data jsonResponse
	data.Error = false
	data.Message = "Login successful"
	data.Data = resp
	pkg.WriteJSON(w, http.StatusAccepted, data)

	var logPayload pkg.LogPayload
	logPayload.Name = "login-service"
	logPayload.Data = resp.Message
	logPayload.Field = "login"
	pkg.LogInformation(w, logPayload)
}
