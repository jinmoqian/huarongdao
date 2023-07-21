package game

import "testing"

func Test_makeTree_and_treeArray(t *testing.T) {
	var n *TreeNode
	n = makeTree([]BoardHash{})
	if n != nil {
		t.Error(n)
	}
	ids := treeArray(n)
	if len(ids) != 0 {
		t.Error(ids)
	}

	n = makeTree([]BoardHash{{hash: 1}})
	if n == nil || n.Left != nil || n.Right != nil || n.Uint32Val.hash != 1 {
		t.Error(n)
	}
	ids = treeArray(n)
	if len(ids) != 1 || ids[0].hash != 1 {
		t.Error(ids)
	}

	n = makeTree([]BoardHash{{hash: 1}, {hash: 2}})
	if n == nil || n.Left == nil || n.Right != nil || n.Uint32Val.hash != 2 ||
		n.Left.Left != nil || n.Left.Right != nil || n.Left.Uint32Val.hash != 1 {
		t.Error(n)
	}
	ids = treeArray(n)
	if len(ids) != 2 || ids[0].hash != 2 || ids[1].hash != 1 {
		t.Error(ids)
	}

	n = makeTree([]BoardHash{{hash: 1}, {hash: 2}, {hash: 3}})
	if n == nil || n.Left == nil || n.Right == nil || n.Uint32Val.hash != 2 ||
		n.Left.Left != nil || n.Left.Right != nil || n.Left.Uint32Val.hash != 1 ||
		n.Right.Left != nil || n.Right.Right != nil || n.Right.Uint32Val.hash != 3 {
		t.Error(n)
	}
	ids = treeArray(n)
	if len(ids) != 3 || ids[0].hash != 2 || ids[1].hash != 1 || ids[2].hash != 3 {
		t.Error(ids)
	}

	n = makeTree([]BoardHash{{hash: 1}, {hash: 2}, {hash: 3}, {hash: 4}})
	if n == nil || n.Left == nil || n.Right == nil || n.Uint32Val.hash != 3 ||
		n.Left.Left == nil || n.Left.Right != nil || n.Left.Uint32Val.hash != 2 ||
		n.Right.Left != nil || n.Right.Right != nil || n.Right.Uint32Val.hash != 4 ||
		n.Left.Left.Left != nil || n.Left.Left.Right != nil || n.Left.Left.Uint32Val.hash != 1 {
		t.Error(n)
	}
	ids = treeArray(n)
	if len(ids) != 4 || ids[0].hash != 3 || ids[1].hash != 2 || ids[2].hash != 4 || ids[3].hash != 1 {
		t.Error(ids)
	}

	n = makeTree([]BoardHash{{hash: 1}, {hash: 2}, {hash: 3}, {hash: 4}, {hash: 5}})
	if n == nil || n.Left == nil || n.Right == nil || n.Uint32Val.hash != 4 ||
		n.Left.Left == nil || n.Left.Right == nil || n.Left.Uint32Val.hash != 2 ||
		n.Right.Left != nil || n.Right.Right != nil || n.Right.Uint32Val.hash != 5 ||
		n.Left.Left.Left != nil || n.Left.Left.Right != nil || n.Left.Left.Uint32Val.hash != 1 ||
		n.Left.Right.Left != nil || n.Left.Right.Right != nil || n.Left.Right.Uint32Val.hash != 3 {
		t.Error(n)
	}
	ids = treeArray(n)
	if len(ids) != 5 || ids[0].hash != 4 || ids[1].hash != 2 || ids[2].hash != 5 || ids[3].hash != 1 || ids[4].hash != 3 {
		t.Error(ids)
	}

	n = makeTree([]BoardHash{{hash: 1}, {hash: 2}, {hash: 3}, {hash: 4}, {hash: 5}, {hash: 6}})
	if n == nil || n.Left == nil || n.Right == nil || n.Uint32Val.hash != 4 ||
		n.Left.Left == nil || n.Left.Right == nil || n.Left.Uint32Val.hash != 2 ||
		n.Right.Left == nil || n.Right.Right != nil || n.Right.Uint32Val.hash != 6 ||
		n.Left.Left.Left != nil || n.Left.Left.Right != nil || n.Left.Left.Uint32Val.hash != 1 ||
		n.Left.Right.Left != nil || n.Left.Right.Right != nil || n.Left.Right.Uint32Val.hash != 3 ||
		n.Right.Left.Left != nil || n.Right.Left.Right != nil || n.Right.Left.Uint32Val.hash != 5 {
		t.Error(n)
	}
	ids = treeArray(n)
	if len(ids) != 6 || ids[0].hash != 4 || ids[1].hash != 2 || ids[2].hash != 6 || ids[3].hash != 1 || ids[4].hash != 3 || ids[5].hash != 5 {
		t.Error(ids)
	}

	n = makeTree([]BoardHash{{hash: 1}, {hash: 2}, {hash: 3}, {hash: 4}, {hash: 5}, {hash: 6}, {hash: 7}})
	if n == nil || n.Left == nil || n.Right == nil || n.Uint32Val.hash != 4 ||
		n.Left.Left == nil || n.Left.Right == nil || n.Left.Uint32Val.hash != 2 ||
		n.Right.Left == nil || n.Right.Right == nil || n.Right.Uint32Val.hash != 6 ||
		n.Left.Left.Left != nil || n.Left.Left.Right != nil || n.Left.Left.Uint32Val.hash != 1 ||
		n.Left.Right.Left != nil || n.Left.Right.Right != nil || n.Left.Right.Uint32Val.hash != 3 ||
		n.Right.Left.Left != nil || n.Right.Left.Right != nil || n.Right.Left.Uint32Val.hash != 5 ||
		n.Right.Right.Left != nil || n.Right.Right.Right != nil || n.Right.Right.Uint32Val.hash != 7 {
		t.Error(n)
	}
	ids = treeArray(n)
	if len(ids) != 7 || ids[0].hash != 4 || ids[1].hash != 2 || ids[2].hash != 6 || ids[3].hash != 1 || ids[4].hash != 3 || ids[5].hash != 5 || ids[6].hash != 7 {
		t.Error(ids)
	}

	n = makeTree([]BoardHash{{hash: 1}, {hash: 2}, {hash: 3}, {hash: 4}, {hash: 5}, {hash: 6}, {hash: 7}, {hash: 8}})
	if n == nil || n.Left == nil || n.Right == nil || n.Uint32Val.hash != 5 ||
		n.Left.Left == nil || n.Left.Right == nil || n.Left.Uint32Val.hash != 3 ||
		n.Right.Left == nil || n.Right.Right == nil || n.Right.Uint32Val.hash != 7 ||
		n.Left.Left.Left == nil || n.Left.Left.Right != nil || n.Left.Left.Uint32Val.hash != 2 ||
		n.Left.Right.Left != nil || n.Left.Right.Right != nil || n.Left.Right.Uint32Val.hash != 4 ||
		n.Right.Left.Left != nil || n.Right.Left.Right != nil || n.Right.Left.Uint32Val.hash != 6 ||
		n.Right.Right.Left != nil || n.Right.Right.Right != nil || n.Right.Right.Uint32Val.hash != 8 ||
		n.Left.Left.Left.Left != nil || n.Left.Left.Left.Right != nil || n.Left.Left.Left.Uint32Val.hash != 1 {
		t.Error(n)
	}
	ids = treeArray(n)
	if len(ids) != 8 || ids[0].hash != 5 || ids[1].hash != 3 || ids[2].hash != 7 || ids[3].hash != 2 || ids[4].hash != 4 || ids[5].hash != 6 || ids[6].hash != 8 || ids[7].hash != 1 {
		t.Error(ids)
	}
}
