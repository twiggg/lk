package validators

import (
	"sort"
	"strings"
)

var validGdaxPairs = map[string]string{
	"btc-usd": "bitcoin", "eth-usd": "ethereum", "eth-btc": "ethereum", "ltc-usd": "litecoin", "ltc-btc": "litecoin",
	"btc-eur": "bitcoin", "ltc-eur": "litecoin", "eth-eur": "ethereum",
	"btc-gbp": "bitcoin",
}

var validBittrexPairs = map[string]string{
	//btc
	"ada-btc": "ada", "eth-btc": "ethereum", "srn-btc": "sirin token", "trx-btc": "tron", "xvg-btc": "verge", "rdd-btc": "reddcoin", "cvc-btc": "civic", "xrp-btc": "ripple", "zrx-btc": "Ox protocol", "strat-btc": "stratis", "bch-btc": "bitcoin cash", "dcr-btc": "decred", "steem-btc": "steem", "xlm-btc": "lumen", "poly-btc": "polymath", "sc-btc": "siacoin", "zen-btc": "zencash", "aid-btc": "aidcoin", "ltc-btc": "litecoin", "neo-btc": "neo",
	"zec-btc": "zcash", "doge-btc": "dogecoin", "waves-btc": "waves", "part-btc": "particl", "zcl-btc": "zclassic", "gnt-btc": "golem", "dgb-btc": "digibyte", "gto-btc": "gifto", "dct-btc": "decent", "qtum-btc": "qtum", "grs-btc": "groestlcoin", "tusd-btc": "trueusd", "xem-btc": "nem", "omg-btc": "omisego", "dash-btc": "dash", "rep-btc": "augur", "lsk-btc": "lisk", "wax-btc": "worlwide asset exchange", "etc-btc": "ethereum classic", "xmr-btc": "monero",
	/*"rvr-btc": "revolution vr", "vbr-btc": "viberate", "rlc-btc": "iex.ec", "pay-btc": "tenx pay token", "vtc-btc": "vertcoin", "nxs-btc": "nexus", "ardr-btc": "ardor", "powr-btc": "powerledger", "storm-btc": "storm", "nxt-btc": "nxt", "ngc-btc": "naga", "sys-btc": "syscoin", "snt-btc": "status network token", "kmd-btc": "komodo", "btg-btc": "bitcoin-gold", "bat-btc": "basic attention token", "ark-btc": "ark", "salt-btc": "salt", "ion-btc": "ion", "bcpt-btc": "blockmason credit protocol",
	"nmr-btc": "numeraire", "pivx-btc": "pivx", "vee-btc": "blockv", "mana-btc": "decentraland", "burst-btc": "burstcoin", "edg-btc": "edgeless", "adx-btc": "adex", "eng-btc": "enigma", "block-btc": "blocknet", "tx-btc": "transfercoin", "xcp-btc": "counterparty", "gbyte-btc": "bytes", "xwc-btc": "whitecoin", "xdn-btc": "digitalnote", "bnt-btc": "bancor", "lun-btc": "lunyr",
	"emc2-btc": "einsteinium", "xzc-btc": "zcoin", "mco-btc": "monaco", "fct-btc": "factom", "ubq-btc": "ubiq", "cfi-btc": "cofound.it", "xel-btc": "elastic", "wings-btc": "wings dao", "ebst-btc": "eboost", "storj-btc": "storj", "trst-btc": "trustcoin", "ok-btc": "okcash", "aeon-btc": "aeon", "ant-btc": "aragon", "gno-btc": "gnosis", "nav-btc": "navcoin", "qrl-btc": "quantum resistant ledger", "lrc-btc": "loopring",
	*/
	//eth
	"trx-eth": "tron", "ada-eth": "ada", "bch-eth": "bitcoin cash", "zrx-eth": "Ox protocol", "xlm-eth": "lumen", "xrp-eth": "ripple", "ltc-eth": "litecoin", "vib-eth": "viberate", "poly-eth": "polymath", "xmr-eth": "monero", "pay-eth": "tenx pay token", "neo-eth": "neo", "cvc-eth": "civic", "tusd-eth": "trueusd", "omg-eth": "omisego", "dash-eth": "dash", "zec-eth": "zcash", "dgb-eth": "digibyte", "etc-eth": "ethereum classic",
	"qtum-eth": "qtum", "strat-eth": "stratis" /*"vee-eth": "blockv",*/, "xem-eth": "nem", "storm-eth": "storm", "sc-eth": "siacoin", "gnt-eth": "golem", "wax-eth": "worldwide asset exchange", "rep-eth": "augur", "srn-eth": "sirin token", "snt-eth": "status network token", "mco-eth": "monaco", "rlc-eth": "iex.ec", "bat-eth": "basic attention token", "waves-eth": "waves" /*"btg-eth": "bitcoin gold",*/, "powr-eth": "powerledger", "ant-eth": "aragon",
	/*"blt-eth": "bloom", "salt-eth": "salt", "gno-eth": "gnosis", "tix-eth": "blocktix", "gto-eth": "gifto", "adt-eth": "adtoken", "wings-eth": "wings dao", "fct-eth": "factom", "dmt-eth": "dmarket", "eng-eth": "enigma", "gup-eth": "guppy", "dnt-eth": "districtOx", "cfi-eth": "cofound.it", "lun-eth": "lunyr", "pro-eth": "propy", "ukg-eth": "unikoingold", "trst-eth": "trustcoin", "mana-eth": "decentraland", "storj-eth": "storj", "bnt-eth": "bancor", "ngc-eth": "naga", "qrl-eth": "quantum resistant ledger",
	"adx-eth": "adex", "lrc-eth": "loopring", "ptoy-eth": "patientory", "lgd-eth": "legends", "hmq-eth": "humaniq",
	*/
	//usdt
	"btc-usdt": "bitcoin", "eth-usdt": "ethereum", "ada-usdt": "ada", "bch-usdt": "bitcoin cash", "xrp-usdt": "ripple", "trx-usdt": "tron", "neo-usdt": "neo", "ltc-usdt": "litecoin", "zec-usdt": "zcash", "tusd-usdt": "trueusd", "xvg-usdt": "verge", "omg-usdt": "omisego", "etc-usdt": "ethereum classic", "xmr-usdt": "monero" /* "btg-usdt": "bitcoin gold",*/, "sc-usdt": "siacoin", "dash-usdt": "dash", "nxt-usdt": "nxt", "dcr-usdt": "decred",
}
var validBinancePairs = map[string]string{
	//btc
	"eos-btc": "eos", "eth-btc": "ethereum", "trx-btc": "tron", "zrx-btc": "0x", "bnb-btc": "binance coin", "xrp-btc": "ripple", "ppt-btc": "populous", "ada-btc": "ada", "ont-btc": "ontology", "neo-btc": "neo" /*"bcc-btc": "bitcoin cash",*/, "icx-btc": "icon", "ltc-btc": "litecoin", "gto-btc": "gifto", "arn-btc": "aeron", "xlm-btc": "lumen", "mtl-btc": "metal", "lun-btc": "lunyr", "xvg-btc": "verge", "elf-btc": "aelf", "iota-btc": "iota", "sky-btc": "skycoin", "zil-btc": "zilliqa",
	"ven-btc": "vechain", "wan-btc": "wancoin", "iost-btc": "iostoken" /*"bcn-btc": "bytecoin",*/, "cmt-btc": "cybermiles", "tusd-btc": "trueusd", "waves-btc": "waves", "nav-btc": "navcoin", "nano-btc": "nano", "dash-btc": "dash", "strat-btc": "stratis", "gvt-btc": "genesis vision", "ncash-btc": "nucleus vision", "snt-btc": "status network token", "omg-btc": "omisego", "brd-btc": "bread", "zec-btc": "zcash", "xmr-btc": "monero", "edo-btc": "eidoo", "mco-btc": "monaco", "steem-btc": "steem",
	/*"sub-btc": "substratum", "storm-btc": "storm", "qlc-btc": "qlink", "bqx-btc": "ethos", "eng-btc": "enigma", "ae-btc": "aeternity", "wings-btc": "wings", "gnt-btc": "golem", "qtum-btc": "qtum", "salt-btc": "salt", "powr-btc": "powerledger", "sngls-btc": "singulardtv", "fuel-btc": "etherparty", "etc-btc": "ethereum classic", "btg-btc": "bitcoin gold", "fun-btc": "funfair", "loom-btc": "loom network", "rcn-btc": "ripio credit network", "bcpt-btc": "blockmason credit protocol", "lsk-btc": "lisk",
	"qsp-btc": "quantstamp", "rpx-btc": "red pulse", "wtc-btc": "walton", "aion-btc": "aion", "poa-btc": "poa network", "poe-btc": "po.et", "bat-btc": "basic attention token", "tnb-btc": "time new bank", "req-btc": "request network", "xem-btc": "nem", "grs-btc": "groetlcoin", "amb-btc": "ambrosus", "pivx-btc": "pivx", "wabi-btc": "wabi", "kmd-btc": "komodo", "trig-btc": "triggers", "mana-btc": "decentraland", "icn-btc": "iconomi", "bnt-btc": "bancor", "storj-btc": "storj", "knc-btc": "kybernetwork", "zen-btc": "zencash",
	*/
	//eth
	"eos-eth": "eos", "trx-eth": "tron", "zrx-eth": "0x", "bnb-eth": "binance coin", "xrp-eth": "ripple", "ppt-eth": "populous", "ada-eth": "ada", "ont-eth": "ontology", "neo-eth": "neo" /*"bcc-eth": "bitcoin cash",*/, "icx-eth": "icon", "ltc-eth": "litecoin", "gto-eth": "gifto", "arn-eth": "aeron", "xlm-eth": "lumen", "mtl-eth": "metal", "lun-eth": "lunyr", "xvg-eth": "verge", "elf-eth": "aelf", "iota-eth": "iota", "sky-eth": "skycoin", "zil-eth": "zilliqa",
	"ven-eth": "vechain", "wan-eth": "wancoin", "iost-eth": "iostoken" /*"bcn-eth": "bytecoin",*/, "cmt-eth": "cybermiles", "tusd-eth": "trueusd", "waves-eth": "waves", "nav-eth": "navcoin", "nano-eth": "nano", "dash-eth": "dash", "strat-eth": "stratis", "gvt-eth": "genesis vision", "ncash-eth": "nucleus vision", "snt-eth": "status network token", "omg-eth": "omisego", "brd-eth": "bread", "zec-eth": "zcash", "xmr-eth": "monero", "edo-eth": "eidoo", "mco-eth": "monaco", "steem-eth": "steem",
	/*"sub-eth": "substratum", "storm-eth": "storm", "qlc-eth": "qlink", "bqx-eth": "ethos", "eng-eth": "enigma", "ae-eth": "aeternity", "wings-eth": "wings", "gnt-eth": "golem", "qtum-eth": "qtum", "salt-eth": "salt", "powr-eth": "powerledger", "sngls-eth": "singulardtv", "fuel-eth": "etherparty", "etc-eth": "ethereum classic", "btg-eth": "bitcoin gold", "fun-eth": "funfair", "loom-eth": "loom network", "rcn-eth": "ripio credit network", "bcpt-eth": "blockmason credit protocol", "lsk-eth": "lisk",
	"qsp-eth": "quantstamp", "rpx-eth": "red pulse", "wtc-eth": "walton", "aion-eth": "aion", "poa-eth": "poa network", "poe-eth": "po.et", "bat-eth": "basic attention token", "tnb-eth": "time new bank", "req-eth": "request network", "xem-eth": "nem", "grs-eth": "groetlcoin", "amb-eth": "ambrosus", "pivx-eth": "pivx", "wabi-eth": "wabi", "kmd-eth": "komodo", "trig-eth": "triggers", "mana-eth": "decentraland", "icn-eth": "iconomi", "bnt-eth": "bancor", "storj-eth": "storj", "knc-eth": "kybernetwork", "zen-eth": "zencash",
	*/
	//bnb
	//usdt
}

