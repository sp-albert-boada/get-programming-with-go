package main

import "fmt"

type fahrenheit float64
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

type celsius float64
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0/5.0) + 32.0)
}

type converterFn func (value float64) float64

func main() {
	from := -40.0
	to := 100.0
	step := 5.0

	const celsiusHeader = "ºC"
	const fahrenheitHeader = "ºF"

	drawTable(from, to, step, celsiusHeader, fahrenheitHeader, c2f)
	drawTable(from, to, step, fahrenheitHeader, celsiusHeader, f2c)
}

func drawTable(
	from float64,
	to float64,
	step float64,
	firstColumn string,
	secondColum string,
	converter converterFn,
) {
	drawHeader(firstColumn, secondColum)

	for from <= to {
		drawRow(from, converter(from))
		from+=step
	}

	drawFooter()
}

func c2f(value float64) float64 {
	c := celsius(value)

	return float64(c.fahrenheit())
}

func f2c(value float64) float64 {
	f := fahrenheit(value)

	return float64(f.celsius())
}

func drawHeader(firstColumn, secondColumn string) {
	drawTableSeparator()
	fmt.Printf("| %8s | %8s |\n", firstColumn, secondColumn)
	drawTableSeparator()
}

func drawRow(value1 float64, value2 float64) {
	fmt.Printf("| %8.2f | %8.2f |\n", value1, value2)
}

func drawFooter() {
	drawTableSeparator()
}

func drawTableSeparator() {
	fmt.Println("=======================")
}