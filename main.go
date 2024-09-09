package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func main() {
	http.HandleFunc("/", calculatorHandler)
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func calculatorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		num1, err1 := strconv.ParseFloat(r.FormValue("num1"), 64)
		operator := r.FormValue("operator")
		num2, err2 := strconv.ParseFloat(r.FormValue("num2"), 64)

		var result string
		if err1 != nil || err2 != nil {
			result = "Invalid number"
		} else {
			switch operator {
			case "+":
				result = fmt.Sprintf("Result: %.2f", num1+num2)
			case "-":
				result = fmt.Sprintf("Result: %.2f", num1-num2)
			case "*":
				result = fmt.Sprintf("Result: %.2f", num1*num2)
			case "/":
				if num2 == 0 {
					result = "Error: division by zero"
				} else {
					result = fmt.Sprintf("Result: %.2f", num1/num2)
				}
			default:
				result = "Error: Invalid operator. Use +, -, *, or /."
			}
		}
		tpl.Execute(w, result)
		return
	}
	tpl.Execute(w, nil)
}
