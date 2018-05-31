package controller

import (
	"fmt"
	"io"
	"strings"
	"time"
)

var (
	staticNow = time.Now()
)

type fake struct{}

// NewFakeController returns a new faked controller.
func NewFakeController() Controller {
	return &fake{}
}

func (f *fake) ProjectListPageContext() *ProjectListPageContext {
	return &ProjectListPageContext{
		Projects: []*Project{
			&Project{
				ID:   "1",
				Name: "company1/AAAAAA",
				LastBuilds: []*Build{
					&Build{
						ID:        "lkjdsbfdflkdsnflkjdsbflkjadbflkjaful",
						Version:   "3140d400028b44f9f21f597b0c4d61f537fc51fc",
						State:     SuccessedState,
						EventType: "github:push",
						Started:   staticNow.Add(-9999 * time.Hour),
						Ended:     time.Now().Add(-9998 * time.Hour),
					},
					&Build{State: FailedState},
					&Build{State: SuccessedState},
					&Build{State: SuccessedState},
					&Build{State: FailedState},
				},
			},
			&Project{
				ID:   "2",
				Name: "company1/123456",
				LastBuilds: []*Build{
					&Build{
						ID:        "flkjdsbfuldflkdsnflkjdsbflkjadbflkja",
						Version:   "1f537fc5c1f028b44f9f274d3140d40061f59b0c",
						State:     SuccessedState,
						EventType: "deploy",
						Started:   staticNow.Add(-1 * time.Hour),
						Ended:     time.Now().Add(-50 * time.Minute),
					},
					&Build{State: SuccessedState},
					&Build{State: SuccessedState},
				},
			},
			&Project{
				ID:   "3",
				Name: "company2/987652",
				LastBuilds: []*Build{
					&Build{
						ID:        "24351321ldflkds32kjdsbflkj323dbflkja",
						Version:   "b44f9f21f59b7161f537fc5cf0280c4d3140d400",
						State:     FailedState,
						EventType: "github:push",
						Started:   staticNow.Add(-120 * time.Hour).Add(-19 * time.Minute),
						Ended:     time.Now().Add(-120 * time.Hour).Add(-18 * time.Minute),
					},
					&Build{State: FailedState},
					&Build{State: UnknownState},
					&Build{State: SuccessedState},
					&Build{State: UnknownState},
				},
			},
			&Project{
				ID:   "4",
				Name: "company3/123sads",
				LastBuilds: []*Build{
					&Build{
						ID:        "2oijohpobna123213eewfeflkj323dbflkja",
						Version:   "061f537fc5c71f0d3140d4028b44f9f21f59b0c4",
						State:     PendingState,
						EventType: "github:push",
						Started:   staticNow.Add(-5 * time.Minute),
						Ended:     time.Now().Add(-128 * time.Second),
					},
					&Build{State: SuccessedState},
					&Build{State: SuccessedState},
					&Build{State: FailedState},
					&Build{State: SuccessedState},
				},
			},
			&Project{
				ID:   "5",
				Name: "company4/3fg1",
				LastBuilds: []*Build{
					&Build{
						ID:        "oijohpobna123213eewfeflkj323dbfl2kja",
						Version:   "244f9f0d40537fc5c1f59061fb0c4d31471f028b",
						State:     RunningState,
						EventType: "github:pull_reqest",
						Started:   time.Now().Add(-30 * time.Second),
					},
					&Build{State: FailedState},
					&Build{State: FailedState},
					&Build{State: FailedState},
					&Build{State: FailedState},
				},
			},
			&Project{
				ID:   "6",
				Name: "company4/1234567566345",
				LastBuilds: []*Build{
					&Build{
						ID:        "flkjdsbfuldflkdsnflkjdsbflkjadbflkja",
						Version:   "1f537fc5c1f028b44f9f274d3140d40061f59b0c",
						State:     SuccessedState,
						EventType: "deploy",
						Started:   staticNow.Add(-1 * time.Hour),
						Ended:     time.Now().Add(-50 * time.Minute),
					},
					&Build{State: FailedState},
					&Build{State: FailedState},
					&Build{State: FailedState},
					&Build{State: UnknownState},
				},
			},
			&Project{
				ID:   "7",
				Name: "company5/423df",
				LastBuilds: []*Build{
					&Build{
						ID:        "24351321ldflkds32kjdsbflkj323dbflkja",
						Version:   "b44f9f21f59b7161f537fc5cf0280c4d3140d400",
						State:     UnknownState,
						EventType: "github:push",
						Started:   staticNow.Add(-120 * time.Hour).Add(-19 * time.Minute),
						Ended:     time.Now().Add(-120 * time.Hour).Add(-18 * time.Minute),
					},
					&Build{State: SuccessedState},
					&Build{State: SuccessedState},
					&Build{State: SuccessedState},
					&Build{State: SuccessedState},
				},
			},
			&Project{
				ID:   "8",
				Name: "company1/ggasdasft",
				LastBuilds: []*Build{
					&Build{
						ID:        "2oijohpobna123213eewfeflkj323dbflkja",
						Version:   "061f537fc5c71f0d3140d4028b44f9f21f59b0c4",
						State:     SuccessedState,
						EventType: "github:push",
						Started:   staticNow.Add(-5 * time.Minute),
						Ended:     time.Now().Add(-128 * time.Second),
					},
					&Build{State: FailedState},
					&Build{State: SuccessedState},
					&Build{State: SuccessedState},
					&Build{State: SuccessedState},
				},
			},
			&Project{
				ID:   "8",
				Name: "company1/0184848q1danfubu<s",
			},
			nil,
		},
	}
}

