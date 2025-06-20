package cost_and_usage_actual

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/turbot/tailpipe-plugin-sdk/error_types"
	"github.com/turbot/tailpipe-plugin-sdk/mappers"
)

// CostAndUsageActualMapper is responsible for mapping CSV rows to CostAndUsageActual structs
type CostAndUsageActualMapper struct {
	headers []string
}

// NewCostAndUsageActualMapper creates a new instance of CostAndUsageActualMapper
func NewCostAndUsageActualMapper() *CostAndUsageActualMapper {
	return &CostAndUsageActualMapper{}
}

// Identifier returns a unique identifier for the mapper
func (m *CostAndUsageActualMapper) Identifier() string {
	return "cost_and_usage_actual_mapper"
}

// Map converts the input data to a CostAndUsageActual struct
func (m *CostAndUsageActualMapper) Map(_ context.Context, a any, opts ...mappers.MapOption[*CostAndUsageActual]) (*CostAndUsageActual, error) {
	var input []byte

	// apply opts
	for _, opt := range opts {
		opt(m)
	}

	// validate input type
	switch v := a.(type) {
	case []byte:
		input = v
	case string:
		input = []byte(v)
	default:
		slog.Error("CostAndUsageActualMapper.Map failed to map row due to invalid type", "expected", "[]byte or string", "got", v)
		return nil, error_types.NewRowErrorWithMessage("unable to map row, invalid type received")
	}

	// read CSV line
	reader := csv.NewReader(bytes.NewReader(input))
	record, err := reader.Read()
	if err != nil {
		slog.Error("CostAndUsageActualMapper.Map failed to read CSV line", "error", err)
		return nil, error_types.NewRowErrorWithMessage("failed to read log line")
	}

	// validate header/value count
	if len(record) != len(m.headers) {
		slog.Error("CostAndUsageActualMapper.Map failed to map row due to header/value count mismatch", "expected", len(m.headers), "got", len(record))
		return nil, error_types.NewRowErrorWithMessage("row field count does not match header count")
	}

	// create a new CostAndUsageActual object with initialized maps
	output := NewCostAndUsageActual()

	// map to struct (normalize headers)
	for i, value := range record {
		field := m.headers[i]

		switch field {
		case "invoiceid":
			output.InvoiceId = &value
		case "previousinvoiceid":
			output.PreviousInvoiceId = &value
		case "billingaccountid":
			output.BillingAccountId = &value
		case "billingaccountname":
			output.BillingAccountName = &value
		case "billingprofileid":
			output.BillingProfileId = &value
		case "billingprofilename":
			output.BillingProfileName = &value
		case "invoicesectionid":
			output.InvoiceSectionId = &value
		case "invoicesectionname":
			output.InvoiceSectionName = &value
		case "resellername":
			output.ResellerName = &value
		case "resellermpnid":
			output.ResellerMpnId = &value
		case "costcenter":
			output.CostCenter = &value
		case "billingperiodenddate":
			if t, err := parseAzureDate(value); err == nil {
				output.BillingPeriodEndDate = &t
			}
		case "billingperiodstartdate":
			if t, err := parseAzureDate(value); err == nil {
				output.BillingPeriodStartDate = &t
			}
		case "serviceperiodenddate":
			if t, err := parseAzureDate(value); err == nil {
				output.ServicePeriodEndDate = &t
			}
		case "serviceperiodstartdate":
			if t, err := parseAzureDate(value); err == nil {
				output.ServicePeriodStartDate = &t
			}
		case "date":
			if t, err := parseAzureDate(value); err == nil {
				output.Date = &t
				output.TpTimestamp = t
			}
		case "servicefamily":
			output.ServiceFamily = &value
		case "productorderid":
			output.ProductOrderId = &value
		case "productordername":
			output.ProductOrderName = &value
		case "consumedservice":
			output.ConsumedService = &value
		case "meterid":
			output.MeterId = &value
		case "metername":
			output.MeterName = &value
		case "metercategory":
			output.MeterCategory = &value
		case "metersubcategory":
			output.MeterSubCategory = &value
		case "meterregion":
			output.MeterRegion = &value
		case "productid":
			output.ProductId = &value
		case "productname":
			output.ProductName = &value
		case "subscriptionid":
			output.SubscriptionId = &value
		case "subscriptionname":
			output.SubscriptionName = &value
		case "publishertype":
			output.PublisherType = &value
		case "publisherid":
			output.PublisherId = &value
		case "publishername":
			output.PublisherName = &value
		case "resourcegroupname":
			output.ResourceGroupName = &value
		case "resourceid":
			output.ResourceId = &value
		case "resourcelocation":
			output.ResourceLocation = &value
		case "location":
			output.Location = &value
		case "effectiveprice":
			if f, err := strconv.ParseFloat(value, 64); err == nil {
				output.EffectivePrice = &f
			}
		case "quantity":
			if f, err := strconv.ParseFloat(value, 64); err == nil {
				output.Quantity = &f
			}
		case "unitofmeasure":
			output.UnitOfMeasure = &value
		case "chargetype":
			output.ChargeType = &value
		case "billingcurrency":
			output.BillingCurrency = &value
		case "pricingcurrency":
			output.PricingCurrency = &value
		case "costinbillingcurrency":
			if f, err := strconv.ParseFloat(value, 64); err == nil {
				output.CostInBillingCurrency = &f
			}
		case "costinpricingcurrency":
			if f, err := strconv.ParseFloat(value, 64); err == nil {
				output.CostInPricingCurrency = &f
			}
		case "costinusd":
			if f, err := strconv.ParseFloat(value, 64); err == nil {
				output.CostInUsd = &f
			}
		case "paygcostinbillingcurrency":
			if f, err := strconv.ParseFloat(value, 64); err == nil {
				output.PaygCostInBillingCurrency = &f
			}
		case "paygcostinusd":
			if f, err := strconv.ParseFloat(value, 64); err == nil {
				output.PaygCostInUsd = &f
			}
		case "exchangeratepricing":
			if f, err := strconv.ParseFloat(value, 64); err == nil {
				output.ExchangeRatePricingToBilling = &f
			}
		case "exchangeratedate":
			if t, err := parseAzureDate(value); err == nil {
				output.ExchangeRateDate = &t
			}
		case "isazurecrediteligible":
			if b, err := strconv.ParseBool(value); err == nil {
				output.IsAzureCreditEligible = &b
			}
		case "serviceinfo1":
			output.ServiceInfo1 = &value
		case "serviceinfo2":
			output.ServiceInfo2 = &value
		case "additionalinfo":
			output.AdditionalInfo = &value
		case "tags":
			// Parse tags JSON string into map
			if value != "" {
				tags := make(map[string]interface{})
				err := json.Unmarshal([]byte(value), &tags)
				if err == nil && len(tags) > 0 {
					output.Tags = &tags
				} else if err != nil {
					slog.Error("CostAndUsageActualMapper.Map failed to parse tags JSON", "error", err, "value", value)
				}
			}
		case "paygprice":
			if f, err := strconv.ParseFloat(value, 64); err == nil {
				output.PaygPrice = &f
			}
		case "frequency":
			output.Frequency = &value
		case "term":
			output.Term = &value
		case "reservationid":
			output.ReservationId = &value
		case "reservationname":
			output.ReservationName = &value
		case "pricingmodel":
			output.PricingModel = &value
		case "unitprice":
			if f, err := strconv.ParseFloat(value, 64); err == nil {
				output.UnitPrice = &f
			}
		case "costallocationrulename":
			output.CostAllocationRuleName = &value
		case "benefitid":
			output.BenefitId = &value
		case "benefitname":
			output.BenefitName = &value
		case "provider":
			output.Provider = &value
		}
	}

	return output, nil
}

// OnHeader implements the HeaderRowNotifier interface
func (m *CostAndUsageActualMapper) OnHeader(header []string) {
	newHeaders := make([]string, len(header))
	// set headers but normalize first
	for i, h := range header {
		// Convert to lowercase and remove special characters
		v := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(h, " ", ""), "-", ""), "_", ""))
		newHeaders[i] = v
	}
	m.headers = newHeaders
}

// parseAzureDate parses a date string from Azure cost and usage data (mm/dd/yyyy format)
func parseAzureDate(dateStr string) (time.Time, error) {
	return time.Parse("01/02/2006", dateStr)
}
