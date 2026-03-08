package models

import "time"

// Business logic constants
const (
	// Invoices expire after 1 hour if not paid
	DefaultInvoiceExpirationDuration = 1 * time.Hour
)
