package testdata

import "path/filepath"

const (
	DirTestdata = "testdata"
	DirParent   = ".."

	FilenameMFD   = "newsportal.mfd"
	FilenameXML   = "portal.xml"
	FilenameVTXML = "portal.vt.xml"

	PackageDB         = "db"
	PackageVT         = "vt"
	PackageVTTemplate = "vt-template"
	PackageVTUpdated  = "vt-updated"

	PrefixAll    = "all"
	PrefixEntity = "entities"
)

var (
	PathActual   = filepath.Join(DirParent, DirTestdata, "actual")
	PathExpected = filepath.Join(DirParent, DirTestdata, "expected")

	PathActualMFD                = filepath.Join(PathActual, FilenameMFD)
	PathExpectedMFD              = filepath.Join(PathExpected, FilenameMFD)
	PathActualDB                 = filepath.Join(PathActual, PackageDB)
	PathExpectedDB               = filepath.Join(PathExpected, PackageDB)
	PathActualVT                 = filepath.Join(PathActual, PackageVT)
	PathUpdatedVT                = filepath.Join(PathActual, PackageVTUpdated)
	PathExpectedVT               = filepath.Join(PathExpected, PackageVT)
	PathActualVTTemplateAll      = filepath.Join(PathActual, PackageVTTemplate, PrefixAll)
	PathExpectedVTTemplateAll    = filepath.Join(PathExpected, PackageVTTemplate, PrefixAll)
	PathActualVTTemplateEntity   = filepath.Join(PathActual, PackageVTTemplate, PrefixEntity)
	PathExpectedVTTemplateEntity = filepath.Join(PathExpected, PackageVTTemplate, PrefixEntity)
)
