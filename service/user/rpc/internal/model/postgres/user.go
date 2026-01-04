package postgres

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

// ExtraInfo用来维护推荐系统中与用户相关的一些字段，例如职业，性别等,非必要尽量不要随便添加
type ExtraInfo struct{  
	Nickname   string            `json:"nickname,omitempty"`
}


type User struct{
	gorm.Model
	Username     string  `gorm:"type:varchar(20);not null;"`
	Password     string  `gorm:"type:varchar(80);not null;"`
	Email        string  `gorm:"type:varchar(50);not null;"`
	Score        int16   `gorm:"type:int;not null;"`
	ExtraInfo    ExtraInfo `gorm:"type:jsonb;default:'{}'"`
}

func (e ExtraInfo) Value() (driver.Value, error) {
	return json.Marshal(e)
}

// Scan implements sql.Scanner for reading from the database.
// Parses JSON bytes (or string) into ExtraInfo struct.
func (e *ExtraInfo) Scan(value interface{}) error {
	if value == nil {
		*e = ExtraInfo{}
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New("failed to scan ExtraInfo: unsupported data type")
	}

	return json.Unmarshal(bytes, e)
}
