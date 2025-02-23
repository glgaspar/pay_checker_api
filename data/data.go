package data

import (
	"time"

	"github.com/glgaspar/pay_checker_api/models"
)

func CreateBill(bill *models.Bill) (*models.Bill, error) {
	conn, err := db()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := `
	insert into bills (description,expDay,path,track)
	values ($1, $2, $3, true);`
	_, err = conn.Query(query, bill.Description, bill.ExpDay, bill.Path)
	if err != nil {
		return nil, err
	}

	return bill, nil
}

func UpdateBill(bill *models.Bill) (*models.Bill, error) {
	conn, err := db()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := `
	update bills
	set 
		description = $1
		,expDay = $2
		,path = $3
		,track = $4
	where 
		id = $5 
		`
	_, err = conn.Query(query, bill.Description, bill.ExpDay, bill.Path, bill.Track, bill.Id)
	if err != nil {
		return nil, err
	}

	return bill, nil
}

func PayBill(bill *models.Bill) (*models.Bill, error) {
	conn, err := db()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := `
	update bills
	set
		lastDate $1
	where 
		id = $2
	`
	_, err = conn.Query(query, time.Now(), bill.Id)
	if err != nil {
		return nil, err
	}

	return bill, nil
}

func GetList() (*[]models.Bill, error) {
	conn, err := db()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	var bill []models.Bill
	query := `
	select 
		id, description, expDay, lastDate, path, track
	from bills
	`
	result, err := conn.Query(query)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		var b models.Bill
		result.Scan(&b.Id, &b.Description, &b.ExpDay, &b.LastDate, &b.Path, &b.Track)
		bill = append(bill, b)
	}

	return &bill, nil
}
