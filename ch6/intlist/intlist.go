package intlist

type IntList struct {
	Value int
	Next  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	//递归
	return list.Value + list.Next.Value
}
