package bug

import (
	"github.com/spf13/cobra"
	"github.com/yanxin666/goctlpri/internal/cobrax"
)

// Cmd describes a bug command.
var Cmd = cobrax.NewCommand("bug", cobrax.WithRunE(cobra.NoArgs), cobrax.WithArgs(cobra.NoArgs))
