package ipam

import (
	"context"

	"github.com/digitalocean/go-netbox/netbox"
	"github.com/digitalocean/go-netbox/netbox/client/ipam"
	"github.com/digitalocean/go-netbox/netbox/models"
)

func ListIpAddresses() ([]*models.IPAddress, error) {
	c := netbox.NewNetboxWithAPIKey("localhost:32768", "0123456789abcdef0123456789abcdef01234567")

	params := &ipam.IPAMIPAddressesListParams{}
	params.Context = context.Background()

	// FIXME call c.IPAM.IPAMIPAddressesList, check for errors and return the results

	return nil, nil
}
