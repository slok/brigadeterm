package controller

import "time"

type fake struct{}

// NewFakeController returns a new faked controller.
func NewFakeController() Controller {
	return &fake{}
}

func (f *fake) ProjectListPageContext() *ProjectListPageContext {
	return &ProjectListPageContext{
		Projects: []*project{
			&project{
				ID:                 "1",
				Name:               "company1/AAAAAA",
				LastBuildOK:        true,
				LastBuildEventType: "github:push",
				LastBuildVersion:   "fb3cfc6844e52c677354660d503177fd9d4f7119",
				LastBuildTime:      time.Now().Add(-23 * time.Hour),
			},
			&project{
				ID:                 "2",
				Name:               "company1/123456",
				LastBuildOK:        true,
				LastBuildEventType: "github:pull_request",
				LastBuildVersion:   "1646bd72c15156fe3df95719e9bd91fc00126d2e",
				LastBuildTime:      time.Now().Add(-5 * time.Minute),
			},
			&project{
				ID:                 "3",
				Name:               "company2/987652",
				LastBuildOK:        true,
				LastBuildEventType: "github:push",
				LastBuildVersion:   "c4d31b44f9f21f40d40061f5371f02859b07fc5c",
				LastBuildTime:      time.Now().Add(-12 * time.Hour).Add(-35 * time.Minute),
			},
			&project{
				ID:                 "4",
				Name:               "company3/123sads",
				LastBuildOK:        true,
				LastBuildEventType: "github:pull_request",
				LastBuildVersion:   "1b44f9f21f401f5371f02859b07f1fxsdczfc5c",
				LastBuildTime:      time.Now().Add(-1 * time.Minute).Add(-12 * time.Second),
			},
			&project{
				ID:                 "5",
				Name:               "company4/3fg1",
				LastBuildOK:        false,
				LastBuildEventType: "github:push",
				LastBuildVersion:   "5371f028519b07c49f2161ffd3f40db44fc125c",
				LastBuildTime:      time.Now().Add(-120 * time.Hour).Add(-19 * time.Minute),
			},
			&project{
				ID:                 "6",
				Name:               "company4/1234567566345",
				LastBuildOK:        true,
				LastBuildEventType: "custom_gateway:deploy",
				LastBuildVersion:   "c5ccf40d74d31b534f9f211f02859b0440061f7f",
				LastBuildTime:      time.Now().Add(-1 * time.Hour),
			},
			&project{
				ID:                 "7",
				Name:               "company5/423df",
				LastBuildOK:        true,
				LastBuildEventType: "docker:push",
				LastBuildVersion:   "1f40d7fc5cc4d31b41f02859b04f9f240061f537",
				LastBuildTime:      time.Now().Add(-10 * time.Second),
			},
			&project{
				ID:                 "8",
				Name:               "company1/ggasdasft",
				LastBuildOK:        false,
				LastBuildEventType: "github:push",
				LastBuildVersion:   "71f028b44f9f21f59b0c4d3140d40061f537fc5c",
				LastBuildTime:      time.Now().Add(-9999 * time.Hour),
			},
		},
	}
}

func (f *fake) ProjectBuildListPageContext(projectID string) *ProjectBuildListPageContext {
	return &ProjectBuildListPageContext{
		ProjectName: "company1/AAAAAA",
		ProjectNS:   "ci",
		ProjectURL:  "git@github.com:slok/brigadeterm",
		Builds: []*build{
			&build{
				ID:         "lkjdsbfdflkdsnflkjdsbflkjadbflkjaful",
				Version:    "3140d400028b44f9f21f597b0c4d61f537fc51fc",
				Running:    false,
				FinishedOK: true,
				EventType:  "github:push",
				Started:    time.Now().Add(-9999 * time.Hour),
				Ended:      time.Now().Add(-9998 * time.Hour),
			},
			&build{
				ID:         "flkjdsbfuldflkdsnflkjdsbflkjadbflkja",
				Version:    "1f537fc5c1f028b44f9f274d3140d40061f59b0c",
				Running:    false,
				FinishedOK: true,
				EventType:  "deploy",
				Started:    time.Now().Add(-1 * time.Hour),
				Ended:      time.Now().Add(-50 * time.Minute),
			},
			&build{
				ID:         "24351321ldflkds32kjdsbflkj323dbflkja",
				Version:    "b44f9f21f59b7161f537fc5cf0280c4d3140d400",
				Running:    false,
				FinishedOK: false,
				EventType:  "github:push",
				Started:    time.Now().Add(-120 * time.Hour).Add(-19 * time.Minute),
				Ended:      time.Now().Add(-120 * time.Hour).Add(-18 * time.Minute),
			},
			&build{
				ID:         "2oijohpobna123213eewfeflkj323dbflkja",
				Version:    "061f537fc5c71f0d3140d4028b44f9f21f59b0c4",
				Running:    false,
				FinishedOK: true,
				EventType:  "github:push",
				Started:    time.Now().Add(-5 * time.Minute),
				Ended:      time.Now().Add(-128 * time.Second),
			},
			&build{
				ID:        "oijohpobna123213eewfeflkj323dbfl2kja",
				Version:   "244f9f0d40537fc5c1f59061fb0c4d31471f028b",
				Running:   true,
				EventType: "github:pull_reqest",
				Started:   time.Now().Add(-30 * time.Second),
			},
		},
	}
}

func (f *fake) BuildJobListPageContext(buildID string) *BuildJobListPageContext {
	return &BuildJobListPageContext{
		BuildInfo: &build{
			ID:         "2oijohpobna123213eewfeflkj323dbflkja",
			Version:    "061f537fc5c71f0d3140d4028b44f9f21f59b0c4",
			Running:    false,
			FinishedOK: true,
			EventType:  "github:push",
			Started:    time.Now().Add(-5 * time.Minute),
			Ended:      time.Now().Add(-128 * time.Second),
		},
		Jobs: []*job{
			&job{
				ID:         "1234567890",
				Name:       "build-job-1",
				Running:    false,
				FinishedOK: true,
				Started:    time.Now().Add(-11 * time.Minute),
				Ended:      time.Now().Add(-9 * time.Minute),
			},
			&job{
				ID:         "1234567890",
				Name:       "build-job-2",
				Running:    false,
				FinishedOK: true,
				Started:    time.Now().Add(-9 * time.Minute),
				Ended:      time.Now().Add(-5 * time.Minute),
			},
			&job{
				ID:         "1234567890",
				Name:       "build-job-3",
				Running:    false,
				FinishedOK: true,
				Started:    time.Now().Add(-9 * time.Minute),
				Ended:      time.Now().Add(-5 * time.Minute),
			},
			&job{
				ID:         "1234567890",
				Name:       "build-job-4",
				Running:    false,
				FinishedOK: true,
				Started:    time.Now().Add(-9 * time.Minute),
				Ended:      time.Now().Add(-3 * time.Minute),
			},
			&job{
				ID:         "1234567890",
				Name:       "build-job-5",
				Running:    false,
				FinishedOK: true,
				Started:    time.Now().Add(-3 * time.Minute),
				Ended:      time.Now().Add(-1 * time.Minute),
			},
		},
	}
}
