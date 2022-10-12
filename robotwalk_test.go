package robotwalk_test

import (
	"robotwalk"
	"testing"
)

func Test_SetStartWalking(t *testing.T) {
	expectedX := 0
	expectedY := 0
	expectedRow := 4
	expectedColumn := 4
	expectedPosition := robotwalk.North

	robotwalk.SetStartWalking()

	actualX := robotwalk.X
	actualY := robotwalk.Y
	actualRow := robotwalk.Row
	actualColumn := robotwalk.Column
	actualPosition := robotwalk.Position

	if expectedX != actualX {
		t.Errorf("Expect X %v but got %v", expectedX, actualX)
	}
	if expectedY != actualY {
		t.Errorf("Expect Y %v but got %v", expectedY, actualY)
	}
	if expectedRow != actualRow {
		t.Errorf("Expect Row %v but got %v", expectedRow, actualRow)
	}
	if expectedColumn != actualColumn {
		t.Errorf("Expect Column %v but got %v", expectedColumn, actualColumn)
	}
	if expectedPosition != actualPosition {
		t.Errorf("Expect Position %v but got %v", expectedPosition, actualPosition)
	}
}

func Test_CalculateNextPosition_X_0_Y_0_Row_4_Column_4_Position_North_Walk_W(t *testing.T) {
	robotwalk.X = 0
	robotwalk.Y = 0
	robotwalk.Row = 4
	robotwalk.Column = 4
	robotwalk.Position = robotwalk.North
	walk := "W"

	expectedX := 0
	expectedY := 1
	expectedRow := 3
	expectedColumn := 4
	expectedPosition := robotwalk.North

	robotwalk.CalculateNextPosition(walk)

	actualX := robotwalk.X
	actualY := robotwalk.Y
	actualRow := robotwalk.Row
	actualColumn := robotwalk.Column
	actualPosition := robotwalk.Position

	if expectedX != actualX {
		t.Errorf("Expect X %v but got %v", expectedX, actualX)
	}
	if expectedY != actualY {
		t.Errorf("Expect Y %v but got %v", expectedY, actualY)
	}
	if expectedRow != actualRow {
		t.Errorf("Expect Row %v but got %v", expectedRow, actualRow)
	}
	if expectedColumn != actualColumn {
		t.Errorf("Expect Column %v but got %v", expectedColumn, actualColumn)
	}
	if expectedPosition != actualPosition {
		t.Errorf("Expect Position %v but got %v", expectedPosition, actualPosition)
	}
}

func Test_CalculateNextPosition_X_0_Y_1_Row_3_Column_4_Position_North_Walk_W(t *testing.T) {
	robotwalk.X = 0
	robotwalk.Y = 1
	robotwalk.Row = 3
	robotwalk.Column = 4
	robotwalk.Position = robotwalk.North
	walk := "W"

	expectedX := 0
	expectedY := 2
	expectedRow := 2
	expectedColumn := 4
	expectedPosition := robotwalk.North

	robotwalk.CalculateNextPosition(walk)

	actualX := robotwalk.X
	actualY := robotwalk.Y
	actualRow := robotwalk.Row
	actualColumn := robotwalk.Column
	actualPosition := robotwalk.Position

	if expectedX != actualX {
		t.Errorf("Expect X %v but got %v", expectedX, actualX)
	}
	if expectedY != actualY {
		t.Errorf("Expect Y %v but got %v", expectedY, actualY)
	}
	if expectedRow != actualRow {
		t.Errorf("Expect Row %v but got %v", expectedRow, actualRow)
	}
	if expectedColumn != actualColumn {
		t.Errorf("Expect Column %v but got %v", expectedColumn, actualColumn)
	}
	if expectedPosition != actualPosition {
		t.Errorf("Expect Position %v but got %v", expectedPosition, actualPosition)
	}
}

