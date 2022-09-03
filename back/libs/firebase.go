package libs

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"log"
)

var FirebaseApp *firebase.App
var FirebaseAuth *auth.Client

func init() {
	var err error
	FirebaseApp, err = firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing firebase admin app: %v\n", err)
	}

	FirebaseAuth, err = FirebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing firebase auth: %v\n", err)
	}

}
