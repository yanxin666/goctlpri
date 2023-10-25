package upgrade

import "github.com/yanxin666/goctlpri/internal/cobrax"

// Cmd describes an upgrade command.
var Cmd = cobrax.NewCommand("upgrade", cobrax.WithRunE(upgrade))
