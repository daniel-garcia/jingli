package jingli


import (
    "database/sql"
    _ "github.com/ziutek/mymysql/godrv"
    "time"
)

// The entry point to the control plane.
type ServiceController struct {
    instance_id       string
    name              string
    connection_string string
}


// Create a connection to a ServiceController
func NewServiceController(connection_string string) (controller *ServiceController, err error){
    controller = new(ServiceController)
    controller.connection_string = connection_string
    con, err := sql.Open("mymysql", controller.connection_string)
    if err != nil {
        return controller, err
    }
    defer con.Close()

    return controller, nil
}


// Return all the hosts associated with this controller
func (controller *ServiceController) Hosts() (hosts []*Host, err error){
    con, err := sql.Open("mymysql", controller.connection_string)
    if err != nil {
        return hosts, err
    }
    defer con.Close()

    rows, err := con.Query("select hostid, hostname, private_network, cores, memory, last_updated from host")
    if err != nil {
        return hosts, err
    }
    hosts = make([]*Host, 0, 10)
    var hostid, hostname, private_network string
    var last_updated time.Time
    var cores, memory int
    for rows.Next() {
        err = rows.Scan(&hostid, &hostname, &private_network, &cores, &memory, &last_updated)
        if err != nil { return hosts, err }
        hosts = append(hosts, &Host{hostid, hostname, private_network, cores, memory, last_updated})
    }
    return hosts, nil
}


// Add a new host a controller
func (controller *ServiceController) AddHost(host *Host) (err error){
    con, err := sql.Open("mymysql", controller.connection_string)
    if err != nil {
        return err
    }
    defer con.Close()

    _, err = con.Exec("INSERT INTO host ( hostid, hostname, private_network, cores, memory, last_updated) VALUES (?, ?, ?, ?, ?, ?)",
        host.HostId(), host.Hostname(), host.PrivateNetwork(), host.Cores(), host.Memory(), host.LastUpdated())
    return err
}


