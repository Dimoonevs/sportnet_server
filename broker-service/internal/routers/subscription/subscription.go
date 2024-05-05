package subscription

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Dimoonevs/SportsApp/broker-service/pkg"
	"github.com/Dimoonevs/SportsApp/broker-service/proto/subscription"
	"google.golang.org/grpc"
)

type TypeSub int32

const (
	FIXED_COUNT TypeSub = iota
	TIME_LIMITED
)

type TimeLimited int32

const (
	WEEK TimeLimited = iota
	MONTH
	YEAR
	CUSTOM
)

type StatusSubscription struct {
	TypeSub         TypeSub     `json:"type_sub"`
	TimeLimited     TimeLimited `json:"time_limited"`
	CustomTimeLimit int32       `json:"custom_timeLimit"`
}
type PriceSubscription struct {
	Price    int32  `json:"price"`
	Currency string `json:"currency"`
}

type requestPayload struct {
	Name               string             `json:"name"`
	Discription        string             `json:"discription"`
	StatusSubscription StatusSubscription `json:"status_subscription"`
	PriceSubscription  PriceSubscription  `json:"price_subscription"`
	DaysOfWeek         []string           `json:"days_of_week"`
	UsernameCoach      string             `json:"username_coach"`
	Automatically      bool               `json:"automatically_management"`
	Time               []string           `json:"time"`
}
type requestEditSubscriptionPayload struct {
	Name               string             `json:"name"`
	Discription        string             `json:"discription"`
	StatusSubscription StatusSubscription `json:"status_subscription"`
	PriceSubscription  PriceSubscription  `json:"price_subscription"`
	DaysOfWeek         []string           `json:"days_of_week"`
	UsernameCoach      string             `json:"username_coach"`
	Automatically      bool               `json:"automatically_management"`
	Time               []string           `json:"time"`
	Id                 int32              `json:"id"`
	IdScheduler        int32              `json:"id_scheduler"`
	CronId             int32              `json:"cron_id"`
}

func CreateSubscription(w http.ResponseWriter, r *http.Request) {
	var requestPayload requestPayload

	err := pkg.ReadJSON(w, r, &requestPayload)
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	conn, err := grpc.Dial("typesubscription-service:50002", grpc.WithInsecure())
	if err != nil {
		log.Println("Could not connect to gRPC server")
		pkg.ErrorJSON(w, err)
		return
	}
	defer conn.Close()

	client := subscription.NewTypeSubscriptionServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	id, err := strconv.Atoi(r.Header.Get("id"))
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}

	resp, err := client.CreateSubscription(ctx, &subscription.SubscriptionRequest{
		Name:        requestPayload.Name,
		Description: requestPayload.Discription,
		StatusSubscription: &subscription.StatusSubscription{
			TypeSub:           subscription.TypeSub(requestPayload.StatusSubscription.TypeSub),
			TimeLimited:       subscription.TimeLimited(requestPayload.StatusSubscription.TimeLimited),
			CustomTimeLimited: requestPayload.StatusSubscription.CustomTimeLimit,
		},
		Price: &subscription.PriceSubscription{
			Price:    requestPayload.PriceSubscription.Price,
			Currency: requestPayload.PriceSubscription.Currency,
		},
		DaysOfWeek:              requestPayload.DaysOfWeek,
		CoachId:                 int32(id),
		AutomaticallyManagement: requestPayload.Automatically,
		Time:                    requestPayload.Time,
	})
	if err != nil {
		log.Println("Could not create subscription")
		pkg.ErrorJSON(w, err)
		return
	}
	pkg.WriteJSON(w, http.StatusAccepted, resp)

	var logPayload pkg.LogPayload
	logPayload.Name = "subscription-service"
	logPayload.Data = resp
	logPayload.Field = "subscription"
	pkg.LogInformation(w, logPayload)

}

func GetAllSubscriptions(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial("typesubscription-service:50002", grpc.WithInsecure())
	if err != nil {
		log.Println("Could not connect to gRPC server")
		pkg.ErrorJSON(w, err)
		return
	}
	defer conn.Close()

	client := subscription.NewTypeSubscriptionServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	id, err := strconv.Atoi(r.Header.Get("id"))
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	resp, err := client.GetAllSubscriptions(ctx, &subscription.GetSubscriptionRequest{
		CoachId: int32(id),
	})
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	pkg.WriteJSON(w, http.StatusAccepted, resp)
}
func EditSubscription(w http.ResponseWriter, r *http.Request) {
	var requestPayload requestEditSubscriptionPayload

	err := pkg.ReadJSON(w, r, &requestPayload)
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	conn, err := grpc.Dial("typesubscription-service:50002", grpc.WithInsecure())
	if err != nil {
		log.Println("Could not connect to gRPC server")
		pkg.ErrorJSON(w, err)
		return
	}
	defer conn.Close()

	client := subscription.NewTypeSubscriptionServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	id, err := strconv.Atoi(r.Header.Get("id"))
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	resp, err := client.EditSubscription(ctx, &subscription.SubscriptionEditRequest{
		Name:        requestPayload.Name,
		Description: requestPayload.Discription,
		StatusSubscription: &subscription.StatusSubscription{
			TypeSub:           subscription.TypeSub(requestPayload.StatusSubscription.TypeSub),
			TimeLimited:       subscription.TimeLimited(requestPayload.StatusSubscription.TimeLimited),
			CustomTimeLimited: requestPayload.StatusSubscription.CustomTimeLimit,
		},
		Price: &subscription.PriceSubscription{
			Price:    requestPayload.PriceSubscription.Price,
			Currency: requestPayload.PriceSubscription.Currency,
		},
		DaysOfWeek:              requestPayload.DaysOfWeek,
		CoachId:                 int32(id),
		AutomaticallyManagement: requestPayload.Automatically,
		Time:                    requestPayload.Time,
		Id:                      requestPayload.Id,
		IdScheduler:             requestPayload.IdScheduler,
		CronId:                  requestPayload.CronId,
	})
	if err != nil {
		log.Println("Broker Error")
		pkg.ErrorJSON(w, err)
		return
	}
	pkg.WriteJSON(w, http.StatusAccepted, resp)
}
