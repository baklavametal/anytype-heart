// Code generated by mockery v2.38.0. DO NOT EDIT.

package mock_filesync

import (
	context "context"

	app "github.com/anyproto/any-sync/app"

	filesync "github.com/anyproto/anytype-heart/core/filestorage/filesync"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// MockFileSync is an autogenerated mock type for the FileSync type
type MockFileSync struct {
	mock.Mock
}

type MockFileSync_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFileSync) EXPECT() *MockFileSync_Expecter {
	return &MockFileSync_Expecter{mock: &_m.Mock}
}

// AddFile provides a mock function with given fields: spaceID, fileID, uploadedByUser, imported
func (_m *MockFileSync) AddFile(spaceID string, fileID string, uploadedByUser bool, imported bool) error {
	ret := _m.Called(spaceID, fileID, uploadedByUser, imported)

	if len(ret) == 0 {
		panic("no return value specified for AddFile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, bool, bool) error); ok {
		r0 = rf(spaceID, fileID, uploadedByUser, imported)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFileSync_AddFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFile'
type MockFileSync_AddFile_Call struct {
	*mock.Call
}

// AddFile is a helper method to define mock.On call
//   - spaceID string
//   - fileID string
//   - uploadedByUser bool
//   - imported bool
func (_e *MockFileSync_Expecter) AddFile(spaceID interface{}, fileID interface{}, uploadedByUser interface{}, imported interface{}) *MockFileSync_AddFile_Call {
	return &MockFileSync_AddFile_Call{Call: _e.mock.On("AddFile", spaceID, fileID, uploadedByUser, imported)}
}

func (_c *MockFileSync_AddFile_Call) Run(run func(spaceID string, fileID string, uploadedByUser bool, imported bool)) *MockFileSync_AddFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(bool), args[3].(bool))
	})
	return _c
}

func (_c *MockFileSync_AddFile_Call) Return(err error) *MockFileSync_AddFile_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockFileSync_AddFile_Call) RunAndReturn(run func(string, string, bool, bool) error) *MockFileSync_AddFile_Call {
	_c.Call.Return(run)
	return _c
}

// CalculateFileSize provides a mock function with given fields: ctx, spaceId, fileID
func (_m *MockFileSync) CalculateFileSize(ctx context.Context, spaceId string, fileID string) (int, error) {
	ret := _m.Called(ctx, spaceId, fileID)

	if len(ret) == 0 {
		panic("no return value specified for CalculateFileSize")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (int, error)); ok {
		return rf(ctx, spaceId, fileID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) int); ok {
		r0 = rf(ctx, spaceId, fileID)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, spaceId, fileID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFileSync_CalculateFileSize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CalculateFileSize'
type MockFileSync_CalculateFileSize_Call struct {
	*mock.Call
}

// CalculateFileSize is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceId string
//   - fileID string
func (_e *MockFileSync_Expecter) CalculateFileSize(ctx interface{}, spaceId interface{}, fileID interface{}) *MockFileSync_CalculateFileSize_Call {
	return &MockFileSync_CalculateFileSize_Call{Call: _e.mock.On("CalculateFileSize", ctx, spaceId, fileID)}
}

func (_c *MockFileSync_CalculateFileSize_Call) Run(run func(ctx context.Context, spaceId string, fileID string)) *MockFileSync_CalculateFileSize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockFileSync_CalculateFileSize_Call) Return(_a0 int, _a1 error) *MockFileSync_CalculateFileSize_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFileSync_CalculateFileSize_Call) RunAndReturn(run func(context.Context, string, string) (int, error)) *MockFileSync_CalculateFileSize_Call {
	_c.Call.Return(run)
	return _c
}

// ClearImportEvents provides a mock function with given fields:
func (_m *MockFileSync) ClearImportEvents() {
	_m.Called()
}

// MockFileSync_ClearImportEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClearImportEvents'
type MockFileSync_ClearImportEvents_Call struct {
	*mock.Call
}

// ClearImportEvents is a helper method to define mock.On call
func (_e *MockFileSync_Expecter) ClearImportEvents() *MockFileSync_ClearImportEvents_Call {
	return &MockFileSync_ClearImportEvents_Call{Call: _e.mock.On("ClearImportEvents")}
}

