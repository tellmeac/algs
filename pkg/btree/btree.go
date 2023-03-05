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
	if len(s) == 0 || s[i] == nil {
		return nil
	}

	node := &Node{
		V: *s[i],
	}

	if j := 2*i + 1; j < len(s) && s[j] != nil {
		node.L = fromSlice(s, j)
	}

	if j := 2*i + 2; j < len(s) && s[j] != nil {
		node.R = fromSlice(s, j)
	}

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
