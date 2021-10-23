package main

import "fmt"

type Foo struct {
	Bar int
}

func (f *Foo) Increment() {
	f.Bar += 1
}

func main() {
	var f Foo

	// Calling `f.Increment` is automatically changed to `(&f).Increment` by the compiler.
	f = Foo{}
	fmt.Printf("f.Bar is %d\n", f.Bar)
	f.Increment()
	fmt.Printf("f.Bar is %d\n", f.Bar)

	// As you can see, calling `(&f).Increment` directly does the same thing.
	f = Foo{}
	fmt.Printf("f.Bar is %d\n", f.Bar)
	(&f).Increment()
	fmt.Printf("f.Bar is %d\n", f.Bar)
}
