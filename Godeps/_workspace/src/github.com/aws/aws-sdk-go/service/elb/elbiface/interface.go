// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package elbiface provides an interface for the Elastic Load Balancing.
package elbiface

import (
	"bosun.org/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/request"
	"bosun.org/Godeps/_workspace/src/github.com/aws/aws-sdk-go/service/elb"
)

// ELBAPI is the interface type for elb.ELB.
type ELBAPI interface {
	AddTagsRequest(*elb.AddTagsInput) (*request.Request, *elb.AddTagsOutput)

	AddTags(*elb.AddTagsInput) (*elb.AddTagsOutput, error)

	ApplySecurityGroupsToLoadBalancerRequest(*elb.ApplySecurityGroupsToLoadBalancerInput) (*request.Request, *elb.ApplySecurityGroupsToLoadBalancerOutput)

	ApplySecurityGroupsToLoadBalancer(*elb.ApplySecurityGroupsToLoadBalancerInput) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error)

	AttachLoadBalancerToSubnetsRequest(*elb.AttachLoadBalancerToSubnetsInput) (*request.Request, *elb.AttachLoadBalancerToSubnetsOutput)

	AttachLoadBalancerToSubnets(*elb.AttachLoadBalancerToSubnetsInput) (*elb.AttachLoadBalancerToSubnetsOutput, error)

	ConfigureHealthCheckRequest(*elb.ConfigureHealthCheckInput) (*request.Request, *elb.ConfigureHealthCheckOutput)

	ConfigureHealthCheck(*elb.ConfigureHealthCheckInput) (*elb.ConfigureHealthCheckOutput, error)

	CreateAppCookieStickinessPolicyRequest(*elb.CreateAppCookieStickinessPolicyInput) (*request.Request, *elb.CreateAppCookieStickinessPolicyOutput)

	CreateAppCookieStickinessPolicy(*elb.CreateAppCookieStickinessPolicyInput) (*elb.CreateAppCookieStickinessPolicyOutput, error)

	CreateLBCookieStickinessPolicyRequest(*elb.CreateLBCookieStickinessPolicyInput) (*request.Request, *elb.CreateLBCookieStickinessPolicyOutput)

	CreateLBCookieStickinessPolicy(*elb.CreateLBCookieStickinessPolicyInput) (*elb.CreateLBCookieStickinessPolicyOutput, error)

	CreateLoadBalancerRequest(*elb.CreateLoadBalancerInput) (*request.Request, *elb.CreateLoadBalancerOutput)

	CreateLoadBalancer(*elb.CreateLoadBalancerInput) (*elb.CreateLoadBalancerOutput, error)

	CreateLoadBalancerListenersRequest(*elb.CreateLoadBalancerListenersInput) (*request.Request, *elb.CreateLoadBalancerListenersOutput)

	CreateLoadBalancerListeners(*elb.CreateLoadBalancerListenersInput) (*elb.CreateLoadBalancerListenersOutput, error)

	CreateLoadBalancerPolicyRequest(*elb.CreateLoadBalancerPolicyInput) (*request.Request, *elb.CreateLoadBalancerPolicyOutput)

	CreateLoadBalancerPolicy(*elb.CreateLoadBalancerPolicyInput) (*elb.CreateLoadBalancerPolicyOutput, error)

	DeleteLoadBalancerRequest(*elb.DeleteLoadBalancerInput) (*request.Request, *elb.DeleteLoadBalancerOutput)

	DeleteLoadBalancer(*elb.DeleteLoadBalancerInput) (*elb.DeleteLoadBalancerOutput, error)

	DeleteLoadBalancerListenersRequest(*elb.DeleteLoadBalancerListenersInput) (*request.Request, *elb.DeleteLoadBalancerListenersOutput)

	DeleteLoadBalancerListeners(*elb.DeleteLoadBalancerListenersInput) (*elb.DeleteLoadBalancerListenersOutput, error)

	DeleteLoadBalancerPolicyRequest(*elb.DeleteLoadBalancerPolicyInput) (*request.Request, *elb.DeleteLoadBalancerPolicyOutput)

	DeleteLoadBalancerPolicy(*elb.DeleteLoadBalancerPolicyInput) (*elb.DeleteLoadBalancerPolicyOutput, error)

	DeregisterInstancesFromLoadBalancerRequest(*elb.DeregisterInstancesFromLoadBalancerInput) (*request.Request, *elb.DeregisterInstancesFromLoadBalancerOutput)

	DeregisterInstancesFromLoadBalancer(*elb.DeregisterInstancesFromLoadBalancerInput) (*elb.DeregisterInstancesFromLoadBalancerOutput, error)

	DescribeInstanceHealthRequest(*elb.DescribeInstanceHealthInput) (*request.Request, *elb.DescribeInstanceHealthOutput)

	DescribeInstanceHealth(*elb.DescribeInstanceHealthInput) (*elb.DescribeInstanceHealthOutput, error)

	DescribeLoadBalancerAttributesRequest(*elb.DescribeLoadBalancerAttributesInput) (*request.Request, *elb.DescribeLoadBalancerAttributesOutput)

	DescribeLoadBalancerAttributes(*elb.DescribeLoadBalancerAttributesInput) (*elb.DescribeLoadBalancerAttributesOutput, error)

	DescribeLoadBalancerPoliciesRequest(*elb.DescribeLoadBalancerPoliciesInput) (*request.Request, *elb.DescribeLoadBalancerPoliciesOutput)

	DescribeLoadBalancerPolicies(*elb.DescribeLoadBalancerPoliciesInput) (*elb.DescribeLoadBalancerPoliciesOutput, error)

	DescribeLoadBalancerPolicyTypesRequest(*elb.DescribeLoadBalancerPolicyTypesInput) (*request.Request, *elb.DescribeLoadBalancerPolicyTypesOutput)

	DescribeLoadBalancerPolicyTypes(*elb.DescribeLoadBalancerPolicyTypesInput) (*elb.DescribeLoadBalancerPolicyTypesOutput, error)

	DescribeLoadBalancersRequest(*elb.DescribeLoadBalancersInput) (*request.Request, *elb.DescribeLoadBalancersOutput)

	DescribeLoadBalancers(*elb.DescribeLoadBalancersInput) (*elb.DescribeLoadBalancersOutput, error)

	DescribeLoadBalancersPages(*elb.DescribeLoadBalancersInput, func(*elb.DescribeLoadBalancersOutput, bool) bool) error

	DescribeTagsRequest(*elb.DescribeTagsInput) (*request.Request, *elb.DescribeTagsOutput)

	DescribeTags(*elb.DescribeTagsInput) (*elb.DescribeTagsOutput, error)

	DetachLoadBalancerFromSubnetsRequest(*elb.DetachLoadBalancerFromSubnetsInput) (*request.Request, *elb.DetachLoadBalancerFromSubnetsOutput)

	DetachLoadBalancerFromSubnets(*elb.DetachLoadBalancerFromSubnetsInput) (*elb.DetachLoadBalancerFromSubnetsOutput, error)

	DisableAvailabilityZonesForLoadBalancerRequest(*elb.DisableAvailabilityZonesForLoadBalancerInput) (*request.Request, *elb.DisableAvailabilityZonesForLoadBalancerOutput)

	DisableAvailabilityZonesForLoadBalancer(*elb.DisableAvailabilityZonesForLoadBalancerInput) (*elb.DisableAvailabilityZonesForLoadBalancerOutput, error)

	EnableAvailabilityZonesForLoadBalancerRequest(*elb.EnableAvailabilityZonesForLoadBalancerInput) (*request.Request, *elb.EnableAvailabilityZonesForLoadBalancerOutput)

	EnableAvailabilityZonesForLoadBalancer(*elb.EnableAvailabilityZonesForLoadBalancerInput) (*elb.EnableAvailabilityZonesForLoadBalancerOutput, error)

	ModifyLoadBalancerAttributesRequest(*elb.ModifyLoadBalancerAttributesInput) (*request.Request, *elb.ModifyLoadBalancerAttributesOutput)

	ModifyLoadBalancerAttributes(*elb.ModifyLoadBalancerAttributesInput) (*elb.ModifyLoadBalancerAttributesOutput, error)

	RegisterInstancesWithLoadBalancerRequest(*elb.RegisterInstancesWithLoadBalancerInput) (*request.Request, *elb.RegisterInstancesWithLoadBalancerOutput)

	RegisterInstancesWithLoadBalancer(*elb.RegisterInstancesWithLoadBalancerInput) (*elb.RegisterInstancesWithLoadBalancerOutput, error)

	RemoveTagsRequest(*elb.RemoveTagsInput) (*request.Request, *elb.RemoveTagsOutput)

	RemoveTags(*elb.RemoveTagsInput) (*elb.RemoveTagsOutput, error)

	SetLoadBalancerListenerSSLCertificateRequest(*elb.SetLoadBalancerListenerSSLCertificateInput) (*request.Request, *elb.SetLoadBalancerListenerSSLCertificateOutput)

	SetLoadBalancerListenerSSLCertificate(*elb.SetLoadBalancerListenerSSLCertificateInput) (*elb.SetLoadBalancerListenerSSLCertificateOutput, error)

	SetLoadBalancerPoliciesForBackendServerRequest(*elb.SetLoadBalancerPoliciesForBackendServerInput) (*request.Request, *elb.SetLoadBalancerPoliciesForBackendServerOutput)

	SetLoadBalancerPoliciesForBackendServer(*elb.SetLoadBalancerPoliciesForBackendServerInput) (*elb.SetLoadBalancerPoliciesForBackendServerOutput, error)

	SetLoadBalancerPoliciesOfListenerRequest(*elb.SetLoadBalancerPoliciesOfListenerInput) (*request.Request, *elb.SetLoadBalancerPoliciesOfListenerOutput)

	SetLoadBalancerPoliciesOfListener(*elb.SetLoadBalancerPoliciesOfListenerInput) (*elb.SetLoadBalancerPoliciesOfListenerOutput, error)
}
