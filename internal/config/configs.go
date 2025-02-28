package config

func LoadConfig() string {
	//dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")
	//dbUser := os.Getenv("DB_USER")
	//dbPassword := os.Getenv("DB_PASSWORD")
	//dbname := os.Getenv("DB_NAME")
	//
	//return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbname)

	return "./internal/database/main.db"
}
