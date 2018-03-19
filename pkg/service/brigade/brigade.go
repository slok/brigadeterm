package brigade

import (
	"sort"

	"github.com/Azure/brigade/pkg/storage"

	brigademodel "github.com/slok/brigadeterm/pkg/model/brigade"
)

// Service is the service where all the brigade data will be retrieved
// and prepared with the information that this applications needs.
type Service interface {
	// GetAllProjects will get all the projects that are on brigade.
	GetAllProjects() ([]*brigademodel.Project, error)
	//// GetProjectBuilds will get all the builds of a project.
	//GetProjectBuilds(projectID string) ([]*brigademodel.Build, error)
	//// GetBuildJobs will get all the jobs of a build.
	//GetBuildJobs(BuildID string) ([]*brigademodel.Job, error)
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

// GetAllProjects satisfies Repository interface.
func (s *service) GetAllProjects() ([]*brigademodel.Project, error) {
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

// // GetAllProjects satisfies Repository interface.
// func (r *repository) GetProjectBuilds(projectID string) ([]*brigademodel.Build, error) {
// 	pr, err := r.client.GetProject(projectID)
// 	if err != nil {
// 		return []*brigademodel.Build{}, err
// 	}

// 	builds, err := r.client.GetProjectBuilds(pr)
// 	if err != nil {
// 		return []*brigademodel.Build{}, err
// 	}

// 	res := make([]*brigademodel.Build, len(builds))
// 	for i, build := range builds {
// 		bl := brigademodel.Build(*build)
// 		res[i] = &bl
// 	}

// 	return res, nil
// }

// // GetBuildJobs satisfies Repository interface.
// func (r *repository) GetBuildJobs(BuildID string) ([]*brigademodel.Job, error) {
// 	bl, err := r.client.GetBuild(BuildID)
// 	if err != nil {
// 		return []*brigademodel.Job{}, err
// 	}

// 	jobs, err := r.client.GetBuildJobs(bl)
// 	if err != nil {
// 		return []*brigademodel.Job{}, err
// 	}
// 	res := make([]*brigademodel.Job, len(jobs))
// 	for i, job := range jobs {
// 		j := brigademodel.Job(*job)
// 		res[i] = &j
// 	}

// 	return res, nil
// }

// func (r *repository) transformJobStatusToBuildStatus(status brigade.JobStatus) brigademodel.BuildStatus {
// 	switch status {
// 	case brigade.JobRunning:
// 		return brigademodel.BuildStatusRunning
// 	case brigade.JobSucceeded:
// 		return brigademodel.BuildStatusSucceeded
// 	case brigade.JobFailed:
// 		return brigademodel.BuildStatusFailed
// 	case brigade.JobUnknown:
// 		return brigademodel.BuildStatusUnknown
// 	case brigade.JobPending:
// 		return brigademodel.BuildStatusPending
// 	}

// 	return brigademodel.BuildStatusUnknown
// }
