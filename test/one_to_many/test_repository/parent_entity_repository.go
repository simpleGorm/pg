package test_repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/simpleGorm/pg"
)

// Mandatory declare ID field
type ParentEntity struct {
	ID        int64  `json:"ID"`
	Name      string `json:"Name"`
	Children1 []any  `json:"Children1"`
	Children2 []any  `json:"Children2"`
}

func (p ParentEntity) GetID() int64 {
	return p.ID
}

func (parent *ParentEntity) AddRelatedEntity(related any) {
	parent.Children1 = append(parent.Children1, related)
}

type ParentEntityRepository struct {
	pg.Repository[ParentEntity]
}

const TABLE_NAME = "TEST_PARENT_ENTITY_TABLE "

var (
	ParentEntity_id   = "id"
	ParentEntity_name = "name"
)

var ParentEntity_Fields = []string{
	ParentEntity_id,
	ParentEntity_name,
}

func NewParentEntityRepository(db pg.DbClient) ParentEntityRepository {
	repo := pg.NewRepository(
		db,
		sq.Insert(TABLE_NAME).PlaceholderFormat(sq.Dollar).Columns(ParentEntity_name),
		sq.Select(ParentEntity_Fields...).PlaceholderFormat(sq.Dollar).From(TABLE_NAME),
		sq.UpdateBuilder{},
		sq.DeleteBuilder{},
		parentEntityConverter,
		func(parent ParentEntity) int64 {
			return parent.ID
		})
	child1Rel := pg.WrapRelation(OneToManyChild1EntityRelation(db))
	child2Rel := pg.WrapRelation(OneToManyChild2EntityRelation(db))
	repo.Relations = append(repo.Relations, child1Rel, child2Rel)
	return ParentEntityRepository{repo}
}

func parentEntityConverter(rows pgx.Row) any {
	var parent ParentEntity
	if err := rows.Scan(&parent.ID, &parent.Name); err != nil {
		panic(err)
	}
	return &parent
}
