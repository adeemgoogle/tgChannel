package patterns

import "fmt"

type CarA interface {
	Drive()
}

type ElectricCar struct{}

func (e *ElectricCar) StartSilentRide() {
	fmt.Println("Едем на Тесле")
}

// Адаптер, который приводит ElectricCar к интерфейсу Car
type ElectricCarAdapter struct {
	ECar *ElectricCar
}

func (a *ElectricCarAdapter) Drive() {
	a.ECar.StartSilentRide()
}

func Travel(c CarA) {
	fmt.Println("\nТут мы на  Альфараби шашкуем с BMW")
	c.Drive()
}
