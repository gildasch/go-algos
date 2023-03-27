package priorityqueue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddAndPop(t *testing.T) {
	pq := &PriorityQueue[int, int]{}

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
	assert.Equal(t, 124, i)
	assert.Equal(t, 10, p)

	i, p = pq.Pop()
	fmt.Println("Apres pop2:", pq)
	assert.Equal(t, 126, i)
	assert.Equal(t, 5, p)

	i, p = pq.Pop()
	fmt.Println("Apres pop3:", pq)
	assert.Equal(t, 125, i)
	assert.Equal(t, 3, p)
}
