package mal

import (
	"net/url"
	"strconv"
	"strings"
)

func Fields(fields ...string) OptFn {
	return func(v url.Values) {
		v.Set("fields", strings.Join(fields, ","))
	}
}

func Limit(n int) OptFn {
	return func(v url.Values) {
		v.Set("limit", strconv.FormatInt(int64(n), 10))
	}
}

func Offset(n int) OptFn {
	return func(v url.Values) {
		v.Set("offset", strconv.FormatInt(int64(n), 10))
	}
}

func Nsfw() OptFn {
	return func(v url.Values) {
		v.Set("nsfw", "true")
	}
}

func SortSeasonalAnime(sort SeasonalSort) OptFn {
	return func(v url.Values) {
		v.Set("sort", string(sort))
	}
}

type SeasonalSort string

const (
	SeasonalSortAnimeScore        SeasonalSort = "anime_score"
	SeasonalSortAnimeNumListUsers SeasonalSort = "anime_num_list_users"
)