func Test_CalculateNextPosition_X_0_Y_2_Row_2_Column_4_Position_North_Walk_L(t *testing.T) {
	robotwalk.X = 0
	robotwalk.Y = 2
	robotwalk.Row = 2
	robotwalk.Column = 4
	robotwalk.Position = robotwalk.North
	walk := "L"

	expectedX := 0
	expectedY := 2
	expectedRow := 2
	expectedColumn := 4
	expectedPosition := robotwalk.West

	robotwalk.CalculateNextPosition(walk)

	actualX := robotwalk.X
	actualY := robotwalk.Y
	actualRow := robotwalk.Row
	actualColumn := robotwalk.Column
	actualPosition := robotwalk.Position

	if expectedX != actualX {
		t.Errorf("Expect X %v but got %v", expectedX, actualX)
	}
	if expectedY != actualY {
		t.Errorf("Expect Y %v but got %v", expectedY, actualY)
	}
	if expectedRow != actualRow {
		t.Errorf("Expect Row %v but got %v", expectedRow, actualRow)
	}
	if expectedColumn != actualColumn {
		t.Errorf("Expect Column %v but got %v", expectedColumn, actualColumn)
	}
	if expectedPosition != actualPosition {
		t.Errorf("Expect Position %v but got %v", expectedPosition, actualPosition)
	}
}

func Test_CalculateNextPosition_X_0_Y_2_Row_2_Column_4_Position_West_Walk_W(t *testing.T) {
	robotwalk.X = 0
	robotwalk.Y = 2
	robotwalk.Row = 2
	robotwalk.Column = 4
	robotwalk.Position = robotwalk.West
	walk := "W"

	expectedX := -1
	expectedY := 2
	expectedRow := 2
	expectedColumn := 3
	expectedPosition := robotwalk.West

	robotwalk.CalculateNextPosition(walk)

	actualX := robotwalk.X
	actualY := robotwalk.Y
	actualRow := robotwalk.Row
	actualColumn := robotwalk.Column
	actualPosition := robotwalk.Position

	if expectedX != actualX {
		t.Errorf("Expect X %v but got %v", expectedX, actualX)
	}
	if expectedY != actualY {
		t.Errorf("Expect Y %v but got %v", expectedY, actualY)
	}
	if expectedRow != actualRow {
		t.Errorf("Expect Row %v but got %v", expectedRow, actualRow)
	}
	if expectedColumn != actualColumn {
		t.Errorf("Expect Column %v but got %v", expectedColumn, actualColumn)
	}
	if expectedPosition != actualPosition {
		t.Errorf("Expect Position %v but got %v", expectedPosition, actualPosition)
	}
}

func Test_CalculateNextPosition_X_Minus_1_Y_2_Row_2_Column_3_Position_West_Walk_R(t *testing.T) {
	robotwalk.X = -1
	robotwalk.Y = 2
	robotwalk.Row = 2
	robotwalk.Column = 3
	robotwalk.Position = robotwalk.West
	walk := "R"

	expectedX := -1
	expectedY := 2
	expectedRow := 2
	expectedColumn := 3
	expectedPosition := robotwalk.North

	robotwalk.CalculateNextPosition(walk)

	actualX := robotwalk.X
	actualY := robotwalk.Y
	actualRow := robotwalk.Row
	actualColumn := robotwalk.Column
	actualPosition := robotwalk.Position

	if expectedX != actualX {
		t.Errorf("Expect X %v but got %v", expectedX, actualX)
	}
	if expectedY != actualY {
		t.Errorf("Expect Y %v but got %v", expectedY, actualY)
	}
	if expectedRow != actualRow {
		t.Errorf("Expect Row %v but got %v", expectedRow, actualRow)
	}
	if expectedColumn != actualColumn {
		t.Errorf("Expect Column %v but got %v", expectedColumn, actualColumn)
	}
	if expectedPosition != actualPosition {
		t.Errorf("Expect Position %v but got %v", expectedPosition, actualPosition)
	}
}

