package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*
PullRequestDTO pull request d t o

swagger:model PullRequestDTO
*/
type PullRequestDTO struct {

	/* ActionAt action at
	 */
	ActionAt strfmt.DateTime `json:"action_at,omitempty"`

	/* ActionAuthor action author
	 */
	ActionAuthor UserDTO `json:"action_author,omitempty"`

	/* Author author
	 */
	Author UserDTO `json:"author,omitempty"`

	/* CreatedAt created at
	 */
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	/* DesBranch des branch
	 */
	DesBranch string `json:"desBranch,omitempty"`

	/* ID id
	 */
	ID int32 `json:"id,omitempty"`

	/* Iid iid
	 */
	Iid int32 `json:"iid,omitempty"`

	/* MergeStatus merge status
	 */
	MergeStatus string `json:"merge_status,omitempty"`

	/* Path path
	 */
	Path string `json:"path,omitempty"`

	/* SourceDepot source depot
	 */
	SourceDepot DepotDTO `json:"source_depot,omitempty"`

	/* SrcBranch src branch
	 */
	SrcBranch string `json:"srcBranch,omitempty"`

	/* SrcExist src exist
	 */
	SrcExist bool `json:"srcExist,omitempty"`

	/* SrcOwnerName src owner name
	 */
	SrcOwnerName string `json:"src_owner_name,omitempty"`

	/* SrcProjectName src project name
	 */
	SrcProjectName string `json:"src_project_name,omitempty"`

	/* Title title
	 */
	Title string `json:"title,omitempty"`
}