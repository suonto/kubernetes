package throttle

import (
	"strconv"

	"github.com/suonto/gophercloud"
)

const (
	path   = "loadbalancers"
	ctPath = "connectionthrottle"
)

func rootURL(c *gophercloud.ServiceClient, id int) string {
	return c.ServiceURL(path, strconv.Itoa(id), ctPath)
}
