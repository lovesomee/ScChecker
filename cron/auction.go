package cron

import (
	"fmt"
	"github.com/robfig/cron"
	"sc-profile/service/auction"
)

const itemId = "kqgy"

type ScCron struct {
	AuctionService auction.IService
}

func NewScCron(auctionService auction.IService) *ScCron {
	return &ScCron{AuctionService: auctionService}
}

func (c *ScCron) Start() {
	scCron := cron.New()

	if err := scCron.AddFunc("0 */4 * * *", c.function); err != nil {
		fmt.Println(err)
	}

	scCron.Start()
}

func (c *ScCron) function() {
	if err := c.AuctionService.UpdateItemHistory(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("прости, если трахнул")
	}
}
