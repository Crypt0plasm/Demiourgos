package Bloodshed

type BSJ struct {
	Name        string `json:"name"`        //Bloodshed
	Description string `json:"description"` //Description of Collection
	//Image       string `json:"image"`       //IPFS Link of Image
	Attributes []TTV `json:"attributes"`
}

type TTV struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

var (
	// Rarity trait
	R1 = TTV{TraitType: "Rarity", Value: NftRarity[0]} //Common
	R2 = TTV{TraitType: "Rarity", Value: NftRarity[1]} //Rare
	R3 = TTV{TraitType: "Rarity", Value: NftRarity[2]} //Epic
	R4 = TTV{TraitType: "Rarity", Value: NftRarity[3]} //Legendary

	// Dacian Trait
	D1 = TTV{TraitType: "Dacian", Value: Dacian[0]} //Dac1
	D2 = TTV{TraitType: "Dacian", Value: Dacian[1]} //Dac2
	D3 = TTV{TraitType: "Dacian", Value: Dacian[2]} //Dac3
	D4 = TTV{TraitType: "Dacian", Value: Dacian[3]} //Dac4
	D5 = TTV{TraitType: "Dacian", Value: Dacian[4]} //Dac5
	D6 = TTV{TraitType: "Dacian", Value: Dacian[5]} //Dac6
	D7 = TTV{TraitType: "Dacian", Value: Dacian[6]} //Dac7
	D8 = TTV{TraitType: "Dacian", Value: Dacian[7]} //Dac8

	//Potency
	P1 = TTV{TraitType: "Potency", Value: Potency[0]}
	P2 = TTV{TraitType: "Potency", Value: Potency[1]}
	P3 = TTV{TraitType: "Potency", Value: Potency[2]}

	//Bloodshed
	B0 = TTV{TraitType: "Bloodshed", Value: Bloodshed[0]}
	B1 = TTV{TraitType: "Bloodshed", Value: Bloodshed[1]}
	B2 = TTV{TraitType: "Bloodshed", Value: Bloodshed[2]}
	B3 = TTV{TraitType: "Bloodshed", Value: Bloodshed[3]}
	B4 = TTV{TraitType: "Bloodshed", Value: Bloodshed[4]}
	B5 = TTV{TraitType: "Bloodshed", Value: Bloodshed[5]}

	//Background
	CC1 = TTV{TraitType: "Background", Value: CommonBG[0]}
	CC2 = TTV{TraitType: "Background", Value: CommonBG[1]}
	CC3 = TTV{TraitType: "Background", Value: CommonBG[2]}
	CC4 = TTV{TraitType: "Background", Value: CommonBG[3]}
	CC5 = TTV{TraitType: "Background", Value: CommonBG[4]}
	CC6 = TTV{TraitType: "Background", Value: CommonBG[5]}
	RR1 = TTV{TraitType: "Background", Value: RareBG[0]}
	RR2 = TTV{TraitType: "Background", Value: RareBG[1]}
	RR3 = TTV{TraitType: "Background", Value: RareBG[2]}
	EP1 = TTV{TraitType: "Background", Value: EpicBG[0]}
	EP2 = TTV{TraitType: "Background", Value: EpicBG[1]}
	LG1 = TTV{TraitType: "Background", Value: LegendaryBG[0]}
	LG2 = TTV{TraitType: "Background", Value: LegendaryBG[1]}
	LG3 = TTV{TraitType: "Background", Value: LegendaryBG[2]}
	LG4 = TTV{TraitType: "Background", Value: LegendaryBG[3]}
	LG5 = TTV{TraitType: "Background", Value: LegendaryBG[4]}
	LG6 = TTV{TraitType: "Background", Value: LegendaryBG[5]}
	LG7 = TTV{TraitType: "Background", Value: LegendaryBG[6]}
	LG8 = TTV{TraitType: "Background", Value: LegendaryBG[7]}

	//Main Protection
	MP1 = TTV{TraitType: "1st Protection", Value: Protection[1]} //Bear-Skin
	MP2 = TTV{TraitType: "1st Protection", Value: Protection[3]} //Armor

	//Secondary Protection
	SP1 = TTV{TraitType: "2nd Protection", Value: Protection[0]} //Dacian-Skin
	SP2 = TTV{TraitType: "2nd Protection", Value: Protection[2]} //Shield

	//Main-Hand
	MH1 = TTV{TraitType: "Main-Hand", Value: Weapons[0]} //Cosor
	MH2 = TTV{TraitType: "Main-Hand", Value: Weapons[1]} //Falx
	MH3 = TTV{TraitType: "Main-Hand", Value: Weapons[4]} //Dacian-Draco

	//Off-Hand
	OH1 = TTV{TraitType: "Off-Hand", Value: Weapons[0]} //Cosor
	OH2 = TTV{TraitType: "Off-Hand", Value: Weapons[1]} //Falx
	OH3 = TTV{TraitType: "Off-Hand", Value: Weapons[2]} //Sicae
	OH4 = TTV{TraitType: "Off-Hand", Value: Weapons[3]} //Pavaza
	OH5 = TTV{TraitType: "Off-Hand", Value: Weapons[5]} //Howler
)

