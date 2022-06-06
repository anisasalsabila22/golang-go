package main

import "fmt"

func main() {

	name := "salsabila"

	switch name {
	case "anisa":
		fmt.Println("Hello anisa")
		fmt.Println("Hello anisa")
	case "caca":
		fmt.Println("Hello caca")
		fmt.Println("Hello caca")
	default:
		fmt.Println("Kenalan Donk")
		fmt.Println("Kenalan Donk")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
