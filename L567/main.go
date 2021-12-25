package main

import (
	"L567/config"
	"L567/db"
	"L567/handlers"
	"L567/internal"
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	postgress:= db.NewDB(config.DB_CONNECTION_STRING)
	repo:= internal.NewRepo(postgress)
	userSvc:=internal.NewUserService(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/login", handlers.MakeLoginHandler(userSvc))
	mux.HandleFunc("/signup", handlers.MakeSignUpHandler(userSvc))
	mux.HandleFunc("/checksignup", handlers.MakeCheckSignUpHandler(userSvc))
	mux.HandleFunc("/pcabinet", handlers.MakePersonalCabinetHandler(userSvc))

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS13,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
		//	tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		//	tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			//tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			//tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:         ":443",
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	log.Fatal(srv.ListenAndServeTLS(config.SERTIFICATE_PATH, config.PRIVATE_KEY_PATH))
}


//db/migrations -database "postgresql://root:root@localhost:5432/l567?sslmode=disable" up