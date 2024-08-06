package templates

import "time"

type Invoice struct {
	ID              int           `db:"id"`
	BusinessId      int           `db:"business_id"`
	UID             string        `db:"uid"`
	InvoiceTitle    string        `db:"title"`
	InvoiceSummary  string        `db:"subheading"`
	InvoiceNumber   string        `db:"invoice_number"`
	ReferenceNumber string        `db:"reference_number"`
	InvoiceDate     string        `db:"invoice_date"`
	PaymentDueDate  string        `db:"payment_due_date"`
	Notes           string        `db:"notes"`
	FooterNotes     string        `db:"footer_notes"`
	SubTotal        string        `db:"subtotal"`
	GrandTotal      string        `db:"grand_total"`
	DueAmount       string        `db:"due_amount"`
	SalesTax        SalesTaxField `db:"sales_tax"`
	BilledCustomer  JsonField     `db:"customer"`
	Business        JsonField     `db:"business"`
}

type SalesTax struct {
	ID              int     `json:"id"`
	TaxName         string  `json:"taxName"`
	TaxAbbreviation string  `json:"taxAbbreviation"`
	TaxNumber       string  `json:"taxNumber"`
	TaxRate         int     `json:"taxRate"`
	TaxAmount       float64 `json:"taxAmount"`
}

type InvoiceItem struct {
	Id          int       `json:"uid" db:"id"`
	InvoiceId   int       `json:"-" db:"invoice_id"`
	ProductId   int       `json:"item_id" db:"product_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Quantity    int       `json:"quantity" db:"quantity"`
	IsTaxable   int       `json:"is_taxable" db:"is_taxable"`
	Rate        float64   `json:"price" db:"rate"`
	Amount      float64   `json:"total" db:"amount"`
	Created     time.Time `json:"-" db:"created"`
	Modified    time.Time `json:"-" db:"modified"`
	ModifiedBy  string    `json:"-" db:"modified_by"`
}

type Setting struct {
	InvoiceTemplate string    `json:"invoice_template" db:"invoice_template"`
	Logo            string    `json:"logo" db:"logo"`
	AccentColor     string    `json:"accent_color" db:"accent_color"`
	InvoiceColumns  JsonField `json:"invoice_columns" db:"invoice_columns"`
	DateFormat      string    `json:"date_format" db:"date_format"`
	Currency        string    `json:"currency" db:"currency"`
	CurrencySymbol  string    `json:"currency_symbol" db:"currency_symbol"`
}

type Payment struct {
	ID            int        `json:"id" db:"id"`
	PaymentDate   string     `json:"paymentDate" db:"payment_date"`
	PaymentMethod string     `json:"paymentMethod" db:"payment_method"`
	Memo          NullString `json:"memo" db:"memo"`
	Amount        amount     `json:"amount" db:"amount"`
	Created       string     `json:"created" db:"created"`
}

type Pdfdata struct {
	Invoice  Invoice
	Setting  Setting
	Items    []InvoiceItem
	Payments []Payment
}
