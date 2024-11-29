package day15

import (
	"strconv"
	"strings"

	"github.com/iancoleman/orderedmap"
)

func Hashing(input string) int {
	result := 0
	for _, c := range input {
		result += int(c)
		result *= 17
		result %= 256
	}
	return result
}

func solvePartA(input string) int {
	parts := strings.Split(input, ",")
	result := 0
	for _, part := range parts {
		result += Hashing(part)
	}
	return result
}

type LensBoxes []orderedmap.OrderedMap

func NewLensBoxes(input string) LensBoxes {
	lensBoxes := make([]orderedmap.OrderedMap, 256)
	for i := 0; i < 256; i++ {
		lensBoxes[i] = *orderedmap.New()
	}

	parts := strings.Split(input, ",")
	for _, part := range parts {
		if strings.Contains(part, "-") {
			id := strings.Split(part, "-")[0]
			box := Hashing(id)
			lensBoxes[box].Delete(id)
		} else {
			r := strings.Split(part, "=")
			id := r[0]
			value, _ := strconv.Atoi(r[1])
			box := Hashing(id)
			lensBoxes[box].Set(id, value)
		}
	}
	// for i, box := range lensBoxes {
	// 	if len(box.Keys()) > 0 {
	// 		fmt.Println(i, box.Keys(), box.Values())
	// 	}
	// }
	return lensBoxes
}

func FocusPower(boxId, slot, focalLength int) int {
	return (boxId + 1) * slot * focalLength
}

func solvePartB(input string) int {
	lensBoxes := NewLensBoxes(input)
	power := 0
	for boxId, box := range lensBoxes {
		for slot, key := range box.Keys() {
			value, _ := box.Get(key)
			power += FocusPower(boxId, slot+1, value.(int))
		}
	}
	return power
}
