package athletes

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Dimoonevs/SportsApp/broker-service/pkg"
	"github.com/Dimoonevs/SportsApp/broker-service/proto/athletes"
	"google.golang.org/grpc"
)

type requestPayload struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	SubscriptionId int32  `json:"subscription_id"`
	GroupId        int32  `json:"group_id"`
	DaysLeft       int32  `json:"days_left"`
	DateLast       string `json:"date_last"`
	Id             int32  `json:"id"`
}
type deleteRequestPayload struct {
	Id []int32 `json:"id"`
}
type editRequestPayload struct {
	Atletes []*requestPayload `json:"athletes"`
}
type trainingData struct {
	Id       int32  `json:"id"`
	DaysLeft int32  `json:"days_left"`
	DateLast string `json:"date_last"`
}
type addTrainingRequestPayload struct {
	TrainingData []*trainingData `json:"training"`
}

func CreateAthletes(w http.ResponseWriter, r *http.Request) {
	var payload requestPayload
	err := pkg.ReadJSON(w, r, &payload)
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	conn, err := grpc.Dial("athletes-service:50004", grpc.WithInsecure())
	if err != nil {
		log.Println("Could not connect to gRPC server")
		pkg.ErrorJSON(w, err)
		return
	}

	defer conn.Close()

	client := athletes.NewAthleteServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	resp, err := client.CreateAthlete(ctx, &athletes.AthleteRequest{
		FirstName:      payload.FirstName,
		LastName:       payload.LastName,
		SubscriptionId: payload.SubscriptionId,
		GroupId:        payload.GroupId,
		DaysLeft:       payload.DaysLeft,
		DateLast:       payload.DateLast,
	})
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	pkg.WriteJSON(w, http.StatusAccepted, resp.Message, nil)
}

func DeleteAthletes(w http.ResponseWriter, r *http.Request) {
	var payload deleteRequestPayload
	err := pkg.ReadJSON(w, r, &payload)
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	conn, err := grpc.Dial("athletes-service:50004", grpc.WithInsecure())
	if err != nil {
		log.Println("Could not connect to gRPC server")
		pkg.ErrorJSON(w, err)
		return
	}

	defer conn.Close()

	client := athletes.NewAthleteServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	resp, err := client.DeleteAthletes(ctx, &athletes.DeleteAthletesRequest{
		Id: payload.Id,
	})
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	pkg.WriteJSON(w, http.StatusAccepted, resp.Message, nil)
}

func EditAthlets(w http.ResponseWriter, r *http.Request) {
	var payload editRequestPayload
	err := pkg.ReadJSON(w, r, &payload)
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	conn, err := grpc.Dial("athletes-service:50004", grpc.WithInsecure())
	if err != nil {
		log.Println("Could not connect to gRPC server")
		pkg.ErrorJSON(w, err)
		return
	}

	defer conn.Close()

	client := athletes.NewAthleteServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	var athletesArr []*athletes.AthleteRequest

	for _, athlete := range payload.Atletes {
		athletesArr = append(athletesArr, &athletes.AthleteRequest{
			FirstName:      athlete.FirstName,
			LastName:       athlete.LastName,
			SubscriptionId: athlete.SubscriptionId,
			GroupId:        athlete.GroupId,
			DaysLeft:       athlete.DaysLeft,
			DateLast:       athlete.DateLast,
			Id:             athlete.Id,
		})
	}

	resp, err := client.EditAthlete(ctx, &athletes.EditAthletesRequest{
		Athlete: athletesArr,
	})
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	pkg.WriteJSON(w, http.StatusAccepted, resp.Message, nil)

}

func AddTraining(w http.ResponseWriter, r *http.Request) {
	var payload addTrainingRequestPayload
	err := pkg.ReadJSON(w, r, &payload)
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	conn, err := grpc.Dial("athletes-service:50004", grpc.WithInsecure())
	if err != nil {
		log.Println("Could not connect to gRPC server")
		pkg.ErrorJSON(w, err)
		return
	}

	defer conn.Close()

	client := athletes.NewAthleteServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	var trainingArr []*athletes.TrainingData
	for _, training := range payload.TrainingData {
		trainingArr = append(trainingArr, &athletes.TrainingData{
			Id:       training.Id,
			DaysLeft: training.DaysLeft,
			DateLast: training.DateLast,
		})
	}
	resp, err := client.AddTraining(ctx, &athletes.AddTrainingRequest{
		Training: trainingArr,
	})
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	pkg.WriteJSON(w, http.StatusAccepted, resp.Message, nil)
}