func Test_CalculateNextPosition_X_Minus_1_Y_2_Row_2_Column_3_Position_North_Walk_W(t *testing.T) {
	robotwalk.X = -1
	robotwalk.Y = 2
	robotwalk.Row = 2
	robotwalk.Column = 3
	robotwalk.Position = robotwalk.North
	walk := "W"

	expectedX := -1
	expectedY := 3
	expectedRow := 1
	expectedColumn := 3
	expectedPosition := robotwalk.North

	robotwalk.CalculateNextPosition(walk)

	actualX := robotwalk.X
	actualY := robotwalk.Y
	actualRow := robotwalk.Row
	actualColumn := robotwalk.Column
	actualPosition := robotwalk.Position

	if expectedX != actualX {
		t.Errorf("Expect X %v but got %v", expectedX, actualX)
	}
	if expectedY != actualY {
		t.Errorf("Expect Y %v but got %v", expectedY, actualY)
	}
	if expectedRow != actualRow {
		t.Errorf("Expect Row %v but got %v", expectedRow, actualRow)
	}
	if expectedColumn != actualColumn {
		t.Errorf("Expect Column %v but got %v", expectedColumn, actualColumn)
	}
	if expectedPosition != actualPosition {
		t.Errorf("Expect Position %v but got %v", expectedPosition, actualPosition)
	}
}

func Test_CalculateNextPosition_X_Minus_1_Y_3_Row_1_Column_3_Position_North_Walk_W(t *testing.T) {
	robotwalk.X = -1
	robotwalk.Y = 3
	robotwalk.Row = 1
	robotwalk.Column = 3
	robotwalk.Position = robotwalk.North
	walk := "W"

	expectedX := -1
	expectedY := 4
	expectedRow := 0
	expectedColumn := 3
	expectedPosition := robotwalk.North

	robotwalk.CalculateNextPosition(walk)

	actualX := robotwalk.X
	actualY := robotwalk.Y
	actualRow := robotwalk.Row
	actualColumn := robotwalk.Column
	actualPosition := robotwalk.Position

	if expectedX != actualX {
		t.Errorf("Expect X %v but got %v", expectedX, actualX)
	}
	if expectedY != actualY {
		t.Errorf("Expect Y %v but got %v", expectedY, actualY)
	}
	if expectedRow != actualRow {
		t.Errorf("Expect Row %v but got %v", expectedRow, actualRow)
	}
	if expectedColumn != actualColumn {
		t.Errorf("Expect Column %v but got %v", expectedColumn, actualColumn)
	}
	if expectedPosition != actualPosition {
		t.Errorf("Expect Position %v but got %v", expectedPosition, actualPosition)
	}
}

func Test_CalculateNextPosition_X_0_Y_0_Row_4_Column_4_Position_North_Walk_L(t *testing.T) {
	robotwalk.X = 0
	robotwalk.Y = 0
	robotwalk.Row = 4
	robotwalk.Column = 4
	robotwalk.Position = robotwalk.North
	walk := "L"

	expectedX := 0
	expectedY := 0
	expectedRow := 4
	expectedColumn := 4
	expectedPosition := robotwalk.West

	robotwalk.CalculateNextPosition(walk)

	actualX := robotwalk.X
	actualY := robotwalk.Y
	actualRow := robotwalk.Row
	actualColumn := robotwalk.Column
	actualPosition := robotwalk.Position

	if expectedX != actualX {
		t.Errorf("Expect X %v but got %v", expectedX, actualX)
	}
	if expectedY != actualY {
		t.Errorf("Expect Y %v but got %v", expectedY, actualY)
	}
	if expectedRow != actualRow {
		t.Errorf("Expect Row %v but got %v", expectedRow, actualRow)
	}
	if expectedColumn != actualColumn {
		t.Errorf("Expect Column %v but got %v", expectedColumn, actualColumn)
	}
	if expectedPosition != actualPosition {
		t.Errorf("Expect Position %v but got %v", expectedPosition, actualPosition)
	}
}

func Test_CalculateNextPosition_X_0_Y_0_Row_4_Column_4_Position_West_Walk_W(t *testing.T) {
	robotwalk.X = 0
	robotwalk.Y = 0
	robotwalk.Row = 4
	robotwalk.Column = 4
	robotwalk.Position = robotwalk.West
	walk := "W"

	expectedX := -1
	expectedY := 0
	expectedRow := 4
	expectedColumn := 3
	expectedPosition := robotwalk.West

	robotwalk.CalculateNextPosition(walk)

	actualX := robotwalk.X
	actualY := robotwalk.Y
	actualRow := robotwalk.Row
	actualColumn := robotwalk.Column
	actualPosition := robotwalk.Position

	if expectedX != actualX {
		t.Errorf("Expect X %v but got %v", expectedX, actualX)
	}
	if expectedY != actualY {
		t.Errorf("Expect Y %v but got %v", expectedY, actualY)
	}
	if expectedRow != actualRow {
		t.Errorf("Expect Row %v but got %v", expectedRow, actualRow)
	}
	if expectedColumn != actualColumn {
		t.Errorf("Expect Column %v but got %v", expectedColumn, actualColumn)
	}
	if expectedPosition != actualPosition {
		t.Errorf("Expect Position %v but got %v", expectedPosition, actualPosition)
	}
}

