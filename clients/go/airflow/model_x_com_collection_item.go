/*
 * Airflow API (Stable)
 *
 * Apache Airflow management API.
 *
 * API version: 1.0.0
 * Contact: dev@airflow.apache.org
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package airflow
// XComCollectionItem struct for XComCollectionItem
type XComCollectionItem struct {
	Key string `json:"key,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	ExecutionDate string `json:"execution_date,omitempty"`
	TaskId string `json:"task_id,omitempty"`
	DagId string `json:"dag_id,omitempty"`
}
