package labs

import (
	"bufio"
	"log"

	"github.com/kamilrahmatullin/lab/utils"
	"math/rand"
)

func task1(logger *log.Logger) {
	logger.Println("Камиль")
	logger.Println("Москва")
}

func task2(logger *log.Logger, reader *bufio.Reader) {
	var a, b int

	logger.Println("Введите первое число: ")
	a, err := utils.ReadInt(reader)
	if err != nil {
		logger.Println(err)
		return
	}

	logger.Println("Введите второе число: ")
	b, err = utils.ReadInt(reader)
	if err != nil {
		logger.Println(err)
		return
	}

	logger.Println("Сумма чисел:", a+b)
}

func task3(logger *log.Logger) {
	var (
		a, b int = 15, 25
	)

	logger.Printf("Число a = %d\nЧисло b = %d\n", a, b)

	utils.ExchangeNumbers(&a, &b)

	logger.Printf("Новое число a = %d\nНовое число b = %d\n", a, b)
}

func task4(logger *log.Logger, reader *bufio.Reader) {
	logger.Print("Введите длину: ")
	length, err := utils.ReadInt(reader)
	if err != nil {
		logger.Println(err)
		return
	}

	logger.Print("Введите ширину: ")
	width, err := utils.ReadInt(reader)
	if err != nil {
		logger.Println(err)
		return
	}

	logger.Println("Площадь прямоугольника:", length*width)
}

func task5(logger *log.Logger, reader *bufio.Reader) {
	rate := 90.5

	logger.Print("Введите сумму в рублях: ")
	sum, err := utils.ReadInt(reader)
	if err != nil {
		logger.Println(err)
		return
	}

	logger.Printf("%d рублей = %.2f долларов", sum, float64(sum)/rate)
}

func task6(logger *log.Logger, reader *bufio.Reader) {
	logger.Print("Введите число: ")
	num, err := utils.ReadInt(reader)
	if err != nil {
		logger.Println(err)
		return
	}

	logger.Printf("В числе %d: десятков - %d, единиц - %d", num, num/10, num%10)
}

func task7(logger *log.Logger) {
	x := 5
	x++
	logger.Println(x + x)
}

func task8(logger *log.Logger, reader *bufio.Reader) {
	logger.Print("Введите число: ")
	num, err := utils.ReadInt(reader)
	if err != nil {
		logger.Println(err)
		return
	}

	if num%2 == 0 {
		logger.Printf("%d - четное? - Да.\n", num)
	} else {
		logger.Printf("%d - четное? - Нет.\n", num)
	}

	if num%5 == 0 {
		logger.Println("Делится на 5? - Да.")
	} else {
		logger.Println("Делится на 5? - Нет.")
	}

	if num >= 10 && num <= 50 {
		logger.Println("Принадлежит диапазону [10, 50]? Да.")
	} else {
		logger.Println("Принадлежит диапазону [10, 50]? Нет.")
	}
}

func task9(logger *log.Logger) {
	var (
		a, b, c bool = true, false, true
	)

	logger.Println(((a && !b) || (c && !a)) && (b || !c))
}

func task10(logger *log.Logger) {
	num := rand.Int() % 100

	logger.Println("Выбрано число", num)

	if num >= 10 && num <= 99 {
		logger.Println("Число двузначное")
	} else {
		logger.Println("Число недвузначное")
	}

	if num%5 == 0 && num%3 == 0 {
		logger.Println("Кратно 3 и 5")
	} else {
		logger.Println("Некратно 3 и 5")
	}

	if ((num/10 + num%10) == 7) || ((num/10 + num%10) == 13) {
		logger.Println("Число счастливое")
	} else {
		logger.Println("Число несчастливое")
	}
}

func Run2(logger *log.Logger, reader *bufio.Reader) {
	for {
		logger.Println("Введите номер задачи (1-12); 0 - Выход: ")

		n, err := utils.ReadInt(reader)
		if err != nil {
			logger.Println(err)
			return
		}

		if n != 0 {
			logger.Println("Выбрана задача №", n)
		}

		switch n {
		case 0:
			return
		case 1:
			task1(logger)
		case 2:
			task2(logger, reader)
		case 3:
			task3(logger)
		case 4:
			task4(logger, reader)
		case 5:
			task5(logger, reader)
		case 6:
			task6(logger, reader)
		case 7:
			task7(logger)
		case 8:
			task8(logger, reader)
		case 9:
			task9(logger)
		case 10:
			task10(logger)
		case 11:
		case 12:
		default:
			logger.Println("Задача не найдена!")
		}
	}
}
