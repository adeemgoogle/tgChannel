package patterns

import "fmt"

// Интерфейс машины
type Car interface {
	Drive()
}

// наша тачка
type BasicCar struct{}

func (b *BasicCar) Drive() {
	fmt.Println("Машина едет на обычных оборотах.")
}

// Декоратор для улучшения мощности
type TurboDecorator struct {
	Car Car
}

func (t *TurboDecorator) Drive() {
	fmt.Println("Турбо-режим активирован!")
	t.Car.Drive() // вызываем оригинальную функцию
}

// Декоратор для добавления спортивного выхлопа
type SportsExhaustDecorator struct {
	Car Car
}

func (s *SportsExhaustDecorator) Drive() {
	fmt.Println("Спортивный выхлоп активирован!")
	s.Car.Drive() // вызываем оригинальную функцию
}

// Функция для теста
func TestDrive(c Car) {
	c.Drive()
}
