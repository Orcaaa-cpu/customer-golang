package models

import (
	"net/http"
	"time"

	"github.com/Orcaaa-cpu/customer-golang/db"
	"github.com/Orcaaa-cpu/customer-golang/entities"
)

func GetCustomer() (Response, error) {
	var res Response
	var obj entities.Customer
	var arrObj []entities.Customer

	con := db.CreateCon()

	sql := "SELECT * FROM customer"

	rows, err := con.Query(sql)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telpon, &obj.Tanggal)
		if err != nil {
			return res, err
		}
		arrObj = append(arrObj, obj)
	}

	// tanggal di bawah ini yyyy-mm-dd
	tanggal, _ := time.Parse("2006-01-02", obj.Tanggal)
	// kita ubah jadi tanggal indo dd-mm-yyyy
	obj.Tanggal = tanggal.Format("02-01-2006")

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

func PostCustomer(nama string, alamat string, telepon string, tanggal string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sql := "Insert customer(nama, alamat, telepon, tanggal) values(?,?,?,?)"

	stmt, err := con.Prepare(sql)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, telepon, tanggal)
	if err != nil {
		return res, err
	}

	rows, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"ID": rows,
	}

	return res, nil
}

func UpdateCustomer(id int, nama string, alamat string, telepon string, tanggal string) (Response, error) {
	var res Response

	con := db.CreateCon()

	stmt, err := con.Prepare("update customer set nama = ?, alamat = ?, telepon = ?, tanggal = ? where id = ?")
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, telepon, tanggal, id)
	if err != nil {
		return res, err
	}

	rows, _ := result.RowsAffected()

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"ID": rows,
	}

	return res, nil
}

func DeleteCustomer(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	stmt, _ := con.Prepare("Delete from customer where id = ?")

	result, _ := stmt.Exec(id)

	rows, _ := result.RowsAffected()

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = rows

	return res, nil
}
