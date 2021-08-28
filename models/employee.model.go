package models

import (
	"github.com/satriyoaji/echo-restful-api/database"
	"github.com/satriyoaji/echo-restful-api/helper"
	"net/http"
)

type Employee struct {
	Id      int    `json:"id"`
	Name    string `json:"nama" validate:"required"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone"`
}

func FetchAllEmployees() (Response, error) {
	var obj Employee
	var arrayObj []Employee
	var res Response

	con := database.GetConnection()

	sqlStatement := "SELECT * FROM employees"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	helper.OutputPanicError(err)

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Address, &obj.Phone)
		if err != nil {
			return res, err
		}

		arrayObj = append(arrayObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrayObj

	return res, nil
}
