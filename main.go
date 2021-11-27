package main

import "fmt"

type Bebek struct {
	energi       int
	hidup        bool
	bisaTerbang  bool
	suaraTerbang string
}

func (b *Bebek) Mati() {
	b.hidup = false
}

func (b *Bebek) Terbang() {
	if b.energi > 0 && b.hidup == true && b.bisaTerbang {
		fmt.Println(b.suaraTerbang)
		b.energi -= 1
		if b.energi == 0 {
			b.Mati()
		}
	}
}

func (b *Bebek) Makan() {
	if b.energi > 0 && b.hidup == true {
		b.energi += 1
	}
}

func main() {
	bebek1 := Bebek{energi: 10, hidup: true, bisaTerbang: true, suaraTerbang: "kepak"}
	bebek1.Terbang()
	bebek1.Terbang()
	bebek1.Terbang()
	fmt.Println(bebek1)

	bebek1.Makan()

	fmt.Println(bebek1)

	bebek1.Mati()

	fmt.Println(bebek1)

}
