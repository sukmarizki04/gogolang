package main

import (
	// "bufio"
	"fmt"
	// "io"
	// "os"
)

type Error struct {
	city, Province, negara string
}

//Pass By reference dengan Pointer

//Pointer function digolang
//Saat kita membuat parameter difunction ,secara default adalah Pass by value ,artinya data akan
//data akan dicopy dan dikirim

func changeFunction(error Error) {
	error.city = "Pekan Baru"
}

func changeCountryToAmerican(country Error) {
	country.negara = "American"
}

// package main

// func main() {
// 	file, err := os.Open("go")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()

// 	// Get the file size
// 	stat, err := file.Stat()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// Read the file into a byte slice
// 	bs := make([]byte, stat.Size())
// 	_, err = bufio.NewReader(file).Read(bs)
// 	if err != nil && err != io.EOF {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(bs)
// }

func main() {
	newFunction := Error{"Medan ", "Sumatra Utara", ""}
	changeFunction(newFunction)
	fmt.Println(newFunction)

	var countryError = Error{
		city:     "Medan",
		Province: "North Sumatra",
		negara:   "American Paman Sam",
	}

	changeCountryToAmerican(countryError)
	fmt.Println(countryError)
	address := Error{"Lhokseumawe", "aceh", ""}
	address2 := &address
	var address3 *Error = &address
	address2.city = "Rantau Prapat"
	var address4 *Error = new(Error)
	address4.city = "Kota Baru"
	fmt.Println(address4)
	fmt.Println(address2)
	fmt.Println(address)
	fmt.Println(address3)
}
