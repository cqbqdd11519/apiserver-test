package v1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	tmaxv1 "github.com/tmax-cloud/approval-watcher/pkg/apis/tmax/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/cqbqdd11519/apiserver-test/internal/utils"
	"github.com/cqbqdd11519/apiserver-test/internal/wrapper"
)

func AddApproveApis(parent *wrapper.RouterWrapper) error {
	approveWrapper := wrapper.New("/approve", []string{"PUT"}, approveHandler)
	if err := parent.Add(approveWrapper); err != nil {
		return err
	}

	return nil
}

func AddRejectApis(parent *wrapper.RouterWrapper) error {
	approveWrapper := wrapper.New("/reject", []string{"PUT"}, rejectHandler)
	if err := parent.Add(approveWrapper); err != nil {
		return err
	}

	return nil
}

func updateDecision(w http.ResponseWriter, req *http.Request, decision tmaxv1.Result) {
	vars := mux.Vars(req)

	ns, nsExist := vars["namespace"]
	approvalName, nameExist := vars["approvalName"]
	if !nsExist || !nameExist {
		_ = utils.RespondError(w, http.StatusBadRequest, "url is malformed")
		return
	}

	opt := client.Options{}
	utils.AddSchemes(&opt, schema.GroupVersion{Group: "tmax.io", Version: "v1"}, &tmaxv1.Approval{})

	c, err := utils.Client(opt)
	if err != nil {
		log.Error(err, "cannot get client")
		_ = utils.RespondError(w, http.StatusInternalServerError, "could not make k8s client")
		return
	}

	approval := &tmaxv1.Approval{}
	if err := c.Get(context.TODO(), types.NamespacedName{Name: approvalName, Namespace: ns}, approval); err != nil {
		log.Error(err, "cannot get approval")
		if errors.IsNotFound(err) {
			_ = utils.RespondError(w, http.StatusNotFound, fmt.Sprintf("there is no Approval %s/%s", ns, approvalName))
		} else {
			_ = utils.RespondError(w, http.StatusInternalServerError, "cannot get approval")
		}
		return
	}

	log.Info(string(decision))
	// TODO : actual logic to approve

	_ = utils.RespondJSON(w, approval)
}

func approveHandler(w http.ResponseWriter, req *http.Request) {
	updateDecision(w, req, tmaxv1.ResultApproved)
}

func rejectHandler(w http.ResponseWriter, req *http.Request) {
	updateDecision(w, req, tmaxv1.ResultRejected)
}
