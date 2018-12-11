package errors

import (
	"fmt"

	"twiggg.com/lk/trading/shared/constants"
	langs "twiggg.com/lk/trading/shared/langs/v1.0.0"
)

//lang codes
const (
	langFR  = 0
	langENG = 1
)

//error constants
const (
	ErrEmptySerie = iota
	ErrInvalidIndex
	ErrNotEnoughPoints
	ErrMaxPeriod
	ErrFailedSelectByName
	ErrInvalidSelection
	ErrMustSelectAnOutput
	ErrNoData
	ErrOnlyOneOutput
	ErrNotSelectable
	ErrInvalidFuncName
	ErrArgsNotFound
	ErrOutputCategNotFound
	ErrInvalidFuncArgs
	ErrInvalidMACDPeriods
	ErrInvalidSignalPeriod
	ErrTimeMismatch
	ErrInvalidFuncCallString
	ErrInvalidSequenceString
	ErrInvalidArgsPeriod
	ErrInvalidArgsIndPeriod
	ErrInvalidArgsWindowBollinger
	ErrInvalidArgsBollingerBands
	ErrInvalidArgsWindowMACD
	ErrInvalidArgsMACD
	ErrCanNotApplyFilterOnValue
	ErrNilPointerPoint
	ErrNilPointerBollingerPoint
	ErrNilPointerMACDPoint
	ErrNilPointerStatsPoint
	ErrNilPointer
	ErrNilSelectableSerie
	ErrNilMultiValue
	ErrNilValue
	ErrNilPointerSerie
	ErrNoCommonPoints
	ErrNotComparable
	ErrTimestampMismatch
	ErrNoDescrForFuncName
	ErrInvalidArgType
	ErrInvalidArgName
	ErrInvalidOutputType
	ErrInvalidOutputName
	ErrInvalidOutputCateg
	ErrEmptyLengthsMap
	ErrInvalidInitialLength
	ErrFastGreaterThanSlow
	ErrPeriodIsZero
	ErrNilDataPointer
	ErrEmptySequence
	ErrOutputsMismatch
	ErrMissingOutput
	ErrCanNotGetItems
	ErrZeroItems
	ErrFailedToGetOutputByName
	ErrSequenceOutputNotValueOrSerie
	ErrCanNotDivideByZero
	ErrInvalidPatternString
	ErrInvalidPatternCode
	ErrInvalidPatternSize
	ErrInvalidMovePct
	ErrNotEnoughPointsForPattern
	ErrInvalidSegmentInds
	ErrCanNotAddUpEmptySerie
	ErrIncompatibleSize
	ErrInvalidArgsWindowTopsBottoms
	ErrInvalidArgsTopsBottoms
	ErrInvalidConditionOperator
	ErrInvalidRefValue
	ErrInsufficientAssets
	ErrInsufficientFunds
	ErrTooManyBuyOrs
	ErrTooManyBuyAnds
	ErrTooManySellOrs
	ErrTooManySellAnds
	ErrTooManySequences
	ErrTooManyFilterLevels
	ErrStrategySeqNameMismatch
	ErrInvalidOperatorAdvanced
	ErrMissingSequenceDef
	ErrInvalidCondition
	ErrInvalidSuccessRate
	ErrInvalidConditionWidth
	ErrInvalidAdvancedComplex
	ErrMissingOperator
	ErrDuplicateSequenceName
	ErrIncompleteCompactStrategy
	ErrIncompleteStrategy
	ErrMustSelectInSeq
	ErrCantSelectInSeq
	ErrMissingDataForSequence
	ErrFailedToComputeSequences
	ErrNilStrategy
	ErrMissingDatesForBacktest
	ErrReservedSequenceName
	ErrEmptySequenceName
	//ErrIncompleteStrategy
	ErrNoOverlap
)

//Err holds the error definition and implements the error interface
type Err struct {
	Code int
	IndI int
	//IndJ int
	Period   int
	Lang     string
	ToSelect string
	FuncName string
	FuncArgs string
}

func (e Err) Error() string {
	return GetErrorMsg(e.Code, e.IndI, e.Period, e.ToSelect, e.FuncName, e.FuncArgs, e.Lang)
}

