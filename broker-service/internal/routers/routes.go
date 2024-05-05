package internal

import (
	"net/http"

	"github.com/Dimoonevs/SportsApp/broker-service/internal/routers/athletes"
	"github.com/Dimoonevs/SportsApp/broker-service/internal/routers/groups"
	"github.com/Dimoonevs/SportsApp/broker-service/internal/routers/security"
	"github.com/Dimoonevs/SportsApp/broker-service/internal/routers/security/confirm"
	"github.com/Dimoonevs/SportsApp/broker-service/internal/routers/security/login"
	"github.com/Dimoonevs/SportsApp/broker-service/internal/routers/security/registration"
	"github.com/Dimoonevs/SportsApp/broker-service/internal/routers/security/user"
	"github.com/Dimoonevs/SportsApp/broker-service/internal/routers/subscription"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type AppBroker struct {
}

func (app *AppBroker) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/registrtion", app.RegistrationHeandler)
	mux.Post("/login", app.LoginHeandler)
	mux.Post("/confirm/email", app.ConfirmEmail)
	mux.Post("/reset/password", app.ResetPassword)
	mux.Post("/reset/password/confirm", app.ConfirmResetPassword)
	mux.Group(func(mux chi.Router) {
		mux.Use(security.AuthRequired)
		mux.Post("/sub/create", app.CreateSubscriptionHeandler)
		mux.Post("/group", app.CreateGroup)
		mux.Get("/group", app.GetGroups)
		mux.Put("/group", app.EditGroup)
		mux.Post("/athlete", app.CreateAthletes)
		mux.Delete("/athlete", app.DeleteAthletes)
		mux.Put("/athlete", app.EditAthletes)
		mux.Put("/athlete/training", app.AddTraining)
		mux.Get("/user", app.GetUser)
		mux.Put("/user", app.UpdateUser)
		mux.Put("/user/email", app.ChangeEmail)
		mux.Get("/user/sub", app.GetAllSubscription)
		mux.Put("/user/sub", app.UpdateSubscription)
	})

	return mux
}

func (app *AppBroker) RegistrationHeandler(w http.ResponseWriter, r *http.Request) {
	registration.Registration(w, r)
}

func (app *AppBroker) LoginHeandler(w http.ResponseWriter, r *http.Request) {
	login.Login(w, r)
}

func (app *AppBroker) CreateSubscriptionHeandler(w http.ResponseWriter, r *http.Request) {
	subscription.CreateSubscription(w, r)
}

func (app *AppBroker) CreateGroup(w http.ResponseWriter, r *http.Request) {
	groups.CreateGroup(w, r)
}

func (app *AppBroker) CreateAthletes(w http.ResponseWriter, r *http.Request) {
	athletes.CreateAthletes(w, r)
}

func (app *AppBroker) ConfirmEmail(w http.ResponseWriter, r *http.Request) {
	confirm.ConfirmEmail(w, r)
}

func (app *AppBroker) ResetPassword(w http.ResponseWriter, r *http.Request) {
	confirm.ResetPassword(w, r)
}

func (app *AppBroker) ConfirmResetPassword(w http.ResponseWriter, r *http.Request) {
	confirm.ConfirmResetPassword(w, r)
}
func (app *AppBroker) GetUser(w http.ResponseWriter, r *http.Request) {
	user.GetUser(w, r)
}
func (app *AppBroker) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user.UpdateUser(w, r)
}

func (app *AppBroker) ChangeEmail(w http.ResponseWriter, r *http.Request) {
	user.ChangeEmail(w, r)
}

func (app *AppBroker) GetAllSubscription(w http.ResponseWriter, r *http.Request) {
	subscription.GetAllSubscriptions(w, r)
}

func (app *AppBroker) UpdateSubscription(w http.ResponseWriter, r *http.Request) {
	subscription.EditSubscription(w, r)
}

func (app *AppBroker) GetGroups(w http.ResponseWriter, r *http.Request) {
	groups.GetGroups(w, r)
}

func (app *AppBroker) EditGroup(w http.ResponseWriter, r *http.Request) {
	groups.EditGroup(w, r)
}

func (app *AppBroker) DeleteAthletes(w http.ResponseWriter, r *http.Request) {
	athletes.DeleteAthletes(w, r)
}

func (app *AppBroker) EditAthletes(w http.ResponseWriter, r *http.Request) {
	athletes.EditAthlets(w, r)
}

func (app *AppBroker) AddTraining(w http.ResponseWriter, r *http.Request) {
	athletes.AddTraining(w, r)
}
