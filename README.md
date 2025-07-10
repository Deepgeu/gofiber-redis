# gofiber-redis
reduce delay by using redis along go- fiber to making api calling faster

for practise purpose :- seeding 1000 to 1000000 random records
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
