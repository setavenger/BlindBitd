package src

import (
	"encoding/json"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/setavenger/blindbitd/src/logging"
	"github.com/setavenger/go-bip352"
	"github.com/tyler-smith/go-bip39"
)

type Keys struct {
	ScanSecretKey  [32]byte
	SpendSecretKey [32]byte
	Mnemonic       string
}

func (k *Keys) Serialise() ([]byte, error) {
	return json.Marshal(k)
}

func (k *Keys) DeSerialise(data []byte) error {
	err := json.Unmarshal(data, k)
	if err != nil {
		logging.ErrorLogger.Println(err)
		return err
	}

	return nil
}

func CreateNewKeys(seedPassphrase string) (*Keys, error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		logging.ErrorLogger.Println(err)
		return nil, err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		logging.ErrorLogger.Println(err)
		return nil, err
	}

	return KeysFromMnemonic(mnemonic, seedPassphrase)
}

func KeysFromMnemonic(mnemonic, seedPassphrase string) (*Keys, error) {
	var result Keys
	result.Mnemonic = mnemonic

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, seedPassphrase)
	if err != nil {
		logging.ErrorLogger.Println(err)
		return nil, err
	}

	master, err := hdkeychain.NewMaster(seed, ChainParams)
	if err != nil {
		logging.ErrorLogger.Println(err)
		return nil, err
	}

	keys, err := DeriveKeysFromMaster(master)
	if err != nil {
		logging.ErrorLogger.Println(err)
		return nil, err
	}

	result.ScanSecretKey = keys.ScanSecretKey
	result.SpendSecretKey = keys.SpendSecretKey

	return &result, err

}

func DeriveKeysFromMaster(master *hdkeychain.ExtendedKey) (*Keys, error) {
	/*
		ScanDerivationPath = "m/352'/0'/0'/1'/0";
		SpendDerivationPath = "m/352'/0'/0'/0'/0";
	*/

	// m/352'
	purpose, err := master.Derive(352 + hdkeychain.HardenedKeyStart)
	if err != nil {
		logging.ErrorLogger.Println(err)
		return nil, err
	}

	var coinType *hdkeychain.ExtendedKey

	if ChainParams.Name == "mainnet" {
		// m/352'/0'
		coinType, err = purpose.Derive(0 + hdkeychain.HardenedKeyStart)
		if err != nil {
			logging.ErrorLogger.Println(err)
			return nil, err
		}
	} else {
		// m/352'/1'
		coinType, err = purpose.Derive(1 + hdkeychain.HardenedKeyStart)
		if err != nil {
			logging.ErrorLogger.Println(err)
			return nil, err
		}
	}

	// m/352'/0'/0'
	acct0, err := coinType.Derive(0 + hdkeychain.HardenedKeyStart)
	if err != nil {
		logging.ErrorLogger.Println(err)
		return nil, err
	}

	// m/352'/0'/0'/1'
	scanExternal, err := acct0.Derive(1 + hdkeychain.HardenedKeyStart)
	if err != nil {
		logging.ErrorLogger.Println(err)
		return nil, err
	}

	// m/352'/0'/0'/0'
	spendExternal, err := acct0.Derive(0 + hdkeychain.HardenedKeyStart)
	if err != nil {
		logging.ErrorLogger.Println(err)
		return nil, err
	}

	scanKey, err := scanExternal.Derive(0)
	if err != nil {
		logging.ErrorLogger.Println(err)
		return nil, err
	}

	spendKey, err := spendExternal.Derive(0)
	if err != nil {
		logging.ErrorLogger.Println(err)
		return nil, err
	}

	secretKeyScan, err := scanKey.ECPrivKey()
	if err != nil {
		logging.ErrorLogger.Println(err)
		return nil, err
	}

	secretKeySpend, err := spendKey.ECPrivKey()
	if err != nil {
		logging.ErrorLogger.Println(err)
		return nil, err
	}

	return &Keys{
		ScanSecretKey:  bip352.ConvertToFixedLength32(secretKeyScan.Serialize()),
		SpendSecretKey: bip352.ConvertToFixedLength32(secretKeySpend.Serialize()),
	}, nil
}
