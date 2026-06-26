package main

import (
	"fmt"
	"html/template"
	"math/cmplx"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

type data struct {
	Result string
	Error  string
}

// a, b, c := 1.0, 4.0, 10.0
// firstRoot, secondRoot, err := complexxQuadraticEquation(a, b, c)
// if err != nil {
// 	fmt.Println("Error:", err)
// } else {
// 	fmt.Printf("Roots: %v, %v\n", firstRoot, secondRoot)
// }

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", home)
	http.HandleFunc("/solve", quadraticHandler)
	fmt.Println("server running on http://localhost:8001")
	http.ListenAndServe(":8001", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl.Execute(w, nil)
}

func quadraticHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	if r.FormValue("a") == "" || r.FormValue("b") == "" || r.FormValue("c") == "" {
		http.Error(w, "Missing input values", http.StatusBadRequest)
		return
	}

	a := r.FormValue("a")
	b := r.FormValue("b")
	c := r.FormValue("c")

	// data := data{
	// 	Result: result,
	// }
	// tmpl.Execute(w, data)

	var aFloat, bFloat, cFloat float64

	fmt.Sscanf(a, "%f", &aFloat)
	fmt.Sscanf(b, "%f", &bFloat)
	fmt.Sscanf(c, "%f", &cFloat)

	firstRoot, secondRoot, err := complexxQuadraticEquation(aFloat, bFloat, cFloat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := fmt.Sprintf("Roots: %v, %v", firstRoot, secondRoot)
	data := data{
		Result: result,
	}

	tmpl.Execute(w, data)

}

func complexxQuadraticEquation(a, b, c float64) (complex128, complex128, error) {
	if a == 0 {
		return 0, 0, fmt.Errorf("coefficient \"a\" cannot be zero")
	}

	A := complex(a, 0)
	B := complex(b, 0)
	C := complex(c, 0)

	discriminant := B*B - 4*A*C

	// if discriminant < 0 {
	// 	return 0, 0, fmt.Errorf("No correct roots")
	// } else if discriminant == 0 {
	// 	sqrtroot := -b / (2 * a)
	// 	return sqrtroot, sqrtroot, nil
	// }

	squarerootDiscriminant := cmplx.Sqrt(discriminant)
	firstRoot := (-B + squarerootDiscriminant) / (2 * A)
	secondRoot := (-B - squarerootDiscriminant) / (2 * A)
	return firstRoot, secondRoot, nil
}