// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	"github.com/33cn/chain33/types"
	gt "github.com/33cn/plugin/plugin/dapp/game/types"
)

func (g *Game) execDelLocal(receiptData *types.ReceiptData) (*types.LocalDBSet, error) {
	dbSet := &types.LocalDBSet{}
	if receiptData.GetTy() != types.ExecOk {
		return dbSet, nil
	}
	for _, log := range receiptData.Logs {
		switch log.GetTy() {
		case gt.TyLogCreateGame, gt.TyLogMatchGame, gt.TyLogCloseGame, gt.TyLogCancleGame:
			receiptGame := &gt.ReceiptGame{}
			if err := types.Decode(log.Log, receiptGame); err != nil {
				return nil, err
			}
			kv := g.rollbackIndex(receiptGame)
			dbSet.KV = append(dbSet.KV, kv...)
		}
	}
	return dbSet, nil
}

func (g *Game) ExecDelLocal_Create(payload *gt.GameCreate, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return g.execDelLocal(receiptData)
}

func (g *Game) ExecDelLocal_Cancel(payload *gt.GameCancel, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return g.execDelLocal(receiptData)
}

func (g *Game) ExecDelLocal_Close(payload *gt.GameClose, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return g.execDelLocal(receiptData)
}

func (g *Game) ExecDelLocal_Match(payload *gt.GameMatch, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return g.execDelLocal(receiptData)
}
