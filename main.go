package main

import (
	"fmt"
	"tgChannel/patterns"
)

func main() {
	// Создаём базовую машину
	car := &patterns.BasicCar{}

	// Создаём декораторы для улучшений
	turboCar := &patterns.TurboDecorator{Car: car}
	sportsCar := &patterns.SportsExhaustDecorator{Car: car}

	// Тестируем машину с декораторами
	fmt.Println("Test decorator")
	fmt.Println("\nТестируем базовую машину:")
	patterns.TestDrive(car)

	fmt.Println("\nТестируем машину с турбонаддувом:")
	patterns.TestDrive(turboCar)

	fmt.Println("\nТестируем машину с турбонаддувом и спортивным выхлопом jjjjjj:")
	patterns.TestDrive(sportsCar)

	// adapter
	tesla := &patterns.ElectricCar{}
	adapter := &patterns.ElectricCarAdapter{ECar: tesla}
	fmt.Println("\ntest adapter")
	fmt.Println("want to test")
	patterns.Travel(adapter)
}
