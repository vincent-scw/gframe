package contracts

// GetID return id
func (ue *UserEvent) GetID() string {
	//TODO: should use match ID
	return ue.User.GetID()
}