var (
	NftRarity     = []string{"Common", "Rare", "Epic", "Legendary"}
	Dacian        = []string{"Comati", "Ursoi", "Pileati", "Smardoi", "Carpian", "Tarabostes", "Costoboc", "Buridavens"}
	Potency       = []string{"Standard", "Premium", "Elite"}
	Bloodshed     = []string{"Tier 0", "Tier 1", "Tier 2", "Tier 3", "Tier 4", "Tier 5"}
	CommonBG      = []string{"Storm", "Grainfield", "Alpine", "Swamp", "Panonic", "Continental"}
	RareBG        = []string{"Rain", "Steppe", "Pontic"}
	EpicBG        = []string{"Fire", "Lightning"}
	LegendaryBG   = []string{"Wolven Trinity", "Dacian Gryphon", "Gryphon Phalera", "Unicorn Bird", "Eight-Legged Stag", "Ram Sacrifice", "Swirling Horses", "Twin Rams"}
	Protection    = []string{"Dacian-Skin", "Bear-Skin", "Shield", "Armor"}
	Weapons       = []string{"Cosor", "Falx", "Sicae", "Pavaza", "Dacian-Draco", "Howler"}
	BSName        = "Bloodshed"
	BSDescription = "A collection of 8861 NFTs depicting 272 unique Dacian Warriors representing the Bloodshed.gg, a Gaming Guild within Age of Zalmoxis, the 1st blockchain-based, Triple-A, Unreal-Engine 5 based, MMORPG."
)

