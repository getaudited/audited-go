# Audited Go SDK

[![Go Coverage](https://github.com/getaudited/audited-go/wiki/coverage.svg)](https://raw.githack.com/wiki/getaudited/audited-go/coverage.html)
[![main](https://github.com/getaudited/audited-go/actions/workflows/main.yml/badge.svg)](https://github.com/getaudited/audited-go/actions/workflows/main.yml)
![GitHub tag](https://img.shields.io/github/tag/getaudited/audited-go.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/getaudited/audited-go/v1.svg)](https://pkg.go.dev/github.com/getaudited/audited-go/v1)

## Installation

To install the Go SDK, simply execute the following command on a terminal:

```sh 
go get github.com/getaudited/audited-go/v1
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"github.com/getaudited/audited"
	"github.com/getaudited/audited/v1"
)

func main() {
	client := audited.NewClient(audited.Config{
		BaseAPI: "https://audited.yourinstance.com",
		APIToken: "secure-api-token",
    })

	event := audited.Event{
		Action: "transfer.initiated",
		Actor: audited.Actor{
			Id: "7f1b2c3d-4e5f-6a7b-8c9d-0e1f2a3b4c5d",
			Metadata: &map[string]interface{}{
				"account_tier":       "premium",
				"compliance_status": "kyc_verified",
			},
			Name: new("Alex Rivera"),
			Type: "customer",
		},
		Context: audited.Context{
			Location:  "90.135.168.140",
			UserAgent: new("Mozilla/5.0 (Linux; Android 10; BLA-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.141 Mobile Safari/537.36"),
		},
		SourceID: "01H7XRA8WWGPQX996AXWG5KXYZ",
		Targets: []audited.Target{
			{
				ID: "acc_92a1b4c3d5e6",
				Metadata: &map[string]interface{}{
					"routing_number": "123456789",
					"bank_country":   "Germany",
				},
				Name: new("Elena Rostova"),
				Type: "external_bank_account",
			},
		},
		Version: 1,
		Metadata: &map[string]interface{}{
			"service_name":    "svc-payment-routing",
			"amount":          2500.00,
			"currency":        "EUR",
			"idempotency_key": "idem-9a8b7c6d5e4f",
			"risk_score":      0.08,
		},
		OccurredAt: time.Now(),
	}

	err := client.CreateEvent(context.Background(), event)
	if err != nil {
		fmt.Printf("Error creating event: %v\n", err)
		return
	}
	
	fmt.Println("Event created successfully")
}

```

## License

MIT