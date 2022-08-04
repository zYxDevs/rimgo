package utils

import "os"

var Config map[string]interface{}

func LoadConfig() {
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	if os.Getenv("RIMGU_PORT") != "" {
		port = os.Getenv("RIMGU_PORT")
	}

	fiberPrefork := false
	if os.Getenv("FIBER_PREFORK") == "true" {
		fiberPrefork = true
	}

	addr := "0.0.0.0"
	if os.Getenv("ADDRESS") != "" {
		addr = os.Getenv("ADDRESS")
	}
	if os.Getenv("RIMGU_ADDRESS") != "" {
		addr = os.Getenv("RIMGU_ADDRESS")
	}

	imgurId := "546c25a59c58ad7"
	if os.Getenv("IMGUR_CLIENT_ID") != "" {
		imgurId = os.Getenv("IMGUR_CLIENT_ID")
	}
	if os.Getenv("RIMGU_IMGUR_CLIENT_ID") != "" {
		imgurId = os.Getenv("RIMGU_IMGUR_CLIENT_ID")
	}

	Config = map[string]interface{}{
		"port":         port,
		"addr":         addr,
		"imgurId":      imgurId,
		"fiberPrefork": fiberPrefork,
		"privacy": map[string]interface{}{
			"set":					 os.Getenv("PRIVACY_NOT_COLLECTED") != "",
			"policy":	       os.Getenv("PRIVACY_POLICY"),
			"message":       os.Getenv("PRIVACY_MESSAGE"),
			"country":       os.Getenv("PRIVACY_COUNTRY"),
			"provider":      os.Getenv("PRIVACY_PROVIDER"),
			"cloudflare":    os.Getenv("PRIVACY_CLOUDFLARE") == "true",
			"not_collected": os.Getenv("PRIVACY_NOT_COLLECTED") == "true",
			"ip":            os.Getenv("PRIVACY_IP") == "true",
			"url":           os.Getenv("PRIVACY_URL") == "true",
			"device":        os.Getenv("PRIVACY_DEVICE") == "true",
			"diagnostics":   os.Getenv("PRIVACY_DIAGNOSTICS") == "true",
		},
	}
}
