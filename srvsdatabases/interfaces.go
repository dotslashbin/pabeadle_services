package srvsdatabases

type DatabaseINT interface {
	InitDB()
	GetClient()
}

type CreatorINT interface {
	Create(interface{}) interface{}
}
