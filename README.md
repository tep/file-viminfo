
# viminfo
`import "toolman.org/file/viminfo"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package viminfo reads and parses Vim swap files into a well-formed structure.

## Install

``` sh
  go get toolman.org/file/viminfo
```

## <a name="pkg-index">Index</a>
* [type CryptMethod](#CryptMethod)
  * [func (cm CryptMethod) String() string](#CryptMethod.String)
* [type FileFormat](#FileFormat)
  * [func (ff FileFormat) String() string](#FileFormat.String)
* [type VimInfo](#VimInfo)
  * [func Parse(filename string) (*VimInfo, error)](#Parse)


#### <a name="pkg-files">Package files</a>
[block0.go](/src/toolman.org/file/viminfo/block0.go) [cryptmethod.go](/src/toolman.org/file/viminfo/cryptmethod.go) [format.go](/src/toolman.org/file/viminfo/format.go) [viminfo.go](/src/toolman.org/file/viminfo/viminfo.go) 






## <a name="CryptMethod">type</a> [CryptMethod](/src/target/cryptmethod.go?s=55:76#L1)
``` go
type CryptMethod byte
```
CryptMethod is a vim crypto method


``` go
const (
    CMnone      CryptMethod = '0' // No encryption
    CMzip       CryptMethod = 'c' // Default encryption
    CMblowfish  CryptMethod = 'C' // Blowfish encryption
    CMblowfish2 CryptMethod = 'd' // Blowfish2 encryption
)
```
Supported Crypto Methods










### <a name="CryptMethod.String">func</a> (CryptMethod) [String](/src/target/cryptmethod.go?s=327:364#L4)
``` go
func (cm CryptMethod) String() string
```



## <a name="FileFormat">type</a> [FileFormat](/src/target/format.go?s=74:94#L1)
``` go
type FileFormat byte
```
FileFormat indicates the EOL token for an edited file


``` go
const (
    FFnone FileFormat = iota // Format is not known
    FFunix                   // Unix Format "\n"
    FFdos                    // DOS Format  "\r\n"
    FFmac                    // MAC Format  "\r"
)
```
Supported formats










### <a name="FileFormat.String">func</a> (FileFormat) [String](/src/target/format.go?s=317:353#L4)
``` go
func (ff FileFormat) String() string
```



## <a name="VimInfo">type</a> [VimInfo](/src/target/viminfo.go?s=269:1405#L4)
``` go
type VimInfo struct {
    // Version indicates which version of Vim wrote this swap file
    Version string
    // LastMod is the modification time for the file being edited
    LastMod time.Time
    // Inode is the filesystem inode of the file being edited
    Inode uint32
    // PID is the process ID for the vim session editing the file
    PID uint32
    // User is the username for the vim session's process owner (or, UID of username is unavailable)
    User string
    // Hostname is the hostname where the vim session is/was running
    Hostname string
    // Filename reflects the name of the file being edited
    Filename string
    // Encoding is the file encoding for the file being edited (or, the word "encrypted" if the file is encrypted)
    Encoding string
    // Crypto indicates the "cryptmethod" for the file being edited (or, "plaintext" if the file is not encrypted)
    Crypto CryptMethod
    // Format is the FileFormat for the edited file (e.g. unix, dos, mac)
    Format FileFormat
    // Modified indicates whether the edit session has unsaved changes
    Modified bool
    // SameDir indicates whether the edited file is in the same directory as the swap file
    SameDir bool
}
```
VimInfo reflects the meta-data stored in a vim swapfile







### <a name="Parse">func</a> [Parse](/src/target/viminfo.go?s=1554:1599#L33)
``` go
func Parse(filename string) (*VimInfo, error)
```
Parse reads and parses the vim swapfile specified by filename and returns
a populated *VimInfo, or an error if the file could not be parsed.




