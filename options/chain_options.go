package options

import "github.com/spf13/pflag"

var _ IOptions = (*ChainOptions)(nil)

type ChainOptions struct {
	ChainId   int32  `json:"chain-id" mapstructure:"chain-id"`
	RPC       string `json:"rpc" mapstructure:"rpc"`
	SignerKey string `json:"signer-key" mapstructure:"signer-key"`
}

func NewChainOptions() *ChainOptions {
	return &ChainOptions{}
}

func (o *ChainOptions) Validate() []error {
	errs := []error{}

	return errs
}

func (o *ChainOptions) AddFlags(fs *pflag.FlagSet, prefixes ...string) {
	fs.Int32Var(&o.ChainId, join(prefixes...)+"chain.chain-id", o.ChainId, "")
	fs.StringVar(&o.RPC, join(prefixes...)+"chain.rpc", o.RPC, "")
	fs.StringVar(&o.SignerKey, join(prefixes...)+"chain.signer-key", o.SignerKey, "")
}
