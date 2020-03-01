package transactions

type Response struct {
	Data data `json:"data"`
	Meta meta `json:"meta"`
}

type data struct {
	Payments  []payments `json:"payments"`
	NextUrl   *string    `json:"next_url"`
	Completed *bool      `json:"completed"`
}

type payments struct {
	Id                     *int64  `json:"id"`
	PaymentNo              *string `json:"payment_no"`
	CreatedAt              *string `json:"created_at"`
	UpdatedAt              *string `json:"updated_at"`
	ParentPaymentCreatedAt *string `json:"parent_payment_created_at"`
	TotalCollected         *int64  `json:"total_collected"`
	TotalItemPriceAmount   *int64  `json:"total_item_price_amount"`
	Name                   *string `json:"name"`
	ParentPaymentId        *int64  `json:"parent_payment_id"`
	PaymentType            *string `json:"payment_type"`
	PaymentTypeLabel       *string `json:"payment_type_label"`
	//"customer_id": null,
	PaymentNote        *string `json:"payment_note"`
	Discounts          *int64  `json:"discounts"`
	Subtotal           *int64  `json:"subtotal"`
	gratuities         *int64
	taxes              *int64
	tendered           *int64
	change             *int64
	IncludeGratuityTax *bool
	EnableTax          *bool
	EnableGratuity     *bool `json:"enable_gratuity"`
	//"card_type": null,
	//"card_no": null,
	TransactionDate *string `json:"transaction_date"`
	TransactionTime *string `json:"transaction_time"`
	CollectedBy     *string `json:"collected_by"`
	//"served_by": null,
	SynchronizedAt *string `json:"synchronized_at"`
	OutletId       *bool   `json:"outlet_id"`
	//"pg_mid": null,
	//"mpos_device_id": null,
	//"transaction_certificate": null,
	//"transaction_status_info": null,
	//"mpos_transaction_date": null,
	//"merchant_id": null,
	//"transaction_reference": null,
	//"transaction_number": null,
	//"order_info": null,
	//"cc_name": null,
	//"pii": null,
	Guid   *string `json:"guid"`
	UniqId *string `json:"uniq_id"`
	//"invoice_due_date": null,
	//"invoice_no": null,
	//"invoice_deposit_amount": null,
	//"invoice_payment_records": [],
	//"payment_discounts": [],
	//"payment_taxes": [],
	//"payment_gratuities": [],
	Checkouts []checkouts `json:"checkouts"`
	//"payment_refunds": [],
	IsRefundBreakdown *bool  `json:"is_refund_breakdown"`
	TotalRefund       *int64 `json:"total_refund"`
	IsRefunded        *bool  `json:"is_refunded"`
	RefundAmount      *int64 `json:"refund_amount"`
	//"customer_name": null,
	//"customer_phone": null,
	//"customer_email": null,
	TableName           *string `json:"table_name"`
	TotalRoundingAmount *int64  `json:"total_rounding_amount"`
}

type checkouts struct {
	Id                        *int64  `json:"id"`
	CustomAmount              *int64  `json:"custom_amount"`
	ItemVariantId             *int64  `json:"item_variant_id"`
	Quantity                  *int64  `json:"quantity"`
	DiscountAmount            *int64  `json:"discount_amount"`
	TaxAmount                 *int64  `json:"tax_amount"`
	BusinessId                *int64  `json:"business_id"`
	PaymentId                 *int64  `json:"payment_id"`
	IsDeleted                 *bool   `json:"is_deleted"`
	CreatedAt                 *string `json:"created_at"`
	UpdatedAt                 *string `json:"updated_at"`
	ItemId                    *int64  `json:"item_id"`
	ItemDiscount              *int64  `json:"item_discount"`
	ItemPriceLibrary          *int64  `json:"item_price_library"`
	ItemPrice                 *int64  `json:"item_price"`
	ItemPriceDiscount         *int64  `json:"item_price_discount"`
	GratuityAmount            *int64  `json:"gratuity_amount"`
	ItemPriceDiscountGratuity *int64  `json:"item_price_discount_gratuity"`
	TotalPrice                *int64  `json:"total_price"`
	ItemPriceQuantity         *int64  `json:"item_price_quantity"`
	CategoryName              *string `json:"category_name"`
	CategoryId                *int64  `json:"category_id"`
	ItemName                  *string `json:"item_name"`
	ItemVariantName           *string `json:"item_variant_name"`
	Sku                       *string `json:"sku"`
	//"note": null,
	Cogs       *int64  `json:"cogs"`
	GrossSales *int64  `json:"gross_sales"`
	OutletId   *int64  `json:"outlet_id"`
	Position   *int64  `json:"position"`
	NetSales   *int64  `json:"net_sales"`
	TrackStock *bool   `json:"track_stock"`
	ItemImage  *string `json:"item_image"`
	//"parent_checkout_id": null,
	IsRecipe         *bool   `json:"is_recipe"`
	RedeemAmount     *int64  `json:"redeem_amount"`
	SalesTypeId      *int64  `json:"sales_type_id"`
	SalesTypeName    *string `json:"sales_type_name"`
	IsProgramItem    *bool   `json:"is_program_item"`
	Price            *int64  `json:"price"`
	RefundedQuantity *int64  `json:"refunded_quantity"`
	//"discounts": [],
	//"modifiers": []
}

type meta struct {
	Code   *string `json:"code"`
	Errors *string `json:"errors"`
}
