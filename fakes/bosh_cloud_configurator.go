package fakes

import "github.com/pivotal-cf-experimental/bosh-bootloader/aws/cloudformation"

type BoshCloudConfigurator struct {
	ConfigureCall struct {
		CallCount int
		Returns   struct {
			Error error
		}
	}
}

func (b *BoshCloudConfigurator) Configure(stack cloudformation.Stack) error {
	b.ConfigureCall.CallCount++

	return b.ConfigureCall.Returns.Error
}
