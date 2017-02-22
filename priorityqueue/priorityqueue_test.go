package priorityqueue

import (
	"fmt"
	"testing"

	"github.com.bak/stretchr/testify/assert"
)

func TestAddAndPop(t *testing.T) {
	var pq PriorityQueue

	pq.Add(123, 1)
	fmt.Println(pq)
	pq.Add(124, 10)
	fmt.Println(pq)
	pq.Add(125, 3)
	fmt.Println(pq)
	pq.Add(126, 5)
	fmt.Println(pq)
	pq.Add(127, 2)
	fmt.Println(pq)

	i, p := pq.Pop()
	fmt.Println("Apres pop1:", pq)
	assert.Equal(t, i, 124)
	assert.Equal(t, p, 10)

	i, p = pq.Pop()
	fmt.Println("Apres pop2:", pq)
	assert.Equal(t, i, 126)
	assert.Equal(t, p, 5)

	i, p = pq.Pop()
	fmt.Println("Apres pop3:", pq)
	assert.Equal(t, i, 125)
	assert.Equal(t, p, 3)
}
