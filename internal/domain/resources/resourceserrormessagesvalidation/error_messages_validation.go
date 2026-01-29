package resourceserrormessagesvalidation

const (
	NameCannotBeEmptyOrHaveSpaces                       = "The name cannot be empty, and should not start or end with spaces"
	NameCannotBeLessThanThreeOrMoreThanTwentyLetters    = "The name cannot be less than 3 letters and/or more than 20 letters"
	DescriptionCannotStartOrEndWithSpaces               = "The description cannot start or end with spaces"
	DescriptionCannotBeLargerThanOneHundredFiftyLetters = "The description cannot be larger than 150 letters"
	ValueMustBeGreaterThanZero                          = "The value must be greater than 0"
	InsertedValueIsIncorrect                            = "The inserted value is incorrect"
	DueDateMustBeGreaterThanYesterday                   = "The due date must be greater than yesterday"
)
