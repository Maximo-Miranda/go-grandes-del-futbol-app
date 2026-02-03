package inertia

import (
	"encoding/json"
	"errors"
	"fmt"
	htmltemplate "html/template"
	"maps"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"grandesdelfutbol/app/models"
)

var (
	ErrMarshalPageData = errors.New("failed to marshal page data")
	ErrParseTemplate   = errors.New("failed to parse template")
	ErrExecuteTemplate = errors.New("failed to execute template")
)

var (
	instance *Inertia
	once     sync.Once
)

type Inertia struct {
	rootTemplate string
	version      string
	sharedProps  map[string]any
	mu           sync.RWMutex
}

type page struct {
	Component string         `json:"component"`
	Props     map[string]any `json:"props"`
	URL       string         `json:"url"`
	Version   string         `json:"version"`
}

type manifest map[string]struct {
	File string   `json:"file"`
	CSS  []string `json:"css"`
}

func New() *Inertia {
	once.Do(func() {
		instance = &Inertia{
			rootTemplate: "resources/views/app.gohtml",
			sharedProps:  make(map[string]any),
		}
	})
	return instance
}

func (i *Inertia) Share(key string, value any) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.sharedProps[key] = value
}

func (i *Inertia) mergeProps(props map[string]any) map[string]any {
	i.mu.RLock()
	merged := make(map[string]any, len(i.sharedProps)+len(props))
	maps.Copy(merged, i.sharedProps)
	i.mu.RUnlock()

	maps.Copy(merged, props)
	return merged
}

func (i *Inertia) Render(ctx http.Context, component string, props map[string]any) http.Response {
	merged := i.mergeProps(props)

	if user, ok := ctx.Value("user").(*models.User); ok && user != nil {
		merged["auth"] = map[string]any{
			"user": map[string]any{
				"id":    user.ID,
				"email": user.Email,
				"name":  user.Name,
				"role":  user.Role,
			},
		}
	}

	pageData := page{
		Component: component,
		Props:     merged,
		URL:       ctx.Request().Path(),
		Version:   i.version,
	}

	if ctx.Request().Header("X-Inertia") == "true" {
		return ctx.Response().Header("X-Inertia", "true").Json(200, pageData)
	}

	return i.renderHTML(ctx, pageData)
}

func (i *Inertia) renderHTML(ctx http.Context, pageData page) http.Response {
	pageJSON, err := json.Marshal(pageData)
	if err != nil {
		facades.Log().Error(fmt.Errorf("%w: %w", ErrMarshalPageData, err))
		return ctx.Response().String(500, "Internal Server Error")
	}

	isDev := facades.Config().GetString("app.env") == "local"

	data := map[string]any{
		"title":    "Grandes del Fútbol",
		"pageData": string(pageJSON),
		"isDev":    isDev,
	}

	if !isDev {
		data["jsFiles"], data["cssFiles"] = i.getManifestAssets()
	}

	tmpl, err := htmltemplate.ParseFiles(i.rootTemplate)
	if err != nil {
		facades.Log().Error(fmt.Errorf("%w: %w", ErrParseTemplate, err))
		return ctx.Response().String(500, "Internal Server Error")
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		facades.Log().Error(fmt.Errorf("%w: %w", ErrExecuteTemplate, err))
		return ctx.Response().String(500, "Internal Server Error")
	}

	return ctx.Response().Data(200, "text/html; charset=utf-8", []byte(buf.String()))
}

func (i *Inertia) getManifestAssets() ([]string, []string) {
	manifestPath := filepath.Join("public", "build", ".vite", "manifest.json")
	data, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, nil
	}

	var m manifest
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, nil
	}

	entry, ok := m["resources/js/app.ts"]
	if !ok {
		return nil, nil
	}

	jsFiles := []string{"/build/" + entry.File}
	cssFiles := make([]string, len(entry.CSS))
	for idx, css := range entry.CSS {
		cssFiles[idx] = "/build/" + css
	}

	return jsFiles, cssFiles
}
