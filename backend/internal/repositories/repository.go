package repositories

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/AntonKhPI2/nim-one-pile/internal/game"
	"github.com/AntonKhPI2/nim-one-pile/internal/models"
)

var ErrNotFound = errors.New("not found")

type GameRepo interface {
	Save(ctx context.Context, s *game.State) error
	Get(ctx context.Context, id string) (*game.State, error)
}

type GormGameRepo struct {
	DB *gorm.DB
}

func NewGormGameRepo(db *gorm.DB) GameRepo {
	return &GormGameRepo{DB: db}
}

func (r *GormGameRepo) Save(ctx context.Context, s *game.State) error {
	row := toRow(s)

	return r.DB.WithContext(ctx).
		Clauses(onConflictGames()).
		Model(&models.Game{}).
		Create(&row).Error
}

func (r *GormGameRepo) Get(ctx context.Context, id string) (*game.State, error) {
	var row models.Game
	err := r.DB.WithContext(ctx).First(&row, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return toDomain(&row), nil
}

func toRow(s *game.State) models.Game {
	return models.Game{
		ID:         s.ID,
		Variant:    string(s.Variant),
		K:          s.K,
		Remaining:  s.Remaining,
		PlayerTurn: s.PlayerTurn,
		Winner:     s.Winner,
	}
}

func toDomain(r *models.Game) *game.State {
	return &game.State{
		ID:         r.ID,
		Variant:    game.Variant(r.Variant),
		K:          r.K,
		Remaining:  r.Remaining,
		PlayerTurn: r.PlayerTurn,
		Winner:     r.Winner,
	}
}

func onConflictGames() clause.OnConflict {
	return clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"variant",
			"k",
			"remaining",
			"player_turn",
			"winner",
			"updated_at",
		}),
	}
}
