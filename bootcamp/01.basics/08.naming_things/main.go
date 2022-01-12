// Why naming is important? : You can't read it, you can't handle it!
// Critical for Readability = Maintainability

// Use the first few letters of the words
var fl string // flag

// Use fewer letters in smaller scopes
var bytesRead int // number of bytes read > NO
var n int // number of bytes read > YES

// Use the complete words in larger scopes
package file
var fileClosed bool

// Use mixedCaps like this
type PlayerScore struct

// Use all captials for acronyms
var localApi string // NO
var localAPI string // YES

// Do not stutter
player.Score // NO
player.PlayerScore // YES

// Do not use under_scores oR LIKE_THIS
const MAX_TIME int // NO
const MaxTime int // YES
const n int // YES