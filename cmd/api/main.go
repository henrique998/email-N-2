package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	campaing "github.com/henrique998/email-N/internal/domain/campaign"
	"github.com/henrique998/email-N/internal/endpoints"
	"github.com/henrique998/email-N/internal/infra/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.NewDb()
	campaignService := campaing.ServiceImp{
		Repo: &database.CampaignRepository{Db: db},
	}
	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}

	r.Route("/campaigns", func(r chi.Router) {
		r.Use(endpoints.Auth)
		r.Post("/", endpoints.HandlerError(handler.CampaignPost))
		r.Get("/{id}", endpoints.HandlerError(handler.CampaignFindById))
		r.Delete("/delete/{id}", endpoints.HandlerError(handler.CampaignsDelete))
	})

	http.ListenAndServe(":3333", r)
}
