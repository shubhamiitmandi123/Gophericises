package main

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler : will return an http.HandlerFunc (which also
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		if url, ok := pathsToUrls[path]; ok {
			http.Redirect(w, req, url, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, req)
	}
}

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// YAMLHandler : will parse the provided YAML and then return
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	var pathUrls []pathURL
	err := yaml.Unmarshal(yml, &pathUrls)
	if err != nil {
		return nil, err
	}
	pathsToUrls := make(map[string]string)

	for _, v := range pathUrls {
		pathsToUrls[v.Path] = v.URL
	}
	return MapHandler(pathsToUrls, fallback), nil
}
