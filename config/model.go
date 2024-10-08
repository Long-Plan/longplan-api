package config

type LocalConfig struct {
	Application ApplicationConfig `mapstructure:"app"`
	Database    DatabaseConfig    `mapstructure:"database"`
	CmuOauth    CmuOauthConfig    `mapstructure:"cmu_oauth"`
}

type ApplicationConfig struct {
	Name         string `mapstructure:"name"`
	ClientOrigin string `mapstructure:"client_origin"`
	Port         string `mapstructure:"port"`
	Secret       string `mapstructure:"secret"`
}

type DatabaseConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Name string `mapstructure:"name"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
}

type CmuOauthConfig struct {
	CmuOauthInfo             string `mapstructure:"cmu_oauth_info"`
	CmuOauthToken            string `mapstructure:"cmu_oauth_token"`
	CmuOauthRedirectURL      string `mapstructure:"cmu_oauth_redirect_url"`
	CmuOauthRedirectURLLocal string `mapstructure:"cmu_oauth_redirect_url_local"`
	CmuOauthClientID         string `mapstructure:"cmu_oauth_client_id"`
	CmuOauthClientSecret     string `mapstructure:"cmu_oauth_client_secret"`
}
