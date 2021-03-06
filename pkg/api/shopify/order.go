package shopify

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"shopify-manager/pkg/config"
	"shopify-manager/pkg/constants"
	"shopify-manager/pkg/infrastructure/http"

	"github.com/tealeg/xlsx"
)

type GetOrdersResponse struct {
	Orders []struct {
		ID                    int64       `json:"id"`
		Email                 string      `json:"email"`
		ClosedAt              interface{} `json:"closed_at"`
		CreatedAt             string      `json:"created_at"`
		UpdatedAt             string      `json:"updated_at"`
		Number                int         `json:"number"`
		Note                  interface{} `json:"note"`
		Token                 string      `json:"token"`
		Gateway               string      `json:"gateway"`
		Test                  bool        `json:"test"`
		TotalPrice            string      `json:"total_price"`
		SubtotalPrice         string      `json:"subtotal_price"`
		TotalWeight           int         `json:"total_weight"`
		TotalTax              string      `json:"total_tax"`
		TaxesIncluded         bool        `json:"taxes_included"`
		Currency              string      `json:"currency"`
		FinancialStatus       string      `json:"financial_status"`
		Confirmed             bool        `json:"confirmed"`
		TotalDiscounts        string      `json:"total_discounts"`
		TotalLineItemsPrice   string      `json:"total_line_items_price"`
		CartToken             string      `json:"cart_token"`
		BuyerAcceptsMarketing bool        `json:"buyer_accepts_marketing"`
		Name                  string      `json:"name"`
		ReferringSite         string      `json:"referring_site"`
		LandingSite           string      `json:"landing_site"`
		CancelledAt           interface{} `json:"cancelled_at"`
		CancelReason          interface{} `json:"cancel_reason"`
		TotalPriceUsd         string      `json:"total_price_usd"`
		CheckoutToken         string      `json:"checkout_token"`
		Reference             string      `json:"reference"`
		UserID                interface{} `json:"user_id"`
		LocationID            interface{} `json:"location_id"`
		SourceIdentifier      string      `json:"source_identifier"`
		SourceURL             interface{} `json:"source_url"`
		ProcessedAt           string      `json:"processed_at"`
		DeviceID              interface{} `json:"device_id"`
		Phone                 string      `json:"phone"`
		CustomerLocale        interface{} `json:"customer_locale"`
		AppID                 interface{} `json:"app_id"`
		BrowserIP             string      `json:"browser_ip"`
		LandingSiteRef        string      `json:"landing_site_ref"`
		OrderNumber           int         `json:"order_number"`
		DiscountApplications  []struct {
			Type             string `json:"type"`
			Value            string `json:"value"`
			ValueType        string `json:"value_type"`
			AllocationMethod string `json:"allocation_method"`
			TargetSelection  string `json:"target_selection"`
			TargetType       string `json:"target_type"`
			Code             string `json:"code"`
		} `json:"discount_applications"`
		DiscountCodes []struct {
			Code   string `json:"code"`
			Amount string `json:"amount"`
			Type   string `json:"type"`
		} `json:"discount_codes"`
		NoteAttributes []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"note_attributes"`
		PaymentGatewayNames []string    `json:"payment_gateway_names"`
		ProcessingMethod    string      `json:"processing_method"`
		CheckoutID          int         `json:"checkout_id"`
		SourceName          string      `json:"source_name"`
		FulfillmentStatus   interface{} `json:"fulfillment_status"`
		TaxLines            []struct {
			Price    string  `json:"price"`
			Rate     float64 `json:"rate"`
			Title    string  `json:"title"`
			PriceSet struct {
				ShopMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"shop_money"`
				PresentmentMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"presentment_money"`
			} `json:"price_set"`
		} `json:"tax_lines"`
		Tags                   string `json:"tags"`
		ContactEmail           string `json:"contact_email"`
		OrderStatusURL         string `json:"order_status_url"`
		PresentmentCurrency    string `json:"presentment_currency"`
		TotalLineItemsPriceSet struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"total_line_items_price_set"`
		TotalDiscountsSet struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"total_discounts_set"`
		TotalShippingPriceSet struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"total_shipping_price_set"`
		SubtotalPriceSet struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"subtotal_price_set"`
		TotalPriceSet struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"total_price_set"`
		TotalTaxSet struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"total_tax_set"`
		TotalTipReceived       string      `json:"total_tip_received"`
		OriginalTotalDutiesSet interface{} `json:"original_total_duties_set"`
		CurrentTotalDutiesSet  interface{} `json:"current_total_duties_set"`
		AdminGraphqlAPIID      string      `json:"admin_graphql_api_id"`
		ShippingLines          []struct {
			ID                            int         `json:"id"`
			Title                         string      `json:"title"`
			Price                         string      `json:"price"`
			Code                          string      `json:"code"`
			Source                        string      `json:"source"`
			Phone                         interface{} `json:"phone"`
			RequestedFulfillmentServiceID interface{} `json:"requested_fulfillment_service_id"`
			DeliveryCategory              interface{} `json:"delivery_category"`
			CarrierIdentifier             interface{} `json:"carrier_identifier"`
			DiscountedPrice               string      `json:"discounted_price"`
			PriceSet                      struct {
				ShopMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"shop_money"`
				PresentmentMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"presentment_money"`
			} `json:"price_set"`
			DiscountedPriceSet struct {
				ShopMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"shop_money"`
				PresentmentMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"presentment_money"`
			} `json:"discounted_price_set"`
			DiscountAllocations []interface{} `json:"discount_allocations"`
			TaxLines            []interface{} `json:"tax_lines"`
		} `json:"shipping_lines"`
		BillingAddress struct {
			FirstName    string      `json:"first_name"`
			Address1     string      `json:"address1"`
			Phone        string      `json:"phone"`
			City         string      `json:"city"`
			Zip          string      `json:"zip"`
			Province     string      `json:"province"`
			Country      string      `json:"country"`
			LastName     string      `json:"last_name"`
			Address2     string      `json:"address2"`
			Company      interface{} `json:"company"`
			Latitude     float64     `json:"latitude"`
			Longitude    float64     `json:"longitude"`
			Name         string      `json:"name"`
			CountryCode  string      `json:"country_code"`
			ProvinceCode string      `json:"province_code"`
		} `json:"billing_address"`
		ShippingAddress struct {
			FirstName    string      `json:"first_name"`
			Address1     string      `json:"address1"`
			Phone        string      `json:"phone"`
			City         string      `json:"city"`
			Zip          string      `json:"zip"`
			Province     string      `json:"province"`
			Country      string      `json:"country"`
			LastName     string      `json:"last_name"`
			Address2     string      `json:"address2"`
			Company      interface{} `json:"company"`
			Latitude     float64     `json:"latitude"`
			Longitude    float64     `json:"longitude"`
			Name         string      `json:"name"`
			CountryCode  string      `json:"country_code"`
			ProvinceCode string      `json:"province_code"`
		} `json:"shipping_address"`
		ClientDetails struct {
			BrowserIP      string      `json:"browser_ip"`
			AcceptLanguage interface{} `json:"accept_language"`
			UserAgent      interface{} `json:"user_agent"`
			SessionHash    interface{} `json:"session_hash"`
			BrowserWidth   interface{} `json:"browser_width"`
			BrowserHeight  interface{} `json:"browser_height"`
		} `json:"client_details"`
		PaymentDetails struct {
			CreditCardBin     interface{} `json:"credit_card_bin"`
			AvsResultCode     interface{} `json:"avs_result_code"`
			CvvResultCode     interface{} `json:"cvv_result_code"`
			CreditCardNumber  string      `json:"credit_card_number"`
			CreditCardCompany string      `json:"credit_card_company"`
		} `json:"payment_details"`
		Customer struct {
			ID                        int           `json:"id"`
			Email                     string        `json:"email"`
			AcceptsMarketing          bool          `json:"accepts_marketing"`
			CreatedAt                 string        `json:"created_at"`
			UpdatedAt                 string        `json:"updated_at"`
			FirstName                 string        `json:"first_name"`
			LastName                  string        `json:"last_name"`
			OrdersCount               int           `json:"orders_count"`
			State                     string        `json:"state"`
			TotalSpent                string        `json:"total_spent"`
			LastOrderID               int           `json:"last_order_id"`
			Note                      interface{}   `json:"note"`
			VerifiedEmail             bool          `json:"verified_email"`
			MultipassIdentifier       interface{}   `json:"multipass_identifier"`
			TaxExempt                 bool          `json:"tax_exempt"`
			Phone                     string        `json:"phone"`
			Tags                      string        `json:"tags"`
			LastOrderName             string        `json:"last_order_name"`
			Currency                  string        `json:"currency"`
			AcceptsMarketingUpdatedAt string        `json:"accepts_marketing_updated_at"`
			MarketingOptInLevel       interface{}   `json:"marketing_opt_in_level"`
			TaxExemptions             []interface{} `json:"tax_exemptions"`
			AdminGraphqlAPIID         string        `json:"admin_graphql_api_id"`
			DefaultAddress            struct {
				ID           int         `json:"id"`
				CustomerID   int         `json:"customer_id"`
				FirstName    interface{} `json:"first_name"`
				LastName     interface{} `json:"last_name"`
				Company      interface{} `json:"company"`
				Address1     string      `json:"address1"`
				Address2     string      `json:"address2"`
				City         string      `json:"city"`
				Province     string      `json:"province"`
				Country      string      `json:"country"`
				Zip          string      `json:"zip"`
				Phone        string      `json:"phone"`
				Name         string      `json:"name"`
				ProvinceCode string      `json:"province_code"`
				CountryCode  string      `json:"country_code"`
				CountryName  string      `json:"country_name"`
				Default      bool        `json:"default"`
			} `json:"default_address"`
		} `json:"customer"`
		LineItems []struct {
			ID                         int         `json:"id"`
			VariantID                  int         `json:"variant_id"`
			Title                      string      `json:"title"`
			Quantity                   int         `json:"quantity"`
			Sku                        string      `json:"sku"`
			VariantTitle               string      `json:"variant_title"`
			Vendor                     interface{} `json:"vendor"`
			FulfillmentService         string      `json:"fulfillment_service"`
			ProductID                  int         `json:"product_id"`
			RequiresShipping           bool        `json:"requires_shipping"`
			Taxable                    bool        `json:"taxable"`
			GiftCard                   bool        `json:"gift_card"`
			Name                       string      `json:"name"`
			VariantInventoryManagement string      `json:"variant_inventory_management"`
			Properties                 []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"properties"`
			ProductExists       bool        `json:"product_exists"`
			FulfillableQuantity int         `json:"fulfillable_quantity"`
			Grams               int         `json:"grams"`
			Price               string      `json:"price"`
			TotalDiscount       string      `json:"total_discount"`
			FulfillmentStatus   interface{} `json:"fulfillment_status"`
			PriceSet            struct {
				ShopMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"shop_money"`
				PresentmentMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"presentment_money"`
			} `json:"price_set"`
			TotalDiscountSet struct {
				ShopMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"shop_money"`
				PresentmentMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"presentment_money"`
			} `json:"total_discount_set"`
			DiscountAllocations []struct {
				Amount                   string `json:"amount"`
				DiscountApplicationIndex int    `json:"discount_application_index"`
				AmountSet                struct {
					ShopMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"shop_money"`
					PresentmentMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"presentment_money"`
				} `json:"amount_set"`
			} `json:"discount_allocations"`
			Duties            []interface{} `json:"duties"`
			AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
			TaxLines          []struct {
				Title    string  `json:"title"`
				Price    string  `json:"price"`
				Rate     float64 `json:"rate"`
				PriceSet struct {
					ShopMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"shop_money"`
					PresentmentMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"presentment_money"`
				} `json:"price_set"`
			} `json:"tax_lines"`
		} `json:"line_items"`
		Fulfillments []struct {
			ID              int         `json:"id"`
			OrderID         int         `json:"order_id"`
			Status          string      `json:"status"`
			CreatedAt       string      `json:"created_at"`
			Service         string      `json:"service"`
			UpdatedAt       string      `json:"updated_at"`
			TrackingCompany string      `json:"tracking_company"`
			ShipmentStatus  interface{} `json:"shipment_status"`
			LocationID      int         `json:"location_id"`
			TrackingNumber  string      `json:"tracking_number"`
			TrackingNumbers []string    `json:"tracking_numbers"`
			TrackingURL     string      `json:"tracking_url"`
			TrackingUrls    []string    `json:"tracking_urls"`
			Receipt         struct {
				Testcase      bool   `json:"testcase"`
				Authorization string `json:"authorization"`
			} `json:"receipt"`
			Name              string `json:"name"`
			AdminGraphqlAPIID string `json:"admin_graphql_api_id"`
			LineItems         []struct {
				ID                         int         `json:"id"`
				VariantID                  int         `json:"variant_id"`
				Title                      string      `json:"title"`
				Quantity                   int         `json:"quantity"`
				Sku                        string      `json:"sku"`
				VariantTitle               string      `json:"variant_title"`
				Vendor                     interface{} `json:"vendor"`
				FulfillmentService         string      `json:"fulfillment_service"`
				ProductID                  int         `json:"product_id"`
				RequiresShipping           bool        `json:"requires_shipping"`
				Taxable                    bool        `json:"taxable"`
				GiftCard                   bool        `json:"gift_card"`
				Name                       string      `json:"name"`
				VariantInventoryManagement string      `json:"variant_inventory_management"`
				Properties                 []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"properties"`
				ProductExists       bool        `json:"product_exists"`
				FulfillableQuantity int         `json:"fulfillable_quantity"`
				Grams               int         `json:"grams"`
				Price               string      `json:"price"`
				TotalDiscount       string      `json:"total_discount"`
				FulfillmentStatus   interface{} `json:"fulfillment_status"`
				PriceSet            struct {
					ShopMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"shop_money"`
					PresentmentMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"presentment_money"`
				} `json:"price_set"`
				TotalDiscountSet struct {
					ShopMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"shop_money"`
					PresentmentMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"presentment_money"`
				} `json:"total_discount_set"`
				DiscountAllocations []struct {
					Amount                   string `json:"amount"`
					DiscountApplicationIndex int    `json:"discount_application_index"`
					AmountSet                struct {
						ShopMoney struct {
							Amount       string `json:"amount"`
							CurrencyCode string `json:"currency_code"`
						} `json:"shop_money"`
						PresentmentMoney struct {
							Amount       string `json:"amount"`
							CurrencyCode string `json:"currency_code"`
						} `json:"presentment_money"`
					} `json:"amount_set"`
				} `json:"discount_allocations"`
				Duties            []interface{} `json:"duties"`
				AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
				TaxLines          []struct {
					Title    string  `json:"title"`
					Price    string  `json:"price"`
					Rate     float64 `json:"rate"`
					PriceSet struct {
						ShopMoney struct {
							Amount       string `json:"amount"`
							CurrencyCode string `json:"currency_code"`
						} `json:"shop_money"`
						PresentmentMoney struct {
							Amount       string `json:"amount"`
							CurrencyCode string `json:"currency_code"`
						} `json:"presentment_money"`
					} `json:"price_set"`
				} `json:"tax_lines"`
			} `json:"line_items"`
		} `json:"fulfillments"`
		Refunds []struct {
			ID                int           `json:"id"`
			OrderID           int           `json:"order_id"`
			CreatedAt         string        `json:"created_at"`
			Note              string        `json:"note"`
			UserID            int           `json:"user_id"`
			ProcessedAt       string        `json:"processed_at"`
			Restock           bool          `json:"restock"`
			Duties            []interface{} `json:"duties"`
			AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
			RefundLineItems   []struct {
				ID          int     `json:"id"`
				Quantity    int     `json:"quantity"`
				LineItemID  int     `json:"line_item_id"`
				LocationID  int     `json:"location_id"`
				RestockType string  `json:"restock_type"`
				Subtotal    float64 `json:"subtotal"`
				TotalTax    float64 `json:"total_tax"`
				SubtotalSet struct {
					ShopMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"shop_money"`
					PresentmentMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"presentment_money"`
				} `json:"subtotal_set"`
				TotalTaxSet struct {
					ShopMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"shop_money"`
					PresentmentMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"presentment_money"`
				} `json:"total_tax_set"`
				LineItem struct {
					ID                         int           `json:"id"`
					VariantID                  int           `json:"variant_id"`
					Title                      string        `json:"title"`
					Quantity                   int           `json:"quantity"`
					Sku                        string        `json:"sku"`
					VariantTitle               string        `json:"variant_title"`
					Vendor                     interface{}   `json:"vendor"`
					FulfillmentService         string        `json:"fulfillment_service"`
					ProductID                  int           `json:"product_id"`
					RequiresShipping           bool          `json:"requires_shipping"`
					Taxable                    bool          `json:"taxable"`
					GiftCard                   bool          `json:"gift_card"`
					Name                       string        `json:"name"`
					VariantInventoryManagement string        `json:"variant_inventory_management"`
					Properties                 []interface{} `json:"properties"`
					ProductExists              bool          `json:"product_exists"`
					FulfillableQuantity        int           `json:"fulfillable_quantity"`
					Grams                      int           `json:"grams"`
					Price                      string        `json:"price"`
					TotalDiscount              string        `json:"total_discount"`
					FulfillmentStatus          interface{}   `json:"fulfillment_status"`
					PriceSet                   struct {
						ShopMoney struct {
							Amount       string `json:"amount"`
							CurrencyCode string `json:"currency_code"`
						} `json:"shop_money"`
						PresentmentMoney struct {
							Amount       string `json:"amount"`
							CurrencyCode string `json:"currency_code"`
						} `json:"presentment_money"`
					} `json:"price_set"`
					TotalDiscountSet struct {
						ShopMoney struct {
							Amount       string `json:"amount"`
							CurrencyCode string `json:"currency_code"`
						} `json:"shop_money"`
						PresentmentMoney struct {
							Amount       string `json:"amount"`
							CurrencyCode string `json:"currency_code"`
						} `json:"presentment_money"`
					} `json:"total_discount_set"`
					DiscountAllocations []struct {
						Amount                   string `json:"amount"`
						DiscountApplicationIndex int    `json:"discount_application_index"`
						AmountSet                struct {
							ShopMoney struct {
								Amount       string `json:"amount"`
								CurrencyCode string `json:"currency_code"`
							} `json:"shop_money"`
							PresentmentMoney struct {
								Amount       string `json:"amount"`
								CurrencyCode string `json:"currency_code"`
							} `json:"presentment_money"`
						} `json:"amount_set"`
					} `json:"discount_allocations"`
					Duties            []interface{} `json:"duties"`
					AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
					TaxLines          []struct {
						Title    string  `json:"title"`
						Price    string  `json:"price"`
						Rate     float64 `json:"rate"`
						PriceSet struct {
							ShopMoney struct {
								Amount       string `json:"amount"`
								CurrencyCode string `json:"currency_code"`
							} `json:"shop_money"`
							PresentmentMoney struct {
								Amount       string `json:"amount"`
								CurrencyCode string `json:"currency_code"`
							} `json:"presentment_money"`
						} `json:"price_set"`
					} `json:"tax_lines"`
				} `json:"line_item"`
			} `json:"refund_line_items"`
			Transactions []struct {
				ID            int         `json:"id"`
				OrderID       int         `json:"order_id"`
				Kind          string      `json:"kind"`
				Gateway       string      `json:"gateway"`
				Status        string      `json:"status"`
				Message       interface{} `json:"message"`
				CreatedAt     string      `json:"created_at"`
				Test          bool        `json:"test"`
				Authorization string      `json:"authorization"`
				LocationID    interface{} `json:"location_id"`
				UserID        interface{} `json:"user_id"`
				ParentID      int         `json:"parent_id"`
				ProcessedAt   string      `json:"processed_at"`
				DeviceID      interface{} `json:"device_id"`
				Receipt       struct {
				} `json:"receipt"`
				ErrorCode                  interface{} `json:"error_code"`
				SourceName                 string      `json:"source_name"`
				CurrencyExchangeAdjustment interface{} `json:"currency_exchange_adjustment"`
				Amount                     string      `json:"amount"`
				Currency                   string      `json:"currency"`
				AdminGraphqlAPIID          string      `json:"admin_graphql_api_id"`
			} `json:"transactions"`
			OrderAdjustments []interface{} `json:"order_adjustments"`
		} `json:"refunds"`
	} `json:"orders"`
}

type CancelOrderRequest struct {
	Email bool `json:"email"`
}

type CancelOrderResponse struct {
	Order struct {
		ID                    int         `json:"id"`
		Email                 string      `json:"email"`
		ClosedAt              interface{} `json:"closed_at"`
		CreatedAt             string      `json:"created_at"`
		UpdatedAt             string      `json:"updated_at"`
		Number                int         `json:"number"`
		Note                  interface{} `json:"note"`
		Token                 string      `json:"token"`
		Gateway               string      `json:"gateway"`
		Test                  bool        `json:"test"`
		TotalPrice            string      `json:"total_price"`
		SubtotalPrice         string      `json:"subtotal_price"`
		TotalWeight           int         `json:"total_weight"`
		TotalTax              string      `json:"total_tax"`
		TaxesIncluded         bool        `json:"taxes_included"`
		Currency              string      `json:"currency"`
		FinancialStatus       string      `json:"financial_status"`
		Confirmed             bool        `json:"confirmed"`
		TotalDiscounts        string      `json:"total_discounts"`
		TotalLineItemsPrice   string      `json:"total_line_items_price"`
		CartToken             string      `json:"cart_token"`
		BuyerAcceptsMarketing bool        `json:"buyer_accepts_marketing"`
		Name                  string      `json:"name"`
		ReferringSite         string      `json:"referring_site"`
		LandingSite           string      `json:"landing_site"`
		CancelledAt           interface{} `json:"cancelled_at"`
		CancelReason          interface{} `json:"cancel_reason"`
		TotalPriceUsd         string      `json:"total_price_usd"`
		CheckoutToken         string      `json:"checkout_token"`
		Reference             string      `json:"reference"`
		UserID                interface{} `json:"user_id"`
		LocationID            interface{} `json:"location_id"`
		SourceIdentifier      string      `json:"source_identifier"`
		SourceURL             interface{} `json:"source_url"`
		ProcessedAt           string      `json:"processed_at"`
		DeviceID              interface{} `json:"device_id"`
		Phone                 string      `json:"phone"`
		CustomerLocale        interface{} `json:"customer_locale"`
		AppID                 interface{} `json:"app_id"`
		BrowserIP             string      `json:"browser_ip"`
		LandingSiteRef        string      `json:"landing_site_ref"`
		OrderNumber           int         `json:"order_number"`
		DiscountApplications  []struct {
			Type             string `json:"type"`
			Value            string `json:"value"`
			ValueType        string `json:"value_type"`
			AllocationMethod string `json:"allocation_method"`
			TargetSelection  string `json:"target_selection"`
			TargetType       string `json:"target_type"`
			Code             string `json:"code"`
		} `json:"discount_applications"`
		DiscountCodes []struct {
			Code   string `json:"code"`
			Amount string `json:"amount"`
			Type   string `json:"type"`
		} `json:"discount_codes"`
		NoteAttributes []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"note_attributes"`
		PaymentGatewayNames []string    `json:"payment_gateway_names"`
		ProcessingMethod    string      `json:"processing_method"`
		CheckoutID          int         `json:"checkout_id"`
		SourceName          string      `json:"source_name"`
		FulfillmentStatus   interface{} `json:"fulfillment_status"`
		TaxLines            []struct {
			Price    string  `json:"price"`
			Rate     float64 `json:"rate"`
			Title    string  `json:"title"`
			PriceSet struct {
				ShopMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"shop_money"`
				PresentmentMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"presentment_money"`
			} `json:"price_set"`
		} `json:"tax_lines"`
		Tags                   string `json:"tags"`
		ContactEmail           string `json:"contact_email"`
		OrderStatusURL         string `json:"order_status_url"`
		PresentmentCurrency    string `json:"presentment_currency"`
		TotalLineItemsPriceSet struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"total_line_items_price_set"`
		TotalDiscountsSet struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"total_discounts_set"`
		TotalShippingPriceSet struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"total_shipping_price_set"`
		SubtotalPriceSet struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"subtotal_price_set"`
		TotalPriceSet struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"total_price_set"`
		TotalTaxSet struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"total_tax_set"`
		TotalTipReceived       string      `json:"total_tip_received"`
		OriginalTotalDutiesSet interface{} `json:"original_total_duties_set"`
		CurrentTotalDutiesSet  interface{} `json:"current_total_duties_set"`
		AdminGraphqlAPIID      string      `json:"admin_graphql_api_id"`
		ShippingLines          []struct {
			ID                            int         `json:"id"`
			Title                         string      `json:"title"`
			Price                         string      `json:"price"`
			Code                          string      `json:"code"`
			Source                        string      `json:"source"`
			Phone                         interface{} `json:"phone"`
			RequestedFulfillmentServiceID interface{} `json:"requested_fulfillment_service_id"`
			DeliveryCategory              interface{} `json:"delivery_category"`
			CarrierIdentifier             interface{} `json:"carrier_identifier"`
			DiscountedPrice               string      `json:"discounted_price"`
			PriceSet                      struct {
				ShopMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"shop_money"`
				PresentmentMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"presentment_money"`
			} `json:"price_set"`
			DiscountedPriceSet struct {
				ShopMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"shop_money"`
				PresentmentMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"presentment_money"`
			} `json:"discounted_price_set"`
			DiscountAllocations []interface{} `json:"discount_allocations"`
			TaxLines            []interface{} `json:"tax_lines"`
		} `json:"shipping_lines"`
		BillingAddress struct {
			FirstName    string      `json:"first_name"`
			Address1     string      `json:"address1"`
			Phone        string      `json:"phone"`
			City         string      `json:"city"`
			Zip          string      `json:"zip"`
			Province     string      `json:"province"`
			Country      string      `json:"country"`
			LastName     string      `json:"last_name"`
			Address2     string      `json:"address2"`
			Company      interface{} `json:"company"`
			Latitude     float64     `json:"latitude"`
			Longitude    float64     `json:"longitude"`
			Name         string      `json:"name"`
			CountryCode  string      `json:"country_code"`
			ProvinceCode string      `json:"province_code"`
		} `json:"billing_address"`
		ShippingAddress struct {
			FirstName    string      `json:"first_name"`
			Address1     string      `json:"address1"`
			Phone        string      `json:"phone"`
			City         string      `json:"city"`
			Zip          string      `json:"zip"`
			Province     string      `json:"province"`
			Country      string      `json:"country"`
			LastName     string      `json:"last_name"`
			Address2     string      `json:"address2"`
			Company      interface{} `json:"company"`
			Latitude     float64     `json:"latitude"`
			Longitude    float64     `json:"longitude"`
			Name         string      `json:"name"`
			CountryCode  string      `json:"country_code"`
			ProvinceCode string      `json:"province_code"`
		} `json:"shipping_address"`
		ClientDetails struct {
			BrowserIP      string      `json:"browser_ip"`
			AcceptLanguage interface{} `json:"accept_language"`
			UserAgent      interface{} `json:"user_agent"`
			SessionHash    interface{} `json:"session_hash"`
			BrowserWidth   interface{} `json:"browser_width"`
			BrowserHeight  interface{} `json:"browser_height"`
		} `json:"client_details"`
		PaymentDetails struct {
			CreditCardBin     interface{} `json:"credit_card_bin"`
			AvsResultCode     interface{} `json:"avs_result_code"`
			CvvResultCode     interface{} `json:"cvv_result_code"`
			CreditCardNumber  string      `json:"credit_card_number"`
			CreditCardCompany string      `json:"credit_card_company"`
		} `json:"payment_details"`
		Customer struct {
			ID                        int           `json:"id"`
			Email                     string        `json:"email"`
			AcceptsMarketing          bool          `json:"accepts_marketing"`
			CreatedAt                 string        `json:"created_at"`
			UpdatedAt                 string        `json:"updated_at"`
			FirstName                 string        `json:"first_name"`
			LastName                  string        `json:"last_name"`
			OrdersCount               int           `json:"orders_count"`
			State                     string        `json:"state"`
			TotalSpent                string        `json:"total_spent"`
			LastOrderID               int           `json:"last_order_id"`
			Note                      interface{}   `json:"note"`
			VerifiedEmail             bool          `json:"verified_email"`
			MultipassIdentifier       interface{}   `json:"multipass_identifier"`
			TaxExempt                 bool          `json:"tax_exempt"`
			Phone                     interface{}   `json:"phone"`
			Tags                      string        `json:"tags"`
			LastOrderName             string        `json:"last_order_name"`
			Currency                  string        `json:"currency"`
			AcceptsMarketingUpdatedAt string        `json:"accepts_marketing_updated_at"`
			MarketingOptInLevel       interface{}   `json:"marketing_opt_in_level"`
			TaxExemptions             []interface{} `json:"tax_exemptions"`
			AdminGraphqlAPIID         string        `json:"admin_graphql_api_id"`
			DefaultAddress            struct {
				ID           int         `json:"id"`
				CustomerID   int         `json:"customer_id"`
				FirstName    interface{} `json:"first_name"`
				LastName     interface{} `json:"last_name"`
				Company      interface{} `json:"company"`
				Address1     string      `json:"address1"`
				Address2     string      `json:"address2"`
				City         string      `json:"city"`
				Province     string      `json:"province"`
				Country      string      `json:"country"`
				Zip          string      `json:"zip"`
				Phone        string      `json:"phone"`
				Name         string      `json:"name"`
				ProvinceCode string      `json:"province_code"`
				CountryCode  string      `json:"country_code"`
				CountryName  string      `json:"country_name"`
				Default      bool        `json:"default"`
			} `json:"default_address"`
		} `json:"customer"`
		LineItems []struct {
			ID                         int         `json:"id"`
			VariantID                  int         `json:"variant_id"`
			Title                      string      `json:"title"`
			Quantity                   int         `json:"quantity"`
			Sku                        string      `json:"sku"`
			VariantTitle               string      `json:"variant_title"`
			Vendor                     interface{} `json:"vendor"`
			FulfillmentService         string      `json:"fulfillment_service"`
			ProductID                  int         `json:"product_id"`
			RequiresShipping           bool        `json:"requires_shipping"`
			Taxable                    bool        `json:"taxable"`
			GiftCard                   bool        `json:"gift_card"`
			Name                       string      `json:"name"`
			VariantInventoryManagement string      `json:"variant_inventory_management"`
			Properties                 []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"properties"`
			ProductExists       bool        `json:"product_exists"`
			FulfillableQuantity int         `json:"fulfillable_quantity"`
			Grams               int         `json:"grams"`
			Price               string      `json:"price"`
			TotalDiscount       string      `json:"total_discount"`
			FulfillmentStatus   interface{} `json:"fulfillment_status"`
			PriceSet            struct {
				ShopMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"shop_money"`
				PresentmentMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"presentment_money"`
			} `json:"price_set"`
			TotalDiscountSet struct {
				ShopMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"shop_money"`
				PresentmentMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"presentment_money"`
			} `json:"total_discount_set"`
			DiscountAllocations []struct {
				Amount                   string `json:"amount"`
				DiscountApplicationIndex int    `json:"discount_application_index"`
				AmountSet                struct {
					ShopMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"shop_money"`
					PresentmentMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"presentment_money"`
				} `json:"amount_set"`
			} `json:"discount_allocations"`
			Duties            []interface{} `json:"duties"`
			AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
			TaxLines          []struct {
				Title    string  `json:"title"`
				Price    string  `json:"price"`
				Rate     float64 `json:"rate"`
				PriceSet struct {
					ShopMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"shop_money"`
					PresentmentMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"presentment_money"`
				} `json:"price_set"`
			} `json:"tax_lines"`
		} `json:"line_items"`
		Fulfillments []struct {
			ID              int         `json:"id"`
			OrderID         int         `json:"order_id"`
			Status          string      `json:"status"`
			CreatedAt       string      `json:"created_at"`
			Service         string      `json:"service"`
			UpdatedAt       string      `json:"updated_at"`
			TrackingCompany string      `json:"tracking_company"`
			ShipmentStatus  interface{} `json:"shipment_status"`
			LocationID      int         `json:"location_id"`
			TrackingNumber  string      `json:"tracking_number"`
			TrackingNumbers []string    `json:"tracking_numbers"`
			TrackingURL     string      `json:"tracking_url"`
			TrackingUrls    []string    `json:"tracking_urls"`
			Receipt         struct {
				Testcase      bool   `json:"testcase"`
				Authorization string `json:"authorization"`
			} `json:"receipt"`
			Name              string `json:"name"`
			AdminGraphqlAPIID string `json:"admin_graphql_api_id"`
			LineItems         []struct {
				ID                         int         `json:"id"`
				VariantID                  int         `json:"variant_id"`
				Title                      string      `json:"title"`
				Quantity                   int         `json:"quantity"`
				Sku                        string      `json:"sku"`
				VariantTitle               string      `json:"variant_title"`
				Vendor                     interface{} `json:"vendor"`
				FulfillmentService         string      `json:"fulfillment_service"`
				ProductID                  int         `json:"product_id"`
				RequiresShipping           bool        `json:"requires_shipping"`
				Taxable                    bool        `json:"taxable"`
				GiftCard                   bool        `json:"gift_card"`
				Name                       string      `json:"name"`
				VariantInventoryManagement string      `json:"variant_inventory_management"`
				Properties                 []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"properties"`
				ProductExists       bool        `json:"product_exists"`
				FulfillableQuantity int         `json:"fulfillable_quantity"`
				Grams               int         `json:"grams"`
				Price               string      `json:"price"`
				TotalDiscount       string      `json:"total_discount"`
				FulfillmentStatus   interface{} `json:"fulfillment_status"`
				PriceSet            struct {
					ShopMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"shop_money"`
					PresentmentMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"presentment_money"`
				} `json:"price_set"`
				TotalDiscountSet struct {
					ShopMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"shop_money"`
					PresentmentMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"presentment_money"`
				} `json:"total_discount_set"`
				DiscountAllocations []struct {
					Amount                   string `json:"amount"`
					DiscountApplicationIndex int    `json:"discount_application_index"`
					AmountSet                struct {
						ShopMoney struct {
							Amount       string `json:"amount"`
							CurrencyCode string `json:"currency_code"`
						} `json:"shop_money"`
						PresentmentMoney struct {
							Amount       string `json:"amount"`
							CurrencyCode string `json:"currency_code"`
						} `json:"presentment_money"`
					} `json:"amount_set"`
				} `json:"discount_allocations"`
				Duties            []interface{} `json:"duties"`
				AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
				TaxLines          []struct {
					Title    string  `json:"title"`
					Price    string  `json:"price"`
					Rate     float64 `json:"rate"`
					PriceSet struct {
						ShopMoney struct {
							Amount       string `json:"amount"`
							CurrencyCode string `json:"currency_code"`
						} `json:"shop_money"`
						PresentmentMoney struct {
							Amount       string `json:"amount"`
							CurrencyCode string `json:"currency_code"`
						} `json:"presentment_money"`
					} `json:"price_set"`
				} `json:"tax_lines"`
			} `json:"line_items"`
		} `json:"fulfillments"`
		Refunds []struct {
			ID                int           `json:"id"`
			OrderID           int           `json:"order_id"`
			CreatedAt         string        `json:"created_at"`
			Note              string        `json:"note"`
			UserID            int           `json:"user_id"`
			ProcessedAt       string        `json:"processed_at"`
			Restock           bool          `json:"restock"`
			Duties            []interface{} `json:"duties"`
			AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
			RefundLineItems   []struct {
				ID          int     `json:"id"`
				Quantity    int     `json:"quantity"`
				LineItemID  int     `json:"line_item_id"`
				LocationID  int     `json:"location_id"`
				RestockType string  `json:"restock_type"`
				Subtotal    float64 `json:"subtotal"`
				TotalTax    float64 `json:"total_tax"`
				SubtotalSet struct {
					ShopMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"shop_money"`
					PresentmentMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"presentment_money"`
				} `json:"subtotal_set"`
				TotalTaxSet struct {
					ShopMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"shop_money"`
					PresentmentMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"presentment_money"`
				} `json:"total_tax_set"`
				LineItem struct {
					ID                         int           `json:"id"`
					VariantID                  int           `json:"variant_id"`
					Title                      string        `json:"title"`
					Quantity                   int           `json:"quantity"`
					Sku                        string        `json:"sku"`
					VariantTitle               string        `json:"variant_title"`
					Vendor                     interface{}   `json:"vendor"`
					FulfillmentService         string        `json:"fulfillment_service"`
					ProductID                  int           `json:"product_id"`
					RequiresShipping           bool          `json:"requires_shipping"`
					Taxable                    bool          `json:"taxable"`
					GiftCard                   bool          `json:"gift_card"`
					Name                       string        `json:"name"`
					VariantInventoryManagement string        `json:"variant_inventory_management"`
					Properties                 []interface{} `json:"properties"`
					ProductExists              bool          `json:"product_exists"`
					FulfillableQuantity        int           `json:"fulfillable_quantity"`
					Grams                      int           `json:"grams"`
					Price                      string        `json:"price"`
					TotalDiscount              string        `json:"total_discount"`
					FulfillmentStatus          interface{}   `json:"fulfillment_status"`
					PriceSet                   struct {
						ShopMoney struct {
							Amount       string `json:"amount"`
							CurrencyCode string `json:"currency_code"`
						} `json:"shop_money"`
						PresentmentMoney struct {
							Amount       string `json:"amount"`
							CurrencyCode string `json:"currency_code"`
						} `json:"presentment_money"`
					} `json:"price_set"`
					TotalDiscountSet struct {
						ShopMoney struct {
							Amount       string `json:"amount"`
							CurrencyCode string `json:"currency_code"`
						} `json:"shop_money"`
						PresentmentMoney struct {
							Amount       string `json:"amount"`
							CurrencyCode string `json:"currency_code"`
						} `json:"presentment_money"`
					} `json:"total_discount_set"`
					DiscountAllocations []struct {
						Amount                   string `json:"amount"`
						DiscountApplicationIndex int    `json:"discount_application_index"`
						AmountSet                struct {
							ShopMoney struct {
								Amount       string `json:"amount"`
								CurrencyCode string `json:"currency_code"`
							} `json:"shop_money"`
							PresentmentMoney struct {
								Amount       string `json:"amount"`
								CurrencyCode string `json:"currency_code"`
							} `json:"presentment_money"`
						} `json:"amount_set"`
					} `json:"discount_allocations"`
					Duties            []interface{} `json:"duties"`
					AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
					TaxLines          []struct {
						Title    string  `json:"title"`
						Price    string  `json:"price"`
						Rate     float64 `json:"rate"`
						PriceSet struct {
							ShopMoney struct {
								Amount       string `json:"amount"`
								CurrencyCode string `json:"currency_code"`
							} `json:"shop_money"`
							PresentmentMoney struct {
								Amount       string `json:"amount"`
								CurrencyCode string `json:"currency_code"`
							} `json:"presentment_money"`
						} `json:"price_set"`
					} `json:"tax_lines"`
				} `json:"line_item"`
			} `json:"refund_line_items"`
			Transactions []struct {
				ID            int         `json:"id"`
				OrderID       int         `json:"order_id"`
				Kind          string      `json:"kind"`
				Gateway       string      `json:"gateway"`
				Status        string      `json:"status"`
				Message       interface{} `json:"message"`
				CreatedAt     string      `json:"created_at"`
				Test          bool        `json:"test"`
				Authorization string      `json:"authorization"`
				LocationID    interface{} `json:"location_id"`
				UserID        interface{} `json:"user_id"`
				ParentID      int         `json:"parent_id"`
				ProcessedAt   string      `json:"processed_at"`
				DeviceID      interface{} `json:"device_id"`
				Receipt       struct {
				} `json:"receipt"`
				ErrorCode         interface{} `json:"error_code"`
				SourceName        string      `json:"source_name"`
				Amount            string      `json:"amount"`
				Currency          string      `json:"currency"`
				AdminGraphqlAPIID string      `json:"admin_graphql_api_id"`
			} `json:"transactions"`
			OrderAdjustments []interface{} `json:"order_adjustments"`
		} `json:"refunds"`
	} `json:"order"`
	Notice string `json:"notice"`
}

