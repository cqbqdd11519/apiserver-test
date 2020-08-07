package v1

import (
	"fmt"

	"github.com/cqbqdd11519/apiserver-test/internal/wrapper"
)

func AddApprovalApis(parent *wrapper.RouterWrapper) error {
	approvalWrapper := wrapper.New(fmt.Sprintf("/%s/{approvalName}", ApprovalKind), nil, nil)
	if err := parent.Add(approvalWrapper); err != nil {
		return err
	}

	// TODO: middleware - authenticate

	if err := AddApproveApis(approvalWrapper); err != nil {
		return err
	}
	if err := AddRejectApis(approvalWrapper); err != nil {
		return err
	}
	return nil
}
