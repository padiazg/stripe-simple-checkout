package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
)

func (a *App) handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func (a *App) handleCreateCheckoutSession(w http.ResponseWriter, r *http.Request) {
	log.Printf("Create checkout session")

	stripe.Key = a.Config.StripeKey

	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{{
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency: stripe.String(string(stripe.CurrencyUSD)),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name: stripe.String("T-shirt"),
				},
				UnitAmount: stripe.Int64(2000),
			},
			Quantity: stripe.Int64(1),
		},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(a.Domain + "/success.html"),
		CancelURL:  stripe.String(a.Domain + "/cancel.html"),
	}

	s, err := session.New(params)
	if err != nil {
		log.Printf("session.New: %v", err)
		fmt.Fprintf(w, `{"error": "%s"}`, err.Error())
		return
	}

	log.Printf("Session created: %s => %s", s.ID, s.URL)
	fmt.Fprintf(w, `{"url": "%s"}`, s.URL)
}
