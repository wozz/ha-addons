package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type addonInfo struct {
	IngressURL string `json:"ingress_url"`
}

type supervisorResp struct {
	Data addonInfo `json:"data"`
}

func fetchIngressURL() string {
	req, err := http.NewRequest(http.MethodGet, "http://supervisor/addons/12c9acea_gollum/info", nil)
	if err != nil {
		log.Fatalf("could not create request: %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+os.Getenv("SUPERVISOR_TOKEN"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("could not do request: %v", err)
	}
	defer resp.Body.Close()
	var info supervisorResp
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		log.Fatalf("could not determine ingress url: %v", err)
	}
	return info.Data.IngressURL
}

func main() {
	ingressURL := fetchIngressURL()
	prxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:4567",
		Path:   ingressURL,
	})
	ln, err := net.Listen("tcp", ":8099")
	if err != nil {
		log.Fatalf("could not listen: %v", err)
	}
	err = http.Serve(ln, prxy)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
		log.Println("clean exit")
	}
}