func Test_CalculateNextPosition_X_Minus_1_Y_0_Row_4_Column_3_Position_West_Walk_W(t *testing.T) {
	robotwalk.X = -1
	robotwalk.Y = 0
	robotwalk.Row = 4
	robotwalk.Column = 3
	robotwalk.Position = robotwalk.West
	walk := "W"

	expectedX := -2
	expectedY := 0
	expectedRow := 4
	expectedColumn := 2
	expectedPosition := robotwalk.West

	robotwalk.CalculateNextPosition(walk)

	actualX := robotwalk.X
	actualY := robotwalk.Y
	actualRow := robotwalk.Row
	actualColumn := robotwalk.Column
	actualPosition := robotwalk.Position

	if expectedX != actualX {
		t.Errorf("Expect X %v but got %v", expectedX, actualX)
	}
	if expectedY != actualY {
		t.Errorf("Expect Y %v but got %v", expectedY, actualY)
	}
	if expectedRow != actualRow {
		t.Errorf("Expect Row %v but got %v", expectedRow, actualRow)
	}
	if expectedColumn != actualColumn {
		t.Errorf("Expect Column %v but got %v", expectedColumn, actualColumn)
	}
	if expectedPosition != actualPosition {
		t.Errorf("Expect Position %v but got %v", expectedPosition, actualPosition)
	}
}

func Test_CalculateNextPosition_X_Minus_2_Y_0_Row_4_Column_2_Position_West_Walk_W(t *testing.T) {
	robotwalk.X = -2
	robotwalk.Y = 0
	robotwalk.Row = 4
	robotwalk.Column = 2
	robotwalk.Position = robotwalk.West
	walk := "W"

	expectedX := -3
	expectedY := 0
	expectedRow := 4
	expectedColumn := 1
	expectedPosition := robotwalk.West

	robotwalk.CalculateNextPosition(walk)

	actualX := robotwalk.X
	actualY := robotwalk.Y
	actualRow := robotwalk.Row
	actualColumn := robotwalk.Column
	actualPosition := robotwalk.Position

	if expectedX != actualX {
		t.Errorf("Expect X %v but got %v", expectedX, actualX)
	}
	if expectedY != actualY {
		t.Errorf("Expect Y %v but got %v", expectedY, actualY)
	}
	if expectedRow != actualRow {
		t.Errorf("Expect Row %v but got %v", expectedRow, actualRow)
	}
	if expectedColumn != actualColumn {
		t.Errorf("Expect Column %v but got %v", expectedColumn, actualColumn)
	}
	if expectedPosition != actualPosition {
		t.Errorf("Expect Position %v but got %v", expectedPosition, actualPosition)
	}
}

func Test_CalculateNextPosition_X_Minus_3_Y_0_Row_4_Column_1_Position_West_Walk_W(t *testing.T) {
	robotwalk.X = -3
	robotwalk.Y = 0
	robotwalk.Row = 4
	robotwalk.Column = 1
	robotwalk.Position = robotwalk.West
	walk := "W"

	expectedX := -4
	expectedY := 0
	expectedRow := 4
	expectedColumn := 0
	expectedPosition := robotwalk.West

	robotwalk.CalculateNextPosition(walk)

	actualX := robotwalk.X
	actualY := robotwalk.Y
	actualRow := robotwalk.Row
	actualColumn := robotwalk.Column
	actualPosition := robotwalk.Position

	if expectedX != actualX {
		t.Errorf("Expect X %v but got %v", expectedX, actualX)
	}
	if expectedY != actualY {
		t.Errorf("Expect Y %v but got %v", expectedY, actualY)
	}
	if expectedRow != actualRow {
		t.Errorf("Expect Row %v but got %v", expectedRow, actualRow)
	}
	if expectedColumn != actualColumn {
		t.Errorf("Expect Column %v but got %v", expectedColumn, actualColumn)
	}
	if expectedPosition != actualPosition {
		t.Errorf("Expect Position %v but got %v", expectedPosition, actualPosition)
	}
}

