package main

import (
	configuration "GoServerlessTest/config"
	models "GoServerlessTest/models"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
)

//ListClientsResponse is a representation of a list of clients
type ListClientsResponse struct {
	Clients []models.Cliente `json:"clients"`
}

func getClients() []models.Cliente {
	var connectionString = configuration.GetConnectionString()

	db, err := sql.Open("mysql", connectionString)
	checkErr(err)
	// query
	rows, err := db.Query("SELECT * FROM cliente")
	checkErr(err)
	var clients []models.Cliente
	for rows.Next() {
		client := models.Cliente{}
		err = rows.Scan(&client.ID, &client.Nit, &client.RazonSocial,
			&client.NombreComercial, &client.ServiciosPrestados, &client.CreadoPorID,
			&client.ActualizadoPorID, &client.FechaCreacion, &client.FechaActualizacion)
		checkErr(err)
		clients = append(clients, client)
	}
	db.Close()
	return clients
}

// HandleRequest is the Handler to get the list of Clients
func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var clients = getClients()
	body, _ := json.Marshal(&ListClientsResponse{
		Clients: clients,
	})

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
