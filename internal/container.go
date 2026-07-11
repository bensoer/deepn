package internal

import (
	"deepn/internal/providers"
	"fmt"
)

type Container struct {
	allPRoviders []providers.Provider
}

func NewContainer(allProviders []providers.Provider) *Container {
	return &Container{
		allPRoviders: allProviders,
	}
}

func (c *Container) GetProvider() (providers.Provider, error) {
	// Do whatever parsing match necessary to return only the
	// provider desired by the user. This could be from a configuration
	// file or from a command line argument etc

	if len(c.allPRoviders) == 0 {
		return nil, fmt.Errorf("no providers available")
	}

	return c.allPRoviders[0], nil
}
