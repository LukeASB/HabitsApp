package validation

import (
	"dohabits/data"
	"dohabits/helper"
	"dohabits/logger"
	"fmt"
	"regexp"
)

type habitForValidation struct {
	name       string
	daysTarget int
	days       int
}

func ValidateHabit(value interface{}, logger logger.ILogger) error {
	habitForValidation := habitForValidation{}

	if err := processHabit(value, &habitForValidation); err != nil {
		logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("%s", err))
		return err
	}

	if err := validateHabitName(habitForValidation.name); err != nil {
		logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("%s", err))
		return fmt.Errorf("%s - %s", helper.GetFunctionName(), err)
	}

	if err := validateHabitDays(habitForValidation.daysTarget); err != nil {
		logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("%s", err))
		return fmt.Errorf("%s - %s", helper.GetFunctionName(), err)
	}

	if err := validateHabitDays(habitForValidation.days); err != nil {
		logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("%s", err))
		return fmt.Errorf("%s - %s", helper.GetFunctionName(), err)
	}

	return nil
}

func processHabit(value interface{}, habitForValidation *habitForValidation) error {
	if newHabit, ok := value.(data.NewHabit); ok {
		habitForValidation.name = newHabit.Name
		habitForValidation.daysTarget = newHabit.DaysTarget
		habitForValidation.days = newHabit.Days
	} else if habit, ok := value.(data.Habit); ok {
		habitForValidation.name = habit.Name
		habitForValidation.daysTarget = habit.DaysTarget
		habitForValidation.days = habit.Days
	} else {
		return fmt.Errorf("validationModel.processHabit - value type is not a habit")
	}

	return nil
}

func validateHabitName(name string) error {

	lengthCheck := func(name string) error {
		if len(name) <= 0 {
			return fmt.Errorf("validateHabitName - No Habit Name Supplied")
		}

		maxCharacterLength := 255

		if len(name) >= maxCharacterLength {
			return fmt.Errorf("validateHabitName - Habit Name exceeds max character length of %d", maxCharacterLength)
		}

		return nil
	}

	validateHabitNameCharactersSpaceOnly := func(name string) error {
		matchUpperLowerCaseLettersOnly := regexp.MustCompile(`^[a-zA-Z ]+$`)

		if !matchUpperLowerCaseLettersOnly.MatchString(name) {
			return fmt.Errorf("validateHabitName - Habit Name is invalid")
		}

		return nil
	}

	if err := lengthCheck(name); err != nil {
		return err
	}

	if err := validateHabitNameCharactersSpaceOnly(name); err != nil {
		return err
	}

	return nil
}

func validateHabitDays(days int) error {
	if err := validateDay(days); err != nil {
		return err
	}

	return nil
}

func validateHabitDaysTarget(days int) error {
	if err := validateDay(days); err != nil {
		return err
	}

	return nil
}

func validateDay(day int) error {
	if day < 0 {
		return fmt.Errorf("validateDay - Habit Days cannot be less than 0")
	}

	if day >= 9999 {
		return fmt.Errorf("validateDay - Habit Days cannot be more than 9999 days")
	}

	return nil
}
