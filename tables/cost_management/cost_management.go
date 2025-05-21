package cost_management

import (
	"time"

	"github.com/turbot/tailpipe-plugin-sdk/schema"
)

// CostManagement represents the Azure Cost Management data structure
type CostManagement struct {
	schema.CommonFields

	InvoiceId                    *string                 `json:"invoice_id,omitempty" parquet:"name=invoice_id"`
	PreviousInvoiceId            *string                 `json:"previous_invoice_id,omitempty" parquet:"name=previous_invoice_id"`
	BillingAccountId             *string                 `json:"billing_account_id,omitempty" parquet:"name=billing_account_id"`
	BillingAccountName           *string                 `json:"billing_account_name,omitempty" parquet:"name=billing_account_name"`
	BillingProfileId             *string                 `json:"billing_profile_id,omitempty" parquet:"name=billing_profile_id"`
	BillingProfileName           *string                 `json:"billing_profile_name,omitempty" parquet:"name=billing_profile_name"`
	InvoiceSectionId             *string                 `json:"invoice_section_id,omitempty" parquet:"name=invoice_section_id"`
	InvoiceSectionName           *string                 `json:"invoice_section_name,omitempty" parquet:"name=invoice_section_name"`
	ResellerName                 *string                 `json:"reseller_name,omitempty" parquet:"name=reseller_name"`
	ResellerMpnId                *string                 `json:"reseller_mpn_id,omitempty" parquet:"name=reseller_mpn_id"`
	CostCenter                   *string                 `json:"cost_center,omitempty" parquet:"name=cost_center"`
	BillingPeriodEndDate         *time.Time              `json:"billing_period_end_date,omitempty" parquet:"name=billing_period_end_date"`
	BillingPeriodStartDate       *time.Time              `json:"billing_period_start_date,omitempty" parquet:"name=billing_period_start_date"`
	ServicePeriodEndDate         *time.Time              `json:"service_period_end_date,omitempty" parquet:"name=service_period_end_date"`
	ServicePeriodStartDate       *time.Time              `json:"service_period_start_date,omitempty" parquet:"name=service_period_start_date"`
	Date                         *time.Time              `json:"date,omitempty" parquet:"name=date"`
	ServiceFamily                *string                 `json:"service_family,omitempty" parquet:"name=service_family"`
	ProductOrderId               *string                 `json:"product_order_id,omitempty" parquet:"name=product_order_id"`
	ProductOrderName             *string                 `json:"product_order_name,omitempty" parquet:"name=product_order_name"`
	ConsumedService              *string                 `json:"consumed_service,omitempty" parquet:"name=consumed_service"`
	MeterId                      *string                 `json:"meter_id,omitempty" parquet:"name=meter_id"`
	MeterName                    *string                 `json:"meter_name,omitempty" parquet:"name=meter_name"`
	MeterCategory                *string                 `json:"meter_category,omitempty" parquet:"name=meter_category"`
	MeterSubCategory             *string                 `json:"meter_sub_category,omitempty" parquet:"name=meter_sub_category"`
	MeterRegion                  *string                 `json:"meter_region,omitempty" parquet:"name=meter_region"`
	ProductId                    *string                 `json:"product_id,omitempty" parquet:"name=product_id"`
	ProductName                  *string                 `json:"product_name,omitempty" parquet:"name=product_name"`
	SubscriptionId               *string                 `json:"subscription_id,omitempty" parquet:"name=subscription_id"`
	SubscriptionName             *string                 `json:"subscription_name,omitempty" parquet:"name=subscription_name"`
	PublisherType                *string                 `json:"publisher_type,omitempty" parquet:"name=publisher_type"`
	PublisherId                  *string                 `json:"publisher_id,omitempty" parquet:"name=publisher_id"`
	PublisherName                *string                 `json:"publisher_name,omitempty" parquet:"name=publisher_name"`
	ResourceGroupName            *string                 `json:"resource_group_name,omitempty" parquet:"name=resource_group_name"`
	ResourceId                   *string                 `json:"resource_id,omitempty" parquet:"name=resource_id"`
	ResourceLocation             *string                 `json:"resource_location,omitempty" parquet:"name=resource_location"`
	Location                     *string                 `json:"location,omitempty" parquet:"name=location"`
	EffectivePrice               *float64                `json:"effective_price,omitempty" parquet:"name=effective_price"`
	Quantity                     *float64                `json:"quantity,omitempty" parquet:"name=quantity"`
	UnitOfMeasure                *string                 `json:"unit_of_measure,omitempty" parquet:"name=unit_of_measure"`
	ChargeType                   *string                 `json:"charge_type,omitempty" parquet:"name=charge_type"`
	BillingCurrency              *string                 `json:"billing_currency,omitempty" parquet:"name=billing_currency"`
	PricingCurrency              *string                 `json:"pricing_currency,omitempty" parquet:"name=pricing_currency"`
	CostInBillingCurrency        *float64                `json:"cost_in_billing_currency,omitempty" parquet:"name=cost_in_billing_currency"`
	CostInPricingCurrency        *float64                `json:"cost_in_pricing_currency,omitempty" parquet:"name=cost_in_pricing_currency"`
	CostInUsd                    *float64                `json:"cost_in_usd,omitempty" parquet:"name=cost_in_usd"`
	PaygCostInBillingCurrency    *float64                `json:"payg_cost_in_billing_currency,omitempty" parquet:"name=payg_cost_in_billing_currency"`
	PaygCostInUsd                *float64                `json:"payg_cost_in_usd,omitempty" parquet:"name=payg_cost_in_usd"`
	ExchangeRatePricingToBilling *float64                `json:"exchange_rate_pricing_to_billing,omitempty" parquet:"name=exchange_rate_pricing_to_billing"`
	ExchangeRateDate             *time.Time              `json:"exchange_rate_date,omitempty" parquet:"name=exchange_rate_date"`
	IsAzureCreditEligible        *bool                   `json:"is_azure_credit_eligible,omitempty" parquet:"name=is_azure_credit_eligible"`
	ServiceInfo1                 *string                 `json:"service_info1,omitempty" parquet:"name=service_info1"`
	ServiceInfo2                 *string                 `json:"service_info2,omitempty" parquet:"name=service_info2"`
	AdditionalInfo               *string                 `json:"additional_info,omitempty" parquet:"name=additional_info"`
	Tags                         *map[string]interface{} `json:"tags,omitempty" parquet:"name=tags, type=JSON"`
	PaygPrice                    *float64                `json:"payg_price,omitempty" parquet:"name=payg_price"`
	Frequency                    *string                 `json:"frequency,omitempty" parquet:"name=frequency"`
	Term                         *string                 `json:"term,omitempty" parquet:"name=term"`
	ReservationId                *string                 `json:"reservation_id,omitempty" parquet:"name=reservation_id"`
	ReservationName              *string                 `json:"reservation_name,omitempty" parquet:"name=reservation_name"`
	PricingModel                 *string                 `json:"pricing_model,omitempty" parquet:"name=pricing_model"`
	UnitPrice                    *float64                `json:"unit_price,omitempty" parquet:"name=unit_price"`
	CostAllocationRuleName       *string                 `json:"cost_allocation_rule_name,omitempty" parquet:"name=cost_allocation_rule_name"`
	BenefitId                    *string                 `json:"benefit_id,omitempty" parquet:"name=benefit_id"`
	BenefitName                  *string                 `json:"benefit_name,omitempty" parquet:"name=benefit_name"`
	Provider                     *string                 `json:"provider,omitempty" parquet:"name=provider"`
}

