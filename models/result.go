package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*
Result result

swagger:model Result
*/
type Result struct {

	/* Code code
	 */
	Code int32 `json:"code,omitempty"`

	/* Data data
	 */
	Data map[string]interface{} `json:"data,omitempty"`

	/* Msg msg
	 */
	Msg map[string]interface{} `json:"msg,omitempty"`
}
