package labs

import (
	"bufio"
	"encoding/json"
	"log"
	"sort"
	"strings"

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
			for runAgain {
				task4_2(logger, reader)
				runAgain = postMenu(logger, reader)
			}
		case 3:
			students := make([]StudentDB, 0)
			for runAgain {
				task4_3(&students, logger, reader)
				runAgain = postMenu(logger, reader)
			}
		case 4:
			books := make([]BookDB, 0)

			books = addBooks(books)

			for runAgain {
				task4_4(books, logger, reader)
				runAgain = postMenu(logger, reader)
			}
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

const alfabetStart = 97
const alfabetEnd = 122

const alfabetStartRu = 1072
const alfabetEndRu = 1103

func task4_2(logger *log.Logger, reader *bufio.Reader) {
	logger.Println("Выберите язык: 1 - EN; 2 - RU")
	lang, err := utils.ReadInt(reader)
	if err != nil {
		logger.Println(err)
		return
	}

	logger.Println("Введите текст от 3 до 25 символов:")

	text, err := utils.ReadString(reader)
	if err != nil {
		logger.Println(err)
		return
	}

	if len(text) < 3 || len(text) > 25 {
		logger.Println("Длина текста должна быть от 3 до 25 символов")
		return
	}

	logger.Println("Введите сдвиг от 0 до 25:")
	move, err := utils.ReadInt(reader)
	if err != nil {
		logger.Println(err)
		return
	}

	if move < 0 || move > 25 {
		logger.Println("Сдвиг должен быть от 0 до 25")
		return
	}
	var cipheredText string
	switch lang {
	case 1:
		cipheredText = enCipher(text, move)
	case 2:
		cipheredText = ruCipher(text, move)
	default:
		cipheredText = enCipher(text, move)
	}

	logger.Printf("Зашифрованный текст: %s", cipheredText)

	letters := utils.TextLettersAnalyze(text)
	for letter, count := range letters {
		logger.Printf(" | Буква %s использована %d раз\n", letter, count)
	}

	wordsPalindromes := utils.SearchPalindromes(text)
	if len(wordsPalindromes) == 0 {
		logger.Println("В тексте нет палиндромов")
	} else {
		logger.Printf("В тексте %d палидромов:\n", len(wordsPalindromes))
		for _, palindrome := range wordsPalindromes {
			logger.Println(" >", palindrome)
		}
	}
}

func enCipher(text string, move int) string {
	newText := ""

	for _, ch := range text {
		newR := int(ch) + move
		if newR > alfabetEnd {
			newR = alfabetStart + (newR - alfabetEnd - 1)
		}

		newText += string(rune(newR))
	}

	return newText
}

func ruCipher(text string, move int) string {
	newText := ""

	for _, ch := range text {
		newR := int(ch) + move
		if newR > alfabetEndRu {
			newR = alfabetStartRu + (newR - alfabetEnd - 1)
		}

		newText += string(rune(newR))
	}

	return newText
}

type StudentDB struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	AvgRate  float64  `json:"avg_rate"`
	Subjects []string `json:"subjects"`
}

func task4_3(students *[]StudentDB, logger *log.Logger, reader *bufio.Reader) {
	logger.Println("Выберите ваше действие 1 - Добавление, 2 - Просмотр, 3 - Поиск, 4 - Экспорт:")
	move, err := utils.ReadInt(reader)
	if err != nil {
		logger.Println(err)
		return
	}

	switch move {
	case 1:
		*students = addStudent(*students, logger, reader)
	case 2:
		readStudents(*students, logger, reader)
	case 3:
		searchStudents(*students, logger, reader)
	case 4:
		exportStudentsFile(*students, logger, reader)
	default:
		logger.Println("Действие не найдено!")
	}
}

func addStudent(students []StudentDB, logger *log.Logger, reader *bufio.Reader) []StudentDB {
	logger.Println("Введите имя студента:")
	name, err := utils.ReadString(reader)
	if err != nil {
		logger.Println(err)
		return students
	}

	logger.Println("Введите возраст студента:")
	age, err := utils.ReadInt(reader)
	if err != nil {
		logger.Println(err)
		return students
	}

	logger.Println("Введите средний балл:")
	avgRate, err := utils.ReadFloat(reader)
	if err != nil {
		logger.Println(err)
		return students
	}

	subjects := make([]string, 0)
	logger.Println("Вводите названия предметов по очереди, пустая строка = Выход:")
	for {
		logger.Print(" > ")
		subject, err := utils.ReadString(reader)
		if err != nil {
			logger.Println(err)
			break
		}

		if subject == "" {
			break
		}

		subjects = append(subjects, subject)
	}

	student := StudentDB{
		Name:     name,
		Age:      age,
		AvgRate:  avgRate,
		Subjects: subjects,
	}

	students = append(students, student)
	logger.Println("Студент успешно добавлен в список!")

	return students
}

func readStudents(students []StudentDB, logger *log.Logger, reader *bufio.Reader) {
	logger.Println("Выберите как будет проходить сортировка студентов (1 - По умолчанию; 2 - По среднему баллу; 3 - По возрасту):")
	ch, err := utils.ReadInt(reader)
	if err != nil {
		logger.Println(err)
		return
	}

	switch ch {
	case 1:
	case 2:
		sort.Slice(students, func(i, j int) bool {
			return students[i].AvgRate > students[j].AvgRate
		})
	case 3:
		sort.Slice(students, func(i, j int) bool {
			return students[i].Age > students[j].Age
		})
	}

	printAllStudents(students, logger)
}

func printAllStudents(students []StudentDB, logger *log.Logger) {
	for i, student := range students {
		printStudent(student, i, logger)
	}
}

