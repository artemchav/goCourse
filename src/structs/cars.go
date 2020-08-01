package structs

type Car struct {
	Firm            string
	Model           string
	Year            int
	Volume          float32
	FilledVolume    byte // % заполненного объема багажника
	IsStarted       bool
	IsOpenedWindows bool
	Sits            int
	Color           string
}

//Допустим, та же машина, но у спорткара должна быть максимальная скорость
type SportCar struct {
	*Car
	MaxSpeed int // km/h
}

//У грузовика - максимальная грузоподъемность и высота кузова, чтоб не сшибать мосты )
type Truck struct {
	*Car
	MaxWeight int
	Height    float32
}
