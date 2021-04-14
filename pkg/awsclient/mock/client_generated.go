// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock is a generated GoMock package.
package mock

import (
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	elbv2 "github.com/aws/aws-sdk-go/service/elbv2"
	resourcegroupstaggingapi "github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	route53 "github.com/aws/aws-sdk-go/service/route53"
	s3iface "github.com/aws/aws-sdk-go/service/s3/s3iface"
	s3manager "github.com/aws/aws-sdk-go/service/s3/s3manager"
	sts "github.com/aws/aws-sdk-go/service/sts"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// DescribeAvailabilityZones mocks base method
func (m *MockClient) DescribeAvailabilityZones(arg0 *ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeAvailabilityZones", arg0)
	ret0, _ := ret[0].(*ec2.DescribeAvailabilityZonesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAvailabilityZones indicates an expected call of DescribeAvailabilityZones
func (mr *MockClientMockRecorder) DescribeAvailabilityZones(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAvailabilityZones", reflect.TypeOf((*MockClient)(nil).DescribeAvailabilityZones), arg0)
}

// DescribeSubnets mocks base method
func (m *MockClient) DescribeSubnets(arg0 *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSubnets", arg0)
	ret0, _ := ret[0].(*ec2.DescribeSubnetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSubnets indicates an expected call of DescribeSubnets
func (mr *MockClientMockRecorder) DescribeSubnets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSubnets", reflect.TypeOf((*MockClient)(nil).DescribeSubnets), arg0)
}

// DescribeRouteTables mocks base method
func (m *MockClient) DescribeRouteTables(arg0 *ec2.DescribeRouteTablesInput) (*ec2.DescribeRouteTablesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeRouteTables", arg0)
	ret0, _ := ret[0].(*ec2.DescribeRouteTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRouteTables indicates an expected call of DescribeRouteTables
func (mr *MockClientMockRecorder) DescribeRouteTables(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRouteTables", reflect.TypeOf((*MockClient)(nil).DescribeRouteTables), arg0)
}

// DescribeInstances mocks base method
func (m *MockClient) DescribeInstances(arg0 *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeInstances", arg0)
	ret0, _ := ret[0].(*ec2.DescribeInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstances indicates an expected call of DescribeInstances
func (mr *MockClientMockRecorder) DescribeInstances(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstances", reflect.TypeOf((*MockClient)(nil).DescribeInstances), arg0)
}

// StopInstances mocks base method
func (m *MockClient) StopInstances(arg0 *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopInstances", arg0)
	ret0, _ := ret[0].(*ec2.StopInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopInstances indicates an expected call of StopInstances
func (mr *MockClientMockRecorder) StopInstances(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopInstances", reflect.TypeOf((*MockClient)(nil).StopInstances), arg0)
}

// StartInstances mocks base method
func (m *MockClient) StartInstances(arg0 *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartInstances", arg0)
	ret0, _ := ret[0].(*ec2.StartInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartInstances indicates an expected call of StartInstances
func (mr *MockClientMockRecorder) StartInstances(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartInstances", reflect.TypeOf((*MockClient)(nil).StartInstances), arg0)
}

// CreateVpcEndpointServiceConfiguration mocks base method
func (m *MockClient) CreateVpcEndpointServiceConfiguration(arg0 *ec2.CreateVpcEndpointServiceConfigurationInput) (*ec2.CreateVpcEndpointServiceConfigurationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVpcEndpointServiceConfiguration", arg0)
	ret0, _ := ret[0].(*ec2.CreateVpcEndpointServiceConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVpcEndpointServiceConfiguration indicates an expected call of CreateVpcEndpointServiceConfiguration
func (mr *MockClientMockRecorder) CreateVpcEndpointServiceConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVpcEndpointServiceConfiguration", reflect.TypeOf((*MockClient)(nil).CreateVpcEndpointServiceConfiguration), arg0)
}

// DescribeVpcEndpointServiceConfigurations mocks base method
func (m *MockClient) DescribeVpcEndpointServiceConfigurations(arg0 *ec2.DescribeVpcEndpointServiceConfigurationsInput) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeVpcEndpointServiceConfigurations", arg0)
	ret0, _ := ret[0].(*ec2.DescribeVpcEndpointServiceConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVpcEndpointServiceConfigurations indicates an expected call of DescribeVpcEndpointServiceConfigurations
func (mr *MockClientMockRecorder) DescribeVpcEndpointServiceConfigurations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcEndpointServiceConfigurations", reflect.TypeOf((*MockClient)(nil).DescribeVpcEndpointServiceConfigurations), arg0)
}

// ModifyVpcEndpointServiceConfiguration mocks base method
func (m *MockClient) ModifyVpcEndpointServiceConfiguration(arg0 *ec2.ModifyVpcEndpointServiceConfigurationInput) (*ec2.ModifyVpcEndpointServiceConfigurationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyVpcEndpointServiceConfiguration", arg0)
	ret0, _ := ret[0].(*ec2.ModifyVpcEndpointServiceConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyVpcEndpointServiceConfiguration indicates an expected call of ModifyVpcEndpointServiceConfiguration
func (mr *MockClientMockRecorder) ModifyVpcEndpointServiceConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyVpcEndpointServiceConfiguration", reflect.TypeOf((*MockClient)(nil).ModifyVpcEndpointServiceConfiguration), arg0)
}

// DeleteVpcEndpointServiceConfigurations mocks base method
func (m *MockClient) DeleteVpcEndpointServiceConfigurations(arg0 *ec2.DeleteVpcEndpointServiceConfigurationsInput) (*ec2.DeleteVpcEndpointServiceConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVpcEndpointServiceConfigurations", arg0)
	ret0, _ := ret[0].(*ec2.DeleteVpcEndpointServiceConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteVpcEndpointServiceConfigurations indicates an expected call of DeleteVpcEndpointServiceConfigurations
func (mr *MockClientMockRecorder) DeleteVpcEndpointServiceConfigurations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVpcEndpointServiceConfigurations", reflect.TypeOf((*MockClient)(nil).DeleteVpcEndpointServiceConfigurations), arg0)
}

// DescribeVpcEndpointServicePermissions mocks base method
func (m *MockClient) DescribeVpcEndpointServicePermissions(arg0 *ec2.DescribeVpcEndpointServicePermissionsInput) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeVpcEndpointServicePermissions", arg0)
	ret0, _ := ret[0].(*ec2.DescribeVpcEndpointServicePermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVpcEndpointServicePermissions indicates an expected call of DescribeVpcEndpointServicePermissions
func (mr *MockClientMockRecorder) DescribeVpcEndpointServicePermissions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcEndpointServicePermissions", reflect.TypeOf((*MockClient)(nil).DescribeVpcEndpointServicePermissions), arg0)
}

// ModifyVpcEndpointServicePermissions mocks base method
func (m *MockClient) ModifyVpcEndpointServicePermissions(arg0 *ec2.ModifyVpcEndpointServicePermissionsInput) (*ec2.ModifyVpcEndpointServicePermissionsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyVpcEndpointServicePermissions", arg0)
	ret0, _ := ret[0].(*ec2.ModifyVpcEndpointServicePermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyVpcEndpointServicePermissions indicates an expected call of ModifyVpcEndpointServicePermissions
func (mr *MockClientMockRecorder) ModifyVpcEndpointServicePermissions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyVpcEndpointServicePermissions", reflect.TypeOf((*MockClient)(nil).ModifyVpcEndpointServicePermissions), arg0)
}

// DescribeVpcEndpointServices mocks base method
func (m *MockClient) DescribeVpcEndpointServices(arg0 *ec2.DescribeVpcEndpointServicesInput) (*ec2.DescribeVpcEndpointServicesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeVpcEndpointServices", arg0)
	ret0, _ := ret[0].(*ec2.DescribeVpcEndpointServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVpcEndpointServices indicates an expected call of DescribeVpcEndpointServices
func (mr *MockClientMockRecorder) DescribeVpcEndpointServices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcEndpointServices", reflect.TypeOf((*MockClient)(nil).DescribeVpcEndpointServices), arg0)
}

// DescribeVpcEndpoints mocks base method
func (m *MockClient) DescribeVpcEndpoints(arg0 *ec2.DescribeVpcEndpointsInput) (*ec2.DescribeVpcEndpointsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeVpcEndpoints", arg0)
	ret0, _ := ret[0].(*ec2.DescribeVpcEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVpcEndpoints indicates an expected call of DescribeVpcEndpoints
func (mr *MockClientMockRecorder) DescribeVpcEndpoints(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcEndpoints", reflect.TypeOf((*MockClient)(nil).DescribeVpcEndpoints), arg0)
}

// CreateVpcEndpoint mocks base method
func (m *MockClient) CreateVpcEndpoint(arg0 *ec2.CreateVpcEndpointInput) (*ec2.CreateVpcEndpointOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVpcEndpoint", arg0)
	ret0, _ := ret[0].(*ec2.CreateVpcEndpointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVpcEndpoint indicates an expected call of CreateVpcEndpoint
func (mr *MockClientMockRecorder) CreateVpcEndpoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVpcEndpoint", reflect.TypeOf((*MockClient)(nil).CreateVpcEndpoint), arg0)
}

// DeleteVpcEndpoints mocks base method
func (m *MockClient) DeleteVpcEndpoints(arg0 *ec2.DeleteVpcEndpointsInput) (*ec2.DeleteVpcEndpointsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVpcEndpoints", arg0)
	ret0, _ := ret[0].(*ec2.DeleteVpcEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteVpcEndpoints indicates an expected call of DeleteVpcEndpoints
func (mr *MockClientMockRecorder) DeleteVpcEndpoints(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVpcEndpoints", reflect.TypeOf((*MockClient)(nil).DeleteVpcEndpoints), arg0)
}

// DescribeLoadBalancers mocks base method
func (m *MockClient) DescribeLoadBalancers(arg0 *elbv2.DescribeLoadBalancersInput) (*elbv2.DescribeLoadBalancersOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeLoadBalancers", arg0)
	ret0, _ := ret[0].(*elbv2.DescribeLoadBalancersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLoadBalancers indicates an expected call of DescribeLoadBalancers
func (mr *MockClientMockRecorder) DescribeLoadBalancers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancers", reflect.TypeOf((*MockClient)(nil).DescribeLoadBalancers), arg0)
}

// Upload mocks base method
func (m *MockClient) Upload(arg0 *s3manager.UploadInput) (*s3manager.UploadOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", arg0)
	ret0, _ := ret[0].(*s3manager.UploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upload indicates an expected call of Upload
func (mr *MockClientMockRecorder) Upload(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockClient)(nil).Upload), arg0)
}

// GetS3API mocks base method
func (m *MockClient) GetS3API() s3iface.S3API {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetS3API")
	ret0, _ := ret[0].(s3iface.S3API)
	return ret0
}

// GetS3API indicates an expected call of GetS3API
func (mr *MockClientMockRecorder) GetS3API() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetS3API", reflect.TypeOf((*MockClient)(nil).GetS3API))
}

// CreateHostedZone mocks base method
func (m *MockClient) CreateHostedZone(input *route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHostedZone", input)
	ret0, _ := ret[0].(*route53.CreateHostedZoneOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHostedZone indicates an expected call of CreateHostedZone
func (mr *MockClientMockRecorder) CreateHostedZone(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHostedZone", reflect.TypeOf((*MockClient)(nil).CreateHostedZone), input)
}

// GetHostedZone mocks base method
func (m *MockClient) GetHostedZone(arg0 *route53.GetHostedZoneInput) (*route53.GetHostedZoneOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostedZone", arg0)
	ret0, _ := ret[0].(*route53.GetHostedZoneOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostedZone indicates an expected call of GetHostedZone
func (mr *MockClientMockRecorder) GetHostedZone(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostedZone", reflect.TypeOf((*MockClient)(nil).GetHostedZone), arg0)
}

// ListTagsForResource mocks base method
func (m *MockClient) ListTagsForResource(arg0 *route53.ListTagsForResourceInput) (*route53.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResource", arg0)
	ret0, _ := ret[0].(*route53.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource
func (mr *MockClientMockRecorder) ListTagsForResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockClient)(nil).ListTagsForResource), arg0)
}

// ChangeTagsForResource mocks base method
func (m *MockClient) ChangeTagsForResource(input *route53.ChangeTagsForResourceInput) (*route53.ChangeTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeTagsForResource", input)
	ret0, _ := ret[0].(*route53.ChangeTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeTagsForResource indicates an expected call of ChangeTagsForResource
func (mr *MockClientMockRecorder) ChangeTagsForResource(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeTagsForResource", reflect.TypeOf((*MockClient)(nil).ChangeTagsForResource), input)
}

// DeleteHostedZone mocks base method
func (m *MockClient) DeleteHostedZone(input *route53.DeleteHostedZoneInput) (*route53.DeleteHostedZoneOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHostedZone", input)
	ret0, _ := ret[0].(*route53.DeleteHostedZoneOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteHostedZone indicates an expected call of DeleteHostedZone
func (mr *MockClientMockRecorder) DeleteHostedZone(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHostedZone", reflect.TypeOf((*MockClient)(nil).DeleteHostedZone), input)
}

// ListResourceRecordSets mocks base method
func (m *MockClient) ListResourceRecordSets(input *route53.ListResourceRecordSetsInput) (*route53.ListResourceRecordSetsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResourceRecordSets", input)
	ret0, _ := ret[0].(*route53.ListResourceRecordSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourceRecordSets indicates an expected call of ListResourceRecordSets
func (mr *MockClientMockRecorder) ListResourceRecordSets(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceRecordSets", reflect.TypeOf((*MockClient)(nil).ListResourceRecordSets), input)
}

// ListHostedZonesByName mocks base method
func (m *MockClient) ListHostedZonesByName(input *route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHostedZonesByName", input)
	ret0, _ := ret[0].(*route53.ListHostedZonesByNameOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHostedZonesByName indicates an expected call of ListHostedZonesByName
func (mr *MockClientMockRecorder) ListHostedZonesByName(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHostedZonesByName", reflect.TypeOf((*MockClient)(nil).ListHostedZonesByName), input)
}

// ListHostedZonesByVPC mocks base method
func (m *MockClient) ListHostedZonesByVPC(input *route53.ListHostedZonesByVPCInput) (*route53.ListHostedZonesByVPCOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHostedZonesByVPC", input)
	ret0, _ := ret[0].(*route53.ListHostedZonesByVPCOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHostedZonesByVPC indicates an expected call of ListHostedZonesByVPC
func (mr *MockClientMockRecorder) ListHostedZonesByVPC(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHostedZonesByVPC", reflect.TypeOf((*MockClient)(nil).ListHostedZonesByVPC), input)
}

// ChangeResourceRecordSets mocks base method
func (m *MockClient) ChangeResourceRecordSets(arg0 *route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeResourceRecordSets", arg0)
	ret0, _ := ret[0].(*route53.ChangeResourceRecordSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeResourceRecordSets indicates an expected call of ChangeResourceRecordSets
func (mr *MockClientMockRecorder) ChangeResourceRecordSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeResourceRecordSets", reflect.TypeOf((*MockClient)(nil).ChangeResourceRecordSets), arg0)
}

// CreateVPCAssociationAuthorization mocks base method
func (m *MockClient) CreateVPCAssociationAuthorization(arg0 *route53.CreateVPCAssociationAuthorizationInput) (*route53.CreateVPCAssociationAuthorizationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVPCAssociationAuthorization", arg0)
	ret0, _ := ret[0].(*route53.CreateVPCAssociationAuthorizationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVPCAssociationAuthorization indicates an expected call of CreateVPCAssociationAuthorization
func (mr *MockClientMockRecorder) CreateVPCAssociationAuthorization(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVPCAssociationAuthorization", reflect.TypeOf((*MockClient)(nil).CreateVPCAssociationAuthorization), arg0)
}

// DeleteVPCAssociationAuthorization mocks base method
func (m *MockClient) DeleteVPCAssociationAuthorization(arg0 *route53.DeleteVPCAssociationAuthorizationInput) (*route53.DeleteVPCAssociationAuthorizationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVPCAssociationAuthorization", arg0)
	ret0, _ := ret[0].(*route53.DeleteVPCAssociationAuthorizationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteVPCAssociationAuthorization indicates an expected call of DeleteVPCAssociationAuthorization
func (mr *MockClientMockRecorder) DeleteVPCAssociationAuthorization(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVPCAssociationAuthorization", reflect.TypeOf((*MockClient)(nil).DeleteVPCAssociationAuthorization), arg0)
}

// AssociateVPCWithHostedZone mocks base method
func (m *MockClient) AssociateVPCWithHostedZone(arg0 *route53.AssociateVPCWithHostedZoneInput) (*route53.AssociateVPCWithHostedZoneOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssociateVPCWithHostedZone", arg0)
	ret0, _ := ret[0].(*route53.AssociateVPCWithHostedZoneOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssociateVPCWithHostedZone indicates an expected call of AssociateVPCWithHostedZone
func (mr *MockClientMockRecorder) AssociateVPCWithHostedZone(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssociateVPCWithHostedZone", reflect.TypeOf((*MockClient)(nil).AssociateVPCWithHostedZone), arg0)
}

// DisassociateVPCFromHostedZone mocks base method
func (m *MockClient) DisassociateVPCFromHostedZone(input *route53.DisassociateVPCFromHostedZoneInput) (*route53.DisassociateVPCFromHostedZoneOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisassociateVPCFromHostedZone", input)
	ret0, _ := ret[0].(*route53.DisassociateVPCFromHostedZoneOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisassociateVPCFromHostedZone indicates an expected call of DisassociateVPCFromHostedZone
func (mr *MockClientMockRecorder) DisassociateVPCFromHostedZone(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisassociateVPCFromHostedZone", reflect.TypeOf((*MockClient)(nil).DisassociateVPCFromHostedZone), input)
}

// GetResourcesPages mocks base method
func (m *MockClient) GetResourcesPages(input *resourcegroupstaggingapi.GetResourcesInput, fn func(*resourcegroupstaggingapi.GetResourcesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourcesPages", input, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetResourcesPages indicates an expected call of GetResourcesPages
func (mr *MockClientMockRecorder) GetResourcesPages(input, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourcesPages", reflect.TypeOf((*MockClient)(nil).GetResourcesPages), input, fn)
}

// GetCallerIdentity mocks base method
func (m *MockClient) GetCallerIdentity(input *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCallerIdentity", input)
	ret0, _ := ret[0].(*sts.GetCallerIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCallerIdentity indicates an expected call of GetCallerIdentity
func (mr *MockClientMockRecorder) GetCallerIdentity(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCallerIdentity", reflect.TypeOf((*MockClient)(nil).GetCallerIdentity), input)
}
