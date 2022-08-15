package server

type LedControlDTO struct {
	Target string `json:"target"`
	Order  string `json:"order"`
}

func (d *LedControlDTO) Validate() (bool, string) {

	if d.Order != "on" && d.Order != "off" {
		return false, "please check order again"
	}

	if d.Target != "led" {
		return false, "please check target again"
	}
	return true, ""
}