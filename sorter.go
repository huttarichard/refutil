package refutil

import "sort"

// ValueSorter will sort values in slice
// this is good for searching maps and structs
// which don't have to preserve order
type ValueSorter []Value

// Sort will sort values in slice
func (sorter ValueSorter) Sort() {
	sort.Sort(sorter)
}

// Len return ren of slice
func (sorter ValueSorter) Len() int {
	return len(sorter)
}

// Less will check if value is 'less' then other
func (sorter ValueSorter) Less(i, j int) bool {
	return sorter[i].String() < sorter[j].String()
}

// Swap performs swaping of values in slice
func (sorter ValueSorter) Swap(i, j int) {
	sorter[i], sorter[j] = sorter[j], sorter[i]
}
