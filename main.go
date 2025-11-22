package main

import (
	"fmt"
	"strconv"
)

// function && parameter function
func sayHello(firstName, lastName string){
	fmt.Println("say hello run brow", firstName, lastName)
}

// return value func
func asalDaerah(daerah string) string{
	return "halo saya dari" + daerah
}

// multipel return 
func multiValueFunc(gg1 string, gg2 string) (string,string){
	return gg1, gg2
}

// vaiadic func
func sum(numberize ...int)int{
	total := 0
	for _, num := range numberize {
		total += num
	}
	return total
}

// func value 
func goodBye(namaLU string)string{
	return "selamat tinggal"+ namaLU
}

// anonymus func
type BlackList func(string) bool

func registerUser(namaaa string, blackList BlackList){
	if blackList(namaaa){
		fmt.Println("u are blocked", namaaa)
	}else{
		fmt.Println("wellcome", namaaa)
	}
}

// recrusif func
// pake for loop
func factorialLoop(values int)int{
	result := 1
	for i := values; i > 0; i--{
		result = result*i
	}
	return result
}

// defer
func logging(){
	fmt.Println("selesai memanggil func")
}

func runApp(){
	defer logging()
	fmt.Println("run app...")
}

// panic & recover(mengambil pesan dari panic)
func endApp(){
	fmt.Println("menghentikan app..")
	message := recover()
	fmt.Println(message)
}
func jalankanApp(error bool){
	defer endApp()
	if error{
		panic("error")
	}
}

	// struct (template data/ prototype data)
	type Customer struct{     //untuk struct biasanya pake Pascalcase
		Name, Address string
		Age int
	}

// struct methode
func(customer Customer) ngomongHello(name string){
	fmt.Println("hellow " + customer.Name + "my name is" + name)
}

// Nil
func NewMap(nameLUBRO string) map[string] string{
	if nameLUBRO == ""{
		return nil
	}else{
		return map[string]string{
			"name" : nameLUBRO,
		}
	}
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

	// return multipel value
	gg1 := "asade"
	gg2 := "kontol"
	hasil1, hasil2 := multiValueFunc(gg1,gg2)
	fmt.Println(hasil1,hasil2)

	// return varadic func
	total1 := sum(1,2,3,4,5)
	fmt.Println(total1)

	// variabel as a value
	goodbye := goodBye
	fmt.Println(goodbye("adi"))

	// annonymus func
	blacklist := func(namaaa string) bool{
		return namaaa == "anjing"
	}

	registerUser("anjing",blacklist)
	//bisajuga =>
	registerUser("rayhan", func(namaaa string) bool{
		return namaaa == "anjing"
	})

	// recrusive menggunakan for
	fmt.Println(factorialLoop(3))

	// closure
	nilai:= 0
	menghitung := func(){
		fmt.Println("menghitung", nilai)
		nilai++
		fmt.Println("menghitung", nilai)
	}

	menghitung()
	// defer

	runApp()

	// panic
	jalankanApp(true)


	var identitas Customer
	identitas.Name = "rayhan"
	identitas.Address = "borneo paradiso"
	identitas.Age = 20

	fmt.Println(identitas)

	//struct literals
	joko := Customer{
		Name : "joko",
		Address: "sama kek rayhan",
		Age: 50,
	}
	fmt.Println(joko)

	// struct methode
	joko.ngomongHello("agus")
	
	// nil
	data := NewMap("jordan")
	if data == nil {
		fmt.Println("data kosong")
	}else{
		fmt.Println(data["name"])
	}

	// pointer (reference ke data yang sama dengan menggunakan (&))
	type AddressUser struct{
		City,Province,Country string
	}

	address1 := AddressUser{"balikpapan", "kalimantan", "indonesia"}
	address2 := &address1 //pointer
	address2.City = "kontolodon"

	fmt.Println(address1)
	fmt.Println(address2)

	// asterisk operator
	address2 = &AddressUser{"jakarta", "jawa", "indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}	
