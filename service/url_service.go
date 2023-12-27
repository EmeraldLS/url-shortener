package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/EmeraldLS/url-shortener/db"
	"github.com/EmeraldLS/url-shortener/model"
	"github.com/golang-module/carbon"
	"go.mongodb.org/mongo-driver/bson"
)

var url_collection = db.URL_COLLECTIONS

func SaveUrl(url *model.URL) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	_, err := url_collection.InsertOne(ctx, url)
	if err != nil {
		return fmt.Errorf("unable to save url -> %v", err)
	}
	return nil
}

func GetURlId(id string) (model.URL, error) {
	url, err := CheckExpired(id)
	if err != nil {
		return model.URL{}, err
	}
	return url, nil
}

func CheckExpired(id string) (model.URL, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "id", Value: id}}
	var url model.URL
	err := url_collection.FindOne(ctx, filter).Decode(&url)
	if err != nil {
		return model.URL{}, errors.New("invalid url")
	}

	if time.Now().After(url.ExpirationTime) {
		return model.URL{}, errors.New("url expired")

	}
	return url, nil
}

// Used bson.D with $lt to get all urls with expiration time less than now
// Then delete all of them
func DeleteExpiredURLS() error {
	currentTime := time.Now()

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "expiration_time", Value: bson.D{{Key: "$lt", Value: currentTime}}}}

	res, err := url_collection.DeleteMany(ctx, filter)
	if err != nil {
		return fmt.Errorf("unable to delete expired urls -> %v", err)
	}
	fmt.Printf("Deleted expired urls -> %d\n", res.DeletedCount)
	return nil
}

// It periodically checks the database with expiration_url less than now's time every 10mins
// then it deletes the expired urls
func AutoExpiredDeleteURL() {
	ticker := time.NewTicker(time.Minute * 10)
	defer ticker.Stop()
	for range ticker.C {
		DeleteExpiredURLS()
	}
}

func UpdateURLClicks(id string) error {
	url, err := GetURlId(id)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	var updateObj = make(bson.M)
	filter := bson.D{{Key: "id", Value: id}}
	url.NoOfClicks++
	updateObj["no_of_clicks"] = url.NoOfClicks
	updateObj["updated_at"] = carbon.Now().ToDateTimeString()
	update := bson.M{"$set": updateObj}

	_, err = url_collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("unable to update this url content -> %v", err)
	}
	return nil
}
