package main

import (
	"time"
	"math/rand"
	"fmt"
)

// Datatypes and methods of the linked list
type List struct {
	first, last *Element
}

// Interner Typ fuer die Listenelemente
type Element struct {
	next, prev *Element
	Data       any
}

func (e *Element) Next() *Element {
    if e == nil {
        return nil
    }
    return e.next
}

func (e *Element) Prev() *Element {
    if e == nil {
        return nil
    }
    return e.prev
}

func (e *Element) forward(idx int, yield func(idx int, dat any) bool) bool {
	return e == nil || yield(idx, e.Data) && e.next.forward(idx+1, yield)
}

func (e *Element) backward(idx int, yield func(idx int, dat any) bool) bool {
	return e == nil || yield(idx, e.Data) && e.prev.backward(idx+1, yield)
}

// Erstellt eine neue (leere) Liste.
func NewList() *List {
	return &List{}
}

func (l *List) Init() {
    l.first = nil
    l.last = nil
}

func (l *List) First() *Element {
    return l.first
}

func (l *List) Last() *Element {
    return l.last
}

// Push haengt ein neues Element am Ende der Liste an.
func (l *List) Push(d any) {
	ele := &Element{Data: d}
	if l.first == nil {
		l.first = ele
		l.last = ele
	} else {
		ele.prev = l.last
		l.last.next = ele
		l.last = ele
	}
}

// Pop entfernt ein Element vom Anfang der Liste und gibt es zurueck.
func (l *List) Pop() (any) {
	if l.first == nil {
		return nil
	}
    ele := l.first
	l.first = l.first.next
    if l.first == nil {
        l.last = l.first
    } else {
        l.first.prev = nil
	}
	return ele.Data
}

func (l *List) Remove(e *Element) (d any) {
    d = e.Data
    if e.prev != nil {
        e.prev.next = e.next
    } else {
        l.first = e.next
    }
    if e.next != nil {
        e.next.prev = e.prev
    } else {
        l.last = e.prev
    }
    e.next = nil
    e.prev = nil
    return d
}

func (l *List) Swap(e1, e2 *Element) {
    if e1 == e2 {
        return
    }
    e1.prev, e2.prev = e2.prev, e1.prev
    if e1.prev != nil {
        e1.prev.next = e1
    } else {
        l.first = e1
    }
    if e2.prev != nil {
        e2.prev.next = e2
    } else {
        l.first = e2
    }
    e1.next, e2.next = e2.next, e1.next
    if e1.next != nil {
        e1.next.prev = e1
    } else {
        l.last = e1
    }
    if e2.next != nil {
        e2.next.prev = e2
    } else {
        l.last = e2
    }
}

func (l *List) Len() int {
    i := 0
    for range l.Forward {
        i++
    }
    return i
}

func (l *List) Forward(yield func(idx int, dat any) bool) {
	l.first.forward(0, yield)
}

func (l *List) Backward(yield func(idx int, dat any) bool) {
	l.last.backward(0, yield)
}


// ---------------------------------------------------------------------------

type Point struct {
	X, Y, Z float64
}

func (p *Point) String() string {
	return fmt.Sprintf("(%.3f; %.3f; %.3f)", p.X, p.Y, p.Z)
}

func Cmp(p, q Point) int {
	if p.X < q.X {
		return -1
	} else if p.X == q.X {
		return 0
	} else {
		return +1
	}
}

func main() {
    numElements := 10
	lst := NewList()

    rand.Seed(123_456)
	for range numElements {
		pt := &Point{rand.Float64(), rand.Float64(), rand.Float64()}
		lst.Push(pt)
	}

    fmt.Printf("%d elements in list\n", lst.Len())

	res := Point{}
    t0 := time.Now()
    for e := lst.First(); e != nil; e = e.Next() {
        p := e.Data.(*Point)
		res.X += p.X
		res.Y += p.Y
		res.Z += p.Z
    }
    dur := time.Since(t0)
	fmt.Printf("result: %+v\n", res)
    fmt.Printf("took  : %v\n", dur)

	res = Point{}
    t0 = time.Now()
    for _, d := range lst.Forward {
        p := d.(*Point)
		res.X += p.X
		res.Y += p.Y
		res.Z += p.Z
	}
    dur = time.Since(t0)
	fmt.Printf("result: %+v\n", res)
    fmt.Printf("took  : %v\n", dur)

    for i, d := range lst.Forward {
        p := d.(*Point)
		fmt.Printf("%d: %v\n", i, p)
	}



 //    fmt.Printf("Iterate forward\n")
	// for d := range lst.Forward {
	// 	fmt.Printf("%v\n", d)
	// }

 //    fmt.Printf("Iterate backward\n")
	// for d := range lst.Backward {
	// 	fmt.Printf("%v\n", d)
	// }

    // e1 := lst.last
    // e2 := e1.prev
    // lst.Swap(e1, e2)

    // fmt.Printf("Iterate forward\n")
	// for d := range lst.Forward {
		// fmt.Printf("%v\n", d)
	// }

    // fmt.Printf("Pop all elements\n")
    // for d := lst.Pop(); d != nil; d = lst.Pop() {
    //     fmt.Printf("%v\n", d)
    // }
}
