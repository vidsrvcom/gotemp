package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// HomeHandler handles requests for the home page.
type HomeHandler struct {
	templates *template.Template
}

// NewHomeHandler creates a new HomeHandler instance.
func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

// Index handles GET / - renders the home page.
func (h *HomeHandler) Index(w http.ResponseWriter, r *http.Request) {
	// Try to load templates if not already loaded
	if h.templates == nil {
		tmplPath := filepath.Join("templates", "*.html")
		tmpl, err := template.ParseGlob(tmplPath)
		if err != nil {
			// Fallback to simple HTML response if templates not found
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(defaultHomeHTML))
			return
		}
		h.templates = tmpl
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

// Health handles GET /health - returns health check status.
func (h *HomeHandler) Health(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{
		"status": "healthy",
	})
}

const defaultHomeHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go Web Template</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body { 
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        .container {
            text-align: center;
            color: white;
            padding: 2rem;
        }
        h1 { font-size: 3rem; margin-bottom: 1rem; }
        p { font-size: 1.2rem; opacity: 0.9; margin-bottom: 2rem; }
        .features {
            display: flex;
            gap: 2rem;
            flex-wrap: wrap;
            justify-content: center;
            margin-top: 2rem;
        }
        .feature {
            background: rgba(255,255,255,0.1);
            padding: 1.5rem;
            border-radius: 10px;
            max-width: 200px;
        }
        .feature h3 { margin-bottom: 0.5rem; }
        a { color: white; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🚀 Go Web Template</h1>
        <p>A production-ready Go web application template</p>
        <div class="features">
            <div class="feature">
                <h3>📦 Clean Architecture</h3>
                <p>Organized with handlers, services, and repositories</p>
            </div>
            <div class="feature">
                <h3>🔧 RESTful API</h3>
                <p>Well-structured API endpoints with JSON responses</p>
            </div>
            <div class="feature">
                <h3>✅ Testing</h3>
                <p>Comprehensive unit tests included</p>
            </div>
        </div>
        <p style="margin-top: 2rem;">
            <a href="/health">Health Check</a> | 
            <a href="/api/users">API Users</a>
        </p>
    </div>
</body>
</html>`
