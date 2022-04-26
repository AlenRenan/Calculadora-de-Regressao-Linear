package calculos

import (
	"fmt"
	"math"
)

/*
Função responsável pelo cálculo do desvio padrão de x e y
*/
func Pearson(n int, x []float64, y []float64, miX float64, miY float64, cov float64) {

	var devioX, desvioY, somQuadradoX, somQuadradoY, r float64

	/*
		Realiza a somatória quadrada de x e y
	*/
	for i := 0; i < n; i++ {

		somQuadradoX = somQuadradoX + (x[i] * x[i])
		somQuadradoY = somQuadradoY + (y[i] * y[i])
	}

	/*
		Realiza os cálculos de desvio de x e y
	*/
	devioX = math.Sqrt((somQuadradoX / float64(n)) - (miX * miX))
	desvioY = math.Sqrt((somQuadradoY / float64(n)) - (miY * miY))

	/*
		Realiza o cálculo do coeficiente  de  correlação  de Pearson
	*/
	r = cov / (devioX * desvioY)

	/*
		Imprime os valores
	*/
	fmt.Println("\n\nRHO de x: ", devioX, "\n\nRHO de y:", desvioY, "\n\nCoeficiente  de  correlação  de Pearson: ", r)

}
