package mal

import (
	"fmt"
	"net/url"
	"strconv"

	"go.uber.org/zap"
)

func (client *Client) SearchAnime(q string, opts ...OptFn) (*AnimeSearchPage, error) {
	log.Debug("SearchAnime()", zap.String("query", q))
	if q == "" {
		return nil, fmt.Errorf("query string is empty")
	}

	values := make(url.Values)
	values.Set("q", q)
	for _, opt := range opts {
		opt(values)
	}

	u := baseURL + "/anime?" + values.Encode()

	var page AnimeSearchPage
	err := client.doGet(u, &page)
	if err != nil {
		return nil, err
	}

	return &page, nil
}

type RankingType string

const (
	RankingAll          RankingType = "all"
	RankingAiring       RankingType = "airing"
	RankingUpcoming     RankingType = "upcoming"
	RankingTV           RankingType = "tv"
	RankingOVA          RankingType = "ova"
	RankingMovie        RankingType = "movie"
	RankingSpecial      RankingType = "special"
	RankingByPopularity RankingType = "bypopularity"
	RankingFavorite     RankingType = "favorite"
)

func (client *Client) GetAnimeDetails(animeID int, opts ...OptFn) (*Anime, error) {
	log.Debug("GetAnimeDetails()", zap.Int("animeID", animeID))
	values := make(url.Values)
	for _, opt := range opts {
		opt(values)
	}

	u := baseURL + "/anime/" + strconv.FormatInt(int64(animeID), 10) + "?" + values.Encode()

	var anime Anime
	err := client.doGet(u, &anime)
	if err != nil {
		return nil, err
	}

	return &anime, nil
}

func (client *Client) GetAnimeRanking(ranking RankingType, opts ...OptFn) (*AnimeRankingPage, error) {
	log.Debug("GetAnimeRanking()", zap.String("ranking", string(ranking)))
	if ranking == "" {
		return nil, fmt.Errorf("ranking is empty")
	}

	values := make(url.Values)
	for _, opt := range opts {
		opt(values)
	}

	u := baseURL + "/anime/ranking?" + values.Encode()

	var rankList AnimeRankingPage
	err := client.doGet(u, &rankList)
	if err != nil {
		return nil, err
	}

	return &rankList, nil
}

func (client *Client) GetSeasonalAnime(
	year int,
	season Season,
	opts ...OptFn,
) (*SeasonalAnimePage, error) {
	values := make(url.Values)
	for _, opt := range opts {
		opt(values)
	}

	u := fmt.Sprintf("%s/anime/season/%d/%s", baseURL, year, season)
	valuesStr := values.Encode()
	if valuesStr != "" {
		u += "?" + valuesStr
	}

	var list SeasonalAnimePage
	err := client.doGet(u, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}

func (client *Client) GetSuggestedAnime(opts ...OptFn) (*SuggestedAnimePage, error) {
	values := make(url.Values)
	for _, opt := range opts {
		opt(values)
	}
	valuesStr := values.Encode()

	u := baseURL + "/anime/suggestions"
	if valuesStr != "" {
		u += "?" + valuesStr
	}

	var list SuggestedAnimePage
	err := client.doGet(u, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
