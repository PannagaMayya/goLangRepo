package main

import "fmt"

type Outergroup []string

/*Reciever Functions*/
func (g Outergroup) addA() {
	for i, o := range g {
		g[i] = o + "A"
	}
}

type rrr struct {
	a string
	b string
}

func (p *rrr) update(a string) {
	(*p).a = a
}
func main() {
	//SLICE / similar to ARRAY
	fmt.Println("SLICE")
	cards := []string{"Hi", "How", "are"}
	cards[2] = "ahah"
	i, card := 0, "Diff scope"
	for i, card := range cards { //i and card are block scope
		fmt.Println(i, card)
	}
	fmt.Println("New cards slice", append(cards, "you?"))
	fmt.Println("Unmodified cards slice", cards)
	fmt.Println("Scope - i=", i, "card=", card)
	//TYPE
	fmt.Println("TYPE")
	type Innergroup []string
	/*Reciever Functions*/
	print := func(g Innergroup) {
		fmt.Println(g)
	}
	toPrintIn := Innergroup{"How", "are", "you", "from Inner Type"}
	print(toPrintIn)
	toPrintOut := Outergroup{"How", "are", "you", "from Outer Type"}
	toPrintOut.addA()
	fmt.Println(toPrintOut)
	//STRUCT / OBJECT
	fmt.Println("STRUCT")
	type rr struct {
		a string
		b string
	}
	r := rr{
		b: "b",
		a: "a", //comma is must for multi line
	}
	var e struct {
		a string
		b string
		c rr
	}
	e.a = "q"
	e.b = "w"
	e.a = r.a
	e.c = r
	fmt.Println(r.a)
	fmt.Printf("%+v\n", e)
	/*Modify Struct*/
	qqq := rrr{"a", "b"}
	qqq.update("ModifiedwithNoPointer")
	fmt.Println(qqq)
	(&qqq).update("ModifiedwithPointer")
	fmt.Println(qqq)
}
