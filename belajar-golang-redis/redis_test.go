package belajar_golang_redis

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   3,
})

func TestConnection(t *testing.T) {
	assert.NotNil(t, client)

	// err := client.Close()
	// assert.Nil(t, err)
}

var ctx = context.Background()

func TestPing(t *testing.T) {
	result, err := client.Ping(ctx).Result()
	assert.Nil(t, err)
	assert.Equal(t, "PONG", result)
}

func TestString(t *testing.T) {
	client.SetEx(ctx, "name", "Azmi", time.Second*3)

	result, err := client.Get(ctx, "name").Result()
	assert.Nil(t, err)
	assert.Equal(t, "Azmi", result)

	time.Sleep(time.Second * 5)
	_, err = client.Get(ctx, "name").Result()
	assert.NotNil(t, err)
}

func TestList(t *testing.T) {
	client.RPush(ctx, "names", "Muhamad")
	client.RPush(ctx, "names", "Habibi")
	client.RPush(ctx, "names", "Azmi")

	assert.Equal(t, "Muhamad", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Habibi", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Azmi", client.LPop(ctx, "names").Val())

	client.Del(ctx, "names")
}

func TestSet(t *testing.T) {
	client.SAdd(ctx, "students", "Muhamad")
	client.SAdd(ctx, "students", "Muhamad")
	client.SAdd(ctx, "students", "Habibi")
	client.SAdd(ctx, "students", "Habibi")
	client.SAdd(ctx, "students", "Azmi")
	client.SAdd(ctx, "students", "Azmi")

	assert.Equal(t, int64(3), client.SCard(ctx, "students").Val())
	assert.Equal(t, []string{"Habibi", "Muhamad", "Azmi"}, client.SMembers(ctx, "students").Val())
}

func TestSortedSet(t *testing.T) {
	client.ZAdd(ctx, "scores", redis.Z{Score: 100, Member: "Muhamad"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 85, Member: "Habibi"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 95, Member: "Azmi"})

	assert.Equal(t, []string{"Habibi", "Azmi", "Muhamad"}, client.ZRange(ctx, "scores", 0, -1).Val())

	assert.Equal(t, "Muhamad", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Azmi", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Habibi", client.ZPopMax(ctx, "scores").Val()[0].Member)
}

func TestGeoPoint(t *testing.T) {
	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Toko A",
		Latitude:  -6.897492117466138,
		Longitude: 107.61688446863279,
	})
	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Toko B",
		Latitude:  -6.9066381017327805,
		Longitude: 107.62285880833143,
	})

	assert.Equal(t, 1.2122, client.GeoDist(ctx, "sellers", "Toko A", "Toko B", "km").Val())

	sellers := client.GeoSearch(ctx, "sellers", &redis.GeoSearchQuery{
		Latitude:   -6.900887946224864,
		Longitude:  107.61601158031317,
		Radius:     1000,
		RadiusUnit: "km",
	}).Val()

	// fmt.Println(sellers)
	assert.Equal(t, []string{"Toko A", "Toko B"}, sellers)
}

func TestPublishStream(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := client.XAdd(ctx, &redis.XAddArgs{
			Stream: "members",
			Values: map[string]interface{}{
				"name":    "Azmi " + strconv.Itoa(i),
				"address": "Indonesia",
			},
		}).Err()
		assert.Nil(t, err)
	}
}

func TestCreateConsumerGroup(t *testing.T) {
	client.XGroupCreate(ctx, "members", "group-1", "0")
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-1")
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-2")
}

func TestGetStream(t *testing.T) {
	streams := client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    "group-1",
		Consumer: "consumer-1",
		Streams:  []string{"members", ">"},
		Count:    2,
		Block:    5 * time.Second,
		NoAck:    false,
	}).Val()

	for _, stream := range streams {
		for _, message := range stream.Messages {
			fmt.Println(message.ID)
			fmt.Println(message.Values)
		}
	}
}

func TestSubscribePubSub(t *testing.T) {
	subscriber := client.Subscribe(ctx, "channel-1")
	defer subscriber.Close()
	for i := 0; i < 10; i++ {
		message, err := subscriber.ReceiveMessage(ctx)
		assert.Nil(t, err)
		fmt.Println(message.Payload)
	}
}

func TestPublishPubSub(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := client.Publish(ctx, "channel-1", "Hello "+strconv.Itoa(i)).Err()
		assert.Nil(t, err)
	}
}
