package probe

import (
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	"math/rand"
	"sync"

	"github.com/golang/glog"
)

const (
	defaultCPUCapacity    float64 = 4000.0
	defaultMemoryCapacity float64 = 4096000.0
	defaultVCPUCapacity   float64 = 1000.0
	defaultVMemCapacity   float64 = 1024000.0
)

// TopologyGenerator is used to generate example topology for probe to discover.
type TopologyGenerator struct {
	pmSet map[string]*PhysicalMachine
	vmSet map[string]*VirtualMachine

	pmUUIDs []string
	vmUUIDs []string

	lock sync.RWMutex
}

func NewTopologyGenerator(numPM, numVM int) (*TopologyGenerator, error) {
	g := &TopologyGenerator{
		pmSet:   make(map[string]*PhysicalMachine),
		vmSet:   make(map[string]*VirtualMachine),
		pmUUIDs: []string{},
		vmUUIDs: []string{},
	}

	err := g.generatePMs(numPM)
	if err != nil {
		return nil, fmt.Errorf("error creating pms: %v", err)
	}
	err = g.generateVMs(numVM)
	if err != nil {
		return nil, fmt.Errorf("error creating vms: %v", err)
	}

	return g, nil
}

func (this *TopologyGenerator) GetPMs() map[string]*PhysicalMachine {
	this.lock.RLock()
	defer this.lock.RUnlock()

	return this.pmSet
}

func (this *TopologyGenerator) GetVMs() map[string]*VirtualMachine {
	this.lock.RLock()
	defer this.lock.RUnlock()

	return this.vmSet
}

func (this *TopologyGenerator) UpdateResource() error {
	this.lock.Lock()
	defer this.lock.Unlock()

	glog.V(3).Infof("Updating resource usage in current topology.")

	// Update PM resource overhead.
	for _, pm := range this.pmSet {
		pm.ResourceStat = generatePMResourceStat()
	}

	for _, vm := range this.vmSet {
		provider, exists := this.pmSet[vm.providerID]
		if !exists {
			return fmt.Errorf("Non exist provider with UUID %s.", vm.providerID)
		}
		vm.ResourceStat = generateVMResourceStat(provider.ResourceStat)
	}

	return nil
}

func (this *TopologyGenerator) generatePMs(num int) error {
	for i := 0; i < num; i++ {
		name := fmt.Sprintf("PM_%d", i)

		uuid, err := generateUUID()
		if err != nil {
			return err
		}
		this.pmUUIDs = append(this.pmUUIDs, uuid)

		resourceStat := generatePMResourceStat()

		this.pmSet[uuid] = &PhysicalMachine{
			ObjectMeta{
				Name: name,
				UUID: uuid,
			},
			resourceStat,
			[]string{},
		}
	}
	return nil
}

func generatePMResourceStat() *PhysicalMachineResourceStat {
	// Generate with 0 to 10 percent overhead.
	return &PhysicalMachineResourceStat{
		cpuCapacity: defaultCPUCapacity,
		cpuUsed:     rand.Float64() * defaultCPUCapacity / 10,
		memCapacity: defaultMemoryCapacity,
		memUsed:     rand.Float64() * defaultMemoryCapacity / 10,
	}
}

func (this *TopologyGenerator) generateVMs(num int) error {
	for i := 0; i < num; i++ {
		name := fmt.Sprintf("VM_%d", i)

		uuid, err := generateUUID()
		if err != nil {
			return err
		}
		this.vmUUIDs = append(this.vmUUIDs, uuid)

		// Choose a provider
		pmUUID := this.pmUUIDs[i%len(this.pmUUIDs)]
		pmProvider, exist := this.pmSet[pmUUID]
		if !exist {
			return fmt.Errorf("Non exist PM provider with UUID %s", pmUUID)
		}
		pmProvider.consumersID = append(pmProvider.consumersID, uuid)
		pmProviderStat := pmProvider.ResourceStat

		resourceStat := generateVMResourceStat(pmProviderStat)

		this.vmSet[uuid] = &VirtualMachine{
			ObjectMeta{
				Name: name,
				UUID: uuid,
			},
			resourceStat,
			pmUUID,
		}
	}
	return nil
}

func generateVMResourceStat(pmProviderStat *PhysicalMachineResourceStat) *VirtualMachineResourceStat {
	randVCPUBase := pmProviderStat.cpuCapacity - pmProviderStat.cpuUsed
	if randVCPUBase > defaultVCPUCapacity {
		randVCPUBase = defaultVCPUCapacity
	}
	randVCPUUsed := rand.Float64() * randVCPUBase
	pmProviderStat.cpuUsed += randVCPUUsed

	randVMemBase := pmProviderStat.memCapacity - pmProviderStat.memUsed
	if randVMemBase > defaultVMemCapacity {
		randVMemBase = defaultVMemCapacity
	}
	randVMemUsed := rand.Float64() * randVMemBase
	pmProviderStat.memUsed += randVMemUsed

	return &VirtualMachineResourceStat{
		vCpuCapacity: defaultVCPUCapacity,
		vCpuUsed:     randVCPUUsed,
		vMemCapacity: defaultVMemCapacity,
		vMemUsed:     randVMemUsed,
	}

}

func generateUUID() (string, error) {
	size := 16
	rb := make([]byte, size)
	_, err := crand.Read(rb)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(rb), nil
}
