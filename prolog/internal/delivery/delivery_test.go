package delivery

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/ndnhuy/proglog/internal/utils"
)

func Init() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "olah123",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:6606",
		DBName:               "delivery",
		AllowNativePasswords: true,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	// delete data
	sqls := []string{
		"DELETE FROM delivery_order",
		"alter table delivery_order AUTO_INCREMENT = 1",

		"DELETE FROM delivery_item",
		"alter table delivery_item AUTO_INCREMENT = 1",
	}
	for _, sql := range sqls {
		_, err = db.Exec(sql)
		if err != nil {
			panic(err)
		}
	}
}

func TestCreateDeliveryOrder(t *testing.T) {
	Init()
	createDeliveryOrder()
}

func TestCreateTrips(t *testing.T) {
	Init()
	createDeliveryOrder()
	createTrips()
}

func createTrips() {
	trip := newTrip()
	_, err := trip.Create()
	if err != nil {
		panic(err)
	}
	// find all DO
	do := newDeliveryOrder("", nil)
	deliveryOrders, err := do.FindAll()
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("%+v\n", utils.Map(deliveryOrders, func(t *DeliveryOrder) DeliveryOrder {
		return *t
	}))
	// add stop
}

func createDeliveryOrder() {
	newDOs := []*DeliveryOrder{
		newDeliveryOrder("SO-001", []*DeliveryItem{
			newDeliveryItem("SKU-A", 100),
			newDeliveryItem("SKU-B", 50),
		}),
		newDeliveryOrder("SO-002", []*DeliveryItem{
			newDeliveryItem("SKU-C", 100),
		}),
	}
	for _, newDO := range newDOs {
		_, err := newDO.Create()
		if err != nil {
			panic(err)
		}
	}
}
