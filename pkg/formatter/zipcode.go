package formatter

import (
	"errors"
	"fmt"
	"regexp"
)

func Zipcode(zipcode string) string {
	match, _ := regexp.MatchString(`\d{5}-\d{3}`, zipcode)

	if match == true {
		return zipcode
	}

	match, _ = regexp.MatchString(`\d{5}\d{3}`, zipcode)
	if match == false {
		panic(errors.New("Zipcode is invalid"))
	}

	pattern := regexp.MustCompile(`\d{5}|\d{3}`)
	zipcodes := pattern.FindAllString(zipcode, 2)

	return fmt.Sprintf("%s-%s", zipcodes[0], zipcodes[1])
}
