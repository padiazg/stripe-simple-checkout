# Stripe Simple Checkout

Example for Stripe simple checkout

## Pre-requisits
You need a Secret Key from your Stripe account: Developers -> API keys -> Secret key

## Config file
Create an .env file with this content
```ini
STRIPE_KEY=sk_test_***
PORT=9000
DOMAIN=http://localhost:9000/static
```

## Run
Run the app with
```bash
$ go run main.go
```

Point your browser to `http://localhost:9000/static`, then click on the **Checkout** button.


## Aknowledgements
This example is based on the [How To Accept Payments With Stripe](https://www.youtube.com/watch?v=1r-F3FIONl8&t=1056s) example from https://www.youtube.com/@WebDevSimplified.
Repo from that video is [here](https://github.com/WebDevSimplified/stripe-checkout-simple)