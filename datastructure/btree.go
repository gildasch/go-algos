package datastructure

// Access	Search	Insertion	Deletion

const order = 5

type BTree []int

func (b BTree) Get(n int) int {
	return b.get(n, 0)
}

func (b BTree) get(n, i int) int {
	children := b.children(i)
	if len(children) == 0 {
		if n >= order-1 {
			return 0
		}

		return b[i+n]
	}

	left := 0
	for k, c := range children {
		l := left + b.len(c)
		if n < l {
			return b.get(n-left, c)
		}
		if n == l {
			return b[k]
		}
		left += b.len(c)
	}

	return 0
}

func (b BTree) len(i int) int {
	var len int

	for k := 0; k < order-1; k++ {
		if b[i+k] != 0 {
			len++
		}
	}

	for _, c := range b.children(i) {
		len += b.len(c)
	}

	return len
}

func (b BTree) children(i int) []int {
	var children []int
	for k := 0; k < order; k++ {
		position := i + order - 1 + k*(order-1)
		if position >= len(b) || b[position] == 0 {
			break
		}

		children = append(children, position)
	}

	// fmt.Println(children)
	return children
}
