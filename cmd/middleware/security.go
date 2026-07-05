package middleware

import (
	"github.com/labstack/echo/v5"
	"github.com/spf13/viper"
	"strings"
)

func SecurityHeaders() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {

			// Prevent MIME type sniffing
			// Protects against: Drive-by downloads, content type confusion attacks
			c.Response().Header().Set("X-Content-Type-Options", "nosniff")

			// Prevent clickjacking attacks
			// Protects against: UI redress attacks where malicious site frames application within an iframe
			c.Response().Header().Set("X-Frame-Options", "DENY")

			// Enable XSS filtering in older browsers
			// Note: Modern browsers have this build-in, but doesn't hurt
			c.Response().Header().Set("X-XSS-Protection", "1; mode=block")

			// Force HTTPS for 1 year
			// Protects against: MITM attacks, protocol downgrade attacks
			// Note: Only set if running on HTTPS in production
			if viper.GetBool("ENABLE_HSTS") {
				c.Response().Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
			}

			// Content Security Policy
			// Protects against: XSS, data injection attacks
			// Note: Adjust based on your frontend requirements
			csp := viper.GetString("CONTENT_SECURITY_POLICY")
			if csp == "" {
				// Default CSP - restrictive but safe
				csp = "default-src 'self'; script-src 'self'; style-src 'self' 'unsafe-inline'; img-src 'self' data: https:; font-src 'self'; connect-src 'self'; frame-ancestors 'none'"
			}
			c.Response().Header().Set("Content-Security-Policy", csp)

			// Referrer policy - control referrer information
			// Protects against: Information leakage via referrer header
			c.Response().Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

			// Permissions policy (formerly Feature-Policy)
			// Restrict browser features that can be used
			c.Response().Header().Set("Permissions-Policy", "geolocation=(), microphone=(), camera=(), payment=()")

			// Prevent browser caching of sensitive data
			// For API responses, we don't want caching
			if strings.HasPrefix(c.Path(), "/record-flow/") && c.Path() != "/record-flow/health/live" {
				c.Response().Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, private")
				c.Response().Header().Set("Pragma", "no-cache")
				c.Response().Header().Set("Expires", "0")
			}
			return next(c)
		}
	}
}
