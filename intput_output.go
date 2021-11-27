package main 

import (
	"fmt"
	"bufio"
	"os"

)

func main()  {
	edad := 15
	fmt.Printf("Mi edad es %d\n", edad)

	nombre := "Jeremias"
	fmt.Printf("Mi edad es %v\n", nombre)

	bander := true
	fmt.Printf("True or false? %t\n", bander)

	var edad2 int 
	fmt.Println("Coloca tu edad: ")
	fmt.Scanf("%d\n", &edad2)
	fmt.Printf("Mi edad es %d\n",edad2)

	var nombre2 string
	fmt.Println("Coloca tu nombre: ")
	fmt.Scanf("%s\n", &nombre2)
	fmt.Printf("Hola %s\n",nombre2)

	//Bufio

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre:")
	text, err := reader.ReadString('\n')

	if (err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println("Hola "+ text)
	}
	}