package utils

import (
	configuration "GoServerlessTest/config"

	"github.com/astaxie/beego/orm"
)

// CheckErr function to check errors ...CheckErr
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// InitORM function is going to configure the ORM connection
func InitORM() {
	var connectionString = configuration.GetConnectionString()
	orm.RegisterDataBase("default", "mysql", connectionString)
}
