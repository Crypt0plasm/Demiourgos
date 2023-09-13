package Vesta

import (
    p "Firefly-APD"
    mvx "MvxApiScanner"
    sm "SuperMath"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type VestaHoldings struct {
    Address mvx.MvxAddress
    Gold    int64
    Silver  int64
    Bronze  int64
}

type LpHoldings struct {
    Address  mvx.MvxAddress
    VestaLiq VestaDexLPs
    KosonLiq VestaDexLPs
    BloodLiq VestaDexLPs
    Vault    VestaDexVault
    SnakeLiq OuroLP
}

type OuroLP struct {
    Native *p.Decimal
    Frozen *p.Decimal
}

type VestaDexVault struct {
    Vesta    *p.Decimal
    Blessed  *p.Decimal
    Sleeping *p.Decimal
    Frozen   *p.Decimal
}

type VestaDexLPs struct {
    Gold    *p.Decimal
    Silver  *p.Decimal
    Bronze  *p.Decimal
    UGold   *p.Decimal
    USilver *p.Decimal
    UBronze *p.Decimal
}

type VestaLPDesignation struct {
    Gold    string
    Silver  string
    Bronze  string
    UGold   string
    USilver string
    UBronze string
}

type FarmMx struct {
    DEB *p.Decimal
    VLM *p.Decimal
    TM  *p.Decimal
    D   *p.Decimal
    UM  *p.Decimal
    IM  *p.Decimal
    DM  *p.Decimal
    VM  *p.Decimal
}

var (
    Neutral = p.NFS("1")

    AncientDEB   = p.NFS("2.5")
    BloodshedDEB = p.NFS("1.6")

    VestaTMSix = p.NFS("6")

    Platinum = p.NFS("1.4")

    Zero  = p.NFS("0")
    Empty = "empty"

    AncientHodler = mvx.AH
    TrDaniel      = mvx.MvxAddress("erd1hg9q84tyzxretw2a8nce6q3lwgfzku587ndwr5k7202xt5pw4vyqp76vxe")
    DRX           = mvx.MvxAddress("erd1gqh79mxfr4al0gfpplvm79lxrx93tgclm3kvlfklmfgr3sp3yx8s7d3qjp")
    Patryx        = mvx.MvxAddress("erd1k64gwm43pqtydv978a5h27zfs48av7hq79j6e6uxz3dlg704vdyq6kuzym")
    Lavinia       = mvx.MvxAddress("erd1xjutp9aj4amd8r5mtedl4aad6jrxv2ajzw5d4x0c3rkwj7cxsgmssdjr49")
    Sandu         = mvx.MvxAddress("erd1pkxygrm9dswrludcxjk9hqmep62gutz7vfknlfpwzufxyara27sqjyhf4h")
    Cuciorva      = mvx.MvxAddress("erd1knx4hu2u6zxyt2mqk9zfyf8m9pz980puw98hqsmv26d3eggxvfmsv4xlpq")
    Codarcea      = mvx.MvxAddress("erd1fvddlyu2jtlhkfvc88qswh7x5vf3583knmhcl4vgwxncht4j9qjspqxguq")
    Pulecs        = mvx.MvxAddress("erd1wa7h9ea30j0enjm8k9x8ymf0s334qde8x4c3fpzkgj28xfkyt7nqdttpst")
    Laurentiu     = mvx.MvxAddress("erd1mf2wtc4jh2zujhu4nynvrpaua9e98c4lwdyjnpp57qcx2gyqfy6q8l8ccp")
    Frostedk9     = mvx.MvxAddress("erd1ktu3qy5ehe42sk6z7ygfvwna6wull2suq49qj8urx9nd6dw79s2qn5qqea")
    IonutDRD      = mvx.MvxAddress("erd1ez8ww48xj9gr6yyzem7g7gdvknjdh3te7lcyanz9tkrzyazyzxeqzssrhx")
    Buhaici       = mvx.MvxAddress("erd18n5y3884pdkeq6gl0wng22z2yqexhemwkyewcksesn2vgmgcqxes0e93cc")
    TheKid        = mvx.MvxAddress("erd1zl890854dweghll9faf67ft26965v8u20d6z63cntr9a5ykhcekqmyzcsf")
    RaulTM        = mvx.MvxAddress("erd1640c9n2cck2326jy0tg87nawhgxdxvzqg9psezg2yjzkxvg6gkcqzfdp6j")
    MakeAStep     = mvx.MvxAddress("erd1heus28d80kkengfz4ltn2m9xhvd4th8pajrsg9hkeeu202cjpfwq03m38w")
    Paul          = mvx.MvxAddress("erd1vj40fxw0yah34mmdxly7l28w097ju6hf8pczpcdxs05n2vyx8hcspyxm2c")
    Florian       = mvx.MvxAddress("erd1005uhtflxhql2cqyvw2y064k3fdjtjpcku0g3z25lr9znx8sd26sjulgt0")
    Coding        = mvx.MvxAddress("erd1qe8kudxwzen5hgxcmws9jrrtg6au97j974gtgrml6amnzlmmcetsv02gps")
    Bloodshed     = mvx.MvxAddress("erd16f3qfmpdhcgdv2ygwj43n2x08vnmfckvq8z976cs85ued5tx003scj60vd")
    Elanul        = mvx.MvxAddress("erd1phh72v0evuphdk8uwsg9ph9mr3gm8ucklvpr808ky3jqnnf3uqkqn6n8xg")
    Buguletu      = mvx.MvxAddress("erd1uq6kl4qtzd2fy42ad7puyf29jf6l25kzfmthv3wqu5jmz8dm2fmqk08n35")
    Carlig        = mvx.MvxAddress("erd13gjrsayl6atc660jyz9anujrepjzslg4q6xvp55p73nh8f793r6qxswpkn")
    Ursea         = mvx.MvxAddress("erd1gm70t49zh83fjdv48n49986fknrhrenkx47rn8dgw7202jvad3kq256c8e")
    BailWrite     = mvx.MvxAddress("erd1qf0n50w0k7xy5ydq4sx0zzfzucsn0w8pah7l3ppcpm2h9hv06wls324563")
    Rogojan       = mvx.MvxAddress("erd1gdp3cne23v6mrx9l3yj9k92tqqk5ycwmy0t0x6tz4hw3tmv3mgksn0pvrh")
    DemiFarm      = mvx.MvxAddress("erd1qrp72xhyn6vd5jrsvzkx32cecszuvh9yka0escx7j977cw8yk77qyup3zq")
    Xlauncher     = mvx.MvxAddress("erd1xa39h8q20gy25449vw2qt4dm38pp3nnxp7kzga2pt54z4u2rgjlqadlgdl")
    Dezvoltatorul = mvx.MvxAddress("erd1xt5ullndellcuxfjycpeds97lv5cwlekwn48l437v7f76lx0zfwq8ycjj0")
    EmmaB         = mvx.MvxAddress("erd1hfjt72u0gvs9p288yskvuerdyrszf48ugldared6myzcgvywpz0sp7f2rr")
    Up4Nothing    = mvx.MvxAddress("erd1mfw3j6fjfu6p36tatlaq5qxydh799xkw2egdq2lls3ve7qq8tvsqsltu7s")
    Daniel282     = mvx.MvxAddress("erd1nuyx3dv78e2ymzdd43z8aektcmyckq3qz4a4nhu3dh2mq5me7mlqkdcqt5")
    Uzkat         = mvx.MvxAddress("erd1w6vaq464stwtfses2yq6skaz5v5gz90vn22f2zm0mjt7aayya09sdtvu2r")
    VioTm         = mvx.MvxAddress("erd1t2732ajvwkzsym3vykmzh4d5fdu6dn4rqkh777x9kfmanpmcmdxqme0heq")
    Valentin      = mvx.MvxAddress("erd1z4v57r2cu6867pzlhcnazps37kx7rgwf8dn4falx86274ry6y7esnxgmy8")

    UserNameList = []string{"AncientHodler", "TrDaniel", "DRX", "Patryx",
	"Lavinia", "Sandu", "Cuciorva", "Codarcea", "Pulecs", "Laurentiu",
	"Frostedk9", "IonutDRD", "Buhaici", "TheKid", "RaultTM", "MakeAStep", "Paul", "Florian",
	"Coding", "Bloodshed", "Elanul", "Buguletu", "Carlig", "Ursea", "BailWrite", "Rogojan",
	"DemiFarm", "Xlauncher", "Dezvoltatorul", "EmmaB", "Up4Nothing", "Daniel282", "Uzkat", "VioTm",
	"Valentin"}

    //Users
    UserChain = []VestaHoldings{User000,
	User001, User002, User003, User004, User005, User006, User007, User008, User009, User010,
	User011, User012, User013, User014, User015, User016, User017, User018, User019, User020,
	User021, User022, User023, User024, User025, User026, User027, User028, User029, User030,
	User031, User032, User033, User034}

    User000 = VestaHoldings{AncientHodler, 35, 151, 592}
    User001 = VestaHoldings{TrDaniel, 16, 0, 50}
    User002 = VestaHoldings{DRX, 9, 3, 15}
    User003 = VestaHoldings{Patryx, 5, 6, 9}
    User004 = VestaHoldings{Lavinia, 1, 6, 6}
    User005 = VestaHoldings{Sandu, 4, 21, 2}
    User006 = VestaHoldings{Cuciorva, 1, 3, 60}
    User007 = VestaHoldings{Codarcea, 3, 3, 0}
    User008 = VestaHoldings{Pulecs, 7, 10, 5}
    User009 = VestaHoldings{Laurentiu, 11, 1, 0}
    User010 = VestaHoldings{Frostedk9, 18, 2, 0}
    User011 = VestaHoldings{IonutDRD, 2, 0, 9}
    User012 = VestaHoldings{Buhaici, 0, 0, 2}
    User013 = VestaHoldings{TheKid, 0, 0, 50}
    User014 = VestaHoldings{RaulTM, 0, 0, 9}
    User015 = VestaHoldings{MakeAStep, 6, 10, 31}

    //User016 = VestaHoldings{Paul, 178, 186, 172}
    //User017 = VestaHoldings{Florian, 0, 49, 697}
    //User018 = VestaHoldings{Coding, 0, 400, 0}
    //User026 = VestaHoldings{DemiFarm, 0, 200, 500}

    User016 = VestaHoldings{Paul, 0, 0, 0}
    User017 = VestaHoldings{Florian, 0, 0, 0}
    User018 = VestaHoldings{Coding, 0, 0, 0}
    User019 = VestaHoldings{Bloodshed, 0, 0, 0}
    User020 = VestaHoldings{Elanul, 1, 5, 56}
    User021 = VestaHoldings{Buguletu, 5, 0, 0}
    User022 = VestaHoldings{Carlig, 1, 0, 0}
    User023 = VestaHoldings{Ursea, 5, 0, 0}
    User024 = VestaHoldings{BailWrite, 34, 0, 0}
    User025 = VestaHoldings{Rogojan, 9, 2, 6}
    User026 = VestaHoldings{DemiFarm, 0, 0, 0}
    User027 = VestaHoldings{Xlauncher, 0, 0, 0}
    User028 = VestaHoldings{Dezvoltatorul, 0, 0, 0}
    User029 = VestaHoldings{EmmaB, 0, 0, 10}
    User030 = VestaHoldings{Up4Nothing, 16, 7, 28}
    User031 = VestaHoldings{Daniel282, 16, 6, 2}
    User032 = VestaHoldings{Uzkat, 11, 13, 2}
    User033 = VestaHoldings{VioTm, 0, 33, 16}
    User034 = VestaHoldings{Valentin, 0, 0, 3}

    //Liquidity From Users
    LiquidityUserChain = []LpHoldings{VLQUser000,
	VLQUser001, VLQUser002, VLQUser003, VLQUser004, VLQUser005, VLQUser006, VLQUser007, VLQUser008, VLQUser009, VLQUser010,
	VLQUser011, VLQUser012, VLQUser013, VLQUser014, VLQUser015, VLQUser016, VLQUser017, VLQUser018, VLQUser019, VLQUser020,
	VLQUser021, VLQUser022, VLQUser023, VLQUser024, VLQUser025, VLQUser026, VLQUser027, VLQUser028, VLQUser029, VLQUser030,
	VLQUser031, VLQUser032, VLQUser033, VLQUser034}
    LQDEmpty = VestaDexLPs{Zero, Zero, Zero, Zero, Zero, Zero}
    VltEmpty = VestaDexVault{Zero, Zero, Zero, Zero}

    VLQUser000      = LpHoldings{AncientHodler, LQDUser000, KosonLiqUser000, BloodLiqUser000, VaultLiqUser000, SnakeLiqUser000}
    LQDUser000      = VestaDexLPs{p.NFS("28215.889749851923648510"), Zero, Zero, p.NFS("0.3242"), Zero, Zero}
    KosonLiqUser000 = VestaDexLPs{p.NFS("45137.278081125541251284"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser000 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser000 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser000 = OuroLP{p.NFS("0"), p.NFS("590441.163898165389832435")}
    //
    VLQUser001      = LpHoldings{TrDaniel, LQDUser001, KosonLiqUser001, BloodLiqUser001, VaultLiqUser001, SnakeLiqUser001}
    LQDUser001      = VestaDexLPs{p.NFS("1835.570798787343702618"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser001 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser001 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser001 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser001 = OuroLP{p.NFS("0"), p.NFS("16888.880734495260617049")}
    //
    VLQUser002      = LpHoldings{DRX, LQDUser002, KosonLiqUser002, BloodLiqUser002, VaultLiqUser002, SnakeLiqUser002}
    LQDUser002      = VestaDexLPs{p.NFS("93.27259156861914193"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser002 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser002 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser002 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser002 = OuroLP{p.NFS("0"), p.NFS("5998.181028398943605902")}
    //
    VLQUser003      = LpHoldings{Patryx, LQDUser003, KosonLiqUser003, BloodLiqUser003, VaultLiqUser003, SnakeLiqUser003}
    LQDUser003      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser003 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser003 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser003 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser003 = OuroLP{p.NFS("0"), p.NFS("4061.891505389062456945")}
    //
    VLQUser004      = LpHoldings{Lavinia, LQDUser004, KosonLiqUser004, BloodLiqUser004, VaultLiqUser004, SnakeLiqUser004}
    LQDUser004      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser004 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser004 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser004 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser004 = OuroLP{p.NFS("0"), p.NFS("0")}
    //
    VLQUser005      = LpHoldings{Sandu, LQDUser005, KosonLiqUser005, BloodLiqUser005, VaultLiqUser005, SnakeLiqUser005}
    LQDUser005      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser005 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser005 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser005 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser005 = OuroLP{p.NFS("0"), p.NFS("0")}
    //
    VLQUser006      = LpHoldings{Cuciorva, LQDUser006, KosonLiqUser006, BloodLiqUser006, VaultLiqUser006, SnakeLiqUser006}
    LQDUser006      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser006 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser006 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser006 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser006 = OuroLP{p.NFS("0"), p.NFS("5959.925753976562407994")}
    //
    VLQUser007      = LpHoldings{Codarcea, LQDUser007, KosonLiqUser007, BloodLiqUser007, VaultLiqUser007, SnakeLiqUser007}
    LQDUser007      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser007 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser007 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser007 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser007 = OuroLP{p.NFS("0"), p.NFS("26400")}
    //
    VLQUser008      = LpHoldings{Pulecs, LQDUser008, KosonLiqUser008, BloodLiqUser008, VaultLiqUser008, SnakeLiqUser008}
    LQDUser008      = VestaDexLPs{p.NFS("4380.479572405876167673"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser008 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser008 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser008 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("1613000"), p.NFS("0")}
    SnakeLiqUser008 = OuroLP{p.NFS("0"), p.NFS("1927.084307414062450010")}
    //
    VLQUser009      = LpHoldings{Laurentiu, LQDUser009, KosonLiqUser009, BloodLiqUser009, VaultLiqUser009, SnakeLiqUser009}
    LQDUser009      = VestaDexLPs{p.NFS("1077.940543973583389333"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser009 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser009 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser009 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser009 = OuroLP{p.NFS("0"), p.NFS("6527.181274794853601651")}
    //
    VLQUser010      = LpHoldings{Frostedk9, LQDUser010, KosonLiqUser010, BloodLiqUser010, VaultLiqUser010, SnakeLiqUser010}
    LQDUser010      = VestaDexLPs{p.NFS("1058.492448497819050567"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser010 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser010 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser010 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser010 = OuroLP{p.NFS("0"), p.NFS("9435.254344918618067014")}
    //
    VLQUser011      = LpHoldings{IonutDRD, LQDUser011, KosonLiqUser011, BloodLiqUser011, VaultLiqUser011, SnakeLiqUser011}
    LQDUser011      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser011 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser011 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser011 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser011 = OuroLP{p.NFS("0"), p.NFS("0")}
    //
    VLQUser012      = LpHoldings{Buhaici, LQDUser012, KosonLiqUser012, BloodLiqUser012, VaultLiqUser012, SnakeLiqUser012}
    LQDUser012      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser012 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser012 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser012 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser012 = OuroLP{p.NFS("0"), p.NFS("0")}
    //
    VLQUser013      = LpHoldings{TheKid, LQDUser013, KosonLiqUser013, BloodLiqUser013, VaultLiqUser013, SnakeLiqUser013}
    LQDUser013      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser013 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser013 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser013 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser013 = OuroLP{p.NFS("0"), p.NFS("5560.897808499187423050")}
    //
    VLQUser014      = LpHoldings{RaulTM, LQDUser014, KosonLiqUser014, BloodLiqUser014, VaultLiqUser014, SnakeLiqUser014}
    LQDUser014      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser014 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser014 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser014 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser014 = OuroLP{p.NFS("0"), p.NFS("0")}
    //
    VLQUser015      = LpHoldings{MakeAStep, LQDUser015, KosonLiqUser015, BloodLiqUser015, VaultLiqUser015, SnakeLiqUser015}
    LQDUser015      = VestaDexLPs{p.NFS("2650.217342505712876198"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser015 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser015 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser015 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser015 = OuroLP{p.NFS("0"), p.NFS("6595.151966627011704552")}
    //
    VLQUser016      = LpHoldings{Paul, LQDUser016, KosonLiqUser016, BloodLiqUser016, VaultLiqUser016, SnakeLiqUser016}
    LQDUser016      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser016 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser016 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser016 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser016 = OuroLP{p.NFS("0"), p.NFS("252000")}
    //
    VLQUser017      = LpHoldings{Florian, LQDUser017, KosonLiqUser017, BloodLiqUser017, VaultLiqUser017, SnakeLiqUser017}
    LQDUser017      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser017 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser017 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser017 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser017 = OuroLP{p.NFS("0"), p.NFS("0")}
    //
    VLQUser018      = LpHoldings{Coding, LQDUser018, KosonLiqUser018, BloodLiqUser018, VaultLiqUser018, SnakeLiqUser018}
    LQDUser018      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser018 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser018 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser018 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser018 = OuroLP{p.NFS("0"), p.NFS("0")}
    //
    VLQUser019      = LpHoldings{Bloodshed, LQDUser019, KosonLiqUser019, BloodLiqUser019, VaultLiqUser019, SnakeLiqUser019}
    LQDUser019      = VestaDexLPs{p.NFS("24452.663808508794693464"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser019 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser019 = VestaDexLPs{p.NFS("193219.526790820208733109"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser019 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser019 = OuroLP{p.NFS("0"), p.NFS("0")}
    //
    VLQUser020      = LpHoldings{Elanul, LQDUser020, KosonLiqUser020, BloodLiqUser020, VaultLiqUser020, SnakeLiqUser020}
    LQDUser020      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser020 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser020 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser020 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser020 = OuroLP{p.NFS("0"), p.NFS("0")}
    //
    VLQUser021      = LpHoldings{Buguletu, LQDUser021, KosonLiqUser021, BloodLiqUser021, VaultLiqUser021, SnakeLiqUser021}
    LQDUser021      = VestaDexLPs{p.NFS("429.539899170853140842"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser021 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser021 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser021 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser021 = OuroLP{p.NFS("0"), p.NFS("2528.602062767862839725")}
    //
    VLQUser022      = LpHoldings{Carlig, LQDUser022, KosonLiqUser022, BloodLiqUser022, VaultLiqUser022, SnakeLiqUser022}
    LQDUser022      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser022 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser022 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser022 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser022 = OuroLP{p.NFS("0"), p.NFS("0")}
    //
    VLQUser023      = LpHoldings{Ursea, LQDUser023, KosonLiqUser023, BloodLiqUser023, VaultLiqUser023, SnakeLiqUser023}
    LQDUser023      = VestaDexLPs{p.NFS("215.236670944110348465"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser023 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser023 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser023 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser023 = OuroLP{p.NFS("0"), p.NFS("0")}
    //
    VLQUser024      = LpHoldings{BailWrite, LQDUser024, KosonLiqUser024, BloodLiqUser024, VaultLiqUser024, SnakeLiqUser024}
    LQDUser024      = VestaDexLPs{p.NFS("330.735566649713741425"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser024 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser024 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser024 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser024 = OuroLP{p.NFS("0"), p.NFS("14737.270955287499875415")}
    //
    VLQUser025      = LpHoldings{Rogojan, LQDUser025, KosonLiqUser025, BloodLiqUser025, VaultLiqUser025, SnakeLiqUser025}
    LQDUser025      = VestaDexLPs{p.NFS("2766.023226046414456636"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser025 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser025 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser025 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser025 = OuroLP{p.NFS("0"), p.NFS("0")}
    //
    VLQUser026      = LpHoldings{DemiFarm, LQDUser026, KosonLiqUser026, BloodLiqUser026, VaultLiqUser026, SnakeLiqUser026}
    LQDUser026      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser026 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser026 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser026 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser026 = OuroLP{p.NFS("0"), p.NFS("0")}
    //
    VLQUser027      = LpHoldings{Xlauncher, LQDUser027, KosonLiqUser027, BloodLiqUser027, VaultLiqUser027, SnakeLiqUser027}
    LQDUser027      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser027 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser027 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser027 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser027 = OuroLP{p.NFS("0"), p.NFS("363000")}
    //
    VLQUser028      = LpHoldings{Dezvoltatorul, LQDUser028, KosonLiqUser028, BloodLiqUser028, VaultLiqUser028, SnakeLiqUser028}
    LQDUser028      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser028 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser028 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser028 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser028 = OuroLP{p.NFS("0"), p.NFS("600000")}
    //
    VLQUser029      = LpHoldings{EmmaB, LQDUser029, KosonLiqUser029, BloodLiqUser029, VaultLiqUser029, SnakeLiqUser029}
    LQDUser029      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser029 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser029 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser029 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser029 = OuroLP{p.NFS("0"), p.NFS("0")}
    //
    VLQUser030      = LpHoldings{Up4Nothing, LQDUser030, KosonLiqUser030, BloodLiqUser030, VaultLiqUser030, SnakeLiqUser030}
    LQDUser030      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser030 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser030 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser030 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser030 = OuroLP{p.NFS("0"), p.NFS("0")}
    //
    VLQUser031      = LpHoldings{Daniel282, LQDUser031, KosonLiqUser031, BloodLiqUser031, VaultLiqUser031, SnakeLiqUser031}
    LQDUser031      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser031 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser031 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser031 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser031 = OuroLP{p.NFS("0"), p.NFS("0")}
    //
    VLQUser032      = LpHoldings{Uzkat, LQDUser032, KosonLiqUser032, BloodLiqUser032, VaultLiqUser032, SnakeLiqUser032}
    LQDUser032      = VestaDexLPs{p.NFS("310.836641022810311016"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser032 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser032 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser032 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser032 = OuroLP{p.NFS("0"), p.NFS("0")}
    //
    VLQUser033      = LpHoldings{VioTm, LQDUser033, KosonLiqUser033, BloodLiqUser033, VaultLiqUser033, SnakeLiqUser033}
    LQDUser033      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser033 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser033 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser033 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser033 = OuroLP{p.NFS("0"), p.NFS("0")}
    //
    VLQUser034      = LpHoldings{Valentin, LQDUser034, KosonLiqUser034, BloodLiqUser034, VaultLiqUser034, SnakeLiqUser034}
    LQDUser034      = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    KosonLiqUser034 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    BloodLiqUser034 = VestaDexLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    VaultLiqUser034 = VestaDexVault{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
    SnakeLiqUser034 = OuroLP{p.NFS("0"), p.NFS("0")}
)

// Individual Multiplier Computation
func ComputeTotalIM(Input []VestaHoldings) *p.Decimal {
    var (
	TotalIM = p.NFS("0")
    )
    for i := 0; i < len(Input); i++ {
	CycleIM := GetSingleIM(Input[i])
	TotalIM = sm.ADDxc(TotalIM, CycleIM)
    }
    return TotalIM
}

func ComputeFinalTotalIM(Input []VestaHoldings) *p.Decimal {
    return sm.ADDxc(sm.DIVxc(ComputeTotalIM(Input), p.NFS("100")), p.NFS("1"))
}

func ComputeFinalTotalIMVault(Input *p.Decimal) *p.Decimal {
    V1 := sm.SUBxc(Input, p.NFS("1"))
    V2 := sm.MULxc(V1, p.NFS("5"))
    V3 := sm.ADDxc(V2, p.NFS("1"))
    return V3
}

func GetSingleIM(Input VestaHoldings) *p.Decimal {
    V1 := sm.MULxc(p.NFI(Input.Gold), p.NFS("5"))
    V2 := sm.MULxc(p.NFI(Input.Silver), p.NFS("2.5"))
    V3 := sm.MULxc(p.NFI(Input.Bronze), p.NFS("1.25"))
    return sm.SUMxc(V1, V2, V3)
}

// Computes Guest IM, and Percentual Split Chain given a Guest Position
func ComputeVestaSplit(Position int64, Input []VestaHoldings) (GuestIM *p.Decimal, PercentualSplitChain []*p.Decimal) {

    var (
	UnitPercent = new(p.Decimal)
    )
    PercentualSplitChain = make([]*p.Decimal, len(Input))

    GIM := GetSingleIM(Input[Position])
    GuestIM = sm.ADDxc(sm.DIVxc(GIM, p.NFS("100")), p.NFS("1"))
    TIM := ComputeTotalIM(Input)
    TIM2 := sm.SUBxc(TIM, GIM)

    for i := 0; i < len(Input); i++ {
	if int64(i) == Position {
	    PercentualSplitChain[i] = p.NFS("1")
	} else {
	    CurrentIM := GetSingleIM(Input[i])
	    UnitPercent = sm.TruncateCustom(sm.DIVxc(CurrentIM, TIM2), 18)
	    PercentualSplitChain[i] = UnitPercent
	}
    }
    return GuestIM, PercentualSplitChain
}

// Outputs Farm Multiplier Structure
func OutputVVMx(Variant string, InputIM, InputUM *p.Decimal) (Output FarmMx) {
    if Variant == "vesta" {
	Output = FarmMx{AncientDEB, Neutral, VestaTMSix, Neutral, InputUM, InputIM, Neutral, Neutral}
    } else if Variant == "blood" {
	Output = FarmMx{BloodshedDEB, Neutral, Neutral, Platinum, InputUM, InputIM, Neutral, Neutral}
    } else if Variant == "koson" {
	Output = FarmMx{AncientDEB, Neutral, Neutral, Platinum, InputUM, InputIM, Neutral, Neutral}
    } else if Variant == "pulecsvault" {
	Output = FarmMx{Neutral, Neutral, Neutral, Neutral, Neutral, ComputeFinalTotalIMVault(InputIM), Neutral, Neutral}
    }

    return Output
}

// Computes Total Multiplier of a farm Structure
func GetAllMx(Input FarmMx) *p.Decimal {
    return sm.PRDxc(Input.DEB, Input.VLM, Input.TM, Input.D, Input.UM, Input.IM, Input.DM, Input.VM)
}

// Creates a split of Vesta Values, given an amount, an chain of percents.
func MakeVestaSplit(Amount *p.Decimal, VS []*p.Decimal) []*p.Decimal {
    var (
	OutputVestaChain = make([]*p.Decimal, len(VS))
	UnitVesta        = new(p.Decimal)
    )
    Half := sm.TruncateCustom(sm.DIVxc(Amount, p.NFS("2")), 18)
    for i := 0; i < len(VS); i++ {
	UnitVesta = sm.TruncateCustom(sm.MULxc(Half, VS[i]), 18)
	OutputVestaChain[i] = UnitVesta
    }
    return OutputVestaChain
}

// Given an InputUM, a Guest position, and a RawVesta Amount to be split, computes a Chain of Vesta Token Values
// According to the List of SFT Holders.
func AbsolutSplitWithVesta(Variant string, RWAmount, InputUM *p.Decimal, Position int64, Input []VestaHoldings) (Remaining *p.Decimal, VestaAmountChain []*p.Decimal) {
    TotalIM := ComputeFinalTotalIM(Input)

    GIM, VS := ComputeVestaSplit(Position, Input)

    ALLVestaFarmMultipliers := OutputVVMx(Variant, TotalIM, InputUM)
    //fmt.Println("M1: ", ALLVestaFarmMultipliers)
    MyVestaFarmMultipliers := OutputVVMx(Variant, GIM, InputUM)
    //fmt.Println("M2: ", MyVestaFarmMultipliers)

    PersonalAmount := sm.TruncateCustom(sm.MULxc(RWAmount, GetAllMx(MyVestaFarmMultipliers)), 18)
    //fmt.Println("PersonalAmount is, ", PersonalAmount)
    BoostedAmount := sm.TruncateCustom(sm.MULxc(RWAmount, GetAllMx(ALLVestaFarmMultipliers)), 18)
    GainedAmount := sm.SUBxc(BoostedAmount, PersonalAmount)
    //fmt.Println("PA: ", PersonalAmount)
    //fmt.Println("BA: ", BoostedAmount)

    return PersonalAmount, MakeVestaSplit(GainedAmount, VS)
}

//Compute Raw Splits based on Individual Liquidity.

// Computes the VLP for an individual, used later for the VLP Split given individual user Liquidity
func ComputeUserVLP(Input VestaDexLPs) *p.Decimal {
    var (
	Output = new(p.Decimal)
    )
    if Input == LQDEmpty {
	Output = p.NFS("0")
    } else {
	V1 := sm.TruncateCustom(sm.MULxc(Input.Gold, p.NFS("2.5")), 18)
	V2 := sm.TruncateCustom(sm.MULxc(Input.Silver, p.NFS("1.6")), 18)
	V3 := sm.TruncateCustom(sm.MULxc(Input.Bronze, p.NFS("1")), 18)
	V4 := sm.TruncateCustom(sm.MULxc(Input.UGold, p.NFS("1.75")), 18)
	V5 := sm.TruncateCustom(sm.MULxc(Input.USilver, p.NFS("0.96")), 18)
	V6 := sm.TruncateCustom(sm.MULxc(Input.UBronze, p.NFS("0.5")), 18)
	Output = sm.SUMxc(V1, V2, V3, V4, V5, V6)
    }

    return Output
}

func ComputeUserVaultLiq(Input VestaDexVault) *p.Decimal {
    var (
	Output = new(p.Decimal)
    )
    if Input == VltEmpty {
	Output = p.NFS("0")
    } else {
	Output = sm.SUMxc(Input.Vesta, Input.Blessed, Input.Frozen, Input.Sleeping)
    }

    return Output
}

func ComputeUserTokenVLP(Variant string, Input LpHoldings) *p.Decimal {
    var (
	Output = new(p.Decimal)
    )

    if Variant == "vesta" {
	Output = ComputeUserVLP(Input.VestaLiq)
    } else if Variant == "koson" {
	Output = ComputeUserVLP(Input.KosonLiq)
    } else if Variant == "blood" {
	Output = ComputeUserVLP(Input.BloodLiq)
    } else if Variant == "pulecsvault" {
	Output = ComputeUserVaultLiq(Input.Vault)
    }

    return Output
}

// Computes total VLP considering the Liquidity Holdings of all Participants
func ComputeTotalTokenVLP(Variant string, Input []LpHoldings) *p.Decimal {
    var (
	VLPSum = p.NFS("0")
	Unit   = new(p.Decimal)
    )
    for i := 0; i < len(Input); i++ {
	Unit = ComputeUserTokenVLP(Variant, Input[i])
	VLPSum = sm.ADDxc(VLPSum, Unit)
    }
    return VLPSum
}

// Computes the VLP Split given the Liquidity Pooled by all participants.
// Used to compute the RawVesta Amount each individual user would earn
// This is further used when computing individual user yield based on individual liquidity
func ComputeVLPSplit(Variant string, Input []LpHoldings) (*p.Decimal, []*p.Decimal) {
    var (
	VLPSplitChain = make([]*p.Decimal, len(Input))
	UnitVLPSplit  = new(p.Decimal)
    )
    GroupVLP := ComputeTotalTokenVLP(Variant, Input)

    for i := 0; i < len(Input); i++ {
	VLP := ComputeUserTokenVLP(Variant, Input[i])
	if sm.DecimalEqual(VLP, p.NFS("0")) == true {
	    UnitVLPSplit = p.NFS("0")
	} else {
	    UnitVLPSplit = sm.TruncateCustom(sm.DIVxc(VLP, GroupVLP), 18)
	}
	VLPSplitChain[i] = UnitVLPSplit
    }
    return GroupVLP, VLPSplitChain

}

// Variadic Vesta Chain Adder
// Seems i didnt need it
func TotalVestaChainAdder(First []*p.Decimal, Rest ...[]*p.Decimal) []*p.Decimal {
    MakeZeroSlice := func(length int, Item *p.Decimal) []*p.Decimal {
	OutputSlice := make([]*p.Decimal, length)
	for i := range OutputSlice {
	    OutputSlice[i] = Item
	}
	return OutputSlice
    }
    sum := MakeZeroSlice(len(First), p.NFS("0"))
    restsum := MakeZeroSlice(len(First), p.NFS("0"))

    for _, item := range Rest {
	restsum = VestaChainAdder(restsum, item)
    }
    sum = VestaChainAdder(First, restsum)
    return sum
}

// Function for adding two slices of decimals
func VestaChainAdder(C1, C2 []*p.Decimal) []*p.Decimal {
    var (
	OutputChain = make([]*p.Decimal, len(C1))
	Unit        = new(p.Decimal)
    )
    for i := 0; i < len(C1); i++ {
	Unit = sm.TruncateCustom(sm.ADDxc(C1[i], C2[i]), 18)
	OutputChain[i] = Unit
    }
    return OutputChain
}

// Computing Individual Minting Amounts
// Multiplies a decimal with each decimal in a chain of decimals
func CreateRawVestaSplit(RawAmount *p.Decimal, VLPSplit []*p.Decimal) []*p.Decimal {
    var (
	OutputChain = make([]*p.Decimal, len(VLPSplit))
	Unit        = new(p.Decimal)
    )
    for i := 0; i < len(VLPSplit); i++ {
	Unit = sm.TruncateCustom(sm.MULxc(RawAmount, VLPSplit[i]), 18)
	OutputChain[i] = Unit
    }
    return OutputChain
}

// The Final Function that computes individual Vesta yields considering all participants guests.
// Then adds the individual computed Vesta Yields Together
func MultipleAbsoluteSplitWithVesta(Variant string, RawVestaAmount, InputUM *p.Decimal, VestaSFTsChain []VestaHoldings, LPChain []LpHoldings) (TotalVLP *p.Decimal, VLPSplit []*p.Decimal, AncientAmount *p.Decimal, TotalVestaRewardChain []*p.Decimal) {
    TotalVLP, VLPSplit = ComputeVLPSplit(Variant, LPChain) //VLP Split
    RawVestaSplit := CreateRawVestaSplit(RawVestaAmount, VLPSplit)

    MakeZeroSlice := func(length int, Item *p.Decimal) []*p.Decimal {
	OutputSlice := make([]*p.Decimal, length)
	for i := range OutputSlice {
	    OutputSlice[i] = Item
	}
	return OutputSlice
    }
    SummedChain := MakeZeroSlice(len(VestaSFTsChain), p.NFS("0"))

    var (
	VestaRewardForPosition = make([]*p.Decimal, len(VestaSFTsChain))
	PersonalAmountsChain   = make([]*p.Decimal, len(VestaSFTsChain))
	PA                     = new(p.Decimal)
    )

    for i := 0; i < len(VestaSFTsChain); i++ {
	//fmt.Println("**************")
	PA, VestaRewardForPosition = AbsolutSplitWithVesta(Variant, RawVestaSplit[i], InputUM, int64(i), VestaSFTsChain)
	//fmt.Println("Chain on position ", i, " is ", VestaRewardForPosition)
	//fmt.Println("**************")
	SummedChain = VestaChainAdder(SummedChain, VestaRewardForPosition)
	if i == 0 {
	    AncientAmount = PA
	    PersonalAmountsChain[i] = p.NFS("0")
	} else {
	    PersonalAmountsChain[i] = PA
	}
    }
    //Personal Amount Chain is the chain with personal Vesta Amounts owner of sent liquidity is earning
    //with his liquidity and with his SFTs. This amount is his in entirety, and he only splits the extra
    //Vesta Token amounts generated with his boosters.
    TotalVestaRewardChain = VestaChainAdder(SummedChain, PersonalAmountsChain)
    return TotalVLP, VLPSplit, AncientAmount, TotalVestaRewardChain
}

func SumChain(InputChain []*p.Decimal) *p.Decimal {
    var (
	SUM = p.NFS("0")
    )
    for i := 0; i < len(InputChain); i++ {
	SUM = sm.ADDxc(SUM, InputChain[i])
    }
    return SUM
}

func ComputeMintPercent(GuestPosition int, PersonalAmount *p.Decimal, InputChain []*p.Decimal) *p.Decimal {
    ChainSum := SumChain(InputChain)
    TotalVST := sm.ADDxc(ChainSum, PersonalAmount)
    fmt.Println("CMP: Total VST is: ", TotalVST)
    TotalGuestAmount := sm.ADDxc(PersonalAmount, InputChain[GuestPosition])
    fmt.Println("CMP: Total Guest Amount is: ", TotalGuestAmount)
    OutgoingAmount := sm.SUBxc(ChainSum, InputChain[GuestPosition])
    fmt.Println("CMP: Total Outgoing Amount to be sent is: ", OutgoingAmount)

    OutgoingAmountPercent := sm.TruncateCustom(sm.DIVxc(OutgoingAmount, TotalVST), 18)
    fmt.Println("CMP: Decimalic Outgoing Percent := ", OutgoingAmountPercent)
    OutgoingAmountRoundUP := sm.TruncateCustom(sm.MULxc(OutgoingAmountPercent, p.NFS("100")), 0)
    FinalOutgoingAmount := sm.ADDxc(OutgoingAmountRoundUP, p.NFS("1"))
    fmt.Println("CMP: Integer Outgoing Percent for MINT =", FinalOutgoingAmount)
    fmt.Println("=====================================")
    return OutgoingAmountRoundUP
}

func ExportOutgoingVestas(GuestPosition int, MainChain []VestaHoldings, Rewards []*p.Decimal) []mvx.BalanceESDT {
    var (
	OutputChain = make([]mvx.BalanceESDT, len(MainChain))
    )
    for i := 0; i < len(MainChain); i++ {
	OutputChain[i].Address = MainChain[i].Address
	OutputChain[i].Balance = sm.DTS(Rewards[i])
    }
    FinalOutput := append(OutputChain[:GuestPosition], OutputChain[GuestPosition+1:]...)
    //Extracts the Guest Position
    mvx.ConvertToBulkCSV("ExportVesta.csv", FinalOutput)
    //mvx.ConvertToBulkCSV("ExportVesta.csv", OutputChain[1:])
    //return OutputChain[1:]
    return FinalOutput
}

func ExportGroupData(OutputName string, NameList []string, VestaSFTsChain []VestaHoldings, LPChain []LpHoldings) {
    f, err := os.Create(OutputName)
    if err != nil {
	fmt.Println(err)
	_ = f.Close()
	return
    }

    LineToPrint := func(Info0 string, Info1 VestaHoldings, Info2 LpHoldings) string {
	ERD := string(Info1.Address)
	GoldSFT := strconv.Itoa(int(Info1.Gold))
	SilverSFT := strconv.Itoa(int(Info1.Silver))
	BronzeSFT := strconv.Itoa(int(Info1.Bronze))
	GoldLiq := sm.DTS(Info2.VestaLiq.Gold)
	SilverLiq := sm.DTS(Info2.VestaLiq.Silver)
	BronzeLiq := sm.DTS(Info2.VestaLiq.Bronze)
	UGoldLiq := sm.DTS(Info2.VestaLiq.UGold)
	USilverLiq := sm.DTS(Info2.VestaLiq.USilver)
	UBronzeLiq := sm.DTS(Info2.VestaLiq.UBronze)

	L := Info0 + ";" + ERD + ";" + GoldSFT + ";" + SilverSFT + ";" + BronzeSFT + ";" + GoldLiq + ";" + SilverLiq + ";" + BronzeLiq + ";" + UGoldLiq + ";" + USilverLiq + ";" + UBronzeLiq
	return L
    }

    for i := 0; i < len(VestaSFTsChain); i++ {
	_, err1 := fmt.Fprintln(f, LineToPrint(NameList[i], VestaSFTsChain[i], LPChain[i]))
	if err1 != nil {
	    return
	}
    }
}

//(VestaSFTsChain []VestaHoldings, LPChain []LpHoldings)

func ImportGroupData(OutputName string) ([]VestaHoldings, []LpHoldings) {

    StringSlice := mvx.ReadFile(OutputName)
    var (
	Chain1 = make([]VestaHoldings, len(StringSlice))
	Chain2 = make([]LpHoldings, len(StringSlice))
    )

    StrToInt := func(Input string) int64 {
	Output, _ := strconv.Atoi(Input)
	return int64(Output)
    }

    for i := 0; i < len(StringSlice); i++ {
	SeparatedStrings := strings.Split(StringSlice[i], ";")
	Chain1[i].Address = mvx.MvxAddress(SeparatedStrings[1])
	Chain1[i].Gold = StrToInt(SeparatedStrings[2])
	Chain1[i].Silver = StrToInt(SeparatedStrings[3])
	Chain1[i].Bronze = StrToInt(SeparatedStrings[4])
	Chain2[i].Address = mvx.MvxAddress(SeparatedStrings[1])
	Chain2[i].VestaLiq.Gold = p.NFS(SeparatedStrings[5])
	Chain2[i].VestaLiq.Silver = p.NFS(SeparatedStrings[6])
	Chain2[i].VestaLiq.Bronze = p.NFS(SeparatedStrings[7])
	Chain2[i].VestaLiq.UGold = p.NFS(SeparatedStrings[8])
	Chain2[i].VestaLiq.USilver = p.NFS(SeparatedStrings[9])
	Chain2[i].VestaLiq.UBronze = p.NFS(SeparatedStrings[10])
    }

    return Chain1, Chain2
}
