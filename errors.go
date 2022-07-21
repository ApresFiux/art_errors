package art_errors

import (
	"fmt"
	"log"
)

func FailLogHandler(err error, failFirst bool, errorList []string, result string) []string {
	if err != nil {
		switch failFirst {
		case true:
			log.Fatal(err.Error())
		default:
			fmt.Printf("%s  %s\n%s", ColorRed, err.Error(), ColorReset)
			errorList = append(errorList, err.Error())
		}
	} else {
		switch result {
		case "0":
			fmt.Printf("%s  Done.%s\n", ColorGreen, ColorReset)
		default:
			fmt.Printf(result)
		}
	}
	return errorList
}