type CreateTransactionRequest struct {
	Transaction struct {
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
		Kind     string `json:"kind"`
		ParentID int64  `json:"parent_id"`
	} `json:"transaction"`
}

type CreateTransactionResponse struct {
	Transaction struct {
		ID            int         `json:"id"`
		OrderID       int         `json:"order_id"`
		Kind          string      `json:"kind"`
		Gateway       string      `json:"gateway"`
		Status        string      `json:"status"`
		Message       string      `json:"message"`
		CreatedAt     string      `json:"created_at"`
		Test          bool        `json:"test"`
		Authorization interface{} `json:"authorization"`
		LocationID    interface{} `json:"location_id"`
		UserID        interface{} `json:"user_id"`
		ParentID      int         `json:"parent_id"`
		ProcessedAt   string      `json:"processed_at"`
		DeviceID      interface{} `json:"device_id"`
		Receipt       struct {
		} `json:"receipt"`
		ErrorCode                  interface{} `json:"error_code"`
		SourceName                 string      `json:"source_name"`
		CurrencyExchangeAdjustment interface{} `json:"currency_exchange_adjustment"`
		Amount                     string      `json:"amount"`
		Currency                   string      `json:"currency"`
		AdminGraphqlAPIID          string      `json:"admin_graphql_api_id"`
	} `json:"transaction"`
}

