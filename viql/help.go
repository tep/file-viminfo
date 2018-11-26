package viql

import (
	"fmt"
	"io"
)

const qlHelpText = `
VIQL - The VimInfo Query Language

Viql is a simple boolean query language for selecting VimInfo records as
created by examining a Vim swapfile with 'toolman.org/file/viminfo'. Query
statements are boolean expressions composed of declarations or comparisons
combined with boolean operators -- for example:

    cryptmethod = plaintext and not(running or missing)

...and are evaluated against VimInfo structures to derive a boolean value.

In the above, 'cryptmethod = plaintext' is a comparison  while 'running'
and 'missing' are declarations. Therefore, the above statement is true
for VimInfo structs describing a plaintext file whose edit session is no
longer running and it's associated file still exists. Otherwise, it
returns false. See GRAMMAR below for the formal grammar specification.


DECLARATIONS

A declaration is a single token that evaluates to true or false in the
context of a specific VinInfo structure. The following declarations are
currently supported:

    all       -  true for all VinInfo

    none      -  false for all VimInfo

    changed   -  true if the edit session has unsaved changes
    modified

    currhost  -  true if the edit session is on the current host
    thishost

    thisuser  -  true of the edit session was initiated by
    curruser     the current user

    running   -  true if the edit session is currently running
                 (implies "thishost")

    missing   -  true if the file associated with this edit session
                 no longer exists


COMPARISONS

A comparison is used to test the value of a particular VimInfo field.
All comparisons are either for equality or a regular expression match;
there are no inequality comparisons. Regular expression matching is
only available for string fields and other fields may have a limited
set of allowed values. See FIELDS AND VALUES below for details.

Note that boolean fields are not supported via comparisons. Supported
boolean tests are available as declarations.


COMPARITORS

The following comparitors are allowed:

    =   Is equal to
    ==  (note that '=' and '==' mean the same thing).

    !=  Is not equal to

    =~  A regular expression match (value must be a regex)

    !~  A negative regular expression match


FIELDS AND VALUES

Certain VimInfo fields are only comparable against a particular type of
data or a limited set of values. For example, only string fields may be
use with regular expression matches and the PID and Inode fields must be
compared against an integer value.

String values containing space must be double quoted otherwise no quoting
is required. Integer values and values for one of the enumerated fields
(e.g. 'cryptmethod' or 'fileformat') should not be quoted.

The following list enumerates all supported fields and their associated
value restrictions, if any:

    cm           - One of: plaintext, zip, blowfish or blowfish2
    cryptmethod

    filename     - A string field; may be compared using a regexp

    format       - One of: unix, dos or mac
    fileformat

    host         - A string field; may be compared using a regexp
    hostname

    inode        - Integer field; value must be a number

    pid          - Integer field; value must be a number

    user         - A string field; may be compared using a regexp


REGULAR EXPRESSIONS

The supported regular expression syntax is that accepted by RE2 as
implemented by the standard Go "regexp" package and is specified in
a manner inspired by Perl's "qr//" operator as described below.

Regular expression values are surrounded by a pair of '/' characters
and are not anchored; if you'd like your regexp to be anchored to one
end or the other you must use '^' and '$' accordingly.  If a '/' is to
be included in the regular expression, it must be escaped by preceding
it with a backslash '\' character. Likewise, a literal backslash may
also be included by specifying two of them together (e.g. '\\').
Otherwise, no escaping is needed for backslash characters intended as
regexp meta characters.

Consider the following regular expression value:

    /foo\/bar\\..*\.bla/

Here "\/" forstalls the end of the regexp and matches a literal "/", "\\"
matches a literal backslash and does not escape the following ".", but "\."
is escaped and matches a literal "." character.

Regular expression options may be specified with one or more of the
following characters immediately following the closing '/' character:

    i  case-insensitve match

		m  multi-line mode: ^ and $ match begin/end line in addition to
       begin/end text

    s  let . match \n

		U  ungreedy: swap meaning of x* and x*?, x+ and x+?, etc

Capture groups may also be used to limit the scope of certain options to a
subset of the regular expression such as:

    xxx(?flags:re)xxx

...beyond that however, capture groups are ignored.

GRAMMAR

The formal grammer for the language is specified by the following modified
BNF:

     expression  ::=  parenthetic | operation | declaration | comparison

    parenthetic  ::=  '(' expression ')'

      operation  ::=  'not' expression
                   |  expression 'and' expression
                   |  expression 'or'  expression

    declaration  ::=  'all'     | 'changed'  | 'currhost' | 'curruser'
                   | 'missing'  | 'modified' | 'none'     | 'running'
                   | 'thishost' | 'thisuser'

     comparison  ::=  field comparitor value

          field  ::=  'cm'     | 'cryptmethod' | 'filename' | 'fileformat'
                   |  'format' | 'host'        | 'hostname' | 'inode'
                   |  'pid'    | 'user'

     comparitor  ::=  '=' | '==' | '!=' | '=~' | '!~'

          value  ::=  BAREWORD | QUOTED | INTEGER | REGEXP


`

func Help(w io.Writer) {
	fmt.Fprint(w, qlHelpText)
}
