package videos

import (
	"context"
	"errors"
	"github.com/go-pg/pg"
	"time"
)

var NotFound = errors.New("not found")

func NewService(db *pg.DB) *VideosService {
	return &VideosService{
		db: db,
	}
}

type VideosService struct {
	db *pg.DB
}

type Video struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *VideosService) GetAllVideos(ctx context.Context) ([]*Video, error) {
	videos := make([]*Video, 0)
	if err := s.db.WithContext(ctx).Model(&videos).OrderExpr("id DESC").Select(); err != nil {
		return nil, err
	}

	return videos, nil
}

func (s *VideosService) GetVideoById(ctx context.Context, videoID int64) (*Video, error) {
	video := &Video{
		ID: videoID,
	}
	if err := s.db.WithContext(ctx).Select(video); err != nil {
		return nil, err
	}
	return video, nil
}

func (s *VideosService) CreateVideo(ctx context.Context, video *Video) error {
	if err := s.db.WithContext(ctx).Insert(video); err != nil {
		return err
	}
	return nil
}

func (s *VideosService) UpdateVideoById(ctx context.Context, video *Video) error {
	video.UpdatedAt = time.Now()
	res, err := s.db.WithContext(ctx).Model(video).Column("name").WherePK().Returning("*").Update()
	if err == pg.ErrNoRows {
		return NotFound
	}
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return NotFound
	}
	return nil
}

func (s *VideosService) DeleteVideoById(ctx context.Context, videoID int64) error {
	video := Video{
		ID: videoID,
	}
	res, err := s.db.Model(&video).WherePK().Delete()
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return NotFound
	}
	return nil
}
