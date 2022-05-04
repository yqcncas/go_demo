package helper

import uuid "github.com/satori/go.uuid"

func GetUUID() string {
	uuid := uuid.NewV4().String()
	return uuid
}
