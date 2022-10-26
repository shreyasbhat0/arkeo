package keeper

import (
	"mercury/common"
	"mercury/x/mercury/types"

	. "gopkg.in/check.v1"
)

type ModProviderSuite struct{}

var _ = Suite(&ModProviderSuite{})

func (ModProviderSuite) TestValidate(c *C) {
	ctx, k := SetupKeeper(c)

	s := newMsgServer(k)

	// happy path
	msg := types.MsgModProvider{
		MinContractDuration: 10,
		MaxContractDuration: 500,
		Status:              types.ProviderStatus_Online,
	}
	c.Assert(s.ModProviderValidate(ctx, &msg), IsNil)

	// bad min duration
	msg.MinContractDuration = 5256000 * 2
	err := s.ModProviderValidate(ctx, &msg)
	c.Check(err, ErrIs, types.ErrInvalidModProviderMinContractDuration)

	// bad max duration
	msg.MinContractDuration = 10
	msg.MaxContractDuration = 5256000 * 2
	err = s.ModProviderValidate(ctx, &msg)
	c.Check(err, ErrIs, types.ErrInvalidModProviderMaxContractDuration)
}

func (ModProviderSuite) TestHandle(c *C) {
	ctx, k := SetupKeeper(c)

	s := newMsgServer(k)

	// setup
	pubkey := types.GetRandomPubKey()
	acct, err := pubkey.GetMyAddress()
	c.Assert(err, IsNil)

	// happy path
	msg := types.MsgModProvider{
		Creator:             acct.String(),
		PubKey:              pubkey,
		Chain:               common.BTCChain,
		MetadataURI:         "foobar",
		MetadataNonce:       3,
		MinContractDuration: 10,
		MaxContractDuration: 500,
		Status:              types.ProviderStatus_Online,
		SubscriptionRate:    11,
		PayAsYouGoRate:      12,
	}
	c.Assert(s.ModProviderHandle(ctx, &msg), IsNil)

	provider, err := k.GetProvider(ctx, msg.PubKey, msg.Chain)
	c.Assert(err, IsNil)
	c.Check(provider.MetadataURI, Equals, "foobar")
	c.Check(provider.MetadataNonce, Equals, uint64(3))
	c.Check(provider.MinContractDuration, Equals, uint64(10))
	c.Check(provider.MaxContractDuration, Equals, uint64(500))
	c.Check(provider.Status, Equals, types.ProviderStatus_Online)
	c.Check(provider.SubscriptionRate, Equals, int64(11))
	c.Check(provider.PayAsYouGoRate, Equals, int64(12))
}