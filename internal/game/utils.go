package game

import (
	"errors"
	"math/rand"
)

func genRandomList(length int) ([]int, error) {
	if length <= 0 {
		return nil, errors.New("the size of the parameter length illegal")
	}
	var list []int
	for i := 0; i < length; i++ {
		list = append(list, i)
	}
	for i := len(list) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		list[i], list[num] = list[num], list[i]
	}
	return list, nil
}