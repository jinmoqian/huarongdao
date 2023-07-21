package game

const NoDepth = 0xFF

type DepthGetter interface {
	Depth(BoardHash) byte
}

func Search(depth DepthGetter, g *Game, b *Board) []*Board {
	maxDepth := depth.Depth(b.game.Hash(b))
	if maxDepth == NoDepth {
		return BFS(b, g)
	}
	ret := make([]*Board, 0)
	ret = append(ret, b)
	for {
		if b.Win() {
			return ret
		}
		movables := g.Movable(b)
		var minDepth byte = 255
		var minBoard *Board
		for _, move := range movables {
			nexts := g.Move(b, move)
			for _, next := range nexts {
				nh := g.Hash(next)
				nd := depth.Depth(nh)
				if nd == NoDepth {
					paths := BFS(next, g)
					if paths == nil {
						continue
					}
					bDepth := depth.Depth(b.game.Hash(b))
					nd = byte(len(paths))
					if bDepth == nd {
						ret = append(ret, paths...)
						return ret
					}
					nd = NoDepth
				}
				if nd < minDepth {
					minDepth = nd
					minBoard = next
				}
			}
		}
		if minBoard == nil {
			return nil
		}
		ret = append(ret, minBoard)
		b = minBoard
	}
}

type LinkNode struct {
	Value  *Board
	Next   *LinkNode
	Parent *LinkNode
}

func BFS(b *Board, g *Game) []*Board {
	head := &LinkNode{Value: b}
	tail := head
	uniq := make(map[BoardHash]struct{})
	uniq[g.Hash(b)] = struct{}{}
	for {
		if head == nil {
			return nil
		}
		// fmt.Println(head.Value)
		if head.Value.Win() {
			break
		}
		b = head.Value
		movables := g.Movable(b)
		for _, move := range movables {
			nexts := g.Move(b, move)
			for _, next := range nexts {
				// fmt.Println("Next=", next)
				h := g.Hash(next)
				if _, ok := uniq[h]; ok {
					// fmt.Println("BH=", h, "exists")
					continue
				}
				// fmt.Println("BH=", h)
				newNext := next.Copy()
				nn := &LinkNode{Value: newNext, Parent: head}
				tail.Next = nn
				tail = nn
				uniq[h] = struct{}{}
			}
		}
		head = head.Next
	}
	var counter int
	t := head
	for {
		counter++
		t = t.Parent
		if t == nil {
			break
		}
	}
	var i = 0
	ret := make([]*Board, counter)
	for {
		ret[counter-i-1] = head.Value
		head = head.Parent
		if head == nil {
			break
		}
		i++
	}
	return ret
}
