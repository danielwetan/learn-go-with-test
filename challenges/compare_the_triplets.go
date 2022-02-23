package challenges

func compareTheTriplets(a, b []int32) []int32 {

	var alice, bob int32

	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			alice++
		} else if a[i] < b[i] {
			bob++
		} else {
			continue
		}
	}

	return []int32{alice, bob}
}

