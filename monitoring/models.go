/*
Copyright 2013 Rackspace

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS-IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package monitoring

type PaginationMetadata struct {
	Count      int     `json:"count"`
	Limit      int     `json:"limit"`
	Marker     *string `json:"marker"`
	NextMarker *string `json:"next_marker"`
	NextHref   *string `json:"next_href"`
}

type Entity struct {
	Id          string            `json:"id"`
	Label       *string           `json:"label"`
	Metadata    map[string]string `json:"metadata"`
	Managed     bool              `json:"managed"`
	Uri         *string           `json:"uri"`
	AgentId     *string           `json:"agent_id"`
	IPAddresses map[string]string `json:"ip_addresses"`
}

type CpuHostInfo struct {
	Timestamp int64 `json:"timestamp"`
	Info      []struct {
		Name    string `json:"name"`
		Vendor  string `json:"vendor"`
		Model   string `json:"model"`
		Mhz     int    `json:"mhz"`
		Idle    int64  `json:"idle"`
		Irq     int    `json:"irq"`
		SoftIrq int    `json:"soft_irq"`
		Nice    int    `json:"nice"`
		Stolen  int    `json:"stolen"`
		Sys     int    `json:"sys"`
		User    int    `json:"user"`
		Wait    int    `json:"wait"`
	} `json:"info"`
}

type MemoryHostInfo struct {
	Timestamp int64 `json:"timestamp"`
	Info      struct {
		ActualFree  int64 `json:"actual_free"`
		ActualUsed  int   `json:"actual_used"`
		Free        int   `json:"free"`
		Used        int64 `json:"used"`
		Total       int64 `json:"total"`
		Ram         int   `json:"ram"`
		SwapTotal   int   `json:"swap_total"`
		SwapUsed    int   `json:"swap_used"`
		SwapFree    int   `json:"swap_free"`
		SwapPageIn  int   `json:"swap_page_in"`
		SwapPageOut int   `json:"swap_page_out"`
	} `json:"info"`
}

type NetworkInterfaceHostInfo struct {
	Timestamp int64 `json:"timestamp"`
	Info      []struct {
		Name      string `json:"name"`
		Type      string `json:"type"`
		Address   string `json:"address"`
		Netmask   string `json:"netmask"`
		Address6  string `json:"address6"`
		Broadcast string `json:"broadcast"`
		Hwaddr    string `json:"hwaddr"`
		Mtu       int    `json:"mtu"`
		RxPackets int    `json:"rx_packets"`
		RxBytes   int64  `json:"rx_bytes"`
		TxPackets int    `json:"tx_packets"`
		TxBytes   int64  `json:"tx_bytes"`
		Flags     int    `json:"flags"`
	} `json:"info"`
}

type SystemHostInfo struct {
	Timestamp int64 `json:"timestamp"`
	Info      struct {
		Name          string `json:"name"`
		Arch          string `json:"arch"`
		Version       string `json:"version"`
		Vendor        string `json:"vendor"`
		VendorVersion string `json:"vendor_version"`
	} `json:"info"`
}

type DiskHostInfo struct {
	Timestamp int64 `json:"timestamp"`
	Info      []struct {
		ReadBytes  int64  `json:"read_bytes"`
		Reads      int    `json:"reads"`
		Rtime      int    `json:"rtime"`
		WriteBytes int64  `json:"write_bytes"`
		Writes     int    `json:"writes"`
		Wtime      int    `json:"wtime"`
		Time       int    `json:"time"`
		Name       string `json:"name"`
	} `json:"info"`
}

type FilesystemsHostInfo struct {
	Timestamp int64 `json:"timestamp"`
	Info      []struct {
		DirName     string `json:"dir_name"`
		DevName     string `json:"dev_name"`
		SysTypeName string `json:"sys_type_name"`
		Options     string `json:"options"`
		Free        int    `json:"free"`
		Used        int    `json:"used"`
		Avail       int    `json:"avail"`
		Total       int    `json:"total"`
		Files       int    `json:"files"`
		FreeFiles   int    `json:"free_files"`
	} `json:"info"`
}

type ProcessesHostInfo struct {
	Timestamp int64 `json:"timestamp"`
	Info      []struct {
		Pid               int    `json:"pid"`
		ExeName           string `json:"exe_name"`
		ExeCwd            string `json:"exe_cwd"`
		ExeRoot           string `json:"exe_root"`
		TimeTotal         int    `json:"time_total"`
		TimeSys           int    `json:"time_sys"`
		TimeUser          int    `json:"time_user"`
		TimeStartTime     int64  `json:"time_start_time"`
		StateName         string `json:"state_name"`
		StatePriority     int    `json:"state_priority"`
		StateThreads      int    `json:"state_threads"`
		MemorySize        int    `json:"memory_size"`
		MemoryResident    int    `json:"memory_resident"`
		MemoryShare       int    `json:"memory_share"`
		MemoryMajorFaults int    `json:"memory_major_faults"`
		MemoryMinorFaults int    `json:"memory_minor_faults"`
		MemoryPageFaults  int    `json:"memory_page_faults"`
		CredUser          string `json:"cred_user"`
		CredGroup         string `json:"cred_group"`
	} `json:"info"`
}

type PaginatedEntityList struct {
	Values   []Entity
	Metadata PaginationMetadata
}

// A Check represents a check that the user configures on one of his or her entities.
// It contains the following fields:
//
// The Id field uniquely identifies the check amongst all others belonging to the user.
//
// The Label field identifies the check to a human operator.
//
// The Type field indicates what kind of check it is.
//
// The Details field provides a mapping of detail to detail-specific information.
//
// The MonitoringZonesPoll field provides a list of what to poll for this check.
//
// The Timeout field indicates how many seconds to wait for a response before the check fails.
//
// The Period field tells how frequently to perform the check, in seconds.
//
// The TargetAlias field does something; I just don't quite know what it is.
//
// The TargetHostname field identifies the host name of that which is being checked.
//
// TargetResolver field identifies the domain name resolver scoping the target hostname.
//
// The Disabled field is true if the check is disabled for the entity; false otherwise.
//
// The Metadata field provides a generic key/value store of miscellaneous bits of information relevant to this check.
// However, its implementation isn't very efficient at all.  This field is not intended for use as a general purpose
// key/value store.
type Check struct {
	Id                  string                 `json:"id"`
	Label               *string                `json:"label"`
	Type                string                 `json:"type"`
	Details             map[string]interface{} `json:"details"`
	MonitoringZonesPoll []string               `json:"monitoring_zones_poll"`
	Timeout             int                    `json:"timeout"`
	Period              int                    `json:"period"`
	TargetAlias         *string                `json:"target_alias"`
	TargetHostname      *string                `json:"target_hostname"`
	TargetResolver      *string                `json:"target_resolver"`
	Disabled            bool                   `json:"disabled"`
	Metadata            map[string]interface{} `json:"metadata"`
}

// A PaginatedCheckList contains a finite subset of the complete set of checks a cloud monitoring user has configured.
// The Values field contains the array slice representing the set of Check objects.
type PaginatedCheckList struct {
	Values   []Check
	Metadata PaginationMetadata
}
