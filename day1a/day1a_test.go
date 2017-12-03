package day1a

import (
	"reflect"
	"testing"
)

func TestgetTotal(t *testing.T) {
	numbers := []int{1, 2, 3, 4}
	total := getTotal(numbers)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want %d.", total, 10)
	}
}

func TestGetValidNumbers(t *testing.T) {
	captcha1 := "1122"
	captcha1Expected := []int{1, 2}
	if !reflect.DeepEqual(GetValidNumbers(captcha1), captcha1Expected) {
		t.Errorf("getValid method failed on generating expected results. Input: 1122")
	}

	captcha2 := "1111"
	captcha2Expected := []int{1, 1, 1, 1}
	if !reflect.DeepEqual(GetValidNumbers(captcha2), captcha2Expected) {
		t.Errorf("getValid method failed on generating expected results.")
	}

	captcha3 := "1234"
	var captcha3Expected []int
	if !reflect.DeepEqual(GetValidNumbers(captcha3), captcha3Expected) {
		t.Errorf("getValid method failed on generating expected results.")
	}

	captcha4 := "91212129"
	captcha4Expected := []int{9}
	if !reflect.DeepEqual(GetValidNumbers(captcha4), captcha4Expected) {
		t.Errorf("getValid method failed on generating expected results.")
	}

}
