package main

import (
    "fmt"
    "sort"
    "strconv"
    "strings"
)

type c struct {
    i     int
    s, rs string
}

type cc []*c

func (c cc) Len() int           { return len(c) }
func (c cc) Less(i, j int) bool { return c[j].rs < c[i].rs }
func (c cc) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

// Function required by task.  Takes a list of integers, returns big int.
func li(is ...int) string {
    ps := make(cc, len(is))
    ss := make([]c, len(is))
    ml := 0
    for j, i := range is {
        p := &ss[j]
        ps[j] = p
        p.i = i
        p.s = strconv.Itoa(i)
        if len(p.s) > ml {
            ml = len(p.s)
        }
    }
    for _, p := range ps {
        p.rs = strings.Repeat(p.s, (ml+len(p.s)-1)/len(p.s))
    }
    sort.Sort(ps)
    s := make([]string, len(ps))
    for i, p := range ps {
        s[i] = p.s
	}
	var b string
	for i:=0;i<len(s);i++ {
		b+=s[i]
	}
    
    return b
}

func main() {
    fmt.Println(li(1, 34, 3, 98, 9, 76, 45, 4))
	fmt.Println(li(54, 546, 548, 60))
	fmt.Printf("Type:%T",li(54, 546, 548, 60))
}