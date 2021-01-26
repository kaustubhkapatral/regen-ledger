package shacl

import "github.com/regen-network/regen-ledger/x/data/rdf"

const (
	ShPrefix rdf.FullIRI = "http://www.w3.org/ns/shacl#"

	ShValidationReport = ShPrefix + "ValidationReport"
	ShValidationResult = ShPrefix + "ValidationResult"
	ShValue            = ShPrefix + "value"

	ShClassConstraintComponent = ShPrefix + "ClassConstraintComponent"
	ShClass                    = ShPrefix + "class"

	ShDatatypeConstraintComponent = ShPrefix + "DatatypeConstraintComponent"
	ShDatatype                    = ShPrefix + "datatype"

	SHNodeKind           = ShPrefix + "nodeKind"
	ShIRI                = ShPrefix + "IRI"
	ShBlankNodeOrIRI     = ShPrefix + "BlankNodeOrIRI"
	ShIRIOrLiteral       = ShPrefix + "IRIOrLiteral"
	ShBlankNode          = ShPrefix + "BlankNode"
	ShBlankNodeOrLiteral = ShPrefix + "BlankNodeOrLiteral"
	ShLiteral            = ShPrefix + "Literal"
)