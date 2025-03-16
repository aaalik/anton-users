package user

func New(userRP iUserRepo, dbu iDatabaseUtils, ru iRandomUtils, hu iHasherUtils) *UserUsecase {
	return &UserUsecase{
		ur:  userRP,
		dbu: dbu,
		ru:  ru,
		hu:  hu,
	}
}
