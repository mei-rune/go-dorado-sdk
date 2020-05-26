// +build ignore

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/whywaita/go-dorado-sdk/example/lib"
)

func main() {
	ctx := context.Background()

	client, err := lib.GetClient()
	if err != nil {
		log.Fatal(err)
	}

	snapshots, err := client.LocalDevice.GetSnapshots(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range snapshots {
		fmt.Printf("%+v\n", s)
	}

	fmt.Println("operation is done!")
}
