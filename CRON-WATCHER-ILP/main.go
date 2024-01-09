package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/oxxi/watcher/ilp/models"
	"github.com/oxxi/watcher/ilp/utils"
)

var urls = [125]string{
	"https://app.ilpbuscatalento.com/users/{id}/active-breaks",
	"https://app.ilpbuscatalento.com/awards",
	"https://app.ilpbuscatalento.com/users/{id}/financial-advisories",
	"https://app.ilpbuscatalento.com/user-events/count",
	"https://app.ilpbuscatalento.com/users/{id}/medical-insurances",
	"https://app.ilpbuscatalento.com/vacations/{sap_id}",
	"https://app.ilpbuscatalento.com/awards/count",
	"https://app.ilpbuscatalento.com/coupons/{id}/state",
	"https://app.ilpbuscatalento.com/zones/{id}/work-certificates",
	"https://app.ilpbuscatalento.com/upcoming-events/{id}",
	"https://app.ilpbuscatalento.com/ethical-incidents",
	"https://app.ilpbuscatalento.com/companies/{id}/awards",
	"https://app.ilpbuscatalento.com/insurace-documents/{id}/insurance-type",
	"https://app.ilpbuscatalento.com/coupon-categories/{id}/coupons",
	"https://app.ilpbuscatalento.com/ethical-incidents/{id}",
	"https://app.ilpbuscatalento.com/financial-advisories/{id}",
	"https://app.ilpbuscatalento.com/physical-formats/count",
	"https://app.ilpbuscatalento.com/miscellaneous/{id}",
	"https://app.ilpbuscatalento.com/states",
	"https://app.ilpbuscatalento.com/ethical-incidents/{id}/user",
	"https://app.ilpbuscatalento.com/users/{id}",
	"https://app.ilpbuscatalento.com/notifications/{id}",
	"https://app.ilpbuscatalento.com/companies/{id}/users",
	"https://app.ilpbuscatalento.com/zones/{id}",
	"https://app.ilpbuscatalento.com/cities",
	"https://app.ilpbuscatalento.com/users/{id}/work-certificates",
	"https://app.ilpbuscatalento.com/users/{id}/notifications",
	"https://app.ilpbuscatalento.com/coupons/{id}/coupon-category",
	"https://app.ilpbuscatalento.com/states/{id}",
	"https://app.ilpbuscatalento.com/zones/{id}/financial-advisories",
	"https://app.ilpbuscatalento.com/active-breaks/{id}",
	"https://app.ilpbuscatalento.com/ethical-incidents/count",
	"https://app.ilpbuscatalento.com/insurance-types",
	"https://app.ilpbuscatalento.com/insurance-types/count",
	"https://app.ilpbuscatalento.com/users",
	"https://app.ilpbuscatalento.com/coupons",
	"https://app.ilpbuscatalento.com/cities/count",
	"https://app.ilpbuscatalento.com/insurance-types/{id}/insurace-documents",
	"https://app.ilpbuscatalento.com/users/{id}/ethical-incidents",
	"https://app.ilpbuscatalento.com/insurace-documents/{id}",
	"https://app.ilpbuscatalento.com/insurance-types/{id}",
	"https://app.ilpbuscatalento.com/groups",
	"https://app.ilpbuscatalento.com/miscellaneous/count",
	"https://app.ilpbuscatalento.com/coupon-categories",
	"https://app.ilpbuscatalento.com/groups/{id}/companies",
	"https://app.ilpbuscatalento.com/states/{id}/country",
	"https://app.ilpbuscatalento.com/companies",
	"https://app.ilpbuscatalento.com/verify-transporter",
	"https://app.ilpbuscatalento.com/roles",
	"https://app.ilpbuscatalento.com/financial-advisories/{id}/user",
	"https://app.ilpbuscatalento.com/companies/count",
	"https://app.ilpbuscatalento.com/work-certificates/{id}/zone",
	"https://app.ilpbuscatalento.com/notifications/{id}/user",
	"https://app.ilpbuscatalento.com/users",
	"https://app.ilpbuscatalento.com/zones",
	"https://app.ilpbuscatalento.com/countries/count",
	"https://app.ilpbuscatalento.com/coupon-categories/{id}",
	"https://app.ilpbuscatalento.com/awards/{id}",
	"https://app.ilpbuscatalento.com/financial-advisories",
	"https://app.ilpbuscatalento.com/active-breaks",
	"https://app.ilpbuscatalento.com/notifications",
	"https://app.ilpbuscatalento.com/zones/count",
	"https://app.ilpbuscatalento.com/emergency-lines/count",
	"https://app.ilpbuscatalento.com/upcoming-events/{id}/users",
	"https://app.ilpbuscatalento.com/upcoming-events",
	"https://app.ilpbuscatalento.com/users/{id}/city",
	"https://app.ilpbuscatalento.com/financial-advisories/{id}/zone",
	"https://app.ilpbuscatalento.com/emergency-lines/{id}",
	"https://app.ilpbuscatalento.com/users/{id}/upcoming-events",
	"https://app.ilpbuscatalento.com/payment-stubs/{sap_id}",
	"https://app.ilpbuscatalento.com/groups",
	"https://app.ilpbuscatalento.com/states/{id}/cities",
	"https://app.ilpbuscatalento.com/zones/{id}/user",
	"https://app.ilpbuscatalento.com/countries/{id}/states",
	"https://app.ilpbuscatalento.com/countries",
	"https://app.ilpbuscatalento.com/groups/count",
	"https://app.ilpbuscatalento.com/medical-insurances/{id}",
	"https://app.ilpbuscatalento.com/countries/{id}",
	"https://app.ilpbuscatalento.com/groups/{id}/group-values",
	"https://app.ilpbuscatalento.com/insurace-documents",
	"https://app.ilpbuscatalento.com/insurace-documents/count",
	"https://app.ilpbuscatalento.com/work-certificates",
	"https://app.ilpbuscatalento.com/physical-formats",
	"https://app.ilpbuscatalento.com/awards/{id}/company",
	"https://app.ilpbuscatalento.com/coupon-categories/count",
	"https://app.ilpbuscatalento.com/notifications/count",
	"https://app.ilpbuscatalento.com/work-certificates/count",
	"https://app.ilpbuscatalento.com/users/{id}/company",
	"https://app.ilpbuscatalento.com/active-breaks/count",
	"https://app.ilpbuscatalento.com/states/count",
	"https://app.ilpbuscatalento.com/group-values",
	"https://app.ilpbuscatalento.com/emergency-lines/{id}/group",
	"https://app.ilpbuscatalento.com/physical-formats/{id}",
	"https://app.ilpbuscatalento.com/medical-insurances/{id}/insurance-document",
	"https://app.ilpbuscatalento.com/payment-stub/{sap_id}/{period}",
	"https://app.ilpbuscatalento.com/group-values/{id}",
	"https://app.ilpbuscatalento.com/cities/{id}/state",
	"https://app.ilpbuscatalento.com/coupons/count",
	"https://app.ilpbuscatalento.com/zones/{id}/medical-insurances",
	"https://app.ilpbuscatalento.com/emergency-lines",
	"https://app.ilpbuscatalento.com/insurance-documents/{id}/medical-insurances",
	"https://app.ilpbuscatalento.com/group-values/count",
	"https://app.ilpbuscatalento.com/coupons/{id}",
	"https://app.ilpbuscatalento.com/user-events",
	"https://app.ilpbuscatalento.com/companies/{id}/group",
	"https://app.ilpbuscatalento.com/miscellaneous",
	"https://app.ilpbuscatalento.com/medical-insurances/{id}/user",
	"https://app.ilpbuscatalento.com/roles/count",
	"https://app.ilpbuscatalento.com/users/count",
	"https://app.ilpbuscatalento.com/work-certificates/{id}/user",
	"https://app.ilpbuscatalento.com/user-events/",
	"https://app.ilpbuscatalento.com/medical-insurances",
	"https://app.ilpbuscatalento.com/medical-insurances/{id}/zone",
	"https://app.ilpbuscatalento.com/states/{id}/coupons",
	"https://app.ilpbuscatalento.com/groups/{id}/emergency-lines",
	"https://app.ilpbuscatalento.com/group-values/",
	"https://app.ilpbuscatalento.com/medical-insurances/count",
	"https://app.ilpbuscatalento.com/work-certificates/{id}",
	"https://app.ilpbuscatalento.com/roles/{id}",
	"https://app.ilpbuscatalento.com/companies/{id}",
	"https://app.ilpbuscatalento.com/financial-advisories/count",
	"https://app.ilpbuscatalento.com/upcoming-events/count",
	"https://app.ilpbuscatalento.com/cities/{id}",
	"https://app.ilpbuscatalento.com/sync",         //<- mover luego
	"https://app.ilpbuscatalento.com/users/whoami", //<- mover luego
}

