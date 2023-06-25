package main

import(
	"fmt"
)

// Struct for professional organizing service
type ProfessionalOrganizingService struct {
	Name           string
	Address        string
	Province       string
	PostalCode     string
	AreaOfExpertise string
}

// Function to create a new Professional Organizing Service
func NewProfessionalOrganizingService(name string, address string, province string, postalCode string, areaOfExpertise string) *ProfessionalOrganizingService {
	return &ProfessionalOrganizingService{
		Name:           name,
		Address:        address,
		Province:       province,
		PostalCode:     postalCode,
		AreaOfExpertise: areaOfExpertise,
	}
}

// Function to get name of Professional Organizing Service
func (s *ProfessionalOrganizingService) GetName() string {
	return s.Name
}

// Function to get address of Professional Organizing Service
func (s *ProfessionalOrganizingService) GetAddress() string {
	return s.Address
}

// Function to get province of Professional Organizing Service
func (s *ProfessionalOrganizingService) GetProvince() string {
	return s.Province
}

// Function to get postal code of Professional Organizing Service
func (s *ProfessionalOrganizingService) GetPostalCode() string {
	return s.PostalCode
}

// Function to get area of expertise of Professional Organizing Service
func (s *ProfessionalOrganizingService) GetAreaOfExpertise() string {
	return s.AreaOfExpertise
}

// Function to set name of Professional Organizing Service
func (s *ProfessionalOrganizingService) SetName(name string) {
	s.Name = name
}

// Function to set address of Professional Organizing Service
func (s *ProfessionalOrganizingService) SetAddress(address string) {
	s.Address = address
}

// Function to set province of Professional Organizing Service
func (s *ProfessionalOrganizingService) SetProvince(province string) {
	s.Province = province
}

// Function to set postal code of Professional Organizing Service
func (s *ProfessionalOrganizingService) SetPostalCode(postalCode string) {
	s.PostalCode = postalCode
}

// Function to set area of expertise of Professional Organizing Service
func (s *ProfessionalOrganizingService) SetAreaOfExpertise(areaOfExpertise string) {
	s.AreaOfExpertise = areaOfExpertise
}

// Function to print details of Professional Organizing Service
func (s *ProfessionalOrganizingService) PrintDetails() {
	fmt.Printf("Name: %s\nAddress: %s\nProvince: %s\nPostal Code: %s\nArea of Expertise: %s\n", s.Name, s.Address, s.Province, s.PostalCode, s.AreaOfExpertise)
}

func main() {
	// Create a new Professional Organizing Service
	professionalOrganizingService := NewProfessionalOrganizingService("Organize it!", "123 Main St", "Ontario", "K1A 0B1", "Home Organization")

	// Print details of Professional Organizing Service
	professionalOrganizingService.PrintDetails()
}