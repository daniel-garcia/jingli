
package jingli

import "testing"
import "time"

func TestAddHost(t *testing.T) {
    controller, err := NewServiceController("cp/root/")
    if err != nil {
        t.Errorf("Failed to create service controller: %s", err)
    }
    hosts, err := controller.Hosts()
    if err != nil {
        t.Errorf("Could not retrive hosts: %s", err)
    }
    if len(hosts) > 0 {
        t.Errorf("Got %d hosts, expected 0", len(hosts))
    }
    // add a host
    now := time.Now()
    host, _ := NewHost("localhost", "007f0101", "172.12.14.0/24", 24, 1000000000, now)
    err = controller.AddHost(host)
    if err != nil {
        t.Errorf("Got a error after adding a host %s", err)
    }

    hosts, err = controller.Hosts()
    if err != nil {
        t.Errorf("Got a error after inserting a host: %s", err)
    }

    found := false
    hostid := host.HostId()
    t.Logf("Trying to find host[%b]", hostid)
    for _, h := range(hosts) {
        t.Logf("Found host[%b]", h.HostId())
        if hostid == h.HostId() {
            found = true
            break
        }
    }
    if found == false {
        t.Errorf("Could not find host.")
    }
}

