package tokens

type SignedDetails struct {
	Email      string
	First_Name string
	Last_Name  string
	Uid
	jwt.StandardClaims
}

func TokenGenerator(email, first_name, last_name, uid string) (signedtoken, signedrefreshtoken string, err error) {
	claims := &SignedDetails{
		Email:email,
		First_Name: first_name,
		Last_Name: last_name,
		Uid:uid,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour*time.Duration(24)).Unix(),
		} 
	}

	
}

func UpdateAllTokens() {

}

func ValidateToken() {

}
