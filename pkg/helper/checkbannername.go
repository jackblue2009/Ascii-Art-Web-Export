package helper

// function check the banner requested match with the list available in website
func CheckBannerName(banner string) bool {
	banners := []string{"standard", "shadow", "thinkertoy", "3d"}
	for _, f := range banners {
		if f == banner {
			return true
		}
	}
	return false
}
