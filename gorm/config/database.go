package config

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
    // DSN: user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
    dsn := "root:@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}
