package storage

import (
	"context"
	"sthl/constants"
	"sthl/ent"

	"go.uber.org/zap"
)

func WithTx(ctx context.Context, logger *zap.Logger, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		logger.Info("fail to create new transactional client", zap.Error(err))
		return constants.ErrInternalServer
	}
	defer func() {
		v := recover()
		if v != nil {
			logger.Info("panicking, defer recovering and rollingback", zap.Error(err))
			err := tx.Rollback()
			logger.Info("panicking, defer fail to tx.Rollback()", zap.Error(err))
		}
	}()
	if err := fn(tx); err != nil {
		logger.Info("fail to execute txc cb fn", zap.Error(err))
		if rerr := tx.Rollback(); rerr != nil {
			logger.Info("fail to rollback", zap.Error(rerr))
			err = constants.ErrInternalServer
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		logger.Info("fail to commit transaction", zap.Error(err))
		return constants.ErrInternalServer
	}
	return nil
}

func WithTxTest(ctx context.Context, logger *zap.Logger, client *ent.Client, fn func(tx *ent.Tx) error) error {
	mTx := new(ent.Tx)
	err := fn(mTx)
	if err != nil {
		return err
	}
	return nil
}