type GetTransactionResponse struct {
	Transactions []struct {
		ID            int64       `json:"id"`
		OrderID       int         `json:"order_id"`
		Kind          string      `json:"kind"`
		Gateway       string      `json:"gateway"`
		Status        string      `json:"status"`
		Message       interface{} `json:"message"`
		CreatedAt     string      `json:"created_at"`
		Test          bool        `json:"test"`
		Authorization string      `json:"authorization"`
		LocationID    interface{} `json:"location_id"`
		UserID        interface{} `json:"user_id"`
		ParentID      int         `json:"parent_id"`
		ProcessedAt   string      `json:"processed_at"`
		DeviceID      interface{} `json:"device_id"`
		Receipt       struct {
			Testcase      bool   `json:"testcase"`
			Authorization string `json:"authorization"`
		} `json:"receipt"`
		ErrorCode                  interface{} `json:"error_code"`
		SourceName                 string      `json:"source_name"`
		CurrencyExchangeAdjustment interface{} `json:"currency_exchange_adjustment"`
		Amount                     string      `json:"amount"`
		Currency                   string      `json:"currency"`
		AdminGraphqlAPIID          string      `json:"admin_graphql_api_id"`
		PaymentDetails             struct {
			CreditCardBin     interface{} `json:"credit_card_bin"`
			AvsResultCode     interface{} `json:"avs_result_code"`
			CvvResultCode     interface{} `json:"cvv_result_code"`
			CreditCardNumber  string      `json:"credit_card_number"`
			CreditCardCompany string      `json:"credit_card_company"`
		} `json:"payment_details,omitempty"`
	} `json:"transactions"`
}

