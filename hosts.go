
package jingli


import "time"


type HostId string

func (h HostId) String() string {
    return string(h)
}

// Host represents a host (physical or virtual) that can host LXC containers.
type Host struct {
    hostname        string
    host_id         HostId // stable / unique id for host
    private_network string // private network used for running Services
    cores           int
    memory          int
    last_updated    time.Time
}

func NewHost(host_id HostId, hostname, private_network string, cores, memory int, last_updated time.Time) (host *Host, err error) {
    host = new(Host)
    host.hostname = hostname
    host.host_id = host_id
    host.private_network = private_network
    host.cores = cores
    host.memory = memory
    host.last_updated = last_updated
    return host, nil
}

func (h *Host) HostId() HostId {
    return h.host_id
}

// get hotname of a host
func (h *Host) Hostname() string {
    return h.hostname
}

// get the private network that containers use on the host
func (h *Host) PrivateNetwork() string {
    return h.private_network
}


func (h *Host) Cores() int {
    return h.cores
}

func (h *Host) Memory() int {
    return h.memory
}

func (h *Host) LastUpdated() time.Time {
    return h.last_updated
}

func (h *Host) String() string {
    return h.hostname
}
