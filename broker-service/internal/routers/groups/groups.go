package groups

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Dimoonevs/SportsApp/broker-service/pkg"
	"github.com/Dimoonevs/SportsApp/broker-service/proto/groups"
	"google.golang.org/grpc"
)

type requestPayload struct {
	Id            int32  `json:"id"`
	Name          string `json:"name"`
	SubscriptonId int32  `json:"subscripton_id"`
}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	var payload requestPayload
	err := pkg.ReadJSON(w, r, &payload)
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}

	conn, err := grpc.Dial("groups-service:50003", grpc.WithInsecure())
	if err != nil {
		log.Println("Could not connect to gRPC server")
		pkg.ErrorJSON(w, err)
		return
	}
	defer conn.Close()

	client := groups.NewGroupServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	id, err := strconv.Atoi(r.Header.Get("id"))
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}

	resp, err := client.CreateGroup(ctx, &groups.GroupRequest{
		Name:           payload.Name,
		CoachId:        int32(id),
		SubscriptionId: payload.SubscriptonId,
	})

	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	pkg.WriteJSON(w, http.StatusAccepted, resp)
}

func GetGroups(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial("groups-service:50003", grpc.WithInsecure())
	if err != nil {
		log.Println("Could not connect to gRPC server")
		pkg.ErrorJSON(w, err)
		return
	}
	defer conn.Close()

	client := groups.NewGroupServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	id, err := strconv.Atoi(r.Header.Get("id"))
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	resp, err := client.GetGroups(ctx, &groups.GetGroupRequest{
		CoachId: int32(id),
	})
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	pkg.WriteJSON(w, http.StatusAccepted, resp)
}

func EditGroup(w http.ResponseWriter, r *http.Request) {
	var payload requestPayload
	err := pkg.ReadJSON(w, r, &payload)
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	conn, err := grpc.Dial("groups-service:50003", grpc.WithInsecure())
	if err != nil {
		log.Println("Could not connect to gRPC server")
		pkg.ErrorJSON(w, err)
		return
	}
	defer conn.Close()

	client := groups.NewGroupServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()
	log.Println(payload)

	resp, err := client.EditGroup(ctx, &groups.GroupEditRequest{
		Id:             payload.Id,
		Name:           payload.Name,
		SubscriptionId: payload.SubscriptonId,
	})
	if err != nil {
		pkg.ErrorJSON(w, err)
		return
	}
	pkg.WriteJSON(w, http.StatusAccepted, resp)
}
