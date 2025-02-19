// file này tạo 1 instance database
package db

import (
	"engineerpro_ex_week4/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitDB(cfg *config.Config) (*sqlx.DB, error) {

	// Để connect tới DB thì cần 1 DataSourceName. --> Nó là 1 string chứa các dữ liệu để kết nối tới DB
	// 1 DataSourceName sẽ có cấu trúc như sau DBUser:DBPassword @Protocol DBHost DBPort và DBName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	// Dùng DSN để kết nối tới MySQL DB
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	query := `
		CREATE TABLE IF NOT EXISTS user(
			username NVARCHAR(255) PRIMARY KEY,
			password NVARCHAR(255),
			image_path NVARCHAR(255)
		)
	`

	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("failed to create user table: %w", err)
	}

	return db, nil

}
