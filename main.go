package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func removeindex(i []int, index int) []int {
	return append(i[:index], i[index+1:]...)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "asd")
	})

	http.HandleFunc("/sort", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, fastSort(convertString(strings.Split(r.URL.Query().Get("arr"), ","))))
	})

	http.ListenAndServe(":80", nil)
}

func convertString(s []string) []int {
	var ints []int
	for _, k := range s {
		num, err := strconv.Atoi(k)
		if err == nil {
			ints = append(ints, num)
		}
	}
	return ints
}

func fastSort(array []int) []int {
	if len(array) < 2 {
		return array
	} else {
		pivot := array[len(array)/2]
		var less []int
		var greater []int
		for k := 0; k < len(array); k++ {
			if array[k] < pivot {
				less = append(less, array[k])
			} else if array[k] > pivot {
				greater = append(greater, array[k])
			}
		}
		var a []int
		a = append(a, fastSort(less)...)
		a = append(a, pivot)
		a = append(a, fastSort(greater)...)
		return a
	}
}

func choseSort(i []int) []int {
	var si []int
	var id int
	var max int
	for d := 0; d < len(i)-1; {
		max = i[d]
		id = d
		for j := 1; j < len(i); j++ {
			if max < i[j] {
				max = i[j]
				id = j
			}
		}
		si = append(si, max)
		i = removeindex(i, id)
	}
	si = append(si, i[0])
	i = removeindex(i, 0)
	return si
}
