package plan_usecases

import (
	"time"

	"github.com/emanueltimlopez/books-motivation/internal/user"
)

func BooksLeft(user *user.User) int {
	pagesPerBook := 350
	wordsPerPage := 350
	wordsPerBook := wordsPerPage * pagesPerBook

	wordsPerMinute := user.Plan.Words
	minutesPerDay := user.Plan.Minutes
	wordsPerDay := wordsPerMinute * minutesPerDay

	wordsPerYear := wordsPerDay * 365

	booksPerYear := float64(wordsPerYear) / float64(wordsPerBook)

	currentTime := time.Now()
	currentYear := currentTime.Year()
	years := currentYear - user.Birth
	lifeSpectation := 90
	yearsLeft := lifeSpectation - years

	booksLeft := booksPerYear * float64(yearsLeft)

	return int(booksLeft)
}
