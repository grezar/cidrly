package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sort"

	"github.com/xlab/treeprint"
)

type Subnets struct {
	Subnets []Subnet `json:Subnets`
}

type Subnet struct {
	MapPublicIPOnLaunch         bool   `json:"MapPublicIpOnLaunch"`
	AvailabilityZoneID          string `json:"AvailabilityZoneId"`
	Tags                        []Tag  `json:"Tags"`
	AvailableIPAddressCount     int    `json:"AvailableIpAddressCount"`
	DefaultForAz                bool   `json:"DefaultForAz"`
	SubnetArn                   string `json:"SubnetArn"`
	Ipv6CidrBlockAssociationSet []byte `json:"Ipv6CidrBlockAssociationSet"`
	VpcID                       string `json:"VpcId"`
	State                       string `json:"State"`
	AvailabilityZone            string `json:"AvailabilityZone"`
	SubnetID                    string `json:"SubnetId"`
	OwnerID                     string `json:"OwnerId"`
	CidrBlock                   string `json:"CidrBlock"`
	AssignIpv6AddressOnCreation bool   `json:"AssignIpv6AddressOnCreation"`
}

type Tag struct {
	Value string `json:"Value"`
	Key   string `json:"Key"`
}

func NewSubnets(bytes []byte) *Subnets {
	subnets := new(Subnets)
	if err := json.Unmarshal(bytes, &subnets); err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't parse subnets JSON: %v", err)
	}
	return subnets
}

func (s Subnets) renderTree() {
	s.sortByCidrBlock()
	fmt.Printf(s.buildTree())
}

func (s *Subnets) sortByCidrBlock() {
	sort.Slice(s.Subnets, func(i, j int) bool {
		ipv4AddrI, _, _ := net.ParseCIDR(s.Subnets[i].CidrBlock)
		ipv4AddrJ, _, _ := net.ParseCIDR(s.Subnets[j].CidrBlock)
		return bytes.Compare(ipv4AddrI, ipv4AddrJ) < 0
	})
}

func (s *Subnets) buildTree() string {
	var subnetName string
	tree := treeprint.New()
	for _, s := range s.Subnets {
		for _, t := range s.Tags {
			if t.Key == "Name" {
				subnetName = t.Value
			}
		}
		tree.AddNode(fmt.Sprintf("%s(%s)", s.CidrBlock, subnetName))
	}
	return tree.String()
}
