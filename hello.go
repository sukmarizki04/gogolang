package main

import (
	// "bufio"
	_ "first/data"
	_ "first/database"
	"flag"
	"fmt"
	_ "os"
	"strings"
)

func main() {

	//package Strings

	var endString = strings.ToLower("SUKMA RIZKI")
	toSPlit := strings.Split("Sukma Rizki", " ")
	toTrim := strings.Trim(" sukma rizki ", " ")
	fmt.Println(strings.ReplaceAll("Sukma Rizki", "sukma", "Rizki"))
	fmt.Println(toTrim)
	var toUpper = strings.ToUpper("sukmarizki")
	fmt.Println(toSPlit)
	fmt.Println(toUpper)
	fmt.Println(endString)

	// //Program Package Flag
	// hostApplication := flag.String("host", "localhost", "Masukkan database Hos")
	// usernameHos := flag.String("username", "root", "put your database username")
	// password := flag.String("password", "root","Put your database password")

	// flag.Parse()
	// fmt.Println(*hostApplication,*usernameHos,*password)

	var host *string = flag.String("host", "localhost", "put Your host")
	var user *string = flag.String("user", "root", "put Your database user")
	var password *string = flag.String("password", "root", "put Your database password")
	var nilai *int = flag.Int("nilai", 8, "put your nilai database")
	// flag.Parse()
	fmt.Println("Host name :", *host)
	fmt.Println("Username :", *user)
	fmt.Println("Password :", *password)
	fmt.Println("nilai :", *nilai)
	//Menggunkan Package OS
	// args := os.Args
	// fmt.Println("Pembelajaran Argumen digolang Os Arguments")
	// fmt.Println(args)
	// fmt.Println(args[1])
	// fmt.Println(args[2])

	// name, err := os.Hostname()
	// if err == nil {
	// 	fmt.Println("My Name :", name)
	// } else {
	// 	fmt.Println("Could not", err.Error())
	// }

	// var t = data.TrainingCoding
	// fmt.Println(t)
	// // hasil := database.GetDatabase()
	// // fmt.Println(hasil)}
}

// data.HelloMyName("Sukma")
