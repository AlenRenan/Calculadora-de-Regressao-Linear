package calculos

import (
	"fmt"
)

func CovaricanciaXeY(n int, x []float64, y []float64) {

	//var somX, mediaProd int //Armazena o valor da somatoria de x
	//var somY int            //Armazena o valor da somatoria de y
	var miX, miY, mXeY, cov, somY, somX, mediaProd float64

	/*
		Realiza a somatória de x e y
	*/
	for i := 0; i < n; i++ {

		somX = somX + x[i]
		somY = somY + y[i]
	}

	/*
		Calcula os valores de M de x e y (M = somatoria(x/y)/n)
	*/
	miX = (float64(somX) / float64(n))
	miY = (float64(somY) / float64(n))

	for k := 0; k < n; k++ {

		mediaProd = (x[k] * y[k]) + mediaProd

	}

	mXeY = float64(mediaProd) / float64(n)

	/*
		Calcula a covariância de xy
	*/
	cov = mXeY - miX*(miY)

	fmt.Print("Somatória de x = ", somX, "\nMi de x = ", miX, "\n\nSomatória de y = ", somY, "\nMi de y = ", miY, "\n\nMi de x e y = ", mXeY, "\n\nCovariância de x e y =  ", cov) //Imprime os valores gerados

	if cov == 0 {
		fmt.Println("\n\nNÃO existe relação linear entre x e y")
	} else if cov < 0 {
		fmt.Println("\n\nExiste uma relação linerar INVERSA entre x e y")
	} else if cov > 0 {
		fmt.Println("\n\nExiste uma relação linear DIRETA  entre x e y")
	}
	Pearson(n, x, y, miX, miY, cov)
}
