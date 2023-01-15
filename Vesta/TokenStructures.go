package Vesta

// Can be made with the following request url
// "https://api.multiversx.com/accounts/erd1h6lh2tqjscs4n69c4w4wunu4qw2mz708qn8mqk4quzsyz2syn0aq5gu64s/tokens/WEGLD-bd4d79"
// Then paste here: https://mholt.github.io/json-to-go/

type ESDTSuperStructure struct {
	Type          string `json:"type"`
	Identifier    string `json:"identifier"`
	Name          string `json:"name"`
	Ticker        string `json:"ticker"`
	Owner         string `json:"owner"`
	Minted        string `json:"minted"`
	Burnt         string `json:"burnt"`
	InitialMinted string `json:"initialMinted"`
	Decimals      int    `json:"decimals"`
	IsPaused      bool   `json:"isPaused"`
	Assets        struct {
		Website         string `json:"website"`
		Description     string `json:"description"`
		Status          string `json:"status"`
		PngURL          string `json:"pngUrl"`
		SvgURL          string `json:"svgUrl"`
		LedgerSignature string `json:"ledgerSignature"`
		Social          struct {
			Email      string `json:"email"`
			Twitter    string `json:"twitter"`
			Whitepaper string `json:"whitepaper"`
			Coingecko  string `json:"coingecko"`
			Discord    string `json:"discord"`
			Telegram   string `json:"telegram"`
		} `json:"social"`
		LockedAccounts         string   `json:"lockedAccounts"`
		ExtraTokens            []string `json:"extraTokens"`
		PreferredRankAlgorithm string   `json:"preferredRankAlgorithm"`
	} `json:"assets"`
	Transactions             int     `json:"transactions"`
	Accounts                 int     `json:"accounts"`
	CanUpgrade               bool    `json:"canUpgrade"`
	CanMint                  bool    `json:"canMint"`
	CanBurn                  bool    `json:"canBurn"`
	CanChangeOwner           bool    `json:"canChangeOwner"`
	CanPause                 bool    `json:"canPause"`
	CanFreeze                bool    `json:"canFreeze"`
	CanWipe                  bool    `json:"canWipe"`
	CanTransferNftCreateRole bool    `json:"canTransferNftCreateRole"`
	Price                    float64 `json:"price"`
	MarketCap                float64 `json:"marketCap"`
	Supply                   string  `json:"supply"`
	CirculatingSupply        string  `json:"circulatingSupply"`
	Timestamp                int     `json:"timestamp"`
	Balance                  string  `json:"balance"`
	ValueUsd                 float64 `json:"valueUsd"`
	Attributes               string  `json:"attributes"`
}
