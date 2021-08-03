# pwdgen

A small utility for generating secure passwords.

```
Usage: pwdgen [ OPTIONS ]

Options supported by pwdgen:
      --base58         Use base-58 safe characters
  -n, --count int      Total generated passwords (default 1)
  -D, --digits ints    Numeric character limits [min, max]
      --fast           use cytro-seeded PRNG. (not recommended!)
  -l, --length int     Password length in characters (default 16)
  -L, --lowers ints    Lower-case character limits [min, max]
  -S, --symbols ints   Symbolic character limits [min, max] (default [1,2])
  -U, --uppers ints    Upper-case character limits [min, max]
      --word-safe      use word-safe symbols
```

## Recipes

Generate 100 passwords:

```bash
❯ pwdgen -n100
```

Generate using word-safe symbols/punctuation (`/-+\\~_.`):

```bash
❯ pwdgen --word-safe
```

Generate a password with no symbols or digits:

```bash
❯ pwdgen -S0 -D0
```

## Notes

The `--fast` argument doesn't actually _go any faster_ on most platforms, using pseudo-random Source from `math/rand`; in other words, _don't use it_.
