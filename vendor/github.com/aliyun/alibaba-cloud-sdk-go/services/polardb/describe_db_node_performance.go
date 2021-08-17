package polardb

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

// DescribeDBNodePerformance invokes the polardb.DescribeDBNodePerformance API synchronously
func (client *Client) DescribeDBNodePerformance(request *DescribeDBNodePerformanceRequest) (response *DescribeDBNodePerformanceResponse, err error) {
	response = CreateDescribeDBNodePerformanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBNodePerformanceWithChan invokes the polardb.DescribeDBNodePerformance API asynchronously
func (client *Client) DescribeDBNodePerformanceWithChan(request *DescribeDBNodePerformanceRequest) (<-chan *DescribeDBNodePerformanceResponse, <-chan error) {
	responseChan := make(chan *DescribeDBNodePerformanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBNodePerformance(request)
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

// DescribeDBNodePerformanceWithCallback invokes the polardb.DescribeDBNodePerformance API asynchronously
func (client *Client) DescribeDBNodePerformanceWithCallback(request *DescribeDBNodePerformanceRequest, callback func(response *DescribeDBNodePerformanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBNodePerformanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBNodePerformance(request)
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

// DescribeDBNodePerformanceRequest is the request struct for api DescribeDBNodePerformance
type DescribeDBNodePerformanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBNodeId             string           `position:"Query" name:"DBNodeId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	Key                  string           `position:"Query" name:"Key"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Metric               string           `position:"Query" name:"Metric"`
}

// DescribeDBNodePerformanceResponse is the response struct for api DescribeDBNodePerformance
type DescribeDBNodePerformanceResponse struct {
	*responses.BaseResponse
	DBVersion       string                                     `json:"DBVersion" xml:"DBVersion"`
	EndTime         string                                     `json:"EndTime" xml:"EndTime"`
	RequestId       string                                     `json:"RequestId" xml:"RequestId"`
	StartTime       string                                     `json:"StartTime" xml:"StartTime"`
	DBType          string                                     `json:"DBType" xml:"DBType"`
	DBNodeId        string                                     `json:"DBNodeId" xml:"DBNodeId"`
	Engine          string                                     `json:"Engine" xml:"Engine"`
	PerformanceKeys PerformanceKeysInDescribeDBNodePerformance `json:"PerformanceKeys" xml:"PerformanceKeys"`
}

// CreateDescribeDBNodePerformanceRequest creates a request to invoke DescribeDBNodePerformance API
func CreateDescribeDBNodePerformanceRequest() (request *DescribeDBNodePerformanceRequest) {
	request = &DescribeDBNodePerformanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBNodePerformance", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBNodePerformanceResponse creates a response to parse from DescribeDBNodePerformance response
func CreateDescribeDBNodePerformanceResponse() (response *DescribeDBNodePerformanceResponse) {
	response = &DescribeDBNodePerformanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