func (f *fake) ProjectBuildListPageContext(projectID string) *ProjectBuildListPageContext {
	return &ProjectBuildListPageContext{
		ProjectName: "company1/AAAAAA",
		ProjectNS:   "ci",
		ProjectURL:  "git@github.com:slok/brigadeterm",
		Builds: []*Build{
			&Build{
				ID:        "lkjdsbfdflkdsnflkjdsbflkjadbflkjaful",
				Version:   "3140d400028b44f9f21f597b0c4d61f537fc51fc",
				State:     SuccessedState,
				EventType: "github:push",
				Started:   staticNow.Add(-9999 * time.Hour),
				Ended:     time.Now().Add(-9998 * time.Hour),
			},
			&Build{
				ID:        "flkjdsbfuldflkdsnflkjdsbflkjadbflkja",
				Version:   "1f537fc5c1f028b44f9f274d3140d40061f59b0c",
				State:     FailedState,
				EventType: "deploy",
				Started:   staticNow.Add(-1 * time.Hour),
				Ended:     time.Now().Add(-50 * time.Minute),
			},
			&Build{
				ID:        "24351321ldflkds32kjdsbflkj323dbflkja",
				Version:   "b44f9f21f59b7161f537fc5cf0280c4d3140d400",
				State:     PendingState,
				EventType: "github:push",
				Started:   staticNow.Add(-120 * time.Hour).Add(-19 * time.Minute),
				Ended:     time.Now().Add(-120 * time.Hour).Add(-18 * time.Minute),
			},
			&Build{
				ID:        "2oijohpobna123213eewfeflkj323dbflkja",
				Version:   "061f537fc5c71f0d3140d4028b44f9f21f59b0c4",
				State:     UnknownState,
				EventType: "github:push",
				Started:   staticNow.Add(-5 * time.Minute),
				Ended:     time.Now().Add(-128 * time.Second),
			},
			&Build{
				ID:        "oijohpobna123213eewfeflkj323dbfl2kja",
				Version:   "244f9f0d40537fc5c1f59061fb0c4d31471f028b",
				State:     RunningState,
				EventType: "github:pull_reqest",
				Started:   staticNow.Add(-30 * time.Second),
			},
			&Build{},
			nil,
		},
	}
}

