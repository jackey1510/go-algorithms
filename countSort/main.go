package main

func main() {

}

type sortBy []int32

func (a sortBy) Len() int           { return len(a) }
func (a sortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortBy) Less(i, j int) bool { return a[i] < a[j] }

func (a sortBy) median() float64 {
	if a.Len() == 1 {
		return float64(a[0])
	}
	middle := a.Len() / 2
	if a.Len()%2 == 0 {
		return float64(a[middle]+a[middle+1]) / float64(2)
	}
	return float64(a[middle])

}

func activityNotifications(expenditure []int32, d int32) int32 {
	// Write your code here
	//  sb := append(sortBy{}, expenditure[:d]...)
	//  sort.Sort(sb)
	noOfNotifications := int32(0)

	sb := sortBy(expenditure[:d-1])
	// fmt.Println(sb)

	for i := d; i < int32(len(expenditure)); i++ {
		sb = append(sb, expenditure[i-1])
		// fmt.Println(sb)
		sb = countSort(sb)
		// fmt.Println(sb)
		median := sb.median()
		// fmt.Println(expenditure[i], median)
		if float64(expenditure[i]) >= 2*median {
			noOfNotifications++
		}
		sb = sb[1:]
	}
	return noOfNotifications
}

func countSort(s sortBy) sortBy {
	m := make(map[int32]int32)
	max := s[0]
	for _, v := range s {
		m[v]++
		if v > max {
			max = v
		}
	}
	result := sortBy{}

	for i := int32(0); i <= max; i++ {
		count := m[i]
		for count > 0 {
			result = append(result, i)
			count--
		}
	}
	return result
}
