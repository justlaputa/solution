package solution

import (
	"math"
	"sort"
)

type arrayElem struct {
	index int
	value int
}

type adjacentIndex struct {
	low  int
	high int
}

//Solution implements the solution
func Solution(data []int) int {
	arrayElems := []arrayElem{}
	adjacentIndexes := []adjacentIndex{}

	//build input array with index as a separate field
	//need O(n) space
	for i, v := range data {
		arrayElems = append(arrayElems, arrayElem{i, v})
	}

	//in-place sort, not stable but we don't care
	//O(nlogn) time
	sort.Slice(arrayElems, func(i, j int) bool {
		return arrayElems[i].value < arrayElems[j].value
	})

	//go through the new array, chech each neighber elements, if a pair of neighber
	//elements are not equal, they are an adjacent value
	for i, elem := range arrayElems {
		if i == len(arrayElems)-1 {
			break
		}

		j := i + 1
		for ; j < len(arrayElems) && arrayElems[j].value == elem.value; j++ {
		}

		//check i's next neightber
		if arrayElems[j].value != elem.value {
			adjacentIndexes = append(adjacentIndexes, adjacentIndex{elem.index, arrayElems[j].index})
			//check following elements, if they have the same value as (i+1), they form an adjacent value with i as well
			for k := j + 1; k < len(arrayElems) && arrayElems[k].value == arrayElems[j].value; k++ {
				adjacentIndexes = append(adjacentIndexes, adjacentIndex{elem.index, arrayElems[k].index})
			}
		}
	}

	//reverse sort by the distance
	sort.Slice(adjacentIndexes, func(i, j int) bool {
		return math.Abs(float64(adjacentIndexes[i].high-adjacentIndexes[i].low)) > math.Abs(float64(adjacentIndexes[j].high-adjacentIndexes[j].low))
	})

	maxDistance := adjacentIndexes[0].high - adjacentIndexes[0].low
	if maxDistance < 0 {
		maxDistance = -maxDistance
	}

	return maxDistance
}
