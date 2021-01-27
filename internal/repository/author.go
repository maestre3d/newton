package repository

import (
	"context"

	"github.com/maestre3d/newton/internal/aggregate"
	"github.com/maestre3d/newton/internal/valueobject"
	"go.uber.org/zap"
)

// Author handles all persistence interactions
type Author interface {
	// Save stores, update or deletes the given record
	Save(context.Context, aggregate.Author) error
	// Get returns the current aggregate if available, returns nil if not found
	Get(context.Context, valueobject.AuthorID) (*aggregate.Author, error)
	// Search returns a list of the current aggregate filtering and ordering by the given criteria, returns the
	// next page token as second argument and returns nil if not found
	Search(context.Context, Criteria) ([]*aggregate.Author, string, error)
}

// NewAuthor wraps the given Author repository (referenced as `root`) with observability tooling such as
// logging, tracing and monitoring
func NewAuthor(root Author, l *zap.Logger) Author {
	return loggerAuthor{
		log:  l,
		next: root,
	}
}

type loggerAuthor struct {
	log  *zap.Logger
	next Author
}

func (l loggerAuthor) Save(ctx context.Context, author aggregate.Author) (err error) {
	defer func() {
		if err != nil {
			fields := append(l.marshalFields(author), zap.Error(err))
			l.log.Error("failed to write author into database", fields...)
			return
		}
		l.log.Info("successfully wrote author into database", l.marshalFields(author)...)
	}()

	err = l.next.Save(ctx, author)
	return
}

func (l loggerAuthor) Get(ctx context.Context, id valueobject.AuthorID) (author *aggregate.Author, err error) {
	defer func() {
		if err != nil {
			l.log.Error("failed to fetch author from database", zap.String("author_id", id.Value()),
				zap.Error(err))
			return
		}
		l.log.Info("successfully fetched author from database",
			zap.String("author_id", author.ID.Value()),
			zap.String("display_name", author.DisplayName.Value()),
		)
	}()

	author, err = l.next.Get(ctx, id)
	return
}

func (l loggerAuthor) Search(ctx context.Context, criteria Criteria) (authors []*aggregate.Author,
	nextToken string, err error) {
	defer func() {
		fields := []zap.Field{
			zap.Int("criteria_limit", criteria.Limit),
			zap.String("criteria_next_page", criteria.NextPage),
			zap.String("next_page_token", nextToken),
			zap.Int("total_items", len(authors)),
		}
		if err != nil {
			fields = append(fields, zap.Error(err))
			l.log.Error("failed to fetch authors from database", fields...)
			return
		}
		l.log.Info("successfully fetched authors from database",
			fields...)
	}()

	authors, nextToken, err = l.next.Search(ctx, criteria)
	return
}

func (l loggerAuthor) marshalFields(author aggregate.Author) []zap.Field {
	return []zap.Field{
		zap.String("author_id", author.ID.Value()),
		zap.String("display_name", author.DisplayName.Value()),
		zap.String("create_by", author.CreateBy.Value()),
		zap.String("image", author.Image.Value()),
		zap.Time("create_time", author.Metadata.CreateTime),
		zap.Time("update_time", author.Metadata.UpdateTime),
		zap.Bool("state", author.Metadata.State),
		zap.Bool("mark_as_mutation", author.Metadata.MarkAsMutation),
		zap.Bool("mark_as_removal", author.Metadata.MarkAsRemoval),
	}
}
