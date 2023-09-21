package model

import (
	"fmt"
	"os"
	"strings"

	"github.com/invoicepro360/go-common/pdf/templates"
	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

var IsDebug bool = false

//connect database
func Connect() *sqlx.DB {

	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")
	// SecuritySalt := os.Getenv("SECURITY_SALT")

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", DBUser, DBPassword, DBHost, DBPort, DBName))

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Fatalf("Failed to create database connection: %s", err.Error())
	}

	return db
}

//
// InvoiceDetailModel Get invoice datail for PDF
func InvoiceDetailModel(uuid string) (pdfdata templates.Pdfdata, err error) {
	db := Connect()

	columns := []string{
		"i.id", "i.uid",
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
			"phoneNumber",c.phone_number
			) as customer`,
	}

	query := fmt.Sprintf(`
		SELECT %s FROM invoices as i 
		LEFT JOIN customers as c on i.customer_id = c.id 
		WHERE i.uid = '%s' group by i.id`,
		strings.Join(columns, ","), uuid)

	if IsDebug {
		log.Infof("DEBUG: QUERY - %s", query)
	}

	rows, err := db.Queryx(query)

	if err != nil {
		log.Error(err.Error())
		return
	}

	defer rows.Close()

	var invoice templates.Invoice
	var invoiceItem templates.InvoiceItem
	var invoiceItems []templates.InvoiceItem

	for rows.Next() {
		err = rows.StructScan(&invoice)

		if err != nil {
			log.Error(err.Error())
			// return
		}
	}

	query = fmt.Sprintf(`
	SELECT  
		id,name,description,invoice_id,product_id,quantity,is_taxable,rate,amount
  	FROM
  		invoice_items
	where invoice_id = %v`, invoice.ID)

	// if IsDebug {
	// 	log.Infof("DEBUG: QUERY - %s", query)
	// }

	rows, err = db.Queryx(query)

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

	// close db connection
	defer db.Close()

	pdfdata.Invoice = invoice
	pdfdata.Items = invoiceItems

	return

}
