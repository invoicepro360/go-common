package model

import (
	"fmt"
	"strings"

	"github.com/invoicepro360/go-common/config"
	"github.com/invoicepro360/go-common/templates"
	log "github.com/sirupsen/logrus"
)

func GetProductsModel(businessId int) (records []templates.ProductCsv, err error) {
	db := Connect()

	columns := []string{
		"name",
		"description",
		"price",
		"is_taxable",
		"type",
		"status",
	}

	query := fmt.Sprintf(`
		SELECT %s FROM products WHERE business_id=%v ORDER BY id desc`,
		strings.Join(columns, ","), businessId)

	if config.IsDebug {
		log.Infof("DEBUG: QUERY - %s", query)
	}

	rows, err := db.Queryx(query)

	if err != nil {
		log.Error(err.Error())
		return
	}

	defer rows.Close()

	var product templates.ProductCsv

	records = append(records, templates.ProductCsv{
		Name:        "Name",
		Description: "Description",
		Type:        "Type",
		IsTaxable:   "Taxable",
		Status:      "Status",
		Price:       "Price",
	})

	for rows.Next() {

		err = rows.StructScan(&product)
		if err != nil {
			log.Error(err.Error())
			// return
		}
		records = append(records, product)
	}

	return

}

func GetCustomersModel(businessId int) (records []templates.CustomerCsv, err error) {
	db := Connect()

	columns := []string{
		"customer",
		"first_name",
		"last_name",
		"email_address",
		"phone_number",
		"billing_address_line1",
		"billing_address_line2",
		"billing_city",
		"billing_state",
		"billing_zipcode",
		"billing_country",
		"shipping_address_line1",
		"shipping_address_line2",
		"shipping_city",
		"shipping_state",
		"shipping_zipcode",
		"shipping_country",
		"mobile_number",
		"website",
		"status",
	}

	query := fmt.Sprintf(`
		SELECT %s FROM customers WHERE business_id=%v ORDER BY id desc`,
		strings.Join(columns, ","), businessId)

	if config.IsDebug {
		log.Infof("DEBUG: QUERY - %s", query)
	}

	rows, err := db.Queryx(query)

	if err != nil {
		log.Error(err.Error())
		return
	}

	defer rows.Close()

	var customer templates.CustomerCsv

	records = append(records, templates.CustomerCsv{
		Customer:            "Customer",
		FirstName:           "First Name",
		LastName:            "Last Name",
		Email:               "Email",
		PhoneNumber:         "Phone Number",
		MobileNumber:        "Mobile Number",
		AddressLine1:        "Shipping Address",
		AddressLine2:        "Shipping Address 2",
		City:                "Shipping City",
		State:               "Shipping State",
		Zipcode:             "Shipping Zipcode",
		Country:             "Shipping Country",
		BillingAddressLine1: "Billing Address",
		BillingAddressLine2: "Billing Address 2",
		BillingCity:         "Billing City",
		BillingState:        "Billing State",
		BillingZipcode:      "Billing Zipcode",
		BillingCountry:      "Billing Country",
		Website:             "Website",
		Status:              "Status",
	})

	for rows.Next() {

		err = rows.StructScan(&customer)
		if err != nil {
			log.Error(err.Error())
			// return
		}
		records = append(records, customer)
	}

	return

}

func GetEstimatesModel(businessId int) (records []templates.EstimateCsv, err error) {

	db := Connect()

	columns := []string{
		"c.customer as customer_name",
		"e.estimate_number as estimate_number",
		"e.title as title",
		"e.subheading as subheading",
		"e.reference_number as reference_number",
		"e.estimate_date as estimate_date",
		"e.validity_date as validity_date",
		"e.notes as notes",
		"e.footer_notes as footer_notes",
		"e.status as status",
		"e.published as published",
		"e.last_send_date as last_send_date",
		"e.grand_total as grand_total",
	}

	query := fmt.Sprintf(`
		SELECT %s FROM estimates as e 
		LEFT JOIN (select id,customer from customers) as c on e.customer_id = c.id
		WHERE e.business_id=%v ORDER BY e.id desc`,
		strings.Join(columns, ","), businessId)

	if config.IsDebug {
		log.Infof("DEBUG: QUERY - %s", query)
	}

	rows, err := db.Queryx(query)

	if err != nil {
		log.Error(err.Error())
		return
	}

	defer rows.Close()

	var estimate templates.EstimateCsv

	records = append(records, templates.EstimateCsv{
		CustomerName:    "Customer Name",
		EstimateNumber:  "Estimate Number",
		Title:           "Estimate Title",
		Subheading:      "Sub Heading",
		ReferenceNumber: "Reference Number",
		EstimateDate:    "Estimate Date",
		ValidityDate:    "Validity Date",
		Notes:           "Estimate Notes",
		FooterNotes:     "Footer Notes",
		Status:          "Estimate Status",
		PublishedDate:   templates.NullString{String: "Publish Date"},
		LastSendDate:    templates.NullString{String: "Last Send Date"},
		GrandTotal:      "Grand Total",
	})

	for rows.Next() {

		err = rows.StructScan(&estimate)
		if err != nil {
			log.Error(err.Error())
			// return
		}
		records = append(records, estimate)
	}

	return
}

func GetInvoicesModel(businessId int) (records []templates.InvoiceCsv, err error) {

	db := Connect()


	columns := []string{
		"c.customer as customer_name",
		"i.invoice_number as invoice_number",
		"i.title as title",
		"i.subheading as subheading",
		"i.reference_number as reference_number",
		"i.invoice_date as invoice_date",
		"i.payment_due_date as payment_due_date",
		"i.notes as notes",
		"i.footer_notes as footer_notes",
		"i.status as status",
		"i.published as published",
		"i.last_send_date as last_send_date",
		"i.grand_total as grand_total",
		"i.due_amount as due_amount",		
	}

	query := fmt.Sprintf(`
		SELECT %s FROM invoices as i 
		LEFT JOIN (select id,customer from customers) as c on i.customer_id = c.id
		WHERE i.business_id=%v ORDER BY i.id desc`,
		strings.Join(columns, ","), businessId)

	if config.IsDebug {
		log.Infof("DEBUG: QUERY - %s", query)
	}

	rows, err := db.Queryx(query)

	if err != nil {
		log.Error(err.Error())
		return
	}

	defer rows.Close()

	var invoice templates.InvoiceCsv

	records = append(records, templates.InvoiceCsv{
		CustomerName:    "Customer Name",
		InvoiceNumber:   "Invoice Number",
		Title:           "Invoice Title",
		Subheading:      "Sub Heading",
		ReferenceNumber: "Reference Number",
		InvoiceDate:     "Invoice Date",
		PaymentDueDate:  "Payment Due Date",
		Notes:           "Invoice Notes",
		FooterNotes:     "Footer Notes",
		Status:          "Invoice Status",
		PublishedDate:   templates.NullString{String: "Publish Date"},
		LastSendDate:    templates.NullString{String: "Invoice Date"},
		GrandTotal:      "Grand Total",
		DueAmount:       "Due Amount",
	})

	for rows.Next() {

		err = rows.StructScan(&invoice)
		if err != nil {
			log.Error(err.Error())
			// return
		}
		records = append(records, invoice)
	}

	return
}
