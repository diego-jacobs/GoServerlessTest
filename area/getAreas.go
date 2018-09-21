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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//ListAreasResponse is a representation of a list of areas
type ListAreasResponse struct {
	Areas []models.Area `json:"areas"`
}

func getAreas() []models.Area {
	var connectionString = configuration.GetConnectionString()

	db, err := sql.Open("mysql", connectionString)
	checkErr(err)
	// query
	rows, err := db.Query("SELECT * FROM area")
	checkErr(err)
	var areas []models.Area
	for rows.Next() {
		area := models.Area{}
		err = rows.Scan(&area.ID, &area.FechaCreacion, &area.FechaActualizacion, &area.DepartamentoID, &area.CreadoPorID, &area.ActualizadoPorID)
		checkErr(err)
		areas = append(areas, area)
	}
	db.Close()
	return areas
}

// HandleRequest is the Handler to get the list of Clients
func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var areas = getAreas()
	body, _ := json.Marshal(&ListAreasResponse{
		Areas: areas,
	})

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
