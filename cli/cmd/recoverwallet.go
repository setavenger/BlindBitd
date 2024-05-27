package cmd

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/setavenger/blindbitd/cli/lib"
	"github.com/setavenger/blindbitd/pb"
	"golang.org/x/crypto/ssh/terminal"
	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

// allow seed passphrases only after thorough testing
var (
	birthHeight       uint64
	useSeedPassphrase bool
	labelCount        uint32

	// todo recover with label count
	recoverwalletCmd = &cobra.Command{
		Use:   "recoverwallet",
		Short: "Recover a wallet from mnemonic seed",
		Long: `birthheight is required, if you want to scan the entire chain then set it to 1.
You will be prompted to enter your mnemonic.
    The number of labels needs to be set if labels existed.
    Otherwise transactions involving labels will not be found.`, // this could be changed to scan from a certain Bip352 activation height unless explicitly overridden
		Run: func(cmd *cobra.Command, args []string) {
			client, conn := lib.NewClient(socketPath)
			defer func(conn *grpc.ClientConn) {
				err := conn.Close()
				if err != nil {
					panic(err)
				}
			}(conn)

			fmt.Println(
				"NOTE: The encryption password is only to encrypt your wallet data (keys, utxos, etc.) on disk." +
					"\nIt is not used for your seed. To add a passphrase to your seed set the --seedpass flag (not extensively tested yet)",
			)
			fmt.Print("Encryption password: ")
			passwordBytes, err := terminal.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				log.Fatalln("Error reading password")
			}
			fmt.Println()
			fmt.Print("Confirm password: ")
			passworConfirmBytes, err := terminal.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				log.Fatalln("Error reading password")
			}
			fmt.Println()
			if !bytes.Equal(passwordBytes, passworConfirmBytes) {
				log.Fatalln("Passwords do not match")
			}

			// todo bring back once tested thoroughly
			var seedPassphrase string
			if useSeedPassphrase {
				fmt.Print("Enter your seed passphrase: ")
				seedPassphraseBytes, err := terminal.ReadPassword(int(os.Stdin.Fd()))
				if err != nil {
					log.Fatalln("Error reading seed passphrase")
				}
				seedPassphrase = string(seedPassphraseBytes)
				fmt.Println()
			}

			fmt.Print("Input mnemonic: ")
			mnemonicBytes, err := terminal.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				log.Fatalln("Error reading mnemonic")
			}
			fmt.Println()

			mnemonic := string(mnemonicBytes)

			// we always start scanning from block one
			if birthHeight < 1 {
				birthHeight = 1
			}

			response, err := client.RecoverWallet(context.Background(), &pb.RecoverWalletRequest{EncryptionPassword: string(passwordBytes), SeedPassphrase: &seedPassphrase, Mnemonic: mnemonic, BirthHeight: birthHeight, LabelCount: labelCount})
			if err != nil {
				log.Fatalln(err)
			}

			if response.Success {
				fmt.Println("Success")
			} else {
				fmt.Printf("Failed with error: %s", response.Error)
			}
		},
	}
)

func init() {
	RootCmd.AddCommand(recoverwalletCmd)

	recoverwalletCmd.PersistentFlags().Uint64Var(&birthHeight, "birthheight", 0, "set the birth height for a recovered wallet")
	recoverwalletCmd.PersistentFlags().Uint32Var(&labelCount, "labelcount", 0, "set the number of labels which should be created")
	recoverwalletCmd.PersistentFlags().BoolVar(&useSeedPassphrase, "seedpass", false, "add a passphrase to the wallet seed")

	err := cobra.MarkFlagRequired(recoverwalletCmd.PersistentFlags(), "birthheight")
	if err != nil {
		log.Fatalln(err)
	}
}
