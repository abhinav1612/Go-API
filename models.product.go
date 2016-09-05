pacakge main

type Product struct {
    gorm.Model
    Name string `gorm:"type:varchar(100);not null"`
    Searchkey string `gorm:"type:varchar(100);unique"`
    MonthlyRentalPrice float64
}

func createProduct(name string, price float64) Product {

  product := Product{
      //CreatedAt:    time.Now().UnixNano(),
      //UpdatedAt:    time.Now().UnixNano(),
      Name:         name,
      Searchkey:    RandomString(14),
      MonthlyRentalPrice: price ,
  }
  return product
}

// RandomString Generator
func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
