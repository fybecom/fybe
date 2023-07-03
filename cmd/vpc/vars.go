package cmd

// create
var (
	createVpnName        string
	createVpnRegion      string
	createVpnDescription string
)

// list
var (
	listVpcNameFilter string
)

// get
var (
	getvpcId int64
)

// assign
var (
	vpcId            int64
	assignInstanceId int64
)

// unassign
var (
	unassignvpcId      int64
	unassignInstanceId int64
)

// delete
var (
	deletevpcId int64
)

// update
var (
	updateVpnId          int64
	updateVpnName        string
	updateVpnDescription string
)
