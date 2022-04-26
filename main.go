package main

import (
	"calculos/calculos"
	"fmt"
	"time"
)

func main() {

	fmt.Print("\nInsira o valor de n: ") //msg
	var n int                            //Iniciavaliza a variável n, y e y
	fmt.Scanln(&n)                       //Escaneia a entrada com os valores de x

	var x = make([]float64, n) //Cria um slice quer armazenará os valores de x
	var y = make([]float64, n) //Cria um slice quer armazenará os valores de y

	/*
		Estrutura de repetição para que os valores de x e y sejam inseridos
		pelo usuário
	*/
	for i := 0; i < n; i++ {
		fmt.Print("x", i+1, ": ")
		fmt.Scanln(&x[i]) //Escaneia a entrada com os valores de x
	}
	for i := 0; i < n; i++ {
		fmt.Print("y", i+1, ": ")
		fmt.Scanln(&y[i]) //Escaneia a entrada com os valores de y
	}

	fmt.Print("\n\nTabela\nn: ", n, "\nx:", x, "\ny:", y, "\n\n") //Imprime os valores das variáveis

	/*
		Chama a função que realiza os cálculos de Covariância
	*/
	calculos.CovaricanciaXeY(n, x, y)
	//auxiliar.desvioPadraoXeY(n, x, y)
	fmt.Print("\n\nO programa fecha em 1 minuto !!")
	time.Sleep(1 * time.Minute)
}
