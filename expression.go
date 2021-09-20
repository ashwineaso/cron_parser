package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	patternWildcard            = `^\*$`
	patternSingleValue         = `^[1-5]?[0-9]$`
	patternRangeValue          = `^([1-5]?[0-9])-([1-5]?[0-9])$`
	patternWildcardAndInterval = `^\*/(\d+)$`
	patternValueAndInterval    = `^([1-5]?[0-9])/(\d+)$`
	patternRangeAndInterval    = `^([1-5]?[0-9])-([1-5]?[0-9])/(\d+)$`
)

type FieldConfig struct {
	min         int
	max         int
	defaultList []int
}

var minutesConfig = FieldConfig{
	min: 0,
	max: 59,
}

var hoursConfig = FieldConfig{
	min: 0,
	max: 23,
}

var dayOfMonthConfig = FieldConfig{
	min: 1,
	max: 31,
}

type Expression struct {
	MinutesList []int
	HoursList   []int
	DaysList    []int
}

func (e *Expression) minuteFieldParser(s string) error {
	var parseErr error
	e.MinutesList, parseErr = e.parseField(s, minutesConfig)
	if parseErr != nil {
		return fmt.Errorf("error parsing second field")
	}
	return parseErr
}

func (e *Expression) hourFieldParser(s string) error {
	var parseErr error
	e.HoursList, parseErr = e.parseField(s, hoursConfig)
	if parseErr != nil {
		return fmt.Errorf("error parsing hour field")
	}
	return parseErr
}

func (e *Expression) dayOfMonthParser(s string) error {
	var parseErr error
	e.DaysList, parseErr = e.parseField(s, dayOfMonthConfig)
	if parseErr != nil {
		return fmt.Errorf("error parsing day field")
	}
	return parseErr
}

func (e *Expression) parseField(s string, config FieldConfig) ([]int, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("field missing")
	}

	s = strings.ToLower(s)

	// Case 1: * - Wildcard Expression
	match, err := regexp.MatchString(patternWildcard, s)
	if err != nil {
		return nil, fmt.Errorf("error parsing wildcard expression")
	}
	if match {
		return generateRange(config.min, config.max, 1), nil
	}

	// Case 2: 5 - Single Value Expression
	match, err = regexp.MatchString(patternSingleValue, s)
	if err != nil {
		return nil, fmt.Errorf("error parsing single value expression")
	}
	if match {
		val, _ := strconv.Atoi(s)
		return []int{val}, nil
	}

	// Case 3: 7-21 - Range Value expression
	re, _ := regexp.Compile(patternRangeValue)
	parts := re.FindStringSubmatch(s)
	if err != nil {
		return nil, fmt.Errorf("error parsing range value expression")
	}
	if len(parts) > 0 {
		min, _ := strconv.Atoi(parts[1])
		max, _ := strconv.Atoi(parts[2])
		// Validate that the range
		if min < config.min || max > config.max || max < min {
			return nil, fmt.Errorf("error ")
		}
		return generateRange(min, max, 1), nil
	}

	// Case 4: */12 - Wildcard and Interval Expression
	re, _ = regexp.Compile(patternWildcardAndInterval)
	parts = re.FindStringSubmatch(s)
	if err != nil {
		return nil, fmt.Errorf("error parsing wildcard and interval expression")
	}
	if len(parts) > 0 {
		step, _ := strconv.Atoi(parts[1])
		// Check that the step is valid
		return generateRange(config.min, config.max, step), nil
	}

	// Case 5: 2/12 - Value and Interval Expression
	re, _ = regexp.Compile(patternValueAndInterval)
	parts = re.FindStringSubmatch(s)
	if err != nil {
		return nil, fmt.Errorf("error parsing value and internal expression")
	}
	if len(parts) > 0 {
		min, _ := strconv.Atoi(parts[1])
		step, _ := strconv.Atoi(parts[2])
		if min < config.min || step > config.max {
			return nil, fmt.Errorf("invalid values in expression")
		}
		return generateRange(min, config.max, step), nil
	}

	// Case 6: 5-20/2 - Range and Internal Expression
	re, _ = regexp.Compile(patternRangeAndInterval)
	parts = re.FindStringSubmatch(s)
	if err != nil {
		return nil, fmt.Errorf("error parsing range and interval expression")
	}
	if len(parts) > 0 {
		min, _ := strconv.Atoi(parts[1])
		max, _ := strconv.Atoi(parts[2])
		step, _ := strconv.Atoi(parts[3])
		if min < config.min || max > config.max || max < min || step > config.max {
			return nil, fmt.Errorf("invalid values in expression")
		}
		return generateRange(min, max, step), nil
	}

	return nil, nil
}

func generateRange(min, max, step int) []int {
	r := make([]int, 0, 1+(max-min)/step)
	for min <= max {
		r = append(r, min)
		min += step
	}
	return r
}
