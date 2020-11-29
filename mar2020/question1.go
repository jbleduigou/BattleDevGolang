import("bufio")
import("container/heap")
import("strconv")

func contestResponse() {
	scanner := bufio.NewScanner(os.Stdin)
	nbLines := -1
	votes := map[string]int{}

    // Read input data 
	for scanner.Scan() {
		line := scanner.Text()
		if nbLines == -1 {
		  nbLines, _ = strconv.Atoi(line)
		} else {
		    current, found := votes[line]
		    if found {
		      current = current + 1 
		      votes[line] = current
		    } else {
				votes[line] = 1
		    }
		}
		os.Stderr.WriteString(line)
    	os.Stderr.WriteString("\n")
	}
	// Use PriorityQueue for input data
	pq := make(PriorityQueue, len(votes))
	i := 0
	for name, priority := range votes {
		pq[i] = &Colour{
			name:    name,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
	
	// Pop best and second best colours from PriorityQueue
	first := heap.Pop(&pq).(*Colour)
	second := heap.Pop(&pq).(*Colour)
	os.Stderr.WriteString("##########################\n\n")
	
	// Write result to standard output
	fmt.Println(first.name + " " + second.name)
}

// A Colour is something we manage in a priority queue.
type Colour struct {
	name    string // The name of the colour.
	priority int    // The priority of the colour in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the colour in the heap.
}

// A PriorityQueue implements heap.Interface and holds Colours.
type PriorityQueue []*Colour

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	c := x.(*Colour)
	c.index = n
	*pq = append(*pq, c)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	c := old[n-1]
	old[n-1] = nil  // avoid memory leak
	c.index = -1 // for safety
	*pq = old[0 : n-1]
	return c
}

// update modifies the priority and value of a Colour in the queue.
func (pq *PriorityQueue) update(c *Colour, name string, priority int) {
	c.name = name
	c.priority = priority
	heap.Fix(pq, c.index)
}
