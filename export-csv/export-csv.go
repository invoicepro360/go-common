// package exportcsv

package main

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"log"

	"github.com/invoicepro360/go-common/config"
	"github.com/invoicepro360/go-common/model"
)

// func main() {
// config.Initialize()

// GenerateCsv(1,"invoice")
// GenerateCsv(1, "product")
// GenerateCsv(1, "customer")
// GenerateCsv(1, "estimate")

// }

func GenerateCsv(businessId int, t string) (pdfBase64 string, err error) {
	config.Initialize()

	if t == "invoice" {
		pdfBase64, err = InvoiceCsv(businessId)
	}

	if t == "product" {
		pdfBase64, err = ProductCsv(businessId)
	}

	if t == "customer" {
		pdfBase64, err = CustomerCsv(businessId)
	}

	if t == "estimate" {
		pdfBase64, err = EstimateCsv(businessId)
	}

	return
}

func ProductCsv(businessId int) (pdfBase64 string, err error) {

	products, err := model.GetProductsModel(businessId)

	if err != nil {
		return
	}

	buf := new(bytes.Buffer)
	w := csv.NewWriter(buf)
	bb := [][]string{}

	for _, v := range products {
		r := make([]string, 0)
		r = append(r, v.Name)
		r = append(r, v.Description)
		r = append(r, v.Price)
		r = append(r, v.IsTaxable)
		r = append(r, v.Status)
		r = append(r, v.Type)

		bb = append(bb, r)
	}

	w.WriteAll(bb)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	pdfBase64 = base64.StdEncoding.EncodeToString(buf.Bytes())

	// fmt.Println(buf.String())

	return
}

func CustomerCsv(businessId int) (pdfBase64 string, err error) {

	customers, err := model.GetCustomersModel(businessId)

	if err != nil {
		return
	}

	buf := new(bytes.Buffer)
	w := csv.NewWriter(buf)
	bb := [][]string{}

	for _, v := range customers {
		r := make([]string, 0)
		r = append(r, v.Customer)
		r = append(r, v.FirstName)
		r = append(r, v.LastName)
		r = append(r, v.Email)
		r = append(r, v.PhoneNumber)
		r = append(r, v.MobileNumber)
		r = append(r, v.AddressLine1)
		r = append(r, v.AddressLine2)
		r = append(r, v.City)
		r = append(r, v.State)
		r = append(r, v.Country)
		r = append(r, v.BillingAddressLine1)
		r = append(r, v.BillingAddressLine2)
		r = append(r, v.BillingCity)
		r = append(r, v.BillingState)
		r = append(r, v.BillingZipcode)
		r = append(r, v.BillingCountry)
		r = append(r, v.Website)
		r = append(r, v.Status)

		bb = append(bb, r)
	}

	w.WriteAll(bb)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	pdfBase64 = base64.StdEncoding.EncodeToString(buf.Bytes())

	// fmt.Println(buf.String())

	return
}

func EstimateCsv(businessId int) (pdfBase64 string, err error) {

	estimates, err := model.GetEstimatesModel(businessId)

	if err != nil {
		return
	}

	buf := new(bytes.Buffer)
	w := csv.NewWriter(buf)
	bb := [][]string{}

	for _, v := range estimates {
		r := make([]string, 0)
		r = append(r, v.CustomerName)
		r = append(r, v.EstimateNumber)
		r = append(r, v.Title)
		r = append(r, v.Subheading)
		r = append(r, v.ReferenceNumber)
		r = append(r, v.EstimateDate)
		r = append(r, v.ValidityDate)
		r = append(r, v.Notes)
		r = append(r, v.FooterNotes)
		r = append(r, v.Status)
		r = append(r, v.PublishedDate.String)
		r = append(r, v.LastSendDate.String)
		r = append(r, v.GrandTotal)

		bb = append(bb, r)
	}

	w.WriteAll(bb)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	pdfBase64 = base64.StdEncoding.EncodeToString(buf.Bytes())

	fmt.Println(buf.String())

	return
}

func InvoiceCsv(businessId int) (pdfBase64 string, err error) {

	invoices, err := model.GetInvoicesModel(businessId)

	if err != nil {
		return
	}

	b := new(bytes.Buffer)
	w := csv.NewWriter(b)
	bb := [][]string{}

	for _, v := range invoices {
		r := make([]string, 0)
		r = append(r, v.CustomerName)
		r = append(r, v.InvoiceNumber)

		r = append(r, v.Title)
		r = append(r, v.Subheading)
		r = append(r, v.ReferenceNumber)
		r = append(r, v.InvoiceDate)
		r = append(r, v.PaymentDueDate)
		r = append(r, v.Notes)
		r = append(r, v.FooterNotes)
		r = append(r, v.Status)
		r = append(r, v.PublishedDate.String)
		r = append(r, v.LastSendDate.String)
		r = append(r, v.GrandTotal)
		r = append(r, v.DueAmount)

		bb = append(bb, r)
	}

	w.WriteAll(bb)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	pdfBase64 = base64.StdEncoding.EncodeToString(b.Bytes())

	fmt.Println(b.String())

	return
}