func CancelOrders(config *config.Config) error {
	isSuccess := true
	cancelOrderNumberList, err := getOrderNumberList(constants.INPUT_EXCEL_FILE_PATH)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	limitCh := make(chan struct{}, config.Thread.ThreadNum)
	for _, orderNumber := range cancelOrderNumberList {
		wg.Add(1)
		limitCh <- struct{}{}
		go func(orderNumber int) {
			defer wg.Done()

			log.Printf("INFO : Try to get order by orderNumber '%d'\n", orderNumber)
			order, err := getOrder(orderNumber, config)
			if err != nil {
				log.Printf("ERROR : orderNumber '%d' failed to cancel due to coludn't get order. %s\n", orderNumber, err.Error())
				isSuccess = false
				<-limitCh
				return
			}

			log.Printf("INFO : Try to get transactionId by orderId '%d' (orderNumber '%d')\n", order.Orders[0].ID, orderNumber)
			transactionId, err := getTransactionId(order.Orders[0].ID, config)
			if err != nil {
				log.Printf("ERROR : orderNumber '%d' failed to cancel due to couldn't get transactionId. %s\n", orderNumber, err.Error())
				isSuccess = false
				<-limitCh
				return
			}

			log.Printf("INFO : Try to disable authorization by orderId '%d' and transactionId '%d' (orderNumber '%d')\n", order.Orders[0].ID, transactionId, orderNumber)
			_, err = disabeAuthorization(order.Orders[0].ID, transactionId, config)
			if err != nil {
				log.Printf("ERROR : orderNumber '%d' failed to cancel due to coludn't be disable auhtorization. %s\n", orderNumber, err.Error())
				isSuccess = false
				<-limitCh
				return
			}

			log.Printf("INFO : Try to cancel order by orderId '%d' (orderNumber '%d')\n", order.Orders[0].ID, orderNumber)
			_, err = cancelOrder(order.Orders[0].ID, config)
			if err != nil {
				log.Printf("ERROR : orderNumber '%d' failed to cancel. %s\n", orderNumber, err.Error())
				isSuccess = false
				<-limitCh
				return
			}

			log.Printf("orderNumber '%d' successed to cancel.\n", orderNumber)
			<-limitCh

		}(orderNumber)
	}

	wg.Wait()

	if isSuccess {
		return nil
	} else {
		return fmt.Errorf("Failed to cancel any of orders.")
	}
}

