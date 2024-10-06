package loadbalance

type WeightRoundRobinBalance struct {
	curIndex int
	rss      []*WeightNode
}

type WeightNode struct {
	addr            int
	weight          int //权重值
	currentWeight   int //节点当前权重
	effectiveWeight int //有效权重
}

func (r *WeightRoundRobinBalance) Add(addr int, weight int) error {

	node := &WeightNode{addr: addr, weight: weight}
	node.effectiveWeight = node.weight
	r.rss = append(r.rss, node)
	return nil
}

// Next 参考了 nginx 的加权负载均衡的策略
func (r *WeightRoundRobinBalance) Next() int {
	total := 0
	var best *WeightNode
	for i := 0; i < len(r.rss); i++ {
		w := r.rss[i]
		//step 1 统计所有有效权重之和
		total += w.effectiveWeight

		//step 2 变更节点临时权重为的节点临时权重+节点有效权重
		w.currentWeight += w.effectiveWeight

		//step 3 有效权重默认与权重相同，通讯异常时-1, 通讯成功+1，直到恢复到weight大小
		if w.effectiveWeight < w.weight {
			w.effectiveWeight++
		}
		//step 4 选择最大临时权重点节点
		if best == nil || w.currentWeight > best.currentWeight {
			best = w
		}
	}
	if best == nil {
		return -1
	}
	//step 5 变更临时权重为 临时权重-有效权重之和
	best.currentWeight -= total
	return best.addr
}

func (r *WeightRoundRobinBalance) Get() (int, error) {
	return r.Next(), nil
}