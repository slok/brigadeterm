// Code generated by mockery v1.0.0
package storage

import brigade "github.com/Azure/brigade/pkg/brigade"
import io "io"
import mock "github.com/stretchr/testify/mock"

import time "time"

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// BlockUntilAPICacheSynced provides a mock function with given fields: waitUntil
func (_m *Store) BlockUntilAPICacheSynced(waitUntil <-chan time.Time) bool {
	ret := _m.Called(waitUntil)

	var r0 bool
	if rf, ok := ret.Get(0).(func(<-chan time.Time) bool); ok {
		r0 = rf(waitUntil)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CreateBuild provides a mock function with given fields: build
func (_m *Store) CreateBuild(build *brigade.Build) error {
	ret := _m.Called(build)

	var r0 error
	if rf, ok := ret.Get(0).(func(*brigade.Build) error); ok {
		r0 = rf(build)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBuild provides a mock function with given fields: id
func (_m *Store) GetBuild(id string) (*brigade.Build, error) {
	ret := _m.Called(id)

	var r0 *brigade.Build
	if rf, ok := ret.Get(0).(func(string) *brigade.Build); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brigade.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBuildJobs provides a mock function with given fields: build
func (_m *Store) GetBuildJobs(build *brigade.Build) ([]*brigade.Job, error) {
	ret := _m.Called(build)

	var r0 []*brigade.Job
	if rf, ok := ret.Get(0).(func(*brigade.Build) []*brigade.Job); ok {
		r0 = rf(build)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*brigade.Job)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*brigade.Build) error); ok {
		r1 = rf(build)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBuilds provides a mock function with given fields:
func (_m *Store) GetBuilds() ([]*brigade.Build, error) {
	ret := _m.Called()

	var r0 []*brigade.Build
	if rf, ok := ret.Get(0).(func() []*brigade.Build); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*brigade.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJob provides a mock function with given fields: id
func (_m *Store) GetJob(id string) (*brigade.Job, error) {
	ret := _m.Called(id)

	var r0 *brigade.Job
	if rf, ok := ret.Get(0).(func(string) *brigade.Job); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brigade.Job)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobLog provides a mock function with given fields: job
func (_m *Store) GetJobLog(job *brigade.Job) (string, error) {
	ret := _m.Called(job)

	var r0 string
	if rf, ok := ret.Get(0).(func(*brigade.Job) string); ok {
		r0 = rf(job)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*brigade.Job) error); ok {
		r1 = rf(job)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobLogStream provides a mock function with given fields: job
func (_m *Store) GetJobLogStream(job *brigade.Job) (io.ReadCloser, error) {
	ret := _m.Called(job)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(*brigade.Job) io.ReadCloser); ok {
		r0 = rf(job)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*brigade.Job) error); ok {
		r1 = rf(job)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobLogStreamFollow provides a mock function with given fields: job
func (_m *Store) GetJobLogStreamFollow(job *brigade.Job) (io.ReadCloser, error) {
	ret := _m.Called(job)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(*brigade.Job) io.ReadCloser); ok {
		r0 = rf(job)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*brigade.Job) error); ok {
		r1 = rf(job)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProject provides a mock function with given fields: id
func (_m *Store) GetProject(id string) (*brigade.Project, error) {
	ret := _m.Called(id)

	var r0 *brigade.Project
	if rf, ok := ret.Get(0).(func(string) *brigade.Project); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brigade.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectBuilds provides a mock function with given fields: proj
func (_m *Store) GetProjectBuilds(proj *brigade.Project) ([]*brigade.Build, error) {
	ret := _m.Called(proj)

	var r0 []*brigade.Build
	if rf, ok := ret.Get(0).(func(*brigade.Project) []*brigade.Build); ok {
		r0 = rf(proj)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*brigade.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*brigade.Project) error); ok {
		r1 = rf(proj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjects provides a mock function with given fields:
func (_m *Store) GetProjects() ([]*brigade.Project, error) {
	ret := _m.Called()

	var r0 []*brigade.Project
	if rf, ok := ret.Get(0).(func() []*brigade.Project); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*brigade.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorker provides a mock function with given fields: buildID
func (_m *Store) GetWorker(buildID string) (*brigade.Worker, error) {
	ret := _m.Called(buildID)

	var r0 *brigade.Worker
	if rf, ok := ret.Get(0).(func(string) *brigade.Worker); ok {
		r0 = rf(buildID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brigade.Worker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(buildID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkerLog provides a mock function with given fields: job
func (_m *Store) GetWorkerLog(job *brigade.Worker) (string, error) {
	ret := _m.Called(job)

	var r0 string
	if rf, ok := ret.Get(0).(func(*brigade.Worker) string); ok {
		r0 = rf(job)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*brigade.Worker) error); ok {
		r1 = rf(job)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkerLogStream provides a mock function with given fields: job
func (_m *Store) GetWorkerLogStream(job *brigade.Worker) (io.ReadCloser, error) {
	ret := _m.Called(job)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(*brigade.Worker) io.ReadCloser); ok {
		r0 = rf(job)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*brigade.Worker) error); ok {
		r1 = rf(job)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkerLogStreamFollow provides a mock function with given fields: job
func (_m *Store) GetWorkerLogStreamFollow(job *brigade.Worker) (io.ReadCloser, error) {
	ret := _m.Called(job)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(*brigade.Worker) io.ReadCloser); ok {
		r0 = rf(job)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*brigade.Worker) error); ok {
		r1 = rf(job)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
