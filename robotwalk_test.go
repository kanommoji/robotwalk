package robotwalk_test

import (
	"robotwalk"
	"testing"
)

func Test_CalculateNextPosition_X_0_Y_0_Position_North_Walk_W_Should_Be_X_0_Y_1_Position_North(t *testing.T) {
	x := 0
	y := 0
	walk := "W"
	position := robotwalk.North

	expectedX := 0
	expectedY := 1
	expectedPosition := robotwalk.North

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

func Test_CalculateNextPosition_X_0_Y_1_Position_North_Walk_W_Should_Be_X_0_Y_2_Position_North(t *testing.T) {
	x := 0
	y := 1
	walk := "W"
	position := robotwalk.North

	expectedX := 0
	expectedY := 2
	expectedPosition := robotwalk.North

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

func Test_CalculateNextPosition_X_0_Y_2_Position_North_Walk_L_Should_Be_X_0_Y_2_Position_West(t *testing.T) {
	x := 0
	y := 2
	walk := "L"
	position := robotwalk.North

	expectedX := 0
	expectedY := 2
	expectedPosition := robotwalk.West

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

func Test_CalculateNextPosition_X_0_Y_2_Position_West_Walk_W_Should_Be_X_Minus1_Y_2_Position_West(t *testing.T) {
	x := 0
	y := 2
	walk := "W"
	position := robotwalk.West

	expectedX := -1
	expectedY := 2
	expectedPosition := robotwalk.West

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

func Test_CalculateNextPosition_X_Minus1_Y_2_Position_West_Walk_R_Should_Be_X_Minus1_Y_2_Position_North(t *testing.T) {
	x := -1
	y := 2
	walk := "R"
	position := robotwalk.West

	expectedX := -1
	expectedY := 2
	expectedPosition := robotwalk.North

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

func Test_CalculateNextPosition_X_Minus1_Y_2_Position_North_Walk_W_Should_Be_X_Minus1_Y_3_Position_North(t *testing.T) {
	x := -1
	y := 2
	walk := "W"
	position := robotwalk.North

	expectedX := -1
	expectedY := 3
	expectedPosition := robotwalk.North

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

func Test_CalculateNextPosition_X_Minus1_Y_3_Position_North_Walk_W_Should_Be_X_Minus1_Y_4_Position_North(t *testing.T) {
	x := -1
	y := 3
	walk := "W"
	position := robotwalk.North

	expectedX := -1
	expectedY := 4
	expectedPosition := robotwalk.North

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

func Test_CalculateNextPosition_X_0_Y_0_Position_North_Walk_L_Should_Be_X_0_Y_0_Position_West(t *testing.T) {
	x := 0
	y := 0
	walk := "L"
	position := robotwalk.North

	expectedX := 0
	expectedY := 0
	expectedPosition := robotwalk.West

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

func Test_CalculateNextPosition_X_Minus4_Y_0_Position_West_Walk_W_Should_Be_X_Minus4_Y_Minus1_Position_South(t *testing.T) {
	x := -4
	y := 0
	walk := "W"
	position := robotwalk.South

	expectedX := -4
	expectedY := -1
	expectedPosition := robotwalk.South

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
