/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0.0
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package logicmonitor

type RestAlertRulePagination struct {
	Total int32 `json:"total,omitempty"`

	SearchId string `json:"searchId,omitempty"`

	Items []RestAlertRule `json:"items,omitempty"`
}
