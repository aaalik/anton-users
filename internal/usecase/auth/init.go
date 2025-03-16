package auth

func New(
	userRP iUserRepo,
	jcu iJwtConfUtils,
	dbu iDatabaseUtils,
	ru iRandomUtils,
	hu iHasherUtils,
	jwu iJwtUtils,
) *AuthUsecase {
	return &AuthUsecase{
		ur:  userRP,
		jcu: jcu,
		dbu: dbu,
		ru:  ru,
		hu:  hu,
		jwu: jwu,
	}
}
