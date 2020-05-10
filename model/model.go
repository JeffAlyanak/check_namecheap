package model

import "encoding/xml"

// ApiResponse struct holds the returned API data.
type ApiResponse struct {
	XMLName          xml.Name `xml:"ApiResponse"`
	Text             string   `xml:",chardata"`
	Status           string   `xml:"Status,attr"`
	Xmlns            string   `xml:"xmlns,attr"`
	Errors           string   `xml:"Errors"`
	Warnings         string   `xml:"Warnings"`
	RequestedCommand string   `xml:"RequestedCommand"`
	CommandResponse  struct {
		Text                string `xml:",chardata"`
		Type                string `xml:"Type,attr"`
		DomainGetInfoResult struct {
			Text          string `xml:",chardata"`
			Status        string `xml:"Status,attr"`
			ID            string `xml:"ID,attr"`
			DomainName    string `xml:"DomainName,attr"`
			OwnerName     string `xml:"OwnerName,attr"`
			IsOwner       string `xml:"IsOwner,attr"`
			IsPremium     string `xml:"IsPremium,attr"`
			DomainDetails struct {
				Text        string `xml:",chardata"`
				CreatedDate string `xml:"CreatedDate"`
				ExpiredDate string `xml:"ExpiredDate"`
				NumYears    string `xml:"NumYears"`
			} `xml:"DomainDetails"`
			LockDetails string `xml:"LockDetails"`
			Whoisguard  struct {
				Text         string `xml:",chardata"`
				Enabled      string `xml:"Enabled,attr"`
				ID           string `xml:"ID"`
				ExpiredDate  string `xml:"ExpiredDate"`
				EmailDetails struct {
					Text                         string `xml:",chardata"`
					WhoisGuardEmail              string `xml:"WhoisGuardEmail,attr"`
					ForwardedTo                  string `xml:"ForwardedTo,attr"`
					LastAutoEmailChangeDate      string `xml:"LastAutoEmailChangeDate,attr"`
					AutoEmailChangeFrequencyDays string `xml:"AutoEmailChangeFrequencyDays,attr"`
				} `xml:"EmailDetails"`
			} `xml:"Whoisguard"`
			PremiumDnsSubscription struct {
				Text           string `xml:",chardata"`
				UseAutoRenew   string `xml:"UseAutoRenew"`
				SubscriptionId string `xml:"SubscriptionId"`
				CreatedDate    string `xml:"CreatedDate"`
				ExpirationDate string `xml:"ExpirationDate"`
				IsActive       string `xml:"IsActive"`
			} `xml:"PremiumDnsSubscription"`
			DnsDetails struct {
				Text             string   `xml:",chardata"`
				ProviderType     string   `xml:"ProviderType,attr"`
				IsUsingOurDNS    string   `xml:"IsUsingOurDNS,attr"`
				HostCount        string   `xml:"HostCount,attr"`
				EmailType        string   `xml:"EmailType,attr"`
				DynamicDNSStatus string   `xml:"DynamicDNSStatus,attr"`
				IsFailover       string   `xml:"IsFailover,attr"`
				Nameserver       []string `xml:"Nameserver"`
			} `xml:"DnsDetails"`
			Modificationrights struct {
				Text string `xml:",chardata"`
				All  string `xml:"All,attr"`
			} `xml:"Modificationrights"`
		} `xml:"DomainGetInfoResult"`
	} `xml:"CommandResponse"`
	Server            string `xml:"Server"`
	GMTTimeDifference string `xml:"GMTTimeDifference"`
	ExecutionTime     string `xml:"ExecutionTime"`
}
