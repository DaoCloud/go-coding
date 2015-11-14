package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*
TaskDTO task d t o

swagger:model TaskDTO
*/
type TaskDTO struct {

	/* Comments comments
	 */
	Comments int32 `json:"comments,omitempty"`

	/* Content content
	 */
	Content string `json:"content,omitempty"`

	/* CreatedAt created at
	 */
	CreatedAt Timestamp `json:"created_at,omitempty"`

	/* CreatorID creator id
	 */
	CreatorID int32 `json:"creator_id,omitempty"`

	/* CurrentUserRole current user role
	 */
	CurrentUserRole string `json:"current_user_role,omitempty"`

	/* CurrentUserRoleID current user role id
	 */
	CurrentUserRoleID int32 `json:"current_user_role_id,omitempty"`

	/* Deadline deadline
	 */
	Deadline string `json:"deadline,omitempty"`

	/* HasDescription has description
	 */
	HasDescription bool `json:"has_description,omitempty"`

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

	/* Priority priority
	 */
	Priority int32 `json:"priority,omitempty"`

	/* ProjectID project id
	 */
	ProjectID int32 `json:"project_id,omitempty"`

	/* Status status
	 */
	Status int32 `json:"status,omitempty"`

	/* UpdatedAt updated at
	 */
	UpdatedAt Timestamp `json:"updated_at,omitempty"`

	/* Watch watch
	 */
	Watch bool `json:"watch,omitempty"`
}
