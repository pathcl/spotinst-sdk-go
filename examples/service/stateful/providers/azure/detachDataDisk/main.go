package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"log"
)

func main() {
	// All clients require a Session. The Session provides the client with
	// shared configuration such as account and credentials.
	// A Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	sess := session.New()

	// Create a new instance of the service's client with a Session.
	// Optional spotinst.Config values can also be provided as variadic
	// arguments to the New function. This option allows you to provide
	// service specific configuration.
	svc := azure.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Read stateful node configuration.
	_, err := svc.DetachDataDisk(ctx, &azure.DetachStatefulNodeDataDiskInput{
		ID:                        spotinst.String("ssn-01234567"),
		DataDiskName:              spotinst.String("foo"),
		DataDiskResourceGroupName: spotinst.String("foo"),
		ShouldDeallocate:          spotinst.Bool(true),
	})
	if err != nil {
		log.Fatalf("spotinst: failed to detach stateful node data disk: %v", err)
	}
}
