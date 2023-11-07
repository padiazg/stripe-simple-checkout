package main

import (
	a "github.com/padiazg/stripe-simple-checkout/app"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	var (
		stripeKey = viper.GetString("STRIPE_KEY")
		domain    = viper.GetString("DOMAIN")
		port      = viper.GetInt("PORT")
	)

	// Create a new instance of the application.
	app := a.New(&a.Config{
		HostPort:  uint(port),
		StripeKey: stripeKey,
		Domain:    domain,
	})

	// app.PrintRoutes()

	// Run the application.
	app.Run()

}
