/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// StatefulSetUpdateStrategy indicates the strategy that the StatefulSet controller will use to perform updates. It includes any additional parameters necessary to perform the update for the indicated strategy.
type V1beta1StatefulSetUpdateStrategy struct {

	// RollingUpdate is used to communicate parameters when Type is RollingUpdateStatefulSetStrategyType.
	RollingUpdate *V1beta1RollingUpdateStatefulSetStrategy `json:"rollingUpdate,omitempty"`

	// Type indicates the type of the StatefulSetUpdateStrategy.
	Type_ string `json:"type,omitempty"`
}
