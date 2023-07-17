package data

import "fmt"

/* Tidak Boleh meletakkkan function yang sama didalam package yang sama*/

var latihanku = "Sukma"
var TrainingCoding = 34567890

func HelloMyName(name string) {
	fmt.Println("Hello: ", name)
}

func accesModifier(first string) {
	fmt.Println(first)
}
