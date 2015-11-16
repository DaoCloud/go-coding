package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*
ProjectTopicDTO project topic d t o

swagger:model ProjectTopicDTO
*/
type ProjectTopicDTO struct {

	/* ChildCount child count
	 */
	ChildCount int32 `json:"child_count,omitempty"`

	/* Content content
	 */
	Content string `json:"content,omitempty"`

	/* CreatedAt created at
	 */
	CreatedAt Timestamp `json:"created_at,omitempty"`

	/* CurrentUserRole current user role
	 */
	CurrentUserRole string `json:"current_user_role,omitempty"`

	/* CurrentUserRoleID current user role id
	 */
	CurrentUserRoleID int32 `json:"current_user_role_id,omitempty"`

	/* ID id
	 */
	ID int32 `json:"id,omitempty"`

	/* Labels labels
	 */
	Labels []ProjectLabelDTO `json:"labels,omitempty"`

	/* Number number
	 */
	Number int32 `json:"number,omitempty"`

	/* OwnerID owner id
	 */
	OwnerID int32 `json:"owner_id,omitempty"`

	/* ParentID parent id
	 */
	ParentID int32 `json:"parent_id,omitempty"`

	/* ProjectID project id
	 */
	ProjectID int32 `json:"project_id,omitempty"`

	/* Title title
	 */
	Title string `json:"title,omitempty"`

	/* Type type
	 */
	Type int32 `json:"type,omitempty"`

	/* UpdatedAt updated at
	 */
	UpdatedAt Timestamp `json:"updated_at,omitempty"`

	/* Watching watching
	 */
	Watching bool `json:"watching,omitempty"`
}
