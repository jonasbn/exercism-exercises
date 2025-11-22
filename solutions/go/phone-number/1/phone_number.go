package phonenumber

import (
	"fmt"
	"regexp"
)

func Number(phoneNumber string) (string, error) {

	re_clean := regexp.MustCompile(`[\D]+`)
	phoneNumber = re_clean.ReplaceAllString(phoneNumber, "")

	var countryCode string
	var areaCode string
	var exchangeCode string

	if len(phoneNumber) == 11 {
		countryCode = phoneNumber[0:1]
		areaCode = phoneNumber[1:4]
		exchangeCode = phoneNumber[4:]

		if countryCode != "1" {
			return "", fmt.Errorf("invalid country code")
		}

		re_strip_country_code := regexp.MustCompile(`^\d{1}`)
		phoneNumber = re_strip_country_code.ReplaceAllString(phoneNumber, "")

	} else if len(phoneNumber) == 10 {
		areaCode = phoneNumber[0:3]
		exchangeCode = phoneNumber[3:]
	} else {
		return "", fmt.Errorf("invalid phone number length")
	}

	if areaCode[0] == '0' || areaCode[0] == '1' {
		return "", fmt.Errorf("invalid area code")
	}

	if exchangeCode[0] == '0' || exchangeCode[0] == '1' {
		return "", fmt.Errorf("invalid exchange code")
	}

	return phoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	strippedPhoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	areaCode := strippedPhoneNumber[0:3]
	return areaCode, nil
}

func Format(phoneNumber string) (string, error) {
	strippedPhoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", strippedPhoneNumber[:3], strippedPhoneNumber[3:6], strippedPhoneNumber[6:]), nil
}
