package entity

type BaseInfo struct {
	ID        string `yaml:"id" json:"id" bson:"_id"`
	CreatedAt int64  `yaml:"createdAt" json:"createdAt" bson:"createdAt"`
	UpdatedAt int64  `yaml:"updatedAt" json:"updatedAt" bson:"updatedAt"`
}
