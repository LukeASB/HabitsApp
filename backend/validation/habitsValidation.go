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

	if err := validateHabitDaysTarget(habitForValidation.daysTarget); err != nil {
		logger.ErrorLog(helper.GetFunctionName(), fmt.Sprintf("%s", err))
		return fmt.Errorf("%s - %s", helper.GetFunctionName(), err)
	}

	return nil
}

func processHabit(value interface{}, habitForValidation *habitForValidation) error {
	if newHabit, ok := value.(data.NewHabit); ok {
		habitForValidation.name = newHabit.Name
		habitForValidation.daysTarget = newHabit.DaysTarget
	} else if habit, ok := value.(data.Habit); ok {
		habitForValidation.name = habit.Name
		habitForValidation.daysTarget = habit.DaysTarget
	} else {
		return fmt.Errorf("%s - value type is not a habit", helper.GetFunctionName())
	}

	return nil
}

func validateHabitName(name string) error {

	lengthCheck := func(name string) error {
		if len(name) <= 0 {
			return fmt.Errorf("%s -  No Habit Name Supplied", helper.GetFunctionName())
		}

		maxCharacterLength := 255

		if len(name) >= maxCharacterLength {
			return fmt.Errorf("%s -  Habit Name exceeds max character length of %d", helper.GetFunctionName(), maxCharacterLength)
		}

		return nil
	}

	validateHabitLettersNumbersAndColon := func(name string) error {
		matchLettersNumbersAndColon := regexp.MustCompile(`^[a-zA-Z0-9: ]+$`)

		if !matchLettersNumbersAndColon.MatchString(name) {
			return fmt.Errorf("%s - Habit Name is invalid", helper.GetFunctionName())
		}

		return nil
	}

	if err := lengthCheck(name); err != nil {
		return err
	}

	if err := validateHabitLettersNumbersAndColon(name); err != nil {
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
		return fmt.Errorf("%s - Habit Days cannot be less than 0", helper.GetFunctionName())
	}

	if day >= 9999 {
		return fmt.Errorf("%s - Habit Days cannot be more than 9999 days", helper.GetFunctionName())
	}

	return nil
}
