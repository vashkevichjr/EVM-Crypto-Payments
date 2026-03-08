package models

import "time"

// Business logic constants
const (
	// DefaultInvoiceExpirationDuration is the default expiration time for invoices
	// Invoices expire after 1 hour if not paid
	DefaultInvoiceExpirationDuration = 1 * time.Hour
)
