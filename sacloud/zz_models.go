// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-models'; DO NOT EDIT

package sacloud

import (
	"time"

	"github.com/sacloud/libsacloud-v2/pkg/mapconv"
	"github.com/sacloud/libsacloud-v2/sacloud/enums"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

/*************************************************
* Note
*************************************************/

// Note represents API parameter/response structure
type Note struct {
	ID           int64
	Name         string
	Description  string
	Tags         []string
	Availability enums.EAvailability
	Scope        enums.EScope
	Class        string
	Content      string
	Icon         *naked.Icon
	CreatedAt    *time.Time
	ModifiedAt   *time.Time
}

// GetID returns value of ID
func (o *Note) GetID() int64 {
	return o.ID
}

// SetID sets value to ID
func (o *Note) SetID(v int64) {
	o.ID = v
}

// GetName returns value of Name
func (o *Note) GetName() string {
	return o.Name
}

// SetName sets value to Name
func (o *Note) SetName(v string) {
	o.Name = v
}

// GetDescription returns value of Description
func (o *Note) GetDescription() string {
	return o.Description
}

// SetDescription sets value to Description
func (o *Note) SetDescription(v string) {
	o.Description = v
}

// GetTags returns value of Tags
func (o *Note) GetTags() []string {
	return o.Tags
}

// SetTags sets value to Tags
func (o *Note) SetTags(v []string) {
	o.Tags = v
}

// GetAvailability returns value of Availability
func (o *Note) GetAvailability() enums.EAvailability {
	return o.Availability
}

// SetAvailability sets value to Availability
func (o *Note) SetAvailability(v enums.EAvailability) {
	o.Availability = v
}

// GetScope returns value of Scope
func (o *Note) GetScope() enums.EScope {
	return o.Scope
}

// SetScope sets value to Scope
func (o *Note) SetScope(v enums.EScope) {
	o.Scope = v
}

// GetClass returns value of Class
func (o *Note) GetClass() string {
	return o.Class
}

// SetClass sets value to Class
func (o *Note) SetClass(v string) {
	o.Class = v
}

// GetContent returns value of Content
func (o *Note) GetContent() string {
	return o.Content
}

// SetContent sets value to Content
func (o *Note) SetContent(v string) {
	o.Content = v
}

// GetIcon returns value of Icon
func (o *Note) GetIcon() *naked.Icon {
	return o.Icon
}

// SetIcon sets value to Icon
func (o *Note) SetIcon(v *naked.Icon) {
	o.Icon = v
}

// GetCreatedAt returns value of CreatedAt
func (o *Note) GetCreatedAt() *time.Time {
	return o.CreatedAt
}

// SetCreatedAt sets value to CreatedAt
func (o *Note) SetCreatedAt(v *time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns value of ModifiedAt
func (o *Note) GetModifiedAt() *time.Time {
	return o.ModifiedAt
}

// SetModifiedAt sets value to ModifiedAt
func (o *Note) SetModifiedAt(v *time.Time) {
	o.ModifiedAt = v
}

// ToNaked returns naked Note
func (o *Note) ToNaked() (*naked.Note, error) {
	dest := &naked.Note{}
	err := mapconv.ToNaked(o, dest)
	return dest, err
}

// ParseNaked parse values from naked Note
func (o *Note) ParseNaked(naked *naked.Note) error {
	return mapconv.FromNaked(naked, o)
}

/*************************************************
* NoteFindRequest
*************************************************/

// NoteFindRequest represents API parameter/response structure
type NoteFindRequest struct {
	Count   int
	From    int
	Sort    []string
	Filter  map[string]interface{}
	Include []string
	Exclude []string
}

// GetCount returns value of Count
func (o *NoteFindRequest) GetCount() int {
	return o.Count
}

// SetCount sets value to Count
func (o *NoteFindRequest) SetCount(v int) {
	o.Count = v
}

// GetFrom returns value of From
func (o *NoteFindRequest) GetFrom() int {
	return o.From
}

// SetFrom sets value to From
func (o *NoteFindRequest) SetFrom(v int) {
	o.From = v
}

// GetSort returns value of Sort
func (o *NoteFindRequest) GetSort() []string {
	return o.Sort
}

// SetSort sets value to Sort
func (o *NoteFindRequest) SetSort(v []string) {
	o.Sort = v
}

// GetFilter returns value of Filter
func (o *NoteFindRequest) GetFilter() map[string]interface{} {
	return o.Filter
}

// SetFilter sets value to Filter
func (o *NoteFindRequest) SetFilter(v map[string]interface{}) {
	o.Filter = v
}

// GetInclude returns value of Include
func (o *NoteFindRequest) GetInclude() []string {
	return o.Include
}

// SetInclude sets value to Include
func (o *NoteFindRequest) SetInclude(v []string) {
	o.Include = v
}

// GetExclude returns value of Exclude
func (o *NoteFindRequest) GetExclude() []string {
	return o.Exclude
}

// SetExclude sets value to Exclude
func (o *NoteFindRequest) SetExclude(v []string) {
	o.Exclude = v
}

/*************************************************
* NoteCreateRequest
*************************************************/

// NoteCreateRequest represents API parameter/response structure
type NoteCreateRequest struct {
	Name    string
	Tags    []string
	IconID  int64 `mapconv:"Icon.ID"`
	Class   string
	Content string
}

// GetName returns value of Name
func (o *NoteCreateRequest) GetName() string {
	return o.Name
}

// SetName sets value to Name
func (o *NoteCreateRequest) SetName(v string) {
	o.Name = v
}

// GetTags returns value of Tags
func (o *NoteCreateRequest) GetTags() []string {
	return o.Tags
}

// SetTags sets value to Tags
func (o *NoteCreateRequest) SetTags(v []string) {
	o.Tags = v
}

// GetIconID returns value of IconID
func (o *NoteCreateRequest) GetIconID() int64 {
	return o.IconID
}

// SetIconID sets value to IconID
func (o *NoteCreateRequest) SetIconID(v int64) {
	o.IconID = v
}

// GetClass returns value of Class
func (o *NoteCreateRequest) GetClass() string {
	return o.Class
}

// SetClass sets value to Class
func (o *NoteCreateRequest) SetClass(v string) {
	o.Class = v
}

// GetContent returns value of Content
func (o *NoteCreateRequest) GetContent() string {
	return o.Content
}

// SetContent sets value to Content
func (o *NoteCreateRequest) SetContent(v string) {
	o.Content = v
}

// ToNaked returns naked NoteCreateRequest
func (o *NoteCreateRequest) ToNaked() (*naked.Note, error) {
	dest := &naked.Note{}
	err := mapconv.ToNaked(o, dest)
	return dest, err
}

// ParseNaked parse values from naked NoteCreateRequest
func (o *NoteCreateRequest) ParseNaked(naked *naked.Note) error {
	return mapconv.FromNaked(naked, o)
}

/*************************************************
* NoteUpdateRequest
*************************************************/

// NoteUpdateRequest represents API parameter/response structure
type NoteUpdateRequest struct {
	Name    string
	Tags    []string
	IconID  int64 `mapconv:"Icon.ID"`
	Class   string
	Content string
}

// GetName returns value of Name
func (o *NoteUpdateRequest) GetName() string {
	return o.Name
}

// SetName sets value to Name
func (o *NoteUpdateRequest) SetName(v string) {
	o.Name = v
}

// GetTags returns value of Tags
func (o *NoteUpdateRequest) GetTags() []string {
	return o.Tags
}

// SetTags sets value to Tags
func (o *NoteUpdateRequest) SetTags(v []string) {
	o.Tags = v
}

// GetIconID returns value of IconID
func (o *NoteUpdateRequest) GetIconID() int64 {
	return o.IconID
}

// SetIconID sets value to IconID
func (o *NoteUpdateRequest) SetIconID(v int64) {
	o.IconID = v
}

// GetClass returns value of Class
func (o *NoteUpdateRequest) GetClass() string {
	return o.Class
}

// SetClass sets value to Class
func (o *NoteUpdateRequest) SetClass(v string) {
	o.Class = v
}

// GetContent returns value of Content
func (o *NoteUpdateRequest) GetContent() string {
	return o.Content
}

// SetContent sets value to Content
func (o *NoteUpdateRequest) SetContent(v string) {
	o.Content = v
}

// ToNaked returns naked NoteUpdateRequest
func (o *NoteUpdateRequest) ToNaked() (*naked.Note, error) {
	dest := &naked.Note{}
	err := mapconv.ToNaked(o, dest)
	return dest, err
}

// ParseNaked parse values from naked NoteUpdateRequest
func (o *NoteUpdateRequest) ParseNaked(naked *naked.Note) error {
	return mapconv.FromNaked(naked, o)
}
