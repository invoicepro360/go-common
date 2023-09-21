//  package pdf
package pdf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/invoicepro360/go-common/pdf/model"
	log "github.com/sirupsen/logrus"
)

// pdf requestpdf struct
type RequestPdf struct {
	body          string
	styleSheetUrl string
	logoUrl       string
}

// func main() {
// 	GeneratePDF("44C1-6CB0")
// }

func GeneratePDF(uid string) {

	templateName := "modern-template"

	p, err := model.InvoiceDetailModel(uid)

	if err != nil {
		log.Errorf("Failed to fetch invoice detail: %s", err.Error())
	}

	ip, _ := json.Marshal(p)
	fmt.Println(string(ip))

	d := map[string]interface{}{
		"Logo": "logo",
		"Business": p.Business,
		"Invoice": p.Invoice,
		"BilledCustomer": p.Invoice.BilledCustomer,
		"Setting": map[string]interface{}{
			"InvoiceColumns": map[string]interface{}{
				"items": map[string]interface{}{
					"name":  "Item",
					"value": "Item",
					"hide":  false,
				},

				"units": map[string]interface{}{
					"name":  "Quantity",
					"value": "Quantity",
					"hide":  false,
				},
				"price": map[string]interface{}{
					"name":  "Price",
					"value": "Price",
					"hide":  false,
				},
				"taxable": map[string]interface{}{
					"name":  "Taxable",
					"value": "Taxable",
					"hide":  false,
				},
				"amount": map[string]interface{}{
					"name":  "Total",
					"value": "Total",
					"hide":  false,
				},
			},
		},
		"Items": p.Items,
		"SalesTax":p.Invoice.SalesTax,
		"SubTotal":   p.Invoice.SubTotal,
		"GrandTotal": p.Invoice.GrandTotal,
	}

	r := NewRequestPdf("")

	dir, _ := filepath.Abs(".")
	t := time.Now().Unix()

	//html template path
	templatePath := dir + "/invoice-template/" + templateName + ".html"

	//path for download pdf
	outputPath := dir + "/storage/" + strconv.FormatInt(int64(t), 10) + ".pdf"

	if err = r.ParseTemplate(templatePath, d); err == nil {
		// Generate PDF
		ok, _ := r.GeneratePDF(outputPath, []string{})
		fmt.Println(ok, "PDF generated successfully")
	} else {
		fmt.Println(err)
	}
}

// new request to pdf function
func NewRequestPdf(body string) *RequestPdf {
	return &RequestPdf{
		body: body,
	}
}

// parsing template function
func (r *RequestPdf) ParseTemplate(templateFileName string, data interface{}) error {

	// funcMap := template.FuncMap{
	// 	"formatCurrency": func(value float64) string {
	// 		return fmt.Sprintf("%.2f", value)
	// 	},
	// }

	// t, err := template.New("").Funcs(funcMap).ParseFiles(templateFileName)
	t, err := template.ParseFiles(templateFileName)

	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()

	r.styleSheetUrl = "https://techzaa.getappui.com/techinvoice/assets/css/style.css"
	r.logoUrl = "https://techzaa.getappui.com/techinvoice/assets/images/logo.png"

	return nil
}

// generate pdf function
func (r *RequestPdf) GeneratePDF(pdfPath string, args []string) (bool, error) {

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	pdfg.LowQuality.Set(true)
	pdfg.Grayscale.Set(true)
	pdfg.NoPdfCompression.Set(true)

	pdfg.Cover.Allow.Set(r.styleSheetUrl)
	pdfg.Cover.Allow.Set(r.logoUrl)

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(r.body)))

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}
	// TODO: IN FUTURE SAVE IN S3 BUCKET

	err = pdfg.WriteFile(pdfPath)
	if err != nil {
		log.Fatal(err)
	}

	return true, nil
}
