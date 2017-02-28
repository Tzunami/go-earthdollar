package types

import (
	"math/big"
	"testing"

	"bytes"
	"github.com/Tzunami/go-earthdollar/common"
	"github.com/Tzunami/go-earthdollar/crypto"
)

func TestChainId(t *testing.T) {
	key, _ := defaultTestKey()

	tx := NewTransaction(0, common.Address{}, new(big.Int), new(big.Int), new(big.Int), nil)
	tx.SetSigner(NewChainIdSigner(big.NewInt(1)))

	var err error
	tx, err = tx.SignECDSA(key)
	if err != nil {
		t.Fatal(err)
	}

	tx.SetSigner(NewChainIdSigner(big.NewInt(2)))
	_, err = tx.From()
	if err != ErrInvalidChainId {
		t.Error("expected error:", ErrInvalidChainId)
	}

	tx.SetSigner(NewChainIdSigner(big.NewInt(1)))
	_, err = tx.From()
	if err != nil {
		t.Error("expected no error")
	}
}

func TestClassicChainId(t *testing.T) {
	key, _ := defaultTestKey()

	tx := NewTransaction(0, common.Address{}, new(big.Int), new(big.Int), new(big.Int), nil)
	tx.SetSigner(NewChainIdSigner(big.NewInt(61)))

	txs, err := tx.SignECDSA(key)
	if err != nil {
		t.Fatal(err)
	}

	if txs.data.V.Cmp(big.NewInt(157)) != 0 && txs.data.V.Cmp(big.NewInt(158)) != 0 {
		t.Errorf("V %v != 157 || 158", txs.data.V)
	}

	v := normaliseV(NewChainIdSigner(big.NewInt(61)), big.NewInt(157))
	if v != 27 {
		t.Errorf("Invalid V %v", v)
	}

	chainId := deriveChainId(big.NewInt(157))
	if chainId.Cmp(big.NewInt(61)) != 0 {
		t.Errorf("Invalid ChainId %v", chainId)
	}

	if !isProtectedV(big.NewInt(157)) {
		t.Error("Unprotected for 157")
	}

}

func TestMordenChainId(t *testing.T) {
	key, _ := defaultTestKey()

	tx := NewTransaction(0, common.Address{}, new(big.Int), new(big.Int), new(big.Int), nil)
	tx.SetSigner(NewChainIdSigner(big.NewInt(62)))

	txs, err := tx.SignECDSA(key)
	if err != nil {
		t.Fatal(err)
	}

	if txs.data.V.Cmp(big.NewInt(160)) != 0 && txs.data.V.Cmp(big.NewInt(159)) != 0 {
		t.Errorf("V %v != 159 || 160", txs.data.V)
	}

	v := normaliseV(NewChainIdSigner(big.NewInt(62)), big.NewInt(160))
	if v != 28 {
		t.Errorf("Invalid V %v", v)
	}
	v = normaliseV(NewChainIdSigner(big.NewInt(62)), big.NewInt(159))
	if v != 27 {
		t.Errorf("Invalid V %v", v)
	}

	chainId := deriveChainId(big.NewInt(160))
	if chainId.Cmp(big.NewInt(62)) != 0 {
		t.Errorf("Invalid ChainId %v", chainId)
	}

	if !isProtectedV(big.NewInt(160)) {
		t.Error("Unprotected for 160")
	}

}

func TestCompatibleSign(t *testing.T) {
	priv, err := crypto.GenerateKey()
	pub := crypto.FromECDSAPub(&priv.PublicKey)
	addr := crypto.PubkeyToAddress(priv.PublicKey)

	tx := NewTransaction(0, addr, new(big.Int), new(big.Int), new(big.Int), nil)
	tx.SetSigner(NewChainIdSigner(big.NewInt(61)))
	tx2 := NewTransaction(0, addr, new(big.Int), new(big.Int), new(big.Int), nil)
	tx2.SetSigner(BasicSigner{})

	tx, err = tx.SignECDSA(priv)
	if err != nil {
		t.Fatal(err)
	}
	tx2, err = tx2.SignECDSA(priv)
	if err != nil {
		t.Fatal(err)
	}

	pub_tx, err := tx.signer.PublicKey(tx)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(pub_tx, pub) {
		t.Errorf("Incorrect pubkey for ChainId Signer:\n%v\n%v", common.ToHex(pub), common.ToHex(pub_tx))
	}

	pub_tx2, err := tx.signer.PublicKey(tx)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(pub_tx2, pub) {
		t.Errorf("Incorrect pubkey for Basic Signer:\n%v\n%v", common.ToHex(pub), common.ToHex(pub_tx2))
	}
}
