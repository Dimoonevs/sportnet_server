package registration

import (
	"context"
	"log"
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
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
	DateOfBirth string `json:"date_of_birth"`
	Status      Status `json:"status"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	TimeZone    string `json:"time_zone"`
}
type Status int32

const (
	Status_SPORTSMEN           Status = 0
	Status_COACH               Status = 1
	Status_SPORTSMEN_AND_COACH Status = 3
)

func Registration(w http.ResponseWriter, r *http.Request) {
	var requestPayload requestPayload
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

	resp, err := client.Register(ctx, &security.RegisterRequest{
		FirstName:   requestPayload.FirstName,
		LastName:    requestPayload.LastName,
		Email:       requestPayload.Email,
		Username:    requestPayload.Username,
		Password:    requestPayload.Password,
		DateOfBirth: requestPayload.DateOfBirth,
		Status:      security.Status(1),
		TimeZone:    requestPayload.TimeZone,
	})
	if err != nil {
		pkg.ErrorJSON(w, err)
		var logPayload pkg.LogPayload
		logPayload.Name = "registration-service"
		logPayload.Data = err.Error()
		logPayload.Field = "registration-error"
		pkg.LogInformation(w, logPayload)
		return
	}
	var response jsonResponse

	response.Error = false
	response.Message = resp.Message
	response.Data = resp.Id
	pkg.WriteJSON(w, http.StatusAccepted, response)

	var logPayload pkg.LogPayload
	logPayload.Name = "registration-service"
	logPayload.Data = response.Data
	logPayload.Field = "registration"
	pkg.LogInformation(w, logPayload)
}
