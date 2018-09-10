package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	models "github.com/diego-jacobs/GoServerlessTest/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tkanos/gonfig"
)

//Configuration is a representation of the enviroment configuration file
type Configuration struct {
	ConnectionString string
}

func getFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "dev"
	}
	filename := []string{"config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath
}

func getConf() Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf(getFileName(), &configuration)
	if err != nil {
		os.Exit(500)
	}

	return configuration
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getClientByID(ID int) models.Cliente {

func getClientById(ID int) Client {
	db, err := sql.Open("mysql", "")
	checkErr(err)
	// query
	rows, err := db.Query("SELECT * FROM cliente WHERE ID = ?", ID)
	checkErr(err)
	var client models.Cliente
	for rows.Next() {
		temp := models.Cliente{}
		err = rows.Scan(&temp.ID, &temp.Nit, &temp.RazonSocial,
			&temp.NombreComercial, &temp.ServiciosPrestados, &temp.CreadoPorID,
			&temp.ActualizadoPorID, &temp.FechaCreacion, &temp.FechaActualizacion)
		checkErr(err)
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