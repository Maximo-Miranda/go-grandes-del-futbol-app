package config

import (
	"github.com/goravel/framework/contracts/filesystem"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/path"
	miniofacades "github.com/goravel/minio/facades"
)

func init() {
	config := facades.Config()
	config.Add("filesystems", map[string]any{
		"default": config.Env("FILESYSTEM_DISK", "local"),
		"disks": map[string]any{
			"local": map[string]any{
				"driver": "local",
				"root":   path.Storage("app"),
			},
			"public": map[string]any{
				"driver": "local",
				"root":   path.Storage("app/public"),
				"url":    config.Env("APP_URL", "").(string) + "/storage",
			},
			"minio": map[string]any{
				"driver":   "custom",
				"key":      config.Env("MINIO_ACCESS_KEY_ID"),
				"secret":   config.Env("MINIO_ACCESS_KEY_SECRET"),
				"region":   config.Env("MINIO_REGION"),
				"bucket":   config.Env("MINIO_BUCKET"),
				"url":      config.Env("MINIO_URL"),
				"endpoint": config.Env("MINIO_ENDPOINT"),
				"ssl":      config.Env("MINIO_SSL", false),
				"via": func() (filesystem.Driver, error) {
					return miniofacades.Minio("minio")
				},
			},
		},
	})
}
