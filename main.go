package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	hedefSayi := rand.Intn(100) + 1
	var tahmin int
	var denemeSayisi int

	fmt.Println("1 ile 100 arasında bir sayı tahmin edin.")

	for {
		denemeSayisi++
		fmt.Print("Tahmininiz: ")
		fmt.Scanln(&tahmin)

		if tahmin < hedefSayi {
			fmt.Println("Daha büyük bir sayı tahmin edin.")
		} else if tahmin > hedefSayi {
			fmt.Println("Daha küçük bir sayı tahmin edin.")
		} else {
			fmt.Printf("Tebrikler! Doğru tahmin ettiniz. %d denemede doğru tahmin ettiniz.\n", denemeSayisi)
			break
		}
	}
}