func NewCostManagement() *CostManagement {
	return &CostManagement{
		Tags: &map[string]interface{}{},
	}
}

func (c *CostManagement) GetColumnDescriptions() map[string]string {
	return map[string]string{
		"invoice_id":                       "The unique identifier for the invoice.",
		"previous_invoice_id":              "The identifier for the previous invoice, if applicable.",
		"billing_account_id":               "The identifier for the billing account.",
		"billing_account_name":             "The name of the billing account.",
		"billing_profile_id":               "The identifier for the billing profile.",
		"billing_profile_name":             "The name of the billing profile.",
		"invoice_section_id":               "The identifier for the invoice section.",
		"invoice_section_name":             "The name of the invoice section.",
		"reseller_name":                    "The name of the reseller, if applicable.",
		"reseller_mpn_id":                  "The Microsoft Partner Network identifier for the reseller.",
		"cost_center":                      "The cost center associated with the charge.",
		"billing_period_end_date":          "The end date of the billing period.",
		"billing_period_start_date":        "The start date of the billing period.",
		"service_period_end_date":          "The end date of the service period.",
		"service_period_start_date":        "The start date of the service period.",
		"date":                             "The date associated with the usage or charge.",
		"service_family":                   "The family of Azure services (e.g., Compute, Storage).",
		"product_order_id":                 "The identifier for the product order.",
		"product_order_name":               "The name of the product order.",
		"consumed_service":                 "The name of the consumed Azure service.",
		"meter_id":                         "The identifier for the meter used to measure usage.",
		"meter_name":                       "The name of the meter used to measure usage.",
		"meter_category":                   "The category of the meter (e.g., Virtual Machines, Storage).",
		"meter_sub_category":               "The subcategory of the meter.",
		"meter_region":                     "The region associated with the meter.",
		"product_id":                       "The identifier for the Azure product.",
		"product_name":                     "The name of the Azure product.",
		"subscription_id":                  "The identifier for the Azure subscription.",
		"subscription_name":                "The name of the Azure subscription.",
		"publisher_type":                   "The type of publisher for marketplace offerings.",
		"publisher_id":                     "The identifier for the publisher.",
		"publisher_name":                   "The name of the publisher.",
		"resource_group_name":              "The name of the resource group containing the resource.",
		"resource_id":                      "The unique identifier for the Azure resource.",
		"resource_location":                "The location of the resource.",
		"location":                         "The geographic location associated with the charge.",
		"effective_price":                  "The effective price per unit after discounts.",
		"quantity":                         "The quantity of the resource consumed.",
		"unit_of_measure":                  "The unit of measurement for the quantity (e.g., hours, GB).",
		"charge_type":                      "The type of charge (e.g., Usage, Purchase).",
		"billing_currency":                 "The currency used for billing.",
		"pricing_currency":                 "The currency used for pricing.",
		"cost_in_billing_currency":         "The cost in the billing currency.",
		"cost_in_pricing_currency":         "The cost in the pricing currency.",
		"cost_in_usd":                      "The cost in US dollars.",
		"payg_cost_in_billing_currency":    "The pay-as-you-go cost in the billing currency.",
		"payg_cost_in_usd":                 "The pay-as-you-go cost in US dollars.",
		"exchange_rate_pricing_to_billing": "The exchange rate from pricing currency to billing currency.",
		"exchange_rate_date":               "The date of the exchange rate.",
		"is_azure_credit_eligible":         "Indicates whether the charge is eligible for Azure credits.",
		"service_info1":                    "Additional service information field 1.",
		"service_info2":                    "Additional service information field 2.",
		"additional_info":                  "Additional information about the charge.",
		"tags":                             "Resource tags associated with the resource.",
		"payg_price":                       "The pay-as-you-go price per unit.",
		"frequency":                        "The frequency of the charge (e.g., monthly, one-time).",
		"term":                             "The term of the commitment (e.g., 1 year, 3 years).",
		"reservation_id":                   "The identifier for the reservation, if applicable.",
		"reservation_name":                 "The name of the reservation, if applicable.",
		"pricing_model":                    "The pricing model used (e.g., on-demand, reserved).",
		"unit_price":                       "The price per unit before discounts.",
		"cost_allocation_rule_name":        "The name of the cost allocation rule applied.",
		"benefit_id":                       "The identifier for the benefit applied.",
		"benefit_name":                     "The name of the benefit applied.",
		"provider":                         "The provider of the service or resource.",

		// Override table specific tp_* column descriptions
		"tp_timestamp": "The timestamp representing the date of the usage or charge.",
	}
}
