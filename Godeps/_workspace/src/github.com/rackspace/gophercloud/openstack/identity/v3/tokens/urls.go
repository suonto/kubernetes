package tokens

import "github.com/suonto/gophercloud"

func tokenURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
