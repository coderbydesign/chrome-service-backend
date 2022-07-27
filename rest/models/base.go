package models

import (
  "time"
)

// Generic Struct used throughout models in this service.
type BaseModel struct {
  Id uint
  CreatedAt time.Time
  UpdatedAt time.Time
}