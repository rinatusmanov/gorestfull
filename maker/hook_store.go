package maker

import "github.com/rinatusmanov/gorestfull/types"

// hookStore store of hooks
type hookStore struct {
	//
	BeforeDBFind   types.THookFunc
	AfterDBFind    types.THookFunc
	BeforeDBDelete types.THookFunc
	AfterDBDelete  types.THookFunc
	BeforeDBCreate types.THookFunc
	AfterDBCreate  types.THookFunc
	BeforeDBChange types.THookFunc
	AfterDBChange  types.THookFunc
}

var emptyHook types.THookFunc = func(response types.IResponse, db types.IGormDB, rw types.IReadWriter) (resultDD types.IGormDB, err error) {
	return db, nil
}

// конструктор для hookStore
func NewHookStore(data interface{}) *hookStore {
	BeforeDBFind := emptyHook
	iBeforeGet, okIBeforeGet := data.(types.IBeforeDBFind)
	if okIBeforeGet {
		BeforeDBFind = iBeforeGet.BeforeDBFind
	}

	AfterDBFind := emptyHook
	iAfterGet, okIAfterGet := data.(types.IAfterDBFind)
	if okIAfterGet {
		AfterDBFind = iAfterGet.AfterDBFind
	}

	BeforeDBDelete := emptyHook
	iBeforeDelete, okIBeforeDelete := data.(types.IBeforeDBDelete)
	if okIBeforeDelete {
		BeforeDBDelete = iBeforeDelete.BeforeDBDelete
	}

	AfterDBDelete := emptyHook
	iAfterDelete, okIAfterDelete := data.(types.IAfterDBDelete)
	if okIAfterDelete {
		AfterDBDelete = iAfterDelete.AfterDBDelete
	}

	BeforeDBCreate := emptyHook
	iBeforePost, okIBeforePost := data.(types.IBeforeDBCreate)
	if okIBeforePost {
		BeforeDBCreate = iBeforePost.BeforeDBCreate
	}

	AfterDBCreate := emptyHook
	iAfterPost, okIAfterPost := data.(types.IAfterDBCreate)
	if okIAfterPost {
		AfterDBCreate = iAfterPost.AfterDBCreate
	}

	BeforeDBChange := emptyHook
	iBeforePut, okIBeforePut := data.(types.IBeforeDBChange)
	if okIBeforePut {
		BeforeDBChange = iBeforePut.BeforeDBChange
	}

	AfterDBChange := emptyHook
	iAfterPut, okIAfterPut := data.(types.IAfterDBChange)
	if okIAfterPut {
		AfterDBChange = iAfterPut.AfterDBChange
	}

	return &hookStore{
		BeforeDBFind:   BeforeDBFind,
		AfterDBFind:    AfterDBFind,
		BeforeDBDelete: BeforeDBDelete,
		AfterDBDelete:  AfterDBDelete,
		BeforeDBCreate: BeforeDBCreate,
		AfterDBCreate:  AfterDBCreate,
		BeforeDBChange: BeforeDBChange,
		AfterDBChange:  AfterDBChange,
	}
}
