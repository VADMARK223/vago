package router

import (
	"context"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"vago/internal/app"
	"vago/internal/infra/token"
	"vago/internal/transport/http/handler"
	"vago/internal/transport/http/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(goCtx context.Context, ctx *app.Context, tokenProvider *token.JWTProvider) *gin.Engine {
	deps := buildDeps(goCtx, ctx, tokenProvider)

	gin.SetMode(ctx.Cfg.GinMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.SetHTMLTemplate(loadTemplates(ctx, "web/templates"))

	_ = r.SetTrustedProxies(nil)

	r.Static("/static", "web/static")

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("web/static/favicon.ico")
	})

	registerReactApp(r)

	// Middleware
	r.Use(middleware.SessionMiddleware())
	r.Use(middleware.CheckJWT(tokenProvider, ctx.Cfg.RefreshTTLInt()))
	r.Use(middleware.LoadUserContext(deps.Services.User, deps.Cache))

	web := r.Group("")
	web.Use(middleware.NoCache, middleware.TemplateContext)
	registerWebRoutes(web, deps)

	api := r.Group("/api")
	registerAPIRoutes(api, deps)

	r.NoRoute(middleware.NoCache, middleware.TemplateContext, handler.NotFoundHandler)

	return r
}

func loadTemplates(ctx *app.Context, root string) *template.Template {
	tmpl := template.New("").Funcs(template.FuncMap{
		"dict": dict,
		"menuActive": func(path, href string) bool {
			if href == "/" {
				return path == "/"
			}
			return strings.HasPrefix(path, href)
		},
	})

	walkErr := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			ctx.Log.Fatal(err.Error())
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, ".html") || strings.HasSuffix(path, ".gohtml") {
			relPath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}

			relPath = filepath.ToSlash(relPath)

			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			_, err = tmpl.New(relPath).Parse(string(content))
			if err != nil {
				return err
			}
		}
		return nil
	})

	if walkErr != nil {
		ctx.Log.Fatal(walkErr.Error())
	}

	return tmpl
}

func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("invalid dict call: odd number of args")
	}
	m := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("dict keys must be strings")
		}
		m[key] = values[i+1]
	}
	return m, nil
}

func registerReactApp(r *gin.Engine) {
	distDir := "./web/v2/dist"

	r.GET("/v2/*path", func(c *gin.Context) {
		// Например: "/assets/index-xxxx.js" или "/courses" или "/"
		p := strings.TrimPrefix(c.Param("path"), "/")

		// Если запросили конкретный файл — отдадим его
		if p != "" {
			fp := filepath.Join(distDir, p)
			if st, err := os.Stat(fp); err == nil && !st.IsDir() { // TODO: пофиксить
				c.File(fp)
				return
			}
		}

		// Иначе — SPA fallback
		c.File(filepath.Join(distDir, "index.html"))
	})
}
