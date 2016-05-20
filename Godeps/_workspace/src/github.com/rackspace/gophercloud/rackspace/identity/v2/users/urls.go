package users

import "github.com/suonto/gophercloud"

func resetAPIKeyURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("users", id, "OS-KSADM", "credentials", "RAX-KSKEY:apiKeyCredentials", "RAX-AUTH", "reset")
}
