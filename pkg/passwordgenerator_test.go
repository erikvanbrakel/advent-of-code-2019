package pkg

import (
	"strconv"
	"testing"
)

func GeneratePasswords(lower, upper int) []int {

	validPasswords := []int { }

	validate := func(password string) bool {
		isValid := false
		for i := 1; i < len(password); i++ {
			if password[i] < password[i-1] { return false }
			if password[i] == password[i-1] {
				if i+1 < len(password) && password[i] == password[i+1] {
					for ; i < len(password); i++ {
						if password[i] != password[i-1] { i--; break }
					}
				} else {
					isValid = true
				}
			}
		}

		return isValid
	}

	for p:=lower;p<upper;p++ {
		strval := strconv.Itoa(p)
		if validate(strval) {
			validPasswords = append(validPasswords, p)
		}
	}

	return validPasswords
}

func TestPasswordGenerator(t *testing.T) {
	validPasswords := GeneratePasswords(147981,691423)

	for _, pwd := range validPasswords {
		if pwd == 222222 {
			t.Log("Found valid password 222222")
		}
		if pwd == 402345 {
			t.Error("Found invalid password 402345")
		}
	}

	t.Logf("Valid passwords found: %v", len(validPasswords))
}