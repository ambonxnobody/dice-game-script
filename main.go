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
	var toAddToTheNextPlayers []int
	for _, pemain := range pemains {
		var toAddToTheNextPlayer int
		for i := 0; i < len(pemain.Dadu); i++ {
			switch pemain.Dadu[i] {
			case 6:
				pemain.Poin++
				pemain.Dadu = append(pemain.Dadu[:i], pemain.Dadu[i+1:]...)
				i--
			case 1:
				toAddToTheNextPlayer++
				pemain.Dadu = append(pemain.Dadu[:i], pemain.Dadu[i+1:]...)
				i--
			}
		}

		toAddToTheNextPlayers = append(toAddToTheNextPlayers, toAddToTheNextPlayer)
	}

	for i := 0; i < len(toAddToTheNextPlayers); i++ {
		if toAddToTheNextPlayers[i] > 0 {
			for j := 0; j < toAddToTheNextPlayers[i]; j++ {
				if i+1 < len(pemains) && len(pemains[i+1].Dadu) > 0 {
					pemains[i+1].Dadu = append(pemains[i+1].Dadu, 1)
				} else if i+1 < len(pemains) && len(pemains[i+1].Dadu) == 0 {
					found := false
					for k := i + 2; k < len(pemains); k++ {
						if len(pemains[k].Dadu) > 0 {
							pemains[k].Dadu = append(pemains[k].Dadu, 1)
							found = true
							break
						}
					}

					if !found {
						for k := 0; k < i; k++ {
							if len(pemains[k].Dadu) > 0 {
								pemains[k].Dadu = append(pemains[k].Dadu, 1)
								break
							}
						}
					}
				} else {
					pemains[0].Dadu = append(pemains[0].Dadu, 1)
				}
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
	var pemenang []*Pemain
	for _, player := range pemains {
		if player.Poin >= poinTerbanyak {
			poinTerbanyak = player.Poin
		}
	}

	for _, player := range pemains {
		if player.Poin == poinTerbanyak {
			pemenang = append(pemenang, player)
		}
	}

	if len(pemenang) > 1 {
		fmt.Println("Game dimenangkan oleh pemain-pemain berikut karena memiliki poin yang sama:")
		for _, player := range pemenang {
			fmt.Printf("Pemain #%d\n", player.ID+1)
		}
		return
	} else if len(pemenang) == 1 {
		fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.\n", pemenang[0].ID+1)
		return
	}
}
