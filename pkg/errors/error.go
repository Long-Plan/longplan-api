package errors

import "errors"

func AuthErr(msg string) error {
	return errors.New("AUTH: " + msg)
}

func DatabaseErr(msg string) error {
	return errors.New("DB: " + msg)
}

func ServiceErr(msg string) error {
	return errors.New("SERVICE: " + msg)
}

func NotHasConfig() error {
	return errors.New("NOT HAS CONFIGURATION")
}

func ParseErr(msg string) error {
	return errors.New("PARSE: " + msg)
}

func InternalErr(msg string) error {
	return errors.New("INTERNAL: " + msg)
}

func CmuOauthErr(msg string) error {
	return errors.New("CMU OAUTH: " + msg)
}

func FileErr(msg string) error {
	return errors.New("FILE: " + msg)
}
