package internal

import (
	"deepn/internal/providers"
	"deepn/internal/providers/antigravity"
	"deepn/internal/utils"
)

func Bootstrap() (*Container, error) {
	// Initialise all providers and wiring for them here

	subproc := utils.NewSubproc() // Create a new Subproc instance

	agcli, err := antigravity.NewAntigravityCLI(subproc)
	if err != nil {
		// Handle the error appropriately, e.g., log it or return it
		return nil, err // For now, just return the error
	}

	return NewContainer(
		[]providers.Provider{
			// agcli provider
			agcli,
		},
	), nil

}