var (
	db *redis.Client
)

var client = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     30 * time.Second,
		DisableKeepAlives:   false,
	},
}

func main() {

	lambda.Start(handler)

}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var loginUrl = "https://app.ilpbuscatalento.com/users/login"
	login, _, err := getLogin(loginUrl)
	if err != nil {
		log.Printf("Error en login %v\n", err)

	}

	info := getResponses()

	info = append(info, login)

	jsonData := utils.ToJson(info)
	getConnection()
	keyTime := time.Now().String()

	_, err = db.HSet("INFO_REQUEST", keyTime, jsonData).Result()
	if err != nil {
		log.Println(err)
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
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

func getLogin(url string) (models.RequestEntity, string, error) {

	login := []byte(`{"userName": "00000000", "password":"@Abced1234"}`)

	bodyReader := bytes.NewReader(login)
	start := time.Now()
	response, err := client.Post(url, "application/json", bodyReader)
	end := time.Since(start)
	if err != nil {
		return models.RequestEntity{}, "", err
	}

	resultRequest := models.RequestEntity{
		UUID:       uuid.NewString(),
		Resource:   url,
		StatusCode: uint64(response.StatusCode),
		Status:     response.Status,
		TimeStart:  start,
		TimeEnd:    time.Now(),
		Duration:   uint64(end.Milliseconds()),
	}

	return resultRequest, "", nil
}

func getResponses() []models.RequestEntity {
	var res []models.RequestEntity
	wg := sync.WaitGroup{}
	for i := 0; i < len(urls)-1; i++ {
		wg.Add(1)
		url := strings.Replace(urls[i], "{id}", "001", 1)
		url = strings.Replace(url, "{sap_id}", "001", 1)
		url = strings.Replace(url, "{period}", "1", 1)

		go func(currentUrl string) {
			start := time.Now()
			resp, err := client.Get(currentUrl)
			end := time.Since(start)
			if err != nil {
				fmt.Println(err)
				return
			}
			var reqInfo models.RequestEntity
			reqInfo.UUID = uuid.NewString()
			reqInfo.Status = resp.Status
			reqInfo.StatusCode = uint64(resp.StatusCode)
			reqInfo.TimeStart = start
			reqInfo.TimeEnd = time.Now()
			reqInfo.Duration = uint64(end.Milliseconds())
			reqInfo.Resource = url
			res = append(res, reqInfo)
			wg.Done()
		}(url)
	}
	wg.Wait()
	return res
}
