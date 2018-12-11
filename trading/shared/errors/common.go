package errors

import (
	"fmt"

	"twiggg.com/lk/trading/shared/format"
	langs "twiggg.com/lk/trading/shared/langs/v1.0.0"
)

//NewErrorPrivateResource error
func NewErrorPrivateResource(lang string) error {
	switch lang {
	case langs.FR:
		return fmt.Errorf("accès non autorisé: resource privée")
	case langs.ENG:
		return fmt.Errorf("private data: access not allowed")
	default:
		return fmt.Errorf("private data: access not allowed")
	}
}

//NewErrorInvalidStringLength error
func NewErrorInvalidStringLength(lang string, str string, min, max uint) error {
	switch lang {
	case langs.FR:
		return fmt.Errorf("chaine de caractères '%s' invalide: doit compter entre %d et %d caractères", str, min, max)
	case langs.ENG:
		return fmt.Errorf("invalid string '%s': must count %d to %d caracters", str, min, max)
	default:
		return fmt.Errorf("invalid string '%s': must count %d to %d caracters", str, min, max)
	}
}

//NewErrorNilPointer error
func NewErrorNilPointer(lang string) error {
	switch lang {
	case langs.FR:
		return fmt.Errorf("pointeur nul")
	case langs.ENG:
		return fmt.Errorf("nil pointer")
	default:
		return fmt.Errorf("nil pointer")
	}
}

//NewErrorStrConvInt error
func NewErrorStrConvInt(lang string, str string, err error) error {
	switch lang {
	case langs.FR:
		return fmt.Errorf("echec de la conversion en entier de '%s': %s", str, err.Error())
	case langs.ENG:
		return fmt.Errorf("failed to parse '%s' as an integer: %s", str, err.Error())
	default:
		return fmt.Errorf("failed to parse '%s' as an integer: %s", str, err.Error())
	}
}

//NewErrorMissingData error
func NewErrorMissingData(lang string, missing []string) error {
	//lang = langs.Validate(lang)
	conc := format.ConcatenateStrings(missing, "[", "]", ", ")
	switch lang {
	case langs.FR:
		return fmt.Errorf("données manquantes: %s", conc)
	case langs.ENG:
		return fmt.Errorf("missing data: %s", conc)
	default:
		return fmt.Errorf("missing data: %s", conc)
	}
}

//NewErrorMissing error
func NewErrorMissing(lang string) error {
	switch lang {
	case langs.FR:
		return fmt.Errorf("donnée manquante")
	case langs.ENG:
		return fmt.Errorf("missing data")
	default:
		return fmt.Errorf("missing data")
	}
}

//NewErrorInvalidEmail error
func NewErrorInvalidEmail(lang string, email string) error {
	//lang = langs.Validate(lang)
	switch lang {
	case langs.FR:
		return fmt.Errorf("adresse email '%s' invalide", email)
	case langs.ENG:
		return fmt.Errorf("invalid email address '%s'", email)
	default:
		return fmt.Errorf("invalid email address '%s'", email)
	}
}

//NewErrorNotUpdatable error
func NewErrorNotUpdatable(lang string) error {
	switch lang {
	case langs.FR:
		return fmt.Errorf("mise à jour non autorisée, valeur verrouillée")
	case langs.ENG:
		return fmt.Errorf("update forbidden")
	default:
		return fmt.Errorf("update forbidden")
	}
}

//NewErrorMustBeAuthenticated error
func NewErrorMustBeAuthenticated(lang string) error {
	switch lang {
	case langs.FR:
		return fmt.Errorf("authentification requise pour accéder à cette ressource")
	case langs.ENG:
		return fmt.Errorf("authentication required to access this resource")
	default:
		return fmt.Errorf("authentication required to access this resource")
	}
}

//NewErrorInvalidQuery error
func NewErrorInvalidQuery(lang string) error {
	//lang = langs.Validate(lang)
	switch lang {
	case langs.FR:
		return fmt.Errorf("requête base de donnée invalide")
	case langs.ENG:
		return fmt.Errorf("invalid database query")
	default:
		return fmt.Errorf("invalid database query")
	}
}

/*
//NewErrorInvalidEmail error
func NewErrorInvalidEmail(lang string, email string) error {
	switch lang {
	case langs.FR:
		return fmt.Errorf("adresse email '%s' invalide", email)
	case langs.ENG:
		return fmt.Errorf("invalid email address '%s'", email)
	default:
		return fmt.Errorf("invalid email address '%s'", email)
	}
}
*/

