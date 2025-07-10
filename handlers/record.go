package handlers

import (
	"fmt"
	"strconv"
	"time"

	"gofiber-redis/db"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2"
)

// Type definition
type Record struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//  Generate fake 1,000,000 keys into Redis
func SeedRecords(c *fiber.Ctx) error {
	for i := 1; i <= 1000000; i++ {
		key := fmt.Sprintf("key_%d", i)
		value := gofakeit.HipsterSentence(10)

		err := db.RDB.Set(db.Ctx, key, value, 0).Err()
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf(" Insert failed at %d: %v", i, err))
		}
	}
	return c.SendString("Seeded 1,000,000 fake Redis records")
}

//  GET /record/:key
func GetRecord(c *fiber.Ctx) error {
	key := c.Params("key")
	value, err := db.RDB.Get(db.Ctx, key).Result()

	if err != nil {
		return c.Status(404).SendString(" Key not found")
	}
	return c.JSON(fiber.Map{"key": key, "value": value})
}

// OST /record
func CreateRecord(c *fiber.Ctx) error {
	var r Record
	if err := c.BodyParser(&r); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

	err := db.RDB.Set(db.Ctx, r.Key, r.Value, 0).Err()
	if err != nil {
		return c.Status(500).SendString("Insert error: " + err.Error())
	}
	return c.SendString(" Record created")
}

//  PUT /record/:key
func UpdateRecord(c *fiber.Ctx) error {
	key := c.Params("key")
	var payload struct {
		Value string `json:"value"`
	}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).SendString("Invalid body")
	}

	err := db.RDB.Set(db.Ctx, key, payload.Value, 0).Err()
	if err != nil {
		return c.Status(500).SendString("Update failed: " + err.Error())
	}
	return c.SendString("Record updated")
}

//  DELETE /record/:key
func DeleteRecord(c *fiber.Ctx) error {
	key := c.Params("key")
	_, err := db.RDB.Del(db.Ctx, key).Result()
	if err != nil {
		return c.Status(500).SendString(" Delete error: " + err.Error())
	}
	return c.SendString(" Record deleted")
}
