// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-envelope'; DO NOT EDIT

package sacloud

import (
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

// NoteCreateRequestEnvelope is envelop of API request
type NoteCreateRequestEnvelope struct {
	Note *naked.Note `json:",omitempty"`
}

// NoteCreateResponseEnvelope is envelop of API response
type NoteCreateResponseEnvelope struct {
	IsOk    bool        `json:"is_ok,omitempty"` // is_ok項目
	Success bool        `json:",omitempty"`      // success項目
	Note    *naked.Note `json:",omitempty"`
}

// NoteReadResponseEnvelope is envelop of API response
type NoteReadResponseEnvelope struct {
	IsOk    bool        `json:"is_ok,omitempty"` // is_ok項目
	Success bool        `json:",omitempty"`      // success項目
	Note    *naked.Note `json:",omitempty"`
}

// NoteUpdateRequestEnvelope is envelop of API request
type NoteUpdateRequestEnvelope struct {
	Note *naked.Note `json:",omitempty"`
}

// NoteUpdateResponseEnvelope is envelop of API response
type NoteUpdateResponseEnvelope struct {
	IsOk    bool        `json:"is_ok,omitempty"` // is_ok項目
	Success bool        `json:",omitempty"`      // success項目
	Note    *naked.Note `json:",omitempty"`
}