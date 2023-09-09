package main

import "fmt"

type DataFeriado struct {
	Nome string
	Dia string 
	Mes string
}

func main() {
	fmt.Printf("Feriadou")

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
}