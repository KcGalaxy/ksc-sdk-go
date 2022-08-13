// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package memcachediface provides an interface to enable mocking the memcached service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package memcachediface

import (
	"github.com/KcGalaxy/ksc-sdk-go/service/memcached"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// MemcachedAPI provides an interface to enable mocking the
// memcached.Memcached service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // memcached.
//    func myFunc(svc memcachediface.MemcachedAPI) bool {
//        // Make svc.CreateCacheCluster request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := memcached.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMemcachedClient struct {
//        memcachediface.MemcachedAPI
//    }
//    func (m *mockMemcachedClient) CreateCacheCluster(input *map[string]interface{}) (*map[string]interface{}, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMemcachedClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type MemcachedAPI interface {
	CreateCacheCluster(*map[string]interface{}) (*map[string]interface{}, error)
	CreateCacheClusterWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateCacheClusterRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteCacheCluster(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteCacheClusterWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteCacheClusterRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteCacheSecurityRule(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteCacheSecurityRuleWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteCacheSecurityRuleRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAvailabilityZones(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAvailabilityZonesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAvailabilityZonesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCacheCluster(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCacheClusterWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCacheClusterRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCacheClusters(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCacheClustersWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCacheClustersRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCacheSecurityRules(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCacheSecurityRulesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCacheSecurityRulesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRegions(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRegionsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRegionsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	FlushCacheCluster(*map[string]interface{}) (*map[string]interface{}, error)
	FlushCacheClusterWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	FlushCacheClusterRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RenameCacheCluster(*map[string]interface{}) (*map[string]interface{}, error)
	RenameCacheClusterWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RenameCacheClusterRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ResizeCacheCluster(*map[string]interface{}) (*map[string]interface{}, error)
	ResizeCacheClusterWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ResizeCacheClusterRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetCacheSecurityRules(*map[string]interface{}) (*map[string]interface{}, error)
	SetCacheSecurityRulesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetCacheSecurityRulesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdatePassword(*map[string]interface{}) (*map[string]interface{}, error)
	UpdatePasswordWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdatePasswordRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ MemcachedAPI = (*memcached.Memcached)(nil)
