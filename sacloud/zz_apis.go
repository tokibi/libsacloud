// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-interfaces'; DO NOT EDIT

package sacloud

import (
	"context"
)

// NoteAPI is interface for operate Note resource
type NoteAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*Note, error)
	Create(ctx context.Context, zone string, param *NoteCreateRequest) (*Note, error)
	Read(ctx context.Context, zone string, id int64) (*Note, error)
	Update(ctx context.Context, zone string, id int64, param *NoteUpdateRequest) (*Note, error)
	Delete(ctx context.Context, zone string, id int64) error
}

// ZoneAPI is interface for operate Zone resource
type ZoneAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*Zone, error)
	Read(ctx context.Context, zone string, id int64) (*Zone, error)
}
