package config

import (
    "log"
    "github.com/kelseyhightower/envconfig"
)

type DBEnvironment struct {
    Dialect string `default:"mysql";envconfig:"DB_DIALECT"`
    Host string `default:"127.0.0.1";envconfig:"DB_HOST"`
    Port int `default:"3306";envconfig:"DB_PORT"`
    Username string `envconfig:"DB_USERNAME"`
    Password string `envconfig:"DB_PASSWORD"`
    Name string `envconfig:"DB_NAME"`
}

type WebEnvironment struct {
    Port int `default: "8080";envconfig:"WEB_PORT"`
    Host string `default: "127.0.0.1";envconfig:"WEB_HOST"`
}

var DBConfig DBEnvironment
var WebConfig WebEnvironment

func init () {
    if err :=  envconfig.Process("", &DBConfig); err != nil {
        log.Fatalf("Failed to process env var : %s", err)
        return
    }
    if err := envconfig.Process("", &WebConfig); err != nil {
        log.Fatalf("Failed to process env var : %s", err)
        return
    }
}
