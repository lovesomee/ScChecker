package cron

import (
	"context"
	"fmt"
	"github.com/robfig/cron"
	"go.uber.org/zap"
	"sc-profile/service/auction"
	"time"
)

type ScCron struct {
	logger         *zap.Logger
	auctionService auction.IService
}

func NewScCron(logger *zap.Logger, auctionService auction.IService) *ScCron {
	return &ScCron{logger: logger, auctionService: auctionService}
}

func (c *ScCron) Start() {
	scCron := cron.New()

	if err := scCron.AddFunc("0 */4 * * *", c.function); err != nil {
		fmt.Println(err)
	}

	scCron.Start()
}

func (c *ScCron) function() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := c.auctionService.UpdateItemHistory(ctx); err != nil {
		c.logger.Error("error updating item history", zap.Error(err))
	} else {
		fmt.Println("by ПИСЯТДВА")
	}
}
