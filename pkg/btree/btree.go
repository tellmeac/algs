package btree

type Node struct {
	V int
	L *Node
	R *Node
}

func FromSlice(s []*int) *Node {
	return fromSlice(s, 0)
}

func fromSlice(s []*int, i int) *Node {
	if len(s) == 0 || i >= len(s) || s[i] == nil {
		return nil
	}

	node := &Node{
		V: *s[i],
	}

	node.L = fromSlice(s, 2*i+1)
	node.R = fromSlice(s, 2*i+2)

	return node
}

func InOrderDFS(root *Node) []int {
	var (
		result = make([]int, 0)
		dfs    func(*Node)
	)

	dfs = func(r *Node) {
		if r == nil {
			return
		}

		dfs(r.L)

		result = append(result, r.V)

		dfs(r.R)
	}

	dfs(root)

	return result
}

func PreOrderDFS(root *Node) []int {
	panic("")
}

func PostOrderDFS(root *Node) []int {
	panic("")
}
