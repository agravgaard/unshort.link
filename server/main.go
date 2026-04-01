package main

import (
	"database/sql"
	"embed"
	"flag"
	"net/http"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"unshort.link/blacklist"
)

//go:embed static/*
var staticFiles embed.FS

var serveUrl, port, supportUrl string
var blacklistSyncInterval time.Duration
var blacklistUrls []string

func init() {
	flag.StringVar(&serveUrl, "url", "http://localhost:8080", "The server url this server runs on. (Required for the frontend)")
	flag.StringVar(&port, "port", "8080", "Port to run the server on")
	flag.StringVar(&supportUrl, "support-url", "", "Url to display in support notice")
	flag.DurationVar(&blacklistSyncInterval, "sync", time.Hour, "Blacklist synchronization interval")
	rawBlacklistUrls := flag.String("blacklist-sources", "https://hosts.ubuntu101.co.za/domains.list", "Comma separated list of blacklist urls to periodically sync")
	flag.Parse()
	blacklistUrls = strings.Split(*rawBlacklistUrls, ",")
}

func main() {
	db, err := sql.Open("sqlite3", "file:blacklist.db")
	if err != nil {
		panic("Couldn't create database for host blacklist")
	}
	blacklistSource := blacklist.NewSqliteRepository(db)
	go blacklist.NewLoader(blacklistUrls, blacklistSource, blacklistSyncInterval).StartSync()

	handler := func(rw http.ResponseWriter, req *http.Request) {
		switch path := req.URL.Path; {
		case path == "" || path == "/" ||
			path == "/d/" || path == "/d" ||
			path == "/api/" || path == "/api" ||
			path == "/nb/" || path == "/nb":
			handleIndex(rw, true)
		case path == "/about":
			browserExtension := false
			if req.URL.Query().Get("extension") == "true" {
				browserExtension = true
			}
			handleAbout(rw, browserExtension)
		case strings.HasPrefix(path, "favicon.ico"):
			rw.WriteHeader(http.StatusNotFound)
		case strings.HasPrefix(path, "/providers"):
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			rw.Header().Set("Access-Control-Allow-Headers", "*")
			rw.Header().Set("Content-type", "application/json")
			handleProviders(rw)
		case strings.HasPrefix(path, "/api/"):
			req.URL.Path = strings.TrimPrefix(req.URL.Path, "/api")
			handleUnShort(rw, req, false, true, true, blacklistSource, supportUrl)
		case strings.HasPrefix(path, "/d/"):
			req.URL.Path = strings.TrimPrefix(req.URL.Path, "/d")
			handleUnShort(rw, req, false, false, true, blacklistSource, supportUrl)
		case strings.HasPrefix(path, "/nb/"):
			req.URL.Path = strings.TrimPrefix(req.URL.Path, "/nb")
			handleUnShort(rw, req, false, false, false, blacklistSource, supportUrl)
		default:
			handleUnShort(rw, req, true, false, true, blacklistSource, supportUrl)
		}
	}

	http.Handle("/static/", http.FileServer(http.FS(staticFiles)))
	http.HandleFunc("/", handler)

	log.Infof("Run server on port '%s', with url '%s'", port, serveUrl)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
