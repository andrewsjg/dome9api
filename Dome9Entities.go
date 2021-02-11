package dome9api

import "time"

//AWSCloudAccounts Object contains all accounts know to Dome9
type AWSCloudAccounts []struct {
	ID                    string      `json:"id"`
	Vendor                string      `json:"vendor"`
	Name                  string      `json:"name"`
	ExternalAccountNumber string      `json:"externalAccountNumber"`
	Error                 interface{} `json:"error"`
	IsFetchingSuspended   bool        `json:"isFetchingSuspended"`
	CreationDate          time.Time   `json:"creationDate"`
	Credentials           struct {
		Apikey     interface{} `json:"apikey"`
		Arn        string      `json:"arn"`
		Secret     interface{} `json:"secret"`
		IamUser    interface{} `json:"iamUser"`
		Type       string      `json:"type"`
		IsReadOnly bool        `json:"isReadOnly"`
	} `json:"credentials"`
	IamSafe interface{} `json:"iamSafe"`
	NetSec  struct {
		Regions []struct {
			Region           string `json:"region"`
			Name             string `json:"name"`
			Hidden           bool   `json:"hidden"`
			NewGroupBehavior string `json:"newGroupBehavior"`
		} `json:"regions"`
	} `json:"netSec"`
	Magellan               bool        `json:"magellan"`
	FullProtection         bool        `json:"fullProtection"`
	AllowReadOnly          bool        `json:"allowReadOnly"`
	OrganizationalUnitID   interface{} `json:"organizationalUnitId"`
	OrganizationalUnitPath string      `json:"organizationalUnitPath"`
	OrganizationalUnitName string      `json:"organizationalUnitName"`
	LambdaScanner          bool        `json:"lambdaScanner"`
	Serverless             struct {
		CodeAnalyzerEnabled           bool `json:"codeAnalyzerEnabled"`
		CodeDependencyAnalyzerEnabled bool `json:"codeDependencyAnalyzerEnabled"`
	} `json:"serverless"`
}

//AzureCloudAccounts - Represents and Azure Subscription
type AzureCloudAccounts []struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	SubscriptionID string `json:"subscriptionId"`
	TenantID       string `json:"tenantId"`
	Credentials    struct {
		ClientID       string      `json:"clientId"`
		ClientPassword interface{} `json:"clientPassword"`
	} `json:"credentials"`
	OperationMode          string      `json:"operationMode"`
	Error                  interface{} `json:"error"`
	CreationDate           time.Time   `json:"creationDate"`
	OrganizationalUnitID   string      `json:"organizationalUnitId"`
	OrganizationalUnitPath string      `json:"organizationalUnitPath"`
	OrganizationalUnitName string      `json:"organizationalUnitName"`
	Vendor                 string      `json:"vendor"`
	Magellan               bool        `json:"magellan"`
}

//ProtectedAssetRequest is the payload for the protected assett search
type ProtectedAssetRequest struct {
	PageSize int `json:"pageSize"`
	Filter   struct {
		FreeTextPhrase string `json:"freeTextPhrase"`
		Fields         []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"fields"`
		Tags []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"tags"`
		IncludedEntityTypes []string `json:"includedEntityTypes"`
		ExcludedEntityTypes []string `json:"excludedEntityTypes"`
	} `json:"filter"`
	SearchAfter []string `json:"searchAfter"`
}

//SearchResult - Contains results of a search request
type SearchResult struct {
	SearchRequest struct {
		PageSize int `json:"pageSize"`
		Filter   struct {
			FreeTextPhrase string `json:"freeTextPhrase"`
			Fields         []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"fields"`
			Tags []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"tags"`
			IncludedEntityTypes []string `json:"includedEntityTypes"`
			ExcludedEntityTypes []string `json:"excludedEntityTypes"`
		} `json:"filter"`
		SearchAfter []string `json:"searchAfter"`
	} `json:"searchRequest"`
	Assets []struct {
		ID                     string `json:"id"`
		EntityID               string `json:"entityId"`
		ExternalCloudAccountID string `json:"externalCloudAccountId"`
		CloudAccountID         string `json:"cloudAccountId"`
		Srl                    string `json:"srl"`
		Type                   string `json:"type"`
		Name                   string `json:"name"`
		Tags                   []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"tags"`
		Platform         string `json:"platform"`
		Network          string `json:"network"`
		Region           string `json:"region"`
		ResourceGroup    string `json:"resourceGroup"`
		AdditionalFields []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"additionalFields"`
	} `json:"assets"`
	TotalCount   int `json:"totalCount"`
	Aggregations struct {
		Property1 []struct {
			Value struct {
			} `json:"value"`
			Count int `json:"count"`
		} `json:"property1"`
		Property2 []struct {
			Value struct {
			} `json:"value"`
			Count int `json:"count"`
		} `json:"property2"`
	} `json:"aggregations"`
	SearchAfter []string `json:"searchAfter"`
}