func searchStudents(students []StudentDB, logger *log.Logger, reader *bufio.Reader) {
	logger.Print("Введите имя студента: ")
	name, err := utils.ReadString(reader)
	if err != nil {
		logger.Println(err)
		return
	}

	foundedStudents := make([]StudentDB, 0)

	for _, student := range students {
		if strings.EqualFold(name, student.Name) {
			foundedStudents = append(foundedStudents, student)
		}
	}

	if len(foundedStudents) == 0 {
		logger.Println("Студент с таким именем не найден!")
		return
	}

	for index, student := range foundedStudents {
		printStudent(student, index, logger)
	}
}

func printStudent(student StudentDB, index int, logger *log.Logger) {
	logger.Printf(" %d) Имя: %s  |  Возраст: %d  |  Средний балл: %.2f  |  Предметы %s\n",
		index, student.Name, student.Age, student.AvgRate, student.Subjects)
}

func exportStudentsFile(students []StudentDB, logger *log.Logger, reader *bufio.Reader) {
	logger.Print("Введите желаемое имя файла:")
	fileName, err := utils.ReadString(reader)
	if err != nil {
		logger.Println(err)
		return
	}
	studentsJson, err := json.Marshal(students)
	if err != nil {
		logger.Println("Ошибка при форматировании студентов: ", err.Error())
		return
	}

	if err = utils.ExportFile(studentsJson, fileName); err != nil {
		logger.Println(err)
		return
	}

	logger.Println("Данные успешно сохранены в файл", fileName)
}

type BookDB struct {
	Title       string
	Author      string
	Publishment int
	Genre       string
	IsAvailable bool
}

func addBooks(books []BookDB) []BookDB {
	book1 := BookDB{
		Title:       "The Great Gatsby",
		Author:      "F. Scott Fitzgerald",
		Publishment: 1925,
		Genre:       "Classic",
		IsAvailable: true,
	}

	book2 := BookDB{
		Title:       "To Kill a Mockingbird",
		Author:      "Harper Lee",
		Publishment: 1960,
		Genre:       "Classic",
		IsAvailable: true,
	}

	book3 := BookDB{
		Title:       "Pride and Prejudice",
		Author:      "Jane Austen",
		Publishment: 1813,
		Genre:       "Classic",
		IsAvailable: true,
	}

	book4 := BookDB{
		Title:       "The Catcher in the Rye",
		Author:      "J.D. Salinger",
		Genre:       "Coming of Age",
		IsAvailable: true,
		Publishment: 1951,
	}

	books = append(books, book1)
	books = append(books, book2)
	books = append(books, book3)
	books = append(books, book4)

	return books
}

func task4_4(books []BookDB, logger *log.Logger, reader *bufio.Reader) {
	logger.Println("1. Просмотр каталога")
	logger.Println("2. Поиск книги")
	logger.Println("3. Получить книгу")
	logger.Print("Выберите действие: ")
	move, err := utils.ReadInt(reader)
	if err != nil {
		logger.Println(err)
		return
	}

	switch move {
	case 1:
		logger.Println("1. По умолчанию")
		logger.Println("2. По жанрам")
		logger.Println("3. По годам")
		logger.Println("4. По Доступности")
		logger.Print("Выберите сортировку: ")
		ch, err := utils.ReadInt(reader)
		if err != nil {
			logger.Println(err)
			return
		}

		switch ch {
		case 1:
		case 2:
			sort.Slice(books, func(i, j int) bool {
				return books[i].Genre > books[j].Genre
			})
		case 3:
			sort.Slice(books, func(i, j int) bool {
				return books[i].Publishment > books[j].Publishment
			})
		case 4:
			sort.Slice(books, func(i, j int) bool {
				if books[i].IsAvailable == true {
					return true
				}

				return false
			})
		}

		readBooks(books, logger)
	case 2:
		logger.Print("Введите название / автора / жанр книги для поиска: ")
		name, err := utils.ReadString(reader)
		if err != nil {
			logger.Println(err)
			return
		}

		searchBook(name, books, logger)
	case 3:
		readBooks(books, logger)

		logger.Print("Введите номер книги, которую хотите забрать: ")
		num, err := utils.ReadInt(reader)
		if err != nil {
			logger.Println(err)
			return
		}

		if num < 1 || num > len(books) {
			logger.Println("Такой книги не найдено!")
			return
		}

		for i, book := range books {
			if i == num-1 {
				if !book.IsAvailable {
					logger.Println("Данной книги нет в наличии")
					break
				}

				book.IsAvailable = false
				books[i] = book

				logger.Printf("Книга %s была успешно получена!\n", book.Title)
				break
			}
		}
	}
}

func readBooks(books []BookDB, logger *log.Logger) {
	for i, book := range books {
		logger.Printf("%d) Название: %s  |  Автор: %s  |  Жанр:  %s  |  Год выпуска: %d  |  Доступность:  %t  |\n",
			i+1, book.Title, book.Author, book.Genre, book.Publishment, book.IsAvailable)
	}
}

func searchBook(name string, books []BookDB, logger *log.Logger) {
	foundedBooks := make([]BookDB, 0)
	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Author), strings.ToLower(name)) || strings.Contains(strings.ToLower(book.Genre), strings.ToLower(name)) || strings.Contains(strings.ToLower(book.Title), strings.ToLower(name)) {
			foundedBooks = append(foundedBooks, book)
		}
	}

	if len(foundedBooks) == 0 {
		logger.Println("Ни одной книги не было найдено!")
		return
	}

	readBooks(foundedBooks, logger)
}
