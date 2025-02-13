package scapi

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"sc-profile/models"
	"strconv"
)

type IScApi interface {
	GetAuctionHistory(ctx context.Context, itemId, region string, limit int) (models.AuctionHistoryResponse, error)
}

type ScApi struct {
	logger     *zap.Logger
	httpClient *http.Client
	domain     string
}

func NewScApi(logger *zap.Logger, httpClient *http.Client, domain string) *ScApi {
	return &ScApi{logger: logger, httpClient: httpClient, domain: domain}
}

const bearerToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiI3MzIiLCJqdGkiOiI1N2YyNmQ1Zjk4YWI4YzkyNDVhZGMwNmJiMWZiYmMxNDQ1MWE5Mzc2NjdlOTRiMTRhMDdkYzc4YTA0MTBmYWFmNDU4YWFjY2M5MjJjOWI1NSIsImlhdCI6MTczNjYyNDgzNy4xNzMzNDcsIm5iZiI6MTczNjYyNDgzNy4xNzMzNDksImV4cCI6MTc2ODE2MDgzNy4xNjEwNTcsInN1YiI6IiIsInNjb3BlcyI6W119.Ced-MtCg6HB0cUCKxES4abhBmyvpfQEs4c-8jYxL9x1C5FWE0t6pE7dBnS4lDIMRU0qZiB5MHwJBdxF_5L7s90soRJukpR384mKb1F_65JA8LFpo9WCF83fY55fSdNCLX1T2HNRDFAqxnpY1zbqOPR0goz5D3q2ssxl6lXD9ywhxg-ugZOQRtGhTDyOhn-UVup5x-TUaLi-UDzlKxt3KrIqt5J9-qAI39OYqxvBf3pKP_SRV_rgBADX3FR9Y3L9V2aZRtYRaFR7QidYRjLEcm6cDO06DX8JxD9wdsS0S6o8CWBR-O617zqRjEiJbShhJJffj4e-1Acjbhr1W39wMwjQdmzEypc9jr2_q48B7en4oFtOcTxiHEY4E6P6r6cFvsuYD9x6zprzWVr8c3eWeNPz3ZV7VtUmTjpg0v995_Rz7nq-bSOG7qMlLMx34T0aeA4asiI8jmf9SpByFWKPx29ZvCZvd-8lJW3-0lp9aQMlFpr65J-olDS4f8bL2CgMfj9Sz85LCVsFlDijMJIhjvl1SelvsEIaSF49QkWx3TqWCvcVBXGD6citEEtiWbxmTVVo1oaz6UopKd0NiY-2rtbaF_6xQArH2SOA_Fv5sUVTEjITrzFg_vB05tuPp-2jhKfQtzcHOJYOM1ayqhcAFU33fLtmWXZu8F4_UarsoNUU"

func (s *ScApi) GetAuctionHistory(ctx context.Context, itemId, region string, limit int) (models.AuctionHistoryResponse, error) {
	strUrl := fmt.Sprintf("%s/%s/auction/%s/history", s.domain, region, itemId)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, strUrl, nil)
	if err != nil {
		return models.AuctionHistoryResponse{}, fmt.Errorf("new request error: %w", err)
	}

	q := req.URL.Query()
	q.Add("limit", strconv.Itoa(limit))
	q.Add("additional", strconv.FormatBool(true))
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", "Bearer "+bearerToken)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return models.AuctionHistoryResponse{}, fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	var auctionHistoryResponse models.AuctionHistoryResponse

	if err = json.NewDecoder(resp.Body).Decode(&auctionHistoryResponse); err != nil {
		return models.AuctionHistoryResponse{}, fmt.Errorf("json decode error: %w", err)
	}

	return auctionHistoryResponse, nil
}
