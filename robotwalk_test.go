package robotwalk_test

import (
	"robotwalk"
	"testing"
)

func Test_CalculateNextPosition_X_0_Y_0_Position_Up_Walk_W_Should_Be_X_0_Y_1_Position_Up(t *testing.T) {
	x := 0
	y := 0
	position := "up"

	expectedX := 0
	expectedY := 1
	expectedPosition := "up"

	actualX, actualY, actualPosition := robotwalk.CalculateNextPosition(x, y, position)

	if expectedX != actualX {
		t.Errorf("Expect X %v but got %v", expectedX, actualX)
	}
	if expectedY != actualY {
		t.Errorf("Expect Y %v but got %v", expectedY, actualY)
	}
	if expectedPosition != actualPosition {
		t.Errorf("Expect Position %v but got %v", expectedPosition, actualPosition)
	}
}
