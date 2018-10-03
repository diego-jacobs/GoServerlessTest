package main

import (
	"GoServerlessTest/beetest/models"
	"GoServerlessTest/utils"
	"context"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
)

// HandleRequest is the Handler to get a Client filtered by ID
func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var ID, err = strconv.Atoi(request.PathParameters["id"])
	utils.InitORM()
	client, err := models.GetClienteById(ID)

	if err != nil {
		response := utils.APIResponse{
			Data:    nil,
			Success: false,
			Message: "",
			Error:   err,
		}
		body, _ := json.Marshal(response)

		return events.APIGatewayProxyResponse{
			Body:       string(body),
			StatusCode: 400,
		}, nil
	}

	response := utils.APIResponse{
		Data:    client,
		Success: true,
		Message: "",
		Error:   err,
	}
	body, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
