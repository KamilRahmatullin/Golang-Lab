package labs

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"

	"github.com/kamilrahmatullin/lab/utils"
)

func postMenu(logger *log.Logger, reader *bufio.Reader) bool {
	logger.Print("Выполнить задачу ещё раз? 1 - Да; 0 - Выход: ")

	choice, err := utils.ReadInt(reader)
	if err != nil {
		log.Println(err)
		return false
	}

	switch choice {
	case 1:
		logger.Println("Начинаю выполнение заново...")
		return true
	case 0:
		logger.Println("Осуществляю выход...")
		return false
	default:
		logger.Println("Выбор не распознан! Осуществляю выход...")
		return false
	}
}

func Run3(logger *log.Logger, reader *bufio.Reader) {
	for {
		logger.Println("Введите номер задачи (1-5); 0 - Выход: ")

		n, err := utils.ReadInt(reader)
		if err != nil {
			logger.Println(err)
			return
		}

		if n != 0 {
			logger.Println("Выбрана задача №", n)
		}

		runAgain := true

		switch n {
		case 0:
			return
		case 1:
			for runAgain {
				task3_1(logger)
				runAgain = postMenu(logger, reader)
			}
		case 2:
			operationLogger := &OperationLogger{}

			for runAgain {
				task3_2(logger, reader, operationLogger)
				runAgain = postMenu(logger, reader)
			}
		case 3:
		case 4:
		case 5:
		default:
			logger.Println("Задача не найдена!")
		}
	}
}

func task3_1(logger *log.Logger) {
	students := [15]int{}

	for i := 0; i < 15; i++ {
		students[i] = rand.Int()%4 + 2
	}

	total := 0
	fools := 0
	genius := 0

	two := 0
	three := 0
	four := 0
	five := 0

	for _, grade := range students {
		total += grade

		if grade < 3 {
			fools++
		}

		if grade == 5 {
			genius++
		}

		switch grade {
		case 2:
			two++
			break
		case 3:
			three++
			break
		case 4:
			four++
			break
		default:
			five++
		}
	}

	avg := float64(total) / 15

	logger.Printf("Средний балл группы: %.2f\n", avg)
	logger.Printf("%d учеников отстают.\n", fools)
	logger.Printf("%d учеников получили пятерку.\n", genius)

	logger.Printf("Количество двоек: %d\n", two)
	logger.Printf("Количество троек: %d\n", three)
	logger.Printf("Количество четверок: %d\n", four)
	logger.Printf("Количество пятёрок: %d\n", five)

	if five > four && five > three && five > two {
		logger.Println("Самая частая оценка - 5")
	} else if four > five && four > three && four > two {
		logger.Println("Самая частая оценка - 4")
	} else if three > five && three > four && three > two {
		logger.Println("Самая частая оценка - 4")
	} else {
		logger.Println("Самая частая оценка - 2")
	}
}

type OperationLogger struct {
	operations [5]string
	currentIdx int
}

func (op *OperationLogger) Add(logger *log.Logger, move string) {
	op.operations[op.currentIdx] = move
	op.currentIdx = (op.currentIdx + 1) % len(op.operations)

	logger.Println("\n===================")
	logger.Println("Последние операции:")
	for i, op := range op.operations {
		if op == "" {
			op = "Пусто."
		}
		logger.Printf("%d. %s", i+1, op)
	}
}

func task3_2(logger *log.Logger, reader *bufio.Reader, operationLog *OperationLogger) {
	logger.Print("Введите первое число: ")
	num1, err := utils.ReadInt(reader)
	if err != nil {
		logger.Println(err)
		return
	}

	logger.Print("Введите второе число: ")
	num2, err := utils.ReadInt(reader)
	if err != nil {
		logger.Println(err)
		return
	}

	logger.Print("Введите арифметический знак (+-*/): ")
	ch, err := utils.ReadString(reader)

	var move string

	switch ch {
	case "+":
		move = fmt.Sprintf("%d %s %d = %d", num1, ch, num2, num1+num2)
	case "-":
		move = fmt.Sprintf("%d %s %d = %d", num1, ch, num2, num1-num2)
	case "*":
		move = fmt.Sprintf("%d %s %d = %d", num1, ch, num2, num1*num2)
	case "/":
		if num2 == 0 {
			logger.Println("Делить на ноль нельзя")
			return
		} else {
			move = fmt.Sprintf("%d %s %d = %d", num1, ch, num2, num1/num2)
		}
	default:
		logger.Println("Операция не найдена!")
		return
	}

	logger.Println(move)

	operationLog.Add(logger, move)
}
