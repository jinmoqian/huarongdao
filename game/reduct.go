package game

func Reduct(boardHashes map[BoardHash]*BoardLinkNode, ratio, batch float64, reductIndex func(allEdges map[BoardHash]struct{}) *BoardHash, batchAdjust func(incr bool, ratio float64, reducts int) float64) map[BoardHash]*BoardLinkNode {
	if ratio >= 1.0 {
		return boardHashes
	}
	max := int(float64(len(boardHashes)) * ratio)
	allEdges := make(map[BoardHash]struct{})
	var allReduct, allEdge, emptyLoopCounter int
	for {
		if emptyLoopCounter >= 100 {
			break
		}
		idx := reductIndex(allEdges)
		if idx == nil {
			break
		}
		reducts := make(map[BoardHash]struct{})
		reductNode := boardHashes[*idx]
		head := reductNode
		tail := head
		head.Next = nil
		for {
			if head == nil {
				break
			}
			if _, ok := allEdges[head.Uint32Val]; !ok {
				for _, next := range head.Nexts {
					if _, ok := allEdges[next]; ok {
						continue
					}
					if _, ok := reducts[next]; ok {
						continue
					}
					// fmt.Println(boardHashes[head.Uint32Val].Board, "=>", boardHashes[next].Board)
					reducts[next] = struct{}{}
					tail.Next = boardHashes[next]
					tail = tail.Next
					tail.Next = nil
				}
				reducts[head.Uint32Val] = struct{}{}
				if float64(len(reducts)) >= batch {
					break
				}
			}
			head = head.Next
		}
		if len(reducts) == 0 {
			emptyLoopCounter++
			continue
		}
		emptyLoopCounter = 0
		allReduct += len(reducts)
		for r := range reducts {
			for _, next := range boardHashes[r].Nexts {
				if _, ok := reducts[next]; !ok {
					if _, ok := allEdges[next]; !ok {
						allEdges[next] = struct{}{}
						allEdge++
					}
				}
			}
		}
		for r := range reducts {
			delete(boardHashes, r)
		}
		r := float64(allEdge) / float64(allReduct+allEdge)
		if r > ratio {
			v := batchAdjust(true, r, len(reducts))
			// fmt.Println("Batch1=", batch, "=>", v, "reducts=", len(reducts), "allEdges=", len(allEdges))
			batch = v
			// batch++
		} else {
			v := batchAdjust(false, r, len(reducts))
			// fmt.Println("Batch2=", batch, "=>", v, "reducts=", len(reducts), "allEdges=", len(allEdges))
			batch = v
			// batch--
		}
		if len(boardHashes) <= max {
			break
		}
	}
	return boardHashes
}

// func Search(depth map[BoardHash]byte, g *Game, b *Board) []*Board {
// 	localDepth := make(map[BoardHash]byte)
// 	path := make([]*Board, 0)
// 	for {

// 		_, next := calcDepth(depth, g, b, localDepth)
// 		path = append(path, next)
// 		if next.Win() {
// 			break
// 		}
// 		b = next
// 	}
// 	return path
// }
// func calcDepth(depth map[BoardHash]byte, g *Game, b *Board, localDepth map[BoardHash]byte) (byte, *Board) {
// 	dependicies := make(map[BoardHash]map[BoardHash]struct{})
// 	revDependicies := make(map[BoardHash]map[BoardHash]struct{})
// 	type queueNode struct {
// 		b     *Board
// 		next  *queueNode
// 		depth byte
// 	}
// 	head := &queueNode{b: b}
// 	// tail := head
// 	for {
// 		if head == nil {
// 			break
// 		}
// 		movables := g.Movable(b)
// 		headHash := g.Hash(head.b)
// 		for _, movable := range movables {
// 			nexts := g.Move(b, movable)
// 			for _, next := range nexts {
// 				var d byte
// 				var ok bool = true
// 				h := g.Hash(next)
// 				if !next.Win() {
// 					if d, ok = depth[h]; !ok {
// 						if d, ok = localDepth[h]; !ok {
// 							if dependicies[headHash] == nil {
// 								dependicies[headHash] = make(map[BoardHash]struct{})
// 							}
// 							dependicies[headHash][h] = struct{}{}
// 							if revDependicies[h] == nil {
// 								revDependicies[h] = make(map[BoardHash]struct{})
// 							}
// 							revDependicies[h][headHash] = struct{}{}
// 						}
// 					}
// 				} else {
// 					localDepth[h] = 0
// 				}
// 				if ok {
// 					type triggernode struct {
// 						h    BoardHash
// 						next *triggernode
// 					}
// 					triggerHead := &triggernode{
// 						h: h,
// 					}
// 					triggerTail := triggerHead
// 					for {
// 						if triggerHead == nil {
// 							break
// 						}
// 						for depend := range revDependicies[triggerHead.h] {
// 							var all bool = true
// 							var minDepth byte = 0xFF
// 							for dep := range dependicies[depend] {
// 								if d, ok = depth[dep]; !ok {
// 									if d, ok = localDepth[dep]; !ok {
// 										all = false
// 										break
// 									}
// 								}
// 								if d < minDepth {
// 									minDepth = d
// 								}
// 							}
// 							if all {
// 								localDepth[depend] = minDepth + 1
// 								triggerTail.next = &triggernode{h: depend}
// 								triggerTail = triggerTail.next
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 		head = head.next
// 	}
// 	return 0, nil
// }
