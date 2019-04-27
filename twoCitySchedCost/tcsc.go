package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

func main() {
	costs := [][]int{}
	str := "[[10,20],[30,200],[400,50],[30,20]]"
	json.Unmarshal([]byte(str), &costs)

	fmt.Println(twoCitySchedCost(costs))

	str = "[[623,388],[481,171],[306,200],[498,347],[950,282],[688,472],[343,777],[995,173],[716,748],[36,410],[50,540],[769,138],[873,453],[589,787],[189,169],[813,400],[22,156],[388,410],[880,265],[800,619],[604,667],[106,338],[469,501],[394,930],[780,323],[225,378],[229,391],[102,430],[872,125],[378,316],[816,998],[179,49],[967,473],[148,426],[8,14],[308,339],[356,473],[195,99],[382,549],[120,342],[199,382],[5,762],[202,726],[689,874],[177,897],[981,758],[494,744],[792,433],[275,306],[555,882],[475,791],[102,779],[454,267],[257,413],[737,401],[300,562],[276,365],[975,850],[218,121],[922,911],[852,769],[31,472],[623,959],[475,258],[945,599],[929,538],[664,284],[94,420],[499,338],[814,312],[935,2],[14,87],[963,377],[936,978],[692,448],[754,380],[378,621],[351,963]]"
	//[10,30],[30,50],[40,60]
	json.Unmarshal([]byte(str), &costs)
	fmt.Println(twoCitySchedCost(costs))

	str = "[[623,388],[481,171],[306,200],[498,347],[950,282],[688,472],[343,777],[995,173],[716,748],[36,410],[50,540],[769,138],[873,453],[589,787],[189,169],[813,400],[22,156],[388,410],[880,265],[800,619]]"
	//[10,30],[30,50],[40,60]
	json.Unmarshal([]byte(str), &costs)
	fmt.Println(twoCitySchedCost(costs))

	// tmp := []int{0, -1, 2, 2, 2, 3, 4, 1, -2, 3}
	// sort.Ints(tmp)
	// fmt.Println(tmp)
}

func twoCitySchedCost(costs [][]int) int {
	var res int
	fmt.Println(costs)
	costs = sortArr(costs)
	fmt.Println(costs)
	//middle := middle(costs)
	l := len(costs)
	for i := 0; i < l; i++ {
		if i <= l/2 {
			res += costs[i][0]
		} else {
			res += costs[i][1]
		}
	}

	return res
}

func middle(costs [][]int) int {
	temp := []int{}
	for i := 0; i < len(costs); i++ {
		temp = append(temp, (costs[i][0] - costs[i][1]))
	}
	sort.Ints(temp)
	fmt.Printf("%v\nlen(%d),%d\n", temp, len(temp), temp[len(temp)/2-1])
	return temp[len(temp)/2-1]
}

func sortArr(arr [][]int) [][]int {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if arr[i][0] > arr[j][0] {
				arr[i], arr[j] = arr[j], arr[i]
			} else if arr[i][0] == arr[j][0] && arr[i][1] > arr[j][1] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
