package redisctl

const (
	// GroupFormat is the key format
	GroupFormat = "GF_G_%s"
	// PlayerCountFormat player count
	PlayerCountFormat = "GF_%s_PCNT"
	// GroupCountFormat group count
	GroupCountFormat = "GF_%s_GCNT"
	// GameEventFormat is the key format
	GameEventFormat = "GF_GE_%s"
	// GameEventCountFormat is the key format
	GameEventCountFormat = "GF_GECNT_%s"

	// GroupChannel Redis channel
	GroupChannel = "group_chan"
	// PlayerChannel Redis channel
	PlayerChannel = "player_chann"
)
