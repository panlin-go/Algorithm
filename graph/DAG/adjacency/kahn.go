package adjacency

import (
	"fmt"
)

func MapCopy(m map[string]int) map[string]int {
	m1 := make(map[string]int)

	for k, v := range m {
		m1[k] = v
	}

	return m1
}

func (ad *Adjacency) Kahn() bool {

	zero := make(map[string]int, 0)

	input := MapCopy(ad.InputDU)

	fmt.Println("kahn sort: ")
	for len(input) > 0 {
		for k, v := range input {
			if v == 0 {
				zero[k] = 0
				delete(input, k)
			}
		}

		for k, _ := range zero {
			fmt.Printf(" [%s] ", k)
			delete(zero, k)

			var mem Member
			for i := 0; i < len(ad.Member); i++ {
				if ad.Member[i].Name == k {
					mem = ad.Member[i]
					break
				}
			}

			var p *List = mem.Next
			for i := 0; i < mem.Count; i++ {
				input[p.Name]--
				p = p.Next
			}
		}

	}

	fmt.Println()

	return true
}
