package main

import (
	configuration "GoServerlessTest/config"
	models "GoServerlessTest/models"
	"context"
	"database/sql"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
)

func getClientByID(ID int) models.Cliente {
	var connectionString = configuration.GetConnectionString()

	db, err := sql.Open("mysql", connectionString)
	// query
	rows, err := db.Query("SELECT * FROM cliente WHERE ID = ?", ID)
	configuration.CheckErr(err)
	var client models.Cliente
	for rows.Next() {
		temp := models.Cliente{}
		err = rows.Scan(&temp.ID, &temp.Nit, &temp.RazonSocial,
			&temp.NombreComercial, &temp.ServiciosPrestados, &temp.CreadoPorID,
			&temp.ActualizadoPorID, &temp.FechaCreacion, &temp.FechaActualizacion)
		configuration.CheckErr(err)
		client = temp
	}
	db.Close()
	return client
}

// HandleRequest is the Handler to get a Client filtered by ID
func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var ID, err = strconv.Atoi(request.PathParameters["id"])
	var client = getClientByID(ID)

	body, _ := json.Marshal(client)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
