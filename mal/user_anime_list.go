package mal

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func (client *Client) GetMyAnimeList(opts AnimeListOptions, commonOpts ...OptFn) (*UserAnimeListPage, error) {
	return client.GetUserAnimeList("@me", opts, commonOpts...)
}

func (client *Client) GetUserAnimeList(user string, opts AnimeListOptions, commonOpts ...OptFn) (*UserAnimeListPage, error) {
	log := log.Named("AnimeList()")
	log.Debug("call",
		zap.String("user", user),
		zap.Any("opts", opts),
		zap.Int("num common opts", len(commonOpts)))

	vals := make(url.Values)

	opts.Set(vals)
	log.Debug("url values", zap.Any("values", vals))

	for _, opt := range commonOpts {
		opt(vals)
	}

	u := baseURL + "/users/" + user + "/animelist?" + vals.Encode()

	var page UserAnimeListPage
	err := client.doGet(u, &page)
	if err != nil {
		return nil, err
	}

	return &page, nil
}

func (client *Client) UpdateAnimeList(animeID int, opts *UpdateListOptsBuilder) (*AnimeUpdateResponse, error) {
	log := log.Named("UpdateList()")
	log.Debug("call",
		zap.Int("animeID", animeID),
		zap.Any("opts", opts.values))

	u := fmt.Sprintf("%s/anime/%d/my_list_status", baseURL, animeID)

	body := opts.values.Encode()

	var updateResp AnimeUpdateResponse
	err := client.doPatch(u, strings.NewReader(body), &updateResp)
	if err != nil {
		return nil, err
	}

	return &updateResp, nil
}

func (client *Client) DeleteAnimeListItem(animeID int) error {
	log := log.Named("DeleteAnimeListItem()")
	log.Debug("call",
		zap.Int("animeID", animeID))

	u := fmt.Sprintf("%s/anime/%d/my_list_status", baseURL, animeID)

	return client.doDelete(u)
}

func UpdateListOpts() *UpdateListOptsBuilder {
	return &UpdateListOptsBuilder{values: make(url.Values)}
}

type UpdateListOptsBuilder struct {
	values url.Values
}

func (b *UpdateListOptsBuilder) Status(s StatusOpt) *UpdateListOptsBuilder {
	b.values.Set("status", string(s))
	return b
}

func (b *UpdateListOptsBuilder) IsRewatching(rewatching bool) *UpdateListOptsBuilder {
	b.values.Set("is_rewatching", strconv.FormatBool(rewatching))
	return b
}

func (b *UpdateListOptsBuilder) Score(score int) *UpdateListOptsBuilder {
	score = clamp(score, 0, 10)

	b.values.Set("score", strconv.FormatInt(int64(score), 10))
	return b
}

func (b *UpdateListOptsBuilder) NumWatchedEpisodes(eps int) *UpdateListOptsBuilder {
	b.values.Set("num_watched_episodes", strconv.FormatInt(int64(eps), 10))
	return b
}

func (b *UpdateListOptsBuilder) Priority(p int) *UpdateListOptsBuilder {
	p = clamp(p, 0, 2)
	b.values.Set("priority", strconv.FormatInt(int64(p), 10))
	return b
}

func (b *UpdateListOptsBuilder) NumTimesRewatched(n int) *UpdateListOptsBuilder {
	b.values.Set("num_times_rewatched", strconv.FormatInt(int64(n), 10))
	return b
}

func (b *UpdateListOptsBuilder) RewatchValue(v int) *UpdateListOptsBuilder {
	v = clamp(v, 0, 5)

	b.values.Set("rewatch_value", strconv.FormatInt(int64(v), 10))
	return b
}

func (b *UpdateListOptsBuilder) Tags(tags string) *UpdateListOptsBuilder {
	b.values.Set("tags", tags)
	return b
}

func (b *UpdateListOptsBuilder) Comments(comments string) *UpdateListOptsBuilder {
	b.values.Set("comments", comments)
	return b
}

type AnimeListOptions struct {
	Status StatusOpt
	Sort   SortOpt
}

func (opts AnimeListOptions) Set(v url.Values) {
	if opts.Status != "" {
		v.Set("status", string(opts.Status))
	}

	if opts.Sort != "" {
		v.Set("sort", string(opts.Sort))
	}
}

type StatusOpt string

const (
	StatusWatching    StatusOpt = "watching"
	StatusCompleted   StatusOpt = "completed"
	StatusOnHold      StatusOpt = "on_hold"
	StatusDropped     StatusOpt = "dropped"
	StatusPlanToWatch StatusOpt = "plan_to_watch"
)

type SortOpt string

const (
	SortListScore      SortOpt = "list_score"
	SortListUpdatedAt  SortOpt = "list_updated_at"
	SortAnimeTitle     SortOpt = "anime_title"
	SortAnimeStartDate SortOpt = "anime_start_date"
	SortAnimeID        SortOpt = "anime_id"
)

func clamp(i, min, max int) int {
	if i > max {
		return max
	}
	if i < min {
		return min
	}
	return i
}
