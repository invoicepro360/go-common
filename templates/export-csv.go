package templates

type ProductCsv struct {
	Name        string `db:"name"`
	Price       string `db:"price"`
	Description string `db:"description"`
	IsTaxable   string `db:"is_taxable"`
	Type        string `db:"type"`
	Status      string `db:"status"`
}

type CustomerCsv struct {
	Customer            string `db:"customer"`
	FirstName           string `db:"first_name"`
	LastName            string `db:"last_name"`
	Email               string `db:"email_address"`
	PhoneNumber         string `db:"phone_number"`
	MobileNumber        string `db:"mobile_number"`
	BillingAddressLine1 string `db:"billing_address_line1"`
	BillingAddressLine2 string `db:"billing_address_line2"`
	BillingCity         string `db:"billing_city"`
	BillingState        string `db:"billing_state"`
	BillingZipcode      string `db:"billing_zipcode"`
	BillingCountry      string `db:"billing_country"`
	AddressLine1        string `db:"shipping_address_line1"`
	AddressLine2        string `db:"shipping_address_line2"`
	City                string `db:"shipping_city"`
	State               string `db:"shipping_state"`
	Zipcode             string `db:"shipping_zipcode"`
	Country             string `db:"shipping_country"`
	Website             string `db:"website"`
	Status              string `db:"status"`
}

type EstimateCsv struct {
	CustomerName    string     `db:"customer_name"`
	Title           string     `db:"title"`
	Subheading      string     `db:"subheading"`
	EstimateNumber  string     `db:"estimate_number"`
	ReferenceNumber string     `db:"reference_number"`
	EstimateDate    string     `db:"estimate_date"`
	ValidityDate    string     `db:"validity_date"`
	Notes           string     `db:"notes"`
	FooterNotes     string     `db:"footer_notes"`
	GrandTotal      string     `db:"grand_total"`
	Status          string     `db:"status"`
	PublishedDate   NullString `db:"published"`
	LastSendDate    NullString `db:"last_send_date"`
}

type InvoiceCsv struct {
	CustomerName    string     `db:"customer_name"`
	Title           string     `db:"title"`
	Subheading      string     `db:"subheading"`
	InvoiceNumber   string     `db:"invoice_number"`
	ReferenceNumber string     `db:"reference_number"`
	InvoiceDate     string     `db:"invoice_date"`
	PaymentDueDate  string     `db:"payment_due_date"`
	Notes           string     `db:"notes"`
	FooterNotes     string     `db:"footer_notes"`
	GrandTotal      string     `db:"grand_total"`
	DueAmount       string     `db:"due_amount"`
	Status          string     `db:"status"`
	PublishedDate   NullString `db:"published"`
	LastSendDate    NullString `db:"last_send_date"`
}
