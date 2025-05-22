package options

import "github.com/spf13/pflag"

var _ IOptions = (*SocialOptions)(nil)

type SocialOptions struct {
	TelegramBotToken string   `json:"telegram-bot-token" mapstructure:"telegram-bot-token"`
	XBearerToken     string   `json:"x-bearer-token" mapstructure:"x-bearer-token"`
	XClientID        string   `json:"x-client-id" mapstructure:"x-client-id"`
	XClientSecret    string   `json:"x-client-secret" mapstructure:"x-client-secret"`
	XRedirectURL     string   `json:"x-redirect-url" mapstructure:"x-redirect-url"`
	XOfficials       []string `json:"x-officials" mapstructure:"x-officials"`
}

func NewSocialOptions() *SocialOptions {
	return &SocialOptions{}
}

func (o *SocialOptions) Validate() []error {
	errs := []error{}

	return errs
}

func (o *SocialOptions) AddFlags(fs *pflag.FlagSet, prefixes ...string) {
	fs.StringVar(&o.TelegramBotToken, join(prefixes...)+"social.telegram-bot-token", o.TelegramBotToken, "")
	fs.StringVar(&o.XBearerToken, join(prefixes...)+"social.x-bearer-token", o.XBearerToken, "")
	fs.StringVar(&o.XClientID, join(prefixes...)+"social.x-client-id", o.XClientID, "")
	fs.StringVar(&o.XClientSecret, join(prefixes...)+"social.x-client-secret", o.XClientSecret, "")
	fs.StringVar(&o.XRedirectURL, join(prefixes...)+"social.x-redirect-url", o.XRedirectURL, "")
	fs.StringSliceVar(&o.XOfficials, join(prefixes...)+"social.x-officials", o.XOfficials, "")
}
