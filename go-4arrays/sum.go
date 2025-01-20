package main

func Sum(numbers []int) int{
	sum := 0
	for _, number := range numbers{
		sum+=number
	}
	return sum
}

func SumAll(arrays ...[]int) []int{
	var finalArray []int

	for _,currentArray := range arrays{
		finalArray = append(finalArray, Sum(currentArray))
	}
	return finalArray
}

func SumAllTails(arrays ...[]int) []int{
	var finalArray []int

	for _,currentArray := range arrays{
		if len(currentArray) == 0{
			finalArray = append(finalArray, 0)
		}else{
			tail:= currentArray[1:]
			finalArray = append(finalArray, Sum(tail))
		}
	}
	return finalArray
}