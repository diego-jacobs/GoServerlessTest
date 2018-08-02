package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
)

//Client is a representation of a client
type Client struct {
	ID                 int            `json:"id"`
	Nit                string         `json:"nit"`
	RazonSocial        string         `json:"razonSocial"`
	NombreComercial    sql.NullString `json:"nombreComercial"`
	ServiciosPrestados sql.NullString `json:"serviciosPrestados"`
	CreadoPorID        sql.NullInt64  `json:"creadoPorID"`
	ActualizadoPorID   sql.NullInt64  `json:"actualizadoPorID"`
	FechaCreacion      time.Time      `json:"fechaCreacion"`
	FechaActualizacion time.Time      `json:"fechaActualizacion"`
}

//ListClientsResponse is a representation of a list of clients
type ListClientsResponse struct {
	Clients []Client `json:"clients"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getCLients() []Client {
	db, err := sql.Open("mysql", "")
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
