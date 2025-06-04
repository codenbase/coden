package options

import "github.com/spf13/pflag"

var _ IOptions = (*AuthenticationOptions)(nil)

type AuthenticationOptions struct {
	VerifyUrl string `json:"verify-url" mapstructure:"verify-url"`
}

func NewAuthenticationOptions() *AuthenticationOptions {
	return &AuthenticationOptions{
		VerifyUrl: "https://127.0.0.1/auth/verify",
	}
}

func (s *AuthenticationOptions) Validate() []error {
	var errs []error

	return errs
}

func (s *AuthenticationOptions) AddFlags(fs *pflag.FlagSet, prefixes ...string) {
	if fs == nil {
		return
	}

	fs.StringVar(&s.VerifyUrl, "authn.verify-url", s.VerifyUrl, "Verify URL used to verify jwt token.")
}
