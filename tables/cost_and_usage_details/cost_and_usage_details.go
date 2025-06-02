package cost_and_usage_details

import (
	"time"

	"github.com/turbot/tailpipe-plugin-sdk/schema"
)

// CostAndUsageDetails represents the Azure Cost and Usage Details data structure
type CostAndUsageDetails struct {
	schema.CommonFields

	AdditionalInfo               *string                 `json:"additional_info,omitempty" parquet:"name=additional_info"`
	BenefitId                    *string                 `json:"benefit_id,omitempty" parquet:"name=benefit_id"`
	BenefitName                  *string                 `json:"benefit_name,omitempty" parquet:"name=benefit_name"`
	BillingAccountId             *string                 `json:"billing_account_id,omitempty" parquet:"name=billing_account_id"`
	BillingAccountName           *string                 `json:"billing_account_name,omitempty" parquet:"name=billing_account_name"`
	BillingCurrency              *string                 `json:"billing_currency,omitempty" parquet:"name=billing_currency"`
	BillingPeriodEndDate         *time.Time              `json:"billing_period_end_date,omitempty" parquet:"name=billing_period_end_date"`
	BillingPeriodStartDate       *time.Time              `json:"billing_period_start_date,omitempty" parquet:"name=billing_period_start_date"`
	BillingProfileId             *string                 `json:"billing_profile_id,omitempty" parquet:"name=billing_profile_id"`
	BillingProfileName           *string                 `json:"billing_profile_name,omitempty" parquet:"name=billing_profile_name"`
	ChargeType                   *string                 `json:"charge_type,omitempty" parquet:"name=charge_type"`
	ConsumedService              *string                 `json:"consumed_service,omitempty" parquet:"name=consumed_service"`
	CostAllocationRuleName       *string                 `json:"cost_allocation_rule_name,omitempty" parquet:"name=cost_allocation_rule_name"`
	CostCenter                   *string                 `json:"cost_center,omitempty" parquet:"name=cost_center"`
	CostInBillingCurrency        *float64                `json:"cost_in_billing_currency,omitempty" parquet:"name=cost_in_billing_currency"`
	CostInPricingCurrency        *float64                `json:"cost_in_pricing_currency,omitempty" parquet:"name=cost_in_pricing_currency"`
	CostInUsd                    *float64                `json:"cost_in_usd,omitempty" parquet:"name=cost_in_usd"`
	Date                         *time.Time              `json:"date,omitempty" parquet:"name=date"`
	EffectivePrice               *float64                `json:"effective_price,omitempty" parquet:"name=effective_price"`
	ExchangeRateDate             *time.Time              `json:"exchange_rate_date,omitempty" parquet:"name=exchange_rate_date"`
	ExchangeRatePricingToBilling *float64                `json:"exchange_rate_pricing_to_billing,omitempty" parquet:"name=exchange_rate_pricing_to_billing"`
	Frequency                    *string                 `json:"frequency,omitempty" parquet:"name=frequency"`
	InvoiceId                    *string                 `json:"invoice_id,omitempty" parquet:"name=invoice_id"`
	InvoiceSectionId             *string                 `json:"invoice_section_id,omitempty" parquet:"name=invoice_section_id"`
	InvoiceSectionName           *string                 `json:"invoice_section_name,omitempty" parquet:"name=invoice_section_name"`
	IsAzureCreditEligible        *bool                   `json:"is_azure_credit_eligible,omitempty" parquet:"name=is_azure_credit_eligible"`
	Location                     *string                 `json:"location,omitempty" parquet:"name=location"`
	MeterCategory                *string                 `json:"meter_category,omitempty" parquet:"name=meter_category"`
	MeterId                      *string                 `json:"meter_id,omitempty" parquet:"name=meter_id"`
	MeterName                    *string                 `json:"meter_name,omitempty" parquet:"name=meter_name"`
	MeterRegion                  *string                 `json:"meter_region,omitempty" parquet:"name=meter_region"`
	MeterSubCategory             *string                 `json:"meter_sub_category,omitempty" parquet:"name=meter_sub_category"`
	PaygCostInBillingCurrency    *float64                `json:"payg_cost_in_billing_currency,omitempty" parquet:"name=payg_cost_in_billing_currency"`
	PaygCostInUsd                *float64                `json:"payg_cost_in_usd,omitempty" parquet:"name=payg_cost_in_usd"`
	PaygPrice                    *float64                `json:"payg_price,omitempty" parquet:"name=payg_price"`
	PreviousInvoiceId            *string                 `json:"previous_invoice_id,omitempty" parquet:"name=previous_invoice_id"`
	PricingCurrency              *string                 `json:"pricing_currency,omitempty" parquet:"name=pricing_currency"`
	PricingModel                 *string                 `json:"pricing_model,omitempty" parquet:"name=pricing_model"`
	ProductId                    *string                 `json:"product_id,omitempty" parquet:"name=product_id"`
	ProductName                  *string                 `json:"product_name,omitempty" parquet:"name=product_name"`
	ProductOrderId               *string                 `json:"product_order_id,omitempty" parquet:"name=product_order_id"`
	ProductOrderName             *string                 `json:"product_order_name,omitempty" parquet:"name=product_order_name"`
	Provider                     *string                 `json:"provider,omitempty" parquet:"name=provider"`
	PublisherId                  *string                 `json:"publisher_id,omitempty" parquet:"name=publisher_id"`
	PublisherName                *string                 `json:"publisher_name,omitempty" parquet:"name=publisher_name"`
	PublisherType                *string                 `json:"publisher_type,omitempty" parquet:"name=publisher_type"`
	Quantity                     *float64                `json:"quantity,omitempty" parquet:"name=quantity"`
	ReservationId                *string                 `json:"reservation_id,omitempty" parquet:"name=reservation_id"`
	ReservationName              *string                 `json:"reservation_name,omitempty" parquet:"name=reservation_name"`
	ResellerMpnId                *string                 `json:"reseller_mpn_id,omitempty" parquet:"name=reseller_mpn_id"`
	ResellerName                 *string                 `json:"reseller_name,omitempty" parquet:"name=reseller_name"`
	ResourceGroupName            *string                 `json:"resource_group_name,omitempty" parquet:"name=resource_group_name"`
	ResourceId                   *string                 `json:"resource_id,omitempty" parquet:"name=resource_id"`
	ResourceLocation             *string                 `json:"resource_location,omitempty" parquet:"name=resource_location"`
	ServiceFamily                *string                 `json:"service_family,omitempty" parquet:"name=service_family"`
	ServiceInfo1                 *string                 `json:"service_info1,omitempty" parquet:"name=service_info1"`
	ServiceInfo2                 *string                 `json:"service_info2,omitempty" parquet:"name=service_info2"`
	ServicePeriodEndDate         *time.Time              `json:"service_period_end_date,omitempty" parquet:"name=service_period_end_date"`
	ServicePeriodStartDate       *time.Time              `json:"service_period_start_date,omitempty" parquet:"name=service_period_start_date"`
	SubscriptionId               *string                 `json:"subscription_id,omitempty" parquet:"name=subscription_id"`
	SubscriptionName             *string                 `json:"subscription_name,omitempty" parquet:"name=subscription_name"`
	Tags                         *map[string]interface{} `json:"tags,omitempty" parquet:"name=tags, type=JSON"`
	Term                         *string                 `json:"term,omitempty" parquet:"name=term"`
	UnitOfMeasure                *string                 `json:"unit_of_measure,omitempty" parquet:"name=unit_of_measure"`
	UnitPrice                    *float64                `json:"unit_price,omitempty" parquet:"name=unit_price"`
}

