package main

import (
	"GoServerlessTest/beetest/models"
	"GoServerlessTest/utils"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
)

// HandleGetAllClientsRequest is the Handler to get the list of Clients
func HandleGetAllClientsRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	utils.InitORM()

	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	list, err := models.GetAllCliente(query, fields, sortby, order, offset, limit)

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
		Data:    list,
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
	lambda.Start(HandleGetAllClientsRequest)
}
