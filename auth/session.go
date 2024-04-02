package auth

import (
	"encoding/json"
	"pagoda/database"
	"pagoda/models"
	"time"

	"github.com/google/uuid"
)

func GenerateSession(user models.User) (string, error) {
	sessionId := uuid.NewString()

	rdb := database.Cache
	jsonData, err := json.Marshal(user)

	if err != nil {
		return "", err
	}

	err = rdb.Set(database.Ctx, sessionId, string(jsonData), time.Hour*24*30).Err()

	if err != nil {
		return "", err
	}

	return sessionId, nil
}

func GetSession(sessionId string) (models.User, error) {
	rdb := database.Cache
	val, err := rdb.Get(database.Ctx, sessionId).Result()

	if err != nil {
		return models.User{}, err
	}

	var user models.User
	err = json.Unmarshal([]byte(val), &user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func DeleteSession(sessionId string) error {
	rdb := database.Cache
	_, err := rdb.Del(database.Ctx, sessionId).Result()

	if err != nil {
		return err
	}

	return nil
}
