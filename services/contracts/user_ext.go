package contracts

// GetId return id
func (ue *UserEvent) GetId() string {
	return ue.User.GetId()
}