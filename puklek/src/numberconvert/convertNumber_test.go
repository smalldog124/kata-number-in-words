package numberconvert_test

import (
	. "numberconvert"
	"testing"
)

func Test_ConvertNumber_Input_100000_Dot_00_Should_Be_One_Hundred_Thousand_Baht(t *testing.T) {
	expected := "หนึ่งแสนบาทถ้วน"

	money := 100000.00

	actual := ConvertNumber(money)

	if expected != actual {
		t.Errorf("expected %s but got it %s", expected, actual)
	}

}