//GetValidPairsMap for an exchange
func GetValidPairsMap(exch string) map[string]string {
	switch exch {
	case Bittrex:
		return validBittrexPairs
	case Binance:
		return validBinancePairs
	case Gdax:
		return validGdaxPairs
	default:
		return map[string]string{}
	}
}

//ListValidPairs list all valid pairs
func ListValidPairs(lang string, exch string, quote string) ([]string, error) {
	if !IsValidExch(exch) {
		return nil, NewErrorInvalidExchange(lang, exch)
	}
	switch exch {
	case Bittrex:
		res := []string{} //make([]string, len(validBittrexPairs))
		i := 0
		for k := range validBittrexPairs {
			if quote != "" {
				if strings.HasSuffix(k, quote) {
					res = append(res, k)
					i++
				}
			} else {
				//res[i] = k
				res = append(res, k)
				i++
			}
		}
		sort.Slice(res, func(i, j int) bool {
			return res[i] < res[j]
		})
		return res, nil
	case Gdax:
		res := []string{} //make([]string, len(validBittrexPairs))
		i := 0
		for k := range validGdaxPairs {
			if quote != "" {
				if strings.HasSuffix(k, quote) {
					res = append(res, k)
					i++
				}
			} else {
				//res[i] = k
				res = append(res, k)
				i++
			}
		}
		sort.Slice(res, func(i, j int) bool {
			return res[i] < res[j]
		})
		return res, nil
	case Binance:
		res := []string{} //make([]string, len(validBittrexPairs))
		i := 0
		for k := range validBinancePairs {
			if quote != "" {
				if strings.HasSuffix(k, quote) {
					res = append(res, k)
					i++
				}
			} else {
				//res[i] = k
				res = append(res, k)
				i++
			}
		}
		sort.Slice(res, func(i, j int) bool {
			return res[i] < res[j]
		})
		return res, nil
	default:
		return nil, NewErrorInvalidExchange(lang, exch)
	}

}

//ListValidExchanges returns the names of the valid exchanges
func ListValidExchanges() []string {
	return listValidExchanges()
}
