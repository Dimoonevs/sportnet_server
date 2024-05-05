package user

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Dimoonevs/SportsApp/broker-service/pkg"
	"github.com/Dimoonevs/SportsApp/broker-service/proto/security"
	"google.golang.org/grpc"
)

type requestUpdatePayload struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
	DateOfBirth string `json:"date_of_birth"`
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type requestChangeEmailPayload struct {
	Email string `json:"email"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {

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

	id, err := strconv.Atoi(r.Header.Get("id"))
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}

	resp, err := client.GetCoach(ctx, &security.GetCoachRequest{Id: int32(id)})
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	var response jsonResponse

	response.Error = false
	response.Message = "success"
	response.Data = resp
	pkg.WriteJSON(w, http.StatusOK, response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var requestPayload requestUpdatePayload
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

	id, err := strconv.Atoi(r.Header.Get("id"))
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}

	resp, err := client.UpdateCoach(ctx, &security.CoachData{
		Id:          int32(id),
		FirstName:   requestPayload.FirstName,
		LastName:    requestPayload.LastName,
		Username:    requestPayload.Username,
		DateOfBirth: requestPayload.DateOfBirth,
	})
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	var response jsonResponse

	response.Error = false
	response.Message = "success"
	response.Data = resp.Token
	pkg.WriteJSON(w, http.StatusOK, response)
}
func ChangeEmail(w http.ResponseWriter, r *http.Request) {
	var requestPayload requestChangeEmailPayload
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

	id, err := strconv.Atoi(r.Header.Get("id"))
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}

	resp, err := client.ChangeEmail(ctx, &security.ChangeEmailReq{
		Id:    int32(id),
		Email: requestPayload.Email,
	})
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}

	var response jsonResponse
	response.Error = false
	response.Message = "success"
	response.Data = resp.Id
	pkg.WriteJSON(w, http.StatusOK, response)
}
