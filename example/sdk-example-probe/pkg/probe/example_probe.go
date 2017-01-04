package probe

import (
	"fmt"

	"github.com/turbonomic/turbo-go-sdk/pkg/builder"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

type ExampleProbe struct {
	topoSource *TopologyGenerator
}

func NewExampleProbe(source *TopologyGenerator) *ExampleProbe {
	return &ExampleProbe{
		topoSource: source,
	}
}

func (probe *ExampleProbe) Discover() ([]*proto.EntityDTO, error) {
	var discoveryResults []*proto.EntityDTO

	pmDTOs, err := probe.discoverPMs()
	if err != nil {
		return nil, fmt.Errorf("Error found during PM discovery: %s", err)
	}
	discoveryResults = append(discoveryResults, pmDTOs...)

	vmDTOs, err := probe.discoverVMs()
	if err != nil {
		return nil, fmt.Errorf("Error found during VM discovery: %s", err)
	}
	discoveryResults = append(discoveryResults, vmDTOs...)

	return discoveryResults, nil
}

func (probe *ExampleProbe) discoverPMs() ([]*proto.EntityDTO, error) {
	var result []*proto.EntityDTO

	pms := probe.topoSource.GetPMs()
	for _, pm := range pms {
		commoditiesSold, err := createPMCommoditiesSold(pm)
		if err != nil {
			return nil, err
		}
		entityDTO, err := builder.NewEntityDTOBuilder(proto.EntityDTO_PHYSICAL_MACHINE, pm.UUID).
			DisplayName(pm.Name).
			SellsCommodities(commoditiesSold).
			Create()
		if err != nil {
			return nil, fmt.Errorf("Error creating entityDTO for PM %s: %v", pm.Name, err)
		}
		result = append(result, entityDTO)
	}

	return result, nil
}

func createPMCommoditiesSold(pm *PhysicalMachine) ([]*proto.CommodityDTO, error) {
	var commoditiesSold []*proto.CommodityDTO
	pmResourceStat := pm.ResourceStat

	cpuComm, err := builder.NewCommodityDTOBuilder(proto.CommodityDTO_CPU).
		Capacity(pmResourceStat.cpuCapacity).
		Used(pmResourceStat.cpuUsed).
		Create()
	if err != nil {
		return nil, fmt.Errorf("Error creating CPU commodity sold by PM: %s", err)
	}
	commoditiesSold = append(commoditiesSold, cpuComm)

	memComm, err := builder.NewCommodityDTOBuilder(proto.CommodityDTO_MEM).
		Capacity(pmResourceStat.memCapacity).
		Used(pmResourceStat.memUsed).
		Create()
	if err != nil {
		return nil, fmt.Errorf("Error creating Memory commodity sold by PM: %s", err)
	}
	commoditiesSold = append(commoditiesSold, memComm)

	return commoditiesSold, nil
}

func (probe *ExampleProbe) discoverVMs() ([]*proto.EntityDTO, error) {
	var result []*proto.EntityDTO

	vms := probe.topoSource.GetVMs()
	for _, vm := range vms {
		commoditiesSold, err := createVMCommoditiesSold(vm)
		if err != nil {
			return nil, err
		}
		commoditiesBought, err := createVMCommoditiesBought(vm)
		if err != nil {
			return nil, err
		}
		entityDTO, err := builder.NewEntityDTOBuilder(proto.EntityDTO_VIRTUAL_MACHINE, vm.UUID).
			DisplayName(vm.Name).
			SellsCommodities(commoditiesSold).
			Provider(proto.EntityDTO_PHYSICAL_MACHINE, vm.providerID).
			BuysCommodities(commoditiesBought).
			Create()
		if err != nil {
			return nil, fmt.Errorf("Error creating entityDTO for VM %s: %v", vm.Name, err)
		}
		result = append(result, entityDTO)
	}

	return result, nil
}

func createVMCommoditiesSold(vm *VirtualMachine) ([]*proto.CommodityDTO, error) {
	var commoditiesSold []*proto.CommodityDTO
	vmResourceStat := vm.ResourceStat

	vCpuComm, err := builder.NewCommodityDTOBuilder(proto.CommodityDTO_VCPU).
		Capacity(vmResourceStat.vCpuCapacity).
		Used(vmResourceStat.vCpuUsed).
		Create()
	if err != nil {
		return nil, fmt.Errorf("Error creating VCPU commodity sold by VM: %s", err)
	}
	commoditiesSold = append(commoditiesSold, vCpuComm)

	vMemComm, err := builder.NewCommodityDTOBuilder(proto.CommodityDTO_VMEM).
		Capacity(vmResourceStat.vMemCapacity).
		Used(vmResourceStat.vMemUsed).
		Create()
	if err != nil {
		return nil, fmt.Errorf("Error creating VMEM commodity sold by VM: %s", err)
	}
	commoditiesSold = append(commoditiesSold, vMemComm)

	return commoditiesSold, nil
}

func createVMCommoditiesBought(vm *VirtualMachine) ([]*proto.CommodityDTO, error) {
	var commoditiesBought []*proto.CommodityDTO
	vCpuCommBought, err := builder.NewCommodityDTOBuilder(proto.CommodityDTO_CPU).
		Used(vm.ResourceStat.vCpuUsed).
		Create()
	if err != nil {
		return nil, fmt.Errorf("Error creating VCPU commodity bought by VM: %s", err)
	}
	commoditiesBought = append(commoditiesBought, vCpuCommBought)

	vMemCommBought, err := builder.NewCommodityDTOBuilder(proto.CommodityDTO_MEM).
		Used(vm.ResourceStat.vMemCapacity).
		Create()
	if err != nil {
		return nil, fmt.Errorf("Error creating VMEM commodity bought by VM: %s", err)
	}
	commoditiesBought = append(commoditiesBought, vMemCommBought)

	return commoditiesBought, nil
}
