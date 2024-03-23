package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Pemain struct untuk menyimpan data pemain
type Pemain struct {
	ID   int
	Dadu []int
	Poin int
}

// LemparDadu lempar dadu untuk semua pemain
func LemparDadu(pemains []*Pemain) {
	for _, pemain := range pemains {
		for i := range pemain.Dadu {
			pemain.Dadu[i] = rand.Intn(6) + 1
		}
	}
}

// EvaluasiDadu evaluasi dadu untuk semua pemain
func EvaluasiDadu(pemains []*Pemain) {
	for _, pemain := range pemains {
		if len(pemain.Dadu) == 0 {
			continue
		}

		for i := 0; i < len(pemain.Dadu); i++ {
			switch pemain.Dadu[i] {
			case 6:
				pemain.Poin++
				pemain.Dadu = append(pemain.Dadu[:i], pemain.Dadu[i+1:]...)
				i--
			case 1:
				pemain.Dadu = append(pemain.Dadu[:i], pemain.Dadu[i+1:]...)
				i--
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var jumlahPemain, jumlahDadu int
	fmt.Print("Pemain = ")
	scanPemain, err := fmt.Scanln(&jumlahPemain)
	if err != nil || scanPemain != 1 || jumlahPemain < 2 {
		fmt.Println("Input harus berupa angka dan minimal 2 pemain.")
		return
	}
	fmt.Print("Dadu = ")
	scanDadu, err := fmt.Scanln(&jumlahDadu)
	if err != nil || scanDadu != 1 || jumlahDadu < 1 {
		fmt.Println("Input harus berupa angka dan minimal 1 dadu.")
		return
	}

	// Inisialisasi pemain
	pemains := make([]*Pemain, jumlahPemain)
	for i := range pemains {
		pemains[i] = &Pemain{
			ID:   i,
			Dadu: make([]int, jumlahDadu),
		}
	}

	// Main game
	round := 1
	for {
		fmt.Printf("==================\nGiliran %d lempar dadu:\n", round)

		LemparDadu(pemains)
		for _, pemain := range pemains {
			dicesStr := ""
			if len(pemain.Dadu) == 0 {
				dicesStr = "_ (Berhenti bermain karena tidak memiliki dadu)"
			} else {
				for i, dice := range pemain.Dadu {
					if i > 0 {
						dicesStr += ", "
					}
					dicesStr += fmt.Sprintf("%d", dice)
				}
			}
			fmt.Printf("Pemain #%d (%d): %v\n", pemain.ID+1, pemain.Poin, dicesStr)
		}

		EvaluasiDadu(pemains)
		fmt.Println("Setelah evaluasi:")
		for _, pemain := range pemains {
			dicesStr := ""
			if len(pemain.Dadu) == 0 {
				dicesStr = "_ (Berhenti bermain karena tidak memiliki dadu)"
			} else {
				for i, dice := range pemain.Dadu {
					if i > 0 {
						dicesStr += ", "
					}
					dicesStr += fmt.Sprintf("%d", dice)
				}
			}
			fmt.Printf("Pemain #%d (%d): %v\n", pemain.ID+1, pemain.Poin, dicesStr)
		}

		// Cek apakah hanya ada satu pemain yang memiliki dadu, jika ya maka game berakhir
		activePlayers := 0
		var pemainTersisa *Pemain
		for _, pemain := range pemains {
			if len(pemain.Dadu) > 0 {
				activePlayers++
				pemainTersisa = pemain
			}
		}
		if activePlayers == 1 {
			fmt.Printf("==================\nGame berakhir karena hanya pemain #%d yang memiliki dadu.\n", pemainTersisa.ID+1)
			break
		}

		round++
	}

	// Cari pemain dengan poin terbanyak
	poinTerbanyak := 0
	var pemenang *Pemain
	for _, player := range pemains {
		if player.Poin > poinTerbanyak {
			poinTerbanyak = player.Poin
			pemenang = player
		}
	}
	fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.\n", pemenang.ID+1)
}
