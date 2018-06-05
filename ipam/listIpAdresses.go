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

	res, err := c.IPAM.IPAMIPAddressesList(params, nil)
	if err != nil {
		return nil, err
	}

	return res.Payload.Results, nil
}