func getOrder(orderNumber int, config *config.Config) (*GetOrdersResponse, error) {

	getOrderUrl := fmt.Sprintf(constants.GET_ORDER_URL_TEMPLATE, config.ApiInfo.ApiKey, config.ApiInfo.ApiPassword, orderNumber)
	httpReqHeader := map[string]string{}
	httpReqHeader["Content-Type"] = "application/json"
	jsonRes, err := http.Get(getOrderUrl, nil, httpReqHeader)
	if err != nil {
		return nil, err
	}

	orderResponse := new(GetOrdersResponse)
	err = json.Unmarshal(jsonRes, &orderResponse)
	if err != nil {
		log.Println("Cancel order response json unmarshal err")
		return nil, err
	}

	if len(orderResponse.Orders) < 1 {
		return nil, fmt.Errorf("Not found Order by order number '%d'", orderNumber)
	}

	return orderResponse, nil
}

func cancelOrder(cancelOrderId int64, config *config.Config) (*CancelOrderResponse, error) {
	cancelOrderReq := new(CancelOrderRequest)
	cancelOrderReq.Email = true
	reqJsonBytes, err := json.MarshalIndent(cancelOrderReq, "", "  ")
	if err != nil {
		log.Println("Cancel order request json marshal error")
		return nil, err
	}

	cancelOrderUrl := fmt.Sprintf(constants.CANCEL_ORDER_URL_TEMPLATE, config.ApiInfo.ApiKey, config.ApiInfo.ApiPassword, cancelOrderId)
	httpReqHeader := map[string]string{}
	httpReqHeader["Content-Type"] = "application/json"
	jsonRes, err := http.Post(cancelOrderUrl, reqJsonBytes, httpReqHeader)
	if err != nil {
		return nil, err
	}

	cancelOrderResponse := new(CancelOrderResponse)
	err = json.Unmarshal(jsonRes, &cancelOrderResponse)
	if err != nil {
		log.Println("Cancel order response json unmarshal err")
		return nil, err
	}

	return cancelOrderResponse, nil
}

