package robotwalk_test

import (
	"robotwalk"
	"testing"
)

func Test_CalculateNextPosition_X_0_Y_0_Position_0_Walk_W_Should_Be_X_0_Y_1_Position_0(t *testing.T) {
	x := 0
	y := 0
	walk := "W"
	position := 0

	expectedX := 0
	expectedY := 1
	expectedPosition := 0

	actualX, actualY, actualPosition := robotwalk.CalculateNextPosition(x, y, position, walk)

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

func Test_CalculateNextPosition_X_0_Y_1_Position_0_Walk_W_Should_Be_X_0_Y_2_Position_0(t *testing.T) {
	x := 0
	y := 1
	walk := "W"
	position := 0

	expectedX := 0
	expectedY := 2
	expectedPosition := 0

	actualX, actualY, actualPosition := robotwalk.CalculateNextPosition(x, y, position, walk)

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

func Test_CalculateNextPosition_X_0_Y_2_Position_0_Walk_L_Should_Be_X_0_Y_2_Position_3(t *testing.T) {
	x := 0
	y := 2
	walk := "L"
	position := 0

	expectedX := 0
	expectedY := 2
	expectedPosition := 3

	actualX, actualY, actualPosition := robotwalk.CalculateNextPosition(x, y, position, walk)

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

func Test_CalculateNextPosition_X_0_Y_2_Position_3_Walk_W_Should_Be_X_Minus1_Y_2_Position_3(t *testing.T) {
	x := 0
	y := 2
	walk := "W"
	position := 3

	expectedX := -1
	expectedY := 2
	expectedPosition := 3

	actualX, actualY, actualPosition := robotwalk.CalculateNextPosition(x, y, position, walk)

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

func Test_CalculateNextPosition_X_Minus1_Y_2_Position_3_Walk_R_Should_Be_X_Minus1_Y_2_Position_0(t *testing.T) {
	x := -1
	y := 2
	walk := "R"
	position := 3

	expectedX := -1
	expectedY := 2
	expectedPosition := 0

	actualX, actualY, actualPosition := robotwalk.CalculateNextPosition(x, y, position, walk)

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

func Test_CalculateNextPosition_X_Minus1_Y_2_Position_0_Walk_W_Should_Be_X_Minus1_Y_3_Position_0(t *testing.T) {
	x := -1
	y := 2
	walk := "W"
	position := 0

	expectedX := -1
	expectedY := 3
	expectedPosition := 0

	actualX, actualY, actualPosition := robotwalk.CalculateNextPosition(x, y, position, walk)

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

func Test_CalculateNextPosition_X_Minus1_Y_3_Position_0_Walk_W_Should_Be_X_Minus1_Y_4_Position_0(t *testing.T) {
	x := -1
	y := 3
	walk := "W"
	position := 0

	expectedX := -1
	expectedY := 4
	expectedPosition := 0

	actualX, actualY, actualPosition := robotwalk.CalculateNextPosition(x, y, position, walk)

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