func (_c *MockFileSync_ClearImportEvents_Call) Run(run func()) *MockFileSync_ClearImportEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFileSync_ClearImportEvents_Call) Return() *MockFileSync_ClearImportEvents_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockFileSync_ClearImportEvents_Call) RunAndReturn(run func()) *MockFileSync_ClearImportEvents_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields: ctx
func (_m *MockFileSync) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFileSync_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockFileSync_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockFileSync_Expecter) Close(ctx interface{}) *MockFileSync_Close_Call {
	return &MockFileSync_Close_Call{Call: _e.mock.On("Close", ctx)}
}

func (_c *MockFileSync_Close_Call) Run(run func(ctx context.Context)) *MockFileSync_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockFileSync_Close_Call) Return(err error) *MockFileSync_Close_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockFileSync_Close_Call) RunAndReturn(run func(context.Context) error) *MockFileSync_Close_Call {
	_c.Call.Return(run)
	return _c
}

// DebugQueue provides a mock function with given fields: _a0
func (_m *MockFileSync) DebugQueue(_a0 *http.Request) (*filesync.QueueInfo, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for DebugQueue")
	}

	var r0 *filesync.QueueInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(*http.Request) (*filesync.QueueInfo, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*http.Request) *filesync.QueueInfo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*filesync.QueueInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(*http.Request) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFileSync_DebugQueue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DebugQueue'
type MockFileSync_DebugQueue_Call struct {
	*mock.Call
}

// DebugQueue is a helper method to define mock.On call
//   - _a0 *http.Request
func (_e *MockFileSync_Expecter) DebugQueue(_a0 interface{}) *MockFileSync_DebugQueue_Call {
	return &MockFileSync_DebugQueue_Call{Call: _e.mock.On("DebugQueue", _a0)}
}

func (_c *MockFileSync_DebugQueue_Call) Run(run func(_a0 *http.Request)) *MockFileSync_DebugQueue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*http.Request))
	})
	return _c
}

func (_c *MockFileSync_DebugQueue_Call) Return(_a0 *filesync.QueueInfo, _a1 error) *MockFileSync_DebugQueue_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFileSync_DebugQueue_Call) RunAndReturn(run func(*http.Request) (*filesync.QueueInfo, error)) *MockFileSync_DebugQueue_Call {
	_c.Call.Return(run)
	return _c
}

// FileListStats provides a mock function with given fields: ctx, spaceId, fileIDs
func (_m *MockFileSync) FileListStats(ctx context.Context, spaceId string, fileIDs []string) ([]filesync.FileStat, error) {
	ret := _m.Called(ctx, spaceId, fileIDs)

	if len(ret) == 0 {
		panic("no return value specified for FileListStats")
	}

	var r0 []filesync.FileStat
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) ([]filesync.FileStat, error)); ok {
		return rf(ctx, spaceId, fileIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) []filesync.FileStat); ok {
		r0 = rf(ctx, spaceId, fileIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]filesync.FileStat)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []string) error); ok {
		r1 = rf(ctx, spaceId, fileIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFileSync_FileListStats_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FileListStats'
type MockFileSync_FileListStats_Call struct {
	*mock.Call
}

// FileListStats is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceId string
//   - fileIDs []string
func (_e *MockFileSync_Expecter) FileListStats(ctx interface{}, spaceId interface{}, fileIDs interface{}) *MockFileSync_FileListStats_Call {
	return &MockFileSync_FileListStats_Call{Call: _e.mock.On("FileListStats", ctx, spaceId, fileIDs)}
}

func (_c *MockFileSync_FileListStats_Call) Run(run func(ctx context.Context, spaceId string, fileIDs []string)) *MockFileSync_FileListStats_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]string))
	})
	return _c
}

func (_c *MockFileSync_FileListStats_Call) Return(_a0 []filesync.FileStat, _a1 error) *MockFileSync_FileListStats_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFileSync_FileListStats_Call) RunAndReturn(run func(context.Context, string, []string) ([]filesync.FileStat, error)) *MockFileSync_FileListStats_Call {
	_c.Call.Return(run)
	return _c
}

