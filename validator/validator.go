package validator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nicolaseliasx/movie-apirest-golang/models"
)

func ValidateMovie(movie *models.Movie) error {
	var validationErrors []string

	if err := requiredField("Title", movie.Title); err != nil {
		validationErrors = append(validationErrors, err.Error())
	}

	if err := requiredField("Ditector", movie.Director); err != nil {
		validationErrors = append(validationErrors, err.Error())
	}

	if err := requiredField("Year", fmt.Sprintf("%d", movie.Year)); err != nil {
		validationErrors = append(validationErrors, err.Error())
	}

	if err := isYearValid(movie.Year); err != nil {
		validationErrors = append(validationErrors, err.Error())
	}

	if err := isYearValid(movie.Year); err != nil {
		validationErrors = append(validationErrors, err.Error())
	}

	if movie.Genre != "" {
		if err := ValidateGenre(movie.Genre); err != nil {
			validationErrors = append(validationErrors, err.Error())
		}
	}

	rating, err := strconv.Atoi(movie.Rating)
	if err != nil {
		validationErrors = append(validationErrors, "Rating must be a number")
	} else if !between(rating, 0, 10) {
		validationErrors = append(validationErrors, "Rating must be between 0 and 10")
	}
	
	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, ", "))
	}

	return nil
}

func requiredField(fieldName, fieldValue string) error {
	if fieldValue == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	return nil
}

func isYearValid(year int) error {
	now := time.Now()

	yearActual := now.Year()

	// This year of first movie in the world
	firstMovieYear := 1888

	if !between(year, firstMovieYear, yearActual) {
		return fmt.Errorf("year must be between 1888 and %d", yearActual)
	}

	return nil
}

func ValidateGenre(genre string) error {
	if !models.validGenres[genre] {
		return errors.New("invalid genre: " + genre)
	}
	return nil
}

func between(checker int, num1 int, num2 int) bool {
	if checker < num1 || checker > num2 {
		return false
	}
	return true
}
