package middleware

import (
	"net/http"
	"strings"

	"github.com/PasinduYeshan/go-sample-oauth/internal/common/response" // Added for standardized error responses
	"github.com/PasinduYeshan/go-sample-oauth/internal/config"
	"github.com/labstack/echo/v4"
)

// ScopeAuthMiddleware creates an Echo middleware function that checks if the
// incoming request has the required scopes based on a predefined configuration.
// For now, it simulates JWT parsing and uses a hardcoded set of scopes for the user.
func ScopeAuthMiddleware() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			scopeCfg := config.GetScopeConfig()
			if scopeCfg == nil {
				// Configuration not loaded, deny access as a security precaution
				return response.ErrorResponse(c, http.StatusInternalServerError, "API scope configuration not loaded")
			}

			requestPath := c.Path() // This might need adjustment if using route patterns with parameters
			requestMethod := c.Request().Method

			var requiredScopes []string
			resourceFound := false

			for _, resource := range scopeCfg.APIResources {
				// Simple path matching. For more complex scenarios (e.g., path parameters),
				// you might need a more sophisticated matching logic or rely on Echo's route naming.
				if resource.Path == requestPath && resource.Method == requestMethod {
					requiredScopes = resource.Scopes
					resourceFound = true
					break
				}
			}

			if !resourceFound {
				// No specific scope configuration for this path/method.
				// Depending on policy, either allow or deny. For now, let's allow.
				// If you want to deny by default, return an error here.
				return next(c)
			}

			if len(requiredScopes) == 0 {
				// Resource is configured but requires no specific scopes.
				return next(c)
			}

			// --- JWT Extraction and Parsing (Simulated) ---
			// In a real scenario:
			// 1. Extract token from "Authorization: Bearer <token>" header.
			// 2. Validate the JWT (signature, expiry, issuer) with WSO2 Asgardeo.
			// 3. Parse the validated JWT to get the 'scope' claim.

			// For now, simulate with hardcoded scopes.
			// Replace this with actual JWT parsing logic later.
			// Example: userScopesFromToken := []string{"read:ads", "internal_scope"}

			// Let's simulate getting scopes from a header for easier testing without full JWT parsing yet.
			// Client should send a header like: X-User-Scopes: scope1 scope2
			userScopesHeader := c.Request().Header.Get("X-User-Scopes")
			userScopesFromToken := []string{}
			if userScopesHeader != "" {
				userScopesFromToken = strings.Fields(userScopesHeader) // Splits by whitespace
			}
			// --- End of Simulated JWT Parsing ---

			if !hasRequiredScopes(userScopesFromToken, requiredScopes) {
				return response.ErrorResponse(c, http.StatusForbidden, "Insufficient scope")
			}

			return next(c)
		}
	}
}

// hasRequiredScopes checks if the userScopes list contains all the requiredScopes.
func hasRequiredScopes(userScopes []string, requiredScopes []string) bool {

	userScopeSet := make(map[string]bool)
	for _, s := range userScopes {
		userScopeSet[s] = true
	}

	for _, reqScope := range requiredScopes {
		if !userScopeSet[reqScope] {
			return false
		}
	}
	return true
}
