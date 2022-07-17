package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// bem formatado
	horaAtual := "20220717181537"
	// mal formatado, sao segundos 
	tempoRestante := "269"

	segundosRestantes,_ := strconv.Atoi(tempoRestante)

	atual,_ := time.Parse("20060102150405", horaAtual)
	chegada := atual.Add(time.Duration(segundosRestantes) * time.Second)

	// convert chegada to string
	//chegadaString := strings.Split(chegada.String(), " ")[1]

	fmt.Println(atual)
	fmt.Println(chegada.Format("150405"))
	fmt.Println(chegada.Sub(atual).String())


	server := chegada.Format("150405")
	fmt.Println(time.Parse("150405", server))


}