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
import (
	"time"
)
// DagDetailAllOf struct for DagDetailAllOf
type DagDetailAllOf struct {
	Timezone string `json:"timezone,omitempty"`
	Catchup bool `json:"catchup,omitempty"`
	Orientation string `json:"orientation,omitempty"`
	Concurrency float32 `json:"concurrency,omitempty"`
	StartDate time.Time `json:"start_date,omitempty"`
	DagRunTimeout TimeDelta `json:"dag_run_timeout,omitempty"`
	DocMd string `json:"doc_md,omitempty"`
	DefaultView string `json:"default_view,omitempty"`
}
