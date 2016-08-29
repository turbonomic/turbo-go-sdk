package probe

type ObjectMeta struct {
	Name string
	UUID string
}

type PhysicalMachine struct {
	ObjectMeta
	ResourceStat *PhysicalMachineResourceStat

	consumersID []string
}

type PhysicalMachineResourceStat struct {
	cpuCapacity float64
	cpuUsed     float64
	memCapacity float64
	memUsed     float64
}

type VirtualMachine struct {
	ObjectMeta
	ResourceStat *VirtualMachineResourceStat

	providerID string
}

type VirtualMachineResourceStat struct {
	vCpuCapacity float64
	vCpuUsed     float64
	vMemCapacity float64
	vMemUsed     float64
}
