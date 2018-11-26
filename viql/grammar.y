
%{
  package viql

  import "toolman.org/text/scanner"
%}

%union{
  lxr    *lexer
  pos    scanner.Position

  expr   Expression
  oper   *operation
  decl   declaration
  comp   *comparison
  value  *value
  fld    field

  cmp    comparitor
}

%start filter

/* Primitives */
%token ERROR
/* Comparitors */
%token CMP EQ NEQ REM NRE
/* Operators */
%token AND OR NOT
/* Declarations */
%token DECL ALL MISSING MODIFIED NONE RUNNING THISHOST THISUSER
/* Fields */
%token FIELD CRYPTMETHOD FILEFORMAT FILENAME HOSTNAME INODE PID USER
/* File Formats */
%token FFDOS FFMAC FFUNIX
/* Crypt Methods */
%token CMBLOWFISH CMBLOWFISH2 CMPLAINTEXT CMZIP
%token VALUE

%type <expr>   filter expression parenthetic
%type <oper>   operation
%type <decl>   declaration
%type <value>  value
%type <comp>   comparison
%type <cmp>    comparitor
%type <fld>    field

%left AND
%left OR
%right NOT
%right '('

%%

       filter: expression { yyVAL.lxr.expr = $1 }
             ;

   expression: parenthetic
             | operation   { $$ = $1 }
             | declaration { $$ = $1 }
             | comparison  { $$ = $1 }
             ;

  parenthetic: '(' expression ')' { $$ = $2 }
             ;

    operation: NOT expression            { $$ = &operation{ opNot, $2, nil } }
             | expression AND expression { $$ = &operation{ opAnd, $1, $3  } }
             | expression OR  expression { $$ = &operation{ opOr,  $1, $3  } }
             ;

  declaration: DECL { $$ = yyVAL.decl }
             ;

   comparison: field comparitor value {
               c, err := mkComparison($1, $2, $3)
               if err != nil {
                 yyVAL.lxr.err = err
                 return 1
               }
               $$ = c
             }
             ;

        field: FIELD { $$ = yyVAL.fld }
             ;

   comparitor: CMP { $$ = yyVAL.cmp }
             ;

        value: VALUE { $$ = yyVAL.value }
             ;

%%
