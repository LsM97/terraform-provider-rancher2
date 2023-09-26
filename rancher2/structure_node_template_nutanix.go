package rancher2

// Flatteners

func flattenNutanixConfig(in *nutanixConfig) []interface{} {
	obj := make(map[string]interface{})
	if in == nil {
		return []interface{}{}
	}
	if len(in.Cloudinit) > 0 {
		obj["cloudinit"] = in.Cloudinit
	}
	if len(in.Password) > 0 {
		obj["password"] = in.Password
	}
	if len(in.Username) > 0 {
		obj["username"] = in.Username
	}
	obj["insecure"] = in.Insecure
	if len(in.Cluster) > 0 {
		obj["cluster"] = in.Cluster
	}
	if len(in.DiskSize) > 0 {
		obj["diskSize"] = in.DiskSize
	}
	if len(in.Endpoint) > 0 {
		obj["endpoint"] = in.Endpoint
	}
	if len(in.Port) > 0 {
		obj["port"] = in.Port
	}
	if len(in.Project) > 0 {
		obj["project"] = in.Project
	}
	if len(in.StorageContainer) > 0 {
		obj["storageContainer"] = in.StorageContainer
	}
	if len(in.VmCategories) > 0 {
		obj["vmCategories"] = toArrayInterface(in.VmCategories)
	}
	if len(in.VmCores) > 0 {
		obj["vmCores"] = in.VmCores
	}
	obj["vmCpuPassthrough"] = in.VmCpuPassthrough
	if len(in.VmCpus) > 0 {
		obj["vmCpus"] = in.VmCpus
	}
	if len(in.VmImage) > 0 {
		obj["vmImage"] = in.VmImage
	}
	if len(in.VmImageSize) > 0 {
		obj["vmImageSize"] = in.VmImageSize
	}
	if len(in.VmMem) > 0 {
		obj["vmMem"] = in.VmMem
	}
	if len(in.VmNetwork) > 0 {
		obj["vmNetwork"] = toArrayInterface(in.VmNetwork)
	}
	obj["vmSerialPort"] = in.VmSerialPort

	return []interface{}{obj}
}

// Expanders

func expandNutanixConfig(p []interface{}) *nutanixConfig {
	obj := &nutanixConfig{}
	if len(p) == 0 || p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["cloudinit"].(string); ok && len(v) > 0 {
		obj.Cloudinit = v
	}
	if v, ok := in["password"].(string); ok && len(v) > 0 {
		obj.Password = v
	}
	if v, ok := in["username"].(string); ok && len(v) > 0 {
		obj.Username = v
	}
	if v, ok := in["insecure"].(bool); ok {
		obj.Insecure = v
	}
	if v, ok := in["cluster"].(string); ok && len(v) > 0 {
		obj.Cluster = v
	}
	if v, ok := in["diskSize"].(string); ok && len(v) > 0 {
		obj.DiskSize = v
	}
	if v, ok := in["endpoint"].(string); ok && len(v) > 0 {
		obj.Endpoint = v
	}
	if v, ok := in["port"].(string); ok && len(v) > 0 {
		obj.Port = v
	}
	if v, ok := in["project"].(string); ok && len(v) > 0 {
		obj.Project = v
	}
	if v, ok := in["storageContainer"].(string); ok && len(v) > 0 {
		obj.StorageContainer = v
	}
	if v, ok := in["vmCategories"].([]interface{}); ok && len(v) > 0 {
		obj.VmCategories = toArrayString(v)
	}
	if v, ok := in["vmCores"].(string); ok && len(v) > 0 {
		obj.VmCores = v
	}
	if v, ok := in["vmCpuPassthrough"].(bool); ok {
		obj.VmCpuPassthrough = v
	}
	if v, ok := in["vmCpus"].(string); ok && len(v) > 0 {
		obj.VmCpus = v
	}
	if v, ok := in["vmImage"].(string); ok && len(v) > 0 {
		obj.VmImage = v
	}
	if v, ok := in["vmImageSize"].(string); ok && len(v) > 0 {
		obj.VmImageSize = v
	}
	if v, ok := in["vmMem"].(string); ok && len(v) > 0 {
		obj.VmMem = v
	}
	if v, ok := in["vmNetwork"].([]interface{}); ok && len(v) > 0 {
		obj.VmNetwork = toArrayString(v)
	}
	if v, ok := in["vmSerialPort"].(bool); ok {
		obj.VmSerialPort = v
	}

	return obj
}
