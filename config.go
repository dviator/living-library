package main

type Config struct {
	GoodreadsDeveloperKey string
	GoodreadsSecret       string
}

func GetConfig() Config {
	conf := Config{
		GoodreadsDeveloperKey: "Jt1pNyo35UbSkSlSTnA2Sg",
		GoodreadsSecret:       "acYLWf8fU6eVQqmuzmIWXOh3QhwABwkRJzXuxlPA",
	}
	return conf
}