//NewErrorInternal error
func NewErrorInternal(lang string, err error) error {
	//lang = langs.Validate(lang)
	switch lang {
	case langs.FR:
		return fmt.Errorf("une erreur interne s'est produite: '%s'", err.Error())
	case langs.ENG:
		return fmt.Errorf("internal server error: '%s'", err.Error())
	default:
		return fmt.Errorf("internal server error: '%s'", err.Error())
	}
}

//NewErrorNotFound error
func NewErrorNotFound(lang string) error {
	switch lang {
	case langs.FR:
		return fmt.Errorf("aucune resource correspondant à cet identifiant")
	case langs.ENG:
		return fmt.Errorf("no resource matching this ID")
	default:
		return fmt.Errorf("no resource matching this ID")
	}
}

//NewErrorGcdContext error
func NewErrorGcdContext(lang string, err error) error {
	//lang = langs.Validate(lang)
	switch lang {
	case langs.FR:
		return fmt.Errorf("erreur interne: impossible de créer le context: '%s'", err.Error())
	case langs.ENG:
		return fmt.Errorf("internal error: failed to create context: '%s'", err.Error())
	default:
		return fmt.Errorf("internal error: failed to create context: '%s'", err.Error())
	}
}

//NewErrorNoSuchEntity error
func NewErrorNoSuchEntity(lang string, keyString string) error {
	//lang = langs.Validate(lang)
	switch lang {
	case langs.FR:
		return fmt.Errorf("aucune entité correspondant à l'identifiant '%s'", keyString)
	case langs.ENG:
		return fmt.Errorf("no entity matching key '%s'", keyString)
	default:
		return fmt.Errorf("no entity matching key '%s'", keyString)
	}
}

//NewErrorWrongNumberOfKeys error
func NewErrorWrongNumberOfKeys(lang string, count, expected int) error {
	//lang = langs.Validate(lang)
	switch lang {
	case langs.FR:
		return fmt.Errorf("%d clés au lieu de %d attendues", count, expected)
	case langs.ENG:
		return fmt.Errorf("%d keys instead of %d expected", count, expected)
	default:
		return fmt.Errorf("%d keys instead of %d expected", count, expected)
	}
}

//NewErrorInvalidKeyID error
func NewErrorInvalidKeyID(lang string, keyID string) error {
	//lang = langs.Validate(lang)
	switch lang {
	case langs.FR:
		return fmt.Errorf("clé invalide: '%s'", keyID)
	case langs.ENG:
		return fmt.Errorf("invalid key: '%s'", keyID)
	default:
		return fmt.Errorf("invalid key: '%s'", keyID)
	}
}

//NewErrorInvalidIntID error
func NewErrorInvalidIntID(lang string, intID int64) error {
	//lang = langs.Validate(lang)
	switch lang {
	case langs.FR:
		return fmt.Errorf("identifiant invalide: '%d'", intID)
	case langs.ENG:
		return fmt.Errorf("invalid id: '%d'", intID)
	default:
		return fmt.Errorf("invalid id: '%d'", intID)
	}
}

//NewErrorInvalidID error
func NewErrorInvalidID(lang string, id string) error {
	//lang = langs.Validate(lang)
	switch lang {
	case langs.FR:
		return fmt.Errorf("identifiant invalide: '%s'", id)
	case langs.ENG:
		return fmt.Errorf("invalid id: '%s'", id)
	default:
		return fmt.Errorf("invalid id: '%s'", id)
	}
}

//NewErrorNotStringID error
func NewErrorNotStringID(lang string) error {
	//lang = langs.Validate(lang)
	switch lang {
	case langs.FR:
		return fmt.Errorf("la clé n'est pas une clé de type StringID")
	case langs.ENG:
		return fmt.Errorf("key type is not StringID")
	default:
		return fmt.Errorf("key type is not StringID")
	}
}

//NewErrorNotIntID error
func NewErrorNotIntID(lang string) error {
	//lang = langs.Validate(lang)
	switch lang {
	case langs.FR:
		return fmt.Errorf("la clé n'est pas de type IntID")
	case langs.ENG:
		return fmt.Errorf("key type is not IntID")
	default:
		return fmt.Errorf("key type is not IntID")
	}
}

//NewErrorInvalidJSONRequestBody error
func NewErrorInvalidJSONRequestBody(lang string) error {
	//lang = langs.Validate(lang)
	switch lang {
	case langs.FR:
		return fmt.Errorf("corps de requête invalide")
	case langs.ENG:
		return fmt.Errorf("invalid request body")
	default:
		return fmt.Errorf("invalid request body")
	}
}
