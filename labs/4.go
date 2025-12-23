package labs

import (
	"bufio"
	"log"
	"sort"

	"github.com/kamilrahmatullin/lab/utils"
)

func Run4(logger *log.Logger, reader *bufio.Reader) {
	for {
		logger.Println("Введите номер задачи: ")
		logger.Println("1 - Умный анализатор успеваемости")
		logger.Println("2 - Текст-шифровальщик")
		logger.Println("3 - База данных студентов")
		logger.Println("4 - Система управления библиотекой")
		logger.Println("5 - Игра на память с динамическими структурами")
		logger.Println("0 - Выход")

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
				task4_1(logger, reader)
				runAgain = postMenu(logger, reader)
			}
		case 2:
		case 3:
		case 4:
		case 5:
		default:
			logger.Println("Задача не найдена!")
		}
	}
}

type Student struct {
	Name     string
	Subjects []Subject
}

type Subject struct {
	SubjectName string
	Grade       int
}

type SubjectInfo struct {
	TotalCount  int
	TotalGrades int
}

const minAvgGrade = 3.5

func task4_1(logger *log.Logger, reader *bufio.Reader) {
	students := getStudents(logger, reader)

	subjectInfos := getSubjectInfos(students)

	for name, info := range subjectInfos {

		if info.TotalCount == 0 {
			logger.Printf(" | Предмет %s пуст |\n", name)
			continue
		}

		avg := float64(info.TotalGrades) / float64(info.TotalCount)
		logger.Printf(" | Предмет %s  |  Средний балл %.2f |\n", name, avg)
	}

	logger.Print("\n\n")

	sortStudents(students)

	for _, student := range students {
		totalCount := len(student.Subjects)
		totalGrade := 0

		for _, subject := range student.Subjects {
			totalGrade += int(subject.Grade)
		}

		if totalCount == 0 {
			logger.Printf(" | Студент %s | У СТУДЕНТА НЕТ ОЦЕНОК |\n", student.Name)
			continue
		}

		avg := float64(totalGrade) / float64(totalCount)

		if avg > minAvgGrade {
			logger.Printf(" | Студент %s | Средний балл %.2f |\n", student.Name, avg)
		} else {
			logger.Printf(" | Студент %s | Средний балл %.2f | ПРОБЛЕМА С ОЦЕНКОЙ!!! \n", student.Name, avg)
		}
	}
}

func sortStudents(students []Student) {
	sort.Slice(students, func(i, j int) bool {
		si := 0
		sj := 0

		for _, subject := range students[i].Subjects {
			si += int(subject.Grade)
		}
		for _, subject := range students[j].Subjects {
			sj += int(subject.Grade)
		}

		return float64(si)/float64(len(students[i].Subjects)) > float64(sj)/float64(len(students[j].Subjects))
	})
}

func getStudents(logger *log.Logger, reader *bufio.Reader) []Student {
	students := []Student{}
	for {
		logger.Print("Введите имя студента: ")
		name, err := utils.ReadString(reader)
		if err != nil {
			logger.Println(err)
			return students
		}

		subjects := getSubjects(logger, reader)

		students = append(students, Student{
			Name:     name,
			Subjects: subjects,
		})

		logger.Println("Хотите продолжить ввод студентов? 1 - ДА / 0 - НЕТ")
		ch, err := utils.ReadInt(reader)
		if err != nil {
			logger.Println(err)
			return students
		}

		if ch != 1 {
			break
		}
	}

	return students
}

func getSubjects(logger *log.Logger, reader *bufio.Reader) []Subject {
	subjects := []Subject{}
	for {
		logger.Print("Введите название предмета: ")
		subjectName, err := utils.ReadString(reader)
		if err != nil {
			logger.Println(err)
			return subjects
		}
		logger.Printf("Введите оценку для предмета (1-5) %s: ", subjectName)
		grade, err := utils.ReadInt(reader)
		if err != nil {
			logger.Println(err)
			return subjects
		}

		if grade < 1 || grade > 5 {
			logger.Println("Оценки должны быть указаны в диапазоне от 1 до 5")
			continue
		}

		subjects = append(subjects, Subject{
			SubjectName: subjectName,
			Grade:       grade,
		})

		logger.Println("Хотите продолжить ввод предметов? 1 - ДА / 0 - НЕТ")
		ch, err := utils.ReadInt(reader)
		if err != nil {
			logger.Println(err)
			return subjects
		}

		if ch != 1 {
			break
		}
	}
	return subjects
}

func getSubjectInfos(students []Student) map[string]*SubjectInfo {
	subjectInfos := make(map[string]*SubjectInfo)

	for _, student := range students {
		for _, subject := range student.Subjects {

			info, ok := subjectInfos[subject.SubjectName]

			if !ok {
				info = &SubjectInfo{TotalCount: 1, TotalGrades: subject.Grade}
			} else {
				info.TotalCount += 1
				info.TotalGrades += subject.Grade
			}

			subjectInfos[subject.SubjectName] = info
		}
	}

	return subjectInfos
}
