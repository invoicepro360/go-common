package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/invoicepro360/go-common/pdf/config"
	"github.com/invoicepro360/go-common/pdf/templates"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

// InvoiceDetailModel Get invoice datail for PDF
func InvoiceDetailModel(uuid string) (pdfdata templates.Pdfdata, err error) {
	db := Connect()

	columns := []string{
		"i.id", "i.uid",
		"i.business_id",
		"i.title", "i.subheading",
		"i.invoice_number", "i.reference_number",
		"i.invoice_date", "i.payment_due_date",
		"i.notes", "i.footer_notes",
		"i.subtotal", "i.grand_total", "i.due_amount",
		"i.sales_tax",
		`JSON_OBJECT(
			"customer",c.customer,
			"firstName",c.first_name,
			"lastName",c.last_name,
			"email",c.email_address,
			"mobileNumber",c.mobile_number,
			"phoneNumber",c.phone_number,
			"billingAddressLine1",c.billing_address_line1,
			"billingAddressLine2",c.billing_address_line2,
			"billingCity",c.billing_city,
			"billingState",c.billing_state,
			"billingZipcode",c.billing_zipcode,
			"shippingAddressLine1",c.shipping_address_line1,
			"shippingAddressLine2",c.shipping_address_line2,
			"shippingCity",c.shipping_city,
			"shippingState",c.shipping_state,
			"shippingZipcode",c.shipping_zipcode
			) as customer`,
		`JSON_OBJECT(
			"companyName",b.name,
			"addressLine1",b.address_line1,
			"addressLine2",b.address_line2,			
			"phoneNumber",b.phone_number
			) as business`,
	}

	query := fmt.Sprintf(`
		SELECT %s FROM invoices as i 
		LEFT JOIN customers as c on i.customer_id = c.id 
		LEFT jOIN businesses as b on i.business_id = b.id 
		WHERE i.uid = '%s' group by i.id`,
		strings.Join(columns, ","), uuid)

	if config.IsDebug {
		log.Infof("DEBUG: QUERY - %s", query)
	}

	rows, err := db.Queryx(query)

	if err != nil {
		log.Error(err.Error())
		return
	}

	defer rows.Close()

	var invoice templates.Invoice

	for rows.Next() {
		err = rows.StructScan(&invoice)

		if err != nil {
			log.Error(err.Error())
			// return
		}
	}

	//invoice items
	invoiceItems, err := GetInvoiceItems(invoice.ID, db)

	if err != nil {
		log.Error(err.Error())
		return
	}

	//invoice settings
	settings, err := GetSettings(invoice.BusinessId, db)

	if err != nil {
		log.Error(err.Error())
		return
	}

	//invoice payments
	payments, err := GetPayments(invoice.UID, invoice.BusinessId, settings.DateFormat, db)

	if err != nil {
		log.Error(err.Error())
		return
	}

	// close db connection
	defer db.Close()

	invoiceDate, err := time.Parse("2006-01-02", invoice.InvoiceDate)

	if err == nil {
		invoice.InvoiceDate = invoiceDate.Format(config.DateFormat[settings.DateFormat])
	} else {
		log.Error(err.Error())
	}

	paymentDueDate, err := time.Parse("2006-01-02", invoice.PaymentDueDate)

	if err == nil {
		invoice.PaymentDueDate = paymentDueDate.Format(config.DateFormat[settings.DateFormat])
	} else {
		log.Error(err.Error())
	}

	pdfdata.Invoice = invoice
	pdfdata.Items = invoiceItems
	pdfdata.Setting = settings
	pdfdata.Payments = payments

	return

}

func GetInvoiceItems(invoiceId int, db *sqlx.DB) (invoiceItems []templates.InvoiceItem, err error) {

	var invoiceItem templates.InvoiceItem

	query := `
		SELECT  
			id,name,description,invoice_id,product_id,quantity,is_taxable,rate,amount
		FROM
			invoice_items
		where invoice_id = :invoiceId`

	rows, err := db.NamedQuery(query, map[string]interface{}{
		"invoiceId": invoiceId,
	})

	if err != nil {
		log.Error(err.Error())
		return
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&invoiceItem)

		if err != nil {
			log.Error(err.Error())
		}
		invoiceItems = append(invoiceItems, invoiceItem)
	}

	return
}

// GetSettings buisness settings
func GetSettings(businessId int, db *sqlx.DB) (setting templates.Setting, err error) {

	query := `SELECT  
				invoice_template,logo,accent_color,invoice_columns,date_format,currency,currency_symbol 
  				FROM settings 
				where business_id = :businessId`

	if config.IsDebug {
		log.Infof("DEBUG: QUERY - %s", query)
	}

	rows, err := db.NamedQuery(query, map[string]interface{}{
		"businessId": businessId,
	})

	if err != nil {
		log.Error(err.Error())
		return
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&setting)

		if err != nil {
			log.Error(err.Error())
		}
	}

	return
}

// GetPayments list of  invoice payments
func GetPayments(invoiceId string, businessId int, DateFmt string, db *sqlx.DB) (payments []templates.Payment, err error) {

	var payment templates.Payment

	// Execute the query
	columns := []string{
		"id",
		"payment_date",
		"payment_method",
		"amount",
		"memo",
		"created",
	}

	query := fmt.Sprintf("SELECT %s FROM payments where invoice_id = :invoiceId and business_id = :businessId", strings.Join(columns, ","))

	if config.IsDebug {
		log.Infof("DEBUG: QUERY - %s", query)
	}

	rows, err := db.NamedQuery(query, map[string]interface{}{
		"invoiceId":  invoiceId,
		"businessId": businessId,
	})

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.StructScan(&payment)

		if err != nil {
			log.Error(err.Error())
		}

		PaymentDate, err := time.Parse("2006-01-02 00:00:00", payment.PaymentDate)

		if err == nil {
			payment.PaymentDate = PaymentDate.Format(config.DateFormat[DateFmt])
		} else {
			log.Error(err.Error())

		}

		payments = append(payments, payment)
	}

	return
}