func disabeAuthorization(orderId, transactionId int64, config *config.Config) (*CreateTransactionResponse, error) {

	createTransactionReq := new(CreateTransactionRequest)
	createTransactionReq.Transaction.Kind = "void"
	createTransactionReq.Transaction.Currency = "JPY"
	createTransactionReq.Transaction.ParentID = transactionId
	reqJsonBytes, err := json.MarshalIndent(createTransactionReq, "", "  ")
	if err != nil {
		log.Println("Create transaction request json marshal error")
		return nil, err
	}
	createTransactionUrl := fmt.Sprintf(constants.TRANSACTIONS_URL_TEMPLATE, config.ApiInfo.ApiKey, config.ApiInfo.ApiPassword, orderId)
	httpReqHeader := map[string]string{}
	httpReqHeader["Content-Type"] = "application/json"
	jsonRes, err := http.Post(createTransactionUrl, reqJsonBytes, httpReqHeader)
	if err != nil {
		return nil, err
	}

	createTransactionRes := new(CreateTransactionResponse)
	err = json.Unmarshal(jsonRes, &createTransactionRes)
	if err != nil {
		log.Println("Create transaction response json unmarshal err")
		return nil, err
	}

	if createTransactionRes.Transaction.Status != "success" {
		log.Printf("Create transaction response status is not success. actual %s\n", createTransactionRes.Transaction.Status)
		return nil, err
	}

	return createTransactionRes, nil
}