//GetErrorMsg generates the error message from the code, indices in a specific language
func GetErrorMsg(code int, i int, period int, toSelect string, funcName string, args string, lang string) string {
	switch code {
	case ErrEmptySerie:
		switch lang {
		case langs.FR:
			return "la série est vide"
		case langs.ENG:
			return "empty serie"
		default:
			return "empty serie"
		}
	case ErrInvalidIndex:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("indice '%d' invalide", i)
		case langs.ENG:
			return fmt.Sprintf("invalid index '%d'", i)
		default:
			return fmt.Sprintf("invalid index '%d'", i)
		}
	case ErrNotEnoughPoints:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("nombre de points insuffisant (<%d)", period)
		case langs.ENG:
			return fmt.Sprintf("not enough points (<%d)", period)
		default:
			return fmt.Sprintf("not enough points (<%d)", period)
		}
	case ErrMaxPeriod:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("la période doit être <= %d", constants.MaxPeriod)
		case langs.ENG:
			return fmt.Sprintf("period must be <= %d", constants.MaxPeriod)
		default:
			return fmt.Sprintf("period must be <= %d", constants.MaxPeriod)
		}
	case ErrFailedSelectByName:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("impossible de sélectionner '%s', nom invalide: valeurs permises %s", toSelect, args)
		case langs.ENG:
			return fmt.Sprintf("failed to select '%s' by name, invalid name: expected values %s", toSelect, args)
		default:
			return fmt.Sprintf("failed to select '%s' by name, invalid name: expected values %s", toSelect, args)
		}
	case ErrInvalidSelection:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("impossible de sélectionner '%s' depuis la source '%s', nom invalide", toSelect, args)
		case langs.ENG:
			return fmt.Sprintf("failed to select '%s' from source '%s', invalid name", toSelect, args)
		default:
			return fmt.Sprintf("failed to select '%s' from source '%s', invalid name", toSelect, args)
		}
	case ErrMustSelectAnOutput:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("veuillez sélectionner l'une des sorties de traitement")
		case langs.ENG:
			return fmt.Sprintf("please select an output by name")
		default:
			return fmt.Sprintf("please select an output by name")
		}
	case ErrNoData:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("aucune donnée")
		case langs.ENG:
			return fmt.Sprintf("no data")
		default:
			return fmt.Sprintf("no data")
		}
	case ErrOnlyOneOutput:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("donnée simple (série ou point), inutile de sélectionner")
		case langs.ENG:
			return fmt.Sprintf("simple data (serie or point), no need to select")
		default:
			return fmt.Sprintf("simple data (serie or point), no need to select")
		}
	case ErrNotSelectable:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("impossible de sélectionner une donnée sur un simple point")
		case langs.ENG:
			return fmt.Sprintf("unable to select item on a simple point")
		default:
			return fmt.Sprintf("unable to select item on a simple point")
		}
	case ErrInvalidFuncName:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("fonction '%s' invalide", funcName)
		case langs.ENG:
			return fmt.Sprintf("invalid function call '%s'", funcName)
		default:
			return fmt.Sprintf("invalid function call '%s'", funcName)
		}
	case ErrArgsNotFound:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("aucun argument trouvé pour la fonction '%s'", funcName)
		case langs.ENG:
			return fmt.Sprintf("no arguments found for func '%s'", funcName)
		default:
			return fmt.Sprintf("no arguments found for func '%s'", funcName)
		}
	case ErrOutputCategNotFound:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("aucune sortie trouvée pour la fonction '%s'", funcName)
		case langs.ENG:
			return fmt.Sprintf("no output found for func '%s'", funcName)
		default:
			return fmt.Sprintf("no output found for func '%s'", funcName)
		}
	case ErrReservedSequenceName:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("le nom de séquence '%s' est réservé", funcName)
		case langs.ENG:
			return fmt.Sprintf("sequence name '%s' is reserved", funcName)
		default:
			return fmt.Sprintf("sequence name '%s' is reserved", funcName)
		}
	case ErrEmptySequenceName:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("définir un nom pour la sequence '%s'", funcName)
		case langs.ENG:
			return fmt.Sprintf("set a name for the sequence '%s'", funcName)
		default:
			return fmt.Sprintf("set a name for the sequence '%s'", funcName)
		}
	case ErrInvalidFuncArgs:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("argument(s) '%s' invalide(s) pour la fonction '%s'", args, funcName)
		case langs.ENG:
			return fmt.Sprintf("invalid argument(s) '%s' for function call '%s'", args, funcName)
		default:
			return fmt.Sprintf("invalid argument(s) '%s' for function call '%s'", args, funcName)
		}
	case ErrInvalidMACDPeriods:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("les périodes rapide (fast) et lente (slow) du MACD doivent vérifier fast<=slow-2 et fast,slow>=2")
		case langs.ENG:
			return fmt.Sprintf("fast and slow periods for the MACD must verify fast<slow-2 and fast,slow>=2")
		default:
			return fmt.Sprintf("fast and slow periods for the MACD must verify fast<slow-2 and fast,slow>=2")
		}
	case ErrInvalidSignalPeriod:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("la période de la 'signal line' du MACD doit être >=2")
		case langs.ENG:
			return fmt.Sprintf("MACD signal line period must be >=2")
		default:
			return fmt.Sprintf("MACD signal line period must be >=2")
		}
	case ErrTimeMismatch:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("les dates des deux points à comparer à l'indice %d ne correspondent pas", i)
		case langs.ENG:
			return fmt.Sprintf("dates do not match on two compared points at index %d", i)
		default:
			return fmt.Sprintf("dates do not match on two compared points at index %d", i)
		}
	case ErrInvalidFuncCallString:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("appel de fonction invalide '%s': doit être au format 'nom%sparametres', les paramètres doivent être séparés par des '%s'", funcName, constants.NameAndArgsSeparator, constants.ArgsSeparator)
		case langs.ENG:
			return fmt.Sprintf("invalid function call '%s': must be formated as 'name%sparameters', with parameters separated by '%s'", funcName, constants.NameAndArgsSeparator, constants.ArgsSeparator)
		default:
			return fmt.Sprintf("invalid function call '%s': must be formated as 'name%sparameters', with parameters separated by '%s'", funcName, constants.NameAndArgsSeparator, constants.ArgsSeparator)
		}
	case ErrInvalidSequenceString:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("séquence '%s' invalide: doit être au format 'nom:fonction1%sparametres%sfonction2%sparametres'", funcName, constants.NameAndArgsSeparator, constants.CallsSeparator, constants.NameAndArgsSeparator)
		case langs.ENG:
			return fmt.Sprintf("invalid sequence '%s': should be formated as 'name:function1%sparameters%sfunction2%sparameters'", funcName, constants.NameAndArgsSeparator, constants.CallsSeparator, constants.NameAndArgsSeparator)
		default:
			return fmt.Sprintf("invalid sequence '%s': should be formated as 'name:function1%sparameters%sfunction2%sparameters'", funcName, constants.NameAndArgsSeparator, constants.CallsSeparator, constants.NameAndArgsSeparator)
		}
	case ErrInvalidArgsPeriod:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("paramètre(s) '%s' invalide(s) pour la fonction '%s' -> paramètre attendu: période (entier)", args, funcName)
		case langs.ENG:
			return fmt.Sprintf("invalid parameter(s) '%s' for function '%s' -> expected parameter: period (integer)", args, funcName)
		default:
			return fmt.Sprintf("invalid parameter(s) '%s' for function '%s' -> expected parameter: period (integer)", args, funcName)
		}
	case ErrInvalidArgsIndPeriod:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("paramètre(s) '%s' invalide(s) pour la fonction '%s' -> paramètres attendus: indice,période (entiers)", args, funcName)
		case langs.ENG:
			return fmt.Sprintf("invalid parameter(s) '%s' for function '%s' -> expected parameters: index,period (integers)", args, funcName)
		default:
			return fmt.Sprintf("invalid parameter(s) '%s' for function '%s' -> expected parameters: index,period (integers)", args, funcName)
		}
	case ErrInvalidArgsBollingerBands:
		switch lang {
		case langs.FR:
			//return fmt.Sprintf("paramètre(s) '%s' invalide(s) pour la fonction '%s' -> paramètres attendus: type de moyenne (sma,smma,ema), période (entier) et nb d'écart type (entier positif)", args, funcName)
			return fmt.Sprintf("paramètre(s) '%s' invalide(s) pour la fonction '%s' -> paramètres attendus: type de moyenne (sma), période (entier) et nb d'écart type (entier positif)", args, funcName)
		case langs.ENG:
			//return fmt.Sprintf("invalid parameter(s) '%s' for function '%s' -> expected parameter: average mode (sma,smma,ema), period (integer) and n standard dev (positive integer)", args, funcName)
			return fmt.Sprintf("invalid parameter(s) '%s' for function '%s' -> expected parameter: average mode (sma), period (integer) and n standard dev (positive integer)", args, funcName)
		default:
			//return fmt.Sprintf("invalid parameter(s) '%s' for function '%s' -> expected parameter: average mode (sma,smma,ema), period (integer) and n standard dev (positive integer)", args, funcName)
			return fmt.Sprintf("invalid parameter(s) '%s' for function '%s' -> expected parameter: average mode (sma), period (integer) and n standard dev (positive integer)", args, funcName)
		}
	case ErrInvalidArgsWindowBollinger:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("paramètre(s) '%s' invalide(s) pour la fonction '%s' -> paramètres attendus: indice,période (entiers) et nb d'écart type (entier positif)", args, funcName)
		case langs.ENG:
			return fmt.Sprintf("invalid parameter(s) '%s' for function '%s' -> expected parameters: index,period (integers) and n standard dev (positive integer)", args, funcName)
		default:
			return fmt.Sprintf("invalid parameter(s) '%s' for function '%s' -> expected parameters: index,period (integers) and n standard dev (positive integer)", args, funcName)
		}
	case ErrInvalidArgsWindowTopsBottoms:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("paramètre(s) '%s' invalide(s) pour la fonction '%s' -> paramètres attendus: indice,longueur maxi (entiers) et taille de mouvements en %s (entre %.2f%s et %.2f%s)", args, funcName, "%", constants.MinMovePct, "%", constants.MaxMovePct, "%")
		case langs.ENG:
			return fmt.Sprintf("invalid parameter(s) '%s' for function '%s' -> expected parameters: index,max length (integers) and move size in %s (between %.2f%s and %.2f%s)", args, funcName, "%", constants.MinMovePct, "%", constants.MaxMovePct, "%")
		default:
			return fmt.Sprintf("invalid parameter(s) '%s' for function '%s' -> expected parameters: index,max length (integers) and move size in %s (between %.2f%s and %.2f%s)", args, funcName, "%", constants.MinMovePct, "%", constants.MaxMovePct, "%")
		}
	case ErrInvalidArgsTopsBottoms:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("paramètre(s) '%s' invalide(s) pour la fonction '%s' -> paramètres attendus: longueur maxi (entier) et taille de mouvements en %s (entre %.2f%s et %.2f%s)", args, funcName, "%", constants.MinMovePct, "%", constants.MaxMovePct, "%")
		case langs.ENG:
			return fmt.Sprintf("invalid parameter(s) '%s' for function '%s' -> expected parameters: max length (integer) and move size in %s (between %.2f%s and %.2f%s)", args, funcName, "%", constants.MinMovePct, "%", constants.MaxMovePct, "%")
		default:
			return fmt.Sprintf("invalid parameter(s) '%s' for function '%s' -> expected parameters: max length (integer) and move size in %s (between %.2f%s and %.2f%s)", args, funcName, "%", constants.MinMovePct, "%", constants.MaxMovePct, "%")
		}
	case ErrInvalidArgsMACD:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("paramètres '%s' invalide(s) pour la fonction '%s' -> paramètres attendus: mode (parmis ['sma','smma','ema']), périodes des moyennes rapide et lente (entiers), et période du signal line (entier >=2)", args, funcName)
		case langs.ENG:
			return fmt.Sprintf("invalid parameter(s) '%s' for function '%s' -> expected parameters: mode (in ['sma','ema','smma']), fast and slow moving average periods (integers), signal line period (integer >=2)", args, funcName)
		default:
			return fmt.Sprintf("invalid parameter(s) '%s' for function '%s' -> expected parameters: mode (in ['sma','ema','smma']), fast and slow moving average periods (integers), signal line period (integer >=2)", args, funcName)
		}
	case ErrFastGreaterThanSlow:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("paramètre(s) '%s' invalide(s) pour la fonction '%s' -> la période rapide doit être plus courte que la période lente", args, funcName)
		case langs.ENG:
			return fmt.Sprintf("invalid parameter(s) '%s' for function '%s' -> fast period must be lower than slow period", args, funcName)
		default:
			return fmt.Sprintf("invalid parameter(s) '%s' for function '%s' -> fast period must be lower than slow period", args, funcName)
		}
	case ErrInvalidArgsWindowMACD:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("paramètre(s) '%s' invalides pour la fonction '%s' -> paramètres attendus: mode (parmis ['sma','smma','ema']), indice (entier), périodes des moyennes rapide et lente (entiers), et période du signal line (entier >=2)", args, funcName)
		case langs.ENG:
			return fmt.Sprintf("invalid parameters '%s' for function '%s' -> expected parameters: mode (in ['sma','ema','smma']), index, fast and slow moving average periods (integers), signal line period (integer >=2)", args, funcName)
		default:
			return fmt.Sprintf("invalid parameters '%s' for function '%s' -> expected parameters: mode (in ['sma','ema','smma']), index, fast and slow moving average periods (integers), signal line period (integer >=2)", args, funcName)
		}
	case ErrCanNotApplyFilterOnValue:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("le filtre '%s' peut être appliqué à une valeur simple", funcName)
		case langs.ENG:
			return fmt.Sprintf("can not apply filter '%s' simple value", funcName)
		default:
			return fmt.Sprintf("can not apply filter '%s' simple value", funcName)
		}
	case ErrNilPointer:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("impossible de lire la valeur: pointeur nul")
		case langs.ENG:
			return fmt.Sprintf("unable to read value: nil pointer")
		default:
			return fmt.Sprintf("unable to read value: nil pointer")
		}
	case ErrNilStrategy:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("impossible de lire la stratégie: pointeur nul")
		case langs.ENG:
			return fmt.Sprintf("unable to read the strategy: nil pointer")
		default:
			return fmt.Sprintf("unable to read the strategy: nil pointer")
		}
	case ErrIncompleteStrategy:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("la stratégie doit contenir au moins une condition d'achat, une condition de vente et les définitions des séquences correspondantes")
		case langs.ENG:
			return fmt.Sprintf("the strategy must contain at least one buy condition, one sell condition and the corresponding sequences definitions")
		default:
			return fmt.Sprintf("the strategy must contain at least one buy condition, one sell condition and the corresponding sequences definitions")
		}
	case ErrNilSelectableSerie:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("impossible de lire les données de la source type 'SelectableSerie': pointeur nul")
		case langs.ENG:
			return fmt.Sprintf("unable to read data from source of type 'SelectableSerie': nil pointer")
		default:
			return fmt.Sprintf("unable to read data from source of type 'SelectableSerie': nil pointer")
		}
	case ErrNilMultiValue:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("impossible de lire les données de la source type 'MultiValue': pointeur nul")
		case langs.ENG:
			return fmt.Sprintf("unable to read data from source of type 'MultiValue': nil pointer")
		default:
			return fmt.Sprintf("unable to read data from source of type 'MultiValue': nil pointer")
		}
	case ErrNilValue:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("impossible de lire les données de la source type 'Value': pointeur nul")
		case langs.ENG:
			return fmt.Sprintf("unable to read data from source of type 'Value': nil pointer")
		default:
			return fmt.Sprintf("unable to read data from source of type 'Value': nil pointer")
		}
	case ErrNoOverlap:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("les deux séries '%s' et '%s' ne sont pas comparables: aucune période en commun", funcName, args)
		case langs.ENG:
			return fmt.Sprintf("series '%s' and '%s' can not be compared: no overlap period", funcName, args)
		default:
			return fmt.Sprintf("series '%s' and '%s' can not be compared: no overlap period", funcName, args)
		}
	case ErrNilPointerSerie:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("impossible de lire les données de la source '%s' de type 'Serie': pointeur nul", toSelect)
		case langs.ENG:
			return fmt.Sprintf("unable to read data from source '%s' of type 'Serie': nil pointer", toSelect)
		default:
			return fmt.Sprintf("unable to read data from source '%s' of type 'Serie': nil pointer", toSelect)
		}
	case ErrNoCommonPoints:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("aucun point comparable")
		case langs.ENG:
			return fmt.Sprintf("no comparable point")
		default:
			return fmt.Sprintf("no comparable point")
		}
	case ErrNotComparable:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("comparaison impossible: les séries doivent se correspondre (même longueurs, mêmes périodes, mêmes dates) et chaque point ne doit contenir qu'une propriété/valeur")
		case langs.ENG:
			return fmt.Sprintf("not comparable: series must match (same length, date and time) and each point must only contain one value/field")
		default:
			return fmt.Sprintf("not comparable: series must match (same length, date and time) and each point must only contain one value/field")
		}
	case ErrTimestampMismatch:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("les timestamp à l'indice %d ne correspondent pas", i)
		case langs.ENG:
			return fmt.Sprintf("timestamp mismatch at index %d", i)
		default:
			return fmt.Sprintf("timestamp mismatch at index %d", i)
		}
	case ErrNoDescrForFuncName:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("aucune description pour la fonction '%s'", funcName)
		case langs.ENG:
			return fmt.Sprintf("no description for '%s'", funcName)
		default:
			return fmt.Sprintf("no description for '%s'", funcName)
		}
	case ErrInvalidArgType:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("type de paramètre '%s' invalide", args)
		case langs.ENG:
			return fmt.Sprintf("invalid parameter type '%s'", args)
		default:
			return fmt.Sprintf("invalid parameter type '%s'", args)
		}
	case ErrInvalidArgName:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("nom de paramètre '%s' invalide", args)
		case langs.ENG:
			return fmt.Sprintf("invalid parameter name '%s'", args)
		default:
			return fmt.Sprintf("invalid parameter name '%s'", args)
		}
	case ErrInvalidOutputType:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("type de sortie '%s' invalide", funcName)
		case langs.ENG:
			return fmt.Sprintf("invalid output type '%s'", funcName)
		default:
			return fmt.Sprintf("invalid output type '%s'", funcName)
		}
	case ErrInvalidOutputName:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("nom de sortie '%s' invalide", funcName)
		case langs.ENG:
			return fmt.Sprintf("invalid output name '%s'", funcName)
		default:
			return fmt.Sprintf("invalid output name '%s'", funcName)
		}
	case ErrInvalidOutputCateg:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("catégorie de sortie '%s' invalide", funcName)
		case langs.ENG:
			return fmt.Sprintf("invalid output category '%s'", funcName)
		default:
			return fmt.Sprintf("invalid output category '%s'", funcName)
		}
	case ErrEmptyLengthsMap:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("liste de tailles vide pour la simulation des réductions de séries")
		case langs.ENG:
			return fmt.Sprintf("empty lengths list for series reduction simulation")
		default:
			return fmt.Sprintf("empty lengths list for series reduction simulation")
		}
	case ErrInvalidInitialLength:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("taille d'entrée '%d' invalide pour la simulation des réductions de séries", period)
		case langs.ENG:
			return fmt.Sprintf("invalid input length '%d' for series reduction simulation", period)
		default:
			return fmt.Sprintf("invalid input length '%d' for series reduction simulation", period)
		}
	case ErrPeriodIsZero:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("impossible d'utiliser une période nulle pour la fonction '%s'", funcName)
		case langs.ENG:
			return fmt.Sprintf("unable to use a zero period for function '%s'", funcName)
		default:
			return fmt.Sprintf("unable to use a zero period for function '%s'", funcName)
		}
	case ErrNilDataPointer:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("impossible de lire les données de la source '%s': pointer nul", funcName)
		case langs.ENG:
			return fmt.Sprintf("unable to read data from source '%s': nil pointer", funcName)
		default:
			return fmt.Sprintf("unable to read data from source '%s': nil pointer", funcName)
		}
	case ErrEmptySequence:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("la séquence ne contient aucun appel valide")
		case langs.ENG:
			return fmt.Sprintf("sequence does not contain any valid function call")
		default:
			return fmt.Sprintf("sequence does not contain any valid function call")
		}
	case ErrOutputsMismatch:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("les sorties ne correspondent pas")
		case langs.ENG:
			return fmt.Sprintf("ouputs mismatch")
		default:
			return fmt.Sprintf("outputs mismatch")
		}
	case ErrMissingOutput:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("sortie '%s' manquante pour la fonction '%s'", args, funcName)
		case langs.ENG:
			return fmt.Sprintf("missing output '%s' for function '%s'", args, funcName)
		default:
			return fmt.Sprintf("missing output '%s' for function '%s'", args, funcName)
		}
	case ErrMissingDataForSequence:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("aucune donnée pour la séquence '%s'", funcName)
		case langs.ENG:
			return fmt.Sprintf("no data for sequence '%s'", funcName)
		default:
			return fmt.Sprintf("no data for sequence '%s'", funcName)
		}
	case ErrCanNotGetItems:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("impossible de récupérer la liste des items sélectionnables pour la source '%s'", args)
		case langs.ENG:
			return fmt.Sprintf("can not get the list of selectable items from source '%s'", args)
		default:
			return fmt.Sprintf("can not get the list of selectable items from source '%s'", args)
		}
	case ErrZeroItems:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("aucun item à sélectionner")
		case langs.ENG:
			return fmt.Sprintf("no selectable item")
		default:
			return fmt.Sprintf("no selectable item")
		}
	case ErrSequenceOutputNotValueOrSerie:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("type de sortie '%s' invalide: le dernier filtre de la séquence doit retourner un point ou une série de points", args)
		case langs.ENG:
			return fmt.Sprintf("invalid output type '%s': last filter of the sequence must return a point or a serie of points", args)
		default:
			return fmt.Sprintf("invalid output type '%s': last filter of the sequence must return a point or a serie of points", args)
		}
	case ErrCanNotDivideByZero:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("impossible de diviser par zero")
		case langs.ENG:
			return fmt.Sprintf("can not divide by zero")
		default:
			return fmt.Sprintf("can not divide by zero")
		}
	case ErrInvalidPatternString:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("motif '%s' invalide", args)
		case langs.ENG:
			return fmt.Sprintf("invalid pattern string '%s'", args)
		default:
			return fmt.Sprintf("invalid pattern string '%s'", args)
		}
	case ErrInvalidPatternCode:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("code de motif '%d' invalide", period)
		case langs.ENG:
			return fmt.Sprintf("invalid pattern code '%d'", period)
		default:
			return fmt.Sprintf("invalid pattern code '%d'", period)
		}
	case ErrInvalidPatternSize:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("taille de motif '%d' invalide: doit contenir entre 2 et 4 points", period)
		case langs.ENG:
			return fmt.Sprintf("invalid pattern size '%d': must be between 2 and 4 points", period)
		default:
			return fmt.Sprintf("invalid pattern size '%d': must be between 2 and 4 points", period)
		}
	case ErrInvalidMovePct:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("taille de mouvement %.2f%s invalide: doit être entre %.2f%s et %.2f%s", float64(period)/100.0, "%", constants.MinMovePct, "%", constants.MaxMovePct, "%")
		case langs.ENG:
			return fmt.Sprintf("invalid move size %.2f%s: must be between %.2f%s and %.2f%s", float64(period)/100.0, "%", constants.MinMovePct, "%", constants.MaxMovePct, "%")
		default:
			return fmt.Sprintf("invalid move size %.2f%s: must be between %.2f%s and %.2f%s", float64(period)/100.0, "%", constants.MinMovePct, "%", constants.MaxMovePct, "%")
		}
	case ErrNotEnoughPointsForPattern:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("pas assez de points pour calculer le 'motif à %d point(s)'", period)
		case langs.ENG:
			return fmt.Sprintf("not enough points to compute '%d point(s) pattern'", period)
		default:
			return fmt.Sprintf("not enough points to compute '%d point(s) pattern'", period)
		}
	case ErrInvalidSegmentInds:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("segment invalide: 2 points identiques")
		case langs.ENG:
			return fmt.Sprintf("invalid segment: 2 identical points")
		default:
			return fmt.Sprintf("invalid segment: 2 identical points")
		}
	case ErrCanNotAddUpEmptySerie:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("impossible d'additionner deux séries si l'une est vide")
		case langs.ENG:
			return fmt.Sprintf("unable to add up series if one is empty")
		default:
			return fmt.Sprintf("unable to add up series if one is empty")
		}
	case ErrIncompatibleSize:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("les deux séries ne sont pas de dimensions compatibles")
		case langs.ENG:
			return fmt.Sprintf("series are not of compatible sizes")
		default:
			return fmt.Sprintf("series are not of compatible sizes")
		}
	case ErrInvalidConditionOperator:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("opérateur '%s' invalide, doit être parmis '>,>=,<,<=,='", funcName)
		case langs.ENG:
			return fmt.Sprintf("invalid operator '%s', must be one of '>,>=,<,<=,='", funcName)
		default:
			return fmt.Sprintf("invalid operator '%s', must be one of '>,>=,<,<=,='", funcName)
		}
	case ErrMissingOperator:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("opérateur manquant, doit être parmis '>,>=,<,<=,='")
		case langs.ENG:
			return fmt.Sprintf("missing operator, must be one of '>,>=,<,<=,='")
		default:
			return fmt.Sprintf("missing operator, must be one of '>,>=,<,<=,='")
		}
	case ErrInvalidOperatorAdvanced:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("opérateur invalide dans la condition '%s', doit être parmis '>,>=,<,<=,='", funcName)
		case langs.ENG:
			return fmt.Sprintf("invalid operator in condition '%s', must be one of '>,>=,<,<=,='", funcName)
		default:
			return fmt.Sprintf("invalid operator in condition '%s', must be one of '>,>=,<,<=,='", funcName)
		}
	case ErrMissingSequenceDef:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("définition manquante pour le membre '%s' dans la condition '%s'", funcName, args)
		case langs.ENG:
			return fmt.Sprintf("missing definition for member '%s' in condition '%s'", funcName, args)
		default:
			return fmt.Sprintf("missing definition for member '%s' in condition '%s'", funcName, args)
		}
	case ErrInvalidCondition:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("condition invalide '%s'", funcName)
		case langs.ENG:
			return fmt.Sprintf("invalid condition '%s'", funcName)
		default:
			return fmt.Sprintf("invalid condition '%s'", funcName)
		}
	case ErrInvalidRefValue:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("valeur de référence '%s' invalide", args)
		case langs.ENG:
			return fmt.Sprintf("invalid reference value '%s'", args)
		default:
			return fmt.Sprintf("invalid reference value '%s'", args)
		}
	case ErrInsufficientAssets:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("fonds insuffisants (%s) pour vendre %s", funcName, args)
		case langs.ENG:
			return fmt.Sprintf("insufficient assets (%s) to sell %s", funcName, args)
		default:
			return fmt.Sprintf("insufficient assets (%s) to sell %s", funcName, args)
		}
	case ErrInsufficientFunds:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("fonds insuffisants (%s) pour acheter %s", funcName, args)
		case langs.ENG:
			return fmt.Sprintf("insufficient funds (%s) to buy %s", funcName, args)
		default:
			return fmt.Sprintf("insufficient funds (%s) to buy %s", funcName, args)
		}
	case ErrTooManyBuyOrs:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("nombre de conditions d'achat 'OU' >%d", period)
		case langs.ENG:
			return fmt.Sprintf("number of 'OR' buy conditions >%d", period)
		default:
			return fmt.Sprintf("number of 'OR' buy conditions >%d", period)
		}
	case ErrTooManyBuyAnds:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("nombre de conditions d'achat 'ET' >%d dans le set %d", period, i)
		case langs.ENG:
			return fmt.Sprintf("number of 'AND' buy conditions >%d in set number %d", period, i)
		default:
			return fmt.Sprintf("number of 'AND' buy conditions >%d in set number %d", period, i)
		}
	case ErrTooManySellOrs:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("nombre de conditions de vente 'OU' >%d", period)
		case langs.ENG:
			return fmt.Sprintf("number of 'OR' sell conditions >%d", period)
		default:
			return fmt.Sprintf("number of 'OR' sell conditions >%d", period)
		}
	case ErrTooManySellAnds:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("nombre de conditions de vente 'ET' >%d dans le set %d", period, i)
		case langs.ENG:
			return fmt.Sprintf("number of 'AND' sell conditions >%d in set number %d", period, i)
		default:
			return fmt.Sprintf("number of 'AND' sell conditions >%d in set number %d", period, i)
		}
	case ErrTooManySequences:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("nombre de séquences >%d", period)
		case langs.ENG:
			return fmt.Sprintf("number of sequences >%d", period)
		default:
			return fmt.Sprintf("number of sequences >%d", period)
		}
	case ErrTooManyFilterLevels:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("nombre de filtres >%d dans la séquence '%s'", period, funcName)
		case langs.ENG:
			return fmt.Sprintf("number of filters >%d in sequence '%s'", period, funcName)
		default:
			return fmt.Sprintf("number of filters >%d in sequence '%s'", period, funcName)
		}
	case ErrStrategySeqNameMismatch:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("le nom extrait de la séquence '%s' ne correspond pas au nom attaché '%s'", funcName, args)
		case langs.ENG:
			return fmt.Sprintf("name extracted from sequence '%s' does not match attached name '%s'", funcName, args)
		default:
			return fmt.Sprintf("name extracted from sequence '%s' does not match attached name '%s'", funcName, args)
		}
	case ErrInvalidSuccessRate:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("taux de succès minimal '%s' invalide: doit être un nombre entre 0 et 1.0", funcName)
		case langs.ENG:
			return fmt.Sprintf("invalid min success rate '%s': must be a number between 0 and 1.0", funcName)
		default:
			return fmt.Sprintf("invalid min success rate '%s': must be a number between 0 and 1.0", funcName)
		}
	case ErrInvalidConditionWidth:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("plage de condition '%s points' invalide: doit être un entier entre 1 et 100", funcName)
		case langs.ENG:
			return fmt.Sprintf("invalid condition window '%s points': must be an integer between 1 and 100", funcName)
		default:
			return fmt.Sprintf("invalid condition window '%s points': must be an integer between 1 and 100", funcName)
		}
	case ErrInvalidAdvancedComplex:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("condition complexe invalide '%s': format attendu 'membre1>membre2,taux_succes,taille_plage'", funcName)
		case langs.ENG:
			return fmt.Sprintf("invalid complex condition '%s': expected format 'member1>member2,success_rate,window_size'", funcName)
		default:
			return fmt.Sprintf("invalid complex condition '%s': expected format 'member1>member2,success_rate,window_size'", funcName)
		}
	case ErrDuplicateSequenceName:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("conflit de nom de séquence '%s'", funcName)
		case langs.ENG:
			return fmt.Sprintf("conflicting sequence name '%s'", funcName)
		default:
			return fmt.Sprintf("conflicting sequence name '%s'", funcName)
		}
	case ErrIncompleteCompactStrategy:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("la stratégie doit contenir au moins une séquence, une condition d'achat et une condition de vente")
		case langs.ENG:
			return fmt.Sprintf("strategy must contain at least one sequence, one buy condition and one sell condition")
		default:
			return fmt.Sprintf("strategy must contain at least one sequence, one buy condition and one sell condition")
		}
	/*case ErrIncompleteStrategy:
	switch lang {
	case langs.FR:
		return fmt.Sprintf("la stratégie doit contenir au moins une séquence, une condition d'achat et une condition de vente")
	case langs.ENG:
		return fmt.Sprintf("strategy must contain at least one sequence, one buy condition and one sell condition")
	default:
		return fmt.Sprintf("strategy must contain at least one sequence, one buy condition and one sell condition")
	}*/
	case ErrMustSelectInSeq:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("séquence %d:'%s', sélectionnez un item avant l'appel %d:'%s'", i, funcName, period, args)
		case langs.ENG:
			return fmt.Sprintf("sequence %d:'%s', select an item before function call %d:'%s'", i, funcName, period, args)
		default:
			return fmt.Sprintf("sequence %d:'%s', select an item before function call %d:'%s'", i, funcName, period, args)
		}
	case ErrCantSelectInSeq:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("séquence %d:'%s', impossible de sélectionner l'item '%s'", i, funcName, toSelect)
		case langs.ENG:
			return fmt.Sprintf("sequence %d:'%s', unable to select item '%s'", i, funcName, toSelect)
		default:
			return fmt.Sprintf("sequence %d:'%s', unable to select item '%s'", i, funcName, toSelect)
		}
	case ErrFailedToComputeSequences:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("impossible de calculer les séquences: donnée source nulle ou séquences vides")
		case langs.ENG:
			return fmt.Sprintf("unable to compute séquences: nil input data or empty sequences")
		default:
			return fmt.Sprintf("unable to compute séquences: nil input data or empty sequences")
		}
	case ErrMissingDatesForBacktest:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("dates manquantes pour le backtest")
		case langs.ENG:
			return fmt.Sprintf("missing dates for backtest")
		default:
			return fmt.Sprintf("missing dates for backtest")
		}

	default:
		switch lang {
		case langs.FR:
			return fmt.Sprintf("code d'erreur '%d' à l'indice '%d'", code, i)
		case langs.ENG:
			return fmt.Sprintf("error code '%d' at index '%d'", code, i)
		default:
			return fmt.Sprintf("error code '%d' at index '%d'", code, i)
		}
	}
}

//New returns a custom Err object
func New(code, ind, period int, lang, toSelect, funcName, funcArgs string) Err {
	return Err{Code: code, IndI: ind, Period: period, Lang: lang, ToSelect: toSelect, FuncName: funcName, FuncArgs: funcArgs}
}

//Print prints the inner values
func (e Err) Print() string {
	s := fmt.Sprintf("code: %d , IndI: %d , Period: %d , Lang: '%s' , ToSelect: '%s' , FuncName: '%s', Args: '%s'\n", e.Code, e.IndI, e.Period, e.Lang, e.ToSelect, e.FuncName, e.FuncArgs)
	return s
}
