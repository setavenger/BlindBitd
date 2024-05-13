package cmd

import (
	"context"
	"fmt"
	"github.com/setavenger/blindbitd/cli/lib"
	"github.com/setavenger/blindbitd/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"os"
	"text/tabwriter"
)

// balanceCmd represents the balance command
var (
	listUTXOs bool

	listAll bool

	showUnconfirmed      = false
	showUnspent          = false
	showSpent            = false
	showSpentUnConfirmed = false

	states []pb.UTXOState

	balanceCmd = &cobra.Command{
		Use:   "balance",
		Short: "shows the balance of the wallet",
		Long:  `Daemon needs to be unlocked. Shows the balance of the wallet.`,
		Run: func(cmd *cobra.Command, args []string) {
			client, conn := lib.NewClient(socketPath)
			defer func(conn *grpc.ClientConn) {
				err := conn.Close()
				if err != nil {
					panic(err)
				}
			}(conn)

			if showUnconfirmed {
				states = append(states, pb.UTXOState_UNCONFIRMED)
			}
			if showUnspent {
				states = append(states, pb.UTXOState_UNSPENT)
			}
			if showSpent {
				states = append(states, pb.UTXOState_SPENT)
			}
			if showSpentUnConfirmed {
				states = append(states, pb.UTXOState_SPENT_UNCONFIRMED)
			}
			if len(states) == 0 {
				states = append(states, pb.UTXOState_UNSPENT)
			}

			utxos, err := client.ListUTXOs(context.Background(), &pb.Empty{})
			if err != nil {
				log.Fatalf("Error: Getting UTXOs failed: %v\n", err)
			}

			var filteredUTXOs []*pb.OwnedUTXO

			if listAll {
				filteredUTXOs = utxos.Utxos
			} else {
				for _, state := range states {
					for _, utxo := range utxos.Utxos {
						if utxo.UtxoState == state {
							filteredUTXOs = append(filteredUTXOs, utxo)
						}
					}
				}
			}

			if listUTXOs {
				writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
				_, err := fmt.Fprintln(writer, "UTXO Outpoint\tAmount\tState\tLabel")
				if err != nil {
					log.Fatalln(err)
				}

				for _, utxo := range filteredUTXOs {
					amount := lib.ConvertIntToThousandString(int(utxo.Amount))
					output := fmt.Sprintf("%x:%d\t%s\t%s", utxo.Txid, utxo.Vout, amount, utxo.UtxoState)
					if utxo.Label != nil && utxo.Label.Comment != "" {
						output += fmt.Sprintf("\t%s", utxo.Label.Comment)
					} else {
						output += "\t"
					}
					_, err := fmt.Fprintln(writer, output)
					if err != nil {
						log.Fatalln(err)
					}
				}

				err = writer.Flush()
				if err != nil {
					log.Fatalln(err)
				}
				return
			} else {
				var balance uint64

				for _, utxo := range filteredUTXOs {
					balance += utxo.Amount
				}
				fmt.Printf("Balance is %s\n", lib.ConvertIntToThousandString(int(balance)))
				return
			}

		},
	}
)

func init() {
	RootCmd.AddCommand(balanceCmd)

	balanceCmd.PersistentFlags().BoolVar(&listUTXOs, "list", false, "list utxos instead showing the balance")
	balanceCmd.PersistentFlags().BoolVar(&listAll, "all", false, "list all states")

	balanceCmd.PersistentFlags().BoolVar(&showUnconfirmed, "unconfirmed", false, "add unconfirmed utxos to the filter")
	balanceCmd.PersistentFlags().BoolVar(&showUnspent, "unspent", false, "add unspent utxos to the filter")
	balanceCmd.PersistentFlags().BoolVar(&showSpent, "spent", false, "add spent utxos to the filter")
	balanceCmd.PersistentFlags().BoolVar(&showSpentUnConfirmed, "spentunconf", false, "add spent utxos whose spending transaction is not confirmed to the filter")

}
