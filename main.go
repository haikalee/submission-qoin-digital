package main

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	/* menampilkan hasil akhir */
	const pemain = 3
	const dadu = 4

	ukuran := map[string]map[int]int{}
	result := map[string]int{}

	/* proses membuat dadu acak */
	for i := 0; i < pemain; i++ {
		ukuran[fmt.Sprintf("Pemain#%v", i+1)] = map[int]int{
			1: dadu,
		}
	}

	giliran := 1
	ketemuangka1 := 0
	for {
		totalPemain := 0
		log.Println()
		log.Println("======================")
		log.Println("Giliran", giliran, "Lempar Dadu:")

		/* membuat dadu acak sesuai ukuran */
		battlefield := [][]string{}
		for i := 0; i < pemain; i++ {
			battlefield = append(battlefield, []string{})
			for j := 0; j < ukuran[fmt.Sprintf("Pemain#%v", i+1)][giliran]; j++ {
				battlefield[i] = append(battlefield[i], strconv.Itoa(rand.Intn(6)+1))
			}

			/* menampilkan dadu yang diacak */
			log.Printf("Pemain#%v: %v", i+1, battlefield[i])
		}

		log.Printf("===========Setelah Evaluasi===========")
		tmpData := map[int][]string{}
		for i, perOrang := range battlefield {
			var tmpDadu string
			tmpDadu = strings.Join(perOrang, "")

			/* proses mencari angka 6 */
			find6regex := regexp.MustCompile(`6`)
			found6Number := find6regex.FindAllString(tmpDadu, -1)
			if len(found6Number) > 0 {
				/* proses menambah point */
				result[fmt.Sprintf("Pemain#%v", i+1)] += len(found6Number)

				/* menghapus angka 6 */
				remove := find6regex.ReplaceAllString(tmpDadu, "")
				tmpDadu = remove
				battlefield[i] = strings.Split(remove, "")
			}

			/* proses mencari angka 1 */
			isEdit := false
			find1regex := regexp.MustCompile(`1`)
			found1number := find1regex.FindAllString(tmpDadu, -1)
			if len(found1number) > 0 {

				/* memindah angka 1 ke selanjutnya */
				var next int
				if i+1 < len(battlefield) {
					next = i + 1
				} else if i+1 == len(battlefield) {
					next = 0
				}

				if next != 0 {
					tmpData[next] = append(tmpData[next], battlefield[next]...)
				}

				for j := 0; j < len(found1number); j++ {
					tmpData[next] = append(tmpData[next], "1")
				}

				/* menghapus angka 1 */
				remove := find1regex.ReplaceAllString(tmpDadu, "")

				for j := 0; j < ketemuangka1; j++ {
					remove += "1"
				}
				ketemuangka1 = len(found1number)

				tmpData[i] = strings.Split(remove, "")
				isEdit = true
			} else {
				if ketemuangka1 > 0 {
					remove := find1regex.ReplaceAllString(tmpDadu, "")
					for j := 0; j < ketemuangka1; j++ {
						remove += "1"
					}
					ketemuangka1 = len(found1number)

					tmpData[i] = strings.Split(remove, "")
					isEdit = true
				}
			}

			if !isEdit {
				tmpData[i] = strings.Split(tmpDadu, "")
			}

			/* setup ukuran digiliran berikutnya */
			ukuran[fmt.Sprintf("Pemain#%v", i+1)] = map[int]int{
				giliran + 1: len(tmpData[i]),
			}

			if len(battlefield[i]) > 0 {
				totalPemain++
			}
		}

		log.Println("Pemain#1:", tmpData[0])
		log.Println("Pemain#2:", tmpData[1])
		log.Println("Pemain#3:", tmpData[2])

		giliran++
		/* proses menghentikan permainan */
		if totalPemain <= 1 {
			break
		}
	}

	/* menampilkan hasil akhir */
	type KeyValue struct {
		Key   string
		Value int
	}
	var sortedSlice []KeyValue

	for k, v := range result {
		sortedSlice = append(sortedSlice, KeyValue{k, v})
	}

	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i].Value > sortedSlice[j].Value
	})

	log.Println(result)
	log.Println(sortedSlice)

	for _, value := range sortedSlice {
		log.Println(value.Key, value.Value)
	}

	log.Println("Pemenang:", sortedSlice[0].Key)

}