// FileStat provides a mock function with given fields: ctx, spaceId, fileId
func (_m *MockFileSync) FileStat(ctx context.Context, spaceId string, fileId string) (filesync.FileStat, error) {
	ret := _m.Called(ctx, spaceId, fileId)

	if len(ret) == 0 {
		panic("no return value specified for FileStat")
	}

	var r0 filesync.FileStat
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (filesync.FileStat, error)); ok {
		return rf(ctx, spaceId, fileId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) filesync.FileStat); ok {
		r0 = rf(ctx, spaceId, fileId)
	} else {
		r0 = ret.Get(0).(filesync.FileStat)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, spaceId, fileId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFileSync_FileStat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FileStat'
type MockFileSync_FileStat_Call struct {
	*mock.Call
}

// FileStat is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceId string
//   - fileId string
func (_e *MockFileSync_Expecter) FileStat(ctx interface{}, spaceId interface{}, fileId interface{}) *MockFileSync_FileStat_Call {
	return &MockFileSync_FileStat_Call{Call: _e.mock.On("FileStat", ctx, spaceId, fileId)}
}

func (_c *MockFileSync_FileStat_Call) Run(run func(ctx context.Context, spaceId string, fileId string)) *MockFileSync_FileStat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockFileSync_FileStat_Call) Return(fs filesync.FileStat, err error) *MockFileSync_FileStat_Call {
	_c.Call.Return(fs, err)
	return _c
}

func (_c *MockFileSync_FileStat_Call) RunAndReturn(run func(context.Context, string, string) (filesync.FileStat, error)) *MockFileSync_FileStat_Call {
	_c.Call.Return(run)
	return _c
}

// HasUpload provides a mock function with given fields: spaceId, fileId
func (_m *MockFileSync) HasUpload(spaceId string, fileId string) (bool, error) {
	ret := _m.Called(spaceId, fileId)

	if len(ret) == 0 {
		panic("no return value specified for HasUpload")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (bool, error)); ok {
		return rf(spaceId, fileId)
	}
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(spaceId, fileId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(spaceId, fileId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFileSync_HasUpload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasUpload'
type MockFileSync_HasUpload_Call struct {
	*mock.Call
}

// HasUpload is a helper method to define mock.On call
//   - spaceId string
//   - fileId string
func (_e *MockFileSync_Expecter) HasUpload(spaceId interface{}, fileId interface{}) *MockFileSync_HasUpload_Call {
	return &MockFileSync_HasUpload_Call{Call: _e.mock.On("HasUpload", spaceId, fileId)}
}

func (_c *MockFileSync_HasUpload_Call) Run(run func(spaceId string, fileId string)) *MockFileSync_HasUpload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockFileSync_HasUpload_Call) Return(ok bool, err error) *MockFileSync_HasUpload_Call {
	_c.Call.Return(ok, err)
	return _c
}

func (_c *MockFileSync_HasUpload_Call) RunAndReturn(run func(string, string) (bool, error)) *MockFileSync_HasUpload_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: a
func (_m *MockFileSync) Init(a *app.App) error {
	ret := _m.Called(a)

	if len(ret) == 0 {
		panic("no return value specified for Init")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*app.App) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFileSync_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockFileSync_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - a *app.App
func (_e *MockFileSync_Expecter) Init(a interface{}) *MockFileSync_Init_Call {
	return &MockFileSync_Init_Call{Call: _e.mock.On("Init", a)}
}

func (_c *MockFileSync_Init_Call) Run(run func(a *app.App)) *MockFileSync_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*app.App))
	})
	return _c
}

func (_c *MockFileSync_Init_Call) Return(err error) *MockFileSync_Init_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockFileSync_Init_Call) RunAndReturn(run func(*app.App) error) *MockFileSync_Init_Call {
	_c.Call.Return(run)
	return _c
}

// IsFileUploadLimited provides a mock function with given fields: spaceId, fileId
func (_m *MockFileSync) IsFileUploadLimited(spaceId string, fileId string) (bool, error) {
	ret := _m.Called(spaceId, fileId)

	if len(ret) == 0 {
		panic("no return value specified for IsFileUploadLimited")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (bool, error)); ok {
		return rf(spaceId, fileId)
	}
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(spaceId, fileId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(spaceId, fileId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFileSync_IsFileUploadLimited_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsFileUploadLimited'
type MockFileSync_IsFileUploadLimited_Call struct {
	*mock.Call
}

// IsFileUploadLimited is a helper method to define mock.On call
//   - spaceId string
//   - fileId string
func (_e *MockFileSync_Expecter) IsFileUploadLimited(spaceId interface{}, fileId interface{}) *MockFileSync_IsFileUploadLimited_Call {
	return &MockFileSync_IsFileUploadLimited_Call{Call: _e.mock.On("IsFileUploadLimited", spaceId, fileId)}
}

func (_c *MockFileSync_IsFileUploadLimited_Call) Run(run func(spaceId string, fileId string)) *MockFileSync_IsFileUploadLimited_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockFileSync_IsFileUploadLimited_Call) Return(ok bool, err error) *MockFileSync_IsFileUploadLimited_Call {
	_c.Call.Return(ok, err)
	return _c
}

func (_c *MockFileSync_IsFileUploadLimited_Call) RunAndReturn(run func(string, string) (bool, error)) *MockFileSync_IsFileUploadLimited_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockFileSync) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockFileSync_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockFileSync_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockFileSync_Expecter) Name() *MockFileSync_Name_Call {
	return &MockFileSync_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockFileSync_Name_Call) Run(run func()) *MockFileSync_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFileSync_Name_Call) Return(name string) *MockFileSync_Name_Call {
	_c.Call.Return(name)
	return _c
}

