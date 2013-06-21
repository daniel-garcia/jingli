/*****************************************************************************
* 
* Copyright (C) Zenoss, Inc. 2013, all rights reserved.
* 
* This content is made available according to terms specified in
* LICENSE under the directory where this product is installed.
* 
*****************************************************************************/

package jingli

import "time"

// A ServiceGroup is a logical grouping of services. For example, all the
// services required to deliver a website (the webserver, database, memcached...etc)
// would be considred part of a single ServiceGroup.
type ServiceGroup struct {
	Name           string // User-friendly name for the ServiceGroup
	Id             string // Unique ID for ServiceGroup
	ResourcePoolId string // Unique ID for ResourcePool to run services in
	Priority       int    // relative priority of ServiceGroup (-19..20)
}

// A Port is a definition of a network port, the protocol (tcp/udp),
// the application protocol, and the ServiceGroup the Port belongs to.
type Port struct {
	Port           uint16 // The port number (1 to 65535)
	Protocol       string // tcp/udp
	Application    string // eg http, ssh, smtp...
	ServiceGroupId string // ServiceGroup.ID
}

// A ServicePort is a mapping of a Service to a Port.
type ServicePort struct {
	ServiceId string // unique ID for a Service
	port      Port   // the Port to map to the Service
}

// A Service describes how to launch an application, the endpoint ports it exposes,
// and the ServicePorts it depends on.
type Service struct {
	Id               string          // unique ID for a Service
	GroupId          string          // ServiceGroup.Id that this service participates in
	StartupTemplate  string          // template for script used to start this service
	ShutdownTemplate string          // template for script used to stop this service
	EndPoints        map[uint16]Port // mapping of exposed port
	ServicePorts     []ServicePort   // list of dependent ServicePorts
	ephemeral        bool            // autostart service flag
	image_id         string          // unique id of file system image used
	resource_pool_id string          // resource pool id if it exists
	priority         int             // relative priority of Service (-19..20)
}

// Does the service start when the Service Group is started?
func (s *Service) Ephemeral() bool {
	return s.ephemeral
}

// Return the filesystem image associated with the Service.
func (s *Service) ImageId() string {
	return s.image_id
}

// Return the ResourcePool.Id for this Service
func (s *Service) ResourcePoolId() string {
	return s.resource_pool_id
}

// Return the relative priority of the Service
func (s *Service) Priority() int {
	return s.priority
}

// Container represnts the LXC container that docker launched on a host.
type Container struct {
	docker_id   string
	host_id     string
	ip_addr     string
	launched_at time.Time
}

// Return the docker Id associated with this container
func (c *Container) DockerId() string {
	return c.docker_id
}

// Return the hostid of the host running this container.
func (c *Container) HostId() string {
	return c.host_id
}

// Return the IP address of the container.
func (c *Container) IpAddr() string {
	return c.ip_addr
}

// Return the time.Time the container was launched.
func (c *Container) StartedAt() time.Time {
	return c.launched_at
}

// Host represents a host (physical or virtual) that can host LXC containers.
type Host struct {
	hostname        string
	host_id         string // stable / unique id for host
	private_network string // private network used for running Services
}

// get hotname of a host
func (h *Host) Hostname() string {
	return h.hostname
}

// get the private network that containers use on the host
func (h *Host) PrivateNetwork() string {
	return h.private_network
}

// ResourcePool represents a set of resouce restrictions and the hosts that participate in it.
type ResourcePool struct {
	id       string   // unique id for ResourcePool
	name     string   // friendly resource pool name
	cores    int      // number of cores, -1 unlimited
	memory   int      // bytes of memory, -1 unlimited
	priority int      // relative priority of pool (-19..20)
	hosts    []string // list of hosts that participate in this resource pool
}

// Return the core limit for the resource pool. -1 if no limit.
func (p *ResourcePool) Cores() int {
	return p.cores
}

// Return the memory limit for the resource pool. -1 if no limit.
func (p *ResourcePool) Memory() int {
	return p.memory
}

// Return the priority of services that run the ResourcePool
func (p *ResourcePool) Priority() int {
	return p.priority
}

// Return the hosts that participate in the resource pool
func (p *ResourcePool) Hosts() []string {
	return p.hosts
}

// The entry point to the control plane.
type ServiceController struct {
	instance_id       string
	name              string
	connection_string string
}
