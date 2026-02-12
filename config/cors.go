package config

import "github.com/goravel/framework/facades"

func init() {
	config := facades.Config()
	appURL := config.Env("APP_URL", "http://localhost:3000").(string)
	config.Add("cors", map[string]any{
		"paths":                []string{},
		"allowed_methods":      []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		"allowed_origins":      []string{appURL},
		"allowed_headers":      []string{"Content-Type", "Accept", "X-Requested-With", "X-Inertia", "X-Inertia-Version", "X-Inertia-Partial-Data", "X-Inertia-Partial-Component", "X-XSRF-TOKEN"},
		"exposed_headers":      []string{"X-Inertia"},
		"max_age":              0,
		"supports_credentials": true,
	})
}
