package helper

type Config struct {
	Name        string
	Interactive bool
	Command     string
	Arguments   []string

	Binds           []string
	ContainerIDFile string
	//	LxcConf         []utils.KeyValuePair
	Privileged bool
	//	PortBindings    nat.PortMap
	Links           []string
	PublishAllPorts bool
	Dns             []string
	DnsSearch       []string
	ExtraHosts      []string
	VolumesFrom     []string
	//	Devices         []DeviceMapping
	//	NetworkMode     NetworkMode
	//	IpcMode         IpcMode
	CapAdd  []string
	CapDrop []string
	//	RestartPolicy   RestartPolicy
	SecurityOpt []string

	Hostname          string   `json:"Hostname,omitempty" yaml:"Hostname,omitempty" toml:"Hostname,omitempty"`
	Domainname        string   `json:"Domainname,omitempty" yaml:"Domainname,omitempty" toml:"Domainname,omitempty"`
	User              string   `json:"User,omitempty" yaml:"User,omitempty" toml:"User,omitempty"`
	Memory            int64    `json:"Memory,omitempty" yaml:"Memory,omitempty" toml:"Memory,omitempty"`
	MemorySwap        int64    `json:"MemorySwap,omitempty" yaml:"MemorySwap,omitempty" toml:"MemorySwap,omitempty"`
	MemoryReservation int64    `json:"MemoryReservation,omitempty" yaml:"MemoryReservation,omitempty" toml:"MemoryReservation,omitempty"`
	KernelMemory      int64    `json:"KernelMemory,omitempty" yaml:"KernelMemory,omitempty" toml:"KernelMemory,omitempty"`
	CPUShares         int64    `json:"CpuShares,omitempty" yaml:"CpuShares,omitempty" toml:"CpuShares,omitempty"`
	CPUSet            string   `json:"Cpuset,omitempty" yaml:"Cpuset,omitempty" toml:"Cpuset,omitempty"`
	PortSpecs         []string `json:"PortSpecs,omitempty" yaml:"PortSpecs,omitempty" toml:"PortSpecs,omitempty"`
	//	ExposedPorts      map[Port]struct{} `json:"ExposedPorts,omitempty" yaml:"ExposedPorts,omitempty" toml:"ExposedPorts,omitempty"`
	PublishService string   `json:"PublishService,omitempty" yaml:"PublishService,omitempty" toml:"PublishService,omitempty"`
	StopSignal     string   `json:"StopSignal,omitempty" yaml:"StopSignal,omitempty" toml:"StopSignal,omitempty"`
	StopTimeout    int      `json:"StopTimeout,omitempty" yaml:"StopTimeout,omitempty" toml:"StopTimeout,omitempty"`
	Env            []string `json:"Env,omitempty" yaml:"Env,omitempty" toml:"Env,omitempty"`
	Cmd            []string `json:"Cmd" yaml:"Cmd" toml:"Cmd"`
	Shell          []string `json:"Shell,omitempty" yaml:"Shell,omitempty" toml:"Shell,omitempty"`
	//	Healthcheck       *HealthConfig       `json:"Healthcheck,omitempty" yaml:"Healthcheck,omitempty" toml:"Healthcheck,omitempty"`
	DNS             []string            `json:"Dns,omitempty" yaml:"Dns,omitempty" toml:"Dns,omitempty"` // For Docker API v1.9 and below only
	Image           string              `json:"Image,omitempty" yaml:"Image,omitempty" toml:"Image,omitempty"`
	Volumes         map[string]struct{} `json:"Volumes,omitempty" yaml:"Volumes,omitempty" toml:"Volumes,omitempty"`
	VolumeDriver    string              `json:"VolumeDriver,omitempty" yaml:"VolumeDriver,omitempty" toml:"VolumeDriver,omitempty"`
	WorkingDir      string              `json:"WorkingDir,omitempty" yaml:"WorkingDir,omitempty" toml:"WorkingDir,omitempty"`
	MacAddress      string              `json:"MacAddress,omitempty" yaml:"MacAddress,omitempty" toml:"MacAddress,omitempty"`
	Entrypoint      []string            `json:"Entrypoint" yaml:"Entrypoint" toml:"Entrypoint"`
	SecurityOpts    []string            `json:"SecurityOpts,omitempty" yaml:"SecurityOpts,omitempty" toml:"SecurityOpts,omitempty"`
	OnBuild         []string            `json:"OnBuild,omitempty" yaml:"OnBuild,omitempty" toml:"OnBuild,omitempty"`
	Mounts          []Mount             `json:"Mounts,omitempty" yaml:"Mounts,omitempty" toml:"Mounts,omitempty"`
	Labels          map[string]string   `json:"Labels,omitempty" yaml:"Labels,omitempty" toml:"Labels,omitempty"`
	AttachStdin     bool                `json:"AttachStdin,omitempty" yaml:"AttachStdin,omitempty" toml:"AttachStdin,omitempty"`
	AttachStdout    bool                `json:"AttachStdout,omitempty" yaml:"AttachStdout,omitempty" toml:"AttachStdout,omitempty"`
	AttachStderr    bool                `json:"AttachStderr,omitempty" yaml:"AttachStderr,omitempty" toml:"AttachStderr,omitempty"`
	ArgsEscaped     bool                `json:"ArgsEscaped,omitempty" yaml:"ArgsEscaped,omitempty" toml:"ArgsEscaped,omitempty"`
	Tty             bool                `json:"Tty,omitempty" yaml:"Tty,omitempty" toml:"Tty,omitempty"`
	OpenStdin       bool                `json:"OpenStdin,omitempty" yaml:"OpenStdin,omitempty" toml:"OpenStdin,omitempty"`
	StdinOnce       bool                `json:"StdinOnce,omitempty" yaml:"StdinOnce,omitempty" toml:"StdinOnce,omitempty"`
	NetworkDisabled bool                `json:"NetworkDisabled,omitempty" yaml:"NetworkDisabled,omitempty" toml:"NetworkDisabled,omitempty"`
}

func NewConfig() *Config {
	out := Config{}
	return &out
}

type Mount struct {
	Name        string
	Source      string
	Destination string
	Driver      string
	Mode        string
	RW          bool
}

type PortMap map[string][]PortBinding

type PortBinding struct {
	// HostIP is the host IP Address
	HostIP string `json:"HostIp"`
	// HostPort is the host port number
	HostPort string
}
