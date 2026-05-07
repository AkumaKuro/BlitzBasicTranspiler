

enum TokenKind {
    IDENT(String),
    INT_LIT(u32),
    FLT_LIT(f32),
    STR_LIT(String),

    PLUS,
    DASH,
    STAR,
    SLASH,

    PERCENT,
    HASH,
    DOLLAR,

    BACKSLASH,
    DOT,

    IF,
    THEN,
    ELSEIF,
    ELSE,
    ENDIF,

    WHILE,
    UNTIL,
    DO,
    WEND,

    FUNCTION,
    ENDFUNCTION,

    END,
}

struct Tokenizer {}