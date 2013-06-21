
package jingli


import "time"

// Host represents a host (physical or virtual) that can host LXC containers.
type Host struct {
    hostname        string
    host_id         string // stable / unique id for host
    private_network string // private network used for running Services
    cores           int
    memory          int
    last_updated    time.Time
}

func NewHost(hostname, host_id, private_network string, core, memory int, last_updated time.Time) (host *Host, err error) {
    host = new(Host)
    host.hostname = hostname
    host.host_id = host_id
    host.private_network = private_network
    host.cores = cores
    host.memory
}

func (h *Host) HostId() string {
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