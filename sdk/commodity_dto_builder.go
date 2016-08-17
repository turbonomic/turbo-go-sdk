package sdk

type CommodityDTOBuilder struct {
	commodityType      CommodityDTO_CommodityType
	key                string
	used               float64
	reservation        float64
	capacity           float64
	limit              float64
	peak               float64
	active             bool
	resizable          bool
	displayName        string
	thin               bool
	computedUsed       bool
	usedIncrement      float64
	storageLatencyData CommodityDTO_StorageLatencyData
	storageAccessData  CommodityDTO_StorageAccessData
}

func NewCommodtiyDTOBuilder(commodityType CommodityDTO_CommodityType) *CommodityDTOBuilder {
	return &CommodityDTOBuilder{
		commodityType: commodityType,
	}
}

func (this *CommodityDTOBuilder) Create() *CommodityDTO {
	return &CommodityDTO{
		CommodityType:      &this.commodityType,
		Key:                &this.key,
		Used:               &this.used,
		Reservation:        &this.reservation,
		Capacity:           &this.capacity,
		Limit:              &this.limit,
		Peak:               &this.peak,
		Active:             &this.active,
		Resizable:          &this.resizable,
		DisplayName:        &this.displayName,
		Thin:               &this.thin,
		ComputedUsed:       &this.computedUsed,
		UsedIncrement:      &this.usedIncrement,
		StorageLatencyData: &this.storageLatencyData,
		StorageAccessData:  &this.storageAccessData,
	}
}

func (this *CommodityDTOBuilder) Key(key string) *CommodityDTOBuilder {
	this.key = key
	return this
}

func (this *CommodityDTOBuilder) Capacity(capcity float64) *CommodityDTOBuilder {
	this.capacity = capcity
	return this
}

func (this *CommodityDTOBuilder) Used(used float64) *CommodityDTOBuilder {
	this.used = used
	return this
}

func (this *CommodityDTOBuilder) Active(active bool) *CommodityDTOBuilder {
	this.active = active
	return this
}

func (this *CommodityDTOBuilder) ComputedUsed(computedUsed bool) *CommodityDTOBuilder {
	this.computedUsed = computedUsed
	return this
}

func (this *CommodityDTOBuilder) DisplayName(displayName string) *CommodityDTOBuilder {
	this.displayName = displayName
	return this
}

func (this *CommodityDTOBuilder) Limit(limit float64) *CommodityDTOBuilder {
	this.limit = limit
	return this
}

func (this *CommodityDTOBuilder) Peak(peak float64) *CommodityDTOBuilder {
	this.peak = peak
	return this
}

func (this *CommodityDTOBuilder) Reservation(r float64) *CommodityDTOBuilder {
	this.reservation = r
	return this
}

func (this *CommodityDTOBuilder) Resizable(resizable bool) *CommodityDTOBuilder {
	this.resizable = resizable
	return this
}

func (this *CommodityDTOBuilder) StorageAccessData(data CommodityDTO_StorageAccessData) *CommodityDTOBuilder {
	this.storageAccessData = data
	return this
}

func (this *CommodityDTOBuilder) StorageLatencyData(data CommodityDTO_StorageLatencyData) *CommodityDTOBuilder {
	this.storageLatencyData = data
	return this
}

func (this *CommodityDTOBuilder) Thin(isThin bool) *CommodityDTOBuilder {
	this.thin = isThin
	return this
}

func (this *CommodityDTOBuilder) UsedIncrement(increment float64) *CommodityDTOBuilder {
	this.usedIncrement = increment
	return this
}
