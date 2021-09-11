package refine

import "fmt"

type Resource struct {
	zeny   int
	al     int
	alPlus int
	fire   int
}

func (r *Resource) refine(level int) {
	if level > 9 {
		r.zeny += 95000
	} else {
		r.zeny += (level + 1) * 9500
	}
	if level < 10 {
		r.al++
	} else {
		r.alPlus++
	}
}

func (r *Resource) repair() {
	r.fire++
}

func (r *Resource) sub(o *Resource) *Resource {
	return &Resource{
		zeny:   r.zeny - o.zeny,
		al:     r.al - o.al,
		alPlus: r.alPlus - o.alPlus,
		fire:   r.fire - o.fire,
	}
}

func (r *Resource) add(o *Resource) {
	r.zeny += o.zeny
	r.al += o.al
	r.alPlus += o.alPlus
	r.fire += o.fire
}

func (r *Resource) String() string {
	return fmt.Sprintf("Zeny: %d, 铝: %d, 强化铝: %d, 炉灰: %d", r.zeny, r.al, r.alPlus, r.fire)
}
