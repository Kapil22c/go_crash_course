package main

import (
	"fmt"
	"regexp"
)

func CheckPasswordLever(ps string) error {
	if len(ps) < 9 {
		return fmt.Errorf("password len is < 9")
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		return fmt.Errorf("password need number")
	}
	if b, err := regexp.MatchString(a_z, ps); !b || err != nil {
		return fmt.Errorf("password need Lowercase letter")
	}
	if b, err := regexp.MatchString(A_Z, ps); !b || err != nil {
		return fmt.Errorf("password need Uppercase letter")
	}
	if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
		return fmt.Errorf("password need symbol")
	}
	return nil
}

func main() {
	// err := CheckPasswordLever("kapil")
	// err := CheckPasswordLever("kapil@fklfffd")
	// err := CheckPasswordLever("kapil@adsd44444")
	err := CheckPasswordLever("Kapil@4545")
	fmt.Println(err)
}
