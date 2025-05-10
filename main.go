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
	fmt.Println("Тестируем базовую машину:")
	patterns.TestDrive(car)

	fmt.Println("\nТестируем машину с турбонаддувом:")
	patterns.TestDrive(turboCar)

	fmt.Println("\nТестируем машину с турбонаддувом и спортивным выхлопом:")
	patterns.TestDrive(sportsCar)

	// adapter
	tesla := &patterns.ElectricCar{}
	adapter := &patterns.ElectricCarAdapter{ECar: tesla}

	patterns.Travel(adapter)
}
