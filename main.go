package main

import (
	"encoding/json"
	"log"
	"vslink-linkcategories-lambda/data"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Going to load list of link categories")
	repository := data.NewLinkCategoryRepository()
	categories, err := repository.GetAllActive()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, nil
	}

	json, _ := json.Marshal(categories)

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(json),
	}

	return response, nil
}
