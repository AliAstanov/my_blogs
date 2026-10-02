package storage

import (
	log "github.com/AliAstanov/my_blogs/pkg/logger"
	"github.com/AliAstanov/my_blogs/storage/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StorageI interface {
	GetContentRepo() postgres.ContnentRepoI
	GetOwnerRepo() postgres.OwnerRepoI
	GetCommonRepo() postgres.CommonRepoI
}

type storage struct {
	contentRepo  postgres.ContnentRepoI
	ownerRepo postgres.OwnerRepoI
	commonRepo postgres.CommonRepoI
}

func NewStorage(db *pgxpool.Pool, log log.Log)StorageI {
	return &storage{
		contentRepo: postgres.NewContentRepo(db,log),
		ownerRepo: postgres.NewOwnerRepo(db,log),
		commonRepo: postgres.NewCommonRepo(db,log),
	}
}

func(s *storage)GetContentRepo() postgres.ContnentRepoI{
	return s.contentRepo
}
func(s *storage)GetOwnerRepo() postgres.OwnerRepoI{
	return s.ownerRepo
}
func(s *storage)GetCommonRepo() postgres.CommonRepoI{
return s.commonRepo
}