func NewCostAndUsageDetails() *CostAndUsageDetails {
	return &CostAndUsageDetails{
		Tags: &map[string]interface{}{},
	}
}

func (c *CostAndUsageDetails) GetColumnDescriptions() map[string]string {
	return map[string]string{
		"additional_info":                  "Additional information about the charge.",
		"benefit_id":                       "The identifier for the benefit applied.",
		"benefit_name":                     "The name of the benefit applied.",
		"billing_account_id":               "The identifier for the billing account.",
		"billing_account_name":             "The name of the billing account.",
		"billing_currency":                 "The currency used for billing.",
		"billing_period_end_date":          "The end date of the billing period.",
		"billing_period_start_date":        "The start date of the billing period.",
		"billing_profile_id":               "The identifier for the billing profile.",
		"billing_profile_name":             "The name of the billing profile.",
		"charge_type":                      "The type of charge (e.g., Usage, Purchase).",
		"consumed_service":                 "The name of the consumed Azure service.",
		"cost_allocation_rule_name":        "The name of the cost allocation rule applied.",
		"cost_center":                      "The cost center associated with the charge.",
		"cost_in_billing_currency":         "The cost in the billing currency.",
		"cost_in_pricing_currency":         "The cost in the pricing currency.",
		"cost_in_usd":                      "The cost in US dollars.",
		"date":                             "The date associated with the usage or charge.",
		"effective_price":                  "The effective price per unit after discounts.",
		"exchange_rate_date":               "The date of the exchange rate.",
		"exchange_rate_pricing_to_billing": "The exchange rate from pricing currency to billing currency.",
		"frequency":                        "The frequency of the charge (e.g., monthly, one-time).",
		"invoice_id":                       "The unique identifier for the invoice.",
		"invoice_section_id":               "The identifier for the invoice section.",
		"invoice_section_name":             "The name of the invoice section.",
		"is_azure_credit_eligible":         "Indicates whether the charge is eligible for Azure credits.",
		"location":                         "The geographic location associated with the charge.",
		"meter_category":                   "The category of the meter (e.g., Virtual Machines, Storage).",
		"meter_id":                         "The identifier for the meter used to measure usage.",
		"meter_name":                       "The name of the meter used to measure usage.",
		"meter_region":                     "The region associated with the meter.",
		"meter_sub_category":               "The subcategory of the meter.",
		"payg_cost_in_billing_currency":    "The pay-as-you-go cost in the billing currency.",
		"payg_cost_in_usd":                 "The pay-as-you-go cost in US dollars.",
		"payg_price":                       "The pay-as-you-go price per unit.",
		"previous_invoice_id":              "The identifier for the previous invoice, if applicable.",
		"pricing_currency":                 "The currency used for pricing.",
		"pricing_model":                    "The pricing model used (e.g., on-demand, reserved).",
		"product_id":                       "The identifier for the Azure product.",
		"product_name":                     "The name of the Azure product.",
		"product_order_id":                 "The identifier for the product order.",
		"product_order_name":               "The name of the product order.",
		"provider":                         "The provider of the service or resource.",
		"publisher_id":                     "The identifier for the publisher.",
		"publisher_name":                   "The name of the publisher.",
		"publisher_type":                   "The type of publisher for marketplace offerings.",
		"quantity":                         "The quantity of the resource consumed.",
		"reservation_id":                   "The identifier for the reservation, if applicable.",
		"reservation_name":                 "The name of the reservation, if applicable.",
		"reseller_mpn_id":                  "The Microsoft Partner Network identifier for the reseller.",
		"reseller_name":                    "The name of the reseller, if applicable.",
		"resource_group_name":              "The name of the resource group containing the resource.",
		"resource_id":                      "The unique identifier for the Azure resource.",
		"resource_location":                "The location of the resource.",
		"service_family":                   "The family of Azure services (e.g., Compute, Storage).",
		"service_info1":                    "Additional service information field 1.",
		"service_info2":                    "Additional service information field 2.",
		"service_period_end_date":          "The end date of the service period.",
		"service_period_start_date":        "The start date of the service period.",
		"subscription_id":                  "The identifier for the Azure subscription.",
		"subscription_name":                "The name of the Azure subscription.",
		"tags":                             "Resource tags associated with the resource.",
		"term":                             "The term of the commitment (e.g., 1 year, 3 years).",
		"unit_of_measure":                  "The unit of measurement for the quantity (e.g., hours, GB).",
		"unit_price":                       "The price per unit before discounts.",

		// Override table specific tp_* column descriptions
		"tp_timestamp": "The timestamp representing the date of the usage or charge.",
	}
}
