package main

import (
	"fmt"
	"log"
	"time"
)

type DataFeriado struct {
	Nome string
	Dia string 
	Mes string
}

func main() {
	anoAtual := time.Now().Year()

	feriadosFixos := []DataFeriado{
		{"Confraternização Universal", "01", "01"},
		{"Tiradentes", "21", "04"},
		{"Dia do Trabalho", "01","05"},
		{"Independência do Brasil", "07", "09"},
		{"Nossa Sr.a Aparecida - Padroeira do Brasil", "12", "10"},
		{"Finados", "02", "11"},
		{"Proclamação da República", "15", "11"},
		{"Natal", "25", "12"},
	}

	feriadosMoveis := definirFeriadosMoveis(anoAtual)

	feriados := append(feriadosFixos, feriadosMoveis...)

	for _, feriado := range feriados {
		fmt.Printf("%s/%s - %s\n",feriado.Dia, feriado.Mes, feriado.Nome)
	}
}

//FormulaDeGauss é utilizada para calcular o dia da Páscoa. A fórmula vale para anos entre 1901 e 2099. A fórmula pode ser estendida para outros anos, alterando X e Y (criada por Gauss até 1999 e estendida pelo autor até 2299)
func definirPascoa(ano int) time.Time {
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

	domingoPascoa := fmt.Sprintf("%s/%s/%d", fmt.Sprintf("%02d",dia), fmt.Sprintf("%02d",mes), ano)
	formatoData := "02/01/2006"
	dataPascoa, err := time.Parse(formatoData, domingoPascoa)

	if err != nil {
		log.Fatal("Erro ao fazer o parsing da data:", err)
	}

	return dataPascoa 
}

func definirFeriadosMoveis(ano int) []DataFeriado {
	dataPascoa := definirPascoa(ano)
	
	paixaoCristo := dataPascoa.AddDate(0, 0, -2)
	segundaCarnaval := dataPascoa.AddDate(0, 0, -48)
	tercaCarnaval := dataPascoa.AddDate(0, 0, -47)
	corpusChristi := dataPascoa.AddDate(0, 0, 60)

	feriadosMoveis := []DataFeriado{
		{"Paixão de Cristo", paixaoCristo.Format("02"), paixaoCristo.Format("01")},
		{"Carnaval", segundaCarnaval.Format("02"), segundaCarnaval.Format("01")},
		{"Carnaval", tercaCarnaval.Format("02"), tercaCarnaval.Format("01")},
		{"Corpus Christi", corpusChristi.Format("02"), corpusChristi.Format("01")},
	}

	return feriadosMoveis
}