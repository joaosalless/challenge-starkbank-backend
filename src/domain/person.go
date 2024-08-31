package domain

type Person struct {
	Name    string `fake:"{firstname} {lastname}"`
	Email   string `fake:"{email}"`
	TaxId   string `fake:"{cpf}"`
	Phone   string `fake:"{phone}"`
	Address string `fake:"{streetname}, {city}, {state}"`
}
