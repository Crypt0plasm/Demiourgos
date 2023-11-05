package Bloodshed

import p "Firefly-APD"

func GetBloodshedDacian(Input []TTV) string {
	return Input[1].Value
}

func GetBloodshedPotency(Input []TTV) string {
	return Input[2].Value

}

func GetBloodshedTier(Input []TTV) string {
	return Input[3].Value
}

func GetBloodshedBackground(Input []TTV) string {
	return Input[4].Value
}

func GetBloodshedMainProtection(Input []TTV) string {
	return Input[5].Value
}

func GetBloodshedSecondaryProtection(Input []TTV) string {
	return Input[6].Value
}

func GetBloodshedMainHand(Input []TTV) string {
	return Input[7].Value
}

func GetBloodshedMOffHand(Input []TTV) string {
	return Input[8].Value
}

func GetWeekMultiplier(Input string) *p.Decimal {
	ExistsInSlice := func(Input string, Slice []string) bool {
		for _, element := range Slice {
			if element == Input {
				return true
			}
		}
		return false
	}

	var (
		WM *p.Decimal
	)
	V1 := p.NFS("1.1")
	V2 := p.NFS("1.3")
	V3 := p.NFS("1.6")
	V4 := p.NFS("2.0")
	V5 := p.NFS("2.5")
	V6 := p.NFS("3.1")

	if Input == Potency[0] || Input == Bloodshed[0] || ExistsInSlice(Input, CommonBG) == true || Input == Protection[0] || Input == Weapons[5] {
		WM = V1
	} else if Input == Potency[1] || Input == Bloodshed[1] || ExistsInSlice(Input, RareBG) == true || Input == Protection[1] || Input == Weapons[2] {
		WM = V2
	} else if Input == Potency[2] || Input == Bloodshed[2] || ExistsInSlice(Input, EpicBG) == true || Input == Protection[2] || Input == Weapons[1] {
		WM = V3
	} else if ExistsInSlice(Input, Dacian) == true || Input == Bloodshed[3] || ExistsInSlice(Input, LegendaryBG) == true || Input == Protection[3] || Input == Weapons[3] {
		WM = V4
	} else if Input == Bloodshed[4] || Input == Weapons[0] {
		WM = V5
	} else if Input == Bloodshed[5] || Input == Weapons[4] {
		WM = V6
	} else if Input == "None" || Input == "none" || Input == "NONE" {
		WM = p.NFS("1")
	}

	return WM
}
