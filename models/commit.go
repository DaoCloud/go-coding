package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*
Commit commit

swagger:model Commit
*/
type Commit struct {

	/* Author author
	 */
	Author PersonIdent `json:"author,omitempty"`

	/* AuthorDate author date
	 */
	AuthorDate strfmt.DateTime `json:"authorDate,omitempty"`

	/* CommitDate commit date
	 */
	CommitDate strfmt.DateTime `json:"commitDate,omitempty"`

	/* Committer committer
	 */
	Committer PersonIdent `json:"committer,omitempty"`

	/* FullMessage full message
	 */
	FullMessage string `json:"fullMessage,omitempty"`

	/* Sha sha
	 */
	Sha string `json:"sha,omitempty"`

	/* ShortMessage short message
	 */
	ShortMessage string `json:"shortMessage,omitempty"`
}
