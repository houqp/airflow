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
// WeightRule the model 'WeightRule'
type WeightRule string

// List of WeightRule
const (
	WEIGHTRULE_DOWNSTREAM WeightRule = "downstream"
	WEIGHTRULE_UPSTREAM WeightRule = "upstream"
	WEIGHTRULE_ABSOLUTE WeightRule = "absolute"
)
