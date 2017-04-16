package tekleader

const (
	intraURL   string = "https://intra.epitech.eu/"
	jsonFormat string = "?format=json"
)

var (
	// AuthKey represents the auth token
	AuthKey string
	// Race represents the boolean to active/disable race conditon
	Race bool
	// Location represents the promotion location you want sort
	Location string
	// Promo represents the promotion you want sort
	Promo string
	// Course represents the course of the student
	Course string
	// Find represents the student you want to find
	Find string
	// Timeout represents the timeout to check is intra is alive
	Timeout int
)