func (_c *MockFileSync_Name_Call) RunAndReturn(run func() string) *MockFileSync_Name_Call {
	_c.Call.Return(run)
	return _c
}

// NodeUsage provides a mock function with given fields: ctx
func (_m *MockFileSync) NodeUsage(ctx context.Context) (filesync.NodeUsage, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for NodeUsage")
	}

	var r0 filesync.NodeUsage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (filesync.NodeUsage, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) filesync.NodeUsage); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(filesync.NodeUsage)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFileSync_NodeUsage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NodeUsage'
type MockFileSync_NodeUsage_Call struct {
	*mock.Call
}

// NodeUsage is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockFileSync_Expecter) NodeUsage(ctx interface{}) *MockFileSync_NodeUsage_Call {
	return &MockFileSync_NodeUsage_Call{Call: _e.mock.On("NodeUsage", ctx)}
}

func (_c *MockFileSync_NodeUsage_Call) Run(run func(ctx context.Context)) *MockFileSync_NodeUsage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockFileSync_NodeUsage_Call) Return(usage filesync.NodeUsage, err error) *MockFileSync_NodeUsage_Call {
	_c.Call.Return(usage, err)
	return _c
}

func (_c *MockFileSync_NodeUsage_Call) RunAndReturn(run func(context.Context) (filesync.NodeUsage, error)) *MockFileSync_NodeUsage_Call {
	_c.Call.Return(run)
	return _c
}

// OnUpload provides a mock function with given fields: _a0
func (_m *MockFileSync) OnUpload(_a0 func(string, string) error) {
	_m.Called(_a0)
}

// MockFileSync_OnUpload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnUpload'
type MockFileSync_OnUpload_Call struct {
	*mock.Call
}

// OnUpload is a helper method to define mock.On call
//   - _a0 func(string , string) error
func (_e *MockFileSync_Expecter) OnUpload(_a0 interface{}) *MockFileSync_OnUpload_Call {
	return &MockFileSync_OnUpload_Call{Call: _e.mock.On("OnUpload", _a0)}
}

func (_c *MockFileSync_OnUpload_Call) Run(run func(_a0 func(string, string) error)) *MockFileSync_OnUpload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(string, string) error))
	})
	return _c
}

func (_c *MockFileSync_OnUpload_Call) Return() *MockFileSync_OnUpload_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockFileSync_OnUpload_Call) RunAndReturn(run func(func(string, string) error)) *MockFileSync_OnUpload_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveFile provides a mock function with given fields: spaceId, fileId
func (_m *MockFileSync) RemoveFile(spaceId string, fileId string) error {
	ret := _m.Called(spaceId, fileId)

	if len(ret) == 0 {
		panic("no return value specified for RemoveFile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(spaceId, fileId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFileSync_RemoveFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveFile'
type MockFileSync_RemoveFile_Call struct {
	*mock.Call
}

// RemoveFile is a helper method to define mock.On call
//   - spaceId string
//   - fileId string
func (_e *MockFileSync_Expecter) RemoveFile(spaceId interface{}, fileId interface{}) *MockFileSync_RemoveFile_Call {
	return &MockFileSync_RemoveFile_Call{Call: _e.mock.On("RemoveFile", spaceId, fileId)}
}

func (_c *MockFileSync_RemoveFile_Call) Run(run func(spaceId string, fileId string)) *MockFileSync_RemoveFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockFileSync_RemoveFile_Call) Return(err error) *MockFileSync_RemoveFile_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockFileSync_RemoveFile_Call) RunAndReturn(run func(string, string) error) *MockFileSync_RemoveFile_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: ctx
func (_m *MockFileSync) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFileSync_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockFileSync_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockFileSync_Expecter) Run(ctx interface{}) *MockFileSync_Run_Call {
	return &MockFileSync_Run_Call{Call: _e.mock.On("Run", ctx)}
}

