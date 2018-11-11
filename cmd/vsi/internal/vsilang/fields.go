package vsilang

type field int

const (
	fldNone field = iota
	fldCrypto
	fldFilename
	fldFormat
	fldHost
	fldInode
	fldPID
	fldUser
)

func tokenField(tok int) field {
	switch tok {
	case CRYPTMETHOD:
		return fldCrypto
	case FILEFORMAT:
		return fldFormat
	case FILENAME:
		return fldFilename
	case HOSTNAME:
		return fldHost
	case INODE:
		return fldInode
	case PID:
		return fldPID
	case USER:
		return fldUser
	default:
		return fldNone
	}
}

// A mapping of viminfo fields and the value types to which they're comparable.
var fieldValueTypes = map[field][]valtype{
	fldCrypto:   {vtCryptMethod},
	fldFilename: {vtString, vtRegex},
	fldFormat:   {vtFileFormat},
	fldHost:     {vtString, vtRegex},
	fldInode:    {vtInt},
	fldPID:      {vtInt},
	fldUser:     {vtString, vtRegex},
}

func (f field) comparableTo(v *value) error {
	types, ok := fieldValueTypes[f]
	if !ok {
		return ErrUnknownField
	}

	for _, t := range types {
		if t == v.typ {
			return nil
		}
	}

	return ErrTypeMismatch
}
