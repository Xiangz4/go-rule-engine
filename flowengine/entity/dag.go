package entity

type Dag struct {
	BaseInfo `yaml:",inline" json:",inline" bson:"inline"`
	Name     string `yaml:"name,omitempty" json:"name,omitempty" bson:"name,omitempty"`
}

type DagInstance struct {
	BaseInfo `bson:"inline"`
}
