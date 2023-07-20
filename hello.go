package main

import (
	// "bufio"
	// _ "first/data"
	// _ "first/database"
	// _ "flag"
	"container/list"
	"container/ring"
	"fmt"
	"math"
	"reflect"
	"sort"
	"time"

	// _ "os"
	"strconv"
	// _ "strings"
)

type UserFajar struct {
	name string
	umur int
}
type UserSlice []UserFajar

func (userslice UserSlice) Len() int {
	return len(userslice)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].name < value[j].name
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}
func main() {
	fmt.Println("==== Package Sort di Golang =====")
	/*Package sort adalah package yang berisikan utilitas untuk memproses pengurutan
	Agar data kita bisa di urutkan ,Kita harus mengimplementasikan kontrak interface Sort.Interface
	*/
	users := []UserFajar{
		{"irwan", 10},
		{"mamat", 20},
		{"fajar", 30},
		{"adi", 10},
		{"kuldi", 10},
		{"karmi", 10},
		{"karni", 10},
	}
	fmt.Println(users)
	sort.Sort(UserSlice(users))
	fmt.Println(users)
	fmt.Println("======Batas Pengenalan Package Sort digolang =====")

	////////
	fmt.Println("======= Package Container Ring di GOlang =========")
	/*Package Container Ring di Golang
	adalah Implementasi Struktur data Circular list atau cincin lingkaran List.
	Circular List adalah data ring,dimana diakhir element akan kembali di element awal
	Beserta cara penggunaanya
	*/

	var data *ring.Ring = ring.New(5)
	data.Value = "RIzki"
	var data2 = data.Next()
	data2.Value = "Sukma"

	//Example 1
	dataRing := ring.New(5)
	for i := 0; i < dataRing.Len(); i++ {
		dataRing.Value = " Value " + strconv.FormatInt(int64(i), 10)
		dataRing = dataRing.Next()
	}

	dataRing.Do(func(value interface{}) {
		fmt.Println(value)
	})

	fmt.Println("================================================")
	//Package container dan list
	/*adalah implementasi struktur data double Linked list di golang*/
	//Pengunaan Gouble Linked lIst
	dataLinkedList := list.New()
	dataLinkedList.PushBack(1)
	dataLinkedList.PushBack(2)
	dataLinkedList.PushBack(3)
	// dataLinkedList.PushFront("SukmaRizki")
	dataLinkedList.PushFront("Ani")
	dataLinkedList.PushBack("sukma")
	dataLinkedList.PushBack("Fajar")

	dataLinkedList.Remove(dataLinkedList.Back())
	// dataLinkedList.Front().Next().Next().Next()
	// fmt.Println(dataLinkedList.Front().Prev())
	// fmt.Println(dataLinkedList.Back().Next())
	//Iterasi dari depan kebelakang
	// for i := dataLinkedList.Front(); i != nil; i = i.Next() {
	// 	fmt.Println(i.Value)
	// }

	//Iterasi Dari belakang ke depan di container double linked List

	for item := dataLinkedList.Back(); item != nil; item = item.Prev() {
		fmt.Println(item.Value)
	}
	fmt.Println(" =================================================================")
	//Package math digolang
	/*Package yang merupakan nilai konstan dan terkandung funsi matematika */
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.3))

	//Package strconv

	var value = strconv.FormatInt(100000, 10)
	fmt.Println(value)
	// fmt.Println(strconv.ParseBool("true"))
	number, err := strconv.ParseInt("100000", 16, 32)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println("Error / Panic", err.Error())
	}

	boolean, err := strconv.ParseBool("false")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	// //package Strings
	// fmt.Println(strings.Replace("Sukma", "Sukma", "Rizki", 2))
	// fmt.Println(strings.Contains("Sukma Rizki", "Sukma"))
	// fmt.Println(strings.Trim("  Sukma Rizki ", " "))
	// fmt.Println(strings.ReplaceAll("Sukma Sukma Sukma", "Sukma", "mamat"))
	// var endString = strings.ToLower("SUKMA RIZKI")
	// toSPlit := strings.Split("Sukma Rizki", " ")
	// toTrim := strings.Trim(" sukma rizki ", " ")
	// fmt.Println(strings.ReplaceAll("Sukma Rizki", "sukma", "Rizki"))
	// fmt.Println(toTrim)

	// var toUpper = strings.ToUpper("sukmarizki")
	// fmt.Println(toSPlit)
	// fmt.Println(toUpper)
	// fmt.Println(endString)
	// flag.Parse()
	// fmt.Println(*hostApplication,*usernameHos,*password)

	// var host *string = flag.String("host", "localhost", "put Your host")
	// var user *string = flag.String("user", "root", "put Your database user")
	// var password *string = flag.String("password", "root", "put Your database password")
	// var nilai *int = flag.Int("nilai", 8, "put your nilai database")
	// // flag.Parse()
	// fmt.Println("Host name :", *host)
	// fmt.Println("Username :", *user)
	// fmt.Println("Password :", *password)
	// fmt.Println("nilai :", *nilai)
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

	fmt.Println("=== Package time yang berisikan fungsionalitas untuk memanejemen waktu ==")
	var b = time.Now()

	fmt.Println(b.Date())

	fmt.Println(b.Hour())
	fmt.Println(b.Month())
	fmt.Println(b.Minute())
	fmt.Println(b.Second())
	fmt.Println(b.Nanosecond())
	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2020-10-01")
	fmt.Println(parse)
	fmt.Println("================================")

	fmt.Println("====Package Reflect =============")
	/*Package Reflect atau singkatan dari golang
	Biasanya dalam bahasa pemograman ada fitur reflection
	dimana kita bisa melihat struktur code
	kita pada aplikasi kita sedang berjalan
	Hal ini Bisa kita lakukan di golang dengan menggunakan Fitur reflection
	Fitur ini pelajari secara bebas dan otodidak
	*/
	sample := sample{"sukma"}
	var sampleType reflect.Type = reflect.TypeOf(sample)
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
	//Kode Programm StructTag
}

type sample struct {
	Name string `required:"true" max:"1000"`
}

// data.HelloMyName("Sukma")
