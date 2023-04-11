package types

// Same concept as model in MVC
type User struct {
	ID        string `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
}
