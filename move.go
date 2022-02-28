package gomovearrposition

func MoveByCountFront(array []int, k int) []int {
	var newArr []int
	kn := 1
	for i := 0; i < len(array); i++ {
		var result int
		if i < (len(array) - k) {
			result = array[i]
		} else {
			result = array[len(array)-kn]
			kn += 1
		}
		newArr = append(newArr, result)
	}
	return newArr
}
