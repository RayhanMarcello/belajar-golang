package main

import (
	"fmt"
	"strconv"
)

func main() {
	// membuat variabel cara 1
	var nama string = "rayhan"
	var umur int = 20

	// membuat variabel cara 2
	salam := "assalamualaikum"
	count := 1
	punyaCewe := true

	// variabel const(kayak js
	const pacar = "luna maya"
	fmt.Println("materi variabel :", salam, nama, umur, count, punyaCewe, len(pacar))

	// conversion
	var text string = "100"
	number, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("error :", err.Error())
	}else{
		fmt.Println(number)
	}

	// type declarations
	type NoKTP string

	var ktpEko NoKTP = "123123"
	var contoh string = "555555"
	var contohKtp = NoKTP(contoh)

	fmt.Println(ktpEko)
	fmt.Println(NoKTP("222222"))
	fmt.Println(contohKtp)

	// operasi matematika
	a := 10
	var b = 10
	var c = a+b
	fmt.Println(c)

	// augmented assigments
	b+=1 // b = b + 1
	fmt.Println(b)

	// unary operatir
	a++
	fmt.Println(a) // a = a + 1

	// perbandingan
	var kalimat1 = "kontolodon"
	var kalimat2 = "ambatukam"
	result  := len(kalimat1) > len(kalimat2)
	fmt.Println(result)

	//operasi boolean (&&, ||, !)
	number1 := false
	number2 := true
	resultNum := number1 && number2 
	fmt.Println(resultNum)

	var names [3]string
	names[0] = "rayhan"
	names[1] = "rayhan"
	names[2] = "rayhan"
	fmt.Println(names)

	arr := []int{1,2,3}
	fmt.Println(arr)
}	
