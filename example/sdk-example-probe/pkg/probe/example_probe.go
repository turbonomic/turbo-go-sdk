package probe

import (
	"fmt"

	"github.com/vmturbo/vmturbo-go-sdk/pkg/builder"
	"github.com/vmturbo/vmturbo-go-sdk/pkg/proto"
)

type ExampleProbe struct {
	topoSource *TopologyGenerator
}

func NewExampleProbe(source *TopologyGenerator) *ExampleProbe {
	return &ExampleProbe{
		topoSource: source,
	}
}

func (this *ExampleProbe) Discover() ([]*proto.EntityDTO, error) {
	var discoveryResults []*proto.EntityDTO

	pmDTOs, err := this.discoverPMs()
	if err != nil {
		return nil, fmt.Errorf("Error found during PM discovery: %s", err)
	}
	discoveryResults = append(discoveryResults, pmDTOs...)

	vmDTOs, err := this.discoverVMs()
	if err != nil {
		return nil, fmt.Errorf("Error found during VM discovery: %s", err)
	}
	discoveryResults = append(discoveryResults, vmDTOs...)

	return discoveryResults, nil
}

func (this *ExampleProbe) discoverPMs() ([]*proto.EntityDTO, error) {
	var result []*proto.EntityDTO

	pms := this.topoSource.GetPMs()
	for _, pm := range pms {
		commoditiesSold := createPMCommoditiesSold(pm)

		entityDTO := builder.NewEntityDTOBuilder(proto.EntityDTO_PHYSICAL_MACHINE, pm.UUID).
			DisplayName(pm.Name).
			SellsCommodities(commoditiesSold).
			Create()
		result = append(result, entityDTO)
	}

	return result, nil
}

func createPMCommoditiesSold(pm *PhysicalMachine) []*proto.CommodityDTO {
	var commoditiesSold []*proto.CommodityDTO
	pmResourceStat := pm.ResourceStat

	cpuComm := builder.NewCommodityDTOBuilder(proto.CommodityDTO_CPU).
		Capacity(pmResourceStat.cpuCapacity).
		Used(pmResourceStat.cpuUsed).
		Create()
	commoditiesSold = append(commoditiesSold, cpuComm)

	memComm := builder.NewCommodityDTOBuilder(proto.CommodityDTO_MEM).
		Capacity(pmResourceStat.memCapacity).
		Used(pmResourceStat.memUsed).
		Create()
	commoditiesSold = append(commoditiesSold, memComm)

	return commoditiesSold
}

func (this *ExampleProbe) discoverVMs() ([]*proto.EntityDTO, error) {
	var result []*proto.EntityDTO

	vms := this.topoSource.GetVMs()
	for _, vm := range vms {
		commoditiesSold := createVMCommoditiesSold(vm)
		commoditiesBought := createVMCommoditiesBought(vm)

		entityDTO := builder.NewEntityDTOBuilder(proto.EntityDTO_VIRTUAL_MACHINE, vm.UUID).
			DisplayName(vm.Name).
			SellsCommodities(commoditiesSold).
			SetProviderWithTypeAndID(proto.EntityDTO_PHYSICAL_MACHINE, vm.providerID).
			BuysCommodities(commoditiesBought).
			Create()
		result = append(result, entityDTO)
	}

	return result, nil
}

func createVMCommoditiesSold(vm *VirtualMachine) []*proto.CommodityDTO {
	var commoditiesSold []*proto.CommodityDTO
	vmResourceStat := vm.ResourceStat

	vCpuComm := builder.NewCommodityDTOBuilder(proto.CommodityDTO_VCPU).
		Capacity(vmResourceStat.vCpuCapacity).
		Used(vmResourceStat.vCpuUsed).
		Create()
	commoditiesSold = append(commoditiesSold, vCpuComm)

	vMemComm := builder.NewCommodityDTOBuilder(proto.CommodityDTO_VMEM).
		Capacity(vmResourceStat.vMemCapacity).
		Used(vmResourceStat.vMemUsed).
		Create()
	commoditiesSold = append(commoditiesSold, vMemComm)

	return commoditiesSold
}

func createVMCommoditiesBought(vm *VirtualMachine) []*proto.CommodityDTO {
	var commoditiesBought []*proto.CommodityDTO
	vCpuCommBought := builder.NewCommodityDTOBuilder(proto.CommodityDTO_CPU).
		Used(vm.ResourceStat.vCpuUsed).
		Create()
	commoditiesBought = append(commoditiesBought, vCpuCommBought)

	vMemCommBought := builder.NewCommodityDTOBuilder(proto.CommodityDTO_MEM).
		Used(vm.ResourceStat.vMemCapacity).
		Create()
	commoditiesBought = append(commoditiesBought, vMemCommBought)

	return commoditiesBought
}
