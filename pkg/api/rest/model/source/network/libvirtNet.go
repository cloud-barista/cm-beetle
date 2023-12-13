package network

import (
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

type LibvirtDomain struct {
	DomainName string                       `json:"domain_name"`
	DomainUUID string                       `json:"domain_uuid"`
	Interfaces []libvirtxml.DomainInterface `json:"interfaces"`
}

type LibvirtNet struct {
	Domains []LibvirtDomain `json:"domains"`
}
