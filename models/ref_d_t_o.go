package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*
RefDTO ref d t o

swagger:model RefDTO
*/
type RefDTO struct {

	/* Author author
	 */
	Author CommitterDTO `json:"author,omitempty"`

	/* CreatedAt created at
	 */
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	/* IsDefaultBranch is default branch
	 */
	IsDefaultBranch bool `json:"is_default_branch,omitempty"`

	/* IsProtected is protected
	 */
	IsProtected bool `json:"is_protected,omitempty"`

	/* LastCommit last commit
	 */
	LastCommit CommitDTO `json:"last_commit,omitempty"`

	/* Message message
	 */
	Message string `json:"message,omitempty"`

	/* Name name
	 */
	Name string `json:"name,omitempty"`
}