func (f *fake) BuildJobListPageContext(buildID string) *BuildJobListPageContext {
	return &BuildJobListPageContext{
		BuildInfo: &Build{
			ID:        "2oijohpobna123213eewfeflkj323dbflkja",
			Version:   "061f537fc5c71f0d3140d4028b44f9f21f59b0c4",
			State:     SuccessedState,
			EventType: "github:push",
			Started:   time.Now().Add(-5 * time.Minute),
			Ended:     time.Now().Add(-128 * time.Second),
		},
		Jobs: []*Job{
			&Job{
				ID:      "unit-test-01c8zehre13ht12776hdkms8gf",
				Name:    "unit-test",
				Image:   "golang:1.9",
				State:   FailedState,
				Started: staticNow.Add(-11 * time.Minute),
				Ended:   time.Now().Add(-9 * time.Minute),
			},
			&Job{
				ID:      "build-binary-1-01c8zehre13ht12776hdkms8gf",
				Name:    "build-binary-1",
				Image:   "docker:stable-dind",
				State:   RunningState,
				Started: staticNow.Add(-9 * time.Minute),
				Ended:   time.Now().Add(-5 * time.Minute),
			},
			&Job{
				ID:      "build-binary-1-01c8zehre13ht12776hdkms8gf",
				Name:    "build-binary-2",
				Image:   "docker:stable-dind",
				State:   PendingState,
				Started: staticNow.Add(-9 * time.Minute),
				Ended:   time.Now().Add(-5 * time.Minute),
			},
			&Job{
				ID:      "build-binary-3-01c8zehre13ht12776hdkms8gf",
				Name:    "build-binary-3",
				Image:   "docker:stable-dind",
				State:   UnknownState,
				Started: staticNow.Add(-9 * time.Minute),
				Ended:   time.Now().Add(-3 * time.Minute),
			},
			&Job{
				ID:      "set-github-build-status-01c8zehre13ht12776hdkms8gf",
				Name:    "set-github-build-status",
				Image:   "technosophos/github-notify:latest",
				State:   RunningState,
				Started: staticNow.Add(-3 * time.Minute),
				Ended:   time.Now().Add(-1 * time.Minute),
			},
			nil,
		},
	}
}

