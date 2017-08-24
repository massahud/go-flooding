package asdfg

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

func TestShouldNotFloodIfTheStartIsAWall(t *testing.T) {
	terrain := [][]rune{
		[]rune("X."),
		[]rune(".."),
	}

	expectedFlooding := [][]rune{
		[]rune("X."),
		[]rune(".."),
	}
	Flood(&terrain, 0, 0)

	assert.Equal(t, expectedFlooding, terrain)
}


func TestShouldFloodASinglePoint(t *testing.T) {
	terrain := [][]rune{
		[]rune("XXX"),
		[]rune("X.X"),
		[]rune("XXX"),
	}

	expectedFlooding := [][]rune{
		[]rune("XXX"),
		[]rune("XXX"),
		[]rune("XXX"),
	}
	Flood(&terrain, 1, 1)

	assert.Equal(t, expectedFlooding, terrain)
}

func TestShouldFloodRight(t *testing.T) {
	terrain := [][]rune{
		[]rune("XXXX"),
		[]rune("X..X"),
		[]rune("XXXX"),
	}

	expectedFlooding := [][]rune{
		[]rune("XXXX"),
		[]rune("XXXX"),
		[]rune("XXXX"),
	}
	Flood(&terrain, 1, 1)

	assert.Equal(t, expectedFlooding, terrain)
}

func TestShouldFloodLeft(t *testing.T) {
	terrain := [][]rune{
		[]rune("XXXX"),
		[]rune("X..X"),
		[]rune("XXXX"),
	}

	expectedFlooding := [][]rune{
		[]rune("XXXX"),
		[]rune("XXXX"),
		[]rune("XXXX"),
	}
	Flood(&terrain, 2, 1)

	assert.Equal(t, expectedFlooding, terrain)
}


func TestShouldFloodUp(t *testing.T) {
	terrain := [][]rune{
		[]rune("XXXX"),
		[]rune("X.XX"),
		[]rune("X.XX"),
		[]rune("XXXX"),
	}

	expectedFlooding := [][]rune{
		[]rune("XXXX"),
		[]rune("XXXX"),
		[]rune("XXXX"),
		[]rune("XXXX"),
	}
	Flood(&terrain, 1, 2)

	assert.Equal(t, expectedFlooding, terrain)
}

func TestShouldFloodDown(t *testing.T) {
	terrain := [][]rune{
		[]rune("XXXX"),
		[]rune("X.XX"),
		[]rune("X.XX"),
		[]rune("XXXX"),
	}

	expectedFlooding := [][]rune{
		[]rune("XXXX"),
		[]rune("XXXX"),
		[]rune("XXXX"),
		[]rune("XXXX"),
	}
	Flood(&terrain, 1, 1)

	assert.Equal(t, expectedFlooding, terrain)
}

func TestShouldFloodSquare(t *testing.T) {
	terrain := [][]rune{
		[]rune("XXXX"),
		[]rune("X..X"),
		[]rune("X..X"),
		[]rune("XXXX"),
	}

	expectedFlooding := [][]rune{
		[]rune("XXXX"),
		[]rune("XXXX"),
		[]rune("XXXX"),
		[]rune("XXXX"),
	}
	Flood(&terrain, 1, 1)

	assert.Equal(t, expectedFlooding, terrain)
}

func TestShouldStopOnWalls(t *testing.T) {
	terrain := [][]rune{
		[]rune("......"),
		[]rune(".XXXX."),
		[]rune(".X..X."),
		[]rune(".X..X."),
		[]rune(".XXXX."),
		[]rune("......"),
	}

	expectedFlooding := [][]rune{
		[]rune("......"),
		[]rune(".XXXX."),
		[]rune(".XXXX."),
		[]rune(".XXXX."),
		[]rune(".XXXX."),
		[]rune("......"),
	}
	Flood(&terrain, 2, 2)

	assert.Equal(t, expectedFlooding, terrain)
}

func TestShouldFloodRing(t *testing.T) {
	terrain := [][]rune{
		[]rune("XXXXX"),
		[]rune("X...X"),
		[]rune("X.X.X"),
		[]rune("X...X"),
		[]rune("XXXXX"),
	}

	expectedFlooding := [][]rune{
		[]rune("XXXXX"),
		[]rune("XXXXX"),
		[]rune("XXXXX"),
		[]rune("XXXXX"),
		[]rune("XXXXX"),
	}
	Flood(&terrain, 1, 2)

	assert.Equal(t, expectedFlooding, terrain)
}

func TestShouldFloodZigZag(t *testing.T) {
	terrain := [][]rune{
		[]rune("XXXXXXXXXXX"),
		[]rune("X.X...X...X"),
		[]rune("X...X...X.X"),
		[]rune("XXXXXXXXXXX"),
	}

	expectedFlooding := [][]rune{
		[]rune("XXXXXXXXXXX"),
		[]rune("XXXXXXXXXXX"),
		[]rune("XXXXXXXXXXX"),
		[]rune("XXXXXXXXXXX"),
	}
	Flood(&terrain, 1, 2)

	assert.Equal(t, expectedFlooding, terrain)
}

func TestShouldFloodUntilBorders(t *testing.T) {
	terrain := [][]rune{
		[]rune("...X."),
		[]rune("..X.."),
		[]rune(".X..."),
	}

	expectedFlooding := [][]rune{
		[]rune("XXXX."),
		[]rune("XXX.."),
		[]rune("XX..."),
	}
	Flood(&terrain, 1, 1)

	assert.Equal(t, expectedFlooding, terrain)
}


func printDebug(terrain *[][]rune) {
	for _, line := range(*terrain) {
		fmt.Println(string(line))
	}
	fmt.Println()
}
