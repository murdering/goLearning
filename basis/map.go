package basis.map

import "fmt"

func main() {
	var mm = map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println(mm)

	var b = mm[2]
	fmt.Println("b: ", b)

	mm[2] = b + "222b"
	fmt.Println("mm[2]: ", mm[2])

	mm[4] = ""
	fmt.Println("mm[4]: ", mm[4])

	var e, ok = mm[5]
	fmt.Println("e: ", e, "ok: ", ok)

	delete(mm, 4)
	fmt.Println("mm[4]: ", mm[4])
}
