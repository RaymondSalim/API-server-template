package models

/*
	Each model has its own separate files.
	In the models package there are the request and response packages which contain the request and response types used in the controllers and services.

	The models contain struct tags for easier mapping of database results to objects.
	The models also contain struct tags for json in case they are mashalled to json directly in an endpoint.
	Whenever the HTTP endpoints response is different from the model, a response type is made in the response package, and the mapping is done in the service.
*/

const tableName = "foo-table"

type Foo struct {
	ID      uint   `gorm:"primaryKey"`
	FooName string `gorm:"column:name"`
}

func (Foo) TableName() string {
	return tableName
}
