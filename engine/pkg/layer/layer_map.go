package layer

type LayerMap map[Layer][]interface{}

func CreateLayerMap() LayerMap {
	return make(LayerMap)
}

func (m LayerMap) Add(layer Layer, d interface{}) {
	m[layer] = append(m[layer], d)
}
