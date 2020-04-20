package basis.struct

import "fmt"

func main() {
	gogo := &Saiyan{Name: "gogo", Power: 100}
	PowerPlus(gogo)
	gogo.Super("new gogogo")
	fmt.Println(gogo.Name, gogo.Power)

	// 伪继承
	prosePoem := &ProsePoem{
		Poem: Poem{
			Title:  "Jack",
			Author: "slow",
			intro:  "simple",
		},
		Author: "test",
	}
	fmt.Println(prosePoem)
}

type Saiyan struct {
	Name  string
	Power int
}

func PowerPlus(s *Saiyan) {
	s.Power += 200
	// s = &Saiyan{Name:"new gogo", Power:500}
}

func (s *Saiyan) Super(newName string) {
	s.Name = newName
	s.Power += 600
}

// 伪继承
func (e *Poem) ShowTitle() {
	fmt.Printf(e.Title)
}

type Poem struct {
	Title  string
	Author string
	intro  string
}

type ProsePoem struct {
	Poem
	Author string
}
