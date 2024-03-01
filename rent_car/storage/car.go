package storage

import (
	"database/sql"
	"fmt"
	"rent_car/models"

	"github.com/google/uuid"
)

type carRepo struct {
	db *sql.DB
}

func NewCar(db *sql.DB) carRepo {
	return carRepo{
		db: db,
	}
}

/*
create (body) id,err
update (body) id,err
delete (id) err
get (id) body,err
getAll (search) []body,count,err
*/

func (c *carRepo) Create(car models.Car) (string, error) {

	id := uuid.New()

	query := ` INSERT INTO cars (
		id,
		name,
		brand,
		year,
		model,
		hourse_power,
		colour,
		engine_cap)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8) 
	`

	res, err := c.db.Exec(query,
		id.String(),
		car.Name, car.Brand, car.Year,
		car.Model, car.HoursePower,
		car.Colour, car.EngineCap)

	if err != nil {
		return "", err
	}

	fmt.Printf("%+v\n", res)

	return id.String(), nil
}

func (c *carRepo) Update(car models.Car) error {
	_, err := c.db.Exec(
		`UPDATE cars 
		SET name=$1,
			brand=$2,
			year=$3,
			model=$4,
			hourse_power=$5,
			colour=$6,
			engine_cap=$7,
			updated_at=NOW()
			WHERE id=$8`, car.Name, car.Brand, car.Year, car.Model, car.HoursePower, car.Colour, car.EngineCap, car.Id)
	if err != nil {
		fmt.Println("error while updating car err: ", err)
		return err
	}

	return nil
}

func (i *carRepo) Delete(c models.Car, id string) error {
	_, err := i.db.Exec(
		`DELETE FROM cars
			WHERE id=$1`, id)
	if err != nil {
		fmt.Println("error while deleting country err: ", err)
		return err
	}

	return nil
}

func (c *carRepo) GetAll() ([]models.Car, error) {
	cars := []models.Car{}
	rows, err := c.db.Query(`SELECT 
		id,
		name,
		brand,
		colour,
		model,
		engine_cap,
		hourse_power,
		year,
		created_at 
	FROM 
		cars`)
	if err != nil {
		fmt.Println("Error while getting all cars:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var car models.Car
		if err := rows.Scan(&car.Id, &car.Name, &car.Brand, &car.Colour, &car.Model, &car.EngineCap, &car.HoursePower, &car.Year, &car.CreatedAt); err != nil {
			fmt.Println("Error while scanning cars:", err)
			return nil, err
		}
		cars = append(cars, car)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error after iterating rows:", err)
		return nil, err
	}

	fmt.Println("Cars:", cars)
	return cars, nil
}


