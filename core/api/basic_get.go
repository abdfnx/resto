package api

import (
	"fmt"
	"net/http"

	httpClient "github.com/abdfnx/resto/client"
	"github.com/abdfnx/resto/validation"
)

// BasicGet sends a simple GET request to the url with any potential parameters like `Tokens` or `Basic Auth`
func BasicGet(
		httpURL,
		method,
		authType,
		bearerToken,
		basicAuthUsername,
		basicAuthPassword string,
	) (string, string, string, error) {
	var url, err = validation.CheckURL(httpURL)

	if err != nil {
		return "", "", "", err
	}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "", "", "", fmt.Errorf("Error creating request: %s", err.Error())
	}

	if authType == "bearer" {
		req.Header.Set("Authorization", "Bearer " + bearerToken)
	} else if authType == "basic" {
		req.Header.Set("Authorization", "Basic " + basicAuth(basicAuthUsername, basicAuthPassword))
	}

	client := httpClient.HttpClient()
	res, err := client.Do(req)

	if err != nil {
		return "", "", "", fmt.Errorf("Error sending request: %s", err.Error())
	}

	defer res.Body.Close()

	return formatResponse(res)
}
