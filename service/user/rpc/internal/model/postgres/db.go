package postgres

import (
	"fmt"
	"time"
	"user/rpc/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB // 建议首字母大写以便外部访问

func InitDB(c config.Config) (err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.User,
		c.Postgres.Password,
		c.Postgres.Dbname,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //禁用复数表名
		},
	})

	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}

	// 获取底层 *sql.DB 以设置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}

	// 自动迁移模型
	err= DB.AutoMigrate(
		&User{},
	)
	if err != nil {
		return fmt.Errorf("failed to auto migrate models: %w", err)
	}

	// 设置连接池
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	return nil
}
