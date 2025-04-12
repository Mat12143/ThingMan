package db

type Item struct {
	Id         int64  "json:\"id\";gorm:\"primaryKey\""
	Name       string "json:\"name\";gorm:\"notNull\""
	Quantity   int    "json:\"quantity\";gorm:\"notNull\""
	Location   int    "json:\"location\";gorm:\"notNull\""
	CreatedAt  int64  `json:"createdAt"`
	ModifiedAt int64  "json:\"modifiedAt\";gorm:\"autoUpdateTime:milli\""
}
