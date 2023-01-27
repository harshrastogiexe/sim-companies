package main

import (
	"github.com/harshrastogiexe/pkg/sim-companies-proxy/internal/core"
	"github.com/harshrastogiexe/pkg/sim-companies-proxy/internal/encyclopedia"
)

var (
	Encyclopedia core.Encyclopedia
)

func init() {
	Encyclopedia = encyclopedia.New()
}
