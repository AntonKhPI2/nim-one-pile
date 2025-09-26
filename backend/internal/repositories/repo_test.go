package repositories

import (
	"context"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/AntonKhPI2/nim-one-pile/internal/game"
)

type gameSchemaLite struct {
	ID         string    `gorm:"type:text;primaryKey"`
	UserID     *string   `gorm:"type:text;index"`
	Variant    string    `gorm:"type:text;not null"`
	K          int       `gorm:"not null"`
	Remaining  int       `gorm:"not null"`
	PlayerTurn string    `gorm:"type:text;not null"`
	Winner     string    `gorm:"type:text;not null;default:''"`
	ReasonCode string    `gorm:"type:text;not null;default:''"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

func (gameSchemaLite) TableName() string { return "games" }

func newTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	dsn := "file::memory:?cache=shared&_foreign_keys=on&parseTime=true"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("gorm open: %v", err)
	}
	if err := db.AutoMigrate(&gameSchemaLite{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return db
}

func TestGormGameRepo_SaveAndGet(t *testing.T) {
	db := newTestDB(t)
	repo := NewGormGameRepo(db)

	s := &game.State{
		ID:         "test1",
		Variant:    game.Normal,
		K:          3,
		Remaining:  21,
		PlayerTurn: "human",
		Winner:     "",
	}
	if err := repo.Save(context.Background(), s); err != nil {
		t.Fatalf("save: %v", err)
	}
	got, err := repo.Get(context.Background(), "test1")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.ID != s.ID || got.K != s.K || got.Remaining != s.Remaining || got.PlayerTurn != s.PlayerTurn {
		t.Fatalf("mismatch: got=%+v want=%+v", got, s)
	}

	s.PlayerTurn = "computer"
	if err := repo.Save(context.Background(), s); err != nil {
		t.Fatalf("save upsert: %v", err)
	}
	got2, err := repo.Get(context.Background(), "test1")
	if err != nil {
		t.Fatalf("get2: %v", err)
	}
	if got2.PlayerTurn != "computer" {
		t.Fatalf("expected upsert to update player_turn, got %q", got2.PlayerTurn)
	}
}
