package main

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getCLients() []Client {
	configuration = getConf()

	db, err := sql.Open("mysql", configuration.ConnectionString)
	checkErr(err)

	// query
	rows, err := db.Query("SELECT * FROM cliente")
	checkErr(err)

	var clients []Client

	for rows.Next() {
		client := Client{}

		err = rows.Scan(&client.ID, &client.Nit, &client.RazonSocial,
			&client.NombreComercial, &client.ServiciosPrestados, &client.CreadoPorID,
			&client.ActualizadoPorID, &client.FechaCreacion, &client.FechaActualizacion)

		checkErr(err)
		clients = append(clients, client)
	}

	db.Close()

	return clients
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var clients = getCLients()

	body, _ := json.Marshal(&ClientList{
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