var (
	//Common
	//DAC 1
	Common001 = []TTV{R1, D1, P1, B0, CC1, MP1, SP1, MH3, OH2}
	Common002 = []TTV{R1, D1, P1, B0, CC2, MP1, SP1, MH3, OH2}
	Common003 = []TTV{R1, D1, P1, B0, CC3, MP1, SP1, MH3, OH2}
	Common004 = []TTV{R1, D1, P1, B0, CC4, MP1, SP1, MH3, OH2}
	Common005 = []TTV{R1, D1, P1, B0, CC5, MP1, SP1, MH3, OH2}
	Common006 = []TTV{R1, D1, P1, B0, CC6, MP1, SP1, MH3, OH2}

	Common007 = []TTV{R1, D1, P2, B1, CC1, MP1, SP1, MH3, OH2}
	Common008 = []TTV{R1, D1, P2, B1, CC2, MP1, SP1, MH3, OH2}
	Common009 = []TTV{R1, D1, P2, B1, CC3, MP1, SP1, MH3, OH2}
	Common010 = []TTV{R1, D1, P2, B1, CC4, MP1, SP1, MH3, OH2}
	Common011 = []TTV{R1, D1, P2, B1, CC5, MP1, SP1, MH3, OH2}
	Common012 = []TTV{R1, D1, P2, B1, CC6, MP1, SP1, MH3, OH2}

	Common013 = []TTV{R1, D1, P3, B2, CC1, MP1, SP1, MH3, OH2}
	Common014 = []TTV{R1, D1, P3, B2, CC2, MP1, SP1, MH3, OH2}
	Common015 = []TTV{R1, D1, P3, B2, CC3, MP1, SP1, MH3, OH2}
	Common016 = []TTV{R1, D1, P3, B2, CC4, MP1, SP1, MH3, OH2}
	Common017 = []TTV{R1, D1, P3, B2, CC5, MP1, SP1, MH3, OH2}
	Common018 = []TTV{R1, D1, P3, B2, CC6, MP1, SP1, MH3, OH2}

	//DAC 2
	Common019 = []TTV{R1, D2, P1, B0, CC1, MP1, SP1, MH1, OH1}
	Common020 = []TTV{R1, D2, P1, B0, CC2, MP1, SP1, MH1, OH1}
	Common021 = []TTV{R1, D2, P1, B0, CC3, MP1, SP1, MH1, OH1}
	Common022 = []TTV{R1, D2, P1, B0, CC4, MP1, SP1, MH1, OH1}
	Common023 = []TTV{R1, D2, P1, B0, CC5, MP1, SP1, MH1, OH1}
	Common024 = []TTV{R1, D2, P1, B0, CC6, MP1, SP1, MH1, OH1}

	Common025 = []TTV{R1, D2, P2, B1, CC1, MP1, SP1, MH1, OH1}
	Common026 = []TTV{R1, D2, P2, B1, CC2, MP1, SP1, MH1, OH1}
	Common027 = []TTV{R1, D2, P2, B1, CC3, MP1, SP1, MH1, OH1}
	Common028 = []TTV{R1, D2, P2, B1, CC4, MP1, SP1, MH1, OH1}
	Common029 = []TTV{R1, D2, P2, B1, CC5, MP1, SP1, MH1, OH1}
	Common030 = []TTV{R1, D2, P2, B1, CC6, MP1, SP1, MH1, OH1}

	Common031 = []TTV{R1, D2, P3, B2, CC1, MP1, SP1, MH1, OH1}
	Common032 = []TTV{R1, D2, P3, B2, CC2, MP1, SP1, MH1, OH1}
	Common033 = []TTV{R1, D2, P3, B2, CC3, MP1, SP1, MH1, OH1}
	Common034 = []TTV{R1, D2, P3, B2, CC4, MP1, SP1, MH1, OH1}
	Common035 = []TTV{R1, D2, P3, B2, CC5, MP1, SP1, MH1, OH1}
	Common036 = []TTV{R1, D2, P3, B2, CC6, MP1, SP1, MH1, OH1}

	//DAC 3
	Common037 = []TTV{R1, D3, P1, B0, CC1, MP1, SP2, MH2, OH4}
	Common038 = []TTV{R1, D3, P1, B0, CC2, MP1, SP2, MH2, OH4}
	Common039 = []TTV{R1, D3, P1, B0, CC3, MP1, SP2, MH2, OH4}
	Common040 = []TTV{R1, D3, P1, B0, CC4, MP1, SP2, MH2, OH4}
	Common041 = []TTV{R1, D3, P1, B0, CC5, MP1, SP2, MH2, OH4}
	Common042 = []TTV{R1, D3, P1, B0, CC6, MP1, SP2, MH2, OH4}

	Common043 = []TTV{R1, D3, P2, B1, CC1, MP1, SP2, MH2, OH4}
	Common044 = []TTV{R1, D3, P2, B1, CC2, MP1, SP2, MH2, OH4}
	Common045 = []TTV{R1, D3, P2, B1, CC3, MP1, SP2, MH2, OH4}
	Common046 = []TTV{R1, D3, P2, B1, CC4, MP1, SP2, MH2, OH4}
	Common047 = []TTV{R1, D3, P2, B1, CC5, MP1, SP2, MH2, OH4}
	Common048 = []TTV{R1, D3, P2, B1, CC6, MP1, SP2, MH2, OH4}

	Common049 = []TTV{R1, D3, P3, B2, CC1, MP1, SP2, MH2, OH4}
	Common050 = []TTV{R1, D3, P3, B2, CC2, MP1, SP2, MH2, OH4}
	Common051 = []TTV{R1, D3, P3, B2, CC3, MP1, SP2, MH2, OH4}
	Common052 = []TTV{R1, D3, P3, B2, CC4, MP1, SP2, MH2, OH4}
	Common053 = []TTV{R1, D3, P3, B2, CC5, MP1, SP2, MH2, OH4}
	Common054 = []TTV{R1, D3, P3, B2, CC6, MP1, SP2, MH2, OH4}

	//DAC 4
	Common055 = []TTV{R1, D4, P1, B0, CC1, MP1, SP1, MH2, OH3}
	Common056 = []TTV{R1, D4, P1, B0, CC2, MP1, SP1, MH2, OH3}
	Common057 = []TTV{R1, D4, P1, B0, CC3, MP1, SP1, MH2, OH3}
	Common058 = []TTV{R1, D4, P1, B0, CC4, MP1, SP1, MH2, OH3}
	Common059 = []TTV{R1, D4, P1, B0, CC5, MP1, SP1, MH2, OH3}
	Common060 = []TTV{R1, D4, P1, B0, CC6, MP1, SP1, MH2, OH3}

	Common061 = []TTV{R1, D4, P2, B1, CC1, MP1, SP1, MH2, OH3}
	Common062 = []TTV{R1, D4, P2, B1, CC2, MP1, SP1, MH2, OH3}
	Common063 = []TTV{R1, D4, P2, B1, CC3, MP1, SP1, MH2, OH3}
	Common064 = []TTV{R1, D4, P2, B1, CC4, MP1, SP1, MH2, OH3}
	Common065 = []TTV{R1, D4, P2, B1, CC5, MP1, SP1, MH2, OH3}
	Common066 = []TTV{R1, D4, P2, B1, CC6, MP1, SP1, MH2, OH3}

	Common067 = []TTV{R1, D4, P3, B2, CC1, MP1, SP1, MH2, OH3}
	Common068 = []TTV{R1, D4, P3, B2, CC2, MP1, SP1, MH2, OH3}
	Common069 = []TTV{R1, D4, P3, B2, CC3, MP1, SP1, MH2, OH3}
	Common070 = []TTV{R1, D4, P3, B2, CC4, MP1, SP1, MH2, OH3}
	Common071 = []TTV{R1, D4, P3, B2, CC5, MP1, SP1, MH2, OH3}
	Common072 = []TTV{R1, D4, P3, B2, CC6, MP1, SP1, MH2, OH3}

	//DAC 5
	Common073 = []TTV{R1, D5, P1, B0, CC1, MP2, SP1, MH3, OH5}
	Common074 = []TTV{R1, D5, P1, B0, CC2, MP2, SP1, MH3, OH5}
	Common075 = []TTV{R1, D5, P1, B0, CC3, MP2, SP1, MH3, OH5}
	Common076 = []TTV{R1, D5, P1, B0, CC4, MP2, SP1, MH3, OH5}
	Common077 = []TTV{R1, D5, P1, B0, CC5, MP2, SP1, MH3, OH5}
	Common078 = []TTV{R1, D5, P1, B0, CC6, MP2, SP1, MH3, OH5}

	Common079 = []TTV{R1, D5, P2, B1, CC1, MP2, SP1, MH3, OH5}
	Common080 = []TTV{R1, D5, P2, B1, CC2, MP2, SP1, MH3, OH5}
	Common081 = []TTV{R1, D5, P2, B1, CC3, MP2, SP1, MH3, OH5}
	Common082 = []TTV{R1, D5, P2, B1, CC4, MP2, SP1, MH3, OH5}
	Common083 = []TTV{R1, D5, P2, B1, CC5, MP2, SP1, MH3, OH5}
	Common084 = []TTV{R1, D5, P2, B1, CC6, MP2, SP1, MH3, OH5}

	Common085 = []TTV{R1, D5, P3, B2, CC1, MP2, SP1, MH3, OH5}
	Common086 = []TTV{R1, D5, P3, B2, CC2, MP2, SP1, MH3, OH5}
	Common087 = []TTV{R1, D5, P3, B2, CC3, MP2, SP1, MH3, OH5}
	Common088 = []TTV{R1, D5, P3, B2, CC4, MP2, SP1, MH3, OH5}
	Common089 = []TTV{R1, D5, P3, B2, CC5, MP2, SP1, MH3, OH5}
	Common090 = []TTV{R1, D5, P3, B2, CC6, MP2, SP1, MH3, OH5}

	//DAC 6
	Common091 = []TTV{R1, D6, P1, B0, CC1, MP2, SP2, MH2, OH4}
	Common092 = []TTV{R1, D6, P1, B0, CC2, MP2, SP2, MH2, OH4}
	Common093 = []TTV{R1, D6, P1, B0, CC3, MP2, SP2, MH2, OH4}
	Common094 = []TTV{R1, D6, P1, B0, CC4, MP2, SP2, MH2, OH4}
	Common095 = []TTV{R1, D6, P1, B0, CC5, MP2, SP2, MH2, OH4}
	Common096 = []TTV{R1, D6, P1, B0, CC6, MP2, SP2, MH2, OH4}

	Common097 = []TTV{R1, D6, P2, B1, CC1, MP2, SP2, MH2, OH4}
	Common098 = []TTV{R1, D6, P2, B1, CC2, MP2, SP2, MH2, OH4}
	Common099 = []TTV{R1, D6, P2, B1, CC3, MP2, SP2, MH2, OH4}
	Common100 = []TTV{R1, D6, P2, B1, CC4, MP2, SP2, MH2, OH4}
	Common101 = []TTV{R1, D6, P2, B1, CC5, MP2, SP2, MH2, OH4}
	Common102 = []TTV{R1, D6, P2, B1, CC6, MP2, SP2, MH2, OH4}

	Common103 = []TTV{R1, D6, P3, B2, CC1, MP2, SP2, MH2, OH4}
	Common104 = []TTV{R1, D6, P3, B2, CC2, MP2, SP2, MH2, OH4}
	Common105 = []TTV{R1, D6, P3, B2, CC3, MP2, SP2, MH2, OH4}
	Common106 = []TTV{R1, D6, P3, B2, CC4, MP2, SP2, MH2, OH4}
	Common107 = []TTV{R1, D6, P3, B2, CC5, MP2, SP2, MH2, OH4}
	Common108 = []TTV{R1, D6, P3, B2, CC6, MP2, SP2, MH2, OH4}

	//DAC 7
	Common109 = []TTV{R1, D7, P1, B0, CC1, MP2, SP1, MH2, OH5}
	Common110 = []TTV{R1, D7, P1, B0, CC2, MP2, SP1, MH2, OH5}
	Common111 = []TTV{R1, D7, P1, B0, CC3, MP2, SP1, MH2, OH5}
	Common112 = []TTV{R1, D7, P1, B0, CC4, MP2, SP1, MH2, OH5}
	Common113 = []TTV{R1, D7, P1, B0, CC5, MP2, SP1, MH2, OH5}
	Common114 = []TTV{R1, D7, P1, B0, CC6, MP2, SP1, MH2, OH5}

	Common115 = []TTV{R1, D7, P2, B1, CC1, MP2, SP1, MH2, OH5}
	Common116 = []TTV{R1, D7, P2, B1, CC2, MP2, SP1, MH2, OH5}
	Common117 = []TTV{R1, D7, P2, B1, CC3, MP2, SP1, MH2, OH5}
	Common118 = []TTV{R1, D7, P2, B1, CC4, MP2, SP1, MH2, OH5}
	Common119 = []TTV{R1, D7, P2, B1, CC5, MP2, SP1, MH2, OH5}
	Common120 = []TTV{R1, D7, P2, B1, CC6, MP2, SP1, MH2, OH5}

	Common121 = []TTV{R1, D7, P3, B2, CC1, MP2, SP1, MH2, OH5}
	Common122 = []TTV{R1, D7, P3, B2, CC2, MP2, SP1, MH2, OH5}
	Common123 = []TTV{R1, D7, P3, B2, CC3, MP2, SP1, MH2, OH5}
	Common124 = []TTV{R1, D7, P3, B2, CC4, MP2, SP1, MH2, OH5}
	Common125 = []TTV{R1, D7, P3, B2, CC5, MP2, SP1, MH2, OH5}
	Common126 = []TTV{R1, D7, P3, B2, CC6, MP2, SP1, MH2, OH5}

	//DAC 8
	Common127 = []TTV{R1, D8, P1, B0, CC1, MP2, SP1, MH1, OH1}
	Common128 = []TTV{R1, D8, P1, B0, CC2, MP2, SP1, MH1, OH1}
	Common129 = []TTV{R1, D8, P1, B0, CC3, MP2, SP1, MH1, OH1}
	Common130 = []TTV{R1, D8, P1, B0, CC4, MP2, SP1, MH1, OH1}
	Common131 = []TTV{R1, D8, P1, B0, CC5, MP2, SP1, MH1, OH1}
	Common132 = []TTV{R1, D8, P1, B0, CC6, MP2, SP1, MH1, OH1}

	Common133 = []TTV{R1, D8, P2, B1, CC1, MP2, SP1, MH1, OH1}
	Common134 = []TTV{R1, D8, P2, B1, CC2, MP2, SP1, MH1, OH1}
	Common135 = []TTV{R1, D8, P2, B1, CC3, MP2, SP1, MH1, OH1}
	Common136 = []TTV{R1, D8, P2, B1, CC4, MP2, SP1, MH1, OH1}
	Common137 = []TTV{R1, D8, P2, B1, CC5, MP2, SP1, MH1, OH1}
	Common138 = []TTV{R1, D8, P2, B1, CC6, MP2, SP1, MH1, OH1}

	Common139 = []TTV{R1, D8, P3, B2, CC1, MP2, SP1, MH1, OH1}
	Common140 = []TTV{R1, D8, P3, B2, CC2, MP2, SP1, MH1, OH1}
	Common141 = []TTV{R1, D8, P3, B2, CC3, MP2, SP1, MH1, OH1}
	Common142 = []TTV{R1, D8, P3, B2, CC4, MP2, SP1, MH1, OH1}
	Common143 = []TTV{R1, D8, P3, B2, CC5, MP2, SP1, MH1, OH1}
	Common144 = []TTV{R1, D8, P3, B2, CC6, MP2, SP1, MH1, OH1}

	//Rare
	//Dac 1
	Rare01 = []TTV{R2, D1, P1, B1, RR1, MP1, SP1, MH3, OH2}
	Rare02 = []TTV{R2, D1, P1, B1, RR2, MP1, SP1, MH3, OH2}
	Rare03 = []TTV{R2, D1, P1, B1, RR3, MP1, SP1, MH3, OH2}

	Rare04 = []TTV{R2, D1, P2, B2, RR1, MP1, SP1, MH3, OH2}
	Rare05 = []TTV{R2, D1, P2, B2, RR2, MP1, SP1, MH3, OH2}
	Rare06 = []TTV{R2, D1, P2, B2, RR3, MP1, SP1, MH3, OH2}

	Rare07 = []TTV{R2, D1, P3, B3, RR1, MP1, SP1, MH3, OH2}
	Rare08 = []TTV{R2, D1, P3, B3, RR2, MP1, SP1, MH3, OH2}
	Rare09 = []TTV{R2, D1, P3, B3, RR3, MP1, SP1, MH3, OH2}

	//Dac 2
	Rare10 = []TTV{R2, D2, P1, B1, RR1, MP1, SP1, MH1, OH1}
	Rare11 = []TTV{R2, D2, P1, B1, RR2, MP1, SP1, MH1, OH1}
	Rare12 = []TTV{R2, D2, P1, B1, RR3, MP1, SP1, MH1, OH1}

	Rare13 = []TTV{R2, D2, P2, B2, RR1, MP1, SP1, MH1, OH1}
	Rare14 = []TTV{R2, D2, P2, B2, RR2, MP1, SP1, MH1, OH1}
	Rare15 = []TTV{R2, D2, P2, B2, RR3, MP1, SP1, MH1, OH1}

	Rare16 = []TTV{R2, D2, P3, B3, RR1, MP1, SP1, MH1, OH1}
	Rare17 = []TTV{R2, D2, P3, B3, RR2, MP1, SP1, MH1, OH1}
	Rare18 = []TTV{R2, D2, P3, B3, RR3, MP1, SP1, MH1, OH1}

	//Dac 3
	Rare19 = []TTV{R2, D3, P1, B1, RR1, MP1, SP2, MH2, OH4}
	Rare20 = []TTV{R2, D3, P1, B1, RR2, MP1, SP2, MH2, OH4}
	Rare21 = []TTV{R2, D3, P1, B1, RR3, MP1, SP2, MH2, OH4}

	Rare22 = []TTV{R2, D3, P2, B2, RR1, MP1, SP2, MH2, OH4}
	Rare23 = []TTV{R2, D3, P2, B2, RR2, MP1, SP2, MH2, OH4}
	Rare24 = []TTV{R2, D3, P2, B2, RR3, MP1, SP2, MH2, OH4}

	Rare25 = []TTV{R2, D3, P3, B3, RR1, MP1, SP2, MH2, OH4}
	Rare26 = []TTV{R2, D3, P3, B3, RR2, MP1, SP2, MH2, OH4}
	Rare27 = []TTV{R2, D3, P3, B3, RR3, MP1, SP2, MH2, OH4}

	//Dac 4
	Rare28 = []TTV{R2, D4, P1, B1, RR1, MP1, SP1, MH2, OH3}
	Rare29 = []TTV{R2, D4, P1, B1, RR2, MP1, SP1, MH2, OH3}
	Rare30 = []TTV{R2, D4, P1, B1, RR3, MP1, SP1, MH2, OH3}

	Rare31 = []TTV{R2, D4, P2, B2, RR1, MP1, SP1, MH2, OH3}
	Rare32 = []TTV{R2, D4, P2, B2, RR2, MP1, SP1, MH2, OH3}
	Rare33 = []TTV{R2, D4, P2, B2, RR3, MP1, SP1, MH2, OH3}

	Rare34 = []TTV{R2, D4, P3, B3, RR1, MP1, SP1, MH2, OH3}
	Rare35 = []TTV{R2, D4, P3, B3, RR2, MP1, SP1, MH2, OH3}
	Rare36 = []TTV{R2, D4, P3, B3, RR3, MP1, SP1, MH2, OH3}

	//Dac 5
	Rare37 = []TTV{R2, D5, P1, B1, RR1, MP2, SP1, MH3, OH5}
	Rare38 = []TTV{R2, D5, P1, B1, RR2, MP2, SP1, MH3, OH5}
	Rare39 = []TTV{R2, D5, P1, B1, RR3, MP2, SP1, MH3, OH5}

	Rare40 = []TTV{R2, D5, P2, B2, RR1, MP2, SP1, MH3, OH5}
	Rare41 = []TTV{R2, D5, P2, B2, RR2, MP2, SP1, MH3, OH5}
	Rare42 = []TTV{R2, D5, P2, B2, RR3, MP2, SP1, MH3, OH5}

	Rare43 = []TTV{R2, D5, P3, B3, RR1, MP2, SP1, MH3, OH5}
	Rare44 = []TTV{R2, D5, P3, B3, RR2, MP2, SP1, MH3, OH5}
	Rare45 = []TTV{R2, D5, P3, B3, RR3, MP2, SP1, MH3, OH5}

	//Dac 6
	Rare46 = []TTV{R2, D6, P1, B1, RR1, MP2, SP2, MH2, OH4}
	Rare47 = []TTV{R2, D6, P1, B1, RR2, MP2, SP2, MH2, OH4}
	Rare48 = []TTV{R2, D6, P1, B1, RR3, MP2, SP2, MH2, OH4}

	Rare49 = []TTV{R2, D6, P2, B2, RR1, MP2, SP2, MH2, OH4}
	Rare50 = []TTV{R2, D6, P2, B2, RR2, MP2, SP2, MH2, OH4}
	Rare51 = []TTV{R2, D6, P2, B2, RR3, MP2, SP2, MH2, OH4}

	Rare52 = []TTV{R2, D6, P3, B3, RR1, MP2, SP2, MH2, OH4}
	Rare53 = []TTV{R2, D6, P3, B3, RR2, MP2, SP2, MH2, OH4}
	Rare54 = []TTV{R2, D6, P3, B3, RR3, MP2, SP2, MH2, OH4}

	//Dac 7
	Rare55 = []TTV{R2, D7, P1, B1, RR1, MP2, SP1, MH2, OH5}
	Rare56 = []TTV{R2, D7, P1, B1, RR2, MP2, SP1, MH2, OH5}
	Rare57 = []TTV{R2, D7, P1, B1, RR3, MP2, SP1, MH2, OH5}

	Rare58 = []TTV{R2, D7, P2, B2, RR1, MP2, SP1, MH2, OH5}
	Rare59 = []TTV{R2, D7, P2, B2, RR2, MP2, SP1, MH2, OH5}
	Rare60 = []TTV{R2, D7, P2, B2, RR3, MP2, SP1, MH2, OH5}

	Rare61 = []TTV{R2, D7, P3, B3, RR1, MP2, SP1, MH2, OH5}
	Rare62 = []TTV{R2, D7, P3, B3, RR2, MP2, SP1, MH2, OH5}
	Rare63 = []TTV{R2, D7, P3, B3, RR3, MP2, SP1, MH2, OH5}

	//Dac 8
	Rare64 = []TTV{R2, D8, P1, B1, RR1, MP2, SP1, MH1, OH1}
	Rare65 = []TTV{R2, D8, P1, B1, RR2, MP2, SP1, MH1, OH1}
	Rare66 = []TTV{R2, D8, P1, B1, RR3, MP2, SP1, MH1, OH1}

	Rare67 = []TTV{R2, D8, P2, B2, RR1, MP2, SP1, MH1, OH1}
	Rare68 = []TTV{R2, D8, P2, B2, RR2, MP2, SP1, MH1, OH1}
	Rare69 = []TTV{R2, D8, P2, B2, RR3, MP2, SP1, MH1, OH1}

	Rare70 = []TTV{R2, D8, P3, B3, RR1, MP2, SP1, MH1, OH1}
	Rare71 = []TTV{R2, D8, P3, B3, RR2, MP2, SP1, MH1, OH1}
	Rare72 = []TTV{R2, D8, P3, B3, RR3, MP2, SP1, MH1, OH1}

	//Epic
	//Dac 1
	Epic01 = []TTV{R3, D1, P1, B2, EP1, MP1, SP1, MH3, OH2}
	Epic02 = []TTV{R3, D1, P1, B2, EP2, MP1, SP1, MH3, OH2}

	Epic03 = []TTV{R3, D1, P2, B3, EP1, MP1, SP1, MH3, OH2}
	Epic04 = []TTV{R3, D1, P2, B3, EP2, MP1, SP1, MH3, OH2}

	Epic05 = []TTV{R3, D1, P3, B4, EP1, MP1, SP1, MH3, OH2}
	Epic06 = []TTV{R3, D1, P3, B4, EP2, MP1, SP1, MH3, OH2}

	//Dac 2
	Epic07 = []TTV{R3, D2, P1, B2, EP1, MP1, SP1, MH1, OH1}
	Epic08 = []TTV{R3, D2, P1, B2, EP2, MP1, SP1, MH1, OH1}

	Epic09 = []TTV{R3, D2, P2, B3, EP1, MP1, SP1, MH1, OH1}
	Epic10 = []TTV{R3, D2, P2, B3, EP2, MP1, SP1, MH1, OH1}

	Epic11 = []TTV{R3, D2, P3, B4, EP1, MP1, SP1, MH1, OH1}
	Epic12 = []TTV{R3, D2, P3, B4, EP2, MP1, SP1, MH1, OH1}

	//Dac 3
	Epic13 = []TTV{R3, D3, P1, B2, EP1, MP1, SP2, MH2, OH4}
	Epic14 = []TTV{R3, D3, P1, B2, EP2, MP1, SP2, MH2, OH4}

	Epic15 = []TTV{R3, D3, P2, B3, EP1, MP1, SP2, MH2, OH4}
	Epic16 = []TTV{R3, D3, P2, B3, EP2, MP1, SP2, MH2, OH4}

	Epic17 = []TTV{R3, D3, P3, B4, EP1, MP1, SP2, MH2, OH4}
	Epic18 = []TTV{R3, D3, P3, B4, EP2, MP1, SP2, MH2, OH4}

	//Dac 4
	Epic19 = []TTV{R3, D4, P1, B2, EP1, MP1, SP1, MH2, OH3}
	Epic20 = []TTV{R3, D4, P1, B2, EP2, MP1, SP1, MH2, OH3}

	Epic21 = []TTV{R3, D4, P2, B3, EP1, MP1, SP1, MH2, OH3}
	Epic22 = []TTV{R3, D4, P2, B3, EP2, MP1, SP1, MH2, OH3}

	Epic23 = []TTV{R3, D4, P3, B4, EP1, MP1, SP1, MH2, OH3}
	Epic24 = []TTV{R3, D4, P3, B4, EP2, MP1, SP1, MH2, OH3}

	//Dac 5
	Epic25 = []TTV{R3, D5, P1, B2, EP1, MP2, SP1, MH3, OH5}
	Epic26 = []TTV{R3, D5, P1, B2, EP2, MP2, SP1, MH3, OH5}

	Epic27 = []TTV{R3, D5, P2, B3, EP1, MP2, SP1, MH3, OH5}
	Epic28 = []TTV{R3, D5, P2, B3, EP2, MP2, SP1, MH3, OH5}

	Epic29 = []TTV{R3, D5, P3, B4, EP1, MP2, SP1, MH3, OH5}
	Epic30 = []TTV{R3, D5, P3, B4, EP2, MP2, SP1, MH3, OH5}

	//Dac 6
	Epic31 = []TTV{R3, D6, P1, B2, EP1, MP2, SP2, MH2, OH4}
	Epic32 = []TTV{R3, D6, P1, B2, EP2, MP2, SP2, MH2, OH4}

	Epic33 = []TTV{R3, D6, P2, B3, EP1, MP2, SP2, MH2, OH4}
	Epic34 = []TTV{R3, D6, P2, B3, EP2, MP2, SP2, MH2, OH4}

	Epic35 = []TTV{R3, D6, P3, B4, EP1, MP2, SP2, MH2, OH4}
	Epic36 = []TTV{R3, D6, P3, B4, EP2, MP2, SP2, MH2, OH4}

	//Dac 7
	Epic37 = []TTV{R3, D7, P1, B2, EP1, MP2, SP1, MH2, OH5}
	Epic38 = []TTV{R3, D7, P1, B2, EP2, MP2, SP1, MH2, OH5}

	Epic39 = []TTV{R3, D7, P2, B3, EP1, MP2, SP1, MH2, OH5}
	Epic40 = []TTV{R3, D7, P2, B3, EP2, MP2, SP1, MH2, OH5}

	Epic41 = []TTV{R3, D7, P3, B4, EP1, MP2, SP1, MH2, OH5}
	Epic42 = []TTV{R3, D7, P3, B4, EP2, MP2, SP1, MH2, OH5}

	//Dac 8
	Epic43 = []TTV{R3, D8, P1, B2, EP1, MP2, SP1, MH1, OH1}
	Epic44 = []TTV{R3, D8, P1, B2, EP2, MP2, SP1, MH1, OH1}

	Epic45 = []TTV{R3, D8, P2, B3, EP1, MP2, SP1, MH1, OH1}
	Epic46 = []TTV{R3, D8, P2, B3, EP2, MP2, SP1, MH1, OH1}

	Epic47 = []TTV{R3, D8, P3, B4, EP1, MP2, SP1, MH1, OH1}
	Epic48 = []TTV{R3, D8, P3, B4, EP2, MP2, SP1, MH1, OH1}

	//Legendary all Dacs
	Legendary1 = []TTV{R4, D1, P3, B5, LG1, MP1, SP1, MH3, OH2}
	Legendary2 = []TTV{R4, D2, P3, B5, LG2, MP1, SP1, MH1, OH1}
	Legendary3 = []TTV{R4, D3, P3, B5, LG3, MP1, SP2, MH2, OH4}
	Legendary4 = []TTV{R4, D4, P3, B5, LG4, MP1, SP1, MH2, OH3}
	Legendary5 = []TTV{R4, D5, P3, B5, LG5, MP2, SP1, MH3, OH5}
	Legendary6 = []TTV{R4, D6, P3, B5, LG6, MP2, SP2, MH2, OH4}
	Legendary7 = []TTV{R4, D7, P3, B5, LG7, MP2, SP1, MH2, OH5}
	Legendary8 = []TTV{R4, D8, P3, B5, LG8, MP2, SP1, MH1, OH1}
)

