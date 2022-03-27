package metric

type Resources struct {
	CPUs   int //m, 1c = 1000m
	Memory int //byte, 1<<10 == 1k
	Disk   int //byte, 1<<10 == 1k
}

func (i *Resources) Add(r *Resources) {
	if r == nil {
		return
	}
	i.CPUs += r.CPUs
	i.Memory += r.Memory
	i.Disk += r.Disk
}

func (i *Resources) Times(n int) {
	if i.CPUs > 0 {
		i.CPUs *= n
	}
	if i.Memory > 0 {
		i.Memory *= n
	}
	if i.Disk > 0 {
		i.Disk *= n
	}
}
