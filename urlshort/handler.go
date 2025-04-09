package urlshort

import (
	"net/http"
	yaml "gopkg.in/yaml.v2"
	"encoding/json"
)

/*
	MapHandler returns an http.HandlerFunc which implements http.Handler.
	This will attempt to map any paths (keys in the map) to their corresponding URL (values).
	If the path is not provided in the map then the fallback http.Handler will be called instead.
*/

type pathToURL struct {
	Path string
	URL string
}


func mapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		url := pathsToUrls[request.URL.Path]
		if url != "" {
			http.Redirect(response, request, url, http.StatusPermanentRedirect)
	} else {
		fallback.ServeHTTP(response, request)
	}
    })
}

func buildMap(pathsToURLs []pathToURL) (builtMap map[string]string) {
	builtMap = make(map[string]string)
	for _, ptu := range pathsToURLs {
		builtMap[ptu.Path] = ptu.URL
	}
	return
}

/*
	 YAML Handler will parse the provided YAML and then return an http.HandlerFunc 
	 that will attempt to map any paths to their corresponding URL. 
	 Again if no path is provided the fallback http.Handler will be called instead.

	 YAML is expected to be in the format:
	 	
	 	- path: /some-path
		  url: https://www.some-url.com/example
*/

func parseYAML(yml []byte) (pathsToURLs []pathToURL, err error) {
	err = yaml.Unmarshal(yml, &pathsToURLs)
	return
}

func YAMLHandler(yml []byte, fallback http.Handler) (yamlHandler http.HandlerFunc, err error) {
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return
	}
	pathMap := buildMap(parsedYaml)
	yamlHandler = mapHandler(pathMap, fallback)
	return
}

func parseJSON(jsn []byte) (pathsToURLs []pathToURL, err error) {
	err = json.Unmarshal(jsn, &pathsToURLs)
	return
}

func JSONHandler(jsn []byte, fallback http.Handler) (jsonHandler http.HandlerFunc, err error) {
	parsedJson, err := parseJSON(jsn)
	if err != nil {
		return
	}
	pathMap := buildMap(parsedJson)
	jsonHandler = mapHandler(pathMap, fallback)
	return
}