func getOrderNumberList(excelFilePath string) ([]int, error) {
	excel, err := xlsx.OpenFile(excelFilePath)
	if err != nil {
		log.Printf("%sのオープンに失敗", excelFilePath)
		return nil, err
	}

	var orderIdList []int
	sheet := excel.Sheets[0]
	for i, row := range sheet.Rows {
		if i == 0 {
			continue
		}

		name := row.Cells[0].String()
		if name == "" {
			continue
		}

		orderId, err := row.Cells[0].Int()
		if err != nil {
			log.Printf("%d行目、OrderIDが空、または数値でないためスキップ", i+1)
			return nil, err
		}

		orderIdList = append(orderIdList, orderId)
	}

	return orderIdList, nil
}

func getTransactionId(orderId int64, config *config.Config) (int64, error) {

	getTransactionUrl := fmt.Sprintf(constants.TRANSACTIONS_URL_TEMPLATE, config.ApiInfo.ApiKey, config.ApiInfo.ApiPassword, orderId)
	httpReqHeader := map[string]string{}
	httpReqHeader["Content-Type"] = "application/json"
	jsonRes, err := http.Get(getTransactionUrl, nil, httpReqHeader)
	if err != nil {
		return -1, err
	}

	getTransactionRes := new(GetTransactionResponse)
	err = json.Unmarshal(jsonRes, &getTransactionRes)
	if err != nil {
		log.Println("Get payment response json unmarshal err")
		return -1, err
	}

	if len(getTransactionRes.Transactions) < 1 {
		return -1, fmt.Errorf("Not found transaction by orderId '%d'", orderId)
	}

	for _, transaction := range getTransactionRes.Transactions {
		if transaction.Kind == "authorization" && transaction.Status == "success" {
			return transaction.ID, nil
		}
	}

	return -1, fmt.Errorf("Found transaction but no exists authorization type transaction")
}
