/*******
 * Read input from os.Stdin
 * Use: fmt.Println to ouput your result to STDOUT.
 * Use: os.Stderr.WriteString to ouput debugging information to STDERR.
 * ***/
 import("bufio")
 import("strconv")
 import("sync")
 import("strings")
 func contestResponse() {
	 scanner := bufio.NewScanner(os.Stdin)
	 nbChar := -1
	 var dic map[int]*Node
	 dic = make(map[int]*Node)
	 g := &ItemGraph{}
	 root := &Node{value:0}
	 dic[0] = root
	 g.AddNode(root)
	 for scanner.Scan() {
		 line := scanner.Text()
		 if nbChar == -1 {
		   nbChar, _ = strconv.Atoi(line)
		 } else {
			 s := strings.Fields(line)
			 a, _ := strconv.Atoi(s[0])
			 b, _ := strconv.Atoi(s[1])
			 na, found := dic[a]
			 if !found {
				 na = &Node{value : a}
				 dic[a] = na
				 g.AddNode(na)
			 }
			 nb, found := dic[b]
			 if !found {
				 nb = &Node{value : b}
				 dic[b] = nb
				 g.AddNode(nb)
			 }
			 g.AddEdge(nb, na)
		 }
		 os.Stderr.WriteString(line)
		 os.Stderr.WriteString("\n")
	 }
	 table := make([]int,10)
	 g.Traverse(func(n *Node) {
		 os.Stderr.WriteString(fmt.Sprintf("Visiting id=%d level=%d\n", n.value, n.level))
	 })
	for _, n := range g.nodes {
		current := table[n.level]
		current = current + 1
		table[n.level] = current
	}
	 result := ""
	 for _,k := range table {
		  result = result + " " + strconv.Itoa(k) 
	 }
	 os.Stderr.WriteString("##########################\n\n")
	 fmt.Println(result)
 }
 
 type Node struct {
	 value int
	 children int
	 level int
 }
 
 // ItemGraph the Items graph
 type ItemGraph struct {
	 nodes []*Node
	 edges map[int][]*Node
	 lock  sync.RWMutex
 }
 
 
 // AddNode adds a node to the graph
 func (g *ItemGraph) AddNode(n *Node) {
	 g.lock.Lock()
	 g.nodes = append(g.nodes, n)
	 g.lock.Unlock()
	 
 }
 
 
 // AddEdge adds an edge to the graph
 func (g *ItemGraph) AddEdge(n1, n2 *Node) {
	 g.lock.Lock()
	 if g.edges == nil {
		 g.edges = make(map[int][]*Node)
	 }
	 g.edges[n1.value] = append(g.edges[n1.value], n2)
	 g.lock.Unlock()
 }
 
 
 // NodeQueue the queue of Nodes
 type NodeQueue struct {
	 items []Node
	 lock  sync.RWMutex
 }
 
 // New creates a new NodeQueue
 func (s *NodeQueue) New() *NodeQueue {
	 s.lock.Lock()
	 s.items = []Node{}
	 s.lock.Unlock()
	 return s
 }
 
 // Enqueue adds an Node to the end of the queue
 func (s *NodeQueue) Enqueue(t Node) {
	 s.lock.Lock()
	 s.items = append(s.items, t)
	 s.lock.Unlock()
 }
 
 // Dequeue removes an Node from the start of the queue
 func (s *NodeQueue) Dequeue() *Node {
	 s.lock.Lock()
	 item := s.items[0]
	 s.items = s.items[1:len(s.items)]
	 s.lock.Unlock()
	 return &item
 }
 
 // Front returns the item next in the queue, without removing it
 func (s *NodeQueue) Front() *Node {
	 s.lock.RLock()
	 item := s.items[0]
	 s.lock.RUnlock()
	 return &item
 }
 
 // IsEmpty returns true if the queue is empty
 func (s *NodeQueue) IsEmpty() bool {
	 s.lock.RLock()
	 defer s.lock.RUnlock()
	 return len(s.items) == 0
 }
 
 // Size returns the number of Nodes in the queue
 func (s *NodeQueue) Size() int {
	 s.lock.RLock()
	 defer s.lock.RUnlock()
	 return len(s.items)
 }
 
 // Traverse implements the BFS traversing algorithm
 func (g *ItemGraph) Traverse(f func(*Node)) {
	 g.lock.RLock()
	 q := NodeQueue{}
	 q.New()
	 n := g.nodes[0]
	 q.Enqueue(*n)
	 visited := make(map[int]bool)
	 for {
		 if q.IsEmpty() {
			 break
		 }
		 node := q.Dequeue()
		 visited[node.value] = true
		 near := g.edges[node.value]
 
		 for i := 0; i < len(near); i++ {
			 j := near[i]
			 j.level = node.level + 1
			 if !visited[j.value] {
				 q.Enqueue(*j)
				 visited[j.value] = true
			 } 
		 }
		 if f != nil {
			 f(node)
		 }
	 }
	 g.lock.RUnlock()
 }