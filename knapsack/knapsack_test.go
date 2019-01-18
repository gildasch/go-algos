package knapsack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test cases from https://people.sc.fsu.edu/~jburkardt/datasets/knapsack_01/knapsack_01.html

func TestKnapsack(t *testing.T) {
	selection := Knapsack(
		165,
		[]int{23, 31, 29, 44, 53, 38, 63, 85, 89, 82},
		[]int{92, 57, 49, 68, 60, 43, 67, 84, 87, 72},
	)

	expected := []int{1, 1, 1, 1, 0, 1, 0, 0, 0, 0}

	for i, contains := range expected {
		if contains == 1 {
			assert.Contains(t, selection, i)
		} else {
			assert.NotContains(t, selection, i)
		}
	}
}

func TestKnapsackBig(t *testing.T) {
	selection := Knapsack(
		750,
		[]int{70, 73, 77, 80, 82, 87, 90, 94, 98, 106, 110, 113, 115, 118, 120},
		[]int{135, 139, 149, 150, 156, 163, 173, 184, 192, 201, 210, 214, 221, 229, 240},
	)

	expected := []int{1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 0, 0, 1, 1}

	for i, contains := range expected {
		if contains == 1 {
			assert.Contains(t, selection, i)
		} else {
			assert.NotContains(t, selection, i)
		}
	}
}

func TestKnapsackBigger(t *testing.T) {
	t.SkipNow()

	selection := Knapsack(
		6404180,
		[]int{
			382745, 799601, 909247, 729069, 467902, 44328, 34610, 698150,
			823460, 903959, 853665, 551830, 610856, 670702, 488960, 951111,
			323046, 446298, 931161, 31385, 496951, 264724, 224916, 169684},
		[]int{
			825594, 1677009, 1676628, 1523970, 943972, 97426, 69666, 1296457,
			1679693, 1902996, 1844992, 1049289, 1252836, 1319836, 953277, 2067538,
			675367, 853655, 1826027, 65731, 901489, 577243, 466257, 369261},
	)

	expected := []int{1, 1, 0, 1, 1, 1, 0, 0, 0, 1, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1}

	for i, contains := range expected {
		if contains == 1 {
			assert.Contains(t, selection, i)
		} else {
			assert.NotContains(t, selection, i)
		}
	}
}
