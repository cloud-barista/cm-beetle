package docs

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ScalarConfig holds configuration for the API Explorer (Scalar UI)
type ScalarConfig struct {
	Title           string // Default: "API Explorer"
	SpecURL         string // Default: "/swagger.json"
	AuthEnabled     bool   // Default: false
	AuthDescription string // Default: "" (Hidden if empty or AuthEnabled is false)
	SidebarWidth    string // Default: "450px"
}

// ScalarHandler returns an Echo handler that serves the API Explorer page
func ScalarHandler(config ScalarConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		title := config.Title
		if title == "" {
			title = "API Explorer"
		}

		specURL := config.SpecURL
		if specURL == "" {
			specURL = "/swagger.json"
		}

		sidebarWidth := config.SidebarWidth
		if sidebarWidth == "" {
			sidebarWidth = "450px"
		}

		authMessageElement := ""
		if config.AuthEnabled && config.AuthDescription != "" {
			authMessageElement = fmt.Sprintf(`<div class="auth-message">%s</div>`, config.AuthDescription)
		}

		htmlContent := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
	<title>%s</title>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<style>
		body { margin: 0; }
		.auth-message {
			padding: 10px;
			background-color: #f8f9fa;
			color: #333;
			text-align: center;
			font-family: system-ui, -apple-system, sans-serif;
			border-bottom: 1px solid #e9ecef;
			font-size: 14px;
		}
		/* Resizer handle */
		.sidebar-resizer {
			width: 10px;
			cursor: col-resize;
			position: absolute;
			top: 0;
			right: -5px;
			bottom: 0;
			z-index: 100;
			background-color: transparent;
		}
		/* Visual feedback on hover/active */
		.sidebar-resizer:hover, .sidebar-resizer.resizing {
			background-color: var(--scalar-color-accent, #007bff);
			opacity: 0.5;
		}
		:root {
			--scalar-sidebar-width: %s;
		}
		/* Ensure the sidebar container allows absolute positioning of the resizer */
		.scalar-app aside {
			width: var(--scalar-sidebar-width, %s) !important;
			min-width: var(--scalar-sidebar-width, %s) !important;
			transition: none !important;
		}
		/* Fix for API Group header wrapping */
		@media (min-width: 1200px) {
			.section-header-wrapper {
				grid-template-columns: 1fr auto !important;
			}
		}
	</style>
	<script>
		// Workaround: Clear Scalar auth from localStorage on new session (tab/window open)
		// but keep it on page reload.
		if (!sessionStorage.getItem('scalar_session_active')) {
			const scalarKeys = [
				'scalar-client-auth',
				'scalar-client-selected-security-schemes',
				'scalar-client-config-security-schemes'
			];
			scalarKeys.forEach(key => localStorage.removeItem(key));
			sessionStorage.setItem('scalar_session_active', 'true');
		}

		// Resizable sidebar implementation
		document.addEventListener('DOMContentLoaded', () => {
			// Use MutationObserver to wait for the sidebar to be rendered by Scalar
			const observer = new MutationObserver((mutations, obs) => {
				const sidebar = document.querySelector('.scalar-app aside');
				
				if (sidebar && !sidebar.querySelector('.sidebar-resizer')) {
					const resizer = document.createElement('div');
					resizer.className = 'sidebar-resizer';
					sidebar.appendChild(resizer);

					let isResizing = false;

					resizer.addEventListener('mousedown', (e) => {
						isResizing = true;
						resizer.classList.add('resizing');
						document.body.style.cursor = 'col-resize';
						document.body.style.userSelect = 'none';
						e.preventDefault();
					});

					let animationFrame;
					document.addEventListener('mousemove', (e) => {
						if (!isResizing) return;
						
						if (animationFrame) cancelAnimationFrame(animationFrame);
						
						animationFrame = requestAnimationFrame(() => {
							const newWidth = e.clientX;
							if (newWidth > 200 && newWidth < 1200) {
								const widthStr = newWidth + 'px';
								document.documentElement.style.setProperty('--scalar-sidebar-width', widthStr);
								document.body.style.setProperty('--scalar-sidebar-width', widthStr);
								const app = document.querySelector('.scalar-app');
								if (app) app.style.setProperty('--scalar-sidebar-width', widthStr);
							}
						});
					});

					document.addEventListener('mouseup', () => {
						if (isResizing) {
							isResizing = false;
							resizer.classList.remove('resizing');
							document.body.style.cursor = '';
							document.body.style.userSelect = '';
						}
					});
				}
			});

			observer.observe(document.body, {
				childList: true,
				subtree: true
			});
		});
	</script>
</head>
<body>
	%s
	<script id="api-reference" 
			data-url="%s"
			data-configuration='{"persistAuth": true}'>
	</script>
	<script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
</body>
</html>`, title, sidebarWidth, sidebarWidth, sidebarWidth, authMessageElement, specURL)

		return c.HTML(http.StatusOK, htmlContent)
	}
}
