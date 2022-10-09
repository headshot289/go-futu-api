package futuapi

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/headshot289/go-futu-api/pb/trdcommon"
	"log"
	"testing"

	"github.com/headshot289/go-futu-api/pb/qotcommon"
)

func TestMd5(t *testing.T) {
	pwd := ""
	h := md5.New()
	if _, err := h.Write([]byte(pwd)); err != nil {
		t.Error(err)
	}
	//s := (string)(h.Sum(nil))
	s := hex.EncodeToString(h.Sum(nil))
	t.Logf(s)
}

func TestPlaceOrder(t *testing.T) {
	api := NewFutuAPI()
	defer api.Close(context.Background())

	if err := api.Connect(context.Background(), ":11111"); err != nil {
		t.Error(err)
		return
	}

	if err := api.UnlockTrade(context.Background(), true, "", trdcommon.SecurityFirm_SecurityFirm_FutuSecurities); err != nil {
		t.Error(err)
		return
	} else {
		t.Log("unlock trade success")
	}

	if trdAccList, err := api.GetAccList(context.Background()); err != nil {
		t.Error(err)
		return
	} else {
		for _, acc := range trdAccList {
			log.Printf("accId=%d accType=%d cardNum=%s firm=%d simAccType=%d trdEnv=%d auth=%+v", acc.AccID, acc.AccType, acc.CardNum, acc.SecurityFirm, acc.SimAccType, acc.TrdEnv, acc.TrdMarketAuthList)
		}
	}

	if orderId, err := api.PlaceOrder(context.Background(),
		&TrdHeader{TrdEnv: trdcommon.TrdEnv_TrdEnv_Real, AccID: 281756456005693483, TrdMarket: trdcommon.TrdMarket_TrdMarket_HK},
		trdcommon.TrdSide_TrdSide_Buy,
		trdcommon.OrderType_OrderType_SpecialLimit, "51214", 5000, 0.026,
		false, 0, trdcommon.TrdSecMarket_TrdSecMarket_HK, "", trdcommon.TimeInForce_TimeInForce_DAY, false); err != nil {
		t.Error(err)
		return
	} else {
		t.Log("orderId", orderId)
	}
}

func TestConnect(t *testing.T) {
	api := NewFutuAPI()
	defer api.Close(context.Background())

	api.SetRecvNotify(true)
	nCh, err := api.SysNotify(context.Background())
	if err != nil {
		t.Error(err)
	}

	if err := api.Connect(context.Background(), ":11111"); err != nil {
		t.Error(err)
		return
	}

	if sub, err := api.QuerySubscription(context.Background(), true); err != nil {
		t.Error(err)
	} else {
		t.Error(sub)
	}

	tCh, err := api.UpdateTicker(context.Background())
	if err != nil {
		t.Error(err)
	}
	if err := api.Subscribe(context.Background(), []*Security{{qotcommon.QotMarket_QotMarket_HK_Security, "00700"}}, []qotcommon.SubType{qotcommon.SubType_SubType_Ticker}, true, true, true, true); err != nil {
		t.Error(err)
	}
	select {
	case notify := <-nCh:
		t.Error(notify)
	case ticker := <-tCh:
		t.Error(ticker)
	}

	if sub, err := api.QuerySubscription(context.Background(), true); err != nil {
		t.Error(err)
	} else {
		t.Error(sub)
	}

	secs, err := api.GetUserSecurity(context.Background(), "全部")
	if err != nil {
		t.Error(err)
	} else {
		t.Error(secs)
	}
}
