package main

import (
	"fmt"
	"strconv"
)

// function && parameter function
func sayHello(firstName string, lastName string){
	fmt.Println("say hello run brow", firstName, lastName)
}

// return value func
func asalDaerah(daerah string) string{
	return "halo saya dari" + daerah
}

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

	var names [4]string
	names[0] = "rayhan"
	names[1] = "marcello"
	names[2] = "ananda"
	names[3] = "purnomo"
	fmt.Println(names)

	arr := [...]int{1,2,3}
	fmt.Println(arr)

	//function arr
	panjangNames := len(names)
	fmt.Println(panjangNames)

	hasilArr := arr[0]
	fmt.Println(hasilArr)

	arr[0] = 100
	fmt.Println(arr)

	// slice (kayak aray tapi lebih dinamis)
	slice := names[1:3]
	fmt.Println(slice[0])
	fmt.Println(slice[1])

	// function slice

	panjangSlice := len(slice)
	fmt.Println(panjangSlice)

	kapasitasSlice := cap(slice)
	fmt.Println(kapasitasSlice)

	addSlice := append(slice, "kontolodon")
	fmt.Println(addSlice)

	//bedanya array dan slice
	// iniArray := [...]int{1,2,3,4,5}
	// iniSlice := []int{1,2,3,4,5}

	// tipe data map
	person := map[string]string{
		"name" : "rayhan",
		"address" : "borneo",
	}
	fmt.Println(person["name"])

	// conditional statement
	
	if names[0] == "rayhan" {
		fmt.Println(`halo`, names[0])
	}else{
		fmt.Println("kamu bukan rayhan")
	}

	if length := len(names[0]) ; length>7 {
		fmt.Println("nama terlalu panjang")
	}else{
		fmt.Println("nama terlalu pendek")
	}

	// for loop
	counter := []int{1,2,34,12}
	counting := len(counter)
	for i := 0;i < counting ;i++{
		fmt.Println(counter[i])
	}

	// break & continue
	for i:=1; i <=10; i++{
		if(i == 5){
			break
		}
		fmt.Println(i)
	}

	for i:=1; i <10; i++{
		if i % 2 == 0{
			fmt.Println("perulangan mod 2 habis dibagi ",i)
			continue
		}
	}

	// function
	firstName := "rayhan"
	lastName := "marcello"
	sayHello(firstName,lastName)

	// return value func
	results := asalDaerah(" ambon")
	fmt.Println(results)
}	
