package engine

const (
	GetCharFromCongkit = `
	SELECT tc FROM characters LEFT JOIN radicals 
	ON (characters.idx = radicals.char_idx) 
	WHERE radicals.version = ? AND radicals.radical = ?
	`

	GetCharFromQuick = `
	SELECT tc FROM characters LEFT JOIN radicals 
	ON (characters.idx = radicals.char_idx) 
	WHERE radicals.version = ? AND radicals.radical LIKE ?
	`

	GetCharWithPrediction = `
	SELECT tc FROM characters LEFT JOIN radicals 
	ON (characters.idx = radicals.char_idx) 
	WHERE radicals.version = ? AND radicals.radical LIKE ?
	`

	GetSimplifiedCharFromCongkit = `
	SELECT sc FROM characters LEFT JOIN radicals 
	ON (characters.idx = radicals.char_idx) 
	WHERE radicals.version = ? AND radicals.radical = ?
	`

	GetSimplifiedCharFromQuick = `
	SELECT sc FROM characters LEFT JOIN radicals 
	ON (characters.idx = radicals.char_idx) 
	WHERE radicals.version = ? AND radicals.radical LIKE ?
	`

	GetSimplifiedCharWithPrediction = `
	SELECT sc FROM characters LEFT JOIN radicals 
	ON (characters.idx = radicals.char_idx) 
	WHERE radicals.version = ? AND radicals.radical LIKE ?
	`
)
