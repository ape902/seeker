package format

func Int32ToIntArray(i32 []int32) []int {
	var data []int

	for i := 0; i < len(i32); i++ {
		data = append(data, int(i32[i]))
	}

	return data
}
