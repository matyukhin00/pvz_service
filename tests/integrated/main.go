package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/matyukhin00/pvz_service/internal/model"
)

const (
	url        = "http://localhost:8080"
	token      = "/dummyLogin"
	pvz        = "/pvz"
	receptions = "/receptions"
	products   = "/products"
	close      = "/close_last_reception"
)

type tokenRequest struct {
	Role string `json:"role"`
}

type pvzRequest struct {
	Id   uuid.UUID `json:"id"`
	Date time.Time `json:"registrationDate"`
	City string    `json:"city"`
}

type receptionRequest struct {
	PvzId uuid.UUID `json:"pvzId"`
}

func main() {
	ctx := context.Background()

	tokenRJson, err := json.Marshal(tokenRequest{Role: "moderator"})
	if err != nil {
		log.Fatal(err.Error())
	}

	reqToken, err := http.NewRequestWithContext(ctx, "POST", url+token, bytes.NewBuffer(tokenRJson))
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	reqToken.Header.Set("Content-Type", "application/json")
	reqToken.Header.Set("Accept", "application/json")

	respToken, err := http.DefaultClient.Do(reqToken)
	if err != nil {
		log.Fatal(err.Error())
	}

	tokenModerator := respToken.Header.Get("Authorization")
	respToken.Body.Close()

	pvzJson, err := json.Marshal(pvzRequest{
		Id:   uuid.New(),
		Date: time.Now(),
		City: "Москва",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	reqPvz, err := http.NewRequestWithContext(ctx, "POST", url+pvz, bytes.NewBuffer(pvzJson))
	if err != nil {
		log.Fatal(err.Error())
	}

	reqPvz.Header.Set("Content-Type", "application/json")
	reqPvz.Header.Set("Accept", "application/json")
	reqPvz.Header.Set("Authorization", tokenModerator)

	respPvz, err := http.DefaultClient.Do(reqPvz)
	if err != nil {
		log.Fatal(err.Error())
	}

	pvz := &model.Pvz{}
	err = json.NewDecoder(respPvz.Body).Decode(&pvz)
	if err != nil {
		log.Fatal(err.Error())
	}
	respPvz.Body.Close()

	tokenRJson, err = json.Marshal(tokenRequest{Role: "employee"})
	if err != nil {
		log.Fatal(err.Error())
	}

	reqToken, err = http.NewRequestWithContext(ctx, "POST", url+token, bytes.NewBuffer(tokenRJson))
	if err != nil {
		log.Fatal(err.Error())
	}

	reqToken.Header.Set("Content-Type", "application/json")
	reqToken.Header.Set("Accept", "application/json")

	respToken, err = http.DefaultClient.Do(reqToken)
	if err != nil {
		log.Fatal(err.Error())
	}

	tokenEmployee := respToken.Header.Get("Authorization")
	respToken.Body.Close()

	receptiomJSON, err := json.Marshal(receptionRequest{
		PvzId: pvz.Id,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	reqReception, err := http.NewRequestWithContext(ctx, "POST", url+receptions, bytes.NewBuffer(receptiomJSON))
	if err != nil {
		log.Fatal(err.Error())
	}

	reqReception.Header.Set("Content-Type", "application/json")
	reqReception.Header.Set("Accept", "application/json")
	reqReception.Header.Set("Authorization", tokenEmployee)

	respReception, err := http.DefaultClient.Do(reqReception)
	if err != nil {
		log.Fatal(err.Error())
	}

	reception := &model.Reception{}
	err = json.NewDecoder(respReception.Body).Decode(&reception)
	if err != nil {
		log.Fatal(err.Error())
	}
	respReception.Body.Close()

	types := []string{"электроника", "одежда", "обувь"}
	for i := 0; i < 50; i++ {
		productJson, err := json.Marshal(model.AddProductInc{
			Type:  types[i%3],
			PvzId: pvz.Id,
		})
		if err != nil {
			log.Fatal(err.Error())
		}

		reqProduct, err := http.NewRequestWithContext(ctx, "POST", url+products, bytes.NewBuffer(productJson))
		if err != nil {
			log.Fatal(err.Error())
		}

		reqProduct.Header.Set("Content-Type", "application/json")
		reqProduct.Header.Set("Accept", "application/json")
		reqProduct.Header.Set("Authorization", tokenEmployee)

		respProduct, err := http.DefaultClient.Do(reqProduct)
		if err != nil {
			log.Fatal(err.Error())
		}
		respProduct.Body.Close()
	}

	closeReception, err := http.NewRequestWithContext(ctx, "POST", url+"/"+pvz.Id.String()+close, nil)

	closeReception.Header.Set("Content-Type", "application/json")
	closeReception.Header.Set("Accept", "application/json")
	closeReception.Header.Set("Authorization", tokenEmployee)

	respClose, err := http.DefaultClient.Do(closeReception)
	if err != nil {
		log.Fatal(err.Error())
	}
	respClose.Body.Close()

	log.Println("success integrated test")
}
