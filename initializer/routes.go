package initializer

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/Atipat-CMU/api-gateway/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) *gin.Engine {

	// CORS middleware setup
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},                                       // Allow frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                     // Allow methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"}, // Allow necessary headers
		AllowCredentials: true,                                                                    // Allow credentials (cookies, auth headers)
		MaxAge:           12 * time.Hour,                                                          // Cache preflight request
	}))

	// JWT middleware for authentication
	router.Use(middleware.JWTAuthMiddleware())
	router.Use(requestLogger())

	// Explicitly handling OPTIONS requests (could be optional with proper CORS middleware)
	router.OPTIONS("/*proxyPath", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	projectURL := fmt.Sprintf(
		"http://%s:%s/api/v1",
		os.Getenv("PROJECT_SERVICE_HOST"),
		os.Getenv("PROJECT_SERVICE_PORT"),
	)
	router.GET("/project-management-service/*proxyPath", reverseProxy(projectURL))
	router.POST("/project-management-service/*proxyPath", reverseProxy(projectURL))

	return router
}

func reverseProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetURL, err := url.Parse(target)
		if err != nil {
			log.Printf("Failed to parse target URL: %s, error: %v", target, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header     // Forward headers
			req.Host = targetURL.Host         // Match the Host header
			req.URL.Scheme = targetURL.Scheme // Set the scheme
			req.URL.Host = targetURL.Host     // Set the host

			// Combine target path and proxyPath
			proxyPath := c.Param("proxyPath")
			req.URL.Path = targetURL.Path + proxyPath // Combine base path and proxy path
		}

		// proxy.ModifyResponse = func(resp *http.Response) error {
		// 	// Remove duplicate CORS headers
		// 	resp.Header.Del("Access-Control-Allow-Origin")
		// 	return nil
		// }

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
		log.Printf("Response Status: %d", c.Writer.Status())
	}
}
