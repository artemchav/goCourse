package main

//задачи 1 и 2
import (
	"lession3/task1/structs"
	"log"
)

func main() {
	MyTruck := &structs.Truck{
		Car: &structs.Car{
			Firm:            "Volkswagen",
			Model:           "e-Delivery",
			Year:            2020,
			Volume:          30,
			FilledVolume:    50,
			IsStarted:       false,
			IsOpenedWindows: false,
			Sits:            2,
			Color:           "black",
		},
		Height:    4,
		MaxWeight: 10000,
	}

	log.Printf("%#v", MyTruck)

	MySportCar := &structs.SportCar{
		Car: &structs.Car{
			Firm:   "Ferrary",
			Model:  "125 S",
			Year:   1947,
			Volume: 5,
		},
		MaxSpeed: 210,
	}
	log.Printf("%#v", MySportCar)

	//обращение к полю вложенной структуры, так как своего такого свойства у SportCar нет
	MySportCar.IsOpenedWindows = true

	log.Printf("%v", MySportCar.Car.IsOpenedWindows) //обращение к полю вложенной структуры
	log.Printf("%v", MySportCar.Firm)                //обращение к полю вложенной структуры
}
