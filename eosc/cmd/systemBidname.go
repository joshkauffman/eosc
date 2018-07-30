package cmd

import (
	"fmt"

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/system"
	"github.com/spf13/cobra"
)

var systemBidnameCmd = &cobra.Command{
	Use:   "bidname [bidder_account_name] [premium_account_name] [bid quantity]",
	Short: "Bid on a premium account name",
	Long: `Bid on a premium account name

All fields are required. Example usage:

    eosc system bidname [your_account_name] [name_to_bid_on] [EOS_quantity]

Please note that your funds will be locked up until
you are awarded the namespace, or you are outbid.


Read https://steemit.com/eos/@eos-canada/everything-you-need-to-know-about-namespace-bidding-on-eos for more info.
`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		api := getAPI()

		bidder := toAccount(args[0], "bidder_account_name")
		newname := toAccount(args[1], "premium_account_name")
		bidAsset, err := eos.NewEOSAssetFromString(args[2])
		errorCheck("bid amount invalid", err)

		fmt.Printf("[%s] bidding for: %s , amount=%d precision=%d symbol=%s\n", bidder, newname, bidAsset.Amount, bidAsset.Symbol.Precision, bidAsset.Symbol.Symbol)

		pushEOSCActions(api,
			system.NewBidname(bidder, newname, bidAsset),
		)
	},
}

func init() {
	systemCmd.AddCommand(systemBidnameCmd)
}