func Test_CalculateNextPosition_X_Minus_4_Y_0_Row_4_Column_0_Position_West_Walk_L(t *testing.T) {
	robotwalk.X = -4
	robotwalk.Y = 0
	robotwalk.Row = 4
	robotwalk.Column = 0
	robotwalk.Position = robotwalk.West
	walk := "L"

	expectedX := -4
	expectedY := 0
	expectedRow := 4
	expectedColumn := 0
	expectedPosition := robotwalk.South

	robotwalk.CalculateNextPosition(walk)

	actualX := robotwalk.X
	actualY := robotwalk.Y
	actualRow := robotwalk.Row
	actualColumn := robotwalk.Column
	actualPosition := robotwalk.Position

	if expectedX != actualX {
		t.Errorf("Expect X %v but got %v", expectedX, actualX)
	}
	if expectedY != actualY {
		t.Errorf("Expect Y %v but got %v", expectedY, actualY)
	}
	if expectedRow != actualRow {
		t.Errorf("Expect Row %v but got %v", expectedRow, actualRow)
	}
	if expectedColumn != actualColumn {
		t.Errorf("Expect Column %v but got %v", expectedColumn, actualColumn)
	}
	if expectedPosition != actualPosition {
		t.Errorf("Expect Position %v but got %v", expectedPosition, actualPosition)
	}
}

func Test_CalculateNextPosition_X_Minus_4_Y_0_Row_4_Column_0_Position_South_Walk_W(t *testing.T) {
	robotwalk.X = -4
	robotwalk.Y = 0
	robotwalk.Row = 4
	robotwalk.Column = 0
	robotwalk.Position = robotwalk.South
	walk := "W"

	expectedX := -4
	expectedY := -1
	expectedRow := 5
	expectedColumn := 0
	expectedPosition := robotwalk.South

	robotwalk.CalculateNextPosition(walk)

	actualX := robotwalk.X
	actualY := robotwalk.Y
	actualRow := robotwalk.Row
	actualColumn := robotwalk.Column
	actualPosition := robotwalk.Position

	if expectedX != actualX {
		t.Errorf("Expect X %v but got %v", expectedX, actualX)
	}
	if expectedY != actualY {
		t.Errorf("Expect Y %v but got %v", expectedY, actualY)
	}
	if expectedRow != actualRow {
		t.Errorf("Expect Row %v but got %v", expectedRow, actualRow)
	}
	if expectedColumn != actualColumn {
		t.Errorf("Expect Column %v but got %v", expectedColumn, actualColumn)
	}
	if expectedPosition != actualPosition {
		t.Errorf("Expect Position %v but got %v", expectedPosition, actualPosition)
	}
}

func Test_RobotWalk_WWLWRWW(t *testing.T) {
	walk := "WWLWRWW"
	expectedX := -1
	expectedY := 4

	robotwalk.RobotWalk(walk)

	actualX := robotwalk.X
	actualY := robotwalk.Y

	if expectedX != actualX {
		t.Errorf("Expect X %v but got %v", expectedX, actualX)
	}
	if expectedY != actualY {
		t.Errorf("Expect Y %v but got %v", expectedY, actualY)
	}
}

func Test_RobotWalk_LWWWWLW(t *testing.T) {
	walk := "LWWWWLW"
	expectedX := -4
	expectedY := -1

	robotwalk.RobotWalk(walk)

	actualX := robotwalk.X
	actualY := robotwalk.Y

	if expectedX != actualX {
		t.Errorf("Expect X %v but got %v", expectedX, actualX)
	}
	if expectedY != actualY {
		t.Errorf("Expect Y %v but got %v", expectedY, actualY)
	}
}
