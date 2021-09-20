package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// Parse return a pointer to a new cron Expression
func Parse(cronString string) (*Expression, error) {

	// Check if the expression container 5 parts + 1 part for the command
	cronParts := strings.Split(cronString, " ")
	if len(cronParts) != 6 {
		return nil, errors.New("malformed expression")
	}

	//Initialize a new cron expression
	var expr Expression
	if parsedErr := expr.minuteFieldParser(cronParts[0]); parsedErr != nil {
		log.Fatal(parsedErr.Error())
	}

	if parsedErr := expr.hourFieldParser(cronParts[1]); parsedErr != nil {
		log.Fatal(parsedErr.Error())
	}

	return &expr, nil
}

func main() {
	cronExpression := os.Args[1]
	fmt.Printf("Cron Expression: %v \n", cronExpression)

	cronParts := strings.Split(cronExpression, " ")

	parsedExpression, parseErr := Parse(cronExpression)
	if parseErr != nil {
		log.Fatal(parseErr)
	}

	fmt.Printf("%-14s: %v\n", "minute", strings.Trim(fmt.Sprint(parsedExpression.MinutesList), "[]"))
	fmt.Printf("%-14s: %v\n", "hour", strings.Trim(fmt.Sprint(parsedExpression.HoursList), "[]"))
	fmt.Printf("%-14s: %v\n", "command", cronParts[5])
}
