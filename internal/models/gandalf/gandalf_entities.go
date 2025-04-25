package gandalfmodels

type EntityMetaData struct {
	Name               string `json:"name"`
	Founder            string `json:"founder"`
	CoFounder          string `json:"co_founder"`
	ClubEmail          string `json:"club_email"`
	Description        string `json:"description"`
	YearEstablished    int    `json:"year_established"`
	ClubLogoImageUrl   string `json:"club_logo_image_url"`
	ClubBannerImageUrl string `json:"club_banner_image_url"`

	ClubWebsiteUrl   string `json:"club_website_url"`
	ClubTwitterUrl   string `json:"club_twitter_url"`
	ClubYoutubeUrl   string `json:"club_youtube_url"`
	ClubFacebookUrl  string `json:"club_facebook_url"`
	ClubLinkedinUrl  string `json:"club_linkedin_url"`
	ClubInstagramUrl string `json:"club_instagram_url"`

	CurrentCoreMembers []*CurrentCoreMembersMapper `json:"current_core_members"`
}

type CurrentCoreMembersMapper struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	ImageUrl string `json:"image_url"`
}
