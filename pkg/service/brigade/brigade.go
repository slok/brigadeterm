package brigade

import (
	"fmt"
	"io"
	"sort"

	"github.com/Azure/brigade/pkg/storage"

	brigademodel "github.com/slok/brigadeterm/pkg/model/brigade"
)

// Service is the service where all the brigade data will be retrieved
// and prepared with the information that this applications needs.
type Service interface {
	// GetProjectBuilds will get one project.
	GetProject(projectID string) (*brigademodel.Project, error)
	// GetProjectLastBuild will get projects last builds.
	GetProjectLastBuilds(projectID string, quantity int) ([]*brigademodel.Build, error)
	// GetProjects will get all the projects that are on brigade.
	GetProjects() ([]*brigademodel.Project, error)
	// GetBuild will get one build.
	GetBuild(buildID string) (*brigademodel.Build, error)
	// GetProjectBuilds will get all the builds of a project in descendant or ascendant order.
	GetProjectBuilds(project *brigademodel.Project, desc bool) ([]*brigademodel.Build, error)
	// GetBuildJobs will get all the jobs of a build in descendant or ascendant order.
	GetBuildJobs(BuildID string, desc bool) ([]*brigademodel.Job, error)
	// GetJob will get a job.
	GetJob(jobID string) (*brigademodel.Job, error)
	// GetJobLog will get a job log.
	GetJobLog(jobID string) (string, error)
	// GetJobLogStream will get a job log stream.
	GetJobLogStream(jobID string) (io.ReadCloser, error)
}

// repository will use kubernetes as repository for the brigade objects.
type service struct {
	client storage.Store
}

// NewService returns a new brigade service.
func NewService(brigadestore storage.Store) Service {
	return &service{
		client: brigadestore,
	}
}

func (s *service) GetProject(projectID string) (*brigademodel.Project, error) {
	prj, err := s.client.GetProject(projectID)

	if err != nil {
		return nil, err
	}
	res := brigademodel.Project(*prj)
	return &res, nil
}

func (s *service) GetProjectLastBuilds(projectID string, quantity int) ([]*brigademodel.Build, error) {
	prj, err := s.client.GetProject(projectID)

	if err != nil {
		return nil, err
	}

	// Get the available builds.
	builds, err := s.GetProjectBuilds(prj, true)
	if err != nil {
		return nil, err
	}
	if len(builds) == 0 {
		return nil, fmt.Errorf("no builds available")
	}

	// Get last one.
	if len(builds) > quantity {
		builds = builds[:quantity]
	}
	lastBuilds := make([]*brigademodel.Build, len(builds))

	for i, b := range builds {
		lb := brigademodel.Build(*b)
		lastBuilds[i] = &lb
	}

	return lastBuilds, nil
}

// GetProjects satisfies Service interface.
func (s *service) GetProjects() ([]*brigademodel.Project, error) {
	prjs, err := s.client.GetProjects()

	if err != nil {
		return nil, err
	}

	// Sort projects by name.
	sort.Slice(prjs, func(i, j int) bool {
		return prjs[i].Name < prjs[j].Name
	})

	prjList := make([]*brigademodel.Project, len(prjs))
	for i, prj := range prjs {
		p := brigademodel.Project(*prj)
		prjList[i] = &p
	}

	return prjList, nil
}

func (s *service) GetBuild(buildID string) (*brigademodel.Build, error) {
	bld, err := s.client.GetBuild(buildID)

	if err != nil {
		return nil, err
	}
	res := brigademodel.Build(*bld)
	return &res, nil
}

// GetAllProjects satisfies Service interface.
func (s *service) GetProjectBuilds(project *brigademodel.Project, desc bool) ([]*brigademodel.Build, error) {
	pr, err := s.client.GetProject(project.ID)
	if err != nil {
		return []*brigademodel.Build{}, err
	}

	builds, err := s.client.GetProjectBuilds(pr)
	if err != nil {
		return []*brigademodel.Build{}, err
	}

	res := make([]*brigademodel.Build, len(builds))
	for i, build := range builds {
		bl := brigademodel.Build(*build)
		res[i] = &bl
	}

	// Order builds in descending order (last ones first).
	sort.Slice(res, func(i, j int) bool {
		// If no data move at the end of the list.
		if res[i].Worker == nil || res[j].Worker == nil {
			if desc {
				return true
			}
			return false
		}

		if desc {
			return res[i].Worker.StartTime.After(res[j].Worker.StartTime)
		}
		return res[i].Worker.StartTime.Before(res[j].Worker.StartTime)
	})

	return res, nil
}

// GetBuildJobs satisfies Service interface.
func (s *service) GetBuildJobs(BuildID string, desc bool) ([]*brigademodel.Job, error) {
	bl, err := s.client.GetBuild(BuildID)
	if err != nil {
		return []*brigademodel.Job{}, err
	}

	jobs, err := s.client.GetBuildJobs(bl)
	if err != nil {
		return []*brigademodel.Job{}, err
	}
	res := make([]*brigademodel.Job, len(jobs))
	for i, job := range jobs {
		j := brigademodel.Job(*job)
		res[i] = &j
	}

	// Order jobs in ascending order (first ones first).
	sort.Slice(res, func(i, j int) bool {
		if desc {
			return res[i].StartTime.After(res[j].StartTime)
		}
		return res[i].StartTime.Before(res[j].StartTime)

	})

	return res, nil
}

func (s *service) GetJob(jobID string) (*brigademodel.Job, error) {
	j, err := s.client.GetJob(jobID)

	if err != nil {
		return nil, err
	}
	res := brigademodel.Job(*j)
	return &res, nil
}

// GetJobLog satisfies Service interface.
func (s *service) GetJobLog(jobID string) (string, error) {
	job, err := s.client.GetJob(jobID)
	if err != nil {
		return "", err
	}

	str, err := s.client.GetJobLog(job)
	if err != nil {
		return "", err
	}

	return str, nil
}

// GetJobLog satisfies Service interface.
func (s *service) GetJobLogStream(jobID string) (io.ReadCloser, error) {
	job, err := s.client.GetJob(jobID)
	if err != nil {
		return nil, err
	}

	// TODO: Brigade doesn't set the follow (PR submited: https://github.com/Azure/brigade/pull/492)
	rc, err := s.client.GetJobLogStream(job)
	if err != nil {
		return nil, err
	}

	return rc, nil
}
