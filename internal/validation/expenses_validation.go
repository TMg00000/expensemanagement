package validation

import (
	"expensemanagement/internal/domain/request"
	"expensemanagement/internal/domain/resources/resourceserrormessagesvalidation"
	"strings"
	"time"
)

func nameValidation(e request.Expenses, ListError *[]string) {
	if (len(strings.TrimSpace(e.Name)) == 0) ||
		(len(e.Name) == 0) ||
		strings.HasPrefix(e.Name, " ") || strings.HasSuffix(e.Name, " ") {

		*ListError = append(*ListError, resourceserrormessagesvalidation.NameCannotBeEmptyOrHaveSpaces)
	}

	if len(e.Name) < 2 && len(e.Name) > 20 {
		*ListError = append(*ListError, resourceserrormessagesvalidation.NameCannotBeLessThanThreeOrMoreThanTwentyLetters)
	}
}

func descriptionValidation(e request.Expenses, ListError *[]string) {
	if strings.HasPrefix(e.Description, " ") || strings.HasSuffix(e.Description, " ") {
		*ListError = append(*ListError, resourceserrormessagesvalidation.DescriptionCannotStartOrEndWithSpaces)
	}

	if len(e.Description) >= 150 {
		*ListError = append(*ListError, resourceserrormessagesvalidation.DescriptionCannotBeLargerThanOneHundredFiftyLetters)
	}
}

func valueValidation(e request.Expenses, ListError *[]string) {
	if e.Value < 0 {
		*ListError = append(*ListError, resourceserrormessagesvalidation.ValueMustBeGreaterThanZero)
	}
}

func dateValidation(e request.Expenses, ListError *[]string) {
	if e.DueDate.Before(time.Now().AddDate(0, 0, -1)) {
		*ListError = append(*ListError, resourceserrormessagesvalidation.DueDateMustBeGreaterThanYesterday)
	}
}

func ListErrorMessages(e request.Expenses) []string {
	var ListError []string

	nameValidation(e, &ListError)
	descriptionValidation(e, &ListError)
	valueValidation(e, &ListError)
	dateValidation(e, &ListError)

	return ListError
}
