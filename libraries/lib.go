/*
Created by Todd Chaney 

This is start of a library and API that is used in competitive programming, it contains
general functions that are necessary frequently. And useful information or that will be somewhere else. 
TBD
*/

// The min function acting on integers
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// The max function acting on integers
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// The absolute value acting on integers
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Special mod to give positive results for when a<0
func modulus(a, b int) int {
    return (a + b)%b
}

// Sum of all elements in an array
func sum(nums []int) int {
    res := 0
    for _, x := range nums {
        res+=x
    }
    return res
}

// Euclidean algorithms for the greatest common denominator
func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    r := a%b
    a, b = b, r
    return gcd(a,b)
}

// Find factors in squrt(n) TC, 
func factors(x int) []int {
    var res []int
    for i:=1;i*i<=x;i++ {
        if x%i==0 {
            f1:=i
            f2:=x/i
            res = append(res, f1,f2)
        }
    }
    return res
}

// convert an array into a string that can be used as key in map in golang.
func getKey(arr []int) string {
    var b []byte
    for i:=0;i<len(arr);i++ {
        b = append(b, byte(i), ':', byte(arr[i]))
    }
    return string(b)
}

/*
A general template for heaps or priority queue containing some node.
*/
import "container/heap"

// node could contain any fields useful for the heap
type node struct {
    u int
    dist int
}

type distHeap []node
// The all important functions for the heap, so it can push elements to it.
func (h distHeap) Len() int {return len(h)}
func (h distHeap) Less(i, j int) bool {return h[i].dist<h[j].dist}
func (h distHeap) Swap(i, j int) {h[i],h[j]=h[j],h[i]}
// Push to the heap
func (h *distHeap) Push(x interface{}) {
    *h = append(*h, x.(node))
}
// Pop from the heap
func (h *distHeap) Pop() interface{} {
    n := len(*h)
    x := (*h)[n-1]
    *h = (*h)[:n-1]
    return x
}

// Fill in an array by reference with a specific val
func fill(dp *[]int, val int) {
    memo := *dp
    for i:=0;i<len(memo);i++{
        memo[i]=val
    }
}

// Reverse string by passing byte array by reference
func reverse(s *[]byte) {
    st := *s
    n := len(st)
    i, j := 0, n-1
    for i < j {
        st[i],st[j] = st[j], st[i]
        i++
        j--
    }
}

/*
Just a useful trick for lexicographically sorting strings 
sort.Strings(words)
*/