func (_c *MockFileSync_Run_Call) Run(run func(ctx context.Context)) *MockFileSync_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockFileSync_Run_Call) Return(err error) *MockFileSync_Run_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockFileSync_Run_Call) RunAndReturn(run func(context.Context) error) *MockFileSync_Run_Call {
	_c.Call.Return(run)
	return _c
}

// SendImportEvents provides a mock function with given fields:
func (_m *MockFileSync) SendImportEvents() {
	_m.Called()
}

// MockFileSync_SendImportEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendImportEvents'
type MockFileSync_SendImportEvents_Call struct {
	*mock.Call
}

// SendImportEvents is a helper method to define mock.On call
func (_e *MockFileSync_Expecter) SendImportEvents() *MockFileSync_SendImportEvents_Call {
	return &MockFileSync_SendImportEvents_Call{Call: _e.mock.On("SendImportEvents")}
}

func (_c *MockFileSync_SendImportEvents_Call) Run(run func()) *MockFileSync_SendImportEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFileSync_SendImportEvents_Call) Return() *MockFileSync_SendImportEvents_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockFileSync_SendImportEvents_Call) RunAndReturn(run func()) *MockFileSync_SendImportEvents_Call {
	_c.Call.Return(run)
	return _c
}

// SpaceStat provides a mock function with given fields: ctx, spaceId
func (_m *MockFileSync) SpaceStat(ctx context.Context, spaceId string) (filesync.SpaceStat, error) {
	ret := _m.Called(ctx, spaceId)

	if len(ret) == 0 {
		panic("no return value specified for SpaceStat")
	}

	var r0 filesync.SpaceStat
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (filesync.SpaceStat, error)); ok {
		return rf(ctx, spaceId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) filesync.SpaceStat); ok {
		r0 = rf(ctx, spaceId)
	} else {
		r0 = ret.Get(0).(filesync.SpaceStat)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, spaceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFileSync_SpaceStat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SpaceStat'
type MockFileSync_SpaceStat_Call struct {
	*mock.Call
}

// SpaceStat is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceId string
func (_e *MockFileSync_Expecter) SpaceStat(ctx interface{}, spaceId interface{}) *MockFileSync_SpaceStat_Call {
	return &MockFileSync_SpaceStat_Call{Call: _e.mock.On("SpaceStat", ctx, spaceId)}
}

func (_c *MockFileSync_SpaceStat_Call) Run(run func(ctx context.Context, spaceId string)) *MockFileSync_SpaceStat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockFileSync_SpaceStat_Call) Return(ss filesync.SpaceStat, err error) *MockFileSync_SpaceStat_Call {
	_c.Call.Return(ss, err)
	return _c
}

func (_c *MockFileSync_SpaceStat_Call) RunAndReturn(run func(context.Context, string) (filesync.SpaceStat, error)) *MockFileSync_SpaceStat_Call {
	_c.Call.Return(run)
	return _c
}

// SyncStatus provides a mock function with given fields:
func (_m *MockFileSync) SyncStatus() (filesync.SyncStatus, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SyncStatus")
	}

	var r0 filesync.SyncStatus
	var r1 error
	if rf, ok := ret.Get(0).(func() (filesync.SyncStatus, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() filesync.SyncStatus); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(filesync.SyncStatus)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFileSync_SyncStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SyncStatus'
type MockFileSync_SyncStatus_Call struct {
	*mock.Call
}

// SyncStatus is a helper method to define mock.On call
func (_e *MockFileSync_Expecter) SyncStatus() *MockFileSync_SyncStatus_Call {
	return &MockFileSync_SyncStatus_Call{Call: _e.mock.On("SyncStatus")}
}

func (_c *MockFileSync_SyncStatus_Call) Run(run func()) *MockFileSync_SyncStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFileSync_SyncStatus_Call) Return(ss filesync.SyncStatus, err error) *MockFileSync_SyncStatus_Call {
	_c.Call.Return(ss, err)
	return _c
}

func (_c *MockFileSync_SyncStatus_Call) RunAndReturn(run func() (filesync.SyncStatus, error)) *MockFileSync_SyncStatus_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFileSync creates a new instance of MockFileSync. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFileSync(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFileSync {
	mock := &MockFileSync{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
