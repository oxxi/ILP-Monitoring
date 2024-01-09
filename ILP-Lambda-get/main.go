package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-redis/redis"
)

type Response struct {
	UUID       string    `json:"uuid"`
	Resource   string    `json:"resource"`
	StatusCode uint64    `json:"statusCode"`
	Status     string    `json:"status"`
	TimeStart  time.Time `json:"timeStart"`
	TimeEnd    time.Time `json:"timeEnd"`
	Duration   uint64    `json:"duration"`
}

var (
	db *redis.Client
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	getConnection()
	value, err := db.HGetAll("INFO_REQUEST").Result()
	if err != nil {
		panic(err)
	}
	var resp []Response
	for _, v := range value {
		tmp := []Response{}
		if err := json.Unmarshal([]byte(v), &tmp); err != nil {
			log.Fatal(err)
		}
		resp = append(resp, tmp...)
	}

	jsonResult, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonResult),
	}

	return response, nil

}

func getConnection() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "us1-fast-griffon-40810.upstash.io:40810",
		Password: "d2a3c23fff55456ca7c6d13c6cb4071a", // no password set
		DB:       0,                                  // use default DB
	})

	db = rdb
}
