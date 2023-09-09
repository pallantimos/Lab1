package main

import (
	"fmt"
	"os"
	"regexp"
	"unicode"
	"unicode/utf8"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func checkRegistrate(login string, pass string, pass2 string) (string, bool) {
	listLogin := [5]string{"Aldar", "Aleksey", "Ivan", "Mikhail", "Krug"}

	err := ""
	isCorrect := true

	for i := 0; i < len(listLogin); i++ {
		if login == listLogin[i] {
			err = err + "Логин уже существует "
			isCorrect = false
		}
	}

	if utf8.RuneCountInString(login) < 5 {
		err = err + "Логин меньше 5 символов "
		isCorrect = false
	}

	regex := regexp.MustCompile("^[a-zA-Z0-9_]+$")

	if !regex.MatchString(login) {
		isCorrect = false
		err = err + "Логин содержит некорректные символы "
	}

	if pass != pass2 {
		err = err + "Пароли не совпадают "
		isCorrect = false
	}

	if utf8.RuneCountInString(pass) < 7 {
		err = err + "Пароль меньше 7 символов "
		isCorrect = false
	}

	for _, r := range pass {
		if unicode.Is(unicode.Latin, r) {
			return err + "Пароль содержит латиницу", false
		}
	}

	return err, isCorrect
}

func main() {
	log := logrus.New()
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Не удалось открыть файл логов: %v", err)
	}
	log.SetOutput(file)
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.TextFormatter{})

	log.Info("Приложение запущено")
	log.Info("Логгер сконфигурирован")

	var pass, pass2, login string
	print("Введите логин:\n")
	fmt.Scan(&login)
	log.Info("Введено значение логина: ", login)

	print("Введите пароль:\n")
	fmt.Scan(&pass) // <PASSWORD>
	log.Info("Введено значение пароля: ", pass)

	print("Повторите пароль:\n")
	fmt.Scan(&pass2) // <PASSWORD>
	log.Info("Введено значение повторного пароля: ", pass2)

	errorRegistrate, isCorrect := checkRegistrate(login, pass, pass2)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	if isCorrect {
		log.Info("Логин ", login, " Успешная регистрация")
	} else {
		log.Error("Логин = ", login, " Пароль = ", hashedPassword, " Ошибка = ", errorRegistrate)
	}

	defer file.Close()
}
