package services

import (
	"rest/database"
	"rest/models"
)

type PackageService struct{}

func InitPackageService() PackageService {
	return PackageService{}
}

func (service *PackageService) GetPackages() ([]models.Packages, error) {
	pkgs := []models.Packages{}

	err := database.DB.Find(&pkgs).Error
	if err != nil {
		return nil, err
	}

	return pkgs, nil
}

func (service *PackageService) GetPackageByID(id string) (models.Packages, error) {
	pkg := models.Packages{}

	err := database.DB.First(&pkg, &id).Error
	if err != nil {
		return models.Packages{}, err
	}

	return pkg, nil
}

func (service *PackageService) AddPackage(body models.Packages) (models.Packages, error) {
	pkg := body

	result := database.DB.Create(&pkg)
	if err := result.Error; err != nil {
		return models.Packages{}, err
	}

	return pkg, nil
}

func (service *PackageService) UpdatePackage(body models.Packages) (models.Packages, error) {
	pkg := body

	result := database.DB.Updates(&pkg)
	if err := result.Error; err != nil {
		return models.Packages{}, err
	}

	return pkg, nil
}

func (service *PackageService) DeletePackage(id string) error {
	pkg := models.Packages{}

	result := database.DB.Delete(&pkg, &id)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}
