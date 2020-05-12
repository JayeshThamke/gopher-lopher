package listops

// IntList - custom data type
type IntList []int

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Foldl - refer README.md
func (l IntList) Foldl(fn binFunc, acc int) int {
	for i := 0; i < len(l); i++ {
		acc = fn(acc, l[i])
	}
	return acc
}

// Foldr - refer README.md
func (l IntList) Foldr(fn binFunc, acc int) int {
	for i := len(l) - 1; i >= 0; i-- {
		acc = fn(l[i], acc)
	}
	return acc
}

// Append - refer README.md
func (l *IntList) Append(tobeAppend IntList) IntList {
	return append(*l, tobeAppend...)
}

// Concat - refer README.md
func (l *IntList) Concat(intLists []IntList) IntList {
	concatinatedList := *l
	for i := 0; i <= len(intLists)-1; i++ {
		concatinatedList = append(concatinatedList, intLists[i]...)
	}

	if len(concatinatedList) == 0 {
		return []int{}
	}

	return concatinatedList
}

// Reverse - refer README.md
func (l IntList) Reverse() IntList {
	midway := len(l)/2 - 1
	afterMidway := len(l) - 1

	for i := midway; i >= 0; i-- {
		j := afterMidway - i

		l[i], l[j] = l[j], l[i]
	}
	return l
}

// Length - refer README.md
func (l *IntList) Length() int {
	return len(*l)
}

// Filter - refer README.md
func (l *IntList) Filter(fn predFunc) IntList {
	var filteredList IntList
	for _, i := range *l {
		if fn(i) {
			filteredList = append(filteredList, i)
		}
	}
	if len(filteredList) == 0 {
		return []int{}
	}
	return filteredList
}

// Map - refer README.md
func (l *IntList) Map(fn unaryFunc) IntList {
	var mappedList IntList

	for _, item := range *l {
		mappedList = append(mappedList, fn(item))
	}
	if len(mappedList) == 0 {
		return []int{}
	}
	return mappedList
}
