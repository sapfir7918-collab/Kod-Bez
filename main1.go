package main

import (
	"fmt"           
	"html/template"
	"net/http"
	"strconv"
)
var dollarRate float64 = 90.0
var euroRate float64 = 100.0
var crossRateEuro float64 = 0.89
type PageData struct {
	Result string 
}
const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Конвертер валют</title>
</head>
<body>
    <h2>Калькулятор валют</h2>
    <form method="POST">
        <label>Сумма:</label>
        <input type="number" step="any" name="amount" required><br><br>

        <label>Действие:</label>
        <select name="choice">
            <option value="1">Из рублей в доллары</option>
            <option value="2">Из рублей в евро</option>
            <option value="3">Из долларов в евро</option>
        </select><br><br>

        <input type="submit" value="Посчитать">
    </form>

    {{if .Result}}
        <h3>{{.Result}}</h3>
    {{end}}
</body>
</html>
`
func convertHandler(w http.ResponseWriter, r *http.Request) {
	var Ruble float64
	var Dollar float64
	var choice int
	var result float64
	var data PageData
	if r.Method == http.MethodPost {
		choice, _ = strconv.Atoi(r.FormValue("choice"))
		if choice == 1 {
			Ruble, _ = strconv.ParseFloat(r.FormValue("amount"), 64)
			result = Ruble / dollarRate 
			data.Result = fmt.Sprintf("Итоговая сумма: %.2f $", result) 
			
		} else if choice == 2 {
			Ruble, _ = strconv.ParseFloat(r.FormValue("amount"), 64)
			
			result = Ruble / euroRate 
			data.Result = fmt.Sprintf("Итоговая сумма: %.2f €", result) 
			
		} else if choice == 3 {
			Dollar, _ = strconv.ParseFloat(r.FormValue("amount"), 64)
			
			result = Dollar * crossRateEuro 
			data.Result = fmt.Sprintf("Итоговая сумма: %.2f €", result) 
		}
	}
tmpl, _ := template.New("webpage").Parse(htmlTemplate) 
	tmpl.Execute(w, data)
}
func main() {
	http.HandleFunc("/", convertHandler)
	println("Сервер запущен! Перейдите по ссылке: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
