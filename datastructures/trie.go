/*
Created by Todd Chaney
A implementation of a trie datastructure, needs more work to improve it and generalize. 
This is specifically focused on count pairs of xor for 16 bit integers. 
*/
type data struct {
    count int
    bit int
}

type node struct {
    edges map[int]*node
    data 
}

func (cur *node) add(x int) {
    var nextNode *node
    for i:=15;i>=0;i-- {
        offset := 1<<i
        if (x&offset)>0 {
            if cur.hasEdge(1) {
                nextNode = cur.getEdge(1).(*node)
                nextNode.data.count++
            } else {
                nextNode = &node{map[int]&node{}, data{1,0}}
                cur.edges = append(cur.edges, nextNode)
            }
        } else {
            if cur.hasEdge(0) {
                nextNode = cur.getEdge(0).(*node)
                nextNode.data.count++
            } else {
                nextNode = &node{map[int]&node{},data{1,1}}
                cur.edges = append(cur.edges, nextNode)
            }
        }
        cur = nextNode
    }
}

func (n *node) hasEdge(digit int) bool {
    _, found := n.edges[digit]
    return found
}

func (n *node) getEdge(digit int) interface{} {
    return n.edges[digit]
}

func (cur *node) search(x int, lim int) int {
    cnt := 0
    for i:=15;i>=0;i-- {
        offset := 1<<i
        if (lim&offset)>0 {
            if (x&offset)>0 {
                if cur.hasEdge(1) {
                    cnt+=cur.edges[1].data.count
                }
                if cur.hasEdge(0) {
                    cur = cur.getEdge(0)
                } else {
                    break
                }
            } else {
                if cur.hasEdge(0) {
                    cnt+=cur.edges[0].data.count
                }
                if cur.hasEdge(1) {
                    cur = cur.getEdge(1)
                } else {
                    break
                }
            }
        } else {
            if (x&offset)>0 {
                if cur.hasEdge(1) {
                    cur = cur.getEdge(1)
                } else {
                    break
                }
            } else {
                if cur.hasEdge(0) {
                    cur = cur.getEdge(0)
                } else {
                    break
                }
            }
        }
    }
    return cnt
}


func solve(nums *[]int, lim int) int {
    n, cnt := len(*nums), 0
    root := &node{map[int]&node{}, data{0,0}}
    for i:=n-1;i>=0;i-- {
        cnt += root.search((*nums)[i], lim)
        root.add((*nums)[i])
        root.data.count++
    }
    return cnt
}

func countPairs(nums []int, low int, high int) int {
    return solve(&nums, high+1) - solve(&nums,low)
}