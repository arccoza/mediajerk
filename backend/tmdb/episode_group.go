package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// https://developer.themoviedb.org/reference/tv-episode-group-details
// https://api.themoviedb.org/3/tv/episode_group/{tv_episode_group_id}
func (cl *Client) EpisodesGroupedBy(groupId string) (*EpisodeGroupDetails, error) {
	path := "tv/episode_group/" + groupId

	resp, err := cl.get(path, url.Values{})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var groupDetails EpisodeGroupDetails
	if err := json.Unmarshal(body, &groupDetails); err != nil {
		return nil, err
	}

	return &groupDetails, nil
}
