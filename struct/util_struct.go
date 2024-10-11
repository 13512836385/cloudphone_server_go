package struct_package

type VolcConfig struct {
	AK string `json:"ak"`
	SK string `json:"sk"`
}

type AcepConfig struct {
	ProductID string   `json:"product_id"`
	PodIDList []string `json:"pod_id_list"`
	PodID     string   `json:"pod_id"`
}
