package mal

import "net/url"

type SuggestedAnimeListPage struct {
	Data   []SuggestedAnimeList `json:"data,omitempty"`
	Paging Paging               `json:"paging,omitempty"`
}
type SuggestedAnimeList struct {
	Node Anime `json:"node,omitempty"`
}

type SeasonalAnimeListPage struct {
	Data   []AnimeSearchList `json:"data,omitempty"`
	Paging Paging            `json:"paging,omitempty"`
	Season Season            `json:"season,omitempty"`
}
type SeasonalAnimeList struct {
	Node Anime `json:"node,omitempty"`
}

type AnimeSearchListPage struct {
	Data   []AnimeSearchList `json:"data,omitempty"`
	Paging Paging            `json:"paging,omitempty"`
}
type AnimeSearchList struct {
	Node Anime `json:"node,omitempty"`
}

type UserAnimeListPage struct {
	Data   []UserAnimeList `json:"data,omitempty"`
	Paging Paging          `json:"paging,omitempty"`
}
type UserAnimeList struct {
	Node       Anime        `json:"node,omitempty"`
	ListStatus MyListStatus `json:"list_status,omitempty"`
}

type AnimeRankingListPage struct {
	Data   []AnimeRankingList `json:"data,omitempty"`
	Paging Paging             `json:"paging,omitempty"`
}
type AnimeRankingList struct {
	Node    Anime `json:"node,omitempty"`
	Ranking struct {
		Rank int `json:"rank,omitempty"`
	} `json:"ranking,omitempty"`
}

type Paging struct {
	Previous string `json:"previous,omitempty"`
	Next     string `json:"next,omitempty"`
}

type Anime struct {
	ID                     int               `json:"id,omitempty"`
	Title                  string            `json:"title,omitempty"`
	MainPicture            Picture           `json:"main_picture,omitempty"`
	AlternativeTitles      AlternativeTitles `json:"alternative_titles,omitempty"`
	StartDate              string            `json:"start_date,omitempty"`
	EndDate                string            `json:"end_date,omitempty"`
	Synopsis               string            `json:"synopsis,omitempty"`
	Mean                   float64           `json:"mean,omitempty"`
	Rank                   int               `json:"rank,omitempty"`
	Popularity             int               `json:"popularity,omitempty"`
	NumListUsers           int               `json:"num_list_users,omitempty"`
	NumScoringUsers        int               `json:"num_scoring_users,omitempty"`
	Nsfw                   string            `json:"nsfw,omitempty"`
	Genres                 []Genre           `json:"genres,omitempty"`
	CreatedAt              string            `json:"created_at,omitempty"`
	UpdatedAt              string            `json:"updated_at,omitempty"`
	MediaType              string            `json:"media_type,omitempty"`
	Status                 string            `json:"status,omitempty"`
	MyListStatus           MyListStatus      `json:"my_list_status,omitempty"`
	NumEpisodes            int               `json:"num_episodes,omitempty"`
	StartSeason            SeasonInfo        `json:"start_season,omitempty"`
	Broadcast              Broadcast         `json:"broadcast,omitempty"`
	Source                 string            `json:"source,omitempty"`
	AverageEpisodeDuration int               `json:"average_episode_duration,omitempty"`
	Rating                 string            `json:"rating,omitempty"`
	Studios                []AnimeStudio     `json:"studios,omitempty"`
}

type Picture struct {
	Large  string `json:"large,omitempty"`
	Medium string `json:"medium,omitempty"`
}

type AlternativeTitles struct {
	Synonyms []string `json:"synonyms,omitempty"`
	En       string   `json:"en,omitempty"`
	Ja       string   `json:"ja,omitempty"`
}

type Genre struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SeasonInfo struct {
	Year   int    `json:"year,omitempty"`
	Season string `json:"season,omitempty"`
}

type Season string

const (
	SeasonWinter Season = "winter"
	SeasonSpring Season = "spring"
	SeasonSummer Season = "summer"
	SeasonFall   Season = "fall"
)

type Broadcast struct {
	DayOfTheWeek string `json:"day_of_the_week,omitempty"`
	StartTime    string `json:"start_time,omitempty"`
}

type AnimeStudio struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type MyListStatus struct {
	Status string `json:"status,omitempty"`
	Score  int    `json:"score,omitempty"`
	// TODO: num_episodes_watched or num_watched_episodes
	NumEpisodesWatched int      `json:"num_episodes_watched,omitempty"`
	IsRewatching       bool     `json:"is_rewatching,omitempty"`
	StartDate          string   `json:"start_date,omitempty"`
	FinishDate         string   `json:"finish_date,omitempty"`
	Priority           int      `json:"priority,omitempty"`
	NumTimesRewatched  int      `json:"num_times_rewatched,omitempty"`
	RewatchValue       int      `json:"rewatch_value,omitempty"`
	Tags               []string `json:"tags,omitempty"`
	Comments           string   `json:"comments,omitempty"`
	UpdatedAt          string   `json:"updated_at,omitempty"`
}

type AnimeUpdateResponse struct {
	Status             string   `json:"status,omitempty"`
	Score              int      `json:"score,omitempty"`
	NumEpisodesWatched int      `json:"num_episodes_watched,omitempty"`
	IsRewatching       bool     `json:"is_rewatching,omitempty"`
	UpdatedAt          string   `json:"updated_at,omitempty"`
	Priority           int      `json:"priority,omitempty"`
	NumTimesRewatched  int      `json:"num_times_rewatched,omitempty"`
	RewatchValue       int      `json:"rewatch_value,omitempty"`
	Tags               []string `json:"tags,omitempty"`
	Comments           string   `json:"comments,omitempty"`
}

type MyUserInformation struct {
	ID              int             `json:"id,omitempty"`
	Name            string          `json:"name,omitempty"`
	Location        string          `json:"location,omitempty"`
	JoinedAt        string          `json:"joined_at,omitempty"`
	AnimeStatistics AnimeStatistics `json:"anime_statistics,omitempty"`
}

type AnimeStatistics struct {
	NumItemsWatching    int     `json:"num_items_watching,omitempty"`
	NumItemsCompleted   int     `json:"num_items_completed,omitempty"`
	NumItemsOnHold      int     `json:"num_items_on_hold,omitempty"`
	NumItemsDropped     int     `json:"num_items_dropped,omitempty"`
	NumItemsPlanToWatch int     `json:"num_items_plan_to_watch,omitempty"`
	NumItems            int     `json:"num_items,omitempty"`
	NumDaysWatched      float64 `json:"num_days_watched,omitempty"`
	NumDaysWatching     float64 `json:"num_days_watching,omitempty"`
	NumDaysCompleted    float64 `json:"num_days_completed,omitempty"`
	NumDaysOnHold       int     `json:"num_days_on_hold,omitempty"`
	NumDaysDropped      int     `json:"num_days_dropped,omitempty"`
	NumDays             float64 `json:"num_days,omitempty"`
	NumEpisodes         int     `json:"num_episodes,omitempty"`
	NumTimesRewatched   int     `json:"num_times_rewatched,omitempty"`
	MeanScore           float64 `json:"mean_score,omitempty"`
}

type OptFn func(url.Values)

type ApiError struct {
	ErrorType  string `json:"error,omitempty"`
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"-"`
}

func (apiErr *ApiError) Error() string {
	return apiErr.Message
}
