package generatepdf

import "github.com/go-pdf/fpdf"

//import "github.com/go-pdf/fpdf"

type GeneratePdf struct {
	fileFormat   string // A4, etc...
	oritentation string // P portraint, L landscape
	sizeStr      string // mm...
}

func NewPdfHandler(fileFormat, orientation, sizeStr string) GeneratePdf {
	return GeneratePdf{
		fileFormat:   fileFormat,
		oritentation: orientation,
		sizeStr:      sizeStr,
	}
}

func (g GeneratePdf) NewPdf() fpdf.Pdf {
	pdf := fpdf.New(g.oritentation, g.sizeStr, g.fileFormat, "")
	return pdf
}

func (g GeneratePdf) NewLine()
