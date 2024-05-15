package main

import (
	"fmt"

	"github.com/dslipak/pdf"
	"github.com/go-pdf/fpdf"
)

func main() {

	content, err := readPdf2("./FunProjectsGo/p1.pdf")
	if err != nil {
		panic(err)
	}
	fmt.Println(content)

}

func readPdf2(path string) (string, error) {

	pdfCtx := fpdf.New("P", "mm", "A4", "")
	r, err := pdf.Open(path)

	if err != nil {
		return "", err
	}

	page := r.Page(24)
	pdfCtx.AddPage()

	currentLine := page.Content().Text[0].Y

	complete_word := ""
	for _, v := range page.Content().Text {

		if v.S != " " {
			complete_word += v.S
		} else {
			fmt.Print(complete_word, " ")
			complete_word = ""
			if currentLine != v.Y {

				fmt.Println("")
			}

			currentLine = v.Y

		}

	}

	return "", nil
}
