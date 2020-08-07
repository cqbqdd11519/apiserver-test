package apis

import v1 "github.com/cqbqdd11519/apiserver-test/pkg/apis/v1"

func init() {
	AddApiFuncs = append(AddApiFuncs, v1.AddV1Apis)
}
