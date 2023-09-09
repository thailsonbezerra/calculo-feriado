package main

import (
	"fmt"
	"time"
)

type DataFeriado struct {
	Nome string
	Dia string 
	Mes string
}

func main() {
	fmt.Printf("Feriadou")

	anoAtual := time.Now().Year()

	feriadosFixos := []DataFeriado{
		{"Confraternização Universal", "01", "01"},
		{"Tiradentes", "21", "04"},
		{"Dia do Trabalhador", "01","05"},
		{"Independência do Brasil", "07", "09"},
		{"Nossa Senhora Aparecida", "12", "10"},
		{"Finados", "02", "11"},
		{"Proclamação da República", "15", "11"},
		{"Natal", "25", "12"},
	}

	for _, feriados := range feriadosFixos{
		fmt.Printf("%s/%s - %s\n",feriados.Dia, feriados.Mes, feriados.Nome)
	}

	dataPascoa := FormulaDeGauss(anoAtual)
	
	fmt.Printf("%s/%s - %s\n",dataPascoa.Dia, dataPascoa.Mes, dataPascoa.Nome)
}

//FormulaDeGauss é utilizada para calcular o dia da Páscoa. A fórmula vale para anos entre 1901 e 2099. A fórmula pode ser estendida para outros anos, alterando X e Y (criada por Gauss até 1999 e estendida pelo autor até 2299)
func FormulaDeGauss(ano int) DataFeriado {
	X := 24
	Y := 5
	a := ano%19
	b := ano%4
	c := ano%7
	d := (19*a+X)%30
	e := (2*b+4*c+6*d+Y)%7

	var dia int
	var mes int

	if d+e > 9 {
		dia = d+e-9
		mes = 4
	} else {
		dia = d+e+22
		mes = 3
	}

	switch {
	case mes == 4 && dia == 26:
			dia = 19
	case mes == 4 && dia == 25 && d == 28 && a > 10:
			dia = 18
	}

	return DataFeriado{"Páscoa", fmt.Sprintf("%02d", dia), fmt.Sprintf("%02d", mes)} 

}