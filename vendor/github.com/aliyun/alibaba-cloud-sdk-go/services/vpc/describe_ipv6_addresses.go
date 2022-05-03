package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeIpv6Addresses invokes the vpc.DescribeIpv6Addresses API synchronously
func (client *Client) DescribeIpv6Addresses(request *DescribeIpv6AddressesRequest) (response *DescribeIpv6AddressesResponse, err error) {
	response = CreateDescribeIpv6AddressesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIpv6AddressesWithChan invokes the vpc.DescribeIpv6Addresses API asynchronously
func (client *Client) DescribeIpv6AddressesWithChan(request *DescribeIpv6AddressesRequest) (<-chan *DescribeIpv6AddressesResponse, <-chan error) {
	responseChan := make(chan *DescribeIpv6AddressesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIpv6Addresses(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeIpv6AddressesWithCallback invokes the vpc.DescribeIpv6Addresses API asynchronously
func (client *Client) DescribeIpv6AddressesWithCallback(request *DescribeIpv6AddressesRequest, callback func(response *DescribeIpv6AddressesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIpv6AddressesResponse
		var err error
		defer close(result)
		response, err = client.DescribeIpv6Addresses(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeIpv6AddressesRequest is the request struct for api DescribeIpv6Addresses
type DescribeIpv6AddressesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Ipv6InternetBandwidthId string           `position:"Query" name:"Ipv6InternetBandwidthId"`
	NetworkType             string           `position:"Query" name:"NetworkType"`
	PageNumber              requests.Integer `position:"Query" name:"PageNumber"`
	AssociatedInstanceType  string           `position:"Query" name:"AssociatedInstanceType"`
	PageSize                requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount    string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string           `position:"Query" name:"OwnerAccount"`
	OwnerId                 requests.Integer `position:"Query" name:"OwnerId"`
	VSwitchId               string           `position:"Query" name:"VSwitchId"`
	Ipv6AddressId           string           `position:"Query" name:"Ipv6AddressId"`
	VpcId                   string           `position:"Query" name:"VpcId"`
	Name                    string           `position:"Query" name:"Name"`
	Ipv6Address             string           `position:"Query" name:"Ipv6Address"`
	AssociatedInstanceId    string           `position:"Query" name:"AssociatedInstanceId"`
}

// DescribeIpv6AddressesResponse is the response struct for api DescribeIpv6Addresses
type DescribeIpv6AddressesResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	TotalCount    int           `json:"TotalCount" xml:"TotalCount"`
	PageNumber    int           `json:"PageNumber" xml:"PageNumber"`
	PageSize      int           `json:"PageSize" xml:"PageSize"`
	Ipv6Addresses Ipv6Addresses `json:"Ipv6Addresses" xml:"Ipv6Addresses"`
}

// CreateDescribeIpv6AddressesRequest creates a request to invoke DescribeIpv6Addresses API
func CreateDescribeIpv6AddressesRequest() (request *DescribeIpv6AddressesRequest) {
	request = &DescribeIpv6AddressesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeIpv6Addresses", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeIpv6AddressesResponse creates a response to parse from DescribeIpv6Addresses response
func CreateDescribeIpv6AddressesResponse() (response *DescribeIpv6AddressesResponse) {
	response = &DescribeIpv6AddressesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