//CloudAccountBundleFilters - Filter definitions for querting cloud account bundles
type CloudAccountBundleFilters struct {
	BundleIds        []int    `json:"bundleIds"`
	CloudAccountIds  []string `json:"cloudAccountIds"`
	CloudAccountType string   `json:"cloudAccountType"`
}

//BundleFilters - Used for retrieving asssessment results
type BundleFilters struct {
	CloudAccountBundleFilters []CloudAccountBundleFilters `json:"cloudAccountBundleFilters"`
	From                      time.Time                   `json:"from"`
	To                        time.Time                   `json:"to"`
}

//TestResults - Assessment Test Results
type TestResults []struct {
	Tests []struct {
		Error             string `json:"error"`
		TestedCount       int    `json:"testedCount"`
		RelevantCount     int    `json:"relevantCount"`
		NonComplyingCount int    `json:"nonComplyingCount"`
		ExclusionStats    struct {
			TestedCount       int `json:"testedCount"`
			RelevantCount     int `json:"relevantCount"`
			NonComplyingCount int `json:"nonComplyingCount"`
		} `json:"exclusionStats"`
		EntityResults []struct {
			ValidationStatus string `json:"validationStatus"`
			IsRelevant       bool   `json:"isRelevant"`
			IsValid          bool   `json:"isValid"`
			IsExcluded       bool   `json:"isExcluded"`
			ExclusionID      string `json:"exclusionId"`
			RemediationID    string `json:"remediationId"`
			Error            string `json:"error"`
			TestObj          struct {
			} `json:"testObj"`
		} `json:"entityResults"`
		Rule struct {
			Name          string `json:"name"`
			Severity      string `json:"severity"`
			Logic         string `json:"logic"`
			Description   string `json:"description"`
			Remediation   string `json:"remediation"`
			ComplianceTag string `json:"complianceTag"`
			Domain        string `json:"domain"`
			Priority      string `json:"priority"`
			ControlTitle  string `json:"controlTitle"`
			RuleID        string `json:"ruleId"`
			Category      string `json:"category"`
			LogicHash     string `json:"logicHash"`
			IsDefault     bool   `json:"isDefault"`
		} `json:"rule"`
		TestPassed bool `json:"testPassed"`
	} `json:"tests"`
	TestEntities struct {
		Property1 []struct {
		} `json:"property1"`
		Property2 []struct {
		} `json:"property2"`
	} `json:"testEntities"`
	Exclusions []struct {
		Platform string `json:"platform"`
		ID       string `json:"id"`
		Rules    []struct {
			LogicHash string `json:"logicHash"`
			ID        string `json:"id"`
			Name      string `json:"name"`
		} `json:"rules"`
		LogicExpressions      []string `json:"logicExpressions"`
		RulesetID             int      `json:"rulesetId"`
		CloudAccountIds       []string `json:"cloudAccountIds"`
		OrganizationalUnitIds []string `json:"organizationalUnitIds"`
		Comment               string   `json:"comment"`
		DateRange             struct {
			From time.Time `json:"from"`
			To   time.Time `json:"to"`
		} `json:"dateRange"`
	} `json:"exclusions"`
	Remediations []struct {
		Platform string `json:"platform"`
		ID       string `json:"id"`
		Rules    []struct {
			LogicHash string `json:"logicHash"`
			ID        string `json:"id"`
			Name      string `json:"name"`
		} `json:"rules"`
		LogicExpressions      []string `json:"logicExpressions"`
		RulesetID             int      `json:"rulesetId"`
		CloudAccountIds       []string `json:"cloudAccountIds"`
		Comment               string   `json:"comment"`
		CloudBots             []string `json:"cloudBots"`
		OrganizationalUnitIds []string `json:"organizationalUnitIds"`
		DateRange             struct {
			From time.Time `json:"from"`
			To   time.Time `json:"to"`
		} `json:"dateRange"`
	} `json:"remediations"`
	DataSyncStatus []struct {
		EntityType                   string `json:"entityType"`
		RecentlySuccessfulSync       bool   `json:"recentlySuccessfulSync"`
		GeneralFetchPermissionIssues bool   `json:"generalFetchPermissionIssues"`
		EntitiesWithPermissionIssues []struct {
			ExternalID            string `json:"externalId"`
			Name                  string `json:"name"`
			CloudVendorIdentifier string `json:"cloudVendorIdentifier"`
		} `json:"entitiesWithPermissionIssues"`
	} `json:"dataSyncStatus"`
	CreatedTime      time.Time `json:"createdTime"`
	ID               int       `json:"id"`
	AssessmentID     string    `json:"assessmentId"`
	TriggeredBy      string    `json:"triggeredBy"`
	AssessmentPassed bool      `json:"assessmentPassed"`
	HasErrors        bool      `json:"hasErrors"`
	Stats            struct {
		Passed                int `json:"passed"`
		PassedRulesBySeverity struct {
			Informational int `json:"informational"`
			Low           int `json:"low"`
			Medium        int `json:"medium"`
			High          int `json:"high"`
			Critical      int `json:"critical"`
		} `json:"passedRulesBySeverity"`
		Failed                int `json:"failed"`
		FailedRulesBySeverity struct {
			Informational int `json:"informational"`
			Low           int `json:"low"`
			Medium        int `json:"medium"`
			High          int `json:"high"`
			Critical      int `json:"critical"`
		} `json:"failedRulesBySeverity"`
		Error                   int `json:"error"`
		FailedTests             int `json:"failedTests"`
		LogicallyTested         int `json:"logicallyTested"`
		FailedEntities          int `json:"failedEntities"`
		ExcludedTests           int `json:"excludedTests"`
		ExcludedFailedTests     int `json:"excludedFailedTests"`
		ExcludedRules           int `json:"excludedRules"`
		ExcludedRulesBySeverity struct {
			Informational int `json:"informational"`
			Low           int `json:"low"`
			Medium        int `json:"medium"`
			High          int `json:"high"`
			Critical      int `json:"critical"`
		} `json:"excludedRulesBySeverity"`
	} `json:"stats"`
	Request struct {
		IsTemplate             bool   `json:"isTemplate"`
		ID                     int    `json:"id"`
		Name                   string `json:"name"`
		Description            string `json:"description"`
		Dome9CloudAccountID    string `json:"dome9CloudAccountId"`
		ExternalCloudAccountID string `json:"externalCloudAccountId"`
		CloudAccountID         string `json:"cloudAccountId"`
		Region                 string `json:"region"`
		CloudNetwork           string `json:"cloudNetwork"`
		CloudAccountType       string `json:"cloudAccountType"`
		RequestID              string `json:"requestId"`
		ShouldMinimizeResult   bool   `json:"shouldMinimizeResult"`
	} `json:"request"`
	HasDataSyncStatusIssues bool   `json:"hasDataSyncStatusIssues"`
	ComparisonCustomID      string `json:"comparisonCustomId"`
	AdditionalFields        struct {
		Property1 string `json:"property1"`
		Property2 string `json:"property2"`
	} `json:"additionalFields"`
}

