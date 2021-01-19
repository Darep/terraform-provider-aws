package waiter

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/service/elasticache/finder"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/tfresource"
)

const (
	ReplicationGroupStatusCreating     = "creating"
	ReplicationGroupStatusAvailable    = "available"
	ReplicationGroupStatusModifying    = "modifying"
	ReplicationGroupStatusDeleting     = "deleting"
	ReplicationGroupStatusCreateFailed = "create-failed"
	ReplicationGroupStatusSnapshotting = "snapshotting"
)

// ReplicationGroupStatus fetches the ReplicationGroup and its Status
func ReplicationGroupStatus(conn *elasticache.ElastiCache, replicationGroupID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		rg, err := finder.ReplicationGroupByID(conn, replicationGroupID)
		var e *resource.NotFoundError
		if errors.As(err, &e) {
			return nil, "", nil
		}
		if tfresource.NotFound(err) {
			return nil, "", nil
		}
		if err != nil {
			return nil, "", err
		}

		return rg, aws.StringValue(rg.Status), nil
	}
}
