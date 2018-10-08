package gofunkex

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type score struct {
	class string
	name string
	course string
	score int
}

func Test_IsMap(t *testing.T) {
	sampleGroup := make(map[int][]int)
	falseGroup := make(map[int]int)
	assert.True(t, isGroup(sampleGroup))
	assert.False(t, isGroup([]int{}))
	assert.False(t, isGroup(""))
	assert.False(t, isGroup(falseGroup))
}

func Test_Sums(t *testing.T) {
	input := getGroup()
	groupFunk := NewGroupFunk(input)
	sums := groupFunk.Sums().Map.(map[string]float64)
	assert.Equal(t, 2, len(sums))
	assert.Equal(t, float64(6), sums["a"])
	assert.Equal(t, float64(4), sums["b"])
}

func Test_MapGroup(t *testing.T) {
	input := getGroup()
	groupFunk := NewGroupFunk(input)
	stringMap := groupFunk.MapGroup(func(i int)string{return strconv.Itoa(i)}).Map.(map[string][]string)
	assert.Equal(t, 2, len(stringMap))
	assert.ElementsMatch(t, []string{"2", "4"}, stringMap["a"])
	assert.ElementsMatch(t, []string{"1", "3"}, stringMap["b"])
}

func Test_Group_Map_Sum(t *testing.T) {
	scores := []score{{"C1", "Tom", "English", 100},
					  {"C2", "Jack", "Math", 100},
					  {"C1", "Peter", "Science", 100}}
	scoreSums := NewSliceFunk(scores).GroupBy(func(s score)string{return s.class}).
									  MapGroup(func(s score)int{return s.score}).
									  Sums().
										Map.(map[string]float64)
	assert.Equal(t, 2, len(scoreSums))
	assert.Equal(t, float64(200), scoreSums["C1"])
	assert.Equal(t, float64(100), scoreSums["C2"])
}

func getGroup() map[string][]int {
	input := make(map[string][]int)
	input["a"] = []int{2, 4}
	input["b"] = []int{1, 3}
	return input
}