var (
	LT = [][]TTV{Legendary1, Legendary2, Legendary3, Legendary4, Legendary5, Legendary6, Legendary7, Legendary8}
	ET = [][]TTV{
		Epic01, Epic02, Epic03, Epic04, Epic05, Epic06, Epic07, Epic08, Epic09, Epic10,
		Epic11, Epic12, Epic13, Epic14, Epic15, Epic16, Epic17, Epic18, Epic19, Epic20,
		Epic21, Epic22, Epic23, Epic24, Epic25, Epic26, Epic27, Epic28, Epic29, Epic30,
		Epic31, Epic32, Epic33, Epic34, Epic35, Epic36, Epic37, Epic38, Epic39, Epic40,
		Epic41, Epic42, Epic43, Epic44, Epic45, Epic46, Epic47, Epic48}

	RT = [][]TTV{
		Rare01, Rare02, Rare03, Rare04, Rare05, Rare06, Rare07, Rare08, Rare09, Rare10,
		Rare11, Rare12, Rare13, Rare14, Rare15, Rare16, Rare17, Rare18, Rare19, Rare20,
		Rare21, Rare22, Rare23, Rare24, Rare25, Rare26, Rare27, Rare28, Rare29, Rare30,
		Rare31, Rare32, Rare33, Rare34, Rare35, Rare36, Rare37, Rare38, Rare39, Rare40,
		Rare41, Rare42, Rare43, Rare44, Rare45, Rare46, Rare47, Rare48, Rare49, Rare50,
		Rare51, Rare52, Rare53, Rare54, Rare55, Rare56, Rare57, Rare58, Rare59, Rare60,
		Rare61, Rare62, Rare63, Rare64, Rare65, Rare66, Rare67, Rare68, Rare69, Rare70,
		Rare71, Rare72}

	CT = [][]TTV{
		Common001, Common002, Common003, Common004, Common005, Common006, Common007, Common008, Common009, Common010,
		Common011, Common012, Common013, Common014, Common015, Common016, Common017, Common018, Common019, Common020,
		Common021, Common022, Common023, Common024, Common025, Common026, Common027, Common028, Common029, Common030,
		Common031, Common032, Common033, Common034, Common035, Common036, Common037, Common038, Common039, Common040,
		Common041, Common042, Common043, Common044, Common045, Common046, Common047, Common048, Common049, Common050,
		Common051, Common052, Common053, Common054, Common055, Common056, Common057, Common058, Common059, Common060,
		Common061, Common062, Common063, Common064, Common065, Common066, Common067, Common068, Common069, Common070,
		Common071, Common072, Common073, Common074, Common075, Common076, Common077, Common078, Common079, Common080,
		Common081, Common082, Common083, Common084, Common085, Common086, Common087, Common088, Common089, Common090,
		Common091, Common092, Common093, Common094, Common095, Common096, Common097, Common098, Common099, Common100,
		Common101, Common102, Common103, Common104, Common105, Common106, Common107, Common108, Common109, Common110,
		Common111, Common112, Common113, Common114, Common115, Common116, Common117, Common118, Common119, Common120,
		Common121, Common122, Common123, Common124, Common125, Common126, Common127, Common128, Common129, Common130,
		Common131, Common132, Common133, Common134, Common135, Common136, Common137, Common138, Common139, Common140,
		Common141, Common142, Common143, Common144}
)