// TODO: Look at breaking this into smaller structs possibly to make it easier to follow

//BundleResult - A result back from a Bundle run
type BundleResult []struct {
	Tests []struct {
		Error             interface{} `json:"error"`
		TestedCount       int         `json:"testedCount"`
		RelevantCount     int         `json:"relevantCount"`
		NonComplyingCount int         `json:"nonComplyingCount"`
		ExclusionStats    struct {
			TestedCount       int `json:"testedCount"`
			RelevantCount     int `json:"relevantCount"`
			NonComplyingCount int `json:"nonComplyingCount"`
		} `json:"exclusionStats"`
		EntityResults []struct {
			ValidationStatus string      `json:"validationStatus"`
			IsRelevant       bool        `json:"isRelevant"`
			IsValid          bool        `json:"isValid"`
			IsExcluded       bool        `json:"isExcluded"`
			ExclusionID      interface{} `json:"exclusionId"`
			RemediationID    interface{} `json:"remediationId"`
			Error            string      `json:"error"`
			TestObj          struct {
				ID          string `json:"id"`
				Dome9ID     string `json:"dome9Id"`
				EntityType  string `json:"entityType"`
				EntityIndex int    `json:"entityIndex"`
			} `json:"testObj"`
		} `json:"entityResults"`
		Rule struct {
			Name          string `json:"name"`
			Severity      string `json:"severity"`
			Logic         string `json:"logic"`
			Description   string `json:"description"`
			Remediation   string `json:"remediation"`
			ComplianceTag string `json:"complianceTag"`
			Domain        string `json:"domain"`
			Priority      string `json:"priority"`
			ControlTitle  string `json:"controlTitle"`
			RuleID        string `json:"ruleId"`
			Category      string `json:"category"`
			LogicHash     string `json:"logicHash"`
			IsDefault     bool   `json:"isDefault"`
		} `json:"rule"`
		TestPassed bool `json:"testPassed"`
	} `json:"tests"`
	TestEntities struct {
		Instance []struct {
			Image        string `json:"image"`
			ImageDetails struct {
				Name                string      `json:"name"`
				OwnerID             string      `json:"ownerId"`
				ImageLocation       string      `json:"imageLocation"`
				IsPublic            bool        `json:"isPublic"`
				State               string      `json:"state"`
				Description         string      `json:"description"`
				CreationDate        int         `json:"creationDate"`
				Hypervisor          string      `json:"hypervisor"`
				Platform            interface{} `json:"platform"`
				Architecture        string      `json:"architecture"`
				RootDeviceName      string      `json:"rootDeviceName"`
				RootDeviceType      string      `json:"rootDeviceType"`
				VirtualizationType  string      `json:"virtualizationType"`
				RamdiskID           interface{} `json:"ramdiskId"`
				ImageOwnerAlias     string      `json:"imageOwnerAlias"`
				ImageType           string      `json:"imageType"`
				KernelID            interface{} `json:"kernelId"`
				BlockDeviceMappings []struct {
					DeviceName string `json:"deviceName"`
					Ebs        struct {
						DeleteOnTermination bool   `json:"deleteOnTermination"`
						Encrypted           bool   `json:"encrypted"`
						Iops                int    `json:"iops"`
						SnapshotID          string `json:"snapshotId"`
						VolumeSize          int    `json:"volumeSize"`
						VolumeType          string `json:"volumeType"`
					} `json:"ebs"`
					NoDevice    interface{} `json:"noDevice"`
					VirtualName interface{} `json:"virtualName"`
				} `json:"blockDeviceMappings"`
				ProductCodes    []interface{} `json:"productCodes"`
				EnaSupport      bool          `json:"enaSupport"`
				SriovNetSupport string        `json:"sriovNetSupport"`
				StateReason     interface{}   `json:"stateReason"`
				Tags            []interface{} `json:"tags"`
			} `json:"imageDetails"`
			KernelID     interface{} `json:"kernelId"`
			Platform     string      `json:"platform"`
			LaunchTime   int         `json:"launchTime"`
			InboundRules []struct {
				Protocol      string      `json:"protocol"`
				Port          int         `json:"port"`
				PortTo        int         `json:"portTo"`
				Scope         string      `json:"scope"`
				ScopeMetaData interface{} `json:"scopeMetaData"`
				ServiceType   string      `json:"serviceType"`
			} `json:"inboundRules"`
			OutboundRules []struct {
				Protocol      string      `json:"protocol"`
				Port          int         `json:"port"`
				PortTo        int         `json:"portTo"`
				Scope         string      `json:"scope"`
				ScopeMetaData interface{} `json:"scopeMetaData"`
				ServiceType   string      `json:"serviceType"`
			} `json:"outboundRules"`
			Nics []struct {
				ID     string `json:"id"`
				Name   string `json:"name"`
				Subnet struct {
					Vpc struct {
						CloudAccountID   string        `json:"cloudAccountId"`
						Cidr             string        `json:"cidr"`
						Region           string        `json:"region"`
						ID               string        `json:"id"`
						AccountNumber    string        `json:"accountNumber"`
						VpnGateways      []interface{} `json:"vpnGateways"`
						InternetGateways []struct {
							ExternalID     string `json:"externalId"`
							VpcAttachments []struct {
								State string `json:"state"`
								VpcID string `json:"vpcId"`
							} `json:"vpcAttachments"`
							Name string `json:"name"`
						} `json:"internetGateways"`
						DhcpOptionsID   string `json:"dhcpOptionsId"`
						InstanceTenancy string `json:"instanceTenancy"`
						IsDefault       bool   `json:"isDefault"`
						State           string `json:"state"`
						Tags            []struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"tags"`
						Name                  string      `json:"name"`
						VpcPeeringConnections interface{} `json:"vpcPeeringConnections"`
						Source                string      `json:"source"`
					} `json:"vpc"`
					State                   string `json:"state"`
					AvailabilityZone        string `json:"availabilityZone"`
					DefaultForAz            bool   `json:"defaultForAz"`
					MapPublicIPOnLaunch     bool   `json:"mapPublicIpOnLaunch"`
					AvailableIPAddressCount int    `json:"availableIpAddressCount"`
					ExternalID              string `json:"externalId"`
					Description             string `json:"description"`
					Cidr                    string `json:"cidr"`
					RouteTable              struct {
						Associations []struct {
							IsMain   bool   `json:"isMain"`
							SubnetID string `json:"subnetId"`
						} `json:"associations"`
						PropagatingVgws []interface{} `json:"propagatingVgws"`
						Routes          []struct {
							DestinationCidrBlock        string      `json:"destinationCidrBlock"`
							DestinationIpv6CidrBlock    interface{} `json:"destinationIpv6CidrBlock"`
							DestinationPrefixListID     interface{} `json:"destinationPrefixListId"`
							EgressOnlyInternetGatewayID interface{} `json:"egressOnlyInternetGatewayId"`
							GatewayID                   string      `json:"gatewayId"`
							InstanceID                  interface{} `json:"instanceId"`
							InstanceOwnerID             interface{} `json:"instanceOwnerId"`
							NatGatewayID                interface{} `json:"natGatewayId"`
							NetworkInterfaceID          interface{} `json:"networkInterfaceId"`
							Origin                      string      `json:"origin"`
							State                       string      `json:"state"`
							VpcPeeringConnectionID      interface{} `json:"vpcPeeringConnectionId"`
						} `json:"routes"`
						RouteTableID string `json:"routeTableId"`
						VpcID        string `json:"vpcId"`
						Tags         struct {
						} `json:"tags"`
						Name         string `json:"name"`
						TagsEntities struct {
						} `json:"tagsEntities"`
					} `json:"routeTable"`
					Nacl struct {
						Tags             []interface{} `json:"tags"`
						ExternalFindings interface{}   `json:"externalFindings"`
						Source           string        `json:"source"`
						Type             string        `json:"type"`
						Vpc              struct {
							CloudAccountID   string        `json:"cloudAccountId"`
							Cidr             string        `json:"cidr"`
							Region           string        `json:"region"`
							ID               string        `json:"id"`
							AccountNumber    string        `json:"accountNumber"`
							VpnGateways      []interface{} `json:"vpnGateways"`
							InternetGateways []struct {
								ExternalID     string `json:"externalId"`
								VpcAttachments []struct {
									State string `json:"state"`
									VpcID string `json:"vpcId"`
								} `json:"vpcAttachments"`
								Name string `json:"name"`
							} `json:"internetGateways"`
							DhcpOptionsID   string `json:"dhcpOptionsId"`
							InstanceTenancy string `json:"instanceTenancy"`
							IsDefault       bool   `json:"isDefault"`
							State           string `json:"state"`
							Tags            []struct {
								Key   string `json:"key"`
								Value string `json:"value"`
							} `json:"tags"`
							Name                  string      `json:"name"`
							VpcPeeringConnections interface{} `json:"vpcPeeringConnections"`
							Source                string      `json:"source"`
						} `json:"vpc"`
						Name          string `json:"name"`
						ID            string `json:"id"`
						Dome9ID       string `json:"dome9Id"`
						AccountNumber string `json:"accountNumber"`
						Region        string `json:"region"`
						Inbound       []struct {
							IcmpProtocol      interface{} `json:"icmpProtocol"`
							Name              string      `json:"name"`
							Number            int         `json:"number"`
							Protocol          string      `json:"protocol"`
							Source            string      `json:"source"`
							Destination       string      `json:"destination"`
							DestinationPort   int         `json:"destinationPort"`
							DestinationPortTo int         `json:"destinationPortTo"`
							Direction         string      `json:"direction"`
							Action            string      `json:"action"`
						} `json:"inbound"`
						Outbound []struct {
							IcmpProtocol      interface{} `json:"icmpProtocol"`
							Name              string      `json:"name"`
							Number            int         `json:"number"`
							Protocol          string      `json:"protocol"`
							Source            string      `json:"source"`
							Destination       string      `json:"destination"`
							DestinationPort   int         `json:"destinationPort"`
							DestinationPortTo int         `json:"destinationPortTo"`
							Direction         string      `json:"direction"`
							Action            string      `json:"action"`
						} `json:"outbound"`
						IsDefault bool `json:"isDefault"`
					} `json:"nacl"`
					Tags []struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"tags"`
					ID               string      `json:"id"`
					Type             string      `json:"type"`
					Name             string      `json:"name"`
					Dome9ID          string      `json:"dome9Id"`
					AccountNumber    string      `json:"accountNumber"`
					Region           string      `json:"region"`
					ExternalFindings interface{} `json:"externalFindings"`
				} `json:"subnet"`
				PrivateDNSName   string      `json:"privateDnsName"`
				PublicIPAddress  string      `json:"publicIpAddress"`
				PrivateIPAddress string      `json:"privateIpAddress"`
				ElasticIP        interface{} `json:"elasticIP"`
				MacAddress       string      `json:"macAddress"`
				SecurityGroups   []struct {
					ID            string `json:"id"`
					Type          string `json:"type"`
					Name          string `json:"name"`
					Dome9ID       string `json:"dome9Id"`
					AccountNumber string `json:"accountNumber"`
					Region        string `json:"region"`
					Vpc           struct {
						CloudAccountID   string        `json:"cloudAccountId"`
						Cidr             string        `json:"cidr"`
						Region           string        `json:"region"`
						ID               string        `json:"id"`
						AccountNumber    string        `json:"accountNumber"`
						VpnGateways      []interface{} `json:"vpnGateways"`
						InternetGateways []struct {
							ExternalID     string `json:"externalId"`
							VpcAttachments []struct {
								State string `json:"state"`
								VpcID string `json:"vpcId"`
							} `json:"vpcAttachments"`
							Name string `json:"name"`
						} `json:"internetGateways"`
						DhcpOptionsID   string `json:"dhcpOptionsId"`
						InstanceTenancy string `json:"instanceTenancy"`
						IsDefault       bool   `json:"isDefault"`
						State           string `json:"state"`
						Tags            []struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"tags"`
						Name                  string      `json:"name"`
						VpcPeeringConnections interface{} `json:"vpcPeeringConnections"`
						Source                string      `json:"source"`
					} `json:"vpc"`
					Tags []struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"tags"`
					Source       string `json:"source"`
					Description  string `json:"description"`
					InboundRules []struct {
						Protocol      string      `json:"protocol"`
						Port          int         `json:"port"`
						PortTo        int         `json:"portTo"`
						Scope         string      `json:"scope"`
						ScopeMetaData interface{} `json:"scopeMetaData"`
						ServiceType   string      `json:"serviceType"`
					} `json:"inboundRules"`
					OutboundRules []struct {
						Protocol      string      `json:"protocol"`
						Port          int         `json:"port"`
						PortTo        int         `json:"portTo"`
						Scope         string      `json:"scope"`
						ScopeMetaData interface{} `json:"scopeMetaData"`
						ServiceType   string      `json:"serviceType"`
					} `json:"outboundRules"`
					InboundPrefixes    []interface{} `json:"inboundPrefixes"`
					OutboundPrefixes   []interface{} `json:"outboundPrefixes"`
					NetworkAssetsStats []interface{} `json:"networkAssetsStats"`
					IsProtected        bool          `json:"isProtected"`
					NetworkInterfaces  interface{}   `json:"networkInterfaces"`
				} `json:"securityGroups"`
			} `json:"nics"`
			IsPublic     bool   `json:"isPublic"`
			InstanceType string `json:"instanceType"`
			IsRunning    bool   `json:"isRunning"`
			Volumes      []struct {
				Attachments []struct {
					AttachTime          int    `json:"attachTime"`
					DeleteOnTermination bool   `json:"deleteOnTermination"`
					Device              string `json:"device"`
					InstanceID          string `json:"instanceId"`
					State               string `json:"state"`
					VolumeID            string `json:"volumeId"`
				} `json:"attachments"`
				AvailabilityZone string        `json:"availabilityZone"`
				CreateTime       int           `json:"createTime"`
				Encrypted        bool          `json:"encrypted"`
				Iops             int           `json:"iops"`
				KmsKeyID         interface{}   `json:"kmsKeyId"`
				Size             int           `json:"size"`
				SnapshotID       string        `json:"snapshotId"`
				State            string        `json:"state"`
				Tags             []interface{} `json:"tags"`
				VolumeID         string        `json:"volumeId"`
				VolumeType       string        `json:"volumeType"`
				EncryptionKey    interface{}   `json:"encryptionKey"`
			} `json:"volumes"`
			ProfileArn string        `json:"profileArn"`
			Roles      []interface{} `json:"roles"`
			Scanners   struct {
				Scans    interface{} `json:"scans"`
				Findings interface{} `json:"findings"`
			} `json:"scanners"`
			AutoScalingGroup interface{} `json:"autoScalingGroup"`
			PrivateDNS       string      `json:"privateDns"`
			PublicDNS        string      `json:"publicDns"`
			Tags             []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"tags"`
			Vpc struct {
				CloudAccountID   string        `json:"cloudAccountId"`
				Cidr             string        `json:"cidr"`
				Region           string        `json:"region"`
				ID               string        `json:"id"`
				AccountNumber    string        `json:"accountNumber"`
				VpnGateways      []interface{} `json:"vpnGateways"`
				InternetGateways []struct {
					ExternalID     string `json:"externalId"`
					VpcAttachments []struct {
						State string `json:"state"`
						VpcID string `json:"vpcId"`
					} `json:"vpcAttachments"`
					Name string `json:"name"`
				} `json:"internetGateways"`
				DhcpOptionsID   string `json:"dhcpOptionsId"`
				InstanceTenancy string `json:"instanceTenancy"`
				IsDefault       bool   `json:"isDefault"`
				State           string `json:"state"`
				Tags            []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"tags"`
				Name                  string      `json:"name"`
				VpcPeeringConnections interface{} `json:"vpcPeeringConnections"`
				Source                string      `json:"source"`
			} `json:"vpc"`
			ID               string      `json:"id"`
			Type             string      `json:"type"`
			Name             string      `json:"name"`
			Dome9ID          string      `json:"dome9Id"`
			AccountNumber    string      `json:"accountNumber"`
			Region           string      `json:"region"`
			ExternalFindings interface{} `json:"externalFindings"`
		} `json:"instance"`
	} `json:"testEntities"`
	Exclusions     []interface{} `json:"exclusions"`
	Remediations   []interface{} `json:"remediations"`
	DataSyncStatus []struct {
		EntityType                   string        `json:"entityType"`
		RecentlySuccessfulSync       bool          `json:"recentlySuccessfulSync"`
		GeneralFetchPermissionIssues bool          `json:"generalFetchPermissionIssues"`
		EntitiesWithPermissionIssues []interface{} `json:"entitiesWithPermissionIssues"`
	} `json:"dataSyncStatus"`
	CreatedTime      time.Time `json:"createdTime"`
	ID               int       `json:"id"`
	AssessmentID     string    `json:"assessmentId"`
	TriggeredBy      string    `json:"triggeredBy"`
	AssessmentPassed bool      `json:"assessmentPassed"`
	HasErrors        bool      `json:"hasErrors"`
	Stats            struct {
		Passed                int `json:"passed"`
		PassedRulesBySeverity struct {
			Informational int `json:"informational"`
			Low           int `json:"low"`
			Medium        int `json:"medium"`
			High          int `json:"high"`
			Critical      int `json:"critical"`
		} `json:"passedRulesBySeverity"`
		Failed                int `json:"failed"`
		FailedRulesBySeverity struct {
			Informational int `json:"informational"`
			Low           int `json:"low"`
			Medium        int `json:"medium"`
			High          int `json:"high"`
			Critical      int `json:"critical"`
		} `json:"failedRulesBySeverity"`
		Error                   int `json:"error"`
		FailedTests             int `json:"failedTests"`
		LogicallyTested         int `json:"logicallyTested"`
		FailedEntities          int `json:"failedEntities"`
		ExcludedTests           int `json:"excludedTests"`
		ExcludedFailedTests     int `json:"excludedFailedTests"`
		ExcludedRules           int `json:"excludedRules"`
		ExcludedRulesBySeverity struct {
			Informational int `json:"informational"`
			Low           int `json:"low"`
			Medium        int `json:"medium"`
			High          int `json:"high"`
			Critical      int `json:"critical"`
		} `json:"excludedRulesBySeverity"`
	} `json:"stats"`
	Request struct {
		IsTemplate             bool        `json:"isTemplate"`
		ID                     int         `json:"id"`
		Name                   string      `json:"name"`
		Description            string      `json:"description"`
		Dome9CloudAccountID    string      `json:"dome9CloudAccountId"`
		ExternalCloudAccountID string      `json:"externalCloudAccountId"`
		CloudAccountID         string      `json:"cloudAccountId"`
		Region                 interface{} `json:"region"`
		CloudNetwork           interface{} `json:"cloudNetwork"`
		CloudAccountType       string      `json:"cloudAccountType"`
		RequestID              string      `json:"requestId"`
		ShouldMinimizeResult   bool        `json:"shouldMinimizeResult"`
	} `json:"request"`
	HasDataSyncStatusIssues bool        `json:"hasDataSyncStatusIssues"`
	ComparisonCustomID      interface{} `json:"comparisonCustomId"`
	AdditionalFields        interface{} `json:"additionalFields"`
}

/* All Dome9 Entity Types
"aws|S3Bucket",
"aws|Instance",
"aws|Lambda",
"aws|IamRole",
"aws|Volume",
"aws|EbsSnapshot",
"aws|IamUser",
"aws|DynamoDbTable",
"aws|VPC",
"aws|AcmCertificate",
"aws|KMS"
"aws|ApiGateway"
"aws|RDS"
"aws|CloudFront"
"aws|IamServerCertificate"
"aws|VPNConnection"
"aws|CustomerGateway"
"aws|Kinesis"
"aws|CloudTrail"
"aws|EFS"
"aws|WAFRegional"
"aws|WAFRegionalV2"
"aws|Redshift"
"aws|ELB"
"aws|ApplicationLoadBalancer"
"aws|NetworkLoadBalancer"
"aws|EcsCluster"
"aws|EksCluster"
"aws|VPNGateway"
"aws|Route53HostedZone"
"aws|Route53Domain"
"aws|SageMakerNotebook"
"aws|ElasticSearchDomain"
"aws|EmrCluster"
"aws|VpcEndpoint"
"aws|ElastiCache"
"aws|ConfigSetting"
"aws|Sqs"
"aws|NatGateway"
"aws|GlueSecurityConfiguration"
"aws|GlueConnection"
"aws|StepFunctionStateMachine"
"aws|SnsTopic"
"aws|SnsPlatformApplication"
"aws|MskCluster"
"aws|Region"
"aws|Organization"
"aws|Account"
"aws|NetworkInterface"
"aws|TransitGateway"
"aws|SecretManager"
"aws|RouteTable"
"aws|SageMakerTrainingJob"
"aws|CognitoIdentityPool"
"aws|CognitoUserPool"
"aws|Workspace"
"aws|EcrRepository"
"aws|TranslationJob"
"aws|TranslationTerminology"
"aws|TranscribeJob"
"aws|TranscribeMedicalJob"
"aws|Personalize"
"aws|Transfer"
"aws|SystemManagerParameter"
"aws|RDSDBSnapshot"
"aws|RDSDBCluster"
"aws|MqBroker"
"aws|NetworkFirewall"
"aws|Athena"
"aws|AthenaWorkGroup"
"aws|DmsEndpoint"
"aws|CloudFormationStack"
"aws|KinesisFirehose"
"aws|StorageGateway"
"aws|SystemManagerDocument"
"aws|CustomDomainName"
"azure|LoadBalancer"
"azure|VirtualMachine"
"azure|VirtualNetworkGateway"
"azure|Firewall"
"azure|StorageAccount"
"azure|KeyVault"
"azure|RedisCache"
"azure|SQLServer"
"azure|ApplicationGateway"
"azure|RouteTable"
"azure|NetworkWatcher"
"azure|NsgFlowLog"
"azure|LogicApp"
"azure|PostgreSQL"
"azure|ContainerRegistry"
"azure|CosmosDbAccount"
"azure|PolicyAssignment"
"azure|LogProfile"
"azure|VirtualMachineScaleSet"
"azure|ApplicationSecurityGroup"
"azure|AnalysisServiceServer"
"azure|AksCluster"
"azure|FunctionApp"
"azure|WebApp"
 "azure|User"
"azure|VMSSInstance"
"azure|SqlManagedInstance"
"azure|ActivityLogAlertRule"
"azure|ApiManagementService"
"azure|HDInsight"
"azure|DataExplorerCluster"
"azure|Disk"
"azure|RegionalWAF"
"azure|RoleAssignment"
"azure|RoleDefinition"
"azure|VNet"
"azure|ServiceBus"
"azure|EventHubNamespace"
"google|ServiceAccount"
"google|BigQuery"
"google|StorageBucket"
"google|GkeCluster"
"google|KmsKeyRing"
"google|VMInstance"
"google|CloudSql"
"google|CloudFunction"
"google|BigTable"
"google|Network"
"google|PubSubTopic"
"google|GcpIamUser"
"google|GcpIamGroup"
"google|GcpIamPolicy"
"google|Image"
"google|Redis"
"google|AppEngine"
"kubernetes|KubernetesNode"
 "kubernetes|KubernetesPod"
 "kubernetes|KubernetesService"
 "kubernetes|KubernetesIngress"
 "kubernetes|KubernetesNetworkPolicy"
 "kubernetes|KubernetesRole"
 "kubernetes|KubernetesRoleBinding"
 "kubernetes|KubernetesServiceAccount"
 "kubernetes|KubernetesPodSecurityPolicy"
 "kubernetes|KubernetesNamespace"
"kubernetes|KubernetesCluster"
 "kubernetes|KubernetesImage"
 "kubernetes|KubernetesJob"
 "kubernetes|KubernetesCronJob"
 "kubernetes|KubernetesStatefulSet"
 "kubernetes|KubernetesDaemonSet"
 "kubernetes|KubernetesDeployment"
 "kubernetes|KubernetesReplicaSet""
    ]
  }
}

*/
