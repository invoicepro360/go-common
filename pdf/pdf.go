package pdf

import (
	"bytes"
	"encoding/base64"
	"html/template"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/invoicepro360/go-common/config"
	"github.com/invoicepro360/go-common/model"
	log "github.com/sirupsen/logrus"
)

// func main() {

// 	// init config
// 	config.Initialize()

// 	uid := "44C1-6CB0"

// 	base64Pdf, err := GeneratePDF(uid)

// 	if err != nil {
// 		fmt.Println("Failed to generate pdf", err.Error())
// 		return
// 	}
// 	fmt.Println("PDF Generated Successfully", base64Pdf)

// }

//GeneratePDF return base64 pdf string,error
func GeneratePDF(uid string) (pdfBase64 string, err error) {
	config.Initialize()

	pdfData, err := model.InvoiceDetailModel(uid)
	if err != nil {
		log.Errorf("Failed to fetch invoice detail: %s", err.Error())
	}

	// ip, _ := json.Marshal(pdfData)
	// fmt.Println(string(ip))
	// return "", err

	//html template path
	templatePath := config.PdfInvoiceTemplates + pdfData.Setting.InvoiceTemplate + ".html"

	t, err := template.ParseFiles(templatePath)

	if err != nil {
		return "", err
	}

	templateBody := new(bytes.Buffer)

	if err = t.Execute(templateBody, pdfData); err != nil {
		return "", err
	}

	// Generate PDF
	pdfg, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		log.Fatal(err)
	}

	pdfg.Dpi.Set(300)
	//pdfg.Grayscale.Set(true)
	pdfg.LowQuality.Set(true)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)
	//pdfg.Cover.Allow.Set()
	pdfg.Cover.EnableLocalFileAccess.Set(true)

	page := wkhtmltopdf.NewPageReader(strings.NewReader(templateBody.String()))
	// page.Zoom.Set(0.95)

	page.EnableLocalFileAccess.Set(true)

	pdfg.AddPage(page)

	outBuf := new(bytes.Buffer)
	pdfg.SetOutput(outBuf)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	b := pdfg.Bytes()
	if len(b) != 0 {
		log.Errorf("expected to have zero bytes in internal buffer, have %d", len(b))
	}

	b = outBuf.Bytes()
	if len(b) < 3000 {
		log.Errorf("expected to have > 3000 bytes in output buffer, have %d", len(b))
	}

	pdfBase64 = base64.StdEncoding.EncodeToString(b)
	//	fmt.Println(pdfBase64)

	return

}