func (f *fake) JobLogPageContext(jobID string) *JobLogPageContext {

	log := fmt.Sprintf(`
%v
=========

time="2018-03-19T15:38:46.631153420Z" level=warning msg="could not change group /var/run/docker.sock to docker: group docker not found"
time="2018-03-19T15:38:46.631266657Z" level=warning msg="[!] DON'T BIND ON ANY IP ADDRESS WITHOUT setting --tlsverify IF YOU DON'T KNOW WHAT YOU'RE DOING [!]"
time="2018-03-19T15:38:46.632042877Z" level=info msg="libcontainerd: started new docker-containerd process" pid=25
time="2018-03-19T15:38:46Z" level=info msg="starting containerd" module=containerd revision=9b55aab90508bd389d7654c4baf173a981477d55 version=v1.0.1 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.content.v1.content"..." module=containerd type=io.containerd.content.v1 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.snapshotter.v1.btrfs"..." module=containerd type=io.containerd.snapshotter.v1 
time="2018-03-19T15:38:46Z" level=warning msg="failed to load plugin io.containerd.snapshotter.v1.btrfs" error="path /var/lib/docker/containerd/daemon/io.containerd.snapshotter.v1.btrfs must be a btrfs filesystem to be used with the btrfs snapshotter" module=containerd 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.snapshotter.v1.overlayfs"..." module=containerd type=io.containerd.snapshotter.v1 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.metadata.v1.bolt"..." module=containerd type=io.containerd.metadata.v1 
time="2018-03-19T15:38:46Z" level=warning msg="could not use snapshotter btrfs in metadata plugin" error="path /var/lib/docker/containerd/daemon/io.containerd.snapshotter.v1.btrfs must be a btrfs filesystem to be used with the btrfs snapshotter" module="containerd/io.containerd.metadata.v1.bolt" 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.differ.v1.walking"..." module=containerd type=io.containerd.differ.v1 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.gc.v1.scheduler"..." module=containerd type=io.containerd.gc.v1 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.grpc.v1.containers"..." module=containerd type=io.containerd.grpc.v1 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.grpc.v1.content"..." module=containerd type=io.containerd.grpc.v1 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.grpc.v1.diff"..." module=containerd type=io.containerd.grpc.v1 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.grpc.v1.events"..." module=containerd type=io.containerd.grpc.v1 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.grpc.v1.healthcheck"..." module=containerd type=io.containerd.grpc.v1 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.grpc.v1.images"..." module=containerd type=io.containerd.grpc.v1 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.grpc.v1.leases"..." module=containerd type=io.containerd.grpc.v1 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.grpc.v1.namespaces"..." module=containerd type=io.containerd.grpc.v1 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.grpc.v1.snapshots"..." module=containerd type=io.containerd.grpc.v1 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.monitor.v1.cgroups"..." module=containerd type=io.containerd.monitor.v1 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.runtime.v1.linux"..." module=containerd type=io.containerd.runtime.v1 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.grpc.v1.tasks"..." module=containerd type=io.containerd.grpc.v1 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.grpc.v1.version"..." module=containerd type=io.containerd.grpc.v1 
time="2018-03-19T15:38:46Z" level=info msg="loading plugin "io.containerd.grpc.v1.introspection"..." module=containerd type=io.containerd.grpc.v1 
time="2018-03-19T15:38:46Z" level=info msg=serving... address="/var/run/docker/containerd/docker-containerd-debug.sock" module="containerd/debug" 
time="2018-03-19T15:38:46Z" level=info msg=serving... address="/var/run/docker/containerd/docker-containerd.sock" module="containerd/grpc" 
time="2018-03-19T15:38:46Z" level=info msg="containerd successfully booted in 0.011531s" module=containerd 
time="2018-03-19T15:38:46.656109039Z" level=info msg="Setting the storage driver from the $DOCKER_DRIVER environment variable (overlay)"
time="2018-03-19T15:38:46.708198389Z" level=info msg="Graph migration to content-addressability took 0.00 seconds"
time="2018-03-19T15:38:46.708414414Z" level=warning msg="Your kernel does not support swap memory limit"
time="2018-03-19T15:38:46.708501869Z" level=warning msg="Your kernel does not support cgroup rt period"
time="2018-03-19T15:38:46.708517450Z" level=warning msg="Your kernel does not support cgroup rt runtime"
time="2018-03-19T15:38:46.709142577Z" level=info msg="Loading containers: start."
time="2018-03-19T15:38:46.716968102Z" level=warning msg="Running modprobe bridge br_netfilter failed with message: ip: can't find device 'bridge'\nbridge                135168  1 br_netfilter\nstp                    16384  1 bridge\nllc                    16384  2 bridge,stp\nip: can't find device 'br_netfilter'\nbr_netfilter           24576  0 \nbridge                135168  1 br_netfilter\nmodprobe: can't change directory to '/lib/modules': No such file or directory\n, error: exit status 1"
time="2018-03-19T15:38:46.721850651Z" level=warning msg="Running modprobe nf_nat failed with message: ip: can't find device 'nf_nat'\nnf_nat_ipv6            16384  1 ip6table_nat\nnf_nat_masquerade_ipv4    16384  1 ipt_MASQUERADE\nnf_nat_ipv4            16384  1 iptable_nat\nnf_nat                 24576  4 nf_nat_ipv6,xt_nat,nf_nat_masquerade_ipv4,nf_nat_ipv4\nnf_conntrack          114688  8 nf_conntrack_ipv6,nf_nat_ipv6,nf_conntrack_netlink,nf_nat_masquerade_ipv4,nf_conntrack_ipv4,nf_nat_ipv4,xt_conntrack,nf_nat\nmodprobe: can't change directory to '/lib/modules': No such file or directory, error: exit status 1"
time="2018-03-19T15:38:46.725409453Z" level=warning msg="Running modprobe xt_conntrack failed with message: ip: can't find device 'xt_conntrack'\nxt_conntrack           16384 285 \nnf_conntrack          114688  8 nf_conntrack_ipv6,nf_nat_ipv6,nf_conntrack_netlink,nf_nat_masquerade_ipv4,nf_conntrack_ipv4,nf_nat_ipv4,xt_conntrack,nf_nat\nx_tables               36864 17 ip6_tables,xt_set,xt_multiport,iptable_mangle,iptable_raw,xt_statistic,xt_nat,xt_recent,ipt_REJECT,xt_tcpudp,xt_comment,xt_mark,ipt_MASQUERADE,xt_addrtype,iptable_filter,xt_conntrack,ip_tables\nmodprobe: can't change directory to '/lib/modules': No such file or directory, error: exit status 1"
time="2018-03-19T15:38:46.784168888Z" level=info msg="Default bridge (docker0) is assigned with an IP address 172.17.0.0/16. Daemon option --bip can be used to set a preferred IP address"
time="2018-03-19T15:38:46.815905782Z" level=info msg="Loading containers: done."
time="2018-03-19T15:38:46.821627496Z" level=info msg="Docker daemon" commit=7390fc6 graphdriver(s)=overlay version=17.12.1-ce
time="2018-03-19T15:38:46.821783403Z" level=info msg="Daemon has completed initialization"
time="2018-03-19T15:38:46.829954871Z" level=info msg="API listen on /var/run/docker.sock"
time="2018-03-19T15:38:46.829960185Z" level=info msg="API listen on [::]:2375"
+ IMAGE_VERSION=4a78618c57d613747176845bb067a0dfca91cea9
+ [ -z workspace-operator ]
+ REPOSITORY=dist.spotahome.net:5000/spotahome/
+ IMAGE=toilet-workspace-operator
+ docker build --build-arg operator=workspace-operator -t dist.spotahome.net:5000/spotahome/toilet-workspace-operator:4a78618c57d613747176845bb067a0dfca91cea9 -t dist.spotahome.net:5000/spotahome/toilet-workspace-operator:latest -f ./docker/prod/Dockerfile .
Sending build context to Docker daemon  58.49MB

Step 1/10 : FROM golang:1.9-alpine AS build-stage
1.9-alpine: Pulling from library/golang
605ce1bd3f31: Pulling fs layer
d1e3a1512603: Pulling fs layer
1c84e3e4a1a3: Pulling fs layer
71af08823d79: Pulling fs layer
c1b0d47bcf8b: Pulling fs layer
042f32f99403: Pulling fs layer
71af08823d79: Waiting
c1b0d47bcf8b: Waiting
042f32f99403: Waiting
1c84e3e4a1a3: Verifying Checksum
1c84e3e4a1a3: Download complete
d1e3a1512603: Verifying Checksum
d1e3a1512603: Download complete
605ce1bd3f31: Download complete
605ce1bd3f31: Pull complete
d1e3a1512603: Pull complete
c1b0d47bcf8b: Download complete
042f32f99403: Verifying Checksum
042f32f99403: Download complete
1c84e3e4a1a3: Pull complete
71af08823d79: Download complete
71af08823d79: Pull complete
c1b0d47bcf8b: Pull complete
042f32f99403: Pull complete
Digest: sha256:ab72eb6db0eda32d429becf8bb28f62081821cc1aa2d49430344083f87b1e6a2
Status: Downloaded newer image for golang:1.9-alpine
 ---> a8f345e387a3
Step 2/10 : ARG operator
 ---> Running in 40b42b26ce4c
time="2018-03-19T15:39:23.179384589Z" level=info msg="Layer sha256:2b9f47ab040a8d03e2bc7a33ed5dd2c8ef507bbb4804f55acbb549a52a1d8684 cleaned up"
Removing intermediate container 40b42b26ce4c
 ---> 136307e16e09
Step 3/10 : WORKDIR /go/src/github.com/spotahome/toilet
Removing intermediate container 7ae04c085e03
 ---> e71fdc7e690b
Step 4/10 : COPY . .
 ---> ba129e731b15
Step 5/10 : RUN ./build.sh ${operator}
 ---> Running in 409c3bb2593e
time="2018-03-19T15:39:27Z" level=info msg="shim docker-containerd-shim started" address="/containerd-shim/moby/409c3bb2593e5ae520605ffa8bfc16198172ba7befefcf71d28737ac7ca17cb1/shim.sock" debug=false module="containerd/tasks" pid=264 
Building workspace-operator operator
Built finished at ./bin/workspace-operator
time="2018-03-19T15:40:18Z" level=info msg="shim reaped" id=409c3bb2593e5ae520605ffa8bfc16198172ba7befefcf71d28737ac7ca17cb1 module="containerd/tasks" 
time="2018-03-19T15:40:18.674578337Z" level=info msg="ignoring event" module=libcontainerd namespace=moby topic=/tasks/delete type="*events.TaskDelete"
Removing intermediate container 409c3bb2593e
 ---> 30240ab3dc09
Step 6/10 : FROM alpine:latest
latest: Pulling from library/alpine
ff3a5c916c92: Pulling fs layer
ff3a5c916c92: Verifying Checksum
ff3a5c916c92: Download complete
ff3a5c916c92: Pull complete
Digest: sha256:7b848083f93822dd21b0a2f14a110bd99f6efb4b838d499df6d04a49d0debf8b
Status: Downloaded newer image for alpine:latest
 ---> 3fd9065eaf02
Step 7/10 : ARG operator
 ---> Running in 6843327bf16a
time="2018-03-19T15:40:22.579691821Z" level=info msg="Layer sha256:c332538a772dd349c01ca25f77abff46d6847fcabe79fe0ee31b195a2bace4cd cleaned up"
Removing intermediate container 6843327bf16a
 ---> 7a0e0abbc480
Step 8/10 : RUN apk --no-cache add   ca-certificates
 ---> Running in 954744345f59
time="2018-03-19T15:40:23Z" level=info msg="shim docker-containerd-shim started" address="/containerd-shim/moby/954744345f592d278f1b4a851ba704264a9a1b19b060b3fdd249c2b14e749940/shim.sock" debug=false module="containerd/tasks" pid=2566 
fetch http://dl-cdn.alpinelinux.org/alpine/v3.7/main/x86_64/APKINDEX.tar.gz
fetch http://dl-cdn.alpinelinux.org/alpine/v3.7/community/x86_64/APKINDEX.tar.gz
(1/1) Installing ca-certificates (20171114-r0)
Executing busybox-1.27.2-r7.trigger
Executing ca-certificates-20171114-r0.trigger
OK: 5 MiB in 12 packages
time="2018-03-19T15:40:24Z" level=info msg="shim reaped" id=954744345f592d278f1b4a851ba704264a9a1b19b060b3fdd249c2b14e749940 module="containerd/tasks" 
time="2018-03-19T15:40:24.594725461Z" level=info msg="ignoring event" module=libcontainerd namespace=moby topic=/tasks/delete type="*events.TaskDelete"
Removing intermediate container 954744345f59
 ---> a5dde4f63b3a
Step 9/10 : COPY --from=build-stage /go/src/github.com/spotahome/toilet/bin/${operator} /usr/local/bin/operator
 ---> 472339046426
Step 10/10 : ENTRYPOINT ["/usr/local/bin/operator"]
 ---> Running in 92ada1bc9b3a
time="2018-03-19T15:40:26.349771116Z" level=info msg="Layer sha256:c95e1daa54ef9263d4c5fe949e84432e0f665680ce64ae707a0c4475f8ba1fea cleaned up"
Removing intermediate container 92ada1bc9b3a
 ---> 2c7f00eab992
Successfully built 2c7f00eab992
Successfully tagged dist.spotahome.net:5000/spotahome/toilet-workspace-operator:4a78618c57d613747176845bb067a0dfca91cea9
Successfully tagged dist.spotahome.net:5000/spotahome/toilet-workspace-operator:latest
WARNING! Using --password via the CLI is insecure. Use --password-stdin.
Login Succeeded
The push refers to repository [dist.spotahome.net:5000/spotahome/toilet-workspace-operator]
3c84a9666104: Preparing
a986eb0e7a89: Preparing
cd7100a72410: Preparing
cd7100a72410: Layer already exists
a986eb0e7a89: Pushed
3c84a9666104: Pushed
4a78618c57d613747176845bb067a0dfca91cea9: digest: sha256:ac500411b48ff5ef31efb5d96b543d725dc97493a983cd2430de00d1c5221adb size: 949
3c84a9666104: Preparing
a986eb0e7a89: Preparing
cd7100a72410: Preparing
cd7100a72410: Layer already exists
a986eb0e7a89: Layer already exists
3c84a9666104: Layer already exists
latest: digest: sha256:ac500411b48ff5ef31efb5d96b543d725dc97493a983cd2430de00d1c5221adb size: 949
`, time.Now().UTC())

	// Create a pipe for our fake log.
	r, w := io.Pipe()

	// Stream the log
	go func() {
		defer w.Close()
		splLog := strings.Split(log, "\n")
		for _, line := range splLog {
			_, err := fmt.Fprintf(w, "%s\n", line)
			if err != nil {
				return // Something happenned, we don't mind if it's ended or not, stop.
			}
			sleepMS := time.Duration(time.Now().Nanosecond() % 1000)
			time.Sleep(sleepMS * time.Millisecond)
		}
	}()

	return &JobLogPageContext{
		Job: &Job{
			ID:      "build-binary-3-01c8zehre13ht12776hdkms8gf",
			Name:    "build-binary-3",
			Image:   "docker:stable-dind",
			State:   SuccessedState,
			Started: staticNow.Add(-9 * time.Minute),
			Ended:   time.Now().Add(-3 * time.Minute),
		},
		Log: r,
	}
}